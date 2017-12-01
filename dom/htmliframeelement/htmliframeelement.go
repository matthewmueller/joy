package htmliframeelement

import (
	"github.com/matthewmueller/golly/dom/childnode"
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/domsettabletokenlist"
	"github.com/matthewmueller/golly/dom/domstringmap"
	"github.com/matthewmueller/golly/dom/domtokenlist"
	"github.com/matthewmueller/golly/dom/mszoomtooptions"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.HTMLElement = (*HTMLIFrameElement)(nil)
var _ window.Element = (*HTMLIFrameElement)(nil)
var _ window.GlobalEventHandlers = (*HTMLIFrameElement)(nil)
var _ window.ElementTraversal = (*HTMLIFrameElement)(nil)
var _ window.NodeSelector = (*HTMLIFrameElement)(nil)
var _ childnode.ChildNode = (*HTMLIFrameElement)(nil)
var _ window.Node = (*HTMLIFrameElement)(nil)
var _ window.EventTarget = (*HTMLIFrameElement)(nil)

// HTMLIFrameElement struct
// js:"HTMLIFrameElement,omit"
type HTMLIFrameElement struct {
}

// GetSVGDocument fn
// js:"getSVGDocument"
func (*HTMLIFrameElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$_.getSVGDocument()")
	return w
}

// Blur fn
// js:"blur"
func (*HTMLIFrameElement) Blur() {
	js.Rewrite("$_.blur()")
}

// Click fn
// js:"click"
func (*HTMLIFrameElement) Click() {
	js.Rewrite("$_.click()")
}

// DragDrop fn
// js:"dragDrop"
func (*HTMLIFrameElement) DragDrop() (b bool) {
	js.Rewrite("$_.dragDrop()")
	return b
}

// Focus fn
// js:"focus"
func (*HTMLIFrameElement) Focus() {
	js.Rewrite("$_.focus()")
}

// GetElementsByClassName fn
// js:"getElementsByClassName"
func (*HTMLIFrameElement) GetElementsByClassName(classNames string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

// InsertAdjacentElement fn
// js:"insertAdjacentElement"
func (*HTMLIFrameElement) InsertAdjacentElement(position string, insertedElement window.Element) (w window.Element) {
	js.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

// InsertAdjacentHTML fn
// js:"insertAdjacentHTML"
func (*HTMLIFrameElement) InsertAdjacentHTML(where string, html string) {
	js.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

// InsertAdjacentText fn
// js:"insertAdjacentText"
func (*HTMLIFrameElement) InsertAdjacentText(where string, text string) {
	js.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

// MsGetInputContext fn
// js:"msGetInputContext"
func (*HTMLIFrameElement) MsGetInputContext() (w *window.MSInputMethodContext) {
	js.Rewrite("$_.msGetInputContext()")
	return w
}

// ScrollIntoView fn
// js:"scrollIntoView"
func (*HTMLIFrameElement) ScrollIntoView(top *bool) {
	js.Rewrite("$_.scrollIntoView($1)", top)
}

// GetAttribute fn
// js:"getAttribute"
func (*HTMLIFrameElement) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*HTMLIFrameElement) GetAttributeNode(name string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*HTMLIFrameElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*HTMLIFrameElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*HTMLIFrameElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*HTMLIFrameElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*HTMLIFrameElement) GetElementsByTagName(name string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*HTMLIFrameElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*HTMLIFrameElement) HasAttribute(name string) (b bool) {
	js.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*HTMLIFrameElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*HTMLIFrameElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	js.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*HTMLIFrameElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*HTMLIFrameElement) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*HTMLIFrameElement) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*HTMLIFrameElement) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*HTMLIFrameElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*HTMLIFrameElement) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*HTMLIFrameElement) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*HTMLIFrameElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*HTMLIFrameElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*HTMLIFrameElement) RequestFullscreen() {
	js.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*HTMLIFrameElement) RequestPointerLock() {
	js.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*HTMLIFrameElement) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*HTMLIFrameElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*HTMLIFrameElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*HTMLIFrameElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*HTMLIFrameElement) SetPointerCapture(pointerId int) {
	js.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*HTMLIFrameElement) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*HTMLIFrameElement) WebkitRequestFullscreen() {
	js.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*HTMLIFrameElement) WebkitRequestFullScreen() {
	js.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLIFrameElement) QuerySelector(selectors string) (w window.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*HTMLIFrameElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*HTMLIFrameElement) AppendChild(newChild window.Node) (w window.Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*HTMLIFrameElement) CloneNode(deep *bool) (w window.Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*HTMLIFrameElement) CompareDocumentPosition(other window.Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*HTMLIFrameElement) Contains(child window.Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*HTMLIFrameElement) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*HTMLIFrameElement) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*HTMLIFrameElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*HTMLIFrameElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*HTMLIFrameElement) IsEqualNode(arg window.Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*HTMLIFrameElement) IsSameNode(other window.Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*HTMLIFrameElement) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*HTMLIFrameElement) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*HTMLIFrameElement) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*HTMLIFrameElement) RemoveChild(oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*HTMLIFrameElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLIFrameElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLIFrameElement) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*HTMLIFrameElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
