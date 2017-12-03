package htmltableelement

import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/domstringmap"
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/dom/htmltablecaptionelement"
	"github.com/matthewmueller/joy/dom/htmltablesectionelement"
	"github.com/matthewmueller/joy/dom/mszoomtooptions"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.HTMLElement = (*HTMLTableElement)(nil)
var _ window.Element = (*HTMLTableElement)(nil)
var _ window.GlobalEventHandlers = (*HTMLTableElement)(nil)
var _ window.ElementTraversal = (*HTMLTableElement)(nil)
var _ window.NodeSelector = (*HTMLTableElement)(nil)
var _ childnode.ChildNode = (*HTMLTableElement)(nil)
var _ window.Node = (*HTMLTableElement)(nil)
var _ window.EventTarget = (*HTMLTableElement)(nil)

// HTMLTableElement struct
// js:"HTMLTableElement,omit"
type HTMLTableElement struct {
}

// CreateCaption fn Creates an empty caption element in the table.
// js:"createCaption"
func (*HTMLTableElement) CreateCaption() (w window.HTMLElement) {
	js.Rewrite("$_.createCaption()")
	return w
}

// CreateTBody fn Creates an empty tBody element in the table.
// js:"createTBody"
func (*HTMLTableElement) CreateTBody() (w window.HTMLElement) {
	js.Rewrite("$_.createTBody()")
	return w
}

// CreateTFoot fn Creates an empty tFoot element in the table.
// js:"createTFoot"
func (*HTMLTableElement) CreateTFoot() (w window.HTMLElement) {
	js.Rewrite("$_.createTFoot()")
	return w
}

// CreateTHead fn Returns the tHead element object if successful, or null otherwise.
// js:"createTHead"
func (*HTMLTableElement) CreateTHead() (w window.HTMLElement) {
	js.Rewrite("$_.createTHead()")
	return w
}

// DeleteCaption fn Deletes the caption element and its contents from the table.
// js:"deleteCaption"
func (*HTMLTableElement) DeleteCaption() {
	js.Rewrite("$_.deleteCaption()")
}

// DeleteRow fn Removes the specified row (tr) from the element and from the rows collection.
//     * @param index Number that specifies the zero-based position in the rows collection of the row to remove.
// js:"deleteRow"
func (*HTMLTableElement) DeleteRow(index *int) {
	js.Rewrite("$_.deleteRow($1)", index)
}

// DeleteTFoot fn Deletes the tFoot element and its contents from the table.
// js:"deleteTFoot"
func (*HTMLTableElement) DeleteTFoot() {
	js.Rewrite("$_.deleteTFoot()")
}

// DeleteTHead fn Deletes the tHead element and its contents from the table.
// js:"deleteTHead"
func (*HTMLTableElement) DeleteTHead() {
	js.Rewrite("$_.deleteTHead()")
}

// InsertRow fn Creates a new row (tr) in the table, and adds the row to the rows collection.
//     * @param index Number that specifies where to insert the row in the rows collection. The default value is -1, which appends the new row to the end of the rows collection.
// js:"insertRow"
func (*HTMLTableElement) InsertRow(index *int) (w window.HTMLElement) {
	js.Rewrite("$_.insertRow($1)", index)
	return w
}

// Blur fn
// js:"blur"
func (*HTMLTableElement) Blur() {
	js.Rewrite("$_.blur()")
}

// Click fn
// js:"click"
func (*HTMLTableElement) Click() {
	js.Rewrite("$_.click()")
}

// DragDrop fn
// js:"dragDrop"
func (*HTMLTableElement) DragDrop() (b bool) {
	js.Rewrite("$_.dragDrop()")
	return b
}

// Focus fn
// js:"focus"
func (*HTMLTableElement) Focus() {
	js.Rewrite("$_.focus()")
}

