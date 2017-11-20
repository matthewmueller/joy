package speechsynthesisutterance

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/speechsynthesisvoice"
	"github.com/matthewmueller/golly/js"
)

type SpeechSynthesisUtterance struct {
	*eventtarget.EventTarget
}

func (*SpeechSynthesisUtterance) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*SpeechSynthesisUtterance) SetLang(lang string) {
	js.Rewrite("$<.lang = $1", lang)
}

func (*SpeechSynthesisUtterance) GetOnboundary() (boundary *event.Event) {
	js.Rewrite("$<.onboundary")
	return boundary
}

func (*SpeechSynthesisUtterance) SetOnboundary(boundary *event.Event) {
	js.Rewrite("$<.onboundary = $1", boundary)
}

func (*SpeechSynthesisUtterance) GetOnend() (end *event.Event) {
	js.Rewrite("$<.onend")
	return end
}

func (*SpeechSynthesisUtterance) SetOnend(end *event.Event) {
	js.Rewrite("$<.onend = $1", end)
}

func (*SpeechSynthesisUtterance) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*SpeechSynthesisUtterance) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*SpeechSynthesisUtterance) GetOnmark() (mark *event.Event) {
	js.Rewrite("$<.onmark")
	return mark
}

func (*SpeechSynthesisUtterance) SetOnmark(mark *event.Event) {
	js.Rewrite("$<.onmark = $1", mark)
}

func (*SpeechSynthesisUtterance) GetOnpause() (pause *event.Event) {
	js.Rewrite("$<.onpause")
	return pause
}

func (*SpeechSynthesisUtterance) SetOnpause(pause *event.Event) {
	js.Rewrite("$<.onpause = $1", pause)
}

func (*SpeechSynthesisUtterance) GetOnresume() (resume *event.Event) {
	js.Rewrite("$<.onresume")
	return resume
}

func (*SpeechSynthesisUtterance) SetOnresume(resume *event.Event) {
	js.Rewrite("$<.onresume = $1", resume)
}

func (*SpeechSynthesisUtterance) GetOnstart() (start *event.Event) {
	js.Rewrite("$<.onstart")
	return start
}

func (*SpeechSynthesisUtterance) SetOnstart(start *event.Event) {
	js.Rewrite("$<.onstart = $1", start)
}

func (*SpeechSynthesisUtterance) GetPitch() (pitch float32) {
	js.Rewrite("$<.pitch")
	return pitch
}

func (*SpeechSynthesisUtterance) SetPitch(pitch float32) {
	js.Rewrite("$<.pitch = $1", pitch)
}

func (*SpeechSynthesisUtterance) GetRate() (rate float32) {
	js.Rewrite("$<.rate")
	return rate
}

func (*SpeechSynthesisUtterance) SetRate(rate float32) {
	js.Rewrite("$<.rate = $1", rate)
}

func (*SpeechSynthesisUtterance) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*SpeechSynthesisUtterance) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*SpeechSynthesisUtterance) GetVoice() (voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.voice")
	return voice
}

func (*SpeechSynthesisUtterance) SetVoice(voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.voice = $1", voice)
}

func (*SpeechSynthesisUtterance) GetVolume() (volume float32) {
	js.Rewrite("$<.volume")
	return volume
}

func (*SpeechSynthesisUtterance) SetVolume(volume float32) {
	js.Rewrite("$<.volume = $1", volume)
}
