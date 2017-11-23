package msmediakeyerror

import "github.com/matthewmueller/golly/js"

// js:"MSMediaKeyError,omit"
type MSMediaKeyError struct {
}

// Code
func (*MSMediaKeyError) Code() (code uint8) {
	js.Rewrite("$<.Code")
	return code
}

// SystemCode
func (*MSMediaKeyError) SystemCode() (systemCode uint) {
	js.Rewrite("$<.SystemCode")
	return systemCode
}
