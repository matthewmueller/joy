package msinputmethodcontext

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/clientrect"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type MSInputMethodContext struct {
	*eventtarget.EventTarget
}

func (*MSInputMethodContext) GetCandidateWindowClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.getCandidateWindowClientRect()")
	return c
}

func (*MSInputMethodContext) GetCompositionAlternatives() (s []string) {
	js.Rewrite("$<.getCompositionAlternatives()")
	return s
}

func (*MSInputMethodContext) HasComposition() (b bool) {
	js.Rewrite("$<.hasComposition()")
	return b
}

func (*MSInputMethodContext) IsCandidateWindowVisible() (b bool) {
	js.Rewrite("$<.isCandidateWindowVisible()")
	return b
}

func (*MSInputMethodContext) GetCompositionEndOffset() (compositionEndOffset uint) {
	js.Rewrite("$<.compositionEndOffset")
	return compositionEndOffset
}

func (*MSInputMethodContext) GetCompositionStartOffset() (compositionStartOffset uint) {
	js.Rewrite("$<.compositionStartOffset")
	return compositionStartOffset
}

func (*MSInputMethodContext) GetOncandidatewindowhide() (MSCandidateWindowHide *event.Event) {
	js.Rewrite("$<.oncandidatewindowhide")
	return MSCandidateWindowHide
}

func (*MSInputMethodContext) SetOncandidatewindowhide(MSCandidateWindowHide *event.Event) {
	js.Rewrite("$<.oncandidatewindowhide = $1", MSCandidateWindowHide)
}

func (*MSInputMethodContext) GetOncandidatewindowshow() (MSCandidateWindowShow *event.Event) {
	js.Rewrite("$<.oncandidatewindowshow")
	return MSCandidateWindowShow
}

func (*MSInputMethodContext) SetOncandidatewindowshow(MSCandidateWindowShow *event.Event) {
	js.Rewrite("$<.oncandidatewindowshow = $1", MSCandidateWindowShow)
}

func (*MSInputMethodContext) GetOncandidatewindowupdate() (MSCandidateWindowUpdate *event.Event) {
	js.Rewrite("$<.oncandidatewindowupdate")
	return MSCandidateWindowUpdate
}

func (*MSInputMethodContext) SetOncandidatewindowupdate(MSCandidateWindowUpdate *event.Event) {
	js.Rewrite("$<.oncandidatewindowupdate = $1", MSCandidateWindowUpdate)
}

func (*MSInputMethodContext) GetTarget() (target *htmlelement.HTMLElement) {
	js.Rewrite("$<.target")
	return target
}
