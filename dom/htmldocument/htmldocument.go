package htmldocument

import (
	"github.com/matthewmueller/joy/dom/htmlallcollection"
	"github.com/matthewmueller/joy/dom/location"
	"github.com/matthewmueller/joy/dom/visibilitystate"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/xpathnsresolver"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Document = (*HTMLDocument)(nil)
var _ window.GlobalEventHandlers = (*HTMLDocument)(nil)
var _ window.NodeSelector = (*HTMLDocument)(nil)
var _ window.DocumentEvent = (*HTMLDocument)(nil)
var _ window.Node = (*HTMLDocument)(nil)
var _ window.EventTarget = (*HTMLDocument)(nil)

// HTMLDocument struct
// js:"HTMLDocument,omit"
type HTMLDocument struct {
}

// AdoptNode fn
// js:"adoptNode"
func (*HTMLDocument) AdoptNode(source window.Node) (w window.Node) {
	macro.Rewrite("$_.adoptNode($1)", source)
	return w
}

// CaptureEvents fn
// js:"captureEvents"
func (*HTMLDocument) CaptureEvents() {
	macro.Rewrite("$_.captureEvents()")
}

// CaretRangeFromPoint fn
// js:"caretRangeFromPoint"
func (*HTMLDocument) CaretRangeFromPoint(x float32, y float32) (w *window.Range) {
	macro.Rewrite("$_.caretRangeFromPoint($1, $2)", x, y)
	return w
}

// Clear fn
// js:"clear"
func (*HTMLDocument) Clear() {
	macro.Rewrite("$_.clear()")
}

// Close fn Closes an output stream and forces the sent data to display.
// js:"close"
func (*HTMLDocument) Close() {
	macro.Rewrite("$_.close()")
}

// CreateAttribute fn Creates an attribute object with a specified name.
//     * @param name String that sets the attribute object's name.
// js:"createAttribute"
func (*HTMLDocument) CreateAttribute(name string) (w *window.Attr) {
	macro.Rewrite("$_.createAttribute($1)", name)
	return w
}

// CreateAttributeNS fn
// js:"createAttributeNS"
func (*HTMLDocument) CreateAttributeNS(namespaceURI string, qualifiedName string) (w *window.Attr) {
	macro.Rewrite("$_.createAttributeNS($1, $2)", namespaceURI, qualifiedName)
	return w
}

// CreateCDATASection fn
// js:"createCDATASection"
func (*HTMLDocument) CreateCDATASection(data string) (w *window.CDATASection) {
	macro.Rewrite("$_.createCDATASection($1)", data)
	return w
}

// CreateComment fn Creates a comment object with the specified data.
//     * @param data Sets the comment object's data.
// js:"createComment"
func (*HTMLDocument) CreateComment(data string) (w *window.Comment) {
	macro.Rewrite("$_.createComment($1)", data)
	return w
}

// CreateDocumentFragment fn Creates a new document.
// js:"createDocumentFragment"
func (*HTMLDocument) CreateDocumentFragment() (w *window.DocumentFragment) {
	macro.Rewrite("$_.createDocumentFragment()")
	return w
}

// CreateElement fn Creates an instance of the element for the specified tag.
//     * @param tagName The name of an element.
// js:"createElement"
func (*HTMLDocument) CreateElement(tagName string) (w window.Element) {
	macro.Rewrite("$_.createElement($1)", tagName)
	return w
}

// CreateElementNS fn
// js:"createElementNS"
func (*HTMLDocument) CreateElementNS(namespaceURI string, qualifiedName string) (w window.Element) {
	macro.Rewrite("$_.createElementNS($1, $2)", namespaceURI, qualifiedName)
	return w
}

// CreateExpression fn
// js:"createExpression"
func (*HTMLDocument) CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (w *window.XPathExpression) {
	macro.Rewrite("$_.createExpression($1, $2)", expression, resolver)
	return w
}

// CreateNodeIterator fn Creates a NodeIterator object that you can use to traverse filtered lists of nodes or elements in a document.
//     * @param root The root element or node to start traversing on.
//     * @param whatToShow The type of nodes or elements to appear in the node list
//     * @param filter A custom NodeFilter function to use. For more information, see filter. Use null for no filter.
//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
// js:"createNodeIterator"
func (*HTMLDocument) CreateNodeIterator(root window.Node, whatToShow *uint, filter *window.NodeFilter, entityReferenceExpansion *bool) (w *window.NodeIterator) {
	macro.Rewrite("$_.createNodeIterator($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return w
}

// CreateNSResolver fn
// js:"createNSResolver"
func (*HTMLDocument) CreateNSResolver(nodeResolver window.Node) (x *xpathnsresolver.XPathNSResolver) {
	macro.Rewrite("$_.createNSResolver($1)", nodeResolver)
	return x
}

// CreateProcessingInstruction fn
// js:"createProcessingInstruction"
func (*HTMLDocument) CreateProcessingInstruction(target string, data string) (w *window.ProcessingInstruction) {
	macro.Rewrite("$_.createProcessingInstruction($1, $2)", target, data)
	return w
}

// CreateRange fn Returns an empty range object that has both of its boundary points positioned at the beginning of the document.
// js:"createRange"
func (*HTMLDocument) CreateRange() (w *window.Range) {
	macro.Rewrite("$_.createRange()")
	return w
}

// CreateTextNode fn Creates a text string from the specified value.
//     * @param data String that specifies the nodeValue property of the text node.
// js:"createTextNode"
func (*HTMLDocument) CreateTextNode(data string) (w window.Text) {
	macro.Rewrite("$_.createTextNode($1)", data)
	return w
}

// CreateTouch fn
// js:"createTouch"
func (*HTMLDocument) CreateTouch(view *window.Window, target window.EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (w *window.Touch) {
	macro.Rewrite("$_.createTouch($1, $2, $3, $4, $5, $6, $7)", view, target, identifier, pageX, pageY, screenX, screenY)
	return w
}

// CreateTouchList fn
// js:"createTouchList"
func (*HTMLDocument) CreateTouchList(touches *window.Touch) (w *window.TouchList) {
	macro.Rewrite("$_.createTouchList($1)", touches)
	return w
}

// CreateTreeWalker fn Creates a TreeWalker object that you can use to traverse filtered lists of nodes or elements in a document.
//     * @param root The root element or node to start traversing on.
//     * @param whatToShow The type of nodes or elements to appear in the node list. For more information, see whatToShow.
//     * @param filter A custom NodeFilter function to use.
//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
// js:"createTreeWalker"
func (*HTMLDocument) CreateTreeWalker(root window.Node, whatToShow *uint, filter *window.NodeFilter, entityReferenceExpansion *bool) (w *window.TreeWalker) {
	macro.Rewrite("$_.createTreeWalker($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return w
}

// ElementFromPoint fn Returns the element for the specified x coordinate and the specified y coordinate.
//     * @param x The x-offset
//     * @param y The y-offset
// js:"elementFromPoint"
func (*HTMLDocument) ElementFromPoint(x int, y int) (w window.Element) {
	macro.Rewrite("$_.elementFromPoint($1, $2)", x, y)
	return w
}

// Evaluate fn
// js:"evaluate"
func (*HTMLDocument) Evaluate(expression string, contextNode window.Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *window.XPathResult) (w *window.XPathResult) {
	macro.Rewrite("$_.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}

// ExecCommand fn Executes a command on the current document, current selection, or the given range.
//     * @param commandId String that specifies the command to execute. This command can be any of the command identifiers that can be executed in script.
//     * @param showUI Display the user interface, defaults to false.
//     * @param value Value to assign.
// js:"execCommand"
func (*HTMLDocument) ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool) {
	macro.Rewrite("$_.execCommand($1, $2, $3)", commandId, showUI, value)
	return b
}

// ExecCommandShowHelp fn Displays help information for the given command identifier.
//     * @param commandId Displays help information for the given command identifier.
// js:"execCommandShowHelp"
func (*HTMLDocument) ExecCommandShowHelp(commandId string) (b bool) {
	macro.Rewrite("$_.execCommandShowHelp($1)", commandId)
	return b
}

// ExitFullscreen fn
// js:"exitFullscreen"
func (*HTMLDocument) ExitFullscreen() {
	macro.Rewrite("$_.exitFullscreen()")
}

// ExitPointerLock fn
// js:"exitPointerLock"
func (*HTMLDocument) ExitPointerLock() {
	macro.Rewrite("$_.exitPointerLock()")
}

// Focus fn Causes the element to receive the focus and executes the code specified by the onfocus event.
// js:"focus"
func (*HTMLDocument) Focus() {
	macro.Rewrite("$_.focus()")
}

// GetElementByID fn Returns a reference to the first object with the specified value of the ID or NAME attribute.
//     * @param elementId String that specifies the ID value. Case-insensitive.
// js:"getElementById"
func (*HTMLDocument) GetElementByID(elementId string) (w window.Element) {
	macro.Rewrite("$_.getElementById($1)", elementId)
	return w
}

// GetElementsByClassName fn
// js:"getElementsByClassName"
func (*HTMLDocument) GetElementsByClassName(classNames string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

// GetElementsByName fn Gets a collection of objects based on the value of the NAME or ID attribute.
//     * @param elementName Gets a collection of objects based on the value of the NAME or ID attribute.
// js:"getElementsByName"
func (*HTMLDocument) GetElementsByName(elementName string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByName($1)", elementName)
	return w
}

// GetElementsByTagName fn Retrieves a collection of objects based on the specified element name.
//     * @param name Specifies the name of an element.
// js:"getElementsByTagName"
func (*HTMLDocument) GetElementsByTagName(tagname string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagName($1)", tagname)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*HTMLDocument) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	macro.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// GetSelection fn Returns an object representing the current selection of the document that is loaded into the object displaying a webpage.
// js:"getSelection"
func (*HTMLDocument) GetSelection() (w *window.Selection) {
	macro.Rewrite("$_.getSelection()")
	return w
}

// HasFocus fn Gets a value indicating whether the object currently has focus.
// js:"hasFocus"
func (*HTMLDocument) HasFocus() (b bool) {
	macro.Rewrite("$_.hasFocus()")
	return b
}

// ImportNode fn
// js:"importNode"
func (*HTMLDocument) ImportNode(importedNode window.Node, deep bool) (w window.Node) {
	macro.Rewrite("$_.importNode($1, $2)", importedNode, deep)
	return w
}

// MsElementsFromPoint fn
// js:"msElementsFromPoint"
func (*HTMLDocument) MsElementsFromPoint(x float32, y float32) (w *window.NodeList) {
	macro.Rewrite("$_.msElementsFromPoint($1, $2)", x, y)
	return w
}

// MsElementsFromRect fn
// js:"msElementsFromRect"
func (*HTMLDocument) MsElementsFromRect(left float32, top float32, width float32, height float32) (w *window.NodeList) {
	macro.Rewrite("$_.msElementsFromRect($1, $2, $3, $4)", left, top, width, height)
	return w
}

// Open fn Opens a new window and loads a document specified by a given URL. Also, opens a new window that uses the url parameter and the name parameter to collect the output of the write method and the writeln method.
//     * @param url Specifies a MIME type for the document.
//     * @param name Specifies the name of the window. This name is used as the value for the TARGET attribute on a form or an anchor element.
//     * @param features Contains a list of items separated by commas. Each item consists of an option and a value, separated by an equals sign (for example, "fullscreen=yes, toolbar=yes"). The following values are supported.
//     * @param replace Specifies whether the existing entry for the document is replaced in the history list.
// js:"open"
func (*HTMLDocument) Open(url *string, name *string, features *string, replace *bool) (i interface{}) {
	macro.Rewrite("$_.open($1, $2, $3, $4)", url, name, features, replace)
	return i
}

// QueryCommandEnabled fn Returns a Boolean value that indicates whether a specified command can be successfully executed using execCommand, given the current state of the document.
//     * @param commandId Specifies a command identifier.
// js:"queryCommandEnabled"
func (*HTMLDocument) QueryCommandEnabled(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandEnabled($1)", commandId)
	return b
}

// QueryCommandIndeterm fn Returns a Boolean value that indicates whether the specified command is in the indeterminate state.
//     * @param commandId String that specifies a command identifier.
// js:"queryCommandIndeterm"
func (*HTMLDocument) QueryCommandIndeterm(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandIndeterm($1)", commandId)
	return b
}

// QueryCommandState fn Returns a Boolean value that indicates the current state of the command.
//     * @param commandId String that specifies a command identifier.
// js:"queryCommandState"
func (*HTMLDocument) QueryCommandState(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandState($1)", commandId)
	return b
}

// QueryCommandSupported fn Returns a Boolean value that indicates whether the current command is supported on the current range.
//     * @param commandId Specifies a command identifier.
// js:"queryCommandSupported"
func (*HTMLDocument) QueryCommandSupported(commandId string) (b bool) {
	macro.Rewrite("$_.queryCommandSupported($1)", commandId)
	return b
}

// QueryCommandText fn Retrieves the string associated with a command.
//     * @param commandId String that contains the identifier of a command. This can be any command identifier given in the list of Command Identifiers.
// js:"queryCommandText"
func (*HTMLDocument) QueryCommandText(commandId string) (s string) {
	macro.Rewrite("$_.queryCommandText($1)", commandId)
	return s
}

// QueryCommandValue fn Returns the current value of the document, range, or current selection for the given command.
//     * @param commandId String that specifies a command identifier.
// js:"queryCommandValue"
func (*HTMLDocument) QueryCommandValue(commandId string) (s string) {
	macro.Rewrite("$_.queryCommandValue($1)", commandId)
	return s
}

// ReleaseEvents fn
// js:"releaseEvents"
func (*HTMLDocument) ReleaseEvents() {
	macro.Rewrite("$_.releaseEvents()")
}

// UpdateSettings fn Allows updating the print settings for the page.
// js:"updateSettings"
func (*HTMLDocument) UpdateSettings() {
	macro.Rewrite("$_.updateSettings()")
}

// WebkitCancelFullScreen fn
// js:"webkitCancelFullScreen"
func (*HTMLDocument) WebkitCancelFullScreen() {
	macro.Rewrite("$_.webkitCancelFullScreen()")
}

// WebkitExitFullscreen fn
// js:"webkitExitFullscreen"
func (*HTMLDocument) WebkitExitFullscreen() {
	macro.Rewrite("$_.webkitExitFullscreen()")
}

// Write fn Writes one or more HTML expressions to a document in the specified window.
//     * @param content Specifies the text and HTML tags to write.
// js:"write"
func (*HTMLDocument) Write(content string) {
	macro.Rewrite("$_.write($1)", content)
}

// Writeln fn Writes one or more HTML expressions, followed by a carriage return, to a document in the specified window.
//     * @param content The text and HTML tags to write.
// js:"writeln"
func (*HTMLDocument) Writeln(content string) {
	macro.Rewrite("$_.writeln($1)", content)
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLDocument) QuerySelector(selectors string) (w window.Element) {
	macro.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*HTMLDocument) QuerySelectorAll(selectors string) (w *window.NodeList) {
	macro.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// CreateEvent fn
// js:"createEvent"
func (*HTMLDocument) CreateEvent(eventInterface string) (w window.Event) {
	macro.Rewrite("$_.createEvent($1)", eventInterface)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*HTMLDocument) AppendChild(newChild window.Node) (w window.Node) {
	macro.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*HTMLDocument) CloneNode(deep *bool) (w window.Node) {
	macro.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*HTMLDocument) CompareDocumentPosition(other window.Node) (u uint8) {
	macro.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*HTMLDocument) Contains(child window.Node) (b bool) {
	macro.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*HTMLDocument) HasAttributes() (b bool) {
	macro.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*HTMLDocument) HasChildNodes() (b bool) {
	macro.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*HTMLDocument) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	macro.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*HTMLDocument) IsDefaultNamespace(namespaceURI string) (b bool) {
	macro.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*HTMLDocument) IsEqualNode(arg window.Node) (b bool) {
	macro.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*HTMLDocument) IsSameNode(other window.Node) (b bool) {
	macro.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*HTMLDocument) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*HTMLDocument) LookupPrefix(namespaceURI string) (s string) {
	macro.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*HTMLDocument) Normalize() {
	macro.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*HTMLDocument) RemoveChild(oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*HTMLDocument) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	macro.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLDocument) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLDocument) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*HTMLDocument) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ActiveElement prop Gets the object that has the focus when the parent document has focus.
// js:"activeElement"
func (*HTMLDocument) ActiveElement() (activeElement window.Element) {
	macro.Rewrite("$_.activeElement")
	return activeElement
}

// AlinkColor prop Sets or gets the color of all active links in the document.
// js:"alinkColor"
func (*HTMLDocument) AlinkColor() (alinkColor string) {
	macro.Rewrite("$_.alinkColor")
	return alinkColor
}

// SetAlinkColor prop Sets or gets the color of all active links in the document.
// js:"alinkColor"
func (*HTMLDocument) SetAlinkColor(alinkColor string) {
	macro.Rewrite("$_.alinkColor = $1", alinkColor)
}

// All prop Returns a reference to the collection of elements contained by the object.
// js:"all"
func (*HTMLDocument) All() (all *htmlallcollection.HTMLAllCollection) {
	macro.Rewrite("$_.all")
	return all
}

// Anchors prop Retrieves a collection of all a objects that have a name and/or id property. Objects in this collection are in HTML source order.
// js:"anchors"
func (*HTMLDocument) Anchors() (anchors window.HTMLCollection) {
	macro.Rewrite("$_.anchors")
	return anchors
}

// Applets prop Retrieves a collection of all applet objects in the document.
// js:"applets"
func (*HTMLDocument) Applets() (applets window.HTMLCollection) {
	macro.Rewrite("$_.applets")
	return applets
}

// BgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
// js:"bgColor"
func (*HTMLDocument) BgColor() (bgColor string) {
	macro.Rewrite("$_.bgColor")
	return bgColor
}

// SetBgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
// js:"bgColor"
func (*HTMLDocument) SetBgColor(bgColor string) {
	macro.Rewrite("$_.bgColor = $1", bgColor)
}

// Body prop Specifies the beginning and end of the document body.
// js:"body"
func (*HTMLDocument) Body() (body window.HTMLElement) {
	macro.Rewrite("$_.body")
	return body
}

// SetBody prop Specifies the beginning and end of the document body.
// js:"body"
func (*HTMLDocument) SetBody(body window.HTMLElement) {
	macro.Rewrite("$_.body = $1", body)
}

// CharacterSet prop
// js:"characterSet"
func (*HTMLDocument) CharacterSet() (characterSet string) {
	macro.Rewrite("$_.characterSet")
	return characterSet
}

// Charset prop Gets or sets the character set used to encode the object.
// js:"charset"
func (*HTMLDocument) Charset() (charset string) {
	macro.Rewrite("$_.charset")
	return charset
}

// SetCharset prop Gets or sets the character set used to encode the object.
// js:"charset"
func (*HTMLDocument) SetCharset(charset string) {
	macro.Rewrite("$_.charset = $1", charset)
}

// CompatMode prop Gets a value that indicates whether standards-compliant mode is switched on for the object.
// js:"compatMode"
func (*HTMLDocument) CompatMode() (compatMode string) {
	macro.Rewrite("$_.compatMode")
	return compatMode
}

// Cookie prop
// js:"cookie"
func (*HTMLDocument) Cookie() (cookie string) {
	macro.Rewrite("$_.cookie")
	return cookie
}

// SetCookie prop
// js:"cookie"
func (*HTMLDocument) SetCookie(cookie string) {
	macro.Rewrite("$_.cookie = $1", cookie)
}

// CurrentScript prop
// js:"currentScript"
func (*HTMLDocument) CurrentScript() (currentScript interface{}) {
	macro.Rewrite("$_.currentScript")
	return currentScript
}

// DefaultView prop
// js:"defaultView"
func (*HTMLDocument) DefaultView() (defaultView *window.Window) {
	macro.Rewrite("$_.defaultView")
	return defaultView
}

// DesignMode prop Sets or gets a value that indicates whether the document can be edited.
// js:"designMode"
func (*HTMLDocument) DesignMode() (designMode string) {
	macro.Rewrite("$_.designMode")
	return designMode
}

// SetDesignMode prop Sets or gets a value that indicates whether the document can be edited.
// js:"designMode"
func (*HTMLDocument) SetDesignMode(designMode string) {
	macro.Rewrite("$_.designMode = $1", designMode)
}

// Dir prop Sets or retrieves a value that indicates the reading order of the object.
// js:"dir"
func (*HTMLDocument) Dir() (dir string) {
	macro.Rewrite("$_.dir")
	return dir
}

// SetDir prop Sets or retrieves a value that indicates the reading order of the object.
// js:"dir"
func (*HTMLDocument) SetDir(dir string) {
	macro.Rewrite("$_.dir = $1", dir)
}

// Doctype prop Gets an object representing the document type declaration associated with the current document.
// js:"doctype"
func (*HTMLDocument) Doctype() (doctype *window.DocumentType) {
	macro.Rewrite("$_.doctype")
	return doctype
}

// DocumentElement prop Gets a reference to the root node of the document.
// js:"documentElement"
func (*HTMLDocument) DocumentElement() (documentElement window.Element) {
	macro.Rewrite("$_.documentElement")
	return documentElement
}

// Domain prop Sets or gets the security domain of the document.
// js:"domain"
func (*HTMLDocument) Domain() (domain string) {
	macro.Rewrite("$_.domain")
	return domain
}

// SetDomain prop Sets or gets the security domain of the document.
// js:"domain"
func (*HTMLDocument) SetDomain(domain string) {
	macro.Rewrite("$_.domain = $1", domain)
}

// Embeds prop Retrieves a collection of all embed objects in the document.
// js:"embeds"
func (*HTMLDocument) Embeds() (embeds window.HTMLCollection) {
	macro.Rewrite("$_.embeds")
	return embeds
}

// FgColor prop Sets or gets the foreground (text) color of the document.
// js:"fgColor"
func (*HTMLDocument) FgColor() (fgColor string) {
	macro.Rewrite("$_.fgColor")
	return fgColor
}

// SetFgColor prop Sets or gets the foreground (text) color of the document.
// js:"fgColor"
func (*HTMLDocument) SetFgColor(fgColor string) {
	macro.Rewrite("$_.fgColor = $1", fgColor)
}

// Forms prop Retrieves a collection, in source order, of all form objects in the document.
// js:"forms"
func (*HTMLDocument) Forms() (forms window.HTMLCollection) {
	macro.Rewrite("$_.forms")
	return forms
}

// FullscreenElement prop
// js:"fullscreenElement"
func (*HTMLDocument) FullscreenElement() (fullscreenElement window.Element) {
	macro.Rewrite("$_.fullscreenElement")
	return fullscreenElement
}

// FullscreenEnabled prop
// js:"fullscreenEnabled"
func (*HTMLDocument) FullscreenEnabled() (fullscreenEnabled bool) {
	macro.Rewrite("$_.fullscreenEnabled")
	return fullscreenEnabled
}

// Head prop
// js:"head"
func (*HTMLDocument) Head() (head *window.HTMLHeadElement) {
	macro.Rewrite("$_.head")
	return head
}

// Hidden prop
// js:"hidden"
func (*HTMLDocument) Hidden() (hidden bool) {
	macro.Rewrite("$_.hidden")
	return hidden
}

// Images prop Retrieves a collection, in source order, of img objects in the document.
// js:"images"
func (*HTMLDocument) Images() (images window.HTMLCollection) {
	macro.Rewrite("$_.images")
	return images
}

// Implementation prop Gets the implementation object of the current document.
// js:"implementation"
func (*HTMLDocument) Implementation() (implementation *window.DOMImplementation) {
	macro.Rewrite("$_.implementation")
	return implementation
}

// InputEncoding prop Returns the character encoding used to create the webpage that is loaded into the document object.
// js:"inputEncoding"
func (*HTMLDocument) InputEncoding() (inputEncoding string) {
	macro.Rewrite("$_.inputEncoding")
	return inputEncoding
}

// LastModified prop Gets the date that the page was last modified, if the page supplies one.
// js:"lastModified"
func (*HTMLDocument) LastModified() (lastModified string) {
	macro.Rewrite("$_.lastModified")
	return lastModified
}

// LinkColor prop Sets or gets the color of the document links.
// js:"linkColor"
func (*HTMLDocument) LinkColor() (linkColor string) {
	macro.Rewrite("$_.linkColor")
	return linkColor
}

// SetLinkColor prop Sets or gets the color of the document links.
// js:"linkColor"
func (*HTMLDocument) SetLinkColor(linkColor string) {
	macro.Rewrite("$_.linkColor = $1", linkColor)
}

// Links prop Retrieves a collection of all a objects that specify the href property and all area objects in the document.
// js:"links"
func (*HTMLDocument) Links() (links window.HTMLCollection) {
	macro.Rewrite("$_.links")
	return links
}

// Location prop Contains information about the current URL.
// js:"location"
func (*HTMLDocument) Location() (location *location.Location) {
	macro.Rewrite("$_.location")
	return location
}

// MsCapsLockWarningOff prop
// js:"msCapsLockWarningOff"
func (*HTMLDocument) MsCapsLockWarningOff() (msCapsLockWarningOff bool) {
	macro.Rewrite("$_.msCapsLockWarningOff")
	return msCapsLockWarningOff
}

// SetMsCapsLockWarningOff prop
// js:"msCapsLockWarningOff"
func (*HTMLDocument) SetMsCapsLockWarningOff(msCapsLockWarningOff bool) {
	macro.Rewrite("$_.msCapsLockWarningOff = $1", msCapsLockWarningOff)
}

// MsCSSOMElementFloatMetrics prop
// js:"msCSSOMElementFloatMetrics"
func (*HTMLDocument) MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool) {
	macro.Rewrite("$_.msCSSOMElementFloatMetrics")
	return msCSSOMElementFloatMetrics
}

// SetMsCSSOMElementFloatMetrics prop
// js:"msCSSOMElementFloatMetrics"
func (*HTMLDocument) SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool) {
	macro.Rewrite("$_.msCSSOMElementFloatMetrics = $1", msCSSOMElementFloatMetrics)
}

// Onabort prop Fires when the user aborts the download.
//     * @param ev The event.
// js:"onabort"
func (*HTMLDocument) Onabort() (onabort func(window.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop Fires when the user aborts the download.
//     * @param ev The event.
// js:"onabort"
func (*HTMLDocument) SetOnabort(onabort func(window.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onactivate prop Fires when the object is set as the active element.
//     * @param ev The event.
// js:"onactivate"
func (*HTMLDocument) Onactivate() (onactivate func(window.Event)) {
	macro.Rewrite("$_.onactivate")
	return onactivate
}

// SetOnactivate prop Fires when the object is set as the active element.
//     * @param ev The event.
// js:"onactivate"
func (*HTMLDocument) SetOnactivate(onactivate func(window.Event)) {
	macro.Rewrite("$_.onactivate = $1", onactivate)
}

// Onbeforeactivate prop Fires immediately before the object is set as the active element.
//     * @param ev The event.
// js:"onbeforeactivate"
func (*HTMLDocument) Onbeforeactivate() (onbeforeactivate func(window.Event)) {
	macro.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

// SetOnbeforeactivate prop Fires immediately before the object is set as the active element.
//     * @param ev The event.
// js:"onbeforeactivate"
func (*HTMLDocument) SetOnbeforeactivate(onbeforeactivate func(window.Event)) {
	macro.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

// Onbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The event.
// js:"onbeforedeactivate"
func (*HTMLDocument) Onbeforedeactivate() (onbeforedeactivate func(window.Event)) {
	macro.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

// SetOnbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The event.
// js:"onbeforedeactivate"
func (*HTMLDocument) SetOnbeforedeactivate(onbeforedeactivate func(window.Event)) {
	macro.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

// Onblur prop Fires when the object loses the input focus.
//     * @param ev The focus event.
// js:"onblur"
func (*HTMLDocument) Onblur() (onblur func(window.Event)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop Fires when the object loses the input focus.
//     * @param ev The focus event.
// js:"onblur"
func (*HTMLDocument) SetOnblur(onblur func(window.Event)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop Occurs when playback is possible, but would require further buffering.
//     * @param ev The event.
// js:"oncanplay"
func (*HTMLDocument) Oncanplay() (oncanplay func(window.Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop Occurs when playback is possible, but would require further buffering.
//     * @param ev The event.
// js:"oncanplay"
func (*HTMLDocument) SetOncanplay(oncanplay func(window.Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLDocument) Oncanplaythrough() (oncanplaythrough func(window.Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLDocument) SetOncanplaythrough(oncanplaythrough func(window.Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop Fires when the contents of the object or selection have changed.
//     * @param ev The event.
// js:"onchange"
func (*HTMLDocument) Onchange() (onchange func(window.Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop Fires when the contents of the object or selection have changed.
//     * @param ev The event.
// js:"onchange"
func (*HTMLDocument) SetOnchange(onchange func(window.Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop Fires when the user clicks the left mouse button on the object
//     * @param ev The mouse event.
// js:"onclick"
func (*HTMLDocument) Onclick() (onclick func(window.Event)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop Fires when the user clicks the left mouse button on the object
//     * @param ev The mouse event.
// js:"onclick"
func (*HTMLDocument) SetOnclick(onclick func(window.Event)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

// Oncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
//     * @param ev The mouse event.
// js:"oncontextmenu"
func (*HTMLDocument) Oncontextmenu() (oncontextmenu func(window.Event)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
//     * @param ev The mouse event.
// js:"oncontextmenu"
func (*HTMLDocument) SetOncontextmenu(oncontextmenu func(window.Event)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Ondblclick prop Fires when the user double-clicks the object.
//     * @param ev The mouse event.
// js:"ondblclick"
func (*HTMLDocument) Ondblclick() (ondblclick func(window.Event)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop Fires when the user double-clicks the object.
//     * @param ev The mouse event.
// js:"ondblclick"
func (*HTMLDocument) SetOndblclick(ondblclick func(window.Event)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The UI Event
// js:"ondeactivate"
func (*HTMLDocument) Ondeactivate() (ondeactivate func(window.Event)) {
	macro.Rewrite("$_.ondeactivate")
	return ondeactivate
}

// SetOndeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The UI Event
// js:"ondeactivate"
func (*HTMLDocument) SetOndeactivate(ondeactivate func(window.Event)) {
	macro.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

// Ondrag prop Fires on the source object continuously during a drag operation.
//     * @param ev The event.
// js:"ondrag"
func (*HTMLDocument) Ondrag() (ondrag func(window.Event)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop Fires on the source object continuously during a drag operation.
//     * @param ev The event.
// js:"ondrag"
func (*HTMLDocument) SetOndrag(ondrag func(window.Event)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
//     * @param ev The event.
// js:"ondragend"
func (*HTMLDocument) Ondragend() (ondragend func(window.Event)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
//     * @param ev The event.
// js:"ondragend"
func (*HTMLDocument) SetOndragend(ondragend func(window.Event)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop Fires on the target element when the user drags the object to a valid drop target.
//     * @param ev The drag event.
// js:"ondragenter"
func (*HTMLDocument) Ondragenter() (ondragenter func(window.Event)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop Fires on the target element when the user drags the object to a valid drop target.
//     * @param ev The drag event.
// js:"ondragenter"
func (*HTMLDocument) SetOndragenter(ondragenter func(window.Event)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
//     * @param ev The drag event.
// js:"ondragleave"
func (*HTMLDocument) Ondragleave() (ondragleave func(window.Event)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
//     * @param ev The drag event.
// js:"ondragleave"
func (*HTMLDocument) SetOndragleave(ondragleave func(window.Event)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
//     * @param ev The event.
// js:"ondragover"
func (*HTMLDocument) Ondragover() (ondragover func(window.Event)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
//     * @param ev The event.
// js:"ondragover"
func (*HTMLDocument) SetOndragover(ondragover func(window.Event)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
//     * @param ev The event.
// js:"ondragstart"
func (*HTMLDocument) Ondragstart() (ondragstart func(window.Event)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
//     * @param ev The event.
// js:"ondragstart"
func (*HTMLDocument) SetOndragstart(ondragstart func(window.Event)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*HTMLDocument) Ondrop() (ondrop func(window.Event)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*HTMLDocument) SetOndrop(ondrop func(window.Event)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop Occurs when the duration attribute is updated.
//     * @param ev The event.
// js:"ondurationchange"
func (*HTMLDocument) Ondurationchange() (ondurationchange func(window.Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop Occurs when the duration attribute is updated.
//     * @param ev The event.
// js:"ondurationchange"
func (*HTMLDocument) SetOndurationchange(ondurationchange func(window.Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop Occurs when the media element is reset to its initial state.
//     * @param ev The event.
// js:"onemptied"
func (*HTMLDocument) Onemptied() (onemptied func(window.Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop Occurs when the media element is reset to its initial state.
//     * @param ev The event.
// js:"onemptied"
func (*HTMLDocument) SetOnemptied(onemptied func(window.Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop Occurs when the end of playback is reached.
//     * @param ev The event
// js:"onended"
func (*HTMLDocument) Onended() (onended func(window.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

// SetOnended prop Occurs when the end of playback is reached.
//     * @param ev The event
// js:"onended"
func (*HTMLDocument) SetOnended(onended func(window.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

// Onerror prop Fires when an error occurs during object loading.
//     * @param ev The event.
// js:"onerror"
func (*HTMLDocument) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop Fires when an error occurs during object loading.
//     * @param ev The event.
// js:"onerror"
func (*HTMLDocument) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop Fires when the object receives focus.
//     * @param ev The event.
// js:"onfocus"
func (*HTMLDocument) Onfocus() (onfocus func(window.Event)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop Fires when the object receives focus.
//     * @param ev The event.
// js:"onfocus"
func (*HTMLDocument) SetOnfocus(onfocus func(window.Event)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

// Onfullscreenchange prop
// js:"onfullscreenchange"
func (*HTMLDocument) Onfullscreenchange() (onfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onfullscreenchange")
	return onfullscreenchange
}

// SetOnfullscreenchange prop
// js:"onfullscreenchange"
func (*HTMLDocument) SetOnfullscreenchange(onfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onfullscreenchange = $1", onfullscreenchange)
}

// Onfullscreenerror prop
// js:"onfullscreenerror"
func (*HTMLDocument) Onfullscreenerror() (onfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onfullscreenerror")
	return onfullscreenerror
}

// SetOnfullscreenerror prop
// js:"onfullscreenerror"
func (*HTMLDocument) SetOnfullscreenerror(onfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onfullscreenerror = $1", onfullscreenerror)
}

// Oninput prop
// js:"oninput"
func (*HTMLDocument) Oninput() (oninput func(window.Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*HTMLDocument) SetOninput(oninput func(window.Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*HTMLDocument) Oninvalid() (oninvalid func(window.Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*HTMLDocument) SetOninvalid(oninvalid func(window.Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop Fires when the user presses a key.
//     * @param ev The keyboard event
// js:"onkeydown"
func (*HTMLDocument) Onkeydown() (onkeydown func(window.Event)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop Fires when the user presses a key.
//     * @param ev The keyboard event
// js:"onkeydown"
func (*HTMLDocument) SetOnkeydown(onkeydown func(window.Event)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop Fires when the user presses an alphanumeric key.
//     * @param ev The event.
// js:"onkeypress"
func (*HTMLDocument) Onkeypress() (onkeypress func(window.Event)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop Fires when the user presses an alphanumeric key.
//     * @param ev The event.
// js:"onkeypress"
func (*HTMLDocument) SetOnkeypress(onkeypress func(window.Event)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop Fires when the user releases a key.
//     * @param ev The keyboard event
// js:"onkeyup"
func (*HTMLDocument) Onkeyup() (onkeyup func(window.Event)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop Fires when the user releases a key.
//     * @param ev The keyboard event
// js:"onkeyup"
func (*HTMLDocument) SetOnkeyup(onkeyup func(window.Event)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop Fires immediately after the browser loads the object.
//     * @param ev The event.
// js:"onload"
func (*HTMLDocument) Onload() (onload func(window.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

// SetOnload prop Fires immediately after the browser loads the object.
//     * @param ev The event.
// js:"onload"
func (*HTMLDocument) SetOnload(onload func(window.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop Occurs when media data is loaded at the current playback position.
//     * @param ev The event.
// js:"onloadeddata"
func (*HTMLDocument) Onloadeddata() (onloadeddata func(window.Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop Occurs when media data is loaded at the current playback position.
//     * @param ev The event.
// js:"onloadeddata"
func (*HTMLDocument) SetOnloadeddata(onloadeddata func(window.Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
//     * @param ev The event.
// js:"onloadedmetadata"
func (*HTMLDocument) Onloadedmetadata() (onloadedmetadata func(window.Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
//     * @param ev The event.
// js:"onloadedmetadata"
func (*HTMLDocument) SetOnloadedmetadata(onloadedmetadata func(window.Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop Occurs when Internet Explorer begins looking for media data.
//     * @param ev The event.
// js:"onloadstart"
func (*HTMLDocument) Onloadstart() (onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop Occurs when Internet Explorer begins looking for media data.
//     * @param ev The event.
// js:"onloadstart"
func (*HTMLDocument) SetOnloadstart(onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmousedown prop Fires when the user clicks the object with either mouse button.
//     * @param ev The mouse event.
// js:"onmousedown"
func (*HTMLDocument) Onmousedown() (onmousedown func(window.Event)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop Fires when the user clicks the object with either mouse button.
//     * @param ev The mouse event.
// js:"onmousedown"
func (*HTMLDocument) SetOnmousedown(onmousedown func(window.Event)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmousemove prop Fires when the user moves the mouse over the object.
//     * @param ev The mouse event.
// js:"onmousemove"
func (*HTMLDocument) Onmousemove() (onmousemove func(window.Event)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop Fires when the user moves the mouse over the object.
//     * @param ev The mouse event.
// js:"onmousemove"
func (*HTMLDocument) SetOnmousemove(onmousemove func(window.Event)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
//     * @param ev The mouse event.
// js:"onmouseout"
func (*HTMLDocument) Onmouseout() (onmouseout func(window.Event)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
//     * @param ev The mouse event.
// js:"onmouseout"
func (*HTMLDocument) SetOnmouseout(onmouseout func(window.Event)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop Fires when the user moves the mouse pointer into the object.
//     * @param ev The mouse event.
// js:"onmouseover"
func (*HTMLDocument) Onmouseover() (onmouseover func(window.Event)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop Fires when the user moves the mouse pointer into the object.
//     * @param ev The mouse event.
// js:"onmouseover"
func (*HTMLDocument) SetOnmouseover(onmouseover func(window.Event)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
//     * @param ev The mouse event.
// js:"onmouseup"
func (*HTMLDocument) Onmouseup() (onmouseup func(window.Event)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
//     * @param ev The mouse event.
// js:"onmouseup"
func (*HTMLDocument) SetOnmouseup(onmouseup func(window.Event)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop Fires when the wheel button is rotated.
//     * @param ev The mouse event
// js:"onmousewheel"
func (*HTMLDocument) Onmousewheel() (onmousewheel func(window.Event)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop Fires when the wheel button is rotated.
//     * @param ev The mouse event
// js:"onmousewheel"
func (*HTMLDocument) SetOnmousewheel(onmousewheel func(window.Event)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLDocument) Onmscontentzoom() (onmscontentzoom func(window.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

// SetOnmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLDocument) SetOnmscontentzoom(onmscontentzoom func(window.UIEvent)) {
	macro.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLDocument) Onmsgesturechange() (onmsgesturechange func(window.Event)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLDocument) SetOnmsgesturechange(onmsgesturechange func(window.Event)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLDocument) Onmsgesturedoubletap() (onmsgesturedoubletap func(window.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLDocument) SetOnmsgesturedoubletap(onmsgesturedoubletap func(window.Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*HTMLDocument) Onmsgestureend() (onmsgestureend func(window.Event)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*HTMLDocument) SetOnmsgestureend(onmsgestureend func(window.Event)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLDocument) Onmsgesturehold() (onmsgesturehold func(window.Event)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLDocument) SetOnmsgesturehold(onmsgesturehold func(window.Event)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLDocument) Onmsgesturestart() (onmsgesturestart func(window.Event)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLDocument) SetOnmsgesturestart(onmsgesturestart func(window.Event)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLDocument) Onmsgesturetap() (onmsgesturetap func(window.Event)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLDocument) SetOnmsgesturetap(onmsgesturetap func(window.Event)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLDocument) Onmsinertiastart() (onmsinertiastart func(window.Event)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLDocument) SetOnmsinertiastart(onmsinertiastart func(window.Event)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLDocument) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// SetOnmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLDocument) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	macro.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*HTMLDocument) Onmspointercancel() (onmspointercancel func(window.Event)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*HTMLDocument) SetOnmspointercancel(onmspointercancel func(window.Event)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*HTMLDocument) Onmspointerdown() (onmspointerdown func(window.Event)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*HTMLDocument) SetOnmspointerdown(onmspointerdown func(window.Event)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*HTMLDocument) Onmspointerenter() (onmspointerenter func(window.Event)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*HTMLDocument) SetOnmspointerenter(onmspointerenter func(window.Event)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*HTMLDocument) Onmspointerleave() (onmspointerleave func(window.Event)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*HTMLDocument) SetOnmspointerleave(onmspointerleave func(window.Event)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*HTMLDocument) Onmspointermove() (onmspointermove func(window.Event)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*HTMLDocument) SetOnmspointermove(onmspointermove func(window.Event)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*HTMLDocument) Onmspointerout() (onmspointerout func(window.Event)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*HTMLDocument) SetOnmspointerout(onmspointerout func(window.Event)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*HTMLDocument) Onmspointerover() (onmspointerover func(window.Event)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*HTMLDocument) SetOnmspointerover(onmspointerover func(window.Event)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*HTMLDocument) Onmspointerup() (onmspointerup func(window.Event)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*HTMLDocument) SetOnmspointerup(onmspointerup func(window.Event)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Onmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
//     * @param ev The event.
// js:"onmssitemodejumplistitemremoved"
func (*HTMLDocument) Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*window.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmssitemodejumplistitemremoved")
	return onmssitemodejumplistitemremoved
}

// SetOnmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
//     * @param ev The event.
// js:"onmssitemodejumplistitemremoved"
func (*HTMLDocument) SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*window.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmssitemodejumplistitemremoved = $1", onmssitemodejumplistitemremoved)
}

// Onmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
//     * @param ev The event.
// js:"onmsthumbnailclick"
func (*HTMLDocument) Onmsthumbnailclick() (onmsthumbnailclick func(*window.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmsthumbnailclick")
	return onmsthumbnailclick
}

// SetOnmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
//     * @param ev The event.
// js:"onmsthumbnailclick"
func (*HTMLDocument) SetOnmsthumbnailclick(onmsthumbnailclick func(*window.MSSiteModeEvent)) {
	macro.Rewrite("$_.onmsthumbnailclick = $1", onmsthumbnailclick)
}

// Onpause prop Occurs when playback is paused.
//     * @param ev The event.
// js:"onpause"
func (*HTMLDocument) Onpause() (onpause func(window.Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop Occurs when playback is paused.
//     * @param ev The event.
// js:"onpause"
func (*HTMLDocument) SetOnpause(onpause func(window.Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop Occurs when the play method is requested.
//     * @param ev The event.
// js:"onplay"
func (*HTMLDocument) Onplay() (onplay func(window.Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop Occurs when the play method is requested.
//     * @param ev The event.
// js:"onplay"
func (*HTMLDocument) SetOnplay(onplay func(window.Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop Occurs when the audio or video has started playing.
//     * @param ev The event.
// js:"onplaying"
func (*HTMLDocument) Onplaying() (onplaying func(window.Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop Occurs when the audio or video has started playing.
//     * @param ev The event.
// js:"onplaying"
func (*HTMLDocument) SetOnplaying(onplaying func(window.Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

// Onpointerlockchange prop
// js:"onpointerlockchange"
func (*HTMLDocument) Onpointerlockchange() (onpointerlockchange func(window.Event)) {
	macro.Rewrite("$_.onpointerlockchange")
	return onpointerlockchange
}

// SetOnpointerlockchange prop
// js:"onpointerlockchange"
func (*HTMLDocument) SetOnpointerlockchange(onpointerlockchange func(window.Event)) {
	macro.Rewrite("$_.onpointerlockchange = $1", onpointerlockchange)
}

// Onpointerlockerror prop
// js:"onpointerlockerror"
func (*HTMLDocument) Onpointerlockerror() (onpointerlockerror func(window.Event)) {
	macro.Rewrite("$_.onpointerlockerror")
	return onpointerlockerror
}

// SetOnpointerlockerror prop
// js:"onpointerlockerror"
func (*HTMLDocument) SetOnpointerlockerror(onpointerlockerror func(window.Event)) {
	macro.Rewrite("$_.onpointerlockerror = $1", onpointerlockerror)
}

// Onprogress prop Occurs to indicate progress while downloading media data.
//     * @param ev The event.
// js:"onprogress"
func (*HTMLDocument) Onprogress() (onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop Occurs to indicate progress while downloading media data.
//     * @param ev The event.
// js:"onprogress"
func (*HTMLDocument) SetOnprogress(onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop Occurs when the playback rate is increased or decreased.
//     * @param ev The event.
// js:"onratechange"
func (*HTMLDocument) Onratechange() (onratechange func(window.Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop Occurs when the playback rate is increased or decreased.
//     * @param ev The event.
// js:"onratechange"
func (*HTMLDocument) SetOnratechange(onratechange func(window.Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreadystatechange prop Fires when the state of the object has changed.
//     * @param ev The event
// js:"onreadystatechange"
func (*HTMLDocument) Onreadystatechange() (onreadystatechange func(window.Event)) {
	macro.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

// SetOnreadystatechange prop Fires when the state of the object has changed.
//     * @param ev The event
// js:"onreadystatechange"
func (*HTMLDocument) SetOnreadystatechange(onreadystatechange func(window.Event)) {
	macro.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

// Onreset prop Fires when the user resets a form.
//     * @param ev The event.
// js:"onreset"
func (*HTMLDocument) Onreset() (onreset func(window.Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop Fires when the user resets a form.
//     * @param ev The event.
// js:"onreset"
func (*HTMLDocument) SetOnreset(onreset func(window.Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

// Onscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
//     * @param ev The event.
// js:"onscroll"
func (*HTMLDocument) Onscroll() (onscroll func(window.Event)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
//     * @param ev The event.
// js:"onscroll"
func (*HTMLDocument) SetOnscroll(onscroll func(window.Event)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop Occurs when the seek operation ends.
//     * @param ev The event.
// js:"onseeked"
func (*HTMLDocument) Onseeked() (onseeked func(window.Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop Occurs when the seek operation ends.
//     * @param ev The event.
// js:"onseeked"
func (*HTMLDocument) SetOnseeked(onseeked func(window.Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop Occurs when the current playback position is moved.
//     * @param ev The event.
// js:"onseeking"
func (*HTMLDocument) Onseeking() (onseeking func(window.Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop Occurs when the current playback position is moved.
//     * @param ev The event.
// js:"onseeking"
func (*HTMLDocument) SetOnseeking(onseeking func(window.Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop Fires when the current selection changes.
//     * @param ev The event.
// js:"onselect"
func (*HTMLDocument) Onselect() (onselect func(window.Event)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop Fires when the current selection changes.
//     * @param ev The event.
// js:"onselect"
func (*HTMLDocument) SetOnselect(onselect func(window.Event)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

// Onselectionchange prop Fires when the selection state of a document changes.
//     * @param ev The event.
// js:"onselectionchange"
func (*HTMLDocument) Onselectionchange() (onselectionchange func(window.Event)) {
	macro.Rewrite("$_.onselectionchange")
	return onselectionchange
}

// SetOnselectionchange prop Fires when the selection state of a document changes.
//     * @param ev The event.
// js:"onselectionchange"
func (*HTMLDocument) SetOnselectionchange(onselectionchange func(window.Event)) {
	macro.Rewrite("$_.onselectionchange = $1", onselectionchange)
}

// Onselectstart prop
// js:"onselectstart"
func (*HTMLDocument) Onselectstart() (onselectstart func(window.Event)) {
	macro.Rewrite("$_.onselectstart")
	return onselectstart
}

// SetOnselectstart prop
// js:"onselectstart"
func (*HTMLDocument) SetOnselectstart(onselectstart func(window.Event)) {
	macro.Rewrite("$_.onselectstart = $1", onselectstart)
}

// Onstalled prop Occurs when the download has stopped.
//     * @param ev The event.
// js:"onstalled"
func (*HTMLDocument) Onstalled() (onstalled func(window.Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop Occurs when the download has stopped.
//     * @param ev The event.
// js:"onstalled"
func (*HTMLDocument) SetOnstalled(onstalled func(window.Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

// Onstop prop Fires when the user clicks the Stop button or leaves the Web page.
//     * @param ev The event.
// js:"onstop"
func (*HTMLDocument) Onstop() (onstop func(window.Event)) {
	macro.Rewrite("$_.onstop")
	return onstop
}

// SetOnstop prop Fires when the user clicks the Stop button or leaves the Web page.
//     * @param ev The event.
// js:"onstop"
func (*HTMLDocument) SetOnstop(onstop func(window.Event)) {
	macro.Rewrite("$_.onstop = $1", onstop)
}

// Onsubmit prop
// js:"onsubmit"
func (*HTMLDocument) Onsubmit() (onsubmit func(window.Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*HTMLDocument) SetOnsubmit(onsubmit func(window.Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop Occurs if the load operation has been intentionally halted.
//     * @param ev The event.
// js:"onsuspend"
func (*HTMLDocument) Onsuspend() (onsuspend func(window.Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop Occurs if the load operation has been intentionally halted.
//     * @param ev The event.
// js:"onsuspend"
func (*HTMLDocument) SetOnsuspend(onsuspend func(window.Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop Occurs to indicate the current playback position.
//     * @param ev The event.
// js:"ontimeupdate"
func (*HTMLDocument) Ontimeupdate() (ontimeupdate func(window.Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop Occurs to indicate the current playback position.
//     * @param ev The event.
// js:"ontimeupdate"
func (*HTMLDocument) SetOntimeupdate(ontimeupdate func(window.Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*HTMLDocument) Ontouchcancel() (ontouchcancel func(window.Event)) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*HTMLDocument) SetOntouchcancel(ontouchcancel func(window.Event)) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*HTMLDocument) Ontouchend() (ontouchend func(window.Event)) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*HTMLDocument) SetOntouchend(ontouchend func(window.Event)) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*HTMLDocument) Ontouchmove() (ontouchmove func(window.Event)) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*HTMLDocument) SetOntouchmove(ontouchmove func(window.Event)) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*HTMLDocument) Ontouchstart() (ontouchstart func(window.Event)) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*HTMLDocument) SetOntouchstart(ontouchstart func(window.Event)) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
//     * @param ev The event.
// js:"onvolumechange"
func (*HTMLDocument) Onvolumechange() (onvolumechange func(window.Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
//     * @param ev The event.
// js:"onvolumechange"
func (*HTMLDocument) SetOnvolumechange(onvolumechange func(window.Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
//     * @param ev The event.
// js:"onwaiting"
func (*HTMLDocument) Onwaiting() (onwaiting func(window.Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
//     * @param ev The event.
// js:"onwaiting"
func (*HTMLDocument) SetOnwaiting(onwaiting func(window.Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLDocument) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLDocument) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLDocument) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLDocument) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	macro.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// Plugins prop
// js:"plugins"
func (*HTMLDocument) Plugins() (plugins window.HTMLCollection) {
	macro.Rewrite("$_.plugins")
	return plugins
}

// PointerLockElement prop
// js:"pointerLockElement"
func (*HTMLDocument) PointerLockElement() (pointerLockElement window.Element) {
	macro.Rewrite("$_.pointerLockElement")
	return pointerLockElement
}

// ReadyState prop Retrieves a value that indicates the current state of the object.
// js:"readyState"
func (*HTMLDocument) ReadyState() (readyState string) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Referrer prop Gets the URL of the location that referred the user to the current page.
// js:"referrer"
func (*HTMLDocument) Referrer() (referrer string) {
	macro.Rewrite("$_.referrer")
	return referrer
}

// RootElement prop Gets the root svg element in the document hierarchy.
// js:"rootElement"
func (*HTMLDocument) RootElement() (rootElement *window.SVGSVGElement) {
	macro.Rewrite("$_.rootElement")
	return rootElement
}

// Scripts prop Retrieves a collection of all script objects in the document.
// js:"scripts"
func (*HTMLDocument) Scripts() (scripts window.HTMLCollection) {
	macro.Rewrite("$_.scripts")
	return scripts
}

// ScrollingElement prop
// js:"scrollingElement"
func (*HTMLDocument) ScrollingElement() (scrollingElement window.Element) {
	macro.Rewrite("$_.scrollingElement")
	return scrollingElement
}

// StyleSheets prop Retrieves a collection of styleSheet objects representing the style sheets that correspond to each instance of a link or style object in the document.
// js:"styleSheets"
func (*HTMLDocument) StyleSheets() (styleSheets *window.StyleSheetList) {
	macro.Rewrite("$_.styleSheets")
	return styleSheets
}

// Title prop Contains the title of the document.
// js:"title"
func (*HTMLDocument) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}

// SetTitle prop Contains the title of the document.
// js:"title"
func (*HTMLDocument) SetTitle(title string) {
	macro.Rewrite("$_.title = $1", title)
}

// URL prop Sets or gets the URL for the current document.
// js:"URL"
func (*HTMLDocument) URL() (URL string) {
	macro.Rewrite("$_.URL")
	return URL
}

// URLUnencoded prop Gets the URL for the document, stripped of any character encoding.
// js:"URLUnencoded"
func (*HTMLDocument) URLUnencoded() (URLUnencoded string) {
	macro.Rewrite("$_.URLUnencoded")
	return URLUnencoded
}

// VisibilityState prop
// js:"visibilityState"
func (*HTMLDocument) VisibilityState() (visibilityState *visibilitystate.VisibilityState) {
	macro.Rewrite("$_.visibilityState")
	return visibilityState
}

// VlinkColor prop Sets or gets the color of the links that the user has visited.
// js:"vlinkColor"
func (*HTMLDocument) VlinkColor() (vlinkColor string) {
	macro.Rewrite("$_.vlinkColor")
	return vlinkColor
}

// SetVlinkColor prop Sets or gets the color of the links that the user has visited.
// js:"vlinkColor"
func (*HTMLDocument) SetVlinkColor(vlinkColor string) {
	macro.Rewrite("$_.vlinkColor = $1", vlinkColor)
}

// WebkitCurrentFullScreenElement prop
// js:"webkitCurrentFullScreenElement"
func (*HTMLDocument) WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement window.Element) {
	macro.Rewrite("$_.webkitCurrentFullScreenElement")
	return webkitCurrentFullScreenElement
}

// WebkitFullscreenElement prop
// js:"webkitFullscreenElement"
func (*HTMLDocument) WebkitFullscreenElement() (webkitFullscreenElement window.Element) {
	macro.Rewrite("$_.webkitFullscreenElement")
	return webkitFullscreenElement
}

// WebkitFullscreenEnabled prop
// js:"webkitFullscreenEnabled"
func (*HTMLDocument) WebkitFullscreenEnabled() (webkitFullscreenEnabled bool) {
	macro.Rewrite("$_.webkitFullscreenEnabled")
	return webkitFullscreenEnabled
}

// WebkitIsFullScreen prop
// js:"webkitIsFullScreen"
func (*HTMLDocument) WebkitIsFullScreen() (webkitIsFullScreen bool) {
	macro.Rewrite("$_.webkitIsFullScreen")
	return webkitIsFullScreen
}

// XMLEncoding prop
// js:"xmlEncoding"
func (*HTMLDocument) XMLEncoding() (xmlEncoding string) {
	macro.Rewrite("$_.xmlEncoding")
	return xmlEncoding
}

// XMLStandalone prop
// js:"xmlStandalone"
func (*HTMLDocument) XMLStandalone() (xmlStandalone bool) {
	macro.Rewrite("$_.xmlStandalone")
	return xmlStandalone
}

// SetXMLStandalone prop
// js:"xmlStandalone"
func (*HTMLDocument) SetXMLStandalone(xmlStandalone bool) {
	macro.Rewrite("$_.xmlStandalone = $1", xmlStandalone)
}

// XMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
// js:"xmlVersion"
func (*HTMLDocument) XMLVersion() (xmlVersion string) {
	macro.Rewrite("$_.xmlVersion")
	return xmlVersion
}

// SetXMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
// js:"xmlVersion"
func (*HTMLDocument) SetXMLVersion(xmlVersion string) {
	macro.Rewrite("$_.xmlVersion = $1", xmlVersion)
}

// Onpointercancel prop
// js:"onpointercancel"
func (*HTMLDocument) Onpointercancel() (onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*HTMLDocument) SetOnpointercancel(onpointercancel func(window.Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*HTMLDocument) Onpointerdown() (onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*HTMLDocument) SetOnpointerdown(onpointerdown func(window.Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*HTMLDocument) Onpointerenter() (onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*HTMLDocument) SetOnpointerenter(onpointerenter func(window.Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*HTMLDocument) Onpointerleave() (onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*HTMLDocument) SetOnpointerleave(onpointerleave func(window.Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*HTMLDocument) Onpointermove() (onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*HTMLDocument) SetOnpointermove(onpointermove func(window.Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*HTMLDocument) Onpointerout() (onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*HTMLDocument) SetOnpointerout(onpointerout func(window.Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*HTMLDocument) Onpointerover() (onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*HTMLDocument) SetOnpointerover(onpointerover func(window.Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*HTMLDocument) Onpointerup() (onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*HTMLDocument) SetOnpointerup(onpointerup func(window.Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*HTMLDocument) Onwheel() (onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*HTMLDocument) SetOnwheel(onwheel func(window.Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// Attributes prop
// js:"attributes"
func (*HTMLDocument) Attributes() (attributes *window.NamedNodeMap) {
	macro.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*HTMLDocument) BaseURI() (baseURI string) {
	macro.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLDocument) ChildNodes() (childNodes *window.NodeList) {
	macro.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLDocument) FirstChild() (firstChild window.Node) {
	macro.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLDocument) LastChild() (lastChild window.Node) {
	macro.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*HTMLDocument) LocalName() (localName string) {
	macro.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*HTMLDocument) NamespaceURI() (namespaceURI string) {
	macro.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLDocument) NextSibling() (nextSibling window.Node) {
	macro.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLDocument) NodeName() (nodeName string) {
	macro.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLDocument) NodeType() (nodeType uint8) {
	macro.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLDocument) NodeValue() (nodeValue string) {
	macro.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLDocument) SetNodeValue(nodeValue string) {
	macro.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*HTMLDocument) OwnerDocument() (ownerDocument window.Document) {
	macro.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*HTMLDocument) ParentElement() (parentElement window.HTMLElement) {
	macro.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*HTMLDocument) ParentNode() (parentNode window.Node) {
	macro.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*HTMLDocument) PreviousSibling() (previousSibling window.Node) {
	macro.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*HTMLDocument) TextContent() (textContent string) {
	macro.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*HTMLDocument) SetTextContent(textContent string) {
	macro.Rewrite("$_.textContent = $1", textContent)
}
