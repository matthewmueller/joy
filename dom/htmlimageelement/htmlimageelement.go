package htmlimageelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLImageElement struct
// js:"HTMLImageElement,omit"
type HTMLImageElement struct {
	window.HTMLElement
}

// MsGetAsCastingSource fn
func (*HTMLImageElement) MsGetAsCastingSource() (i interface{}) {
	js.Rewrite("$<.msGetAsCastingSource()")
	return i
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLImageElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLImageElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLImageElement) Alt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLImageElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

// Border prop Specifies the properties of a border drawn around an object.
func (*HTMLImageElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop Specifies the properties of a border drawn around an object.
func (*HTMLImageElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// Complete prop Retrieves whether the object is fully loaded.
func (*HTMLImageElement) Complete() (complete bool) {
	js.Rewrite("$<.complete")
	return complete
}

// CrossOrigin prop
func (*HTMLImageElement) CrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

// CrossOrigin prop
func (*HTMLImageElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

// CurrentSrc prop
func (*HTMLImageElement) CurrentSrc() (currentSrc string) {
	js.Rewrite("$<.currentSrc")
	return currentSrc
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLImageElement) Height() (height uint) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLImageElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

// Hspace prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLImageElement) Hspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

// Hspace prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLImageElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

// IsMap prop Sets or retrieves whether the image is a server-side image map.
func (*HTMLImageElement) IsMap() (isMap bool) {
	js.Rewrite("$<.isMap")
	return isMap
}

// IsMap prop Sets or retrieves whether the image is a server-side image map.
func (*HTMLImageElement) SetIsMap(isMap bool) {
	js.Rewrite("$<.isMap = $1", isMap)
}

// LongDesc prop Sets or retrieves a Uniform Resource Identifier (URI) to a long description of the object.
func (*HTMLImageElement) LongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

// LongDesc prop Sets or retrieves a Uniform Resource Identifier (URI) to a long description of the object.
func (*HTMLImageElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

// Lowsrc prop
func (*HTMLImageElement) Lowsrc() (lowsrc string) {
	js.Rewrite("$<.lowsrc")
	return lowsrc
}

// Lowsrc prop
func (*HTMLImageElement) SetLowsrc(lowsrc string) {
	js.Rewrite("$<.lowsrc = $1", lowsrc)
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLImageElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLImageElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI prop
func (*HTMLImageElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

// MsPlayToPreferredSourceURI prop
func (*HTMLImageElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
func (*HTMLImageElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
func (*HTMLImageElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
func (*HTMLImageElement) MsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLImageElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLImageElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// NaturalHeight prop The original height of the image resource before sizing.
func (*HTMLImageElement) NaturalHeight() (naturalHeight uint) {
	js.Rewrite("$<.naturalHeight")
	return naturalHeight
}

// NaturalWidth prop The original width of the image resource before sizing.
func (*HTMLImageElement) NaturalWidth() (naturalWidth uint) {
	js.Rewrite("$<.naturalWidth")
	return naturalWidth
}

// Sizes prop
func (*HTMLImageElement) Sizes() (sizes string) {
	js.Rewrite("$<.sizes")
	return sizes
}

// Sizes prop
func (*HTMLImageElement) SetSizes(sizes string) {
	js.Rewrite("$<.sizes = $1", sizes)
}

// Src prop The address or URL of the a media resource that is to be considered.
func (*HTMLImageElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop The address or URL of the a media resource that is to be considered.
func (*HTMLImageElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Srcset prop
func (*HTMLImageElement) Srcset() (srcset string) {
	js.Rewrite("$<.srcset")
	return srcset
}

// Srcset prop
func (*HTMLImageElement) SetSrcset(srcset string) {
	js.Rewrite("$<.srcset = $1", srcset)
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLImageElement) UseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLImageElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

// Vspace prop Sets or retrieves the vertical margin for the object.
func (*HTMLImageElement) Vspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

// Vspace prop Sets or retrieves the vertical margin for the object.
func (*HTMLImageElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLImageElement) Width() (width uint) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLImageElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}

// X prop
func (*HTMLImageElement) X() (x int) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*HTMLImageElement) Y() (y int) {
	js.Rewrite("$<.y")
	return y
}
