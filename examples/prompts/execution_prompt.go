
package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type ExecutionInput struct {
	Objective       Objective     `json:"objective"`
	CurrentTask     Task          `json:"current_task"`
	SolvedTasks     Tasks         `json:"solved_tasks"`
	RelevantContext []TaskContext `json:"relevant_context"`
}

type ExecutionOutput struct {
	CurrentTaskResult TaskResult `json:"current_task_result"`
}

var (
	ExecutionPrompt = prompt.NewPrompt(
		"execution",
		`Execute the task (with the given context) to archieve the objective. Output the result of the task to result text.`,
		prompt.NewExample(
			&ExecutionInput{