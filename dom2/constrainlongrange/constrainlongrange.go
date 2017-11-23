package constrainlongrange

type ConstrainLongRange struct {
	*longrange.LongRange

	exact *int
	ideal *int
}