// GetElementsByClassName fn
// js:"getElementsByClassName"
func (*HTMLTableElement) GetElementsByClassName(classNames string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

// InsertAdjacentElement fn
// js:"insertAdjacentElement"
func (*HTMLTableElement) InsertAdjacentElement(position string, insertedElement window.Element) (w window.Element) {
	js.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

// InsertAdjacentHTML fn
// js:"insertAdjacentHTML"
func (*HTMLTableElement) InsertAdjacentHTML(where string, html string) {
	js.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

// InsertAdjacentText fn
// js:"insertAdjacentText"
func (*HTMLTableElement) InsertAdjacentText(where string, text string) {
	js.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

// MsGetInputContext fn
// js:"msGetInputContext"
func (*HTMLTableElement) MsGetInputContext() (w *window.MSInputMethodContext) {
	js.Rewrite("$_.msGetInputContext()")
	return w
}

// ScrollIntoView fn
// js:"scrollIntoView"
func (*HTMLTableElement) ScrollIntoView(top *bool) {
	js.Rewrite("$_.scrollIntoView($1)", top)
}

// GetAttribute fn
// js:"getAttribute"
func (*HTMLTableElement) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*HTMLTableElement) GetAttributeNode(name string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*HTMLTableElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*HTMLTableElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*HTMLTableElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*HTMLTableElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*HTMLTableElement) GetElementsByTagName(name string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*HTMLTableElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*HTMLTableElement) HasAttribute(name string) (b bool) {
	js.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*HTMLTableElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*HTMLTableElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	js.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*HTMLTableElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*HTMLTableElement) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*HTMLTableElement) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*HTMLTableElement) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*HTMLTableElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*HTMLTableElement) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*HTMLTableElement) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*HTMLTableElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*HTMLTableElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*HTMLTableElement) RequestFullscreen() {
	js.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*HTMLTableElement) RequestPointerLock() {
	js.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*HTMLTableElement) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*HTMLTableElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*HTMLTableElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*HTMLTableElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*HTMLTableElement) SetPointerCapture(pointerId int) {
	js.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*HTMLTableElement) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*HTMLTableElement) WebkitRequestFullscreen() {
	js.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*HTMLTableElement) WebkitRequestFullScreen() {
	js.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLTableElement) QuerySelector(selectors string) (w window.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*HTMLTableElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*HTMLTableElement) AppendChild(newChild window.Node) (w window.Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*HTMLTableElement) CloneNode(deep *bool) (w window.Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*HTMLTableElement) CompareDocumentPosition(other window.Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*HTMLTableElement) Contains(child window.Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*HTMLTableElement) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*HTMLTableElement) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*HTMLTableElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*HTMLTableElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*HTMLTableElement) IsEqualNode(arg window.Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*HTMLTableElement) IsSameNode(other window.Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*HTMLTableElement) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*HTMLTableElement) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*HTMLTableElement) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*HTMLTableElement) RemoveChild(oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*HTMLTableElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLTableElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLTableElement) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*HTMLTableElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Align prop Sets or retrieves a value that indicates the table alignment.
// js:"align"
func (*HTMLTableElement) Align() (align string) {
	js.Rewrite("$_.align")
	return align
}

// SetAlign prop Sets or retrieves a value that indicates the table alignment.
// js:"align"
func (*HTMLTableElement) SetAlign(align string) {
	js.Rewrite("$_.align = $1", align)
}

// BgColor prop
// js:"bgColor"
func (*HTMLTableElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$_.bgColor")
	return bgColor
}

// SetBgColor prop
// js:"bgColor"
func (*HTMLTableElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$_.bgColor = $1", bgColor)
}

// Border prop Sets or retrieves the width of the border to draw around the object.
// js:"border"
func (*HTMLTableElement) Border() (border string) {
	js.Rewrite("$_.border")
	return border
}

// SetBorder prop Sets or retrieves the width of the border to draw around the object.
// js:"border"
func (*HTMLTableElement) SetBorder(border string) {
	js.Rewrite("$_.border = $1", border)
}

// BorderColor prop Sets or retrieves the border color of the object.
// js:"borderColor"
func (*HTMLTableElement) BorderColor() (borderColor interface{}) {
	js.Rewrite("$_.borderColor")
	return borderColor
}

