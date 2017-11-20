package mediaerror

import "github.com/matthewmueller/golly/js"

type MediaError struct {
}

func (*MediaError) GetCode() (code int8) {
	js.Rewrite("$<.code")
	return code
}

func (*MediaError) GetMsExtendedCode() (msExtendedCode int) {
	js.Rewrite("$<.msExtendedCode")
	return msExtendedCode
}
