package window

import (
	"github.com/matthewmueller/joy/dom/childnode"
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/domstringmap"
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/dom/mszoomtooptions"
	"github.com/matthewmueller/joy/macro"
)

var _ HTMLElement = (*HTMLHeadElement)(nil)
var _ Element = (*HTMLHeadElement)(nil)
var _ GlobalEventHandlers = (*HTMLHeadElement)(nil)
var _ ElementTraversal = (*HTMLHeadElement)(nil)
var _ NodeSelector = (*HTMLHeadElement)(nil)
var _ childnode.ChildNode = (*HTMLHeadElement)(nil)
var _ Node = (*HTMLHeadElement)(nil)
var _ EventTarget = (*HTMLHeadElement)(nil)

// HTMLHeadElement struct
// js:"HTMLHeadElement,omit"
type HTMLHeadElement struct {
}

// Blur fn
// js:"blur"
func (*HTMLHeadElement) Blur() {
	macro.Rewrite("$_.blur()")
}

// Click fn
// js:"click"
func (*HTMLHeadElement) Click() {
	macro.Rewrite("$_.click()")
}

// DragDrop fn
// js:"dragDrop"
func (*HTMLHeadElement) DragDrop() (b bool) {
	macro.Rewrite("$_.dragDrop()")
	return b
}

// Focus fn
// js:"focus"
func (*HTMLHeadElement) Focus() {
	macro.Rewrite("$_.focus()")
}

// GetElementsByClassName fn
// js:"getElementsByClassName"
func (*HTMLHeadElement) GetElementsByClassName(classNames string) (n *NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return n
}

// InsertAdjacentElement fn
// js:"insertAdjacentElement"
func (*HTMLHeadElement) InsertAdjacentElement(position string, insertedElement Element) (e Element) {
	macro.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return e
}

// InsertAdjacentHTML fn
// js:"insertAdjacentHTML"
func (*HTMLHeadElement) InsertAdjacentHTML(where string, html string) {
	macro.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

// InsertAdjacentText fn
// js:"insertAdjacentText"
func (*HTMLHeadElement) InsertAdjacentText(where string, text string) {
	macro.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

// MsGetInputContext fn
// js:"msGetInputContext"
func (*HTMLHeadElement) MsGetInputContext() (m *MSInputMethodContext) {
	macro.Rewrite("$_.msGetInputContext()")
	return m
}

// ScrollIntoView fn
// js:"scrollIntoView"
func (*HTMLHeadElement) ScrollIntoView(top *bool) {
	macro.Rewrite("$_.scrollIntoView($1)", top)
}

// GetAttribute fn
// js:"getAttribute"
func (*HTMLHeadElement) GetAttribute(qualifiedName string) (s string) {
	macro.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*HTMLHeadElement) GetAttributeNode(name string) (a *Attr) {
	macro.Rewrite("$_.getAttributeNode($1)", name)
	return a
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*HTMLHeadElement) GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr) {
	macro.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return a
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*HTMLHeadElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	macro.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*HTMLHeadElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*HTMLHeadElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*HTMLHeadElement) GetElementsByTagName(name string) (n *NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", name)
	return n
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*HTMLHeadElement) GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

// HasAttribute fn
// js:"hasAttribute"
func (*HTMLHeadElement) HasAttribute(name string) (b bool) {
	macro.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*HTMLHeadElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	macro.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*HTMLHeadElement) MsGetRegionContent() (m *MSRangeCollection) {
	macro.Rewrite("$_.msGetRegionContent()")
	return m
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*HTMLHeadElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*HTMLHeadElement) MsMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*HTMLHeadElement) MsReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*HTMLHeadElement) MsSetPointerCapture(pointerId int) {
	macro.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*HTMLHeadElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	macro.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*HTMLHeadElement) ReleasePointerCapture(pointerId int) {
	macro.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*HTMLHeadElement) RemoveAttribute(qualifiedName string) {
	macro.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*HTMLHeadElement) RemoveAttributeNode(oldAttr *Attr) (a *Attr) {
	macro.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return a
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*HTMLHeadElement) RemoveAttributeNS(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*HTMLHeadElement) RequestFullscreen() {
	macro.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*HTMLHeadElement) RequestPointerLock() {
	macro.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*HTMLHeadElement) SetAttribute(qualifiedName string, value string) {
	macro.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*HTMLHeadElement) SetAttributeNode(newAttr *Attr) (a *Attr) {
	macro.Rewrite("$_.setAttributeNode($1)", newAttr)
	return a
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*HTMLHeadElement) SetAttributeNodeNS(newAttr *Attr) (a *Attr) {
	macro.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return a
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*HTMLHeadElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	macro.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*HTMLHeadElement) SetPointerCapture(pointerId int) {
	macro.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*HTMLHeadElement) WebkitMatchesSelector(selectors string) (b bool) {
	macro.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*HTMLHeadElement) WebkitRequestFullscreen() {
	macro.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*HTMLHeadElement) WebkitRequestFullScreen() {
	macro.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLHeadElement) QuerySelector(selectors string) (e Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return e
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*HTMLHeadElement) QuerySelectorAll(selectors string) (n *NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return n
}

