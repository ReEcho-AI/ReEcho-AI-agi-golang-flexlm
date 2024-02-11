package llmclient

import (
	"context"

	"github.com/zawakin/lightweight-agi/model"
)

type EmbeddingClient interface {
	EmbedText(ctx context.Context, text string) (*model.Em