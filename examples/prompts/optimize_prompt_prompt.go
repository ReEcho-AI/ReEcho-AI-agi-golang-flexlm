package prompts

import "github.com/zawakin/lightweight-agi/prompt"

type OptimizePromptInput struct {
	Original *prompt.Prompt `json:"original_prompt"`
}

type OptimizePrompt