package globaleventhandlers

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type GlobalEventHandlers struct {
}

func (*GlobalEventHandlers) GetOnpointercancel() (e *event.Event) {
	js.Rewrite("$<.onpointercancel")
	return e
}

func (*GlobalEventHandlers) SetOnpointercancel(e *event.Event) {
	js.Rewrite("$<.onpointercancel = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerdown() (e *event.Event) {
	js.Rewrite("$<.onpointerdown")
	return e
}

func (*GlobalEventHandlers) SetOnpointerdown(e *event.Event) {
	js.Rewrite("$<.onpointerdown = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerenter() (e *event.Event) {
	js.Rewrite("$<.onpointerenter")
	return e
}

func (*GlobalEventHandlers) SetOnpointerenter(e *event.Event) {
	js.Rewrite("$<.onpointerenter = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerleave() (e *event.Event) {
	js.Rewrite("$<.onpointerleave")
	return e
}

func (*GlobalEventHandlers) SetOnpointerleave(e *event.Event) {
	js.Rewrite("$<.onpointerleave = $1", e)
}

func (*GlobalEventHandlers) GetOnpointermove() (e *event.Event) {
	js.Rewrite("$<.onpointermove")
	return e
}

func (*GlobalEventHandlers) SetOnpointermove(e *event.Event) {
	js.Rewrite("$<.onpointermove = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerout() (e *event.Event) {
	js.Rewrite("$<.onpointerout")
	return e
}

func (*GlobalEventHandlers) SetOnpointerout(e *event.Event) {
	js.Rewrite("$<.onpointerout = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerover() (e *event.Event) {
	js.Rewrite("$<.onpointerover")
	return e
}

func (*GlobalEventHandlers) SetOnpointerover(e *event.Event) {
	js.Rewrite("$<.onpointerover = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerup() (e *event.Event) {
	js.Rewrite("$<.onpointerup")
	return e
}

func (*GlobalEventHandlers) SetOnpointerup(e *event.Event) {
	js.Rewrite("$<.onpointerup = $1", e)
}

func (*GlobalEventHandlers) GetOnwheel() (e *event.Event) {
	js.Rewrite("$<.onwheel")
	return e
}

func (*GlobalEventHandlers) SetOnwheel(e *event.Event) {
	js.Rewrite("$<.onwheel = $1", e)
}
