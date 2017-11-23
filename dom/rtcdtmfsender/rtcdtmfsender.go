package rtcdtmfsender

import (
	"github.com/matthewmueller/golly/dom2/rtcdtmftonechangeevent"
	"github.com/matthewmueller/golly/dom2/rtcrtpsender"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(sender *rtcrtpsender.RTCRtpSender) *RTCDtmfSender {
	js.Rewrite("RTCDtmfSender")
	return &RTCDtmfSender{}
}

// RTCDtmfSender struct
// js:"RTCDtmfSender,omit"
type RTCDtmfSender struct {
	window.EventTarget
}

// InsertDTMF fn
func (*RTCDtmfSender) InsertDTMF(tones string, duration *int, interToneGap *int) {
	js.Rewrite("$<.insertDTMF($1, $2, $3)", tones, duration, interToneGap)
}

// CanInsertDTMF prop
func (*RTCDtmfSender) CanInsertDTMF() (canInsertDTMF bool) {
	js.Rewrite("$<.canInsertDTMF")
	return canInsertDTMF
}

// Duration prop
func (*RTCDtmfSender) Duration() (duration int) {
	js.Rewrite("$<.duration")
	return duration
}

// InterToneGap prop
func (*RTCDtmfSender) InterToneGap() (interToneGap int) {
	js.Rewrite("$<.interToneGap")
	return interToneGap
}

// Ontonechange prop
func (*RTCDtmfSender) Ontonechange() (ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	js.Rewrite("$<.ontonechange")
	return ontonechange
}

// Ontonechange prop
func (*RTCDtmfSender) SetOntonechange(ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	js.Rewrite("$<.ontonechange = $1", ontonechange)
}

// Sender prop
func (*RTCDtmfSender) Sender() (sender *rtcrtpsender.RTCRtpSender) {
	js.Rewrite("$<.sender")
	return sender
}

// ToneBuffer prop
func (*RTCDtmfSender) ToneBuffer() (toneBuffer string) {
	js.Rewrite("$<.toneBuffer")
	return toneBuffer
}
