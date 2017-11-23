package mediastreamerror

import "github.com/matthewmueller/golly/js"

// MediaStreamError struct
// js:"MediaStreamError,omit"
type MediaStreamError struct {
}

// ConstraintName prop
func (*MediaStreamError) ConstraintName() (constraintName string) {
	js.Rewrite("$<.constraintName")
	return constraintName
}

// Message prop
func (*MediaStreamError) Message() (message string) {
	js.Rewrite("$<.message")
	return message
}

// Name prop
func (*MediaStreamError) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}
