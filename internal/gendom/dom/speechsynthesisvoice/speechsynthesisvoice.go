package speechsynthesisvoice

import "github.com/matthewmueller/golly/js"

type SpeechSynthesisVoice struct {
}

func (*SpeechSynthesisVoice) GetDefault() (def bool) {
	js.Rewrite("$<.default")
	return def
}

func (*SpeechSynthesisVoice) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*SpeechSynthesisVoice) GetLocalService() (localService bool) {
	js.Rewrite("$<.localService")
	return localService
}

func (*SpeechSynthesisVoice) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*SpeechSynthesisVoice) GetVoiceURI() (voiceURI string) {
	js.Rewrite("$<.voiceURI")
	return voiceURI
}
