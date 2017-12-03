package msmediakeyerror

import "github.com/matthewmueller/joy/macro"

// MSMediaKeyError struct
// js:"MSMediaKeyError,omit"
type MSMediaKeyError struct {
}

// Code prop
// js:"code"
func (*MSMediaKeyError) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

// SystemCode prop
// js:"systemCode"
func (*MSMediaKeyError) SystemCode() (systemCode uint) {
	macro.Rewrite("$_.systemCode")
	return systemCode
}
