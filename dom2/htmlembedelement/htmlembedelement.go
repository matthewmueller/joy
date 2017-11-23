package htmlembedelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLEmbedElement struct
// js:"HTMLEmbedElement,omit"
type HTMLEmbedElement struct {
	window.HTMLElement
}

// GetSVGDocument
func (*HTMLEmbedElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.GetSVGDocument()")
	return w
}

// Height Sets or retrieves the height of the object.
func (*HTMLEmbedElement) Height() (height string) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLEmbedElement) SetHeight(height string) {
	js.Rewrite("$<.Height = $1", height)
}

// Hidden
func (*HTMLEmbedElement) Hidden() (hidden string) {
	js.Rewrite("$<.Hidden")
	return hidden
}

// Hidden
func (*HTMLEmbedElement) SetHidden(hidden string) {
	js.Rewrite("$<.Hidden = $1", hidden)
}

// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLEmbedElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.MsPlayToDisabled")
	return msPlayToDisabled
}

// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLEmbedElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.MsPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLEmbedElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.MsPlayToPreferredSourceURI")
	return msPlayToPreferredSourceUri
}

// MsPlayToPreferredSourceURI Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLEmbedElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.MsPlayToPreferredSourceURI = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
func (*HTMLEmbedElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.MsPlayToPrimary")
	return msPlayToPrimary
}

// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
func (*HTMLEmbedElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.MsPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource Gets the source associated with the media element for use by the PlayToManager.
func (*HTMLEmbedElement) MsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.MsPlayToSource")
	return msPlayToSource
}

// Name Sets or retrieves the name of the object.
func (*HTMLEmbedElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLEmbedElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Palette Retrieves the palette used for the embedded document.
func (*HTMLEmbedElement) Palette() (palette string) {
	js.Rewrite("$<.Palette")
	return palette
}

// Pluginspage Retrieves the URL of the plug-in used to view an embedded document.
func (*HTMLEmbedElement) Pluginspage() (pluginspage string) {
	js.Rewrite("$<.Pluginspage")
	return pluginspage
}

// ReadyState
func (*HTMLEmbedElement) ReadyState() (readyState string) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Src Sets or retrieves a URL to be loaded by the object.
func (*HTMLEmbedElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src Sets or retrieves a URL to be loaded by the object.
func (*HTMLEmbedElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Units Sets or retrieves the height and width units of the embed object.
func (*HTMLEmbedElement) Units() (units string) {
	js.Rewrite("$<.Units")
	return units
}

// Units Sets or retrieves the height and width units of the embed object.
func (*HTMLEmbedElement) SetUnits(units string) {
	js.Rewrite("$<.Units = $1", units)
}

// Width Sets or retrieves the width of the object.
func (*HTMLEmbedElement) Width() (width string) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLEmbedElement) SetWidth(width string) {
	js.Rewrite("$<.Width = $1", width)
}
