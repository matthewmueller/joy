package htmlsourceelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLSourceElement struct
// js:"HTMLSourceElement,omit"
type HTMLSourceElement struct {
	window.HTMLElement
}

// Media prop Gets or sets the intended media type of the media source.
func (*HTMLSourceElement) Media() (media string) {
	js.Rewrite("$<.media")
	return media
}

// Media prop Gets or sets the intended media type of the media source.
func (*HTMLSourceElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

// MsKeySystem prop
func (*HTMLSourceElement) MsKeySystem() (msKeySystem string) {
	js.Rewrite("$<.msKeySystem")
	return msKeySystem
}

// MsKeySystem prop
func (*HTMLSourceElement) SetMsKeySystem(msKeySystem string) {
	js.Rewrite("$<.msKeySystem = $1", msKeySystem)
}

// Sizes prop
func (*HTMLSourceElement) Sizes() (sizes string) {
	js.Rewrite("$<.sizes")
	return sizes
}

// Sizes prop
func (*HTMLSourceElement) SetSizes(sizes string) {
	js.Rewrite("$<.sizes = $1", sizes)
}

// Src prop The address or URL of the a media resource that is to be considered.
func (*HTMLSourceElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop The address or URL of the a media resource that is to be considered.
func (*HTMLSourceElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Srcset prop
func (*HTMLSourceElement) Srcset() (srcset string) {
	js.Rewrite("$<.srcset")
	return srcset
}

// Srcset prop
func (*HTMLSourceElement) SetSrcset(srcset string) {
	js.Rewrite("$<.srcset = $1", srcset)
}

// Type prop Gets or sets the MIME type of a media resource.
func (*HTMLSourceElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Gets or sets the MIME type of a media resource.
func (*HTMLSourceElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
