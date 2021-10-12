package inmemory

import (
	"context"
	"math"
	"sort"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-agi/model"
)

var _ datastore.DataStoreProvider = (*InMemoryDataStore)(nil)

type InMemory