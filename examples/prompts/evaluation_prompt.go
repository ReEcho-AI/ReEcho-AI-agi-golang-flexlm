package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type EvaluationTaskInput struct {
	Objective  Objective  `json:"objective"`
	Task       Task       `json:"task"`
	TaskResult TaskR