// js:"align"
func (*HTMLIFrameElement) Align() (align string) {
	js.Rewrite("$_.align")
	return align
}

// SetAlign prop Sets or retrieves how the object is aligned with adjacent text.
// js:"align"
func (*HTMLIFrameElement) SetAlign(align string) {
	js.Rewrite("$_.align = $1", align)
}

// AllowFullscreen prop
// js:"allowFullscreen"
func (*HTMLIFrameElement) AllowFullscreen() (allowFullscreen bool) {
	js.Rewrite("$_.allowFullscreen")
	return allowFullscreen
}

// SetAllowFullscreen prop
// js:"allowFullscreen"
func (*HTMLIFrameElement) SetAllowFullscreen(allowFullscreen bool) {
	js.Rewrite("$_.allowFullscreen = $1", allowFullscreen)
}

// AllowPaymentRequest prop
// js:"allowPaymentRequest"
func (*HTMLIFrameElement) AllowPaymentRequest() (allowPaymentRequest bool) {
	js.Rewrite("$_.allowPaymentRequest")
	return allowPaymentRequest
}

// SetAllowPaymentRequest prop
// js:"allowPaymentRequest"
func (*HTMLIFrameElement) SetAllowPaymentRequest(allowPaymentRequest bool) {
	js.Rewrite("$_.allowPaymentRequest = $1", allowPaymentRequest)
}

// Border prop Specifies the properties of a border drawn around an object.
// js:"border"
func (*HTMLIFrameElement) Border() (border string) {
	js.Rewrite("$_.border")
	return border
}

// SetBorder prop Specifies the properties of a border drawn around an object.
// js:"border"
func (*HTMLIFrameElement) SetBorder(border string) {
	js.Rewrite("$_.border = $1", border)
}

// ContentDocument prop Retrieves the document object of the page or frame.
// js:"contentDocument"
func (*HTMLIFrameElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$_.contentDocument")
	return contentDocument
}

// ContentWindow prop Retrieves the object of the specified.
// js:"contentWindow"
func (*HTMLIFrameElement) ContentWindow() (contentWindow *window.Window) {
	js.Rewrite("$_.contentWindow")
	return contentWindow
}

// FrameBorder prop Sets or retrieves whether to display a border for the frame.
// js:"frameBorder"
func (*HTMLIFrameElement) FrameBorder() (frameBorder string) {
	js.Rewrite("$_.frameBorder")
	return frameBorder
}

// SetFrameBorder prop Sets or retrieves whether to display a border for the frame.
// js:"frameBorder"
func (*HTMLIFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$_.frameBorder = $1", frameBorder)
}

// FrameSpacing prop Sets or retrieves the amount of additional space between the frames.
// js:"frameSpacing"
func (*HTMLIFrameElement) FrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$_.frameSpacing")
	return frameSpacing
}

// SetFrameSpacing prop Sets or retrieves the amount of additional space between the frames.
// js:"frameSpacing"
func (*HTMLIFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$_.frameSpacing = $1", frameSpacing)
}

// Height prop Sets or retrieves the height of the object.
// js:"height"
func (*HTMLIFrameElement) Height() (height string) {
	js.Rewrite("$_.height")
	return height
}

