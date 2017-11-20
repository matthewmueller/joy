package htmltrackelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrack"
	"github.com/matthewmueller/golly/js"
)

type HTMLTrackElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLTrackElement) GetDefault() (def bool) {
	js.Rewrite("$<.default")
	return def
}

func (*HTMLTrackElement) SetDefault(def bool) {
	js.Rewrite("$<.default = $1", def)
}

func (*HTMLTrackElement) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*HTMLTrackElement) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

func (*HTMLTrackElement) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*HTMLTrackElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

func (*HTMLTrackElement) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLTrackElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLTrackElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLTrackElement) GetSrclang() (srclang string) {
	js.Rewrite("$<.srclang")
	return srclang
}

func (*HTMLTrackElement) SetSrclang(srclang string) {
	js.Rewrite("$<.srclang = $1", srclang)
}

func (*HTMLTrackElement) GetTrack() (track *texttrack.TextTrack) {
	js.Rewrite("$<.track")
	return track
}
