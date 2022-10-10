package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type ObjectiveRefinementInput struct {
	Objective Objective `json:"objective"`
}

type ObjectiveRefinementOutput struct {
	RefinedObjective Objective `json:"refined_objective"`
}

var (
	ObjectRefinementPrompt = &prompt.Prompt{
		Name:        "refinement of objective",
		Description: `