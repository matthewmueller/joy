package htmlmarqueeelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLMarqueeElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLMarqueeElement) Start() {
	js.Rewrite("$<.start()")
}

func (*HTMLMarqueeElement) Stop() {
	js.Rewrite("$<.stop()")
}

func (*HTMLMarqueeElement) GetBehavior() (behavior string) {
	js.Rewrite("$<.behavior")
	return behavior
}

func (*HTMLMarqueeElement) SetBehavior(behavior string) {
	js.Rewrite("$<.behavior = $1", behavior)
}

func (*HTMLMarqueeElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLMarqueeElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLMarqueeElement) GetDirection() (direction string) {
	js.Rewrite("$<.direction")
	return direction
}

func (*HTMLMarqueeElement) SetDirection(direction string) {
	js.Rewrite("$<.direction = $1", direction)
}

func (*HTMLMarqueeElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLMarqueeElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLMarqueeElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLMarqueeElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLMarqueeElement) GetLoop() (loop int) {
	js.Rewrite("$<.loop")
	return loop
}

func (*HTMLMarqueeElement) SetLoop(loop int) {
	js.Rewrite("$<.loop = $1", loop)
}

func (*HTMLMarqueeElement) GetOnbounce() (bounce *event.Event) {
	js.Rewrite("$<.onbounce")
	return bounce
}

func (*HTMLMarqueeElement) SetOnbounce(bounce *event.Event) {
	js.Rewrite("$<.onbounce = $1", bounce)
}

func (*HTMLMarqueeElement) GetOnfinish() (finish *event.Event) {
	js.Rewrite("$<.onfinish")
	return finish
}

func (*HTMLMarqueeElement) SetOnfinish(finish *event.Event) {
	js.Rewrite("$<.onfinish = $1", finish)
}

func (*HTMLMarqueeElement) GetOnstart() (start *event.Event) {
	js.Rewrite("$<.onstart")
	return start
}

func (*HTMLMarqueeElement) SetOnstart(start *event.Event) {
	js.Rewrite("$<.onstart = $1", start)
}

func (*HTMLMarqueeElement) GetScrollAmount() (scrollAmount uint) {
	js.Rewrite("$<.scrollAmount")
	return scrollAmount
}

func (*HTMLMarqueeElement) SetScrollAmount(scrollAmount uint) {
	js.Rewrite("$<.scrollAmount = $1", scrollAmount)
}

func (*HTMLMarqueeElement) GetScrollDelay() (scrollDelay uint) {
	js.Rewrite("$<.scrollDelay")
	return scrollDelay
}

func (*HTMLMarqueeElement) SetScrollDelay(scrollDelay uint) {
	js.Rewrite("$<.scrollDelay = $1", scrollDelay)
}

func (*HTMLMarqueeElement) GetTrueSpeed() (trueSpeed bool) {
	js.Rewrite("$<.trueSpeed")
	return trueSpeed
}

func (*HTMLMarqueeElement) SetTrueSpeed(trueSpeed bool) {
	js.Rewrite("$<.trueSpeed = $1", trueSpeed)
}

func (*HTMLMarqueeElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLMarqueeElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLMarqueeElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLMarqueeElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
