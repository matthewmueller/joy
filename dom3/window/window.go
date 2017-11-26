package window

// New window
func New() *Window {
	return &Window{
		EventTarget: eventtarget{},
	}
}

// Window struct
// js:"Window,omit"
type Window struct {
	EventTarget

	Document Document
}
