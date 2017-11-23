package rtcdtmfsender

import (
	"github.com/matthewmueller/golly/dom2/rtcdtmftonechangeevent"
	"github.com/matthewmueller/golly/dom2/rtcrtpsender"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCDtmfSender,omit"
type RTCDtmfSender struct {
	window.EventTarget
}

// InsertDTMF
func (*RTCDtmfSender) InsertDTMF(tones string, duration *int, interToneGap *int) {
	js.Rewrite("$<.InsertDTMF($1, $2, $3)", tones, duration, interToneGap)
}

// CanInsertDTMF
func (*RTCDtmfSender) CanInsertDTMF() (canInsertDTMF bool) {
	js.Rewrite("$<.CanInsertDTMF")
	return canInsertDTMF
}

// Duration
func (*RTCDtmfSender) Duration() (duration int) {
	js.Rewrite("$<.Duration")
	return duration
}

// InterToneGap
func (*RTCDtmfSender) InterToneGap() (interToneGap int) {
	js.Rewrite("$<.InterToneGap")
	return interToneGap
}

// Ontonechange
func (*RTCDtmfSender) Ontonechange() (ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	js.Rewrite("$<.Ontonechange")
	return ontonechange
}

// Ontonechange
func (*RTCDtmfSender) SetOntonechange(ontonechange func(*rtcdtmftonechangeevent.RTCDTMFToneChangeEvent)) {
	js.Rewrite("$<.Ontonechange = $1", ontonechange)
}

// Sender
func (*RTCDtmfSender) Sender() (sender *rtcrtpsender.RTCRtpSender) {
	js.Rewrite("$<.Sender")
	return sender
}

// ToneBuffer
func (*RTCDtmfSender) ToneBuffer() (toneBuffer string) {
	js.Rewrite("$<.ToneBuffer")
	return toneBuffer
}
