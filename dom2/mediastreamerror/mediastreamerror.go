package mediastreamerror

import "github.com/matthewmueller/golly/js"

// js:"MediaStreamError,omit"
type MediaStreamError struct {
}

// ConstraintName
func (*MediaStreamError) ConstraintName() (constraintName string) {
	js.Rewrite("$<.ConstraintName")
	return constraintName
}

// Message
func (*MediaStreamError) Message() (message string) {
	js.Rewrite("$<.Message")
	return message
}

// Name
func (*MediaStreamError) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}
