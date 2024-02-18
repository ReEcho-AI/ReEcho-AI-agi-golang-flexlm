package llmclient

import (
	"context"

	"github.com/zawakin/lightweight-agi/model"
)

type EmbeddingClient interface {
	EmbedText(ctx context.Context, text string) (*model.Embedding, error)
	EmbedTexts(ctx context.Context, text []string) ([]model.Embedding, error)
}

// CompletionClient is an interface that provides a method to complete text.
type CompletionClient interface