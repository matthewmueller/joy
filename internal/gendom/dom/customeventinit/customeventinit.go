package customeventinit

type CustomEventInit struct {
	*EventInit

	detail *interface{}
}
