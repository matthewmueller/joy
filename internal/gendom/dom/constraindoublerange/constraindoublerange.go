package constraindoublerange

type ConstrainDoubleRange struct {
	*DoubleRange

	exact *float32
	ideal *float32
}
