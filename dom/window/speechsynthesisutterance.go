package window

import (
	"github.com/matthewmueller/joy/dom/speechsynthesisvoice"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*SpeechSynthesisUtterance)(nil)

// NewSpeechSynthesisUtterance fn
func NewSpeechSynthesisUtterance(text *string) *SpeechSynthesisUtterance {
	macro.Rewrite("SpeechSynthesisUtterance")
	return &SpeechSynthesisUtterance{}
}

// SpeechSynthesisUtterance struct
// js:"SpeechSynthesisUtterance,omit"
type SpeechSynthesisUtterance struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*SpeechSynthesisUtterance) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SpeechSynthesisUtterance) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SpeechSynthesisUtterance) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Lang prop
// js:"lang"
func (*SpeechSynthesisUtterance) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

// SetLang prop
// js:"lang"
func (*SpeechSynthesisUtterance) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

// Onboundary prop
// js:"onboundary"
func (*SpeechSynthesisUtterance) Onboundary() (onboundary func(Event)) {
	macro.Rewrite("$_.onboundary")
	return onboundary
}

// SetOnboundary prop
// js:"onboundary"
func (*SpeechSynthesisUtterance) SetOnboundary(onboundary func(Event)) {
	macro.Rewrite("$_.onboundary = $1", onboundary)
}

// Onend prop
// js:"onend"
func (*SpeechSynthesisUtterance) Onend() (onend func(Event)) {
	macro.Rewrite("$_.onend")
	return onend
}

// SetOnend prop
// js:"onend"
func (*SpeechSynthesisUtterance) SetOnend(onend func(Event)) {
	macro.Rewrite("$_.onend = $1", onend)
}

// Onerror prop
// js:"onerror"
func (*SpeechSynthesisUtterance) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*SpeechSynthesisUtterance) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onmark prop
// js:"onmark"
func (*SpeechSynthesisUtterance) Onmark() (onmark func(Event)) {
	macro.Rewrite("$_.onmark")
	return onmark
}

// SetOnmark prop
// js:"onmark"
func (*SpeechSynthesisUtterance) SetOnmark(onmark func(Event)) {
	macro.Rewrite("$_.onmark = $1", onmark)
}

// Onpause prop
// js:"onpause"
func (*SpeechSynthesisUtterance) Onpause() (onpause func(Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*SpeechSynthesisUtterance) SetOnpause(onpause func(Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

// Onresume prop
// js:"onresume"
func (*SpeechSynthesisUtterance) Onresume() (onresume func(Event)) {
	macro.Rewrite("$_.onresume")
	return onresume
}

// SetOnresume prop
// js:"onresume"
func (*SpeechSynthesisUtterance) SetOnresume(onresume func(Event)) {
	macro.Rewrite("$_.onresume = $1", onresume)
}

// Onstart prop
// js:"onstart"
func (*SpeechSynthesisUtterance) Onstart() (onstart func(Event)) {
	macro.Rewrite("$_.onstart")
	return onstart
}

// SetOnstart prop
// js:"onstart"
func (*SpeechSynthesisUtterance) SetOnstart(onstart func(Event)) {
	macro.Rewrite("$_.onstart = $1", onstart)
}

// Pitch prop
// js:"pitch"
func (*SpeechSynthesisUtterance) Pitch() (pitch float32) {
	macro.Rewrite("$_.pitch")
	return pitch
}

// SetPitch prop
// js:"pitch"
func (*SpeechSynthesisUtterance) SetPitch(pitch float32) {
	macro.Rewrite("$_.pitch = $1", pitch)
}

// Rate prop
// js:"rate"
func (*SpeechSynthesisUtterance) Rate() (rate float32) {
	macro.Rewrite("$_.rate")
	return rate
}

// SetRate prop
// js:"rate"
func (*SpeechSynthesisUtterance) SetRate(rate float32) {
	macro.Rewrite("$_.rate = $1", rate)
}

// Text prop
// js:"text"
func (*SpeechSynthesisUtterance) Text() (text string) {
	macro.Rewrite("$_.text")
	return text
}

// SetText prop
// js:"text"
func (*SpeechSynthesisUtterance) SetText(text string) {
	macro.Rewrite("$_.text = $1", text)
}

// Voice prop
// js:"voice"
func (*SpeechSynthesisUtterance) Voice() (voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	macro.Rewrite("$_.voice")
	return voice
}

// SetVoice prop
// js:"voice"
func (*SpeechSynthesisUtterance) SetVoice(voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	macro.Rewrite("$_.voice = $1", voice)
}

// Volume prop
// js:"volume"
func (*SpeechSynthesisUtterance) Volume() (volume float32) {
	macro.Rewrite("$_.volume")
	return volume
}

// SetVolume prop
// js:"volume"
func (*SpeechSynthesisUtterance) SetVolume(volume float32) {
	macro.Rewrite("$_.volume = $1", volume)
}