// SetBorderColor prop Sets or retrieves the border color of the object.
// js:"borderColor"
func (*HTMLTableElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$_.borderColor = $1", borderColor)
}

// Caption prop Retrieves the caption object of a table.
// js:"caption"
func (*HTMLTableElement) Caption() (caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$_.caption")
	return caption
}

// SetCaption prop Retrieves the caption object of a table.
// js:"caption"
func (*HTMLTableElement) SetCaption(caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$_.caption = $1", caption)
}

// CellPadding prop Sets or retrieves the amount of space between the border of the cell and the content of the cell.
// js:"cellPadding"
func (*HTMLTableElement) CellPadding() (cellPadding string) {
	js.Rewrite("$_.cellPadding")
	return cellPadding
}

// SetCellPadding prop Sets or retrieves the amount of space between the border of the cell and the content of the cell.
// js:"cellPadding"
func (*HTMLTableElement) SetCellPadding(cellPadding string) {
	js.Rewrite("$_.cellPadding = $1", cellPadding)
}

// CellSpacing prop Sets or retrieves the amount of space between cells in a table.
// js:"cellSpacing"
func (*HTMLTableElement) CellSpacing() (cellSpacing string) {
	js.Rewrite("$_.cellSpacing")
	return cellSpacing
}

// SetCellSpacing prop Sets or retrieves the amount of space between cells in a table.
// js:"cellSpacing"
func (*HTMLTableElement) SetCellSpacing(cellSpacing string) {
	js.Rewrite("$_.cellSpacing = $1", cellSpacing)
}

// Cols prop Sets or retrieves the number of columns in the table.
// js:"cols"
func (*HTMLTableElement) Cols() (cols int) {
	js.Rewrite("$_.cols")
	return cols
}

// SetCols prop Sets or retrieves the number of columns in the table.
// js:"cols"
func (*HTMLTableElement) SetCols(cols int) {
	js.Rewrite("$_.cols = $1", cols)
}

// Frame prop Sets or retrieves the way the border frame around the table is displayed.
// js:"frame"
func (*HTMLTableElement) Frame() (frame string) {
	js.Rewrite("$_.frame")
	return frame
}

// SetFrame prop Sets or retrieves the way the border frame around the table is displayed.
// js:"frame"
func (*HTMLTableElement) SetFrame(frame string) {
	js.Rewrite("$_.frame = $1", frame)
}

// Height prop Sets or retrieves the height of the object.
// js:"height"
func (*HTMLTableElement) Height() (height interface{}) {
	js.Rewrite("$_.height")
	return height
}

// SetHeight prop Sets or retrieves the height of the object.
// js:"height"
func (*HTMLTableElement) SetHeight(height interface{}) {
	js.Rewrite("$_.height = $1", height)
}

// Rows prop Sets or retrieves the number of horizontal rows contained in the object.
// js:"rows"
func (*HTMLTableElement) Rows() (rows window.HTMLCollection) {
	js.Rewrite("$_.rows")
	return rows
}

// Rules prop Sets or retrieves which dividing lines (inner borders) are displayed.
// js:"rules"
func (*HTMLTableElement) Rules() (rules string) {
	js.Rewrite("$_.rules")
	return rules
}

// SetRules prop Sets or retrieves which dividing lines (inner borders) are displayed.
// js:"rules"
func (*HTMLTableElement) SetRules(rules string) {
	js.Rewrite("$_.rules = $1", rules)
}

// Summary prop Sets or retrieves a description and/or structure of the object.
// js:"summary"
func (*HTMLTableElement) Summary() (summary string) {
	js.Rewrite("$_.summary")
	return summary
}

// SetSummary prop Sets or retrieves a description and/or structure of the object.
// js:"summary"
func (*HTMLTableElement) SetSummary(summary string) {
	js.Rewrite("$_.summary = $1", summary)
}

// TBodies prop Retrieves a collection of all tBody objects in the table. Objects in this collection are in source order.
// js:"tBodies"
func (*HTMLTableElement) TBodies() (tBodies window.HTMLCollection) {
	js.Rewrite("$_.tBodies")
	return tBodies
}

