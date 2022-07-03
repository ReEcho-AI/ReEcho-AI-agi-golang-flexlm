package prompts

import (
	"fmt"
	"strings"
)

// Objective is the objective of the task.
// It is used to determine the type of task to create.
type Objective string

func (o Objective) String() string {
	return s