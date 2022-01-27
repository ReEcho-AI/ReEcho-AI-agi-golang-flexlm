package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type EvaluationTaskInput struct {
	Objective  Objective  `json:"objective"`
	Task       Task       `json:"task"`
	TaskResult TaskResult `json:"task_result"`
}

type EvaluationTaskOutput struct {
	Score  int    `json:"score"`
	Reason string `json:"reason"`
}

var (
	EvaluationTasksPrompt = &prompt.Prompt{
		Name:        "evaluation tasks",
		Description: `Evaluate the follow