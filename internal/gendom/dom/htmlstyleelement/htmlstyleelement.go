package htmlstyleelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/linkstyle"
	"github.com/matthewmueller/golly/js"
)

type HTMLStyleElement struct {
	*htmlelement.HTMLElement
	*linkstyle.LinkStyle
}

func (*HTMLStyleElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLStyleElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*HTMLStyleElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*HTMLStyleElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLStyleElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
