package vertex

import (
	"cloud.google.com/go/vertexai/genai"
	"context"
	"encoding/json"
	"errors"
	"github.com/pzierahn/chatbot_services/llm"
	"strings"
)

const (
	RoleUser  = "user"
	RoleModel = "model"
)

func (client *Client) Completion(ctx context.Context, req *llm.CompletionRequest) (*llm.CompletionResponse, error) {
	if len(req.Messages) == 0 {
		return nil, errors.New("no messages")
	}

	modelName, _ := strings.CutPrefix(req.Model, modelPrefix)

	outputTokens := int32(req.MaxTokens)

	model := client.client.GenerativeModel(modelName)
	model.TopP = &req.TopP
	model.Temperature = &req.Temperature
	model.MaxOutputTokens = &outputTokens
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(req.SystemPrompt)},
	}
	model.Tools = client.getTools()

	chat := model.StartChat()

	// Transform the messages to a history
	history, err := transformToHistory(req.Messages)
	if err != nil {
		return nil, err
	}

	// Remove the last message from the history, because the last message needs to be sent to the model
	chat.History = history[:len(history)-1]
	gen, err := chat.SendMessage(ctx, history[len(history)-1].Parts...)
	if err != nil {
		return nil, err
	}

	// Check if the model returned a response
	if len(gen.Candidates) == 0 || len(gen.Candidates[0].Content.Parts) == 0 {
		return nil, nil
	}

	usage := llm.ModelUsage{
		UserId: req.UserId,
		Model:  modelName,
	}

	if gen.UsageMetadata != nil {
		usage.InputTokens = int(gen.UsageMetadata.PromptTokenCount)
		usage.OutputTokens = int(gen.UsageMetadata.CandidatesTokenCount)
	}

	if fun, ok := gen.Candidates[0].Content.Parts[0].(genai.FunctionCall); ok {
		// Add the function call to the history
		history = append(history, &genai.Content{
			Role:  RoleModel,
			Parts: []genai.Part{fun},
		})

		// Call the function to get the result
		resultStr, err := client.callTool(ctx, fun.Name, fun.Args)
		if err != nil {
			return nil, err
		}

		// Parse the result
		var results map[string]interface{}
		err = json.Unmarshal([]byte(resultStr), &results)
		if err != nil {
			return nil, err
		}

		functionResults := genai.FunctionResponse{
			Name:     fun.Name,
			Response: results,
		}
		history = append(history, &genai.Content{
			Role:  RoleUser,
			Parts: []genai.Part{functionResults},
		})
		chat.History = history[:len(history)-1]

		gen, err = chat.SendMessage(ctx, functionResults)
		if err != nil {
			return nil, err
		}

		if gen.UsageMetadata != nil {
			usage.InputTokens += int(gen.UsageMetadata.PromptTokenCount)
			usage.OutputTokens += int(gen.UsageMetadata.CandidatesTokenCount)
		}
	}

	txt, ok := gen.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return nil, nil
	}

	return &llm.CompletionResponse{
		Message: &llm.Message{
			Role:    llm.RoleAssistant,
			Content: string(txt),
		},
		Usage: usage,
	}, nil
}
