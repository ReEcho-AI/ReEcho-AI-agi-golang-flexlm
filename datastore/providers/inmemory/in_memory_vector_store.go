package inmemory

import (
	"context"
	"math"
	"sort"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-agi/model"
)

var _ datastore.DataStoreProvider = (*InMemoryDataStore)(nil)

type InMemoryDataStore struct {
	data map[model.DocumentChunkID]model.DocumentChunk
}

func NewInMemoryDataStore() *InMemoryDataStore {
	return &InMemoryDataStore{
		data: make(map[model.DocumentChunkID]model.DocumentChunk),
	}
}

func (s *InMemoryDataStore) Upsert(ctx context.Context, chunks map[model.DocumentID][]model.DocumentChunk, chunkTokenSize *int) ([]model.DocumentID, error) {
	var result []model.DocumentID
	for docID, v := range chunks {
		for _, chunk := range v {
			s.data[chunk.ID] = chunk
		}
		result = append(result, docID)
	}
	