// AppendChild fn
// js:"appendChild"
func (*HTMLHeadElement) AppendChild(newChild Node) (n Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return n
}

// CloneNode fn
// js:"cloneNode"
func (*HTMLHeadElement) CloneNode(deep *bool) (n Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return n
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*HTMLHeadElement) CompareDocumentPosition(other Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*HTMLHeadElement) Contains(child Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*HTMLHeadElement) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*HTMLHeadElement) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*HTMLHeadElement) InsertBefore(newChild Node, refChild *Node) (n Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return n
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*HTMLHeadElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*HTMLHeadElement) IsEqualNode(arg Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*HTMLHeadElement) IsSameNode(other Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*HTMLHeadElement) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*HTMLHeadElement) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*HTMLHeadElement) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*HTMLHeadElement) RemoveChild(oldChild Node) (n Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return n
}

// ReplaceChild fn
// js:"replaceChild"
func (*HTMLHeadElement) ReplaceChild(newChild Node, oldChild Node) (n Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLHeadElement) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLHeadElement) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*HTMLHeadElement) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Profile prop
// js:"profile"
func (*HTMLHeadElement) Profile() (profile string) {
	macro.Rewrite("$_.profile")
	return profile
}

// SetProfile prop
// js:"profile"
func (*HTMLHeadElement) SetProfile(profile string) {
	macro.Rewrite("$_.profile = $1", profile)
}

// AccessKey prop
// js:"accessKey"
func (*HTMLHeadElement) AccessKey() (accessKey string) {
	macro.Rewrite("$_.accessKey")
	return accessKey
}

// SetAccessKey prop
// js:"accessKey"
func (*HTMLHeadElement) SetAccessKey(accessKey string) {
	macro.Rewrite("$_.accessKey = $1", accessKey)
}

// Children prop
// js:"children"
func (*HTMLHeadElement) Children() (children HTMLCollection) {
	macro.Rewrite("$_.children")
	return children
}

// ContentEditable prop
// js:"contentEditable"
func (*HTMLHeadElement) ContentEditable() (contentEditable string) {
	macro.Rewrite("$_.contentEditable")
	return contentEditable
}

// SetContentEditable prop
// js:"contentEditable"
func (*HTMLHeadElement) SetContentEditable(contentEditable string) {
	macro.Rewrite("$_.contentEditable = $1", contentEditable)
}

// Dataset prop
// js:"dataset"
func (*HTMLHeadElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	macro.Rewrite("$_.dataset")
	return dataset
}

// Dir prop
// js:"dir"
func (*HTMLHeadElement) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

// SetDir prop
// js:"dir"
func (*HTMLHeadElement) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

// Draggable prop
// js:"draggable"
func (*HTMLHeadElement) Draggable() (draggable bool) {
	macro.Rewrite("$_.draggable")
	return draggable
}

// SetDraggable prop
// js:"draggable"
func (*HTMLHeadElement) SetDraggable(draggable bool) {
	macro.Rewrite("$_.draggable = $1", draggable)
}

