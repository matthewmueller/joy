package randomsource

import "github.com/matthewmueller/golly/js"

type RandomSource struct {
}

func (*RandomSource) GetRandomValues(array []byte) (b []byte) {
	js.Rewrite("$<.getRandomValues($1)", array)
	return b
}
