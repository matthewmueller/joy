package speechsynthesisvoice

import "github.com/matthewmueller/golly/js"

// SpeechSynthesisVoice struct
// js:"SpeechSynthesisVoice,omit"
type SpeechSynthesisVoice struct {
}

// Default prop
func (*SpeechSynthesisVoice) Default() (def bool) {
	js.Rewrite("$<.default")
	return def
}

// Lang prop
func (*SpeechSynthesisVoice) Lang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

// LocalService prop
func (*SpeechSynthesisVoice) LocalService() (localService bool) {
	js.Rewrite("$<.localService")
	return localService
}

// Name prop
func (*SpeechSynthesisVoice) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// VoiceURI prop
func (*SpeechSynthesisVoice) VoiceURI() (voiceURI string) {
	js.Rewrite("$<.voiceURI")
	return voiceURI
}
