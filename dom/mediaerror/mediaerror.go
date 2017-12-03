package mediaerror

import "github.com/matthewmueller/joy/macro"

// MediaError struct
// js:"MediaError,omit"
type MediaError struct {
}

// Code prop
// js:"code"
func (*MediaError) Code() (code int8) {
	macro.Rewrite("$_.code")
	return code
}

// MsExtendedCode prop
// js:"msExtendedCode"
func (*MediaError) MsExtendedCode() (msExtendedCode int) {
	macro.Rewrite("$_.msExtendedCode")
	return msExtendedCode
}
