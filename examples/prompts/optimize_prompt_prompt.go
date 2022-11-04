package prompts

import "github.com/zawakin/lightweight-agi/prompt"

type OptimizePromptInput struct {
	Original *prompt.Prompt `json:"original_prompt"`
}

type OptimizePromptOutput struct {
	OptimizedPrompt *prompt.Prompt `json:"optimized_prompt"`
}

var (
	OptimizePromptPrompt = &prompt.Prompt{
		Name:         "AI-Powered Prompt Optimizer",
		Description:  `You are an advanced AI assistant whose goal is to optimize a given prom