// TFoot prop Retrieves the tFoot object of the table.
// js:"tFoot"
func (*HTMLTableElement) TFoot() (tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$_.tFoot")
	return tFoot
}

// SetTFoot prop Retrieves the tFoot object of the table.
// js:"tFoot"
func (*HTMLTableElement) SetTFoot(tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$_.tFoot = $1", tFoot)
}

// THead prop Retrieves the tHead object of the table.
// js:"tHead"
func (*HTMLTableElement) THead() (tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$_.tHead")
	return tHead
}

// SetTHead prop Retrieves the tHead object of the table.
// js:"tHead"
func (*HTMLTableElement) SetTHead(tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$_.tHead = $1", tHead)
}

// Width prop Sets or retrieves the width of the object.
// js:"width"
func (*HTMLTableElement) Width() (width string) {
	js.Rewrite("$_.width")
	return width
}

// SetWidth prop Sets or retrieves the width of the object.
// js:"width"
func (*HTMLTableElement) SetWidth(width string) {
	js.Rewrite("$_.width = $1", width)
}

// AccessKey prop
// js:"accessKey"
func (*HTMLTableElement) AccessKey() (accessKey string) {
	js.Rewrite("$_.accessKey")
	return accessKey
}

// SetAccessKey prop
// js:"accessKey"
func (*HTMLTableElement) SetAccessKey(accessKey string) {
	js.Rewrite("$_.accessKey = $1", accessKey)
}

// Children prop
// js:"children"
func (*HTMLTableElement) Children() (children window.HTMLCollection) {
	js.Rewrite("$_.children")
	return children
}

// ContentEditable prop
// js:"contentEditable"
func (*HTMLTableElement) ContentEditable() (contentEditable string) {
	js.Rewrite("$_.contentEditable")
	return contentEditable
}

// SetContentEditable prop
// js:"contentEditable"
func (*HTMLTableElement) SetContentEditable(contentEditable string) {
	js.Rewrite("$_.contentEditable = $1", contentEditable)
}

// Dataset prop
// js:"dataset"
func (*HTMLTableElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$_.dataset")
	return dataset
}

// Dir prop
// js:"dir"
func (*HTMLTableElement) Dir() (dir string) {
	js.Rewrite("$_.dir")
	return dir
}

// SetDir prop
// js:"dir"
func (*HTMLTableElement) SetDir(dir string) {
	js.Rewrite("$_.dir = $1", dir)
}

// Draggable prop
// js:"draggable"
func (*HTMLTableElement) Draggable() (draggable bool) {
	js.Rewrite("$_.draggable")
	return draggable
}

// SetDraggable prop
// js:"draggable"
func (*HTMLTableElement) SetDraggable(draggable bool) {
	js.Rewrite("$_.draggable = $1", draggable)
}

// Hidden prop
// js:"hidden"
func (*HTMLTableElement) Hidden() (hidden bool) {
	js.Rewrite("$_.hidden")
	return hidden
}

// SetHidden prop
// js:"hidden"
func (*HTMLTableElement) SetHidden(hidden bool) {
	js.Rewrite("$_.hidden = $1", hidden)
}

// HideFocus prop
// js:"hideFocus"
func (*HTMLTableElement) HideFocus() (hideFocus bool) {
	js.Rewrite("$_.hideFocus")
	return hideFocus
}

// SetHideFocus prop
// js:"hideFocus"
func (*HTMLTableElement) SetHideFocus(hideFocus bool) {
	js.Rewrite("$_.hideFocus = $1", hideFocus)
}

// InnerText prop
// js:"innerText"
func (*HTMLTableElement) InnerText() (innerText string) {
	js.Rewrite("$_.innerText")
	return innerText
}

// SetInnerText prop
// js:"innerText"
func (*HTMLTableElement) SetInnerText(innerText string) {
	js.Rewrite("$_.innerText = $1", innerText)
}

// IsContentEditable prop
// js:"isContentEditable"
func (*HTMLTableElement) IsContentEditable() (isContentEditable bool) {
	js.Rewrite("$_.isContentEditable")
	return isContentEditable
}

