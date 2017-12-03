package mediastreamerror

import "github.com/matthewmueller/joy/js"

// MediaStreamError struct
// js:"MediaStreamError,omit"
type MediaStreamError struct {
}

// ConstraintName prop
// js:"constraintName"
func (*MediaStreamError) ConstraintName() (constraintName string) {
	js.Rewrite("$_.constraintName")
	return constraintName
}

// Message prop
// js:"message"
func (*MediaStreamError) Message() (message string) {
	js.Rewrite("$_.message")
	return message
}

// Name prop
// js:"name"
func (*MediaStreamError) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}
