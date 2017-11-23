package htmlmarqueeelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLMarqueeElement struct
// js:"HTMLMarqueeElement,omit"
type HTMLMarqueeElement struct {
	window.HTMLElement
}

// Start fn
func (*HTMLMarqueeElement) Start() {
	js.Rewrite("$<.start()")
}

// Stop fn
func (*HTMLMarqueeElement) Stop() {
	js.Rewrite("$<.stop()")
}

// Behavior prop
func (*HTMLMarqueeElement) Behavior() (behavior string) {
	js.Rewrite("$<.behavior")
	return behavior
}

// Behavior prop
func (*HTMLMarqueeElement) SetBehavior(behavior string) {
	js.Rewrite("$<.behavior = $1", behavior)
}

// BgColor prop
func (*HTMLMarqueeElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

// BgColor prop
func (*HTMLMarqueeElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

// Direction prop
func (*HTMLMarqueeElement) Direction() (direction string) {
	js.Rewrite("$<.direction")
	return direction
}

// Direction prop
func (*HTMLMarqueeElement) SetDirection(direction string) {
	js.Rewrite("$<.direction = $1", direction)
}

// Height prop
func (*HTMLMarqueeElement) Height() (height string) {
	js.Rewrite("$<.height")
	return height
}

// Height prop
func (*HTMLMarqueeElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

// Hspace prop
func (*HTMLMarqueeElement) Hspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

// Hspace prop
func (*HTMLMarqueeElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

// Loop prop
func (*HTMLMarqueeElement) Loop() (loop int) {
	js.Rewrite("$<.loop")
	return loop
}

// Loop prop
func (*HTMLMarqueeElement) SetLoop(loop int) {
	js.Rewrite("$<.loop = $1", loop)
}

// Onbounce prop
func (*HTMLMarqueeElement) Onbounce() (onbounce func(window.Event)) {
	js.Rewrite("$<.onbounce")
	return onbounce
}

// Onbounce prop
func (*HTMLMarqueeElement) SetOnbounce(onbounce func(window.Event)) {
	js.Rewrite("$<.onbounce = $1", onbounce)
}

// Onfinish prop
func (*HTMLMarqueeElement) Onfinish() (onfinish func(window.Event)) {
	js.Rewrite("$<.onfinish")
	return onfinish
}

// Onfinish prop
func (*HTMLMarqueeElement) SetOnfinish(onfinish func(window.Event)) {
	js.Rewrite("$<.onfinish = $1", onfinish)
}

// Onstart prop
func (*HTMLMarqueeElement) Onstart() (onstart func(window.Event)) {
	js.Rewrite("$<.onstart")
	return onstart
}

// Onstart prop
func (*HTMLMarqueeElement) SetOnstart(onstart func(window.Event)) {
	js.Rewrite("$<.onstart = $1", onstart)
}

// ScrollAmount prop
func (*HTMLMarqueeElement) ScrollAmount() (scrollAmount uint) {
	js.Rewrite("$<.scrollAmount")
	return scrollAmount
}

// ScrollAmount prop
func (*HTMLMarqueeElement) SetScrollAmount(scrollAmount uint) {
	js.Rewrite("$<.scrollAmount = $1", scrollAmount)
}

// ScrollDelay prop
func (*HTMLMarqueeElement) ScrollDelay() (scrollDelay uint) {
	js.Rewrite("$<.scrollDelay")
	return scrollDelay
}

// ScrollDelay prop
func (*HTMLMarqueeElement) SetScrollDelay(scrollDelay uint) {
	js.Rewrite("$<.scrollDelay = $1", scrollDelay)
}

// TrueSpeed prop
func (*HTMLMarqueeElement) TrueSpeed() (trueSpeed bool) {
	js.Rewrite("$<.trueSpeed")
	return trueSpeed
}

// TrueSpeed prop
func (*HTMLMarqueeElement) SetTrueSpeed(trueSpeed bool) {
	js.Rewrite("$<.trueSpeed = $1", trueSpeed)
}

// Vspace prop
func (*HTMLMarqueeElement) Vspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

// Vspace prop
func (*HTMLMarqueeElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

// Width prop
func (*HTMLMarqueeElement) Width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// Width prop
func (*HTMLMarqueeElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
