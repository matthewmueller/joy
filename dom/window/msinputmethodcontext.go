package window

import (
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*MSInputMethodContext)(nil)

// MSInputMethodContext struct
// js:"MSInputMethodContext,omit"
type MSInputMethodContext struct {
}

// GetCandidateWindowClientRect fn
// js:"getCandidateWindowClientRect"
func (*MSInputMethodContext) GetCandidateWindowClientRect() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.getCandidateWindowClientRect()")
	return c
}

// GetCompositionAlternatives fn
// js:"getCompositionAlternatives"
func (*MSInputMethodContext) GetCompositionAlternatives() (s []string) {
	macro.Rewrite("$_.getCompositionAlternatives()")
	return s
}

// HasComposition fn
// js:"hasComposition"
func (*MSInputMethodContext) HasComposition() (b bool) {
	macro.Rewrite("$_.hasComposition()")
	return b
}

// IsCandidateWindowVisible fn
// js:"isCandidateWindowVisible"
func (*MSInputMethodContext) IsCandidateWindowVisible() (b bool) {
	macro.Rewrite("$_.isCandidateWindowVisible()")
	return b
}

// AddEventListener fn
// js:"addEventListener"
func (*MSInputMethodContext) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MSInputMethodContext) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MSInputMethodContext) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// CompositionEndOffset prop
// js:"compositionEndOffset"
func (*MSInputMethodContext) CompositionEndOffset() (compositionEndOffset uint) {
	macro.Rewrite("$_.compositionEndOffset")
	return compositionEndOffset
}

// CompositionStartOffset prop
// js:"compositionStartOffset"
func (*MSInputMethodContext) CompositionStartOffset() (compositionStartOffset uint) {
	macro.Rewrite("$_.compositionStartOffset")
	return compositionStartOffset
}

// Oncandidatewindowhide prop
// js:"oncandidatewindowhide"
func (*MSInputMethodContext) Oncandidatewindowhide() (oncandidatewindowhide func(Event)) {
	macro.Rewrite("$_.oncandidatewindowhide")
	return oncandidatewindowhide
}

// SetOncandidatewindowhide prop
// js:"oncandidatewindowhide"
func (*MSInputMethodContext) SetOncandidatewindowhide(oncandidatewindowhide func(Event)) {
	macro.Rewrite("$_.oncandidatewindowhide = $1", oncandidatewindowhide)
}

// Oncandidatewindowshow prop
// js:"oncandidatewindowshow"
func (*MSInputMethodContext) Oncandidatewindowshow() (oncandidatewindowshow func(Event)) {
	macro.Rewrite("$_.oncandidatewindowshow")
	return oncandidatewindowshow
}

// SetOncandidatewindowshow prop
// js:"oncandidatewindowshow"
func (*MSInputMethodContext) SetOncandidatewindowshow(oncandidatewindowshow func(Event)) {
	macro.Rewrite("$_.oncandidatewindowshow = $1", oncandidatewindowshow)
}

// Oncandidatewindowupdate prop
// js:"oncandidatewindowupdate"
func (*MSInputMethodContext) Oncandidatewindowupdate() (oncandidatewindowupdate func(Event)) {
	macro.Rewrite("$_.oncandidatewindowupdate")
	return oncandidatewindowupdate
}

// SetOncandidatewindowupdate prop
// js:"oncandidatewindowupdate"
func (*MSInputMethodContext) SetOncandidatewindowupdate(oncandidatewindowupdate func(Event)) {
	macro.Rewrite("$_.oncandidatewindowupdate = $1", oncandidatewindowupdate)
}

// Target prop
// js:"target"
func (*MSInputMethodContext) Target() (target HTMLElement) {
	macro.Rewrite("$_.target")
	return target
}
