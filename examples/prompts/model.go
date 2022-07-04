package prompts

import (
	"fmt"
	"strings"
)

// Objective is the objective of the task.
// It is used to determine the type of task to create.
type Objective string

func (o Objective) String() string {
	return string(o)
}

type Milestone struct {
	Objective Objective `json:"objective"`
	Name      string    `json:"name"`
}

type Milestones []Milestone

func (ms Milestones) String() string {
	var sb strings.Builder
	for i, m := range ms {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, m.Name))
	}
	return sb.String()
}

var (
	// MaxContextLength is the maximum length of the context that can be
	// returned by the model.
	MaxContextLength = 2000
)

// Task is a struct that contains the base task information.
// It is used to identify the task and to provide a name for the task.
type Task struct {
	Name string `json:"name"`
}

func (t Task) String() string {
	return fmt.Sprintf(`"%s"`, t.Name)
}

type Tasks []Task

func (ts *Tasks) Add(task Task) {
	*ts = append(*ts, task)
}

func (ts Tasks) PopLeft() (Task, Tasks) {
	retu