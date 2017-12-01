package speechsynthesisvoice

import "github.com/matthewmueller/golly/js"

// SpeechSynthesisVoice struct
// js:"SpeechSynthesisVoice,omit"
type SpeechSynthesisVoice struct {
}

// Default prop
// js:"default"
func (*SpeechSynthesisVoice) Default() (def bool) {
	js.Rewrite("$_.default")
	return def
}

// Lang prop
// js:"lang"
func (*SpeechSynthesisVoice) Lang() (lang string) {
	js.Rewrite("$_.lang")
	return lang
}

// LocalService prop
// js:"localService"
func (*SpeechSynthesisVoice) LocalService() (localService bool) {
	js.Rewrite("$_.localService")
	return localService
}

// Name prop
// js:"name"
func (*SpeechSynthesisVoice) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// VoiceURI prop
// js:"voiceURI"
func (*SpeechSynthesisVoice) VoiceURI() (voiceURI string) {
	js.Rewrite("$_.voiceURI")
	return voiceURI
}
