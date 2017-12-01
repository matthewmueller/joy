package mediaerror

import "github.com/matthewmueller/golly/js"

// MediaError struct
// js:"MediaError,omit"
type MediaError struct {
}

// Code prop
// js:"code"
func (*MediaError) Code() (code int8) {
	js.Rewrite("$_.code")
	return code
}

// MsExtendedCode prop
// js:"msExtendedCode"
func (*MediaError) MsExtendedCode() (msExtendedCode int) {
	js.Rewrite("$_.msExtendedCode")
	return msExtendedCode
}
