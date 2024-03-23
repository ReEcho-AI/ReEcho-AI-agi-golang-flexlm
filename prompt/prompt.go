
package prompt

import (
	"encoding/json"
	"fmt"
)

type (
	Input  any
	Output any
)

// Prompter interface defines a common method for generating formatted prompts.
type Prompter interface {
	Format(input Input) (string, error)
}

var _ Prompter = (*Prompt)(nil)

// Prompt is a struct that contains the information of a prompt.
type Prompt struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	InputFormat  string   `json:"input_format"`
	OutputFormat string   `json:"output_format"`
	Template     *Example `json:"template"`
	Examples     Examples `json:"examples"`
}
