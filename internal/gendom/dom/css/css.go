package css

import "github.com/matthewmueller/golly/js"

type CSS struct {
}

func (*CSS) Supports(property string, value string) (b bool) {
	js.Rewrite("$<.supports($1, $2)", property, value)
	return b
}
