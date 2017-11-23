package htmlsourceelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLSourceElement struct
// js:"HTMLSourceElement,omit"
type HTMLSourceElement struct {
	window.HTMLElement
}

// Media Gets or sets the intended media type of the media source.
func (*HTMLSourceElement) Media() (media string) {
	js.Rewrite("$<.Media")
	return media
}

// Media Gets or sets the intended media type of the media source.
func (*HTMLSourceElement) SetMedia(media string) {
	js.Rewrite("$<.Media = $1", media)
}

// MsKeySystem
func (*HTMLSourceElement) MsKeySystem() (msKeySystem string) {
	js.Rewrite("$<.MsKeySystem")
	return msKeySystem
}

// MsKeySystem
func (*HTMLSourceElement) SetMsKeySystem(msKeySystem string) {
	js.Rewrite("$<.MsKeySystem = $1", msKeySystem)
}

// Sizes
func (*HTMLSourceElement) Sizes() (sizes string) {
	js.Rewrite("$<.Sizes")
	return sizes
}

// Sizes
func (*HTMLSourceElement) SetSizes(sizes string) {
	js.Rewrite("$<.Sizes = $1", sizes)
}

// Src The address or URL of the a media resource that is to be considered.
func (*HTMLSourceElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src The address or URL of the a media resource that is to be considered.
func (*HTMLSourceElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Srcset
func (*HTMLSourceElement) Srcset() (srcset string) {
	js.Rewrite("$<.Srcset")
	return srcset
}

// Srcset
func (*HTMLSourceElement) SetSrcset(srcset string) {
	js.Rewrite("$<.Srcset = $1", srcset)
}

// Type Gets or sets the MIME type of a media resource.
func (*HTMLSourceElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Gets or sets the MIME type of a media resource.
func (*HTMLSourceElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
