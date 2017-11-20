package svgnumber

import "github.com/matthewmueller/golly/js"

type SVGNumber struct {
}

func (*SVGNumber) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*SVGNumber) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
