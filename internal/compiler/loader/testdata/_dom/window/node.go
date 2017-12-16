package window

type Node interface {
	DispatchEvent() Event

	ParentElement() HTMLElement
}
