package htmlobjectelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/getsvgdocument"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/validitystate"
	"github.com/matthewmueller/golly/js"
)

type HTMLObjectElement struct {
	*htmlelement.HTMLElement
	*getsvgdocument.GetSVGDocument
}

func (*HTMLObjectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLObjectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLObjectElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLObjectElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLObjectElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLObjectElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLObjectElement) GetAltHTML() (altHtml string) {
	js.Rewrite("$<.altHtml")
	return altHtml
}

func (*HTMLObjectElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.altHtml = $1", altHtml)
}

func (*HTMLObjectElement) GetArchive() (archive string) {
	js.Rewrite("$<.archive")
	return archive
}

func (*HTMLObjectElement) SetArchive(archive string) {
	js.Rewrite("$<.archive = $1", archive)
}

func (*HTMLObjectElement) GetBaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

func (*HTMLObjectElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLObjectElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLObjectElement) GetCode() (code string) {
	js.Rewrite("$<.code")
	return code
}

func (*HTMLObjectElement) SetCode(code string) {
	js.Rewrite("$<.code = $1", code)
}

func (*HTMLObjectElement) GetCodeBase() (codeBase string) {
	js.Rewrite("$<.codeBase")
	return codeBase
}

func (*HTMLObjectElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.codeBase = $1", codeBase)
}

func (*HTMLObjectElement) GetCodeType() (codeType string) {
	js.Rewrite("$<.codeType")
	return codeType
}

func (*HTMLObjectElement) SetCodeType(codeType string) {
	js.Rewrite("$<.codeType = $1", codeType)
}

func (*HTMLObjectElement) GetContentDocument() (contentDocument *document.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLObjectElement) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*HTMLObjectElement) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

func (*HTMLObjectElement) GetDeclare() (declare bool) {
	js.Rewrite("$<.declare")
	return declare
}

func (*HTMLObjectElement) SetDeclare(declare bool) {
	js.Rewrite("$<.declare = $1", declare)
}

func (*HTMLObjectElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLObjectElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLObjectElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLObjectElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLObjectElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLObjectElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLObjectElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLObjectElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLObjectElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLObjectElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLObjectElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLObjectElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLObjectElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLObjectElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLObjectElement) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLObjectElement) GetStandby() (standby string) {
	js.Rewrite("$<.standby")
	return standby
}

func (*HTMLObjectElement) SetStandby(standby string) {
	js.Rewrite("$<.standby = $1", standby)
}

func (*HTMLObjectElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLObjectElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLObjectElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLObjectElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLObjectElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLObjectElement) GetValidity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLObjectElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLObjectElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLObjectElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLObjectElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

func (*HTMLObjectElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
