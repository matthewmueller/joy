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

// MsGetAsCastingSource
func (*HTMLImageElement) MsGetAsCastingSource() (i interface{}) {
	js.Rewrite("$<.MsGetAsCastingSource()")
	return i
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLImageElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLImageElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLImageElement) Alt() (alt string) {
	js.Rewrite("$<.Alt")
	return alt
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLImageElement) SetAlt(alt string) {
	js.Rewrite("$<.Alt = $1", alt)
}

// Border Specifies the properties of a border drawn around an object.
func (*HTMLImageElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border Specifies the properties of a border drawn around an object.
func (*HTMLImageElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// Complete Retrieves whether the object is fully loaded.
func (*HTMLImageElement) Complete() (complete bool) {
	js.Rewrite("$<.Complete")
	return complete
}

// CrossOrigin
func (*HTMLImageElement) CrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.CrossOrigin")
	return crossOrigin
}

// CrossOrigin
func (*HTMLImageElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.CrossOrigin = $1", crossOrigin)
}

// CurrentSrc
func (*HTMLImageElement) CurrentSrc() (currentSrc string) {
	js.Rewrite("$<.CurrentSrc")
	return currentSrc
}

// Height Sets or retrieves the height of the object.
func (*HTMLImageElement) Height() (height uint) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLImageElement) SetHeight(height uint) {
	js.Rewrite("$<.Height = $1", height)
}

// Hspace Sets or retrieves the width of the border to draw around the object.
func (*HTMLImageElement) Hspace() (hspace int) {
	js.Rewrite("$<.Hspace")
	return hspace
}

// Hspace Sets or retrieves the width of the border to draw around the object.
func (*HTMLImageElement) SetHspace(hspace int) {
	js.Rewrite("$<.Hspace = $1", hspace)
}

// IsMap Sets or retrieves whether the image is a server-side image map.
func (*HTMLImageElement) IsMap() (isMap bool) {
	js.Rewrite("$<.IsMap")
	return isMap
}

// IsMap Sets or retrieves whether the image is a server-side image map.
func (*HTMLImageElement) SetIsMap(isMap bool) {
	js.Rewrite("$<.IsMap = $1", isMap)
}

// LongDesc Sets or retrieves a Uniform Resource Identifier (URI) to a long description of the object.
func (*HTMLImageElement) LongDesc() (longDesc string) {
	js.Rewrite("$<.LongDesc")
	return longDesc
}

// LongDesc Sets or retrieves a Uniform Resource Identifier (URI) to a long description of the object.
func (*HTMLImageElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.LongDesc = $1", longDesc)
}

// Lowsrc
func (*HTMLImageElement) Lowsrc() (lowsrc string) {
	js.Rewrite("$<.Lowsrc")
	return lowsrc
}

// Lowsrc
func (*HTMLImageElement) SetLowsrc(lowsrc string) {
	js.Rewrite("$<.Lowsrc = $1", lowsrc)
}

// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLImageElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.MsPlayToDisabled")
	return msPlayToDisabled
}

// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLImageElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.MsPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI
func (*HTMLImageElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.MsPlayToPreferredSourceURI")
	return msPlayToPreferredSourceUri
}

// MsPlayToPreferredSourceURI
func (*HTMLImageElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.MsPlayToPreferredSourceURI = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
func (*HTMLImageElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.MsPlayToPrimary")
	return msPlayToPrimary
}

// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
func (*HTMLImageElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.MsPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource Gets the source associated with the media element for use by the PlayToManager.
func (*HTMLImageElement) MsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.MsPlayToSource")
	return msPlayToSource
}

// Name Sets or retrieves the name of the object.
func (*HTMLImageElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLImageElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// NaturalHeight The original height of the image resource before sizing.
func (*HTMLImageElement) NaturalHeight() (naturalHeight uint) {
	js.Rewrite("$<.NaturalHeight")
	return naturalHeight
}

// NaturalWidth The original width of the image resource before sizing.
func (*HTMLImageElement) NaturalWidth() (naturalWidth uint) {
	js.Rewrite("$<.NaturalWidth")
	return naturalWidth
}

// Sizes
func (*HTMLImageElement) Sizes() (sizes string) {
	js.Rewrite("$<.Sizes")
	return sizes
}

// Sizes
func (*HTMLImageElement) SetSizes(sizes string) {
	js.Rewrite("$<.Sizes = $1", sizes)
}

// Src The address or URL of the a media resource that is to be considered.
func (*HTMLImageElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src The address or URL of the a media resource that is to be considered.
func (*HTMLImageElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Srcset
func (*HTMLImageElement) Srcset() (srcset string) {
	js.Rewrite("$<.Srcset")
	return srcset
}

// Srcset
func (*HTMLImageElement) SetSrcset(srcset string) {
	js.Rewrite("$<.Srcset = $1", srcset)
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLImageElement) UseMap() (useMap string) {
	js.Rewrite("$<.UseMap")
	return useMap
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLImageElement) SetUseMap(useMap string) {
	js.Rewrite("$<.UseMap = $1", useMap)
}

// Vspace Sets or retrieves the vertical margin for the object.
func (*HTMLImageElement) Vspace() (vspace int) {
	js.Rewrite("$<.Vspace")
	return vspace
}

// Vspace Sets or retrieves the vertical margin for the object.
func (*HTMLImageElement) SetVspace(vspace int) {
	js.Rewrite("$<.Vspace = $1", vspace)
}

// Width Sets or retrieves the width of the object.
func (*HTMLImageElement) Width() (width uint) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLImageElement) SetWidth(width uint) {
	js.Rewrite("$<.Width = $1", width)
}

// X
func (*HTMLImageElement) X() (x int) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*HTMLImageElement) Y() (y int) {
	js.Rewrite("$<.Y")
	return y
}
