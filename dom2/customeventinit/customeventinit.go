package customeventinit

type CustomEventInit struct {
	*eventinit.EventInit

	detail *interface{}
}
