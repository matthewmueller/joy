package rtcdtmfsender

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtmftonechangeevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpsender"
	"github.com/matthewmueller/golly/js"
)

type RTCDtmfSender struct {
	*eventtarget.EventTarget
}

func (*RTCDtmfSender) InsertDTMF(tones string, duration int, interToneGap int) {
	js.Rewrite("$<.insertDTMF($1, $2, $3)", tones, duration, interToneGap)
}

func (*RTCDtmfSender) GetCanInsertDTMF() (canInsertDTMF bool) {
	js.Rewrite("$<.canInsertDTMF")
	return canInsertDTMF
}

func (*RTCDtmfSender) GetDuration() (duration int) {
	js.Rewrite("$<.duration")
	return duration
}

func (*RTCDtmfSender) GetInterToneGap() (interToneGap int) {
	js.Rewrite("$<.interToneGap")
	return interToneGap
}

func (*RTCDtmfSender) GetOntonechange() (tonechange *rtcdtmftonechangeevent.RTCDTMFToneChangeEvent) {
	js.Rewrite("$<.ontonechange")
	return tonechange
}

func (*RTCDtmfSender) SetOntonechange(tonechange *rtcdtmftonechangeevent.RTCDTMFToneChangeEvent) {
	js.Rewrite("$<.ontonechange = $1", tonechange)
}

func (*RTCDtmfSender) GetSender() (sender *rtcrtpsender.RTCRtpSender) {
	js.Rewrite("$<.sender")
	return sender
}

func (*RTCDtmfSender) GetToneBuffer() (toneBuffer string) {
	js.Rewrite("$<.toneBuffer")
	return toneBuffer
}
