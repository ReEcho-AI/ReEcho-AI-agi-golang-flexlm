package prompt

import (
	"context"
	"encoding/json"
	"log"

	"github.com/zawakin/lightweight-agi/llmclient"
)

// PromptRunner is a struct that runs a prompt.
type PromptRunner struct {
	llmClient llmclient.CompletionClient

	verbose bool
}

func NewPromptRunner(llmClient llmclient.CompletionClient, verbose bool) *PromptRunner {
	return &PromptRunner{
		llmClient: llmClient,
		verbose:   verbose,
	}
}

// Run runs a prompt.
// It formats the input, sends it to the completion client, and unmarshals the output.
func (a *PromptRunner) Run(ctx context.Context, prompter Prompter, input Input, out Output) error {
	prompt, err := prompter.Format(input)
	if err != nil {
		return err
	}

	if a.verbose {
		log.Println("--------------------------------")
		log.Printf("\033[33mPrompt:\