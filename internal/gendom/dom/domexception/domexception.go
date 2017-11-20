package domexception

import "github.com/matthewmueller/golly/js"

type DOMException struct {
}

func (*DOMException) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*DOMException) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*DOMException) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}

func (*DOMException) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}
