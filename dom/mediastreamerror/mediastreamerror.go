package mediastreamerror

import "github.com/matthewmueller/joy/macro"

// MediaStreamError struct
// js:"MediaStreamError,omit"
type MediaStreamError struct {
}

// ConstraintName prop
// js:"constraintName"
func (*MediaStreamError) ConstraintName() (constraintName string) {
	macro.Rewrite("$_.constraintName")
	return constraintName
}

// Message prop
// js:"message"
func (*MediaStreamError) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}

// Name prop
// js:"name"
func (*MediaStreamError) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}
