package anthropic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/pzierahn/chatbot_services/llm"
	"strings"
)

func (client *Client) invokeRequest(model string, req *ClaudeRequest) (*ClaudeResponse, error) {
	body, _ := json.Marshal(req)
	result, err := client.bedrock.InvokeModel(context.Background(), &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(model),
		ContentType: aws.String("application/json"),
		Accept:      aws.String("application/json"),
		Body:        body,
	})
	if err != nil {
		return nil, err
	}

	var response ClaudeResponse
	err = json.Unmarshal(result.Body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *Client) Completion(ctx context.Context, req *llm.CompletionRequest) (*llm.CompletionResponse, error) {

	messages, err := transformMessages(req.Messages)
	if err != nil {
		return nil, err
	}

	toolList := transformTools(req.Tools)
	tools := make(map[string]llm.FunctionCall)
	for _, tool := range req.Tools {
		tools[tool.Name] = tool.Call
	}

	request := ClaudeRequest{
		AnthropicVersion: "bedrock-2023-05-31",
		Messages:         messages,
		System:           req.SystemPrompt,
		MaxTokens:        req.MaxTokens,
		TopP:             req.TopP,
		Temperature:      req.Temperature,
		Tools:            toolList,
	}

	response, err := client.invokeRequest(req.Model, &request)
	if err != nil {
		return nil, err
	}

	usage := llm.ModelUsage{
		UserId:       req.UserId,
		Model:        response.Model,
		InputTokens:  response.Usage.InputTokens,
		OutputTokens: response.Usage.OutputTokens,
	}

	loops := 0
	for response.StopReason == ContentTypeToolUse && loops < 6 {
		request.Messages = append(request.Messages, ClaudeMessage{
			Role:    ChatMessageRoleAssistant,
			Content: response.Content,
		})

		for _, message := range response.Content {
			if message.Type == ContentTypeToolUse {
				callTool, ok := tools[message.Name]
				if !ok {
					return nil, fmt.Errorf("unknown tool %s", message.Name)
				}

				result, err := callTool(ctx, message.Input)
				if err != nil {
					return nil, err
				}

				request.Messages = append(request.Messages, ClaudeMessage{
					Role: ChatMessageRoleUser,
					Content: []Content{{
						Type:      ContentTypeToolResult,
						ToolUseId: message.ID,
						Content:   result,
					}},
				})
			}
		}

		response, err = client.invokeRequest(req.Model, &request)
		if err != nil {
			return nil, err
		}

		usage.OutputTokens += response.Usage.OutputTokens
		usage.InputTokens += response.Usage.InputTokens

		loops++
	}

	return &llm.CompletionResponse{
		Message: &llm.Message{
			Role:    llm.RoleAssistant,
			Content: strings.TrimSpace(response.Content[0].Text),
		},
		Usage: usage,
	}, nil
}
