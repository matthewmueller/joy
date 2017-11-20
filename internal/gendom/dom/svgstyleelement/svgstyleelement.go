package svgstyleelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGStyleElement struct {
	*svgelement.SVGElement
}

func (*SVGStyleElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*SVGStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*SVGStyleElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*SVGStyleElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*SVGStyleElement) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*SVGStyleElement) SetTitle(title string) {
	js.Rewrite("$<.title = $1", title)
}

func (*SVGStyleElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*SVGStyleElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
