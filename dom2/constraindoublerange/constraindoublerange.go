package constraindoublerange

type ConstrainDoubleRange struct {
	*doublerange.DoubleRange

	exact *float32
	ideal *float32
}
