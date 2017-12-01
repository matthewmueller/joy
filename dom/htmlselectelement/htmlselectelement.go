package htmlselectelement

import (
	"github.com/matthewmueller/golly/dom/childnode"
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/domstringmap"
	"github.com/matthewmueller/golly/dom/domtokenlist"
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/htmloptionscollection"
	"github.com/matthewmueller/golly/dom/mszoomtooptions"
	"github.com/matthewmueller/golly/dom/validitystate"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.HTMLElement = (*HTMLSelectElement)(nil)
var _ window.Element = (*HTMLSelectElement)(nil)
var _ window.GlobalEventHandlers = (*HTMLSelectElement)(nil)
var _ window.ElementTraversal = (*HTMLSelectElement)(nil)
var _ window.NodeSelector = (*HTMLSelectElement)(nil)
var _ childnode.ChildNode = (*HTMLSelectElement)(nil)
var _ window.Node = (*HTMLSelectElement)(nil)
var _ window.EventTarget = (*HTMLSelectElement)(nil)

// HTMLSelectElement struct
// js:"HTMLSelectElement,omit"
type HTMLSelectElement struct {
}

// Add fn Adds an element to the areas, controlRange, or options collection.
//     * @param element Variant of type Number that specifies the index position in the collection where the element is placed. If no value is given, the method places the element at the end of the collection.
//     * @param before Variant of type Object that specifies an element to insert before, or null to append the object to the collection.
// js:"add"
func (*HTMLSelectElement) Add(element window.HTMLElement, before *interface{}) {
	js.Rewrite("$_.add($1, $2)", element, before)
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
// js:"checkValidity"
func (*HTMLSelectElement) CheckValidity() (b bool) {
	js.Rewrite("$_.checkValidity()")
	return b
}

// Item fn Retrieves a select object or an object from an options collection.
//     * @param name Variant of type Number or String that specifies the object or collection to retrieve. If this parameter is an integer, it is the zero-based index of the object. If this parameter is a string, all objects with matching name or id properties are retrieved, and a collection is returned if more than one match is made.
//     * @param index Variant of type Number that specifies the zero-based index of the object to retrieve when a collection is returned.
// js:"item"
func (*HTMLSelectElement) Item(name *interface{}, index *interface{}) (i interface{}) {
	js.Rewrite("$_.item($1, $2)", name, index)
	return i
}

// NamedItem fn Retrieves a select object or an object from an options collection.
//     * @param namedItem A String that specifies the name or id property of the object to retrieve. A collection is returned if more than one match is made.
// js:"namedItem"
func (*HTMLSelectElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$_.namedItem($1)", name)
	return i
}

// Remove fn Removes an element from the collection.
//     * @param index Number that specifies the zero-based index of the element to remove from the collection.
// js:"remove"
func (*HTMLSelectElement) Remove(index *int) {
	js.Rewrite("$_.remove($1)", index)
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
// js:"setCustomValidity"
func (*HTMLSelectElement) SetCustomValidity(err string) {
	js.Rewrite("$_.setCustomValidity($1)", err)
}

// Blur fn
// js:"blur"
func (*HTMLSelectElement) Blur() {
	js.Rewrite("$_.blur()")
}

// Click fn
// js:"click"
func (*HTMLSelectElement) Click() {
	js.Rewrite("$_.click()")
}

// DragDrop fn
// js:"dragDrop"
func (*HTMLSelectElement) DragDrop() (b bool) {
	js.Rewrite("$_.dragDrop()")
	return b
}

// Focus fn
// js:"focus"
func (*HTMLSelectElement) Focus() {
	js.Rewrite("$_.focus()")
}

// GetElementsByClassName fn
// js:"getElementsByClassName"
func (*HTMLSelectElement) GetElementsByClassName(classNames string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

// InsertAdjacentElement fn
// js:"insertAdjacentElement"
func (*HTMLSelectElement) InsertAdjacentElement(position string, insertedElement window.Element) (w window.Element) {
	js.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

// InsertAdjacentHTML fn
// js:"insertAdjacentHTML"
func (*HTMLSelectElement) InsertAdjacentHTML(where string, html string) {
	js.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

// InsertAdjacentText fn
// js:"insertAdjacentText"
func (*HTMLSelectElement) InsertAdjacentText(where string, text string) {
	js.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

// MsGetInputContext fn
// js:"msGetInputContext"
func (*HTMLSelectElement) MsGetInputContext() (w *window.MSInputMethodContext) {
	js.Rewrite("$_.msGetInputContext()")
	return w
}

// ScrollIntoView fn
// js:"scrollIntoView"
func (*HTMLSelectElement) ScrollIntoView(top *bool) {
	js.Rewrite("$_.scrollIntoView($1)", top)
}

// GetAttribute fn
// js:"getAttribute"
func (*HTMLSelectElement) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*HTMLSelectElement) GetAttributeNode(name string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*HTMLSelectElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*HTMLSelectElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*HTMLSelectElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*HTMLSelectElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*HTMLSelectElement) GetElementsByTagName(name string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*HTMLSelectElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*HTMLSelectElement) HasAttribute(name string) (b bool) {
	js.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*HTMLSelectElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*HTMLSelectElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	js.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*HTMLSelectElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*HTMLSelectElement) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*HTMLSelectElement) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*HTMLSelectElement) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*HTMLSelectElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*HTMLSelectElement) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*HTMLSelectElement) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*HTMLSelectElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*HTMLSelectElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*HTMLSelectElement) RequestFullscreen() {
	js.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*HTMLSelectElement) RequestPointerLock() {
	js.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*HTMLSelectElement) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*HTMLSelectElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*HTMLSelectElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*HTMLSelectElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*HTMLSelectElement) SetPointerCapture(pointerId int) {
	js.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*HTMLSelectElement) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*HTMLSelectElement) WebkitRequestFullscreen() {
	js.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*HTMLSelectElement) WebkitRequestFullScreen() {
	js.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLSelectElement) QuerySelector(selectors string) (w window.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*HTMLSelectElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*HTMLSelectElement) AppendChild(newChild window.Node) (w window.Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*HTMLSelectElement) CloneNode(deep *bool) (w window.Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*HTMLSelectElement) CompareDocumentPosition(other window.Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*HTMLSelectElement) Contains(child window.Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*HTMLSelectElement) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*HTMLSelectElement) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*HTMLSelectElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*HTMLSelectElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*HTMLSelectElement) IsEqualNode(arg window.Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*HTMLSelectElement) IsSameNode(other window.Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*HTMLSelectElement) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*HTMLSelectElement) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*HTMLSelectElement) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*HTMLSelectElement) RemoveChild(oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*HTMLSelectElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLSelectElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLSelectElement) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*HTMLSelectElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
