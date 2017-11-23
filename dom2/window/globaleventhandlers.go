package window

import "github.com/matthewmueller/golly/js"

// GlobalEventHandlers struct
// js:"GlobalEventHandlers,omit"
type GlobalEventHandlers struct {
}

// Onpointercancel
func (*GlobalEventHandlers) Onpointercancel() (onpointercancel func(Event)) {
	js.Rewrite("$<.Onpointercancel")
	return onpointercancel
}

// Onpointercancel
func (*GlobalEventHandlers) SetOnpointercancel(onpointercancel func(Event)) {
	js.Rewrite("$<.Onpointercancel = $1", onpointercancel)
}

// Onpointerdown
func (*GlobalEventHandlers) Onpointerdown() (onpointerdown func(Event)) {
	js.Rewrite("$<.Onpointerdown")
	return onpointerdown
}

// Onpointerdown
func (*GlobalEventHandlers) SetOnpointerdown(onpointerdown func(Event)) {
	js.Rewrite("$<.Onpointerdown = $1", onpointerdown)
}

// Onpointerenter
func (*GlobalEventHandlers) Onpointerenter() (onpointerenter func(Event)) {
	js.Rewrite("$<.Onpointerenter")
	return onpointerenter
}

// Onpointerenter
func (*GlobalEventHandlers) SetOnpointerenter(onpointerenter func(Event)) {
	js.Rewrite("$<.Onpointerenter = $1", onpointerenter)
}

// Onpointerleave
func (*GlobalEventHandlers) Onpointerleave() (onpointerleave func(Event)) {
	js.Rewrite("$<.Onpointerleave")
	return onpointerleave
}

// Onpointerleave
func (*GlobalEventHandlers) SetOnpointerleave(onpointerleave func(Event)) {
	js.Rewrite("$<.Onpointerleave = $1", onpointerleave)
}

// Onpointermove
func (*GlobalEventHandlers) Onpointermove() (onpointermove func(Event)) {
	js.Rewrite("$<.Onpointermove")
	return onpointermove
}

// Onpointermove
func (*GlobalEventHandlers) SetOnpointermove(onpointermove func(Event)) {
	js.Rewrite("$<.Onpointermove = $1", onpointermove)
}

// Onpointerout
func (*GlobalEventHandlers) Onpointerout() (onpointerout func(Event)) {
	js.Rewrite("$<.Onpointerout")
	return onpointerout
}

// Onpointerout
func (*GlobalEventHandlers) SetOnpointerout(onpointerout func(Event)) {
	js.Rewrite("$<.Onpointerout = $1", onpointerout)
}

// Onpointerover
func (*GlobalEventHandlers) Onpointerover() (onpointerover func(Event)) {
	js.Rewrite("$<.Onpointerover")
	return onpointerover
}

// Onpointerover
func (*GlobalEventHandlers) SetOnpointerover(onpointerover func(Event)) {
	js.Rewrite("$<.Onpointerover = $1", onpointerover)
}

// Onpointerup
func (*GlobalEventHandlers) Onpointerup() (onpointerup func(Event)) {
	js.Rewrite("$<.Onpointerup")
	return onpointerup
}

// Onpointerup
func (*GlobalEventHandlers) SetOnpointerup(onpointerup func(Event)) {
	js.Rewrite("$<.Onpointerup = $1", onpointerup)
}

// Onwheel
func (*GlobalEventHandlers) Onwheel() (onwheel func(Event)) {
	js.Rewrite("$<.Onwheel")
	return onwheel
}

// Onwheel
func (*GlobalEventHandlers) SetOnwheel(onwheel func(Event)) {
	js.Rewrite("$<.Onwheel = $1", onwheel)
}
