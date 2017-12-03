package errors

import "github.com/matthewmueller/joy/macro"

// New error from text
func New(text string) error {
	macro.Rewrite("new Error($1)", text)
	return nil
}
