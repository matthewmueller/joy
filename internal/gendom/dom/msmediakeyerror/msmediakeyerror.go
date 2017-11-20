package msmediakeyerror

import "github.com/matthewmueller/golly/js"

type MSMediaKeyError struct {
}

func (*MSMediaKeyError) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*MSMediaKeyError) GetSystemCode() (systemCode uint) {
	js.Rewrite("$<.systemCode")
	return systemCode
}
