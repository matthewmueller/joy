package mediaerror

import "github.com/matthewmueller/golly/js"

// MediaError struct
// js:"MediaError,omit"
type MediaError struct {
}

// Code prop
func (*MediaError) Code() (code int8) {
	js.Rewrite("$<.code")
	return code
}

// MsExtendedCode prop
func (*MediaError) MsExtendedCode() (msExtendedCode int) {
	js.Rewrite("$<.msExtendedCode")
	return msExtendedCode
}
