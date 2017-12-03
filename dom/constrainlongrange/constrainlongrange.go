package constrainlongrange

import "github.com/matthewmueller/joy/dom/longrange"

type ConstrainLongRange struct {
	*longrange.LongRange

	exact *int
	ideal *int
}
