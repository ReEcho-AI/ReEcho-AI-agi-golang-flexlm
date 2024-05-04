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

fun