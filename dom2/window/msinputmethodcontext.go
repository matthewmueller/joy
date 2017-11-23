package window

import (
	"github.com/matthewmueller/golly/dom2/clientrect"
	"github.com/matthewmueller/golly/js"
)

// MSInputMethodContext struct
// js:"MSInputMethodContext,omit"
type MSInputMethodContext struct {
	EventTarget
}

// GetCandidateWindowClientRect
func (*MSInputMethodContext) GetCandidateWindowClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.GetCandidateWindowClientRect()")
	return c
}

// GetCompositionAlternatives
func (*MSInputMethodContext) GetCompositionAlternatives() (s []string) {
	js.Rewrite("$<.GetCompositionAlternatives()")
	return s
}

// HasComposition
func (*MSInputMethodContext) HasComposition() (b bool) {
	js.Rewrite("$<.HasComposition()")
	return b
}

// IsCandidateWindowVisible
func (*MSInputMethodContext) IsCandidateWindowVisible() (b bool) {
	js.Rewrite("$<.IsCandidateWindowVisible()")
	return b
}

// CompositionEndOffset
func (*MSInputMethodContext) CompositionEndOffset() (compositionEndOffset uint) {
	js.Rewrite("$<.CompositionEndOffset")
	return compositionEndOffset
}

// CompositionStartOffset
func (*MSInputMethodContext) CompositionStartOffset() (compositionStartOffset uint) {
	js.Rewrite("$<.CompositionStartOffset")
	return compositionStartOffset
}

// Oncandidatewindowhide
func (*MSInputMethodContext) Oncandidatewindowhide() (oncandidatewindowhide func(Event)) {
	js.Rewrite("$<.Oncandidatewindowhide")
	return oncandidatewindowhide
}

// Oncandidatewindowhide
func (*MSInputMethodContext) SetOncandidatewindowhide(oncandidatewindowhide func(Event)) {
	js.Rewrite("$<.Oncandidatewindowhide = $1", oncandidatewindowhide)
}

// Oncandidatewindowshow
func (*MSInputMethodContext) Oncandidatewindowshow() (oncandidatewindowshow func(Event)) {
	js.Rewrite("$<.Oncandidatewindowshow")
	return oncandidatewindowshow
}

// Oncandidatewindowshow
func (*MSInputMethodContext) SetOncandidatewindowshow(oncandidatewindowshow func(Event)) {
	js.Rewrite("$<.Oncandidatewindowshow = $1", oncandidatewindowshow)
}

// Oncandidatewindowupdate
func (*MSInputMethodContext) Oncandidatewindowupdate() (oncandidatewindowupdate func(Event)) {
	js.Rewrite("$<.Oncandidatewindowupdate")
	return oncandidatewindowupdate
}

// Oncandidatewindowupdate
func (*MSInputMethodContext) SetOncandidatewindowupdate(oncandidatewindowupdate func(Event)) {
	js.Rewrite("$<.Oncandidatewindowupdate = $1", oncandidatewindowupdate)
}

// Target
func (*MSInputMethodContext) Target() (target HTMLElement) {
	js.Rewrite("$<.Target")
	return target
}
