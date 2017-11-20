package videotrack

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/sourcebuffer"
	"github.com/matthewmueller/golly/js"
)

type VideoTrack struct {
}

func (*VideoTrack) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*VideoTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*VideoTrack) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

func (*VideoTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*VideoTrack) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*VideoTrack) SetLanguage(language string) {
	js.Rewrite("$<.language = $1", language)
}

func (*VideoTrack) GetSelected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

func (*VideoTrack) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

func (*VideoTrack) GetSourceBuffer() (sourceBuffer *sourcebuffer.SourceBuffer) {
	js.Rewrite("$<.sourceBuffer")
	return sourceBuffer
}
