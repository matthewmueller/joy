package window

type Node interface {
	EventTarget

	ParentElement() HTMLElement
}
