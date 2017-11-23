package window

import (
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/js"
)

// MSInputMethodContext struct
// js:"MSInputMethodContext,omit"
type MSInputMethodContext struct {
	EventTarget
}

// GetCandidateWindowClientRect fn
func (*MSInputMethodContext) GetCandidateWindowClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.getCandidateWindowClientRect()")
	return c
}

// GetCompositionAlternatives fn
func (*MSInputMethodContext) GetCompositionAlternatives() (s []string) {
	js.Rewrite("$<.getCompositionAlternatives()")
	return s
}

// HasComposition fn
func (*MSInputMethodContext) HasComposition() (b bool) {
	js.Rewrite("$<.hasComposition()")
	return b
}

// IsCandidateWindowVisible fn
func (*MSInputMethodContext) IsCandidateWindowVisible() (b bool) {
	js.Rewrite("$<.isCandidateWindowVisible()")
	return b
}

// CompositionEndOffset prop
func (*MSInputMethodContext) CompositionEndOffset() (compositionEndOffset uint) {
	js.Rewrite("$<.compositionEndOffset")
	return compositionEndOffset
}

// CompositionStartOffset prop
func (*MSInputMethodContext) CompositionStartOffset() (compositionStartOffset uint) {
	js.Rewrite("$<.compositionStartOffset")
	return compositionStartOffset
}

// Oncandidatewindowhide prop
func (*MSInputMethodContext) Oncandidatewindowhide() (oncandidatewindowhide func(Event)) {
	js.Rewrite("$<.oncandidatewindowhide")
	return oncandidatewindowhide
}

// Oncandidatewindowhide prop
func (*MSInputMethodContext) SetOncandidatewindowhide(oncandidatewindowhide func(Event)) {
	js.Rewrite("$<.oncandidatewindowhide = $1", oncandidatewindowhide)
}

// Oncandidatewindowshow prop
func (*MSInputMethodContext) Oncandidatewindowshow() (oncandidatewindowshow func(Event)) {
	js.Rewrite("$<.oncandidatewindowshow")
	return oncandidatewindowshow
}

// Oncandidatewindowshow prop
func (*MSInputMethodContext) SetOncandidatewindowshow(oncandidatewindowshow func(Event)) {
	js.Rewrite("$<.oncandidatewindowshow = $1", oncandidatewindowshow)
}

// Oncandidatewindowupdate prop
func (*MSInputMethodContext) Oncandidatewindowupdate() (oncandidatewindowupdate func(Event)) {
	js.Rewrite("$<.oncandidatewindowupdate")
	return oncandidatewindowupdate
}

// Oncandidatewindowupdate prop
func (*MSInputMethodContext) SetOncandidatewindowupdate(oncandidatewindowupdate func(Event)) {
	js.Rewrite("$<.oncandidatewindowupdate = $1", oncandidatewindowupdate)
}

// Target prop
func (*MSInputMethodContext) Target() (target HTMLElement) {
	js.Rewrite("$<.target")
	return target
}
