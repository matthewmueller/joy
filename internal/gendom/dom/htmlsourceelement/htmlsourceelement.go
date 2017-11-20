package htmlsourceelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLSourceElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLSourceElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*HTMLSourceElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*HTMLSourceElement) GetMsKeySystem() (msKeySystem string) {
	js.Rewrite("$<.msKeySystem")
	return msKeySystem
}

func (*HTMLSourceElement) SetMsKeySystem(msKeySystem string) {
	js.Rewrite("$<.msKeySystem = $1", msKeySystem)
}

func (*HTMLSourceElement) GetSizes() (sizes string) {
	js.Rewrite("$<.sizes")
	return sizes
}

func (*HTMLSourceElement) SetSizes(sizes string) {
	js.Rewrite("$<.sizes = $1", sizes)
}

func (*HTMLSourceElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLSourceElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLSourceElement) GetSrcset() (srcset string) {
	js.Rewrite("$<.srcset")
	return srcset
}

func (*HTMLSourceElement) SetSrcset(srcset string) {
	js.Rewrite("$<.srcset = $1", srcset)
}

func (*HTMLSourceElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLSourceElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