// Lang prop
// js:"lang"
func (*HTMLTableElement) Lang() (lang string) {
	js.Rewrite("$_.lang")
	return lang
}

// SetLang prop
// js:"lang"
func (*HTMLTableElement) SetLang(lang string) {
	js.Rewrite("$_.lang = $1", lang)
}

// OffsetHeight prop
// js:"offsetHeight"
func (*HTMLTableElement) OffsetHeight() (offsetHeight int) {
	js.Rewrite("$_.offsetHeight")
	return offsetHeight
}

// OffsetLeft prop
// js:"offsetLeft"
func (*HTMLTableElement) OffsetLeft() (offsetLeft int) {
	js.Rewrite("$_.offsetLeft")
	return offsetLeft
}

// OffsetParent prop
// js:"offsetParent"
func (*HTMLTableElement) OffsetParent() (offsetParent window.Element) {
	js.Rewrite("$_.offsetParent")
	return offsetParent
}

// OffsetTop prop
// js:"offsetTop"
func (*HTMLTableElement) OffsetTop() (offsetTop int) {
	js.Rewrite("$_.offsetTop")
	return offsetTop
}

// OffsetWidth prop
// js:"offsetWidth"
func (*HTMLTableElement) OffsetWidth() (offsetWidth int) {
	js.Rewrite("$_.offsetWidth")
	return offsetWidth
}

// Onabort prop
// js:"onabort"
func (*HTMLTableElement) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*HTMLTableElement) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onactivate prop
// js:"onactivate"
func (*HTMLTableElement) Onactivate() (onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate")
	return onactivate
}

// SetOnactivate prop
// js:"onactivate"
func (*HTMLTableElement) SetOnactivate(onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate = $1", onactivate)
}

// Onbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLTableElement) Onbeforeactivate() (onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

// SetOnbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLTableElement) SetOnbeforeactivate(onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

