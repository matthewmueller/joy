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

// GetSVGDocument fn
func (*HTMLEmbedElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.getSVGDocument()")
	return w
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLEmbedElement) Height() (height string) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLEmbedElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

// Hidden prop
func (*HTMLEmbedElement) Hidden() (hidden string) {
	js.Rewrite("$<.hidden")
	return hidden
}

// Hidden prop
func (*HTMLEmbedElement) SetHidden(hidden string) {
	js.Rewrite("$<.hidden = $1", hidden)
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLEmbedElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLEmbedElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLEmbedElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLEmbedElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
func (*HTMLEmbedElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
func (*HTMLEmbedElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
func (*HTMLEmbedElement) MsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLEmbedElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLEmbedElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Palette prop Retrieves the palette used for the embedded document.
func (*HTMLEmbedElement) Palette() (palette string) {
	js.Rewrite("$<.palette")
	return palette
}

// Pluginspage prop Retrieves the URL of the plug-in used to view an embedded document.
func (*HTMLEmbedElement) Pluginspage() (pluginspage string) {
	js.Rewrite("$<.pluginspage")
	return pluginspage
}

// ReadyState prop
func (*HTMLEmbedElement) ReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Src prop Sets or retrieves a URL to be loaded by the object.
func (*HTMLEmbedElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop Sets or retrieves a URL to be loaded by the object.
func (*HTMLEmbedElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Units prop Sets or retrieves the height and width units of the embed object.
func (*HTMLEmbedElement) Units() (units string) {
	js.Rewrite("$<.units")
	return units
}

// Units prop Sets or retrieves the height and width units of the embed object.
func (*HTMLEmbedElement) SetUnits(units string) {
	js.Rewrite("$<.units = $1", units)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLEmbedElement) Width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLEmbedElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
