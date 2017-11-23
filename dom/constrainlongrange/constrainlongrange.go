package constrainlongrange

import "github.com/matthewmueller/golly/dom/longrange"

type ConstrainLongRange struct {
	*longrange.LongRange

	exact *int
	ideal *int
}
