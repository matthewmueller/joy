package constraindoublerange

import "github.com/matthewmueller/golly/dom2/doublerange"

type ConstrainDoubleRange struct {
	*doublerange.DoubleRange

	exact *float32
	ideal *float32
}