// js:"autofocus"
func (*HTMLSelectElement) Autofocus() (autofocus bool) {
	js.Rewrite("$_.autofocus")
	return autofocus
}

// SetAutofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
// js:"autofocus"
func (*HTMLSelectElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$_.autofocus = $1", autofocus)
}

// Disabled prop
// js:"disabled"
func (*HTMLSelectElement) Disabled() (disabled bool) {
	js.Rewrite("$_.disabled")
	return disabled
}

// SetDisabled prop
// js:"disabled"
func (*HTMLSelectElement) SetDisabled(disabled bool) {
	js.Rewrite("$_.disabled = $1", disabled)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
// js:"form"
func (*HTMLSelectElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$_.form")
	return form
}

// Length prop Sets or retrieves the number of objects in a collection.
// js:"length"
func (*HTMLSelectElement) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// SetLength prop Sets or retrieves the number of objects in a collection.
// js:"length"
func (*HTMLSelectElement) SetLength(length uint) {
	js.Rewrite("$_.length = $1", length)
}

// Multiple prop Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
// js:"multiple"
func (*HTMLSelectElement) Multiple() (multiple bool) {
	js.Rewrite("$_.multiple")
	return multiple
}

// SetMultiple prop Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
// js:"multiple"
func (*HTMLSelectElement) SetMultiple(multiple bool) {
	js.Rewrite("$_.multiple = $1", multiple)
}

// Name prop Sets or retrieves the name of the object.
// js:"name"
func (*HTMLSelectElement) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// SetName prop Sets or retrieves the name of the object.
// js:"name"
func (*HTMLSelectElement) SetName(name string) {
	js.Rewrite("$_.name = $1", name)
}

// Options prop
// js:"options"
func (*HTMLSelectElement) Options() (options *htmloptionscollection.HTMLOptionsCollection) {
	js.Rewrite("$_.options")
	return options
}

// Required prop When present, marks an element that can't be submitted without a value.
// js:"required"
func (*HTMLSelectElement) Required() (required bool) {
	js.Rewrite("$_.required")
	return required
}

// SetRequired prop When present, marks an element that can't be submitted without a value.
// js:"required"
func (*HTMLSelectElement) SetRequired(required bool) {
	js.Rewrite("$_.required = $1", required)
}

// SelectedIndex prop Sets or retrieves the index of the selected option in a select object.
// js:"selectedIndex"
func (*HTMLSelectElement) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$_.selectedIndex")
	return selectedIndex
}