// Hidden prop
// js:"hidden"
func (*HTMLHeadElement) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

// SetHidden prop
// js:"hidden"
func (*HTMLHeadElement) SetHidden(hidden bool) {
	macro.Rewrite("$_.hidden = $1", hidden)
}

// HideFocus prop
// js:"hideFocus"
func (*HTMLHeadElement) HideFocus() (hideFocus bool) {
	macro.Rewrite("$_.hideFocus")
	return hideFocus
}

// SetHideFocus prop
// js:"hideFocus"
func (*HTMLHeadElement) SetHideFocus(hideFocus bool) {
	macro.Rewrite("$_.hideFocus = $1", hideFocus)
}

// InnerText prop
// js:"innerText"
func (*HTMLHeadElement) InnerText() (innerText string) {
	macro.Rewrite("$_.innerText")
	return innerText
}

// SetInnerText prop
// js:"innerText"
func (*HTMLHeadElement) SetInnerText(innerText string) {
	macro.Rewrite("$_.innerText = $1", innerText)
}

// IsContentEditable prop
// js:"isContentEditable"
func (*HTMLHeadElement) IsContentEditable() (isContentEditable bool) {
	macro.Rewrite("$_.isContentEditable")
	return isContentEditable
}

// Lang prop
// js:"lang"
func (*HTMLHeadElement) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

// SetLang prop
// js:"lang"
func (*HTMLHeadElement) SetLang(lang string) {
	macro.Rewrite("$_.lang = $1", lang)
}

// OffsetHeight prop
// js:"offsetHeight"
func (*HTMLHeadElement) OffsetHeight() (offsetHeight int) {
	macro.Rewrite("$_.offsetHeight")
	return offsetHeight
}

// OffsetLeft prop
// js:"offsetLeft"
func (*HTMLHeadElement) OffsetLeft() (offsetLeft int) {
	macro.Rewrite("$_.offsetLeft")
	return offsetLeft
}

// OffsetParent prop
// js:"offsetParent"
func (*HTMLHeadElement) OffsetParent() (offsetParent Element) {
	macro.Rewrite("$_.offsetParent")
	return offsetParent
}

// OffsetTop prop
// js:"offsetTop"
func (*HTMLHeadElement) OffsetTop() (offsetTop int) {
	macro.Rewrite("$_.offsetTop")
	return offsetTop
}

// OffsetWidth prop
// js:"offsetWidth"
func (*HTMLHeadElement) OffsetWidth() (offsetWidth int) {
	macro.Rewrite("$_.offsetWidth")
	return offsetWidth
}