// SetHeight prop Sets or retrieves the height of the object.
// js:"height"
func (*HTMLIFrameElement) SetHeight(height string) {
	js.Rewrite("$_.height = $1", height)
}

// Hspace prop Sets or retrieves the horizontal margin for the object.
// js:"hspace"
func (*HTMLIFrameElement) Hspace() (hspace int) {
	js.Rewrite("$_.hspace")
	return hspace
}

// SetHspace prop Sets or retrieves the horizontal margin for the object.
// js:"hspace"
func (*HTMLIFrameElement) SetHspace(hspace int) {
	js.Rewrite("$_.hspace = $1", hspace)
}

// LongDesc prop Sets or retrieves a URI to a long description of the object.
// js:"longDesc"
func (*HTMLIFrameElement) LongDesc() (longDesc string) {
	js.Rewrite("$_.longDesc")
	return longDesc
}

// SetLongDesc prop Sets or retrieves a URI to a long description of the object.
// js:"longDesc"
func (*HTMLIFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$_.longDesc = $1", longDesc)
}

// MarginHeight prop Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
// js:"marginHeight"
func (*HTMLIFrameElement) MarginHeight() (marginHeight string) {
	js.Rewrite("$_.marginHeight")
	return marginHeight
}

// SetMarginHeight prop Sets or retrieves the top and bottom margin heights before displaying the text in a frame.
// js:"marginHeight"
func (*HTMLIFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$_.marginHeight = $1", marginHeight)
}

// MarginWidth prop Sets or retrieves the left and right margin widths before displaying the text in a frame.
// js:"marginWidth"
func (*HTMLIFrameElement) MarginWidth() (marginWidth string) {
	js.Rewrite("$_.marginWidth")
	return marginWidth
}

// SetMarginWidth prop Sets or retrieves the left and right margin widths before displaying the text in a frame.
// js:"marginWidth"
func (*HTMLIFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$_.marginWidth = $1", marginWidth)
}

// Name prop Sets or retrieves the frame name.
// js:"name"
func (*HTMLIFrameElement) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// SetName prop Sets or retrieves the frame name.
// js:"name"
func (*HTMLIFrameElement) SetName(name string) {
	js.Rewrite("$_.name = $1", name)
}

// NoResize prop Sets or retrieves whether the user can resize the frame.
// js:"noResize"
func (*HTMLIFrameElement) NoResize() (noResize bool) {
	js.Rewrite("$_.noResize")
	return noResize
}

// SetNoResize prop Sets or retrieves whether the user can resize the frame.
// js:"noResize"
func (*HTMLIFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$_.noResize = $1", noResize)
}

