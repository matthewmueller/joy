package errors

import "github.com/matthewmueller/golly/js"

// New error from text
func New(text string) error {
	js.Rewrite("new Error($1)", text)
	return nil
}
