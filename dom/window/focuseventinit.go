package window

type FocusEventInit struct {
	*UIEventInit

	relatedTarget *EventTarget
}
