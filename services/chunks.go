package services

import (
	"context"

	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/model"
)

func GetTextChunks(text string, chunkTokenSize int) []string {
	var