// Onload prop Raised when the object has been completely received from the server.
// js:"onload"
func (*HTMLIFrameElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop Raised when the object has been completely received from the server.
// js:"onload"
func (*HTMLIFrameElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Sandbox prop
// js:"sandbox"
func (*HTMLIFrameElement) Sandbox() (sandbox *domsettabletokenlist.DOMSettableTokenList) {
	js.Rewrite("$_.sandbox")
	return sandbox
}

// Scrolling prop Sets or retrieves whether the frame can be scrolled.
// js:"scrolling"
func (*HTMLIFrameElement) Scrolling() (scrolling string) {
	js.Rewrite("$_.scrolling")
	return scrolling
}

// SetScrolling prop Sets or retrieves whether the frame can be scrolled.
// js:"scrolling"
func (*HTMLIFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$_.scrolling = $1", scrolling)
}

// Src prop Sets or retrieves a URL to be loaded by the object.
// js:"src"
func (*HTMLIFrameElement) Src() (src string) {
	js.Rewrite("$_.src")
	return src
}

// SetSrc prop Sets or retrieves a URL to be loaded by the object.
// js:"src"
func (*HTMLIFrameElement) SetSrc(src string) {
	js.Rewrite("$_.src = $1", src)
}

// Vspace prop Sets or retrieves the vertical margin for the object.
// js:"vspace"
func (*HTMLIFrameElement) Vspace() (vspace int) {
	js.Rewrite("$_.vspace")
	return vspace
}

// SetVspace prop Sets or retrieves the vertical margin for the object.
// js:"vspace"
func (*HTMLIFrameElement) SetVspace(vspace int) {
	js.Rewrite("$_.vspace = $1", vspace)
}

// Width prop Sets or retrieves the width of the object.
// js:"width"
func (*HTMLIFrameElement) Width() (width string) {
	js.Rewrite("$_.width")
	return width
}

// SetWidth prop Sets or retrieves the width of the object.
// js:"width"
func (*HTMLIFrameElement) SetWidth(width string) {
	js.Rewrite("$_.width = $1", width)
}

// AccessKey prop
// js:"accessKey"
func (*HTMLIFrameElement) AccessKey() (accessKey string) {
	js.Rewrite("$_.accessKey")
	return accessKey
}

// SetAccessKey prop
// js:"accessKey"
func (*HTMLIFrameElement) SetAccessKey(accessKey string) {
	js.Rewrite("$_.accessKey = $1", accessKey)
}

// Children prop
// js:"children"
func (*HTMLIFrameElement) Children() (children window.HTMLCollection) {
	js.Rewrite("$_.children")
	return children
}

// ContentEditable prop
// js:"contentEditable"
func (*HTMLIFrameElement) ContentEditable() (contentEditable string) {
	js.Rewrite("$_.contentEditable")
	return contentEditable
}

// SetContentEditable prop
// js:"contentEditable"
func (*HTMLIFrameElement) SetContentEditable(contentEditable string) {
	js.Rewrite("$_.contentEditable = $1", contentEditable)
}

// Dataset prop
// js:"dataset"
func (*HTMLIFrameElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$_.dataset")
	return dataset
}

// Dir prop
// js:"dir"
func (*HTMLIFrameElement) Dir() (dir string) {
	js.Rewrite("$_.dir")
	return dir
}

// SetDir prop
// js:"dir"
func (*HTMLIFrameElement) SetDir(dir string) {
	js.Rewrite("$_.dir = $1", dir)
}

// Draggable prop
// js:"draggable"
func (*HTMLIFrameElement) Draggable() (draggable bool) {
	js.Rewrite("$_.draggable")
	return draggable
}

// SetDraggable prop
// js:"draggable"
func (*HTMLIFrameElement) SetDraggable(draggable bool) {
	js.Rewrite("$_.draggable = $1", draggable)
}

// Hidden prop
// js:"hidden"
func (*HTMLIFrameElement) Hidden() (hidden bool) {
	js.Rewrite("$_.hidden")
	return hidden
}

// SetHidden prop
// js:"hidden"
func (*HTMLIFrameElement) SetHidden(hidden bool) {
	js.Rewrite("$_.hidden = $1", hidden)
}

// HideFocus prop
// js:"hideFocus"
func (*HTMLIFrameElement) HideFocus() (hideFocus bool) {
	js.Rewrite("$_.hideFocus")
	return hideFocus
}

// SetHideFocus prop
// js:"hideFocus"
func (*HTMLIFrameElement) SetHideFocus(hideFocus bool) {
	js.Rewrite("$_.hideFocus = $1", hideFocus)
}

// InnerText prop
// js:"innerText"
func (*HTMLIFrameElement) InnerText() (innerText string) {
	js.Rewrite("$_.innerText")
	return innerText
}

// SetInnerText prop
// js:"innerText"
func (*HTMLIFrameElement) SetInnerText(innerText string) {
	js.Rewrite("$_.innerText = $1", innerText)
}

// IsContentEditable prop
// js:"isContentEditable"
func (*HTMLIFrameElement) IsContentEditable() (isContentEditable bool) {
	js.Rewrite("$_.isContentEditable")
	return isContentEditable
}

// Lang prop
// js:"lang"
func (*HTMLIFrameElement) Lang() (lang string) {
	js.Rewrite("$_.lang")
	return lang
}

// SetLang prop
// js:"lang"
func (*HTMLIFrameElement) SetLang(lang string) {
	js.Rewrite("$_.lang = $1", lang)
}

// OffsetHeight prop
// js:"offsetHeight"
func (*HTMLIFrameElement) OffsetHeight() (offsetHeight int) {
	js.Rewrite("$_.offsetHeight")
	return offsetHeight
}

// OffsetLeft prop
// js:"offsetLeft"
func (*HTMLIFrameElement) OffsetLeft() (offsetLeft int) {
	js.Rewrite("$_.offsetLeft")
	return offsetLeft
}

// OffsetParent prop
// js:"offsetParent"
func (*HTMLIFrameElement) OffsetParent() (offsetParent window.Element) {
	js.Rewrite("$_.offsetParent")
	return offsetParent
}

// OffsetTop prop
// js:"offsetTop"
func (*HTMLIFrameElement) OffsetTop() (offsetTop int) {
	js.Rewrite("$_.offsetTop")
	return offsetTop
}

// OffsetWidth prop
// js:"offsetWidth"
func (*HTMLIFrameElement) OffsetWidth() (offsetWidth int) {
	js.Rewrite("$_.offsetWidth")
	return offsetWidth
}

// Onabort prop
// js:"onabort"
func (*HTMLIFrameElement) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*HTMLIFrameElement) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onactivate prop
// js:"onactivate"
func (*HTMLIFrameElement) Onactivate() (onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate")
	return onactivate
}

// SetOnactivate prop
// js:"onactivate"
func (*HTMLIFrameElement) SetOnactivate(onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate = $1", onactivate)
}

// Onbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLIFrameElement) Onbeforeactivate() (onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

// SetOnbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLIFrameElement) SetOnbeforeactivate(onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

// Onbeforecopy prop
// js:"onbeforecopy"
func (*HTMLIFrameElement) Onbeforecopy() (onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

// SetOnbeforecopy prop
// js:"onbeforecopy"
func (*HTMLIFrameElement) SetOnbeforecopy(onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

// Onbeforecut prop
// js:"onbeforecut"
func (*HTMLIFrameElement) Onbeforecut() (onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut")
	return onbeforecut
}

// SetOnbeforecut prop
// js:"onbeforecut"
func (*HTMLIFrameElement) SetOnbeforecut(onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

// Onbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLIFrameElement) Onbeforedeactivate() (onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

// SetOnbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLIFrameElement) SetOnbeforedeactivate(onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

// Onbeforepaste prop
// js:"onbeforepaste"
func (*HTMLIFrameElement) Onbeforepaste() (onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

// SetOnbeforepaste prop
// js:"onbeforepaste"
func (*HTMLIFrameElement) SetOnbeforepaste(onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

// Onblur prop
// js:"onblur"
func (*HTMLIFrameElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*HTMLIFrameElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*HTMLIFrameElement) Oncanplay() (oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*HTMLIFrameElement) SetOncanplay(oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLIFrameElement) Oncanplaythrough() (oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLIFrameElement) SetOncanplaythrough(oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*HTMLIFrameElement) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*HTMLIFrameElement) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*HTMLIFrameElement) Onclick() (onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*HTMLIFrameElement) SetOnclick(onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick = $1", onclick)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*HTMLIFrameElement) Oncontextmenu() (oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*HTMLIFrameElement) SetOncontextmenu(oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Oncopy prop
// js:"oncopy"
func (*HTMLIFrameElement) Oncopy() (oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy")
	return oncopy
}

// SetOncopy prop
// js:"oncopy"
func (*HTMLIFrameElement) SetOncopy(oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy = $1", oncopy)
}

// Oncuechange prop
// js:"oncuechange"
func (*HTMLIFrameElement) Oncuechange() (oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange")
	return oncuechange
}

// SetOncuechange prop
// js:"oncuechange"
func (*HTMLIFrameElement) SetOncuechange(oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange = $1", oncuechange)
}

// Oncut prop
// js:"oncut"
func (*HTMLIFrameElement) Oncut() (oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut")
	return oncut
}

// SetOncut prop
// js:"oncut"
func (*HTMLIFrameElement) SetOncut(oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut = $1", oncut)
}

// Ondblclick prop
// js:"ondblclick"
func (*HTMLIFrameElement) Ondblclick() (ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*HTMLIFrameElement) SetOndblclick(ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondeactivate prop
// js:"ondeactivate"
func (*HTMLIFrameElement) Ondeactivate() (ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate")
	return ondeactivate
}

// SetOndeactivate prop
// js:"ondeactivate"
func (*HTMLIFrameElement) SetOndeactivate(ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

// Ondrag prop
// js:"ondrag"
func (*HTMLIFrameElement) Ondrag() (ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*HTMLIFrameElement) SetOndrag(ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*HTMLIFrameElement) Ondragend() (ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*HTMLIFrameElement) SetOndragend(ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*HTMLIFrameElement) Ondragenter() (ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*HTMLIFrameElement) SetOndragenter(ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*HTMLIFrameElement) Ondragleave() (ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*HTMLIFrameElement) SetOndragleave(ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*HTMLIFrameElement) Ondragover() (ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*HTMLIFrameElement) SetOndragover(ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*HTMLIFrameElement) Ondragstart() (ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*HTMLIFrameElement) SetOndragstart(ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*HTMLIFrameElement) Ondrop() (ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*HTMLIFrameElement) SetOndrop(ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*HTMLIFrameElement) Ondurationchange() (ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*HTMLIFrameElement) SetOndurationchange(ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*HTMLIFrameElement) Onemptied() (onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*HTMLIFrameElement) SetOnemptied(onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*HTMLIFrameElement) Onended() (onended func(window.Event)) {
	js.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*HTMLIFrameElement) SetOnended(onended func(window.Event)) {
	js.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*HTMLIFrameElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*HTMLIFrameElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*HTMLIFrameElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*HTMLIFrameElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus = $1", onfocus)
}

// Oninput prop
// js:"oninput"
func (*HTMLIFrameElement) Oninput() (oninput func(window.Event)) {
	js.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*HTMLIFrameElement) SetOninput(oninput func(window.Event)) {
	js.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*HTMLIFrameElement) Oninvalid() (oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*HTMLIFrameElement) SetOninvalid(oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*HTMLIFrameElement) Onkeydown() (onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*HTMLIFrameElement) SetOnkeydown(onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*HTMLIFrameElement) Onkeypress() (onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*HTMLIFrameElement) SetOnkeypress(onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*HTMLIFrameElement) Onkeyup() (onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*HTMLIFrameElement) SetOnkeyup(onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*HTMLIFrameElement) Onloadeddata() (onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*HTMLIFrameElement) SetOnloadeddata(onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLIFrameElement) Onloadedmetadata() (onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLIFrameElement) SetOnloadedmetadata(onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*HTMLIFrameElement) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*HTMLIFrameElement) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmousedown prop
// js:"onmousedown"
func (*HTMLIFrameElement) Onmousedown() (onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*HTMLIFrameElement) SetOnmousedown(onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*HTMLIFrameElement) Onmouseenter() (onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*HTMLIFrameElement) SetOnmouseenter(onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*HTMLIFrameElement) Onmouseleave() (onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*HTMLIFrameElement) SetOnmouseleave(onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*HTMLIFrameElement) Onmousemove() (onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*HTMLIFrameElement) SetOnmousemove(onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*HTMLIFrameElement) Onmouseout() (onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*HTMLIFrameElement) SetOnmouseout(onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*HTMLIFrameElement) Onmouseover() (onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*HTMLIFrameElement) SetOnmouseover(onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*HTMLIFrameElement) Onmouseup() (onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*HTMLIFrameElement) SetOnmouseup(onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*HTMLIFrameElement) Onmousewheel() (onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*HTMLIFrameElement) SetOnmousewheel(onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLIFrameElement) Onmscontentzoom() (onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

// SetOnmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLIFrameElement) SetOnmscontentzoom(onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

// Onmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLIFrameElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// SetOnmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLIFrameElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

// Onpaste prop
// js:"onpaste"
func (*HTMLIFrameElement) Onpaste() (onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste")
	return onpaste
}

// SetOnpaste prop
// js:"onpaste"
func (*HTMLIFrameElement) SetOnpaste(onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste = $1", onpaste)
}

// Onpause prop
// js:"onpause"
func (*HTMLIFrameElement) Onpause() (onpause func(window.Event)) {
	js.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*HTMLIFrameElement) SetOnpause(onpause func(window.Event)) {
	js.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*HTMLIFrameElement) Onplay() (onplay func(window.Event)) {
	js.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*HTMLIFrameElement) SetOnplay(onplay func(window.Event)) {
	js.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*HTMLIFrameElement) Onplaying() (onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*HTMLIFrameElement) SetOnplaying(onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying = $1", onplaying)
}

// Onprogress prop
// js:"onprogress"
func (*HTMLIFrameElement) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*HTMLIFrameElement) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*HTMLIFrameElement) Onratechange() (onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*HTMLIFrameElement) SetOnratechange(onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreset prop
// js:"onreset"
func (*HTMLIFrameElement) Onreset() (onreset func(window.Event)) {
	js.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*HTMLIFrameElement) SetOnreset(onreset func(window.Event)) {
	js.Rewrite("$_.onreset = $1", onreset)
}

// Onscroll prop
// js:"onscroll"
func (*HTMLIFrameElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*HTMLIFrameElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*HTMLIFrameElement) Onseeked() (onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*HTMLIFrameElement) SetOnseeked(onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*HTMLIFrameElement) Onseeking() (onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*HTMLIFrameElement) SetOnseeking(onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*HTMLIFrameElement) Onselect() (onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*HTMLIFrameElement) SetOnselect(onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect = $1", onselect)
}

// Onselectstart prop
// js:"onselectstart"
func (*HTMLIFrameElement) Onselectstart() (onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart")
	return onselectstart
}

// SetOnselectstart prop
// js:"onselectstart"
func (*HTMLIFrameElement) SetOnselectstart(onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart = $1", onselectstart)
}

// Onstalled prop
// js:"onstalled"
func (*HTMLIFrameElement) Onstalled() (onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*HTMLIFrameElement) SetOnstalled(onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled = $1", onstalled)
}

// Onsubmit prop
// js:"onsubmit"
func (*HTMLIFrameElement) Onsubmit() (onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*HTMLIFrameElement) SetOnsubmit(onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*HTMLIFrameElement) Onsuspend() (onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*HTMLIFrameElement) SetOnsuspend(onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*HTMLIFrameElement) Ontimeupdate() (ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*HTMLIFrameElement) SetOntimeupdate(ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*HTMLIFrameElement) Onvolumechange() (onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*HTMLIFrameElement) SetOnvolumechange(onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*HTMLIFrameElement) Onwaiting() (onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*HTMLIFrameElement) SetOnwaiting(onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting = $1", onwaiting)
}

// OuterText prop
// js:"outerText"
func (*HTMLIFrameElement) OuterText() (outerText string) {
	js.Rewrite("$_.outerText")
	return outerText
}

// SetOuterText prop
// js:"outerText"
func (*HTMLIFrameElement) SetOuterText(outerText string) {
	js.Rewrite("$_.outerText = $1", outerText)
}

// Spellcheck prop
// js:"spellcheck"
func (*HTMLIFrameElement) Spellcheck() (spellcheck bool) {
	js.Rewrite("$_.spellcheck")
	return spellcheck
}

// SetSpellcheck prop
// js:"spellcheck"
func (*HTMLIFrameElement) SetSpellcheck(spellcheck bool) {
	js.Rewrite("$_.spellcheck = $1", spellcheck)
}

// Style prop
// js:"style"
func (*HTMLIFrameElement) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// TabIndex prop
// js:"tabIndex"
func (*HTMLIFrameElement) TabIndex() (tabIndex int8) {
	js.Rewrite("$_.tabIndex")
	return tabIndex
}

// SetTabIndex prop
// js:"tabIndex"
func (*HTMLIFrameElement) SetTabIndex(tabIndex int8) {
	js.Rewrite("$_.tabIndex = $1", tabIndex)
}

// Title prop
// js:"title"
func (*HTMLIFrameElement) Title() (title string) {
	js.Rewrite("$_.title")
	return title
}

// SetTitle prop
// js:"title"
func (*HTMLIFrameElement) SetTitle(title string) {
	js.Rewrite("$_.title = $1", title)
}

// ClassList prop
// js:"classList"
func (*HTMLIFrameElement) ClassList() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*HTMLIFrameElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLIFrameElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*HTMLIFrameElement) ClientHeight() (clientHeight int) {
	js.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*HTMLIFrameElement) ClientLeft() (clientLeft int) {
	js.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*HTMLIFrameElement) ClientTop() (clientTop int) {
	js.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*HTMLIFrameElement) ClientWidth() (clientWidth int) {
	js.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*HTMLIFrameElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLIFrameElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLIFrameElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLIFrameElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLIFrameElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLIFrameElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*HTMLIFrameElement) MsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*HTMLIFrameElement) Onariarequest() (onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*HTMLIFrameElement) SetOnariarequest(onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*HTMLIFrameElement) Oncommand() (oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*HTMLIFrameElement) SetOncommand(oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLIFrameElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLIFrameElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLIFrameElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLIFrameElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLIFrameElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLIFrameElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLIFrameElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLIFrameElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*HTMLIFrameElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*HTMLIFrameElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLIFrameElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLIFrameElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLIFrameElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLIFrameElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLIFrameElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLIFrameElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLIFrameElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLIFrameElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLIFrameElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLIFrameElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLIFrameElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLIFrameElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*HTMLIFrameElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*HTMLIFrameElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*HTMLIFrameElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*HTMLIFrameElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*HTMLIFrameElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*HTMLIFrameElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*HTMLIFrameElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*HTMLIFrameElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*HTMLIFrameElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*HTMLIFrameElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*HTMLIFrameElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*HTMLIFrameElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*HTMLIFrameElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*HTMLIFrameElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*HTMLIFrameElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*HTMLIFrameElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*HTMLIFrameElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*HTMLIFrameElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*HTMLIFrameElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*HTMLIFrameElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*HTMLIFrameElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*HTMLIFrameElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*HTMLIFrameElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*HTMLIFrameElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLIFrameElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLIFrameElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLIFrameElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLIFrameElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLIFrameElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLIFrameElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*HTMLIFrameElement) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLIFrameElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLIFrameElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLIFrameElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLIFrameElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLIFrameElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*HTMLIFrameElement) ScrollWidth() (scrollWidth int) {
	js.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*HTMLIFrameElement) TagName() (tagName string) {
	js.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*HTMLIFrameElement) Onpointercancel() (onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*HTMLIFrameElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*HTMLIFrameElement) Onpointerdown() (onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*HTMLIFrameElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*HTMLIFrameElement) Onpointerenter() (onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*HTMLIFrameElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*HTMLIFrameElement) Onpointerleave() (onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*HTMLIFrameElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*HTMLIFrameElement) Onpointermove() (onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*HTMLIFrameElement) SetOnpointermove(onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*HTMLIFrameElement) Onpointerout() (onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*HTMLIFrameElement) SetOnpointerout(onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*HTMLIFrameElement) Onpointerover() (onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*HTMLIFrameElement) SetOnpointerover(onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*HTMLIFrameElement) Onpointerup() (onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*HTMLIFrameElement) SetOnpointerup(onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*HTMLIFrameElement) Onwheel() (onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*HTMLIFrameElement) SetOnwheel(onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*HTMLIFrameElement) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*HTMLIFrameElement) FirstElementChild() (firstElementChild window.Element) {
	js.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*HTMLIFrameElement) LastElementChild() (lastElementChild window.Element) {
	js.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*HTMLIFrameElement) NextElementSibling() (nextElementSibling window.Element) {
	js.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*HTMLIFrameElement) PreviousElementSibling() (previousElementSibling window.Element) {
	js.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*HTMLIFrameElement) Attributes() (attributes *window.NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*HTMLIFrameElement) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLIFrameElement) ChildNodes() (childNodes *window.NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLIFrameElement) FirstChild() (firstChild window.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLIFrameElement) LastChild() (lastChild window.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*HTMLIFrameElement) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*HTMLIFrameElement) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLIFrameElement) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLIFrameElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLIFrameElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLIFrameElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLIFrameElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*HTMLIFrameElement) OwnerDocument() (ownerDocument window.Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*HTMLIFrameElement) ParentElement() (parentElement window.HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*HTMLIFrameElement) ParentNode() (parentNode window.Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*HTMLIFrameElement) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*HTMLIFrameElement) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*HTMLIFrameElement) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
