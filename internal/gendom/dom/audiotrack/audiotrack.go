package audiotrack

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/sourcebuffer"
	"github.com/matthewmueller/golly/js"
)

type AudioTrack struct {
}

func (*AudioTrack) GetEnabled() (enabled bool) {
	js.Rewrite("$<.enabled")
	return enabled
}

func (*AudioTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.enabled = $1", enabled)
}

func (*AudioTrack) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*AudioTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*AudioTrack) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

func (*AudioTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*AudioTrack) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*AudioTrack) SetLanguage(language string) {
	js.Rewrite("$<.language = $1", language)
}

func (*AudioTrack) GetSourceBuffer() (sourceBuffer *sourcebuffer.SourceBuffer) {
	js.Rewrite("$<.sourceBuffer")
	return sourceBuffer
}
