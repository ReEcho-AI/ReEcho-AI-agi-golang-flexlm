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
		Name: "milestone creation",
		Description: `You are an AI tasked with creating a milestone for the following objective.

Please provide a milestone that can be used to achieve the objective.`,
		Template: &prompt.Example{
			Input: &MilestoneCreationInput{
				Objective: Objective("original objective"),
			},
			Output: &MilestoneCreationOutput{
				Milestones: Milestones{
					{Objective: Objective("milestone1