// Onabort prop
// js:"onabort"
func (*HTMLHeadElement) Onabort() (onabort func(Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*HTMLHeadElement) SetOnabort(onabort func(Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onactivate prop
// js:"onactivate"
func (*HTMLHeadElement) Onactivate() (onactivate func(UIEvent)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

// SetOnactivate prop
// js:"onactivate"
func (*HTMLHeadElement) SetOnactivate(onactivate func(UIEvent)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

// Onbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLHeadElement) Onbeforeactivate() (onbeforeactivate func(UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

// SetOnbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLHeadElement) SetOnbeforeactivate(onbeforeactivate func(UIEvent)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

// Onbeforecopy prop
// js:"onbeforecopy"
func (*HTMLHeadElement) Onbeforecopy() (onbeforecopy func(*ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

// SetOnbeforecopy prop
// js:"onbeforecopy"
func (*HTMLHeadElement) SetOnbeforecopy(onbeforecopy func(*ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

// Onbeforecut prop
// js:"onbeforecut"
func (*HTMLHeadElement) Onbeforecut() (onbeforecut func(*ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut")
	return onbeforecut
}

// SetOnbeforecut prop
// js:"onbeforecut"
func (*HTMLHeadElement) SetOnbeforecut(onbeforecut func(*ClipboardEvent)) {
	macro.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

// Onbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLHeadElement) Onbeforedeactivate() (onbeforedeactivate func(UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

// SetOnbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLHeadElement) SetOnbeforedeactivate(onbeforedeactivate func(UIEvent)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

// Onbeforepaste prop
// js:"onbeforepaste"
func (*HTMLHeadElement) Onbeforepaste() (onbeforepaste func(*ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

// SetOnbeforepaste prop
// js:"onbeforepaste"
func (*HTMLHeadElement) SetOnbeforepaste(onbeforepaste func(*ClipboardEvent)) {
	macro.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

// Onblur prop
// js:"onblur"
func (*HTMLHeadElement) Onblur() (onblur func(*FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*HTMLHeadElement) SetOnblur(onblur func(*FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*HTMLHeadElement) Oncanplay() (oncanplay func(Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*HTMLHeadElement) SetOncanplay(oncanplay func(Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLHeadElement) Oncanplaythrough() (oncanplaythrough func(Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLHeadElement) SetOncanplaythrough(oncanplaythrough func(Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*HTMLHeadElement) Onchange() (onchange func(Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*HTMLHeadElement) SetOnchange(onchange func(Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*HTMLHeadElement) Onclick() (onclick func(MouseEvent)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*HTMLHeadElement) SetOnclick(onclick func(MouseEvent)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*HTMLHeadElement) Oncontextmenu() (oncontextmenu func(*PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*HTMLHeadElement) SetOncontextmenu(oncontextmenu func(*PointerEvent)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Oncopy prop
// js:"oncopy"
func (*HTMLHeadElement) Oncopy() (oncopy func(*ClipboardEvent)) {
	macro.Rewrite("$_.oncopy")
	return oncopy
}

// SetOncopy prop
// js:"oncopy"
func (*HTMLHeadElement) SetOncopy(oncopy func(*ClipboardEvent)) {
	macro.Rewrite("$_.oncopy = $1", oncopy)
}

// Oncuechange prop
// js:"oncuechange"
func (*HTMLHeadElement) Oncuechange() (oncuechange func(Event)) {
	macro.Rewrite("$_.oncuechange")
	return oncuechange
}

// SetOncuechange prop
// js:"oncuechange"
func (*HTMLHeadElement) SetOncuechange(oncuechange func(Event)) {
	macro.Rewrite("$_.oncuechange = $1", oncuechange)
}

// Oncut prop
// js:"oncut"
func (*HTMLHeadElement) Oncut() (oncut func(*ClipboardEvent)) {
	macro.Rewrite("$_.oncut")
	return oncut
}

// SetOncut prop
// js:"oncut"
func (*HTMLHeadElement) SetOncut(oncut func(*ClipboardEvent)) {
	macro.Rewrite("$_.oncut = $1", oncut)
}

// Ondblclick prop
// js:"ondblclick"
func (*HTMLHeadElement) Ondblclick() (ondblclick func(MouseEvent)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*HTMLHeadElement) SetOndblclick(ondblclick func(MouseEvent)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondeactivate prop
// js:"ondeactivate"
func (*HTMLHeadElement) Ondeactivate() (ondeactivate func(UIEvent)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

// SetOndeactivate prop
// js:"ondeactivate"
func (*HTMLHeadElement) SetOndeactivate(ondeactivate func(UIEvent)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

// Ondrag prop
// js:"ondrag"
func (*HTMLHeadElement) Ondrag() (ondrag func(*DragEvent)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*HTMLHeadElement) SetOndrag(ondrag func(*DragEvent)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*HTMLHeadElement) Ondragend() (ondragend func(*DragEvent)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*HTMLHeadElement) SetOndragend(ondragend func(*DragEvent)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*HTMLHeadElement) Ondragenter() (ondragenter func(*DragEvent)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*HTMLHeadElement) SetOndragenter(ondragenter func(*DragEvent)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*HTMLHeadElement) Ondragleave() (ondragleave func(*DragEvent)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*HTMLHeadElement) SetOndragleave(ondragleave func(*DragEvent)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*HTMLHeadElement) Ondragover() (ondragover func(*DragEvent)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*HTMLHeadElement) SetOndragover(ondragover func(*DragEvent)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*HTMLHeadElement) Ondragstart() (ondragstart func(*DragEvent)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*HTMLHeadElement) SetOndragstart(ondragstart func(*DragEvent)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*HTMLHeadElement) Ondrop() (ondrop func(*DragEvent)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*HTMLHeadElement) SetOndrop(ondrop func(*DragEvent)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*HTMLHeadElement) Ondurationchange() (ondurationchange func(Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*HTMLHeadElement) SetOndurationchange(ondurationchange func(Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*HTMLHeadElement) Onemptied() (onemptied func(Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*HTMLHeadElement) SetOnemptied(onemptied func(Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*HTMLHeadElement) Onended() (onended func(Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*HTMLHeadElement) SetOnended(onended func(Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*HTMLHeadElement) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*HTMLHeadElement) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*HTMLHeadElement) Onfocus() (onfocus func(*FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*HTMLHeadElement) SetOnfocus(onfocus func(*FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

// Oninput prop
// js:"oninput"
func (*HTMLHeadElement) Oninput() (oninput func(Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*HTMLHeadElement) SetOninput(oninput func(Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*HTMLHeadElement) Oninvalid() (oninvalid func(Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*HTMLHeadElement) SetOninvalid(oninvalid func(Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*HTMLHeadElement) Onkeydown() (onkeydown func(*KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*HTMLHeadElement) SetOnkeydown(onkeydown func(*KeyboardEvent)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*HTMLHeadElement) Onkeypress() (onkeypress func(*KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*HTMLHeadElement) SetOnkeypress(onkeypress func(*KeyboardEvent)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*HTMLHeadElement) Onkeyup() (onkeyup func(*KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*HTMLHeadElement) SetOnkeyup(onkeyup func(*KeyboardEvent)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop
// js:"onload"
func (*HTMLHeadElement) Onload() (onload func(Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*HTMLHeadElement) SetOnload(onload func(Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*HTMLHeadElement) Onloadeddata() (onloadeddata func(Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*HTMLHeadElement) SetOnloadeddata(onloadeddata func(Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLHeadElement) Onloadedmetadata() (onloadedmetadata func(Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLHeadElement) SetOnloadedmetadata(onloadedmetadata func(Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*HTMLHeadElement) Onloadstart() (onloadstart func(Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*HTMLHeadElement) SetOnloadstart(onloadstart func(Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmousedown prop
// js:"onmousedown"
func (*HTMLHeadElement) Onmousedown() (onmousedown func(MouseEvent)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*HTMLHeadElement) SetOnmousedown(onmousedown func(MouseEvent)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*HTMLHeadElement) Onmouseenter() (onmouseenter func(MouseEvent)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*HTMLHeadElement) SetOnmouseenter(onmouseenter func(MouseEvent)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*HTMLHeadElement) Onmouseleave() (onmouseleave func(MouseEvent)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*HTMLHeadElement) SetOnmouseleave(onmouseleave func(MouseEvent)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*HTMLHeadElement) Onmousemove() (onmousemove func(MouseEvent)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*HTMLHeadElement) SetOnmousemove(onmousemove func(MouseEvent)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*HTMLHeadElement) Onmouseout() (onmouseout func(MouseEvent)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*HTMLHeadElement) SetOnmouseout(onmouseout func(MouseEvent)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*HTMLHeadElement) Onmouseover() (onmouseover func(MouseEvent)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*HTMLHeadElement) SetOnmouseover(onmouseover func(MouseEvent)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*HTMLHeadElement) Onmouseup() (onmouseup func(MouseEvent)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*HTMLHeadElement) SetOnmouseup(onmouseup func(MouseEvent)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*HTMLHeadElement) Onmousewheel() (onmousewheel func(*WheelEvent)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*HTMLHeadElement) SetOnmousewheel(onmousewheel func(*WheelEvent)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLHeadElement) Onmscontentzoom() (onmscontentzoom func(UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

// SetOnmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLHeadElement) SetOnmscontentzoom(onmscontentzoom func(UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

// Onmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLHeadElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// SetOnmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLHeadElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

// Onpaste prop
// js:"onpaste"
func (*HTMLHeadElement) Onpaste() (onpaste func(*ClipboardEvent)) {
	macro.Rewrite("$_.onpaste")
	return onpaste
}

// SetOnpaste prop
// js:"onpaste"
func (*HTMLHeadElement) SetOnpaste(onpaste func(*ClipboardEvent)) {
	macro.Rewrite("$_.onpaste = $1", onpaste)
}

// Onpause prop
// js:"onpause"
func (*HTMLHeadElement) Onpause() (onpause func(Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*HTMLHeadElement) SetOnpause(onpause func(Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*HTMLHeadElement) Onplay() (onplay func(Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*HTMLHeadElement) SetOnplay(onplay func(Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*HTMLHeadElement) Onplaying() (onplaying func(Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*HTMLHeadElement) SetOnplaying(onplaying func(Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

// Onprogress prop
// js:"onprogress"
func (*HTMLHeadElement) Onprogress() (onprogress func(Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*HTMLHeadElement) SetOnprogress(onprogress func(Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*HTMLHeadElement) Onratechange() (onratechange func(Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*HTMLHeadElement) SetOnratechange(onratechange func(Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreset prop
// js:"onreset"
func (*HTMLHeadElement) Onreset() (onreset func(Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*HTMLHeadElement) SetOnreset(onreset func(Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

// Onscroll prop
// js:"onscroll"
func (*HTMLHeadElement) Onscroll() (onscroll func(UIEvent)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*HTMLHeadElement) SetOnscroll(onscroll func(UIEvent)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*HTMLHeadElement) Onseeked() (onseeked func(Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*HTMLHeadElement) SetOnseeked(onseeked func(Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*HTMLHeadElement) Onseeking() (onseeking func(Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*HTMLHeadElement) SetOnseeking(onseeking func(Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*HTMLHeadElement) Onselect() (onselect func(UIEvent)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*HTMLHeadElement) SetOnselect(onselect func(UIEvent)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

// Onselectstart prop
// js:"onselectstart"
func (*HTMLHeadElement) Onselectstart() (onselectstart func(Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

// SetOnselectstart prop
// js:"onselectstart"
func (*HTMLHeadElement) SetOnselectstart(onselectstart func(Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

// Onstalled prop
// js:"onstalled"
func (*HTMLHeadElement) Onstalled() (onstalled func(Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*HTMLHeadElement) SetOnstalled(onstalled func(Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

// Onsubmit prop
// js:"onsubmit"
func (*HTMLHeadElement) Onsubmit() (onsubmit func(Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*HTMLHeadElement) SetOnsubmit(onsubmit func(Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*HTMLHeadElement) Onsuspend() (onsuspend func(Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*HTMLHeadElement) SetOnsuspend(onsuspend func(Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*HTMLHeadElement) Ontimeupdate() (ontimeupdate func(Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*HTMLHeadElement) SetOntimeupdate(ontimeupdate func(Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*HTMLHeadElement) Onvolumechange() (onvolumechange func(Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*HTMLHeadElement) SetOnvolumechange(onvolumechange func(Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*HTMLHeadElement) Onwaiting() (onwaiting func(Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*HTMLHeadElement) SetOnwaiting(onwaiting func(Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

// OuterText prop
// js:"outerText"
func (*HTMLHeadElement) OuterText() (outerText string) {
	macro.Rewrite("$_.outerText")
	return outerText
}

// SetOuterText prop
// js:"outerText"
func (*HTMLHeadElement) SetOuterText(outerText string) {
	macro.Rewrite("$_.outerText = $1", outerText)
}

// Spellcheck prop
// js:"spellcheck"
func (*HTMLHeadElement) Spellcheck() (spellcheck bool) {
	macro.Rewrite("$_.spellcheck")
	return spellcheck
}

// SetSpellcheck prop
// js:"spellcheck"
func (*HTMLHeadElement) SetSpellcheck(spellcheck bool) {
	macro.Rewrite("$_.spellcheck = $1", spellcheck)
}

// Style prop
// js:"style"
func (*HTMLHeadElement) Style() (style *CSSStyleDeclaration) {
	macro.Rewrite("$_.style")
	return style
}

// TabIndex prop
// js:"tabIndex"
func (*HTMLHeadElement) TabIndex() (tabIndex int8) {
	macro.Rewrite("$_.tabIndex")
	return tabIndex
}

// SetTabIndex prop
// js:"tabIndex"
func (*HTMLHeadElement) SetTabIndex(tabIndex int8) {
	macro.Rewrite("$_.tabIndex = $1", tabIndex)
}

// Title prop
// js:"title"
func (*HTMLHeadElement) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

// SetTitle prop
// js:"title"
func (*HTMLHeadElement) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

// ClassList prop
// js:"classList"
func (*HTMLHeadElement) ClassList() (classList domtokenlist.DOMTokenList) {
	macro.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*HTMLHeadElement) ClassName() (className string) {
	macro.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLHeadElement) SetClassName(className string) {
	macro.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*HTMLHeadElement) ClientHeight() (clientHeight int) {
	macro.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*HTMLHeadElement) ClientLeft() (clientLeft int) {
	macro.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*HTMLHeadElement) ClientTop() (clientTop int) {
	macro.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*HTMLHeadElement) ClientWidth() (clientWidth int) {
	macro.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*HTMLHeadElement) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLHeadElement) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLHeadElement) InnerHTML() (innerHTML string) {
	macro.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLHeadElement) SetInnerHTML(innerHTML string) {
	macro.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLHeadElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLHeadElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	macro.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*HTMLHeadElement) MsRegionOverflow() (msRegionOverflow string) {
	macro.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*HTMLHeadElement) Onariarequest() (onariarequest func(Event)) {
	macro.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*HTMLHeadElement) SetOnariarequest(onariarequest func(Event)) {
	macro.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*HTMLHeadElement) Oncommand() (oncommand func(Event)) {
	macro.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*HTMLHeadElement) SetOncommand(oncommand func(Event)) {
	macro.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLHeadElement) Ongotpointercapture() (ongotpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLHeadElement) SetOngotpointercapture(ongotpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLHeadElement) Onlostpointercapture() (onlostpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLHeadElement) SetOnlostpointercapture(onlostpointercapture func(*PointerEvent)) {
	macro.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLHeadElement) Onmsgesturechange() (onmsgesturechange func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLHeadElement) SetOnmsgesturechange(onmsgesturechange func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLHeadElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLHeadElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*HTMLHeadElement) Onmsgestureend() (onmsgestureend func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*HTMLHeadElement) SetOnmsgestureend(onmsgestureend func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLHeadElement) Onmsgesturehold() (onmsgesturehold func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLHeadElement) SetOnmsgesturehold(onmsgesturehold func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLHeadElement) Onmsgesturestart() (onmsgesturestart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLHeadElement) SetOnmsgesturestart(onmsgesturestart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLHeadElement) Onmsgesturetap() (onmsgesturetap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLHeadElement) SetOnmsgesturetap(onmsgesturetap func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLHeadElement) Onmsgotpointercapture() (onmsgotpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLHeadElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLHeadElement) Onmsinertiastart() (onmsinertiastart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLHeadElement) SetOnmsinertiastart(onmsinertiastart func(*MSGestureEvent)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLHeadElement) Onmslostpointercapture() (onmslostpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLHeadElement) SetOnmslostpointercapture(onmslostpointercapture func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*HTMLHeadElement) Onmspointercancel() (onmspointercancel func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*HTMLHeadElement) SetOnmspointercancel(onmspointercancel func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*HTMLHeadElement) Onmspointerdown() (onmspointerdown func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*HTMLHeadElement) SetOnmspointerdown(onmspointerdown func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*HTMLHeadElement) Onmspointerenter() (onmspointerenter func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*HTMLHeadElement) SetOnmspointerenter(onmspointerenter func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*HTMLHeadElement) Onmspointerleave() (onmspointerleave func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*HTMLHeadElement) SetOnmspointerleave(onmspointerleave func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*HTMLHeadElement) Onmspointermove() (onmspointermove func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*HTMLHeadElement) SetOnmspointermove(onmspointermove func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*HTMLHeadElement) Onmspointerout() (onmspointerout func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*HTMLHeadElement) SetOnmspointerout(onmspointerout func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*HTMLHeadElement) Onmspointerover() (onmspointerover func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*HTMLHeadElement) SetOnmspointerover(onmspointerover func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*HTMLHeadElement) Onmspointerup() (onmspointerup func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*HTMLHeadElement) SetOnmspointerup(onmspointerup func(*MSPointerEvent)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*HTMLHeadElement) Ontouchcancel() (ontouchcancel func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*HTMLHeadElement) SetOntouchcancel(ontouchcancel func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*HTMLHeadElement) Ontouchend() (ontouchend func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*HTMLHeadElement) SetOntouchend(ontouchend func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*HTMLHeadElement) Ontouchmove() (ontouchmove func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*HTMLHeadElement) SetOntouchmove(ontouchmove func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*HTMLHeadElement) Ontouchstart() (ontouchstart func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*HTMLHeadElement) SetOntouchstart(ontouchstart func(*TouchEvent)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLHeadElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLHeadElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLHeadElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLHeadElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLHeadElement) OuterHTML() (outerHTML string) {
	macro.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLHeadElement) SetOuterHTML(outerHTML string) {
	macro.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*HTMLHeadElement) Prefix() (prefix string) {
	macro.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLHeadElement) ScrollHeight() (scrollHeight int) {
	macro.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLHeadElement) ScrollLeft() (scrollLeft int) {
	macro.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLHeadElement) SetScrollLeft(scrollLeft int) {
	macro.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLHeadElement) ScrollTop() (scrollTop int) {
	macro.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLHeadElement) SetScrollTop(scrollTop int) {
	macro.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*HTMLHeadElement) ScrollWidth() (scrollWidth int) {
	macro.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*HTMLHeadElement) TagName() (tagName string) {
	macro.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*HTMLHeadElement) Onpointercancel() (onpointercancel func(Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*HTMLHeadElement) SetOnpointercancel(onpointercancel func(Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*HTMLHeadElement) Onpointerdown() (onpointerdown func(Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*HTMLHeadElement) SetOnpointerdown(onpointerdown func(Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*HTMLHeadElement) Onpointerenter() (onpointerenter func(Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*HTMLHeadElement) SetOnpointerenter(onpointerenter func(Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*HTMLHeadElement) Onpointerleave() (onpointerleave func(Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*HTMLHeadElement) SetOnpointerleave(onpointerleave func(Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*HTMLHeadElement) Onpointermove() (onpointermove func(Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*HTMLHeadElement) SetOnpointermove(onpointermove func(Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*HTMLHeadElement) Onpointerout() (onpointerout func(Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*HTMLHeadElement) SetOnpointerout(onpointerout func(Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*HTMLHeadElement) Onpointerover() (onpointerover func(Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*HTMLHeadElement) SetOnpointerover(onpointerover func(Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*HTMLHeadElement) Onpointerup() (onpointerup func(Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*HTMLHeadElement) SetOnpointerup(onpointerup func(Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*HTMLHeadElement) Onwheel() (onwheel func(Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*HTMLHeadElement) SetOnwheel(onwheel func(Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*HTMLHeadElement) ChildElementCount() (childElementCount uint) {
	macro.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*HTMLHeadElement) FirstElementChild() (firstElementChild Element) {
	macro.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*HTMLHeadElement) LastElementChild() (lastElementChild Element) {
	macro.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*HTMLHeadElement) NextElementSibling() (nextElementSibling Element) {
	macro.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*HTMLHeadElement) PreviousElementSibling() (previousElementSibling Element) {
	macro.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*HTMLHeadElement) Attributes() (attributes *NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*HTMLHeadElement) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLHeadElement) ChildNodes() (childNodes *NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLHeadElement) FirstChild() (firstChild Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLHeadElement) LastChild() (lastChild Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*HTMLHeadElement) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*HTMLHeadElement) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLHeadElement) NextSibling() (nextSibling Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLHeadElement) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLHeadElement) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLHeadElement) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLHeadElement) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*HTMLHeadElement) OwnerDocument() (ownerDocument Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*HTMLHeadElement) ParentElement() (parentElement HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*HTMLHeadElement) ParentNode() (parentNode Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*HTMLHeadElement) PreviousSibling() (previousSibling Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*HTMLHeadElement) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*HTMLHeadElement) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
