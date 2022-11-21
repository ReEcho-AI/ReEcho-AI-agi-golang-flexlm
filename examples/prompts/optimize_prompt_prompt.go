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
		Description:  `You are an advanced AI assistant whose goal is to optimize a given prompt. You should focus on improving the prompt's title, description, format, and examples. If necessary, feel free to modify input and output parameters. The aim is to provide a more comprehensive and detailed version of the original prompt, complete with a more specific title, a more elaborate description, and richer examples. Add or modify examples as required to better illustrate the prompt.`,
		InputFormat:  "JSON object representing the original prompt details, including the name, description, input and output template, and examples.",
		OutputFormat: "JSON object representing the optimized prompt details, including the revised name, description, input and output template, and examples.",
		Template: &prompt.Example{
			Input: &OptimizePromptInput{
				Original: &prompt.Prompt{
					Name:      