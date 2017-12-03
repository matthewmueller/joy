package msmediakeyerror

import "github.com/matthewmueller/joy/js"

// MSMediaKeyError struct
// js:"MSMediaKeyError,omit"
type MSMediaKeyError struct {
}

// Code prop
// js:"code"
func (*MSMediaKeyError) Code() (code uint8) {
	js.Rewrite("$_.code")
	return code
}

// SystemCode prop
// js:"systemCode"
func (*MSMediaKeyError) SystemCode() (systemCode uint) {
	js.Rewrite("$_.systemCode")
	return systemCode
}