// Onbeforecopy prop
// js:"onbeforecopy"
func (*HTMLTableElement) Onbeforecopy() (onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

// SetOnbeforecopy prop
// js:"onbeforecopy"
func (*HTMLTableElement) SetOnbeforecopy(onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

// Onbeforecut prop
// js:"onbeforecut"
func (*HTMLTableElement) Onbeforecut() (onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut")
	return onbeforecut
}

// SetOnbeforecut prop
// js:"onbeforecut"
func (*HTMLTableElement) SetOnbeforecut(onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

// Onbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLTableElement) Onbeforedeactivate() (onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

// SetOnbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLTableElement) SetOnbeforedeactivate(onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

// Onbeforepaste prop
// js:"onbeforepaste"
func (*HTMLTableElement) Onbeforepaste() (onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

// SetOnbeforepaste prop
// js:"onbeforepaste"
func (*HTMLTableElement) SetOnbeforepaste(onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

// Onblur prop
// js:"onblur"
func (*HTMLTableElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*HTMLTableElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*HTMLTableElement) Oncanplay() (oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*HTMLTableElement) SetOncanplay(oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLTableElement) Oncanplaythrough() (oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLTableElement) SetOncanplaythrough(oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*HTMLTableElement) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*HTMLTableElement) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*HTMLTableElement) Onclick() (onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*HTMLTableElement) SetOnclick(onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick = $1", onclick)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*HTMLTableElement) Oncontextmenu() (oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*HTMLTableElement) SetOncontextmenu(oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Oncopy prop
// js:"oncopy"
func (*HTMLTableElement) Oncopy() (oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy")
	return oncopy
}

// SetOncopy prop
// js:"oncopy"
func (*HTMLTableElement) SetOncopy(oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy = $1", oncopy)
}

// Oncuechange prop
// js:"oncuechange"
func (*HTMLTableElement) Oncuechange() (oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange")
	return oncuechange
}

// SetOncuechange prop
// js:"oncuechange"
func (*HTMLTableElement) SetOncuechange(oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange = $1", oncuechange)
}

// Oncut prop
// js:"oncut"
func (*HTMLTableElement) Oncut() (oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut")
	return oncut
}

// SetOncut prop
// js:"oncut"
func (*HTMLTableElement) SetOncut(oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut = $1", oncut)
}

// Ondblclick prop
// js:"ondblclick"
func (*HTMLTableElement) Ondblclick() (ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*HTMLTableElement) SetOndblclick(ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondeactivate prop
// js:"ondeactivate"
func (*HTMLTableElement) Ondeactivate() (ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate")
	return ondeactivate
}

// SetOndeactivate prop
// js:"ondeactivate"
func (*HTMLTableElement) SetOndeactivate(ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

// Ondrag prop
// js:"ondrag"
func (*HTMLTableElement) Ondrag() (ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*HTMLTableElement) SetOndrag(ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*HTMLTableElement) Ondragend() (ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*HTMLTableElement) SetOndragend(ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*HTMLTableElement) Ondragenter() (ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*HTMLTableElement) SetOndragenter(ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*HTMLTableElement) Ondragleave() (ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*HTMLTableElement) SetOndragleave(ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*HTMLTableElement) Ondragover() (ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*HTMLTableElement) SetOndragover(ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*HTMLTableElement) Ondragstart() (ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*HTMLTableElement) SetOndragstart(ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*HTMLTableElement) Ondrop() (ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*HTMLTableElement) SetOndrop(ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*HTMLTableElement) Ondurationchange() (ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*HTMLTableElement) SetOndurationchange(ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*HTMLTableElement) Onemptied() (onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*HTMLTableElement) SetOnemptied(onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*HTMLTableElement) Onended() (onended func(window.Event)) {
	js.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*HTMLTableElement) SetOnended(onended func(window.Event)) {
	js.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*HTMLTableElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*HTMLTableElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*HTMLTableElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*HTMLTableElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus = $1", onfocus)
}

// Oninput prop
// js:"oninput"
func (*HTMLTableElement) Oninput() (oninput func(window.Event)) {
	js.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*HTMLTableElement) SetOninput(oninput func(window.Event)) {
	js.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*HTMLTableElement) Oninvalid() (oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*HTMLTableElement) SetOninvalid(oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*HTMLTableElement) Onkeydown() (onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*HTMLTableElement) SetOnkeydown(onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*HTMLTableElement) Onkeypress() (onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*HTMLTableElement) SetOnkeypress(onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*HTMLTableElement) Onkeyup() (onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*HTMLTableElement) SetOnkeyup(onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop
// js:"onload"
func (*HTMLTableElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*HTMLTableElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*HTMLTableElement) Onloadeddata() (onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*HTMLTableElement) SetOnloadeddata(onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLTableElement) Onloadedmetadata() (onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLTableElement) SetOnloadedmetadata(onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*HTMLTableElement) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*HTMLTableElement) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmousedown prop
// js:"onmousedown"
func (*HTMLTableElement) Onmousedown() (onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*HTMLTableElement) SetOnmousedown(onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*HTMLTableElement) Onmouseenter() (onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*HTMLTableElement) SetOnmouseenter(onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*HTMLTableElement) Onmouseleave() (onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*HTMLTableElement) SetOnmouseleave(onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*HTMLTableElement) Onmousemove() (onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*HTMLTableElement) SetOnmousemove(onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*HTMLTableElement) Onmouseout() (onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*HTMLTableElement) SetOnmouseout(onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*HTMLTableElement) Onmouseover() (onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*HTMLTableElement) SetOnmouseover(onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*HTMLTableElement) Onmouseup() (onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*HTMLTableElement) SetOnmouseup(onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*HTMLTableElement) Onmousewheel() (onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*HTMLTableElement) SetOnmousewheel(onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLTableElement) Onmscontentzoom() (onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

// SetOnmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLTableElement) SetOnmscontentzoom(onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

// Onmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLTableElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// SetOnmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLTableElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

// Onpaste prop
// js:"onpaste"
func (*HTMLTableElement) Onpaste() (onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste")
	return onpaste
}

// SetOnpaste prop
// js:"onpaste"
func (*HTMLTableElement) SetOnpaste(onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste = $1", onpaste)
}

// Onpause prop
// js:"onpause"
func (*HTMLTableElement) Onpause() (onpause func(window.Event)) {
	js.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*HTMLTableElement) SetOnpause(onpause func(window.Event)) {
	js.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*HTMLTableElement) Onplay() (onplay func(window.Event)) {
	js.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*HTMLTableElement) SetOnplay(onplay func(window.Event)) {
	js.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*HTMLTableElement) Onplaying() (onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*HTMLTableElement) SetOnplaying(onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying = $1", onplaying)
}

// Onprogress prop
// js:"onprogress"
func (*HTMLTableElement) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*HTMLTableElement) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*HTMLTableElement) Onratechange() (onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*HTMLTableElement) SetOnratechange(onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreset prop
// js:"onreset"
func (*HTMLTableElement) Onreset() (onreset func(window.Event)) {
	js.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*HTMLTableElement) SetOnreset(onreset func(window.Event)) {
	js.Rewrite("$_.onreset = $1", onreset)
}

// Onscroll prop
// js:"onscroll"
func (*HTMLTableElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*HTMLTableElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*HTMLTableElement) Onseeked() (onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*HTMLTableElement) SetOnseeked(onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*HTMLTableElement) Onseeking() (onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*HTMLTableElement) SetOnseeking(onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*HTMLTableElement) Onselect() (onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*HTMLTableElement) SetOnselect(onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect = $1", onselect)
}

// Onselectstart prop
// js:"onselectstart"
func (*HTMLTableElement) Onselectstart() (onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart")
	return onselectstart
}

// SetOnselectstart prop
// js:"onselectstart"
func (*HTMLTableElement) SetOnselectstart(onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart = $1", onselectstart)
}

// Onstalled prop
// js:"onstalled"
func (*HTMLTableElement) Onstalled() (onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*HTMLTableElement) SetOnstalled(onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled = $1", onstalled)
}

// Onsubmit prop
// js:"onsubmit"
func (*HTMLTableElement) Onsubmit() (onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*HTMLTableElement) SetOnsubmit(onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*HTMLTableElement) Onsuspend() (onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*HTMLTableElement) SetOnsuspend(onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*HTMLTableElement) Ontimeupdate() (ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*HTMLTableElement) SetOntimeupdate(ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*HTMLTableElement) Onvolumechange() (onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*HTMLTableElement) SetOnvolumechange(onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*HTMLTableElement) Onwaiting() (onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*HTMLTableElement) SetOnwaiting(onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting = $1", onwaiting)
}

// OuterText prop
// js:"outerText"
func (*HTMLTableElement) OuterText() (outerText string) {
	js.Rewrite("$_.outerText")
	return outerText
}

// SetOuterText prop
// js:"outerText"
func (*HTMLTableElement) SetOuterText(outerText string) {
	js.Rewrite("$_.outerText = $1", outerText)
}

// Spellcheck prop
// js:"spellcheck"
func (*HTMLTableElement) Spellcheck() (spellcheck bool) {
	js.Rewrite("$_.spellcheck")
	return spellcheck
}

// SetSpellcheck prop
// js:"spellcheck"
func (*HTMLTableElement) SetSpellcheck(spellcheck bool) {
	js.Rewrite("$_.spellcheck = $1", spellcheck)
}

// Style prop
// js:"style"
func (*HTMLTableElement) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// TabIndex prop
// js:"tabIndex"
func (*HTMLTableElement) TabIndex() (tabIndex int8) {
	js.Rewrite("$_.tabIndex")
	return tabIndex
}

// SetTabIndex prop
// js:"tabIndex"
func (*HTMLTableElement) SetTabIndex(tabIndex int8) {
	js.Rewrite("$_.tabIndex = $1", tabIndex)
}

// Title prop
// js:"title"
func (*HTMLTableElement) Title() (title string) {
	js.Rewrite("$_.title")
	return title
}

// SetTitle prop
// js:"title"
func (*HTMLTableElement) SetTitle(title string) {
	js.Rewrite("$_.title = $1", title)
}

// ClassList prop
// js:"classList"
func (*HTMLTableElement) ClassList() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*HTMLTableElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLTableElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*HTMLTableElement) ClientHeight() (clientHeight int) {
	js.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*HTMLTableElement) ClientLeft() (clientLeft int) {
	js.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*HTMLTableElement) ClientTop() (clientTop int) {
	js.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*HTMLTableElement) ClientWidth() (clientWidth int) {
	js.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*HTMLTableElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLTableElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLTableElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLTableElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLTableElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLTableElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*HTMLTableElement) MsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*HTMLTableElement) Onariarequest() (onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*HTMLTableElement) SetOnariarequest(onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*HTMLTableElement) Oncommand() (oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*HTMLTableElement) SetOncommand(oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLTableElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLTableElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLTableElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLTableElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLTableElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLTableElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLTableElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLTableElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*HTMLTableElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*HTMLTableElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLTableElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLTableElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLTableElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLTableElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLTableElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLTableElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLTableElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLTableElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLTableElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLTableElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLTableElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLTableElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*HTMLTableElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*HTMLTableElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*HTMLTableElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*HTMLTableElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*HTMLTableElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*HTMLTableElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*HTMLTableElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*HTMLTableElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*HTMLTableElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*HTMLTableElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*HTMLTableElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*HTMLTableElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*HTMLTableElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*HTMLTableElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*HTMLTableElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*HTMLTableElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*HTMLTableElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*HTMLTableElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*HTMLTableElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*HTMLTableElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*HTMLTableElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*HTMLTableElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*HTMLTableElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*HTMLTableElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLTableElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLTableElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLTableElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLTableElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLTableElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLTableElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*HTMLTableElement) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLTableElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLTableElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLTableElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLTableElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLTableElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*HTMLTableElement) ScrollWidth() (scrollWidth int) {
	js.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*HTMLTableElement) TagName() (tagName string) {
	js.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*HTMLTableElement) Onpointercancel() (onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*HTMLTableElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*HTMLTableElement) Onpointerdown() (onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*HTMLTableElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*HTMLTableElement) Onpointerenter() (onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*HTMLTableElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*HTMLTableElement) Onpointerleave() (onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*HTMLTableElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*HTMLTableElement) Onpointermove() (onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*HTMLTableElement) SetOnpointermove(onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*HTMLTableElement) Onpointerout() (onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*HTMLTableElement) SetOnpointerout(onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*HTMLTableElement) Onpointerover() (onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*HTMLTableElement) SetOnpointerover(onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*HTMLTableElement) Onpointerup() (onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*HTMLTableElement) SetOnpointerup(onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*HTMLTableElement) Onwheel() (onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*HTMLTableElement) SetOnwheel(onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*HTMLTableElement) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*HTMLTableElement) FirstElementChild() (firstElementChild window.Element) {
	js.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*HTMLTableElement) LastElementChild() (lastElementChild window.Element) {
	js.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*HTMLTableElement) NextElementSibling() (nextElementSibling window.Element) {
	js.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*HTMLTableElement) PreviousElementSibling() (previousElementSibling window.Element) {
	js.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*HTMLTableElement) Attributes() (attributes *window.NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*HTMLTableElement) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLTableElement) ChildNodes() (childNodes *window.NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLTableElement) FirstChild() (firstChild window.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLTableElement) LastChild() (lastChild window.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*HTMLTableElement) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*HTMLTableElement) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLTableElement) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLTableElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLTableElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLTableElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLTableElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*HTMLTableElement) OwnerDocument() (ownerDocument window.Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*HTMLTableElement) ParentElement() (parentElement window.HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*HTMLTableElement) ParentNode() (parentNode window.Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*HTMLTableElement) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*HTMLTableElement) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*HTMLTableElement) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
