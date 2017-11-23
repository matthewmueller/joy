package window

import "github.com/matthewmueller/golly/js"

// GlobalEventHandlers struct
// js:"GlobalEventHandlers,omit"
type GlobalEventHandlers struct {
}

// Onpointercancel prop
func (*GlobalEventHandlers) Onpointercancel() (onpointercancel func(Event)) {
	js.Rewrite("$<.onpointercancel")
	return onpointercancel
}

// Onpointercancel prop
func (*GlobalEventHandlers) SetOnpointercancel(onpointercancel func(Event)) {
	js.Rewrite("$<.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
func (*GlobalEventHandlers) Onpointerdown() (onpointerdown func(Event)) {
	js.Rewrite("$<.onpointerdown")
	return onpointerdown
}

// Onpointerdown prop
func (*GlobalEventHandlers) SetOnpointerdown(onpointerdown func(Event)) {
	js.Rewrite("$<.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
func (*GlobalEventHandlers) Onpointerenter() (onpointerenter func(Event)) {
	js.Rewrite("$<.onpointerenter")
	return onpointerenter
}

// Onpointerenter prop
func (*GlobalEventHandlers) SetOnpointerenter(onpointerenter func(Event)) {
	js.Rewrite("$<.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
func (*GlobalEventHandlers) Onpointerleave() (onpointerleave func(Event)) {
	js.Rewrite("$<.onpointerleave")
	return onpointerleave
}

// Onpointerleave prop
func (*GlobalEventHandlers) SetOnpointerleave(onpointerleave func(Event)) {
	js.Rewrite("$<.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
func (*GlobalEventHandlers) Onpointermove() (onpointermove func(Event)) {
	js.Rewrite("$<.onpointermove")
	return onpointermove
}

// Onpointermove prop
func (*GlobalEventHandlers) SetOnpointermove(onpointermove func(Event)) {
	js.Rewrite("$<.onpointermove = $1", onpointermove)
}

// Onpointerout prop
func (*GlobalEventHandlers) Onpointerout() (onpointerout func(Event)) {
	js.Rewrite("$<.onpointerout")
	return onpointerout
}

// Onpointerout prop
func (*GlobalEventHandlers) SetOnpointerout(onpointerout func(Event)) {
	js.Rewrite("$<.onpointerout = $1", onpointerout)
}

// Onpointerover prop
func (*GlobalEventHandlers) Onpointerover() (onpointerover func(Event)) {
	js.Rewrite("$<.onpointerover")
	return onpointerover
}

// Onpointerover prop
func (*GlobalEventHandlers) SetOnpointerover(onpointerover func(Event)) {
	js.Rewrite("$<.onpointerover = $1", onpointerover)
}

// Onpointerup prop
func (*GlobalEventHandlers) Onpointerup() (onpointerup func(Event)) {
	js.Rewrite("$<.onpointerup")
	return onpointerup
}

// Onpointerup prop
func (*GlobalEventHandlers) SetOnpointerup(onpointerup func(Event)) {
	js.Rewrite("$<.onpointerup = $1", onpointerup)
}

// Onwheel prop
func (*GlobalEventHandlers) Onwheel() (onwheel func(Event)) {
	js.Rewrite("$<.onwheel")
	return onwheel
}

// Onwheel prop
func (*GlobalEventHandlers) SetOnwheel(onwheel func(Event)) {
	js.Rewrite("$<.onwheel = $1", onwheel)
}
