package window

import (
	"github.com/matthewmueller/golly/dom/htmlallcollection"
	"github.com/matthewmueller/golly/dom/location"
	"github.com/matthewmueller/golly/dom/visibilitystate"
	"github.com/matthewmueller/golly/dom/xpathnsresolver"
)

// Document interface
// js:"Document"
type Document interface {
	Node

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll"
	QuerySelectorAll(selectors string) (n *NodeList)

	// CreateEvent
	// js:"createEvent"
	CreateEvent(eventInterface string) (e Event)

	// AdoptNode
	// js:"adoptNode"
	AdoptNode(source Node) (n Node)

	// CaptureEvents
	// js:"captureEvents"
	CaptureEvents()

	// CaretRangeFromPoint
	// js:"caretRangeFromPoint"
	CaretRangeFromPoint(x float32, y float32) (r *Range)

	// Clear
	// js:"clear"
	Clear()

	// Close Closes an output stream and forces the sent data to display.
	// js:"close"
	Close()

	// CreateAttribute Creates an attribute object with a specified name.
	//     * @param name String that sets the attribute object's name.
	// js:"createAttribute"
	CreateAttribute(name string) (a *Attr)

	// CreateAttributeNS
	// js:"createAttributeNS"
	CreateAttributeNS(namespaceURI string, qualifiedName string) (a *Attr)

	// CreateCDATASection
	// js:"createCDATASection"
	CreateCDATASection(data string) (c *CDATASection)

	// CreateComment Creates a comment object with the specified data.
	//     * @param data Sets the comment object's data.
	// js:"createComment"
	CreateComment(data string) (c *Comment)

	// CreateDocumentFragment Creates a new document.
	// js:"createDocumentFragment"
	CreateDocumentFragment() (d *DocumentFragment)

	// CreateElement Creates an instance of the element for the specified tag.
	//     * @param tagName The name of an element.
	// js:"createElement"
	CreateElement(tagName string) (e Element)

	// CreateElementNS
	// js:"createElementNS"
	CreateElementNS(namespaceURI string, qualifiedName string) (e Element)

	// CreateExpression
	// js:"createExpression"
	CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (x *XPathExpression)

	// CreateNodeIterator Creates a NodeIterator object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list
	//     * @param filter A custom NodeFilter function to use. For more information, see filter. Use null for no filter.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	// js:"createNodeIterator"
	CreateNodeIterator(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (n *NodeIterator)

	// CreateNSResolver
	// js:"createNSResolver"
	CreateNSResolver(nodeResolver Node) (x *xpathnsresolver.XPathNSResolver)

	// CreateProcessingInstruction
	// js:"createProcessingInstruction"
	CreateProcessingInstruction(target string, data string) (p *ProcessingInstruction)

	// CreateRange Returns an empty range object that has both of its boundary points positioned at the beginning of the document.
	// js:"createRange"
	CreateRange() (r *Range)

	// CreateTextNode Creates a text string from the specified value.
	//     * @param data String that specifies the nodeValue property of the text node.
	// js:"createTextNode"
	CreateTextNode(data string) (t Text)

	// CreateTouch
	// js:"createTouch"
	CreateTouch(view *Window, target EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *Touch)

	// CreateTouchList
	// js:"createTouchList"
	CreateTouchList(touches *Touch) (t *TouchList)

	// CreateTreeWalker Creates a TreeWalker object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list. For more information, see whatToShow.
	//     * @param filter A custom NodeFilter function to use.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	// js:"createTreeWalker"
	CreateTreeWalker(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (t *TreeWalker)

	// ElementFromPoint Returns the element for the specified x coordinate and the specified y coordinate.
	//     * @param x The x-offset
	//     * @param y The y-offset
	// js:"elementFromPoint"
	ElementFromPoint(x int, y int) (e Element)

	// Evaluate
	// js:"evaluate"
	Evaluate(expression string, contextNode Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *XPathResult) (x *XPathResult)

	// ExecCommand Executes a command on the current document, current selection, or the given range.
	//     * @param commandId String that specifies the command to execute. This command can be any of the command identifiers that can be executed in script.
	//     * @param showUI Display the user interface, defaults to false.
	//     * @param value Value to assign.
	// js:"execCommand"
	ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool)

	// ExecCommandShowHelp Displays help information for the given command identifier.
	//     * @param commandId Displays help information for the given command identifier.
	// js:"execCommandShowHelp"
	ExecCommandShowHelp(commandId string) (b bool)

	// ExitFullscreen
	// js:"exitFullscreen"
	ExitFullscreen()

	// ExitPointerLock
	// js:"exitPointerLock"
	ExitPointerLock()

	// Focus Causes the element to receive the focus and executes the code specified by the onfocus event.
	// js:"focus"
	Focus()

	// GetElementByID Returns a reference to the first object with the specified value of the ID or NAME attribute.
	//     * @param elementId String that specifies the ID value. Case-insensitive.
	// js:"getElementById"
	GetElementByID(elementId string) (e Element)

	// GetElementsByClassName
	// js:"getElementsByClassName"
	GetElementsByClassName(classNames string) (n *NodeList)

	// GetElementsByName Gets a collection of objects based on the value of the NAME or ID attribute.
	//     * @param elementName Gets a collection of objects based on the value of the NAME or ID attribute.
	// js:"getElementsByName"
	GetElementsByName(elementName string) (n *NodeList)

	// GetElementsByTagName Retrieves a collection of objects based on the specified element name.
	//     * @param name Specifies the name of an element.
	// js:"getElementsByTagName"
	GetElementsByTagName(tagname string) (n *NodeList)

	// GetElementsByTagNameNS
	// js:"getElementsByTagNameNS"
	GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// GetSelection Returns an object representing the current selection of the document that is loaded into the object displaying a webpage.
	// js:"getSelection"
	GetSelection() (s *Selection)

	// HasFocus Gets a value indicating whether the object currently has focus.
	// js:"hasFocus"
	HasFocus() (b bool)

	// ImportNode
	// js:"importNode"
	ImportNode(importedNode Node, deep bool) (n Node)

	// MsElementsFromPoint
	// js:"msElementsFromPoint"
	MsElementsFromPoint(x float32, y float32) (n *NodeList)

	// MsElementsFromRect
	// js:"msElementsFromRect"
	MsElementsFromRect(left float32, top float32, width float32, height float32) (n *NodeList)

	// Open Opens a new window and loads a document specified by a given URL. Also, opens a new window that uses the url parameter and the name parameter to collect the output of the write method and the writeln method.
	//     * @param url Specifies a MIME type for the document.
	//     * @param name Specifies the name of the window. This name is used as the value for the TARGET attribute on a form or an anchor element.
	//     * @param features Contains a list of items separated by commas. Each item consists of an option and a value, separated by an equals sign (for example, "fullscreen=yes, toolbar=yes"). The following values are supported.
	//     * @param replace Specifies whether the existing entry for the document is replaced in the history list.
	// js:"open"
	Open(url *string, name *string, features *string, replace *bool) (i interface{})

	// QueryCommandEnabled Returns a Boolean value that indicates whether a specified command can be successfully executed using execCommand, given the current state of the document.
	//     * @param commandId Specifies a command identifier.
	// js:"queryCommandEnabled"
	QueryCommandEnabled(commandId string) (b bool)

	// QueryCommandIndeterm Returns a Boolean value that indicates whether the specified command is in the indeterminate state.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandIndeterm"
	QueryCommandIndeterm(commandId string) (b bool)

	// QueryCommandState Returns a Boolean value that indicates the current state of the command.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandState"
	QueryCommandState(commandId string) (b bool)

	// QueryCommandSupported Returns a Boolean value that indicates whether the current command is supported on the current range.
	//     * @param commandId Specifies a command identifier.
	// js:"queryCommandSupported"
	QueryCommandSupported(commandId string) (b bool)

	// QueryCommandText Retrieves the string associated with a command.
	//     * @param commandId String that contains the identifier of a command. This can be any command identifier given in the list of Command Identifiers.
	// js:"queryCommandText"
	QueryCommandText(commandId string) (s string)

	// QueryCommandValue Returns the current value of the document, range, or current selection for the given command.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandValue"
	QueryCommandValue(commandId string) (s string)

	// ReleaseEvents
	// js:"releaseEvents"
	ReleaseEvents()

	// UpdateSettings Allows updating the print settings for the page.
	// js:"updateSettings"
	UpdateSettings()

	// WebkitCancelFullScreen
	// js:"webkitCancelFullScreen"
	WebkitCancelFullScreen()

	// WebkitExitFullscreen
	// js:"webkitExitFullscreen"
	WebkitExitFullscreen()

	// Write Writes one or more HTML expressions to a document in the specified window.
	//     * @param content Specifies the text and HTML tags to write.
	// js:"write"
	Write(content string)

	// Writeln Writes one or more HTML expressions, followed by a carriage return, to a document in the specified window.
	//     * @param content The text and HTML tags to write.
	// js:"writeln"
	Writeln(content string)

	// Onpointercancel prop
	// js:"onpointercancel"
	Onpointercancel() (onpointercancel func(Event))

	// SetOnpointercancel prop
	// js:"onpointercancel"
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown prop
	// js:"onpointerdown"
	Onpointerdown() (onpointerdown func(Event))

	// SetOnpointerdown prop
	// js:"onpointerdown"
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter prop
	// js:"onpointerenter"
	Onpointerenter() (onpointerenter func(Event))

	// SetOnpointerenter prop
	// js:"onpointerenter"
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave prop
	// js:"onpointerleave"
	Onpointerleave() (onpointerleave func(Event))

	// SetOnpointerleave prop
	// js:"onpointerleave"
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove prop
	// js:"onpointermove"
	Onpointermove() (onpointermove func(Event))

	// SetOnpointermove prop
	// js:"onpointermove"
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout prop
	// js:"onpointerout"
	Onpointerout() (onpointerout func(Event))

	// SetOnpointerout prop
	// js:"onpointerout"
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover prop
	// js:"onpointerover"
	Onpointerover() (onpointerover func(Event))

	// SetOnpointerover prop
	// js:"onpointerover"
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup prop
	// js:"onpointerup"
	Onpointerup() (onpointerup func(Event))

	// SetOnpointerup prop
	// js:"onpointerup"
	SetOnpointerup(onpointerup func(Event))

	// Onwheel prop
	// js:"onwheel"
	Onwheel() (onwheel func(Event))

	// SetOnwheel prop
	// js:"onwheel"
	SetOnwheel(onwheel func(Event))

	// ActiveElement prop Gets the object that has the focus when the parent document has focus.
	// js:"activeElement"
	ActiveElement() (activeElement Element)

	// AlinkColor prop Sets or gets the color of all active links in the document.
	// js:"alinkColor"
	AlinkColor() (alinkColor string)

	// SetAlinkColor prop Sets or gets the color of all active links in the document.
	// js:"alinkColor"
	SetAlinkColor(alinkColor string)

	// All prop Returns a reference to the collection of elements contained by the object.
	// js:"all"
	All() (all *htmlallcollection.HTMLAllCollection)

	// Anchors prop Retrieves a collection of all a objects that have a name and/or id property. Objects in this collection are in HTML source order.
	// js:"anchors"
	Anchors() (anchors HTMLCollection)

	// Applets prop Retrieves a collection of all applet objects in the document.
	// js:"applets"
	Applets() (applets HTMLCollection)

	// BgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	// js:"bgColor"
	BgColor() (bgColor string)

	// SetBgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	// js:"bgColor"
	SetBgColor(bgColor string)

	// Body prop Specifies the beginning and end of the document body.
	// js:"body"
	Body() (body HTMLElement)

	// SetBody prop Specifies the beginning and end of the document body.
	// js:"body"
	SetBody(body HTMLElement)

	// CharacterSet prop
	// js:"characterSet"
	CharacterSet() (characterSet string)

	// Charset prop Gets or sets the character set used to encode the object.
	// js:"charset"
	Charset() (charset string)

	// SetCharset prop Gets or sets the character set used to encode the object.
	// js:"charset"
	SetCharset(charset string)

	// CompatMode prop Gets a value that indicates whether standards-compliant mode is switched on for the object.
	// js:"compatMode"
	CompatMode() (compatMode string)

	// Cookie prop
	// js:"cookie"
	Cookie() (cookie string)

	// SetCookie prop
	// js:"cookie"
	SetCookie(cookie string)

	// CurrentScript prop
	// js:"currentScript"
	CurrentScript() (currentScript interface{})

	// DefaultView prop
	// js:"defaultView"
	DefaultView() (defaultView *Window)

	// DesignMode prop Sets or gets a value that indicates whether the document can be edited.
	// js:"designMode"
	DesignMode() (designMode string)

	// SetDesignMode prop Sets or gets a value that indicates whether the document can be edited.
	// js:"designMode"
	SetDesignMode(designMode string)

	// Dir prop Sets or retrieves a value that indicates the reading order of the object.
	// js:"dir"
	Dir() (dir string)

	// SetDir prop Sets or retrieves a value that indicates the reading order of the object.
	// js:"dir"
	SetDir(dir string)

	// Doctype prop Gets an object representing the document type declaration associated with the current document.
	// js:"doctype"
	Doctype() (doctype *DocumentType)

	// DocumentElement prop Gets a reference to the root node of the document.
	// js:"documentElement"
	DocumentElement() (documentElement Element)

	// Domain prop Sets or gets the security domain of the document.
	// js:"domain"
	Domain() (domain string)

	// SetDomain prop Sets or gets the security domain of the document.
	// js:"domain"
	SetDomain(domain string)

	// Embeds prop Retrieves a collection of all embed objects in the document.
	// js:"embeds"
	Embeds() (embeds HTMLCollection)

	// FgColor prop Sets or gets the foreground (text) color of the document.
	// js:"fgColor"
	FgColor() (fgColor string)

	// SetFgColor prop Sets or gets the foreground (text) color of the document.
	// js:"fgColor"
	SetFgColor(fgColor string)

	// Forms prop Retrieves a collection, in source order, of all form objects in the document.
	// js:"forms"
	Forms() (forms HTMLCollection)

	// FullscreenElement prop
	// js:"fullscreenElement"
	FullscreenElement() (fullscreenElement Element)

	// FullscreenEnabled prop
	// js:"fullscreenEnabled"
	FullscreenEnabled() (fullscreenEnabled bool)

	// Head prop
	// js:"head"
	Head() (head *HTMLHeadElement)

	// Hidden prop
	// js:"hidden"
	Hidden() (hidden bool)

	// Images prop Retrieves a collection, in source order, of img objects in the document.
	// js:"images"
	Images() (images HTMLCollection)

	// Implementation prop Gets the implementation object of the current document.
	// js:"implementation"
	Implementation() (implementation *DOMImplementation)

	// InputEncoding prop Returns the character encoding used to create the webpage that is loaded into the document object.
	// js:"inputEncoding"
	InputEncoding() (inputEncoding string)

	// LastModified prop Gets the date that the page was last modified, if the page supplies one.
	// js:"lastModified"
	LastModified() (lastModified string)

	// LinkColor prop Sets or gets the color of the document links.
	// js:"linkColor"
	LinkColor() (linkColor string)

	// SetLinkColor prop Sets or gets the color of the document links.
	// js:"linkColor"
	SetLinkColor(linkColor string)

	// Links prop Retrieves a collection of all a objects that specify the href property and all area objects in the document.
	// js:"links"
	Links() (links HTMLCollection)

	// Location prop Contains information about the current URL.
	// js:"location"
	Location() (location *location.Location)

	// MsCapsLockWarningOff prop
	// js:"msCapsLockWarningOff"
	MsCapsLockWarningOff() (msCapsLockWarningOff bool)

	// SetMsCapsLockWarningOff prop
	// js:"msCapsLockWarningOff"
	SetMsCapsLockWarningOff(msCapsLockWarningOff bool)

	// MsCSSOMElementFloatMetrics prop
	// js:"msCSSOMElementFloatMetrics"
	MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool)

	// SetMsCSSOMElementFloatMetrics prop
	// js:"msCSSOMElementFloatMetrics"
	SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool)

	// Onabort prop Fires when the user aborts the download.
	//     * @param ev The event.
	// js:"onabort"
	Onabort() (onabort func(Event))

	// SetOnabort prop Fires when the user aborts the download.
	//     * @param ev The event.
	// js:"onabort"
	SetOnabort(onabort func(Event))

	// Onactivate prop Fires when the object is set as the active element.
	//     * @param ev The event.
	// js:"onactivate"
	Onactivate() (onactivate func(Event))

	// SetOnactivate prop Fires when the object is set as the active element.
	//     * @param ev The event.
	// js:"onactivate"
	SetOnactivate(onactivate func(Event))

	// Onbeforeactivate prop Fires immediately before the object is set as the active element.
	//     * @param ev The event.
	// js:"onbeforeactivate"
	Onbeforeactivate() (onbeforeactivate func(Event))

	// SetOnbeforeactivate prop Fires immediately before the object is set as the active element.
	//     * @param ev The event.
	// js:"onbeforeactivate"
	SetOnbeforeactivate(onbeforeactivate func(Event))

	// Onbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The event.
	// js:"onbeforedeactivate"
	Onbeforedeactivate() (onbeforedeactivate func(Event))

	// SetOnbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The event.
	// js:"onbeforedeactivate"
	SetOnbeforedeactivate(onbeforedeactivate func(Event))

	// Onblur prop Fires when the object loses the input focus.
	//     * @param ev The focus event.
	// js:"onblur"
	Onblur() (onblur func(Event))

	// SetOnblur prop Fires when the object loses the input focus.
	//     * @param ev The focus event.
	// js:"onblur"
	SetOnblur(onblur func(Event))

	// Oncanplay prop Occurs when playback is possible, but would require further buffering.
	//     * @param ev The event.
	// js:"oncanplay"
	Oncanplay() (oncanplay func(Event))

	// SetOncanplay prop Occurs when playback is possible, but would require further buffering.
	//     * @param ev The event.
	// js:"oncanplay"
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough prop
	// js:"oncanplaythrough"
	Oncanplaythrough() (oncanplaythrough func(Event))

	// SetOncanplaythrough prop
	// js:"oncanplaythrough"
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange prop Fires when the contents of the object or selection have changed.
	//     * @param ev The event.
	// js:"onchange"
	Onchange() (onchange func(Event))

	// SetOnchange prop Fires when the contents of the object or selection have changed.
	//     * @param ev The event.
	// js:"onchange"
	SetOnchange(onchange func(Event))

	// Onclick prop Fires when the user clicks the left mouse button on the object
	//     * @param ev The mouse event.
	// js:"onclick"
	Onclick() (onclick func(Event))

	// SetOnclick prop Fires when the user clicks the left mouse button on the object
	//     * @param ev The mouse event.
	// js:"onclick"
	SetOnclick(onclick func(Event))

	// Oncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
	//     * @param ev The mouse event.
	// js:"oncontextmenu"
	Oncontextmenu() (oncontextmenu func(Event))

	// SetOncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
	//     * @param ev The mouse event.
	// js:"oncontextmenu"
	SetOncontextmenu(oncontextmenu func(Event))

	// Ondblclick prop Fires when the user double-clicks the object.
	//     * @param ev The mouse event.
	// js:"ondblclick"
	Ondblclick() (ondblclick func(Event))

	// SetOndblclick prop Fires when the user double-clicks the object.
	//     * @param ev The mouse event.
	// js:"ondblclick"
	SetOndblclick(ondblclick func(Event))

	// Ondeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The UI Event
	// js:"ondeactivate"
	Ondeactivate() (ondeactivate func(Event))

	// SetOndeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The UI Event
	// js:"ondeactivate"
	SetOndeactivate(ondeactivate func(Event))

	// Ondrag prop Fires on the source object continuously during a drag operation.
	//     * @param ev The event.
	// js:"ondrag"
	Ondrag() (ondrag func(Event))

	// SetOndrag prop Fires on the source object continuously during a drag operation.
	//     * @param ev The event.
	// js:"ondrag"
	SetOndrag(ondrag func(Event))

	// Ondragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
	//     * @param ev The event.
	// js:"ondragend"
	Ondragend() (ondragend func(Event))

	// SetOndragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
	//     * @param ev The event.
	// js:"ondragend"
	SetOndragend(ondragend func(Event))

	// Ondragenter prop Fires on the target element when the user drags the object to a valid drop target.
	//     * @param ev The drag event.
	// js:"ondragenter"
	Ondragenter() (ondragenter func(Event))

	// SetOndragenter prop Fires on the target element when the user drags the object to a valid drop target.
	//     * @param ev The drag event.
	// js:"ondragenter"
	SetOndragenter(ondragenter func(Event))

	// Ondragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	//     * @param ev The drag event.
	// js:"ondragleave"
	Ondragleave() (ondragleave func(Event))

	// SetOndragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	//     * @param ev The drag event.
	// js:"ondragleave"
	SetOndragleave(ondragleave func(Event))

	// Ondragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
	//     * @param ev The event.
	// js:"ondragover"
	Ondragover() (ondragover func(Event))

	// SetOndragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
	//     * @param ev The event.
	// js:"ondragover"
	SetOndragover(ondragover func(Event))

	// Ondragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
	//     * @param ev The event.
	// js:"ondragstart"
	Ondragstart() (ondragstart func(Event))

	// SetOndragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
	//     * @param ev The event.
	// js:"ondragstart"
	SetOndragstart(ondragstart func(Event))

	// Ondrop prop
	// js:"ondrop"
	Ondrop() (ondrop func(Event))

	// SetOndrop prop
	// js:"ondrop"
	SetOndrop(ondrop func(Event))

	// Ondurationchange prop Occurs when the duration attribute is updated.
	//     * @param ev The event.
	// js:"ondurationchange"
	Ondurationchange() (ondurationchange func(Event))

	// SetOndurationchange prop Occurs when the duration attribute is updated.
	//     * @param ev The event.
	// js:"ondurationchange"
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied prop Occurs when the media element is reset to its initial state.
	//     * @param ev The event.
	// js:"onemptied"
	Onemptied() (onemptied func(Event))

	// SetOnemptied prop Occurs when the media element is reset to its initial state.
	//     * @param ev The event.
	// js:"onemptied"
	SetOnemptied(onemptied func(Event))

	// Onended prop Occurs when the end of playback is reached.
	//     * @param ev The event
	// js:"onended"
	Onended() (onended func(Event))

	// SetOnended prop Occurs when the end of playback is reached.
	//     * @param ev The event
	// js:"onended"
	SetOnended(onended func(Event))

	// Onerror prop Fires when an error occurs during object loading.
	//     * @param ev The event.
	// js:"onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop Fires when an error occurs during object loading.
	//     * @param ev The event.
	// js:"onerror"
	SetOnerror(onerror func(Event))

	// Onfocus prop Fires when the object receives focus.
	//     * @param ev The event.
	// js:"onfocus"
	Onfocus() (onfocus func(Event))

	// SetOnfocus prop Fires when the object receives focus.
	//     * @param ev The event.
	// js:"onfocus"
	SetOnfocus(onfocus func(Event))

	// Onfullscreenchange prop
	// js:"onfullscreenchange"
	Onfullscreenchange() (onfullscreenchange func(Event))

	// SetOnfullscreenchange prop
	// js:"onfullscreenchange"
	SetOnfullscreenchange(onfullscreenchange func(Event))

	// Onfullscreenerror prop
	// js:"onfullscreenerror"
	Onfullscreenerror() (onfullscreenerror func(Event))

	// SetOnfullscreenerror prop
	// js:"onfullscreenerror"
	SetOnfullscreenerror(onfullscreenerror func(Event))

	// Oninput prop
	// js:"oninput"
	Oninput() (oninput func(Event))

	// SetOninput prop
	// js:"oninput"
	SetOninput(oninput func(Event))

	// Oninvalid prop
	// js:"oninvalid"
	Oninvalid() (oninvalid func(Event))

	// SetOninvalid prop
	// js:"oninvalid"
	SetOninvalid(oninvalid func(Event))

	// Onkeydown prop Fires when the user presses a key.
	//     * @param ev The keyboard event
	// js:"onkeydown"
	Onkeydown() (onkeydown func(Event))

	// SetOnkeydown prop Fires when the user presses a key.
	//     * @param ev The keyboard event
	// js:"onkeydown"
	SetOnkeydown(onkeydown func(Event))

	// Onkeypress prop Fires when the user presses an alphanumeric key.
	//     * @param ev The event.
	// js:"onkeypress"
	Onkeypress() (onkeypress func(Event))

	// SetOnkeypress prop Fires when the user presses an alphanumeric key.
	//     * @param ev The event.
	// js:"onkeypress"
	SetOnkeypress(onkeypress func(Event))

	// Onkeyup prop Fires when the user releases a key.
	//     * @param ev The keyboard event
	// js:"onkeyup"
	Onkeyup() (onkeyup func(Event))

	// SetOnkeyup prop Fires when the user releases a key.
	//     * @param ev The keyboard event
	// js:"onkeyup"
	SetOnkeyup(onkeyup func(Event))

	// Onload prop Fires immediately after the browser loads the object.
	//     * @param ev The event.
	// js:"onload"
	Onload() (onload func(Event))

	// SetOnload prop Fires immediately after the browser loads the object.
	//     * @param ev The event.
	// js:"onload"
	SetOnload(onload func(Event))

	// Onloadeddata prop Occurs when media data is loaded at the current playback position.
	//     * @param ev The event.
	// js:"onloadeddata"
	Onloadeddata() (onloadeddata func(Event))

	// SetOnloadeddata prop Occurs when media data is loaded at the current playback position.
	//     * @param ev The event.
	// js:"onloadeddata"
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
	//     * @param ev The event.
	// js:"onloadedmetadata"
	Onloadedmetadata() (onloadedmetadata func(Event))

	// SetOnloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
	//     * @param ev The event.
	// js:"onloadedmetadata"
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart prop Occurs when Internet Explorer begins looking for media data.
	//     * @param ev The event.
	// js:"onloadstart"
	Onloadstart() (onloadstart func(Event))

	// SetOnloadstart prop Occurs when Internet Explorer begins looking for media data.
	//     * @param ev The event.
	// js:"onloadstart"
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown prop Fires when the user clicks the object with either mouse button.
	//     * @param ev The mouse event.
	// js:"onmousedown"
	Onmousedown() (onmousedown func(Event))

	// SetOnmousedown prop Fires when the user clicks the object with either mouse button.
	//     * @param ev The mouse event.
	// js:"onmousedown"
	SetOnmousedown(onmousedown func(Event))

	// Onmousemove prop Fires when the user moves the mouse over the object.
	//     * @param ev The mouse event.
	// js:"onmousemove"
	Onmousemove() (onmousemove func(Event))

	// SetOnmousemove prop Fires when the user moves the mouse over the object.
	//     * @param ev The mouse event.
	// js:"onmousemove"
	SetOnmousemove(onmousemove func(Event))

	// Onmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
	//     * @param ev The mouse event.
	// js:"onmouseout"
	Onmouseout() (onmouseout func(Event))

	// SetOnmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
	//     * @param ev The mouse event.
	// js:"onmouseout"
	SetOnmouseout(onmouseout func(Event))

	// Onmouseover prop Fires when the user moves the mouse pointer into the object.
	//     * @param ev The mouse event.
	// js:"onmouseover"
	Onmouseover() (onmouseover func(Event))

	// SetOnmouseover prop Fires when the user moves the mouse pointer into the object.
	//     * @param ev The mouse event.
	// js:"onmouseover"
	SetOnmouseover(onmouseover func(Event))

	// Onmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
	//     * @param ev The mouse event.
	// js:"onmouseup"
	Onmouseup() (onmouseup func(Event))

	// SetOnmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
	//     * @param ev The mouse event.
	// js:"onmouseup"
	SetOnmouseup(onmouseup func(Event))

	// Onmousewheel prop Fires when the wheel button is rotated.
	//     * @param ev The mouse event
	// js:"onmousewheel"
	Onmousewheel() (onmousewheel func(Event))

	// SetOnmousewheel prop Fires when the wheel button is rotated.
	//     * @param ev The mouse event
	// js:"onmousewheel"
	SetOnmousewheel(onmousewheel func(Event))

	// Onmscontentzoom prop
	// js:"onmscontentzoom"
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// SetOnmscontentzoom prop
	// js:"onmscontentzoom"
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsgesturechange prop
	// js:"onmsgesturechange"
	Onmsgesturechange() (onmsgesturechange func(Event))

	// SetOnmsgesturechange prop
	// js:"onmsgesturechange"
	SetOnmsgesturechange(onmsgesturechange func(Event))

	// Onmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	Onmsgesturedoubletap() (onmsgesturedoubletap func(Event))

	// SetOnmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event))

	// Onmsgestureend prop
	// js:"onmsgestureend"
	Onmsgestureend() (onmsgestureend func(Event))

	// SetOnmsgestureend prop
	// js:"onmsgestureend"
	SetOnmsgestureend(onmsgestureend func(Event))

	// Onmsgesturehold prop
	// js:"onmsgesturehold"
	Onmsgesturehold() (onmsgesturehold func(Event))

	// SetOnmsgesturehold prop
	// js:"onmsgesturehold"
	SetOnmsgesturehold(onmsgesturehold func(Event))

	// Onmsgesturestart prop
	// js:"onmsgesturestart"
	Onmsgesturestart() (onmsgesturestart func(Event))

	// SetOnmsgesturestart prop
	// js:"onmsgesturestart"
	SetOnmsgesturestart(onmsgesturestart func(Event))

	// Onmsgesturetap prop
	// js:"onmsgesturetap"
	Onmsgesturetap() (onmsgesturetap func(Event))

	// SetOnmsgesturetap prop
	// js:"onmsgesturetap"
	SetOnmsgesturetap(onmsgesturetap func(Event))

	// Onmsinertiastart prop
	// js:"onmsinertiastart"
	Onmsinertiastart() (onmsinertiastart func(Event))

	// SetOnmsinertiastart prop
	// js:"onmsinertiastart"
	SetOnmsinertiastart(onmsinertiastart func(Event))

	// Onmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// SetOnmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmspointercancel prop
	// js:"onmspointercancel"
	Onmspointercancel() (onmspointercancel func(Event))

	// SetOnmspointercancel prop
	// js:"onmspointercancel"
	SetOnmspointercancel(onmspointercancel func(Event))

	// Onmspointerdown prop
	// js:"onmspointerdown"
	Onmspointerdown() (onmspointerdown func(Event))

	// SetOnmspointerdown prop
	// js:"onmspointerdown"
	SetOnmspointerdown(onmspointerdown func(Event))

	// Onmspointerenter prop
	// js:"onmspointerenter"
	Onmspointerenter() (onmspointerenter func(Event))

	// SetOnmspointerenter prop
	// js:"onmspointerenter"
	SetOnmspointerenter(onmspointerenter func(Event))

	// Onmspointerleave prop
	// js:"onmspointerleave"
	Onmspointerleave() (onmspointerleave func(Event))

	// SetOnmspointerleave prop
	// js:"onmspointerleave"
	SetOnmspointerleave(onmspointerleave func(Event))

	// Onmspointermove prop
	// js:"onmspointermove"
	Onmspointermove() (onmspointermove func(Event))

	// SetOnmspointermove prop
	// js:"onmspointermove"
	SetOnmspointermove(onmspointermove func(Event))

	// Onmspointerout prop
	// js:"onmspointerout"
	Onmspointerout() (onmspointerout func(Event))

	// SetOnmspointerout prop
	// js:"onmspointerout"
	SetOnmspointerout(onmspointerout func(Event))

	// Onmspointerover prop
	// js:"onmspointerover"
	Onmspointerover() (onmspointerover func(Event))

	// SetOnmspointerover prop
	// js:"onmspointerover"
	SetOnmspointerover(onmspointerover func(Event))

	// Onmspointerup prop
	// js:"onmspointerup"
	Onmspointerup() (onmspointerup func(Event))

	// SetOnmspointerup prop
	// js:"onmspointerup"
	SetOnmspointerup(onmspointerup func(Event))

	// Onmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmssitemodejumplistitemremoved"
	Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// SetOnmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmssitemodejumplistitemremoved"
	SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// Onmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmsthumbnailclick"
	Onmsthumbnailclick() (onmsthumbnailclick func(*MSSiteModeEvent))

	// SetOnmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmsthumbnailclick"
	SetOnmsthumbnailclick(onmsthumbnailclick func(*MSSiteModeEvent))

	// Onpause prop Occurs when playback is paused.
	//     * @param ev The event.
	// js:"onpause"
	Onpause() (onpause func(Event))

	// SetOnpause prop Occurs when playback is paused.
	//     * @param ev The event.
	// js:"onpause"
	SetOnpause(onpause func(Event))

	// Onplay prop Occurs when the play method is requested.
	//     * @param ev The event.
	// js:"onplay"
	Onplay() (onplay func(Event))

	// SetOnplay prop Occurs when the play method is requested.
	//     * @param ev The event.
	// js:"onplay"
	SetOnplay(onplay func(Event))

	// Onplaying prop Occurs when the audio or video has started playing.
	//     * @param ev The event.
	// js:"onplaying"
	Onplaying() (onplaying func(Event))

	// SetOnplaying prop Occurs when the audio or video has started playing.
	//     * @param ev The event.
	// js:"onplaying"
	SetOnplaying(onplaying func(Event))

	// Onpointerlockchange prop
	// js:"onpointerlockchange"
	Onpointerlockchange() (onpointerlockchange func(Event))

	// SetOnpointerlockchange prop
	// js:"onpointerlockchange"
	SetOnpointerlockchange(onpointerlockchange func(Event))

	// Onpointerlockerror prop
	// js:"onpointerlockerror"
	Onpointerlockerror() (onpointerlockerror func(Event))

	// SetOnpointerlockerror prop
	// js:"onpointerlockerror"
	SetOnpointerlockerror(onpointerlockerror func(Event))

	// Onprogress prop Occurs to indicate progress while downloading media data.
	//     * @param ev The event.
	// js:"onprogress"
	Onprogress() (onprogress func(Event))

	// SetOnprogress prop Occurs to indicate progress while downloading media data.
	//     * @param ev The event.
	// js:"onprogress"
	SetOnprogress(onprogress func(Event))

	// Onratechange prop Occurs when the playback rate is increased or decreased.
	//     * @param ev The event.
	// js:"onratechange"
	Onratechange() (onratechange func(Event))

	// SetOnratechange prop Occurs when the playback rate is increased or decreased.
	//     * @param ev The event.
	// js:"onratechange"
	SetOnratechange(onratechange func(Event))

	// Onreadystatechange prop Fires when the state of the object has changed.
	//     * @param ev The event
	// js:"onreadystatechange"
	Onreadystatechange() (onreadystatechange func(Event))

	// SetOnreadystatechange prop Fires when the state of the object has changed.
	//     * @param ev The event
	// js:"onreadystatechange"
	SetOnreadystatechange(onreadystatechange func(Event))

	// Onreset prop Fires when the user resets a form.
	//     * @param ev The event.
	// js:"onreset"
	Onreset() (onreset func(Event))

	// SetOnreset prop Fires when the user resets a form.
	//     * @param ev The event.
	// js:"onreset"
	SetOnreset(onreset func(Event))

	// Onscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
	//     * @param ev The event.
	// js:"onscroll"
	Onscroll() (onscroll func(Event))

	// SetOnscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
	//     * @param ev The event.
	// js:"onscroll"
	SetOnscroll(onscroll func(Event))

	// Onseeked prop Occurs when the seek operation ends.
	//     * @param ev The event.
	// js:"onseeked"
	Onseeked() (onseeked func(Event))

	// SetOnseeked prop Occurs when the seek operation ends.
	//     * @param ev The event.
	// js:"onseeked"
	SetOnseeked(onseeked func(Event))

	// Onseeking prop Occurs when the current playback position is moved.
	//     * @param ev The event.
	// js:"onseeking"
	Onseeking() (onseeking func(Event))

	// SetOnseeking prop Occurs when the current playback position is moved.
	//     * @param ev The event.
	// js:"onseeking"
	SetOnseeking(onseeking func(Event))

	// Onselect prop Fires when the current selection changes.
	//     * @param ev The event.
	// js:"onselect"
	Onselect() (onselect func(Event))

	// SetOnselect prop Fires when the current selection changes.
	//     * @param ev The event.
	// js:"onselect"
	SetOnselect(onselect func(Event))

	// Onselectionchange prop Fires when the selection state of a document changes.
	//     * @param ev The event.
	// js:"onselectionchange"
	Onselectionchange() (onselectionchange func(Event))

	// SetOnselectionchange prop Fires when the selection state of a document changes.
	//     * @param ev The event.
	// js:"onselectionchange"
	SetOnselectionchange(onselectionchange func(Event))

	// Onselectstart prop
	// js:"onselectstart"
	Onselectstart() (onselectstart func(Event))

	// SetOnselectstart prop
	// js:"onselectstart"
	SetOnselectstart(onselectstart func(Event))

	// Onstalled prop Occurs when the download has stopped.
	//     * @param ev The event.
	// js:"onstalled"
	Onstalled() (onstalled func(Event))

	// SetOnstalled prop Occurs when the download has stopped.
	//     * @param ev The event.
	// js:"onstalled"
	SetOnstalled(onstalled func(Event))

	// Onstop prop Fires when the user clicks the Stop button or leaves the Web page.
	//     * @param ev The event.
	// js:"onstop"
	Onstop() (onstop func(Event))

	// SetOnstop prop Fires when the user clicks the Stop button or leaves the Web page.
	//     * @param ev The event.
	// js:"onstop"
	SetOnstop(onstop func(Event))

	// Onsubmit prop
	// js:"onsubmit"
	Onsubmit() (onsubmit func(Event))

	// SetOnsubmit prop
	// js:"onsubmit"
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend prop Occurs if the load operation has been intentionally halted.
	//     * @param ev The event.
	// js:"onsuspend"
	Onsuspend() (onsuspend func(Event))

	// SetOnsuspend prop Occurs if the load operation has been intentionally halted.
	//     * @param ev The event.
	// js:"onsuspend"
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate prop Occurs to indicate the current playback position.
	//     * @param ev The event.
	// js:"ontimeupdate"
	Ontimeupdate() (ontimeupdate func(Event))

	// SetOntimeupdate prop Occurs to indicate the current playback position.
	//     * @param ev The event.
	// js:"ontimeupdate"
	SetOntimeupdate(ontimeupdate func(Event))

	// Ontouchcancel prop
	// js:"ontouchcancel"
	Ontouchcancel() (ontouchcancel func(Event))

	// SetOntouchcancel prop
	// js:"ontouchcancel"
	SetOntouchcancel(ontouchcancel func(Event))

	// Ontouchend prop
	// js:"ontouchend"
	Ontouchend() (ontouchend func(Event))

	// SetOntouchend prop
	// js:"ontouchend"
	SetOntouchend(ontouchend func(Event))

	// Ontouchmove prop
	// js:"ontouchmove"
	Ontouchmove() (ontouchmove func(Event))

	// SetOntouchmove prop
	// js:"ontouchmove"
	SetOntouchmove(ontouchmove func(Event))

	// Ontouchstart prop
	// js:"ontouchstart"
	Ontouchstart() (ontouchstart func(Event))

	// SetOntouchstart prop
	// js:"ontouchstart"
	SetOntouchstart(ontouchstart func(Event))

	// Onvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
	//     * @param ev The event.
	// js:"onvolumechange"
	Onvolumechange() (onvolumechange func(Event))

	// SetOnvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
	//     * @param ev The event.
	// js:"onvolumechange"
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
	//     * @param ev The event.
	// js:"onwaiting"
	Onwaiting() (onwaiting func(Event))

	// SetOnwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
	//     * @param ev The event.
	// js:"onwaiting"
	SetOnwaiting(onwaiting func(Event))

	// Onwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// SetOnwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// SetOnwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// Plugins prop
	// js:"plugins"
	Plugins() (plugins HTMLCollection)

	// PointerLockElement prop
	// js:"pointerLockElement"
	PointerLockElement() (pointerLockElement Element)

	// ReadyState prop Retrieves a value that indicates the current state of the object.
	// js:"readyState"
	ReadyState() (readyState string)

	// Referrer prop Gets the URL of the location that referred the user to the current page.
	// js:"referrer"
	Referrer() (referrer string)

	// RootElement prop Gets the root svg element in the document hierarchy.
	// js:"rootElement"
	RootElement() (rootElement *SVGSVGElement)

	// Scripts prop Retrieves a collection of all script objects in the document.
	// js:"scripts"
	Scripts() (scripts HTMLCollection)

	// ScrollingElement prop
	// js:"scrollingElement"
	ScrollingElement() (scrollingElement Element)

	// StyleSheets prop Retrieves a collection of styleSheet objects representing the style sheets that correspond to each instance of a link or style object in the document.
	// js:"styleSheets"
	StyleSheets() (styleSheets *StyleSheetList)

	// Title prop Contains the title of the document.
	// js:"title"
	Title() (title string)

	// SetTitle prop Contains the title of the document.
	// js:"title"
	SetTitle(title string)

	// URL prop Sets or gets the URL for the current document.
	// js:"URL"
	URL() (URL string)

	// URLUnencoded prop Gets the URL for the document, stripped of any character encoding.
	// js:"URLUnencoded"
	URLUnencoded() (URLUnencoded string)

	// VisibilityState prop
	// js:"visibilityState"
	VisibilityState() (visibilityState *visibilitystate.VisibilityState)

	// VlinkColor prop Sets or gets the color of the links that the user has visited.
	// js:"vlinkColor"
	VlinkColor() (vlinkColor string)

	// SetVlinkColor prop Sets or gets the color of the links that the user has visited.
	// js:"vlinkColor"
	SetVlinkColor(vlinkColor string)

	// WebkitCurrentFullScreenElement prop
	// js:"webkitCurrentFullScreenElement"
	WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement Element)

	// WebkitFullscreenElement prop
	// js:"webkitFullscreenElement"
	WebkitFullscreenElement() (webkitFullscreenElement Element)

	// WebkitFullscreenEnabled prop
	// js:"webkitFullscreenEnabled"
	WebkitFullscreenEnabled() (webkitFullscreenEnabled bool)

	// WebkitIsFullScreen prop
	// js:"webkitIsFullScreen"
	WebkitIsFullScreen() (webkitIsFullScreen bool)

	// XMLEncoding prop
	// js:"xmlEncoding"
	XMLEncoding() (xmlEncoding string)

	// XMLStandalone prop
	// js:"xmlStandalone"
	XMLStandalone() (xmlStandalone bool)

	// SetXMLStandalone prop
	// js:"xmlStandalone"
	SetXMLStandalone(xmlStandalone bool)

	// XMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
	// js:"xmlVersion"
	XMLVersion() (xmlVersion string)

	// SetXMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
	// js:"xmlVersion"
	SetXMLVersion(xmlVersion string)
}
