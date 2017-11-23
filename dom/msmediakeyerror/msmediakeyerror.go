package msmediakeyerror

import "github.com/matthewmueller/golly/js"

// MSMediaKeyError struct
// js:"MSMediaKeyError,omit"
type MSMediaKeyError struct {
}

// Code prop
func (*MSMediaKeyError) Code() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

// SystemCode prop
func (*MSMediaKeyError) SystemCode() (systemCode uint) {
	js.Rewrite("$<.systemCode")
	return systemCode
}
