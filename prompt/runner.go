package prompt

import (
	"context"
	"encoding/json"
	"log"

	"github.com/zawakin/lightweight-agi/llmclient"
)

// PromptRunner is a struct that runs a prompt.
type PromptRunner struct {
	llmClient llmclient.Completi