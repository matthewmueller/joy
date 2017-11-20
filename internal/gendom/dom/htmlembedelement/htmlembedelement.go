package htmlembedelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/getsvgdocument"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLEmbedElement struct {
	*htmlelement.HTMLElement
	*getsvgdocument.GetSVGDocument
}

func (*HTMLEmbedElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLEmbedElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLEmbedElement) GetHidden() (hidden string) {
	js.Rewrite("$<.hidden")
	return hidden
}

func (*HTMLEmbedElement) SetHidden(hidden string) {
	js.Rewrite("$<.hidden = $1", hidden)
}

func (*HTMLEmbedElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLEmbedElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLEmbedElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLEmbedElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLEmbedElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLEmbedElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLEmbedElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLEmbedElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLEmbedElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLEmbedElement) GetPalette() (palette string) {
	js.Rewrite("$<.palette")
	return palette
}

func (*HTMLEmbedElement) GetPluginspage() (pluginspage string) {
	js.Rewrite("$<.pluginspage")
	return pluginspage
}

func (*HTMLEmbedElement) GetReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLEmbedElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLEmbedElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLEmbedElement) GetUnits() (units string) {
	js.Rewrite("$<.units")
	return units
}

func (*HTMLEmbedElement) SetUnits(units string) {
	js.Rewrite("$<.units = $1", units)
}

func (*HTMLEmbedElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLEmbedElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
