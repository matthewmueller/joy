package speechsynthesisvoice

import "github.com/matthewmueller/joy/macro"

// SpeechSynthesisVoice struct
// js:"SpeechSynthesisVoice,omit"
type SpeechSynthesisVoice struct {
}

// Default prop
// js:"default"
func (*SpeechSynthesisVoice) Default() (def bool) {
	macro.Rewrite("$_.default")
	return def
}

// Lang prop
// js:"lang"
func (*SpeechSynthesisVoice) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

// LocalService prop
// js:"localService"
func (*SpeechSynthesisVoice) LocalService() (localService bool) {
	macro.Rewrite("$_.localService")
	return localService
}

// Name prop
// js:"name"
func (*SpeechSynthesisVoice) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// VoiceURI prop
// js:"voiceURI"
func (*SpeechSynthesisVoice) VoiceURI() (voiceURI string) {
	macro.Rewrite("$_.voiceURI")
	return voiceURI
}
