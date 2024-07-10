package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pzierahn/chatbot_services/datastore"
	"github.com/pzierahn/chatbot_services/llm"
	"github.com/pzierahn/chatbot_services/search"
	pb "github.com/pzierahn/chatbot_services/services/proto"
	"log"
	"sort"
	"time"
)

type Sources struct {
	Items []*search.Result `json:"sources"`
}

// PostMessage is a gRPC endpoint that receives a prompt and returns a completion.
func (service *Service) PostMessage(ctx context.Context, prompt *pb.Prompt) (*pb.Message, error) {
	log.Printf("PostMessage: %v", prompt)

	userId, err := service.Auth.VerifyFunding(ctx)
	if err != nil {
		return nil, err
	}

	//
	// Integrity check
	//

	collectionId, err := uuid.Parse(prompt.CollectionId)
	if err != nil {
		return nil, errors.New("invalid collection id")
	}

	modelOps := prompt.GetModelOptions()
	if modelOps == nil {
		return nil, fmt.Errorf("options missing")
	}

	retrievalOptions := prompt.GetRetrievalOptions()
	if retrievalOptions == nil {
		return nil, fmt.Errorf("retrieval options missing")
	}

	model, err := service.getModel(modelOps.ModelId)
	if err != nil {
		return nil, err
	}

	//
	// Get the thread messages
	//

	var thread *datastore.Thread
	if prompt.ThreadId != "" {
		//
		// Get the thread from the database
		//

		threadId, err := uuid.Parse(prompt.ThreadId)
		if err != nil {
			return nil, err
		}

		thread, err = service.Database.GetThread(ctx, userId, threadId)
		if err != nil {
			return nil, err
		}
	} else {
		//
		// Create a new thread
		//

		thread = &datastore.Thread{
			Id:           uuid.New(),
			UserId:       userId,
			CollectionId: collectionId,
			Timestamp:    time.Now(),
		}
	}

	//
	// Call the model
	//

	messages := append(thread.Messages, &llm.Message{
		Role:    llm.RoleUser,
		Content: prompt.Prompt,
	})

	request := &llm.CompletionRequest{
		SystemPrompt: "You are a scientific research assistant. Quote sources with \\cite{document_id}.",
		Messages:     messages,
		Model:        modelOps.ModelId,
		MaxTokens:    int(modelOps.MaxTokens),
		TopP:         modelOps.TopP,
		Temperature:  modelOps.Temperature,
		UserId:       userId,
		Tools: []llm.ToolDefinition{{
			Name:        "get_sources",
			Description: "Retrieves the sources for the prompt. The prompt should be optimized for embedding retrieval. The tool will return a list of sources in JSON format with the following fields: SourceID, Content.",
			Parameters: llm.ToolParameters{
				Type: "object",
				Properties: map[string]llm.ParametersProperties{
					"prompt": {
						Type:        "string",
						Description: "The topic for which to retrieve sources. The prompt should be optimized for embedding retrieval.",
					},
				},
				Required: []string{"prompt"},
			},
			Call: func(ctx context.Context, parameters map[string]interface{}) (string, error) {
				query, ok := parameters["prompt"].(string)
				if !ok {
					return "", errors.New("prompt missing")
				}

				log.Printf("get_sources: %v", query)

				response, err := service.Search.Search(ctx, search.Query{
					UserId:       userId,
					CollectionId: prompt.CollectionId,
					Query:        query,
					Limit:        retrievalOptions.Documents,
					Threshold:    retrievalOptions.Threshold,
				})
				if err != nil {
					return "", err
				}

				_ = service.Database.InsertModelUsage(ctx, &datastore.ModelUsage{
					Id:          uuid.New(),
					UserId:      userId,
					Timestamp:   time.Now(),
					ModelId:     response.Usage.ModelId,
					InputTokens: response.Usage.Tokens,
				})

				sources := response.Results

				// Group by document and sort by position
				sort.Slice(sources, func(i, j int) bool {
					if sources[i].DocumentId != sources[j].DocumentId {
						return sources[i].DocumentId < sources[j].DocumentId
					}
					return sources[i].Position < sources[j].Position
				})

				byt, err := json.Marshal(Sources{
					Items: sources,
				})
				if err != nil {
					return "", err
				}

				byt2, _ := json.MarshalIndent(sources, "", "  ")
				log.Printf("get_sources: %s", byt2)

				return string(byt), nil
			},
		}},
	}

	response, err := model.Completion(ctx, request)
	if err != nil {
		return nil, err
	}

	//
	// Save the response
	//

	thread.Messages = response.Messages
	err = service.Database.StoreThread(ctx, thread)
	if err != nil {
		return nil, err
	}

	_ = service.Database.InsertModelUsage(ctx, &datastore.ModelUsage{
		Id:           uuid.New(),
		UserId:       userId,
		Timestamp:    time.Now(),
		ModelId:      modelOps.ModelId,
		InputTokens:  response.Usage.InputTokens,
		OutputTokens: response.Usage.OutputTokens,
	})

	// Get the document names
	sources := getSources(response.Messages)

	for idx, source := range sources {
		docId, err := uuid.Parse(source.DocumentId)
		if err != nil {
			continue
		}

		docName, err := service.Database.GetDocumentName(ctx, userId, docId)
		if err != nil {
			continue
		}
		sources[idx].Name = docName
	}

	return &pb.Message{
		ThreadId:   thread.Id.String(),
		Prompt:     prompt.Prompt,
		Completion: thread.Messages[len(thread.Messages)-1].Content,
		Sources:    sources,
	}, nil
}