
package datastore

import (
	"context"

	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/model"
	"github.com/zawakin/lightweight-agi/services"
)

type DataStore struct {
	provider        DataStoreProvider
	embeddingClient llmclient.EmbeddingClient
}

type DataStoreProvider interface {
	Upsert(ctx context.Context, chunks map[model.DocumentID][]model.DocumentChunk, chunkTokenSize *int) ([]model.DocumentID, error)
	Query(ctx context.Context, queries []model.QueryWithEmbedding) ([]model.QueryResult, error)
}

func NewDataStore(provider DataStoreProvider, embeddingClient llmclient.EmbeddingClient) *DataStore {
	return &DataStore{
		provider:        provider,
		embeddingClient: embeddingClient,
	}
}

func (d *DataStore) Upsert(ctx context.Context, documents []model.Document, chunkTokenSize *int) ([]model.DocumentID, error) {
	chunks := make(map[model.DocumentID][]model.DocumentChunk)