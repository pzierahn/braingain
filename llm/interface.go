package llm

import "context"

const (
	MessageTypeUser = iota
	MessageTypeBot
)

type Message struct {
	Type int
	Text string
}

type GenerateRequest struct {
	Messages    []*Message
	Model       string
	MaxTokens   int
	TopP        float32
	Temperature float32
	UserId      string
}

type GenerateResponse struct {
	Text string
}

type EmbeddingRequest struct {
	Input        string
	UserId       string
	SkipTracking bool
}

type EmbeddingResponse struct {
	Data   []float32
	Tokens int
}

type ModelUsage struct {
	Model            string
	UserId           string
	PromptTokens     int
	CompletionTokens int
}

type Embedding interface {
	CreateEmbeddings(ctx context.Context, req *EmbeddingRequest) (*EmbeddingResponse, error)
	GetModelName() string
}

type Completion interface {
	GenerateCompletion(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error)
	ProvidesModel(model string) bool
}