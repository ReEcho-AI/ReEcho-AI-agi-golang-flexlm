package services

import (
	"context"

	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/model"
)

func GetTextChunks(text string, chunkTokenSize int) []string {
	var chunks []string
	for i := 0; i < len(text); i += chunkTokenSize {
		end := i + chunkTokenSize
		if end > len(text) {
			end = len(text)
		}
		chunks = append(chunks, text[i:end])
	}
	return chunks
}

func CreateDocumentChunks(ctx context.Context, embeddingClient llmclient.EmbeddingClient, doc *model.Document, chunkTokenSize int) ([]model.DocumentChunk, error) {
	var chunks []model.DocumentChunk

	embeddings, err := embeddingClient.EmbedTexts(ctx, GetTextChunks(doc.Text, chunkTokenSize))
	if err != nil {
		return