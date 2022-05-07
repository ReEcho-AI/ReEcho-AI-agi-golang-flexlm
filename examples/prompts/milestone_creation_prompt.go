package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type MilestoneCreationInput struct {
	Objective Objective `json:"objective"`
}

type MilestoneCreationOutput struct {
	Milestones Milestones `json:"milestones"`
}

var (
	MilestoneCreationPrompt = &prompt.Prompt{
		N