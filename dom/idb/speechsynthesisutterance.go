package idb

import (
	"github.com/matthewmueller/golly/dom2/speechsynthesisvoice"
	"github.com/matthewmueller/golly/js"
)

// NewSpeechSynthesisUtterance fn
func NewSpeechSynthesisUtterance(text *string) *SpeechSynthesisUtterance {
	js.Rewrite("SpeechSynthesisUtterance")
	return &SpeechSynthesisUtterance{}
}

// SpeechSynthesisUtterance struct
// js:"SpeechSynthesisUtterance,omit"
type SpeechSynthesisUtterance struct {
	EventTarget
}

// Lang
func (*SpeechSynthesisUtterance) Lang() (lang string) {
	js.Rewrite("$<.Lang")
	return lang
}

// Lang
func (*SpeechSynthesisUtterance) SetLang(lang string) {
	js.Rewrite("$<.Lang = $1", lang)
}

// Onboundary
func (*SpeechSynthesisUtterance) Onboundary() (onboundary func(Event)) {
	js.Rewrite("$<.Onboundary")
	return onboundary
}

// Onboundary
func (*SpeechSynthesisUtterance) SetOnboundary(onboundary func(Event)) {
	js.Rewrite("$<.Onboundary = $1", onboundary)
}

// Onend
func (*SpeechSynthesisUtterance) Onend() (onend func(Event)) {
	js.Rewrite("$<.Onend")
	return onend
}

// Onend
func (*SpeechSynthesisUtterance) SetOnend(onend func(Event)) {
	js.Rewrite("$<.Onend = $1", onend)
}

// Onerror
func (*SpeechSynthesisUtterance) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*SpeechSynthesisUtterance) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onmark
func (*SpeechSynthesisUtterance) Onmark() (onmark func(Event)) {
	js.Rewrite("$<.Onmark")
	return onmark
}

// Onmark
func (*SpeechSynthesisUtterance) SetOnmark(onmark func(Event)) {
	js.Rewrite("$<.Onmark = $1", onmark)
}

// Onpause
func (*SpeechSynthesisUtterance) Onpause() (onpause func(Event)) {
	js.Rewrite("$<.Onpause")
	return onpause
}

// Onpause
func (*SpeechSynthesisUtterance) SetOnpause(onpause func(Event)) {
	js.Rewrite("$<.Onpause = $1", onpause)
}

// Onresume
func (*SpeechSynthesisUtterance) Onresume() (onresume func(Event)) {
	js.Rewrite("$<.Onresume")
	return onresume
}

// Onresume
func (*SpeechSynthesisUtterance) SetOnresume(onresume func(Event)) {
	js.Rewrite("$<.Onresume = $1", onresume)
}

// Onstart
func (*SpeechSynthesisUtterance) Onstart() (onstart func(Event)) {
	js.Rewrite("$<.Onstart")
	return onstart
}

// Onstart
func (*SpeechSynthesisUtterance) SetOnstart(onstart func(Event)) {
	js.Rewrite("$<.Onstart = $1", onstart)
}

// Pitch
func (*SpeechSynthesisUtterance) Pitch() (pitch float32) {
	js.Rewrite("$<.Pitch")
	return pitch
}

// Pitch
func (*SpeechSynthesisUtterance) SetPitch(pitch float32) {
	js.Rewrite("$<.Pitch = $1", pitch)
}

// Rate
func (*SpeechSynthesisUtterance) Rate() (rate float32) {
	js.Rewrite("$<.Rate")
	return rate
}

// Rate
func (*SpeechSynthesisUtterance) SetRate(rate float32) {
	js.Rewrite("$<.Rate = $1", rate)
}

// Text
func (*SpeechSynthesisUtterance) Text() (text string) {
	js.Rewrite("$<.Text")
	return text
}

// Text
func (*SpeechSynthesisUtterance) SetText(text string) {
	js.Rewrite("$<.Text = $1", text)
}

// Voice
func (*SpeechSynthesisUtterance) Voice() (voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.Voice")
	return voice
}

// Voice
func (*SpeechSynthesisUtterance) SetVoice(voice *speechsynthesisvoice.SpeechSynthesisVoice) {
	js.Rewrite("$<.Voice = $1", voice)
}

// Volume
func (*SpeechSynthesisUtterance) Volume() (volume float32) {
	js.Rewrite("$<.Volume")
	return volume
}

// Volume
func (*SpeechSynthesisUtterance) SetVolume(volume float32) {
	js.Rewrite("$<.Volume = $1", volume)
}
