package rtcdtmfsender

import (
	"github.com/matthewmueller/joy/dom/rtcdtmftonechangeevent"
	"github.com/matthewmueller/joy/dom/rtcrtpsender"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*RTCDtmfSender)(nil)

// New fn
func New(sender *rtcrtpsender.RTCRtpSender) *RTCDtmfSender {
	js.Rewrite("RTCDtmfSender")
	return &RTCDtmfSender{}
}

// RTCDtmfSender struct
// js:"RTCDtmfSender,omit"
type RTCDtmfSender struct {
}

// InsertDTMF fn
// js:"insertDTMF"
func (*RTCDtmfSender) InsertDTMF(tones string, duration *int, interToneGap *int) {
	js.Rewrite("$_.insertDTMF($1, $2, $3)", tones, duration, interToneGap)
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCDtmfSender) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCDtmfSender) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCDtmfSender) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// CanInsertDTMF prop
// js:"canInsertDTMF"
func (*RTCDtmfSender) CanInsertDTMF() (canInsertDTMF bool) {
	js.Rewrite("$_.canInsertDTMF")
	return canInsertDTMF
}

// Duration prop
// js:"duration"
func (*RTCDtmfSender) Duration() (duration int) {
	js.Rewrite("$_.duration")
	return duration
}

// InterToneGap prop
// js:"interToneGap"
func (*RTCDtmfSender) InterToneGap() (interToneGap int) {
	js.Rewrite("$_.interToneGap")
	return interToneGap
}

// Ontonechange prop
// js:"ontonechange"
func (*RTCDtmfSender) Ontonechange() (ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	js.Rewrite("$_.ontonechange")
	return ontonechange
}

// SetOntonechange prop
// js:"ontonechange"
func (*RTCDtmfSender) SetOntonechange(ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	js.Rewrite("$_.ontonechange = $1", ontonechange)
}

// Sender prop
// js:"sender"
func (*RTCDtmfSender) Sender() (sender *rtcrtpsender.RTCRtpSender) {
	js.Rewrite("$_.sender")
	return sender
}

// ToneBuffer prop
// js:"toneBuffer"
func (*RTCDtmfSender) ToneBuffer() (toneBuffer string) {
	js.Rewrite("$_.toneBuffer")
	return toneBuffer
}
