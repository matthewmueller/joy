package mediaerror

import "github.com/matthewmueller/golly/js"

// MediaError struct
// js:"MediaError,omit"
type MediaError struct {
}

// Code
func (*MediaError) Code() (code int8) {
	js.Rewrite("$<.Code")
	return code
}

// MsExtendedCode
func (*MediaError) MsExtendedCode() (msExtendedCode int) {
	js.Rewrite("$<.MsExtendedCode")
	return msExtendedCode
}
