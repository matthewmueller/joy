package rtcdtmfsender

import (
	"github.com/matthewmueller/joy/dom/rtcdtmftonechangeevent"
	"github.com/matthewmueller/joy/dom/rtcrtpsender"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*RTCDtmfSender)(nil)

// New fn
func New(sender *rtcrtpsender.RTCRtpSender) *RTCDtmfSender {
	macro.Rewrite("new RTCDtmfSender($1)", sender)
	return &RTCDtmfSender{}
}

// RTCDtmfSender struct
// js:"RTCDtmfSender,omit"
type RTCDtmfSender struct {
}

// InsertDTMF fn
// js:"insertDTMF"
func (*RTCDtmfSender) InsertDTMF(tones string, duration *int, interToneGap *int) {
	macro.Rewrite("$_.insertDTMF($1, $2, $3)", tones, duration, interToneGap)
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCDtmfSender) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCDtmfSender) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCDtmfSender) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// CanInsertDTMF prop
// js:"canInsertDTMF"
func (*RTCDtmfSender) CanInsertDTMF() (canInsertDTMF bool) {
	macro.Rewrite("$_.canInsertDTMF")
	return canInsertDTMF
}

// Duration prop
// js:"duration"
func (*RTCDtmfSender) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

// InterToneGap prop
// js:"interToneGap"
func (*RTCDtmfSender) InterToneGap() (interToneGap int) {
	macro.Rewrite("$_.interToneGap")
	return interToneGap
}

// Ontonechange prop
// js:"ontonechange"
func (*RTCDtmfSender) Ontonechange() (ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	macro.Rewrite("$_.ontonechange")
	return ontonechange
}

// SetOntonechange prop
// js:"ontonechange"
func (*RTCDtmfSender) SetOntonechange(ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	macro.Rewrite("$_.ontonechange = $1", ontonechange)
}

// Sender prop
// js:"sender"
func (*RTCDtmfSender) Sender() (sender *rtcrtpsender.RTCRtpSender) {
	macro.Rewrite("$_.sender")
	return sender
}

// ToneBuffer prop
// js:"toneBuffer"
func (*RTCDtmfSender) ToneBuffer() (toneBuffer string) {
	macro.Rewrite("$_.toneBuffer")
	return toneBuffer
}
