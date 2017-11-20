package htmlappletelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLAppletElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLAppletElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLAppletElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLAppletElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLAppletElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLAppletElement) GetAltHTML() (altHtml string) {
	js.Rewrite("$<.altHtml")
	return altHtml
}

func (*HTMLAppletElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.altHtml = $1", altHtml)
}

func (*HTMLAppletElement) GetArchive() (archive string) {
	js.Rewrite("$<.archive")
	return archive
}

func (*HTMLAppletElement) SetArchive(archive string) {
	js.Rewrite("$<.archive = $1", archive)
}

func (*HTMLAppletElement) GetBaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

func (*HTMLAppletElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLAppletElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLAppletElement) GetCode() (code string) {
	js.Rewrite("$<.code")
	return code
}

func (*HTMLAppletElement) SetCode(code string) {
	js.Rewrite("$<.code = $1", code)
}

func (*HTMLAppletElement) GetCodeBase() (codeBase string) {
	js.Rewrite("$<.codeBase")
	return codeBase
}

func (*HTMLAppletElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.codeBase = $1", codeBase)
}

func (*HTMLAppletElement) GetCodeType() (codeType string) {
	js.Rewrite("$<.codeType")
	return codeType
}

func (*HTMLAppletElement) SetCodeType(codeType string) {
	js.Rewrite("$<.codeType = $1", codeType)
}

func (*HTMLAppletElement) GetContentDocument() (contentDocument *document.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLAppletElement) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*HTMLAppletElement) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

func (*HTMLAppletElement) GetDeclare() (declare bool) {
	js.Rewrite("$<.declare")
	return declare
}

func (*HTMLAppletElement) SetDeclare(declare bool) {
	js.Rewrite("$<.declare = $1", declare)
}

func (*HTMLAppletElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLAppletElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLAppletElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLAppletElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLAppletElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLAppletElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLAppletElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLAppletElement) GetObject() (object string) {
	js.Rewrite("$<.object")
	return object
}

func (*HTMLAppletElement) SetObject(object string) {
	js.Rewrite("$<.object = $1", object)
}

func (*HTMLAppletElement) GetStandby() (standby string) {
	js.Rewrite("$<.standby")
	return standby
}

func (*HTMLAppletElement) SetStandby(standby string) {
	js.Rewrite("$<.standby = $1", standby)
}

func (*HTMLAppletElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLAppletElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLAppletElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLAppletElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLAppletElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLAppletElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLAppletElement) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLAppletElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}
