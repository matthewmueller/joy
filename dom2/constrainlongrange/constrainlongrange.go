package constrainlongrange

import "github.com/matthewmueller/golly/dom2/longrange"

type ConstrainLongRange struct {
	*longrange.LongRange

	exact *int
	ideal *int
}
