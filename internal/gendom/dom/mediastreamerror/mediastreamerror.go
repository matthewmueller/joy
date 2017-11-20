package mediastreamerror

import "github.com/matthewmueller/golly/js"

type MediaStreamError struct {
}

func (*MediaStreamError) GetConstraintName() (constraintName string) {
	js.Rewrite("$<.constraintName")
	return constraintName
}

func (*MediaStreamError) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}

func (*MediaStreamError) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}
