
package agi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-agi/examples/prompts"
	"github.com/zawakin/lightweight-agi/model"
	"github.com/zawakin/lightweight-agi/prompt"
)

const (
	maxTaskIterationCount = 20
)

var (
	chunkSize = 200
)

type AGIAgent struct {
	runner *prompt.PromptRunner

	dataStore *datastore.DataStore
}

func NewAGIAgent(
	runner *prompt.PromptRunner,
	dataStore *datastore.DataStore,
) *AGIAgent {
	return &AGIAgent{
		runner:    runner,
		dataStore: dataStore,
	}
}

func (a *AGIAgent) RunAGIByObjective(ctx context.Context, objective prompts.Objective) error {
	var objectiveRefinementOutput prompts.ObjectiveRefinementOutput
	err := a.runner.Run(ctx, prompts.ObjectRefinementPrompt, &prompts.ObjectiveRefinementInput{
		Objective: objective,
	}, &objectiveRefinementOutput)
	if err != nil {
		return err
	}

	var milestoneCreationOutput prompts.MilestoneCreationOutput
	err = a.runner.Run(ctx, prompts.MilestoneCreationPrompt, &prompts.MilestoneCreationInput{
		Objective: objective,
	}, &milestoneCreationOutput)
	if err != nil {
		return err
	}

	milestones := milestoneCreationOutput.Milestones

	for _, milestone := range milestones {
		err := a.RunAGIByMilestone(ctx, milestone)
		if err != nil {