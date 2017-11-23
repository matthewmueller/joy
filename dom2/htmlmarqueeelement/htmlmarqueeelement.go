package htmlmarqueeelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLMarqueeElement,omit"
type HTMLMarqueeElement struct {
	window.HTMLElement
}

// Start
func (*HTMLMarqueeElement) Start() {
	js.Rewrite("$<.Start()")
}

// Stop
func (*HTMLMarqueeElement) Stop() {
	js.Rewrite("$<.Stop()")
}

// Behavior
func (*HTMLMarqueeElement) Behavior() (behavior string) {
	js.Rewrite("$<.Behavior")
	return behavior
}

// Behavior
func (*HTMLMarqueeElement) SetBehavior(behavior string) {
	js.Rewrite("$<.Behavior = $1", behavior)
}

// BgColor
func (*HTMLMarqueeElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.BgColor")
	return bgColor
}

// BgColor
func (*HTMLMarqueeElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.BgColor = $1", bgColor)
}

// Direction
func (*HTMLMarqueeElement) Direction() (direction string) {
	js.Rewrite("$<.Direction")
	return direction
}

// Direction
func (*HTMLMarqueeElement) SetDirection(direction string) {
	js.Rewrite("$<.Direction = $1", direction)
}

// Height
func (*HTMLMarqueeElement) Height() (height string) {
	js.Rewrite("$<.Height")
	return height
}

// Height
func (*HTMLMarqueeElement) SetHeight(height string) {
	js.Rewrite("$<.Height = $1", height)
}

// Hspace
func (*HTMLMarqueeElement) Hspace() (hspace int) {
	js.Rewrite("$<.Hspace")
	return hspace
}

// Hspace
func (*HTMLMarqueeElement) SetHspace(hspace int) {
	js.Rewrite("$<.Hspace = $1", hspace)
}

// Loop
func (*HTMLMarqueeElement) Loop() (loop int) {
	js.Rewrite("$<.Loop")
	return loop
}

// Loop
func (*HTMLMarqueeElement) SetLoop(loop int) {
	js.Rewrite("$<.Loop = $1", loop)
}

// Onbounce
func (*HTMLMarqueeElement) Onbounce() (onbounce func(window.Event)) {
	js.Rewrite("$<.Onbounce")
	return onbounce
}

// Onbounce
func (*HTMLMarqueeElement) SetOnbounce(onbounce func(window.Event)) {
	js.Rewrite("$<.Onbounce = $1", onbounce)
}

// Onfinish
func (*HTMLMarqueeElement) Onfinish() (onfinish func(window.Event)) {
	js.Rewrite("$<.Onfinish")
	return onfinish
}

// Onfinish
func (*HTMLMarqueeElement) SetOnfinish(onfinish func(window.Event)) {
	js.Rewrite("$<.Onfinish = $1", onfinish)
}

// Onstart
func (*HTMLMarqueeElement) Onstart() (onstart func(window.Event)) {
	js.Rewrite("$<.Onstart")
	return onstart
}

// Onstart
func (*HTMLMarqueeElement) SetOnstart(onstart func(window.Event)) {
	js.Rewrite("$<.Onstart = $1", onstart)
}

// ScrollAmount
func (*HTMLMarqueeElement) ScrollAmount() (scrollAmount uint) {
	js.Rewrite("$<.ScrollAmount")
	return scrollAmount
}

// ScrollAmount
func (*HTMLMarqueeElement) SetScrollAmount(scrollAmount uint) {
	js.Rewrite("$<.ScrollAmount = $1", scrollAmount)
}

// ScrollDelay
func (*HTMLMarqueeElement) ScrollDelay() (scrollDelay uint) {
	js.Rewrite("$<.ScrollDelay")
	return scrollDelay
}

// ScrollDelay
func (*HTMLMarqueeElement) SetScrollDelay(scrollDelay uint) {
	js.Rewrite("$<.ScrollDelay = $1", scrollDelay)
}

// TrueSpeed
func (*HTMLMarqueeElement) TrueSpeed() (trueSpeed bool) {
	js.Rewrite("$<.TrueSpeed")
	return trueSpeed
}

// TrueSpeed
func (*HTMLMarqueeElement) SetTrueSpeed(trueSpeed bool) {
	js.Rewrite("$<.TrueSpeed = $1", trueSpeed)
}

// Vspace
func (*HTMLMarqueeElement) Vspace() (vspace int) {
	js.Rewrite("$<.Vspace")
	return vspace
}

// Vspace
func (*HTMLMarqueeElement) SetVspace(vspace int) {
	js.Rewrite("$<.Vspace = $1", vspace)
}

// Width
func (*HTMLMarqueeElement) Width() (width string) {
	js.Rewrite("$<.Width")
	return width
}

// Width
func (*HTMLMarqueeElement) SetWidth(width string) {
	js.Rewrite("$<.Width = $1", width)
}