// SetSelectedIndex prop Sets or retrieves the index of the selected option in a select object.
// js:"selectedIndex"
func (*HTMLSelectElement) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$_.selectedIndex = $1", selectedIndex)
}

// SelectedOptions prop
// js:"selectedOptions"
func (*HTMLSelectElement) SelectedOptions() (selectedOptions window.HTMLCollection) {
	js.Rewrite("$_.selectedOptions")
	return selectedOptions
}

// Size prop Sets or retrieves the number of rows in the list box.
// js:"size"
func (*HTMLSelectElement) Size() (size uint) {
	js.Rewrite("$_.size")
	return size
}

// SetSize prop Sets or retrieves the number of rows in the list box.
// js:"size"
func (*HTMLSelectElement) SetSize(size uint) {
	js.Rewrite("$_.size = $1", size)
}

// Type prop Retrieves the type of select control based on the value of the MULTIPLE attribute.
// js:"type"
func (*HTMLSelectElement) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
// js:"validationMessage"
func (*HTMLSelectElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$_.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
// js:"validity"
func (*HTMLSelectElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$_.validity")
	return validity
}

// Value prop Sets or retrieves the value which is returned to the server when the form control is submitted.
// js:"value"
func (*HTMLSelectElement) Value() (value string) {
	js.Rewrite("$_.value")
	return value
}

// SetValue prop Sets or retrieves the value which is returned to the server when the form control is submitted.
// js:"value"
func (*HTMLSelectElement) SetValue(value string) {
	js.Rewrite("$_.value = $1", value)
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
// js:"willValidate"
func (*HTMLSelectElement) WillValidate() (willValidate bool) {
	js.Rewrite("$_.willValidate")
	return willValidate
}

// AccessKey prop
// js:"accessKey"
func (*HTMLSelectElement) AccessKey() (accessKey string) {
	js.Rewrite("$_.accessKey")
	return accessKey
}

// SetAccessKey prop
// js:"accessKey"
func (*HTMLSelectElement) SetAccessKey(accessKey string) {
	js.Rewrite("$_.accessKey = $1", accessKey)
}

// Children prop
// js:"children"
func (*HTMLSelectElement) Children() (children window.HTMLCollection) {
	js.Rewrite("$_.children")
	return children
}

// ContentEditable prop
// js:"contentEditable"
func (*HTMLSelectElement) ContentEditable() (contentEditable string) {
	js.Rewrite("$_.contentEditable")
	return contentEditable
}

// SetContentEditable prop
// js:"contentEditable"
func (*HTMLSelectElement) SetContentEditable(contentEditable string) {
	js.Rewrite("$_.contentEditable = $1", contentEditable)
}

// Dataset prop
// js:"dataset"
func (*HTMLSelectElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$_.dataset")
	return dataset
}

// Dir prop
// js:"dir"
func (*HTMLSelectElement) Dir() (dir string) {
	js.Rewrite("$_.dir")
	return dir
}

// SetDir prop
// js:"dir"
func (*HTMLSelectElement) SetDir(dir string) {
	js.Rewrite("$_.dir = $1", dir)
}

// Draggable prop
// js:"draggable"
func (*HTMLSelectElement) Draggable() (draggable bool) {
	js.Rewrite("$_.draggable")
	return draggable
}

// SetDraggable prop
// js:"draggable"
func (*HTMLSelectElement) SetDraggable(draggable bool) {
	js.Rewrite("$_.draggable = $1", draggable)
}

// Hidden prop
// js:"hidden"
func (*HTMLSelectElement) Hidden() (hidden bool) {
	js.Rewrite("$_.hidden")
	return hidden
}

// SetHidden prop
// js:"hidden"
func (*HTMLSelectElement) SetHidden(hidden bool) {
	js.Rewrite("$_.hidden = $1", hidden)
}

// HideFocus prop
// js:"hideFocus"
func (*HTMLSelectElement) HideFocus() (hideFocus bool) {
	js.Rewrite("$_.hideFocus")
	return hideFocus
}

// SetHideFocus prop
// js:"hideFocus"
func (*HTMLSelectElement) SetHideFocus(hideFocus bool) {
	js.Rewrite("$_.hideFocus = $1", hideFocus)
}

// InnerText prop
// js:"innerText"
func (*HTMLSelectElement) InnerText() (innerText string) {
	js.Rewrite("$_.innerText")
	return innerText
}

// SetInnerText prop
// js:"innerText"
func (*HTMLSelectElement) SetInnerText(innerText string) {
	js.Rewrite("$_.innerText = $1", innerText)
}

// IsContentEditable prop
// js:"isContentEditable"
func (*HTMLSelectElement) IsContentEditable() (isContentEditable bool) {
	js.Rewrite("$_.isContentEditable")
	return isContentEditable
}

// Lang prop
// js:"lang"
func (*HTMLSelectElement) Lang() (lang string) {
	js.Rewrite("$_.lang")
	return lang
}

// SetLang prop
// js:"lang"
func (*HTMLSelectElement) SetLang(lang string) {
	js.Rewrite("$_.lang = $1", lang)
}

// OffsetHeight prop
// js:"offsetHeight"
func (*HTMLSelectElement) OffsetHeight() (offsetHeight int) {
	js.Rewrite("$_.offsetHeight")
	return offsetHeight
}

// OffsetLeft prop
// js:"offsetLeft"
func (*HTMLSelectElement) OffsetLeft() (offsetLeft int) {
	js.Rewrite("$_.offsetLeft")
	return offsetLeft
}

// OffsetParent prop
// js:"offsetParent"
func (*HTMLSelectElement) OffsetParent() (offsetParent window.Element) {
	js.Rewrite("$_.offsetParent")
	return offsetParent
}

// OffsetTop prop
// js:"offsetTop"
func (*HTMLSelectElement) OffsetTop() (offsetTop int) {
	js.Rewrite("$_.offsetTop")
	return offsetTop
}

// OffsetWidth prop
// js:"offsetWidth"
func (*HTMLSelectElement) OffsetWidth() (offsetWidth int) {
	js.Rewrite("$_.offsetWidth")
	return offsetWidth
}

// Onabort prop
// js:"onabort"
func (*HTMLSelectElement) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*HTMLSelectElement) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onactivate prop
// js:"onactivate"
func (*HTMLSelectElement) Onactivate() (onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate")
	return onactivate
}

// SetOnactivate prop
// js:"onactivate"
func (*HTMLSelectElement) SetOnactivate(onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate = $1", onactivate)
}

// Onbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLSelectElement) Onbeforeactivate() (onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

// SetOnbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLSelectElement) SetOnbeforeactivate(onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

// Onbeforecopy prop
// js:"onbeforecopy"
func (*HTMLSelectElement) Onbeforecopy() (onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

// SetOnbeforecopy prop
// js:"onbeforecopy"
func (*HTMLSelectElement) SetOnbeforecopy(onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

// Onbeforecut prop
// js:"onbeforecut"
func (*HTMLSelectElement) Onbeforecut() (onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut")
	return onbeforecut
}

// SetOnbeforecut prop
// js:"onbeforecut"
func (*HTMLSelectElement) SetOnbeforecut(onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

// Onbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLSelectElement) Onbeforedeactivate() (onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

// SetOnbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLSelectElement) SetOnbeforedeactivate(onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

// Onbeforepaste prop
// js:"onbeforepaste"
func (*HTMLSelectElement) Onbeforepaste() (onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

// SetOnbeforepaste prop
// js:"onbeforepaste"
func (*HTMLSelectElement) SetOnbeforepaste(onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

// Onblur prop
// js:"onblur"
func (*HTMLSelectElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*HTMLSelectElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*HTMLSelectElement) Oncanplay() (oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*HTMLSelectElement) SetOncanplay(oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLSelectElement) Oncanplaythrough() (oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLSelectElement) SetOncanplaythrough(oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*HTMLSelectElement) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*HTMLSelectElement) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*HTMLSelectElement) Onclick() (onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*HTMLSelectElement) SetOnclick(onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick = $1", onclick)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*HTMLSelectElement) Oncontextmenu() (oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*HTMLSelectElement) SetOncontextmenu(oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Oncopy prop
// js:"oncopy"
func (*HTMLSelectElement) Oncopy() (oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy")
	return oncopy
}

// SetOncopy prop
// js:"oncopy"
func (*HTMLSelectElement) SetOncopy(oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy = $1", oncopy)
}

// Oncuechange prop
// js:"oncuechange"
func (*HTMLSelectElement) Oncuechange() (oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange")
	return oncuechange
}

// SetOncuechange prop
// js:"oncuechange"
func (*HTMLSelectElement) SetOncuechange(oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange = $1", oncuechange)
}

// Oncut prop
// js:"oncut"
func (*HTMLSelectElement) Oncut() (oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut")
	return oncut
}

// SetOncut prop
// js:"oncut"
func (*HTMLSelectElement) SetOncut(oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut = $1", oncut)
}

// Ondblclick prop
// js:"ondblclick"
func (*HTMLSelectElement) Ondblclick() (ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*HTMLSelectElement) SetOndblclick(ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondeactivate prop
// js:"ondeactivate"
func (*HTMLSelectElement) Ondeactivate() (ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate")
	return ondeactivate
}

// SetOndeactivate prop
// js:"ondeactivate"
func (*HTMLSelectElement) SetOndeactivate(ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

// Ondrag prop
// js:"ondrag"
func (*HTMLSelectElement) Ondrag() (ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*HTMLSelectElement) SetOndrag(ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*HTMLSelectElement) Ondragend() (ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*HTMLSelectElement) SetOndragend(ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*HTMLSelectElement) Ondragenter() (ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*HTMLSelectElement) SetOndragenter(ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*HTMLSelectElement) Ondragleave() (ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*HTMLSelectElement) SetOndragleave(ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*HTMLSelectElement) Ondragover() (ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*HTMLSelectElement) SetOndragover(ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*HTMLSelectElement) Ondragstart() (ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*HTMLSelectElement) SetOndragstart(ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*HTMLSelectElement) Ondrop() (ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*HTMLSelectElement) SetOndrop(ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*HTMLSelectElement) Ondurationchange() (ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*HTMLSelectElement) SetOndurationchange(ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*HTMLSelectElement) Onemptied() (onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*HTMLSelectElement) SetOnemptied(onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*HTMLSelectElement) Onended() (onended func(window.Event)) {
	js.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*HTMLSelectElement) SetOnended(onended func(window.Event)) {
	js.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*HTMLSelectElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*HTMLSelectElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*HTMLSelectElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*HTMLSelectElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus = $1", onfocus)
}

// Oninput prop
// js:"oninput"
func (*HTMLSelectElement) Oninput() (oninput func(window.Event)) {
	js.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*HTMLSelectElement) SetOninput(oninput func(window.Event)) {
	js.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*HTMLSelectElement) Oninvalid() (oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*HTMLSelectElement) SetOninvalid(oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*HTMLSelectElement) Onkeydown() (onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*HTMLSelectElement) SetOnkeydown(onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*HTMLSelectElement) Onkeypress() (onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*HTMLSelectElement) SetOnkeypress(onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*HTMLSelectElement) Onkeyup() (onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*HTMLSelectElement) SetOnkeyup(onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop
// js:"onload"
func (*HTMLSelectElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*HTMLSelectElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*HTMLSelectElement) Onloadeddata() (onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*HTMLSelectElement) SetOnloadeddata(onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLSelectElement) Onloadedmetadata() (onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLSelectElement) SetOnloadedmetadata(onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*HTMLSelectElement) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*HTMLSelectElement) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmousedown prop
// js:"onmousedown"
func (*HTMLSelectElement) Onmousedown() (onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*HTMLSelectElement) SetOnmousedown(onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*HTMLSelectElement) Onmouseenter() (onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*HTMLSelectElement) SetOnmouseenter(onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*HTMLSelectElement) Onmouseleave() (onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*HTMLSelectElement) SetOnmouseleave(onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*HTMLSelectElement) Onmousemove() (onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*HTMLSelectElement) SetOnmousemove(onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*HTMLSelectElement) Onmouseout() (onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*HTMLSelectElement) SetOnmouseout(onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*HTMLSelectElement) Onmouseover() (onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*HTMLSelectElement) SetOnmouseover(onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*HTMLSelectElement) Onmouseup() (onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*HTMLSelectElement) SetOnmouseup(onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*HTMLSelectElement) Onmousewheel() (onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*HTMLSelectElement) SetOnmousewheel(onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLSelectElement) Onmscontentzoom() (onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

// SetOnmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLSelectElement) SetOnmscontentzoom(onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

// Onmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLSelectElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// SetOnmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLSelectElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

// Onpaste prop
// js:"onpaste"
func (*HTMLSelectElement) Onpaste() (onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste")
	return onpaste
}

// SetOnpaste prop
// js:"onpaste"
func (*HTMLSelectElement) SetOnpaste(onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste = $1", onpaste)
}

// Onpause prop
// js:"onpause"
func (*HTMLSelectElement) Onpause() (onpause func(window.Event)) {
	js.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*HTMLSelectElement) SetOnpause(onpause func(window.Event)) {
	js.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*HTMLSelectElement) Onplay() (onplay func(window.Event)) {
	js.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*HTMLSelectElement) SetOnplay(onplay func(window.Event)) {
	js.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*HTMLSelectElement) Onplaying() (onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*HTMLSelectElement) SetOnplaying(onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying = $1", onplaying)
}

// Onprogress prop
// js:"onprogress"
func (*HTMLSelectElement) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*HTMLSelectElement) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*HTMLSelectElement) Onratechange() (onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*HTMLSelectElement) SetOnratechange(onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreset prop
// js:"onreset"
func (*HTMLSelectElement) Onreset() (onreset func(window.Event)) {
	js.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*HTMLSelectElement) SetOnreset(onreset func(window.Event)) {
	js.Rewrite("$_.onreset = $1", onreset)
}

// Onscroll prop
// js:"onscroll"
func (*HTMLSelectElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*HTMLSelectElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*HTMLSelectElement) Onseeked() (onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*HTMLSelectElement) SetOnseeked(onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*HTMLSelectElement) Onseeking() (onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*HTMLSelectElement) SetOnseeking(onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*HTMLSelectElement) Onselect() (onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*HTMLSelectElement) SetOnselect(onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect = $1", onselect)
}

// Onselectstart prop
// js:"onselectstart"
func (*HTMLSelectElement) Onselectstart() (onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart")
	return onselectstart
}

// SetOnselectstart prop
// js:"onselectstart"
func (*HTMLSelectElement) SetOnselectstart(onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart = $1", onselectstart)
}

// Onstalled prop
// js:"onstalled"
func (*HTMLSelectElement) Onstalled() (onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*HTMLSelectElement) SetOnstalled(onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled = $1", onstalled)
}

// Onsubmit prop
// js:"onsubmit"
func (*HTMLSelectElement) Onsubmit() (onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*HTMLSelectElement) SetOnsubmit(onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*HTMLSelectElement) Onsuspend() (onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*HTMLSelectElement) SetOnsuspend(onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*HTMLSelectElement) Ontimeupdate() (ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*HTMLSelectElement) SetOntimeupdate(ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*HTMLSelectElement) Onvolumechange() (onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*HTMLSelectElement) SetOnvolumechange(onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*HTMLSelectElement) Onwaiting() (onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*HTMLSelectElement) SetOnwaiting(onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting = $1", onwaiting)
}

// OuterText prop
// js:"outerText"
func (*HTMLSelectElement) OuterText() (outerText string) {
	js.Rewrite("$_.outerText")
	return outerText
}

// SetOuterText prop
// js:"outerText"
func (*HTMLSelectElement) SetOuterText(outerText string) {
	js.Rewrite("$_.outerText = $1", outerText)
}

// Spellcheck prop
// js:"spellcheck"
func (*HTMLSelectElement) Spellcheck() (spellcheck bool) {
	js.Rewrite("$_.spellcheck")
	return spellcheck
}

// SetSpellcheck prop
// js:"spellcheck"
func (*HTMLSelectElement) SetSpellcheck(spellcheck bool) {
	js.Rewrite("$_.spellcheck = $1", spellcheck)
}

// Style prop
// js:"style"
func (*HTMLSelectElement) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// TabIndex prop
// js:"tabIndex"
func (*HTMLSelectElement) TabIndex() (tabIndex int8) {
	js.Rewrite("$_.tabIndex")
	return tabIndex
}

// SetTabIndex prop
// js:"tabIndex"
func (*HTMLSelectElement) SetTabIndex(tabIndex int8) {
	js.Rewrite("$_.tabIndex = $1", tabIndex)
}

// Title prop
// js:"title"
func (*HTMLSelectElement) Title() (title string) {
	js.Rewrite("$_.title")
	return title
}

// SetTitle prop
// js:"title"
func (*HTMLSelectElement) SetTitle(title string) {
	js.Rewrite("$_.title = $1", title)
}

// ClassList prop
// js:"classList"
func (*HTMLSelectElement) ClassList() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*HTMLSelectElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLSelectElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*HTMLSelectElement) ClientHeight() (clientHeight int) {
	js.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*HTMLSelectElement) ClientLeft() (clientLeft int) {
	js.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*HTMLSelectElement) ClientTop() (clientTop int) {
	js.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*HTMLSelectElement) ClientWidth() (clientWidth int) {
	js.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*HTMLSelectElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLSelectElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLSelectElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLSelectElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLSelectElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLSelectElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*HTMLSelectElement) MsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*HTMLSelectElement) Onariarequest() (onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*HTMLSelectElement) SetOnariarequest(onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*HTMLSelectElement) Oncommand() (oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*HTMLSelectElement) SetOncommand(oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLSelectElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLSelectElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLSelectElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLSelectElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLSelectElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLSelectElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLSelectElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLSelectElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*HTMLSelectElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*HTMLSelectElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLSelectElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLSelectElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLSelectElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLSelectElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLSelectElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLSelectElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLSelectElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLSelectElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLSelectElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLSelectElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLSelectElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLSelectElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*HTMLSelectElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*HTMLSelectElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*HTMLSelectElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*HTMLSelectElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*HTMLSelectElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*HTMLSelectElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*HTMLSelectElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*HTMLSelectElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*HTMLSelectElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*HTMLSelectElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*HTMLSelectElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*HTMLSelectElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*HTMLSelectElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*HTMLSelectElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*HTMLSelectElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*HTMLSelectElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*HTMLSelectElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*HTMLSelectElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*HTMLSelectElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*HTMLSelectElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*HTMLSelectElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*HTMLSelectElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*HTMLSelectElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*HTMLSelectElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLSelectElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLSelectElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLSelectElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLSelectElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLSelectElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLSelectElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*HTMLSelectElement) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLSelectElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLSelectElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLSelectElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLSelectElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLSelectElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*HTMLSelectElement) ScrollWidth() (scrollWidth int) {
	js.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*HTMLSelectElement) TagName() (tagName string) {
	js.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*HTMLSelectElement) Onpointercancel() (onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*HTMLSelectElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*HTMLSelectElement) Onpointerdown() (onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*HTMLSelectElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*HTMLSelectElement) Onpointerenter() (onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*HTMLSelectElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*HTMLSelectElement) Onpointerleave() (onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*HTMLSelectElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*HTMLSelectElement) Onpointermove() (onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*HTMLSelectElement) SetOnpointermove(onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*HTMLSelectElement) Onpointerout() (onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*HTMLSelectElement) SetOnpointerout(onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*HTMLSelectElement) Onpointerover() (onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*HTMLSelectElement) SetOnpointerover(onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*HTMLSelectElement) Onpointerup() (onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*HTMLSelectElement) SetOnpointerup(onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*HTMLSelectElement) Onwheel() (onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*HTMLSelectElement) SetOnwheel(onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*HTMLSelectElement) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*HTMLSelectElement) FirstElementChild() (firstElementChild window.Element) {
	js.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*HTMLSelectElement) LastElementChild() (lastElementChild window.Element) {
	js.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*HTMLSelectElement) NextElementSibling() (nextElementSibling window.Element) {
	js.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*HTMLSelectElement) PreviousElementSibling() (previousElementSibling window.Element) {
	js.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*HTMLSelectElement) Attributes() (attributes *window.NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*HTMLSelectElement) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLSelectElement) ChildNodes() (childNodes *window.NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLSelectElement) FirstChild() (firstChild window.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLSelectElement) LastChild() (lastChild window.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*HTMLSelectElement) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*HTMLSelectElement) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLSelectElement) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLSelectElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLSelectElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLSelectElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLSelectElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*HTMLSelectElement) OwnerDocument() (ownerDocument window.Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*HTMLSelectElement) ParentElement() (parentElement window.HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*HTMLSelectElement) ParentNode() (parentNode window.Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*HTMLSelectElement) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*HTMLSelectElement) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*HTMLSelectElement) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
