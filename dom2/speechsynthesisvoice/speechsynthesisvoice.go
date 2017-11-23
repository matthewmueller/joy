package speechsynthesisvoice

import "github.com/matthewmueller/golly/js"

// js:"SpeechSynthesisVoice,omit"
type SpeechSynthesisVoice struct {
}

// Default
func (*SpeechSynthesisVoice) Default() (def bool) {
	js.Rewrite("$<.Default")
	return def
}

// Lang
func (*SpeechSynthesisVoice) Lang() (lang string) {
	js.Rewrite("$<.Lang")
	return lang
}

// LocalService
func (*SpeechSynthesisVoice) LocalService() (localService bool) {
	js.Rewrite("$<.LocalService")
	return localService
}

// Name
func (*SpeechSynthesisVoice) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// VoiceURI
func (*SpeechSynthesisVoice) VoiceURI() (voiceURI string) {
	js.Rewrite("$<.VoiceURI")
	return voiceURI
}
