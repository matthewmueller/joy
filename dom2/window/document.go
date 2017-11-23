package window

import (
	"github.com/matthewmueller/golly/dom2/htmlallcollection"
	"github.com/matthewmueller/golly/dom2/location"
	"github.com/matthewmueller/golly/dom2/visibilitystate"
	"github.com/matthewmueller/golly/dom2/xpathnsresolver"
)

// js:"Document,omit"
type Document interface {
	Node
	GlobalEventHandlers
	NodeSelector
	DocumentEvent

	// AdoptNode
	AdoptNode(source Node) (n Node)

	// CaptureEvents
	CaptureEvents()

	// CaretRangeFromPoint
	CaretRangeFromPoint(x float32, y float32) (r *Range)

	// Clear
	Clear()

	// Close Closes an output stream and forces the sent data to display.
	Close()

	// CreateAttribute Creates an attribute object with a specified name.
	//     * @param name String that sets the attribute object's name.
	CreateAttribute(name string) (a *Attr)

	// CreateAttributeNS
	CreateAttributeNS(namespaceURI string, qualifiedName string) (a *Attr)

	// CreateCDATASection
	CreateCDATASection(data string) (c *CDATASection)

	// CreateComment Creates a comment object with the specified data.
	//     * @param data Sets the comment object's data.
	CreateComment(data string) (c *Comment)

	// CreateDocumentFragment Creates a new document.
	CreateDocumentFragment() (d *DocumentFragment)

	// CreateElement Creates an instance of the element for the specified tag.
	//     * @param tagName The name of an element.
	CreateElement(tagName string) (e Element)

	// CreateElementNS
	CreateElementNS(namespaceURI string, qualifiedName string) (e Element)

	// CreateExpression
	CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (x *XPathExpression)

	// CreateNodeIterator Creates a NodeIterator object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list
	//     * @param filter A custom NodeFilter function to use. For more information, see filter. Use null for no filter.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	CreateNodeIterator(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (n *NodeIterator)

	// CreateNSResolver
	CreateNSResolver(nodeResolver Node) (x *xpathnsresolver.XPathNSResolver)

	// CreateProcessingInstruction
	CreateProcessingInstruction(target string, data string) (p *ProcessingInstruction)

	// CreateRange Returns an empty range object that has both of its boundary points positioned at the beginning of the document.
	CreateRange() (r *Range)

	// CreateTextNode Creates a text string from the specified value.
	//     * @param data String that specifies the nodeValue property of the text node.
	CreateTextNode(data string) (t Text)

	// CreateTouch
	CreateTouch(view *Window, target EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *Touch)

	// CreateTouchList
	CreateTouchList(touches *Touch) (t *TouchList)

	// CreateTreeWalker Creates a TreeWalker object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list. For more information, see whatToShow.
	//     * @param filter A custom NodeFilter function to use.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	CreateTreeWalker(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (t *TreeWalker)

	// ElementFromPoint Returns the element for the specified x coordinate and the specified y coordinate.
	//     * @param x The x-offset
	//     * @param y The y-offset
	ElementFromPoint(x int, y int) (e Element)

	// Evaluate
	Evaluate(expression string, contextNode Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *XPathResult) (x *XPathResult)

	// ExecCommand Executes a command on the current document, current selection, or the given range.
	//     * @param commandId String that specifies the command to execute. This command can be any of the command identifiers that can be executed in script.
	//     * @param showUI Display the user interface, defaults to false.
	//     * @param value Value to assign.
	ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool)

	// ExecCommandShowHelp Displays help information for the given command identifier.
	//     * @param commandId Displays help information for the given command identifier.
	ExecCommandShowHelp(commandId string) (b bool)

	// ExitFullscreen
	ExitFullscreen()

	// ExitPointerLock
	ExitPointerLock()

	// Focus Causes the element to receive the focus and executes the code specified by the onfocus event.
	Focus()

	// GetElementByID Returns a reference to the first object with the specified value of the ID or NAME attribute.
	//     * @param elementId String that specifies the ID value. Case-insensitive.
	GetElementByID(elementId string) (e Element)

	// GetElementsByClassName
	GetElementsByClassName(classNames string) (n *NodeList)

	// GetElementsByName Gets a collection of objects based on the value of the NAME or ID attribute.
	//     * @param elementName Gets a collection of objects based on the value of the NAME or ID attribute.
	GetElementsByName(elementName string) (n *NodeList)

	// GetElementsByTagName Retrieves a collection of objects based on the specified element name.
	//     * @param name Specifies the name of an element.
	GetElementsByTagName(tagname string) (n *NodeList)

	// GetElementsByTagNameNS
	GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// GetSelection Returns an object representing the current selection of the document that is loaded into the object displaying a webpage.
	GetSelection() (s *Selection)

	// HasFocus Gets a value indicating whether the object currently has focus.
	HasFocus() (b bool)

	// ImportNode
	ImportNode(importedNode Node, deep bool) (n Node)

	// MsElementsFromPoint
	MsElementsFromPoint(x float32, y float32) (n *NodeList)

	// MsElementsFromRect
	MsElementsFromRect(left float32, top float32, width float32, height float32) (n *NodeList)

	// Open Opens a new window and loads a document specified by a given URL. Also, opens a new window that uses the url parameter and the name parameter to collect the output of the write method and the writeln method.
	//     * @param url Specifies a MIME type for the document.
	//     * @param name Specifies the name of the window. This name is used as the value for the TARGET attribute on a form or an anchor element.
	//     * @param features Contains a list of items separated by commas. Each item consists of an option and a value, separated by an equals sign (for example, "fullscreen=yes, toolbar=yes"). The following values are supported.
	//     * @param replace Specifies whether the existing entry for the document is replaced in the history list.
	Open(url *string, name *string, features *string, replace *bool) (i interface{})

	// QueryCommandEnabled Returns a Boolean value that indicates whether a specified command can be successfully executed using execCommand, given the current state of the document.
	//     * @param commandId Specifies a command identifier.
	QueryCommandEnabled(commandId string) (b bool)

	// QueryCommandIndeterm Returns a Boolean value that indicates whether the specified command is in the indeterminate state.
	//     * @param commandId String that specifies a command identifier.
	QueryCommandIndeterm(commandId string) (b bool)

	// QueryCommandState Returns a Boolean value that indicates the current state of the command.
	//     * @param commandId String that specifies a command identifier.
	QueryCommandState(commandId string) (b bool)

	// QueryCommandSupported Returns a Boolean value that indicates whether the current command is supported on the current range.
	//     * @param commandId Specifies a command identifier.
	QueryCommandSupported(commandId string) (b bool)

	// QueryCommandText Retrieves the string associated with a command.
	//     * @param commandId String that contains the identifier of a command. This can be any command identifier given in the list of Command Identifiers.
	QueryCommandText(commandId string) (s string)

	// QueryCommandValue Returns the current value of the document, range, or current selection for the given command.
	//     * @param commandId String that specifies a command identifier.
	QueryCommandValue(commandId string) (s string)

	// ReleaseEvents
	ReleaseEvents()

	// UpdateSettings Allows updating the print settings for the page.
	UpdateSettings()

	// WebkitCancelFullScreen
	WebkitCancelFullScreen()

	// WebkitExitFullscreen
	WebkitExitFullscreen()

	// Write Writes one or more HTML expressions to a document in the specified window.
	//     * @param content Specifies the text and HTML tags to write.
	Write(content string)

	// Writeln Writes one or more HTML expressions, followed by a carriage return, to a document in the specified window.
	//     * @param content The text and HTML tags to write.
	Writeln(content string)

	// ActiveElement Gets the object that has the focus when the parent document has focus.
	ActiveElement() (activeElement Element)

	// AlinkColor Sets or gets the color of all active links in the document.
	AlinkColor() (alinkColor string)

	// AlinkColor Sets or gets the color of all active links in the document.
	SetAlinkColor(alinkColor string)

	// All Returns a reference to the collection of elements contained by the object.
	All() (all *htmlallcollection.HTMLAllCollection)

	// Anchors Retrieves a collection of all a objects that have a name and/or id property. Objects in this collection are in HTML source order.
	Anchors() (anchors HTMLCollection)

	// Applets Retrieves a collection of all applet objects in the document.
	Applets() (applets HTMLCollection)

	// BgColor Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	BgColor() (bgColor string)

	// BgColor Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	SetBgColor(bgColor string)

	// Body Specifies the beginning and end of the document body.
	Body() (body HTMLElement)

	// Body Specifies the beginning and end of the document body.
	SetBody(body HTMLElement)

	// CharacterSet
	CharacterSet() (characterSet string)

	// Charset Gets or sets the character set used to encode the object.
	Charset() (charset string)

	// Charset Gets or sets the character set used to encode the object.
	SetCharset(charset string)

	// CompatMode Gets a value that indicates whether standards-compliant mode is switched on for the object.
	CompatMode() (compatMode string)

	// Cookie
	Cookie() (cookie string)

	// Cookie
	SetCookie(cookie string)

	// CurrentScript
	CurrentScript() (currentScript interface{})

	// DefaultView
	DefaultView() (defaultView *Window)

	// DesignMode Sets or gets a value that indicates whether the document can be edited.
	DesignMode() (designMode string)

	// DesignMode Sets or gets a value that indicates whether the document can be edited.
	SetDesignMode(designMode string)

	// Dir Sets or retrieves a value that indicates the reading order of the object.
	Dir() (dir string)

	// Dir Sets or retrieves a value that indicates the reading order of the object.
	SetDir(dir string)

	// Doctype Gets an object representing the document type declaration associated with the current document.
	Doctype() (doctype *DocumentType)

	// DocumentElement Gets a reference to the root node of the document.
	DocumentElement() (documentElement Element)

	// Domain Sets or gets the security domain of the document.
	Domain() (domain string)

	// Domain Sets or gets the security domain of the document.
	SetDomain(domain string)

	// Embeds Retrieves a collection of all embed objects in the document.
	Embeds() (embeds HTMLCollection)

	// FgColor Sets or gets the foreground (text) color of the document.
	FgColor() (fgColor string)

	// FgColor Sets or gets the foreground (text) color of the document.
	SetFgColor(fgColor string)

	// Forms Retrieves a collection, in source order, of all form objects in the document.
	Forms() (forms HTMLCollection)

	// FullscreenElement
	FullscreenElement() (fullscreenElement Element)

	// FullscreenEnabled
	FullscreenEnabled() (fullscreenEnabled bool)

	// Head
	Head() (head *HTMLHeadElement)

	// Hidden
	Hidden() (hidden bool)

	// Images Retrieves a collection, in source order, of img objects in the document.
	Images() (images HTMLCollection)

	// Implementation Gets the implementation object of the current document.
	Implementation() (implementation *DOMImplementation)

	// InputEncoding Returns the character encoding used to create the webpage that is loaded into the document object.
	InputEncoding() (inputEncoding string)

	// LastModified Gets the date that the page was last modified, if the page supplies one.
	LastModified() (lastModified string)

	// LinkColor Sets or gets the color of the document links.
	LinkColor() (linkColor string)

	// LinkColor Sets or gets the color of the document links.
	SetLinkColor(linkColor string)

	// Links Retrieves a collection of all a objects that specify the href property and all area objects in the document.
	Links() (links HTMLCollection)

	// Location Contains information about the current URL.
	Location() (location *location.Location)

	// MsCapsLockWarningOff
	MsCapsLockWarningOff() (msCapsLockWarningOff bool)

	// MsCapsLockWarningOff
	SetMsCapsLockWarningOff(msCapsLockWarningOff bool)

	// MsCSSOMElementFloatMetrics
	MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool)

	// MsCSSOMElementFloatMetrics
	SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool)

	// Onabort Fires when the user aborts the download.
	//     * @param ev The event.
	Onabort() (onabort func(Event))

	// Onabort Fires when the user aborts the download.
	//     * @param ev The event.
	SetOnabort(onabort func(Event))

	// Onactivate Fires when the object is set as the active element.
	//     * @param ev The event.
	Onactivate() (onactivate func(Event))

	// Onactivate Fires when the object is set as the active element.
	//     * @param ev The event.
	SetOnactivate(onactivate func(Event))

	// Onbeforeactivate Fires immediately before the object is set as the active element.
	//     * @param ev The event.
	Onbeforeactivate() (onbeforeactivate func(Event))

	// Onbeforeactivate Fires immediately before the object is set as the active element.
	//     * @param ev The event.
	SetOnbeforeactivate(onbeforeactivate func(Event))

	// Onbeforedeactivate Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The event.
	Onbeforedeactivate() (onbeforedeactivate func(Event))

	// Onbeforedeactivate Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The event.
	SetOnbeforedeactivate(onbeforedeactivate func(Event))

	// Onblur Fires when the object loses the input focus.
	//     * @param ev The focus event.
	Onblur() (onblur func(Event))

	// Onblur Fires when the object loses the input focus.
	//     * @param ev The focus event.
	SetOnblur(onblur func(Event))

	// Oncanplay Occurs when playback is possible, but would require further buffering.
	//     * @param ev The event.
	Oncanplay() (oncanplay func(Event))

	// Oncanplay Occurs when playback is possible, but would require further buffering.
	//     * @param ev The event.
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough
	Oncanplaythrough() (oncanplaythrough func(Event))

	// Oncanplaythrough
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange Fires when the contents of the object or selection have changed.
	//     * @param ev The event.
	Onchange() (onchange func(Event))

	// Onchange Fires when the contents of the object or selection have changed.
	//     * @param ev The event.
	SetOnchange(onchange func(Event))

	// Onclick Fires when the user clicks the left mouse button on the object
	//     * @param ev The mouse event.
	Onclick() (onclick func(Event))

	// Onclick Fires when the user clicks the left mouse button on the object
	//     * @param ev The mouse event.
	SetOnclick(onclick func(Event))

	// Oncontextmenu Fires when the user clicks the right mouse button in the client area, opening the context menu.
	//     * @param ev The mouse event.
	Oncontextmenu() (oncontextmenu func(Event))

	// Oncontextmenu Fires when the user clicks the right mouse button in the client area, opening the context menu.
	//     * @param ev The mouse event.
	SetOncontextmenu(oncontextmenu func(Event))

	// Ondblclick Fires when the user double-clicks the object.
	//     * @param ev The mouse event.
	Ondblclick() (ondblclick func(Event))

	// Ondblclick Fires when the user double-clicks the object.
	//     * @param ev The mouse event.
	SetOndblclick(ondblclick func(Event))

	// Ondeactivate Fires when the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The UI Event
	Ondeactivate() (ondeactivate func(Event))

	// Ondeactivate Fires when the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The UI Event
	SetOndeactivate(ondeactivate func(Event))

	// Ondrag Fires on the source object continuously during a drag operation.
	//     * @param ev The event.
	Ondrag() (ondrag func(Event))

	// Ondrag Fires on the source object continuously during a drag operation.
	//     * @param ev The event.
	SetOndrag(ondrag func(Event))

	// Ondragend Fires on the source object when the user releases the mouse at the close of a drag operation.
	//     * @param ev The event.
	Ondragend() (ondragend func(Event))

	// Ondragend Fires on the source object when the user releases the mouse at the close of a drag operation.
	//     * @param ev The event.
	SetOndragend(ondragend func(Event))

	// Ondragenter Fires on the target element when the user drags the object to a valid drop target.
	//     * @param ev The drag event.
	Ondragenter() (ondragenter func(Event))

	// Ondragenter Fires on the target element when the user drags the object to a valid drop target.
	//     * @param ev The drag event.
	SetOndragenter(ondragenter func(Event))

	// Ondragleave Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	//     * @param ev The drag event.
	Ondragleave() (ondragleave func(Event))

	// Ondragleave Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	//     * @param ev The drag event.
	SetOndragleave(ondragleave func(Event))

	// Ondragover Fires on the target element continuously while the user drags the object over a valid drop target.
	//     * @param ev The event.
	Ondragover() (ondragover func(Event))

	// Ondragover Fires on the target element continuously while the user drags the object over a valid drop target.
	//     * @param ev The event.
	SetOndragover(ondragover func(Event))

	// Ondragstart Fires on the source object when the user starts to drag a text selection or selected object.
	//     * @param ev The event.
	Ondragstart() (ondragstart func(Event))

	// Ondragstart Fires on the source object when the user starts to drag a text selection or selected object.
	//     * @param ev The event.
	SetOndragstart(ondragstart func(Event))

	// Ondrop
	Ondrop() (ondrop func(Event))

	// Ondrop
	SetOndrop(ondrop func(Event))

	// Ondurationchange Occurs when the duration attribute is updated.
	//     * @param ev The event.
	Ondurationchange() (ondurationchange func(Event))

	// Ondurationchange Occurs when the duration attribute is updated.
	//     * @param ev The event.
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied Occurs when the media element is reset to its initial state.
	//     * @param ev The event.
	Onemptied() (onemptied func(Event))

	// Onemptied Occurs when the media element is reset to its initial state.
	//     * @param ev The event.
	SetOnemptied(onemptied func(Event))

	// Onended Occurs when the end of playback is reached.
	//     * @param ev The event
	Onended() (onended func(Event))

	// Onended Occurs when the end of playback is reached.
	//     * @param ev The event
	SetOnended(onended func(Event))

	// Onerror Fires when an error occurs during object loading.
	//     * @param ev The event.
	Onerror() (onerror func(Event))

	// Onerror Fires when an error occurs during object loading.
	//     * @param ev The event.
	SetOnerror(onerror func(Event))

	// Onfocus Fires when the object receives focus.
	//     * @param ev The event.
	Onfocus() (onfocus func(Event))

	// Onfocus Fires when the object receives focus.
	//     * @param ev The event.
	SetOnfocus(onfocus func(Event))

	// Onfullscreenchange
	Onfullscreenchange() (onfullscreenchange func(Event))

	// Onfullscreenchange
	SetOnfullscreenchange(onfullscreenchange func(Event))

	// Onfullscreenerror
	Onfullscreenerror() (onfullscreenerror func(Event))

	// Onfullscreenerror
	SetOnfullscreenerror(onfullscreenerror func(Event))

	// Oninput
	Oninput() (oninput func(Event))

	// Oninput
	SetOninput(oninput func(Event))

	// Oninvalid
	Oninvalid() (oninvalid func(Event))

	// Oninvalid
	SetOninvalid(oninvalid func(Event))

	// Onkeydown Fires when the user presses a key.
	//     * @param ev The keyboard event
	Onkeydown() (onkeydown func(Event))

	// Onkeydown Fires when the user presses a key.
	//     * @param ev The keyboard event
	SetOnkeydown(onkeydown func(Event))

	// Onkeypress Fires when the user presses an alphanumeric key.
	//     * @param ev The event.
	Onkeypress() (onkeypress func(Event))

	// Onkeypress Fires when the user presses an alphanumeric key.
	//     * @param ev The event.
	SetOnkeypress(onkeypress func(Event))

	// Onkeyup Fires when the user releases a key.
	//     * @param ev The keyboard event
	Onkeyup() (onkeyup func(Event))

	// Onkeyup Fires when the user releases a key.
	//     * @param ev The keyboard event
	SetOnkeyup(onkeyup func(Event))

	// Onload Fires immediately after the browser loads the object.
	//     * @param ev The event.
	Onload() (onload func(Event))

	// Onload Fires immediately after the browser loads the object.
	//     * @param ev The event.
	SetOnload(onload func(Event))

	// Onloadeddata Occurs when media data is loaded at the current playback position.
	//     * @param ev The event.
	Onloadeddata() (onloadeddata func(Event))

	// Onloadeddata Occurs when media data is loaded at the current playback position.
	//     * @param ev The event.
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata Occurs when the duration and dimensions of the media have been determined.
	//     * @param ev The event.
	Onloadedmetadata() (onloadedmetadata func(Event))

	// Onloadedmetadata Occurs when the duration and dimensions of the media have been determined.
	//     * @param ev The event.
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart Occurs when Internet Explorer begins looking for media data.
	//     * @param ev The event.
	Onloadstart() (onloadstart func(Event))

	// Onloadstart Occurs when Internet Explorer begins looking for media data.
	//     * @param ev The event.
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown Fires when the user clicks the object with either mouse button.
	//     * @param ev The mouse event.
	Onmousedown() (onmousedown func(Event))

	// Onmousedown Fires when the user clicks the object with either mouse button.
	//     * @param ev The mouse event.
	SetOnmousedown(onmousedown func(Event))

	// Onmousemove Fires when the user moves the mouse over the object.
	//     * @param ev The mouse event.
	Onmousemove() (onmousemove func(Event))

	// Onmousemove Fires when the user moves the mouse over the object.
	//     * @param ev The mouse event.
	SetOnmousemove(onmousemove func(Event))

	// Onmouseout Fires when the user moves the mouse pointer outside the boundaries of the object.
	//     * @param ev The mouse event.
	Onmouseout() (onmouseout func(Event))

	// Onmouseout Fires when the user moves the mouse pointer outside the boundaries of the object.
	//     * @param ev The mouse event.
	SetOnmouseout(onmouseout func(Event))

	// Onmouseover Fires when the user moves the mouse pointer into the object.
	//     * @param ev The mouse event.
	Onmouseover() (onmouseover func(Event))

	// Onmouseover Fires when the user moves the mouse pointer into the object.
	//     * @param ev The mouse event.
	SetOnmouseover(onmouseover func(Event))

	// Onmouseup Fires when the user releases a mouse button while the mouse is over the object.
	//     * @param ev The mouse event.
	Onmouseup() (onmouseup func(Event))

	// Onmouseup Fires when the user releases a mouse button while the mouse is over the object.
	//     * @param ev The mouse event.
	SetOnmouseup(onmouseup func(Event))

	// Onmousewheel Fires when the wheel button is rotated.
	//     * @param ev The mouse event
	Onmousewheel() (onmousewheel func(Event))

	// Onmousewheel Fires when the wheel button is rotated.
	//     * @param ev The mouse event
	SetOnmousewheel(onmousewheel func(Event))

	// Onmscontentzoom
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// Onmscontentzoom
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsgesturechange
	Onmsgesturechange() (onmsgesturechange func(Event))

	// Onmsgesturechange
	SetOnmsgesturechange(onmsgesturechange func(Event))

	// Onmsgesturedoubletap
	Onmsgesturedoubletap() (onmsgesturedoubletap func(Event))

	// Onmsgesturedoubletap
	SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event))

	// Onmsgestureend
	Onmsgestureend() (onmsgestureend func(Event))

	// Onmsgestureend
	SetOnmsgestureend(onmsgestureend func(Event))

	// Onmsgesturehold
	Onmsgesturehold() (onmsgesturehold func(Event))

	// Onmsgesturehold
	SetOnmsgesturehold(onmsgesturehold func(Event))

	// Onmsgesturestart
	Onmsgesturestart() (onmsgesturestart func(Event))

	// Onmsgesturestart
	SetOnmsgesturestart(onmsgesturestart func(Event))

	// Onmsgesturetap
	Onmsgesturetap() (onmsgesturetap func(Event))

	// Onmsgesturetap
	SetOnmsgesturetap(onmsgesturetap func(Event))

	// Onmsinertiastart
	Onmsinertiastart() (onmsinertiastart func(Event))

	// Onmsinertiastart
	SetOnmsinertiastart(onmsinertiastart func(Event))

	// Onmsmanipulationstatechanged
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmsmanipulationstatechanged
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmspointercancel
	Onmspointercancel() (onmspointercancel func(Event))

	// Onmspointercancel
	SetOnmspointercancel(onmspointercancel func(Event))

	// Onmspointerdown
	Onmspointerdown() (onmspointerdown func(Event))

	// Onmspointerdown
	SetOnmspointerdown(onmspointerdown func(Event))

	// Onmspointerenter
	Onmspointerenter() (onmspointerenter func(Event))

	// Onmspointerenter
	SetOnmspointerenter(onmspointerenter func(Event))

	// Onmspointerleave
	Onmspointerleave() (onmspointerleave func(Event))

	// Onmspointerleave
	SetOnmspointerleave(onmspointerleave func(Event))

	// Onmspointermove
	Onmspointermove() (onmspointermove func(Event))

	// Onmspointermove
	SetOnmspointermove(onmspointermove func(Event))

	// Onmspointerout
	Onmspointerout() (onmspointerout func(Event))

	// Onmspointerout
	SetOnmspointerout(onmspointerout func(Event))

	// Onmspointerover
	Onmspointerover() (onmspointerover func(Event))

	// Onmspointerover
	SetOnmspointerover(onmspointerover func(Event))

	// Onmspointerup
	Onmspointerup() (onmspointerup func(Event))

	// Onmspointerup
	SetOnmspointerup(onmspointerup func(Event))

	// Onmssitemodejumplistitemremoved Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	//     * @param ev The event.
	Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// Onmssitemodejumplistitemremoved Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	//     * @param ev The event.
	SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// Onmsthumbnailclick Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	//     * @param ev The event.
	Onmsthumbnailclick() (onmsthumbnailclick func(*MSSiteModeEvent))

	// Onmsthumbnailclick Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	//     * @param ev The event.
	SetOnmsthumbnailclick(onmsthumbnailclick func(*MSSiteModeEvent))

	// Onpause Occurs when playback is paused.
	//     * @param ev The event.
	Onpause() (onpause func(Event))

	// Onpause Occurs when playback is paused.
	//     * @param ev The event.
	SetOnpause(onpause func(Event))

	// Onplay Occurs when the play method is requested.
	//     * @param ev The event.
	Onplay() (onplay func(Event))

	// Onplay Occurs when the play method is requested.
	//     * @param ev The event.
	SetOnplay(onplay func(Event))

	// Onplaying Occurs when the audio or video has started playing.
	//     * @param ev The event.
	Onplaying() (onplaying func(Event))

	// Onplaying Occurs when the audio or video has started playing.
	//     * @param ev The event.
	SetOnplaying(onplaying func(Event))

	// Onpointerlockchange
	Onpointerlockchange() (onpointerlockchange func(Event))

	// Onpointerlockchange
	SetOnpointerlockchange(onpointerlockchange func(Event))

	// Onpointerlockerror
	Onpointerlockerror() (onpointerlockerror func(Event))

	// Onpointerlockerror
	SetOnpointerlockerror(onpointerlockerror func(Event))

	// Onprogress Occurs to indicate progress while downloading media data.
	//     * @param ev The event.
	Onprogress() (onprogress func(Event))

	// Onprogress Occurs to indicate progress while downloading media data.
	//     * @param ev The event.
	SetOnprogress(onprogress func(Event))

	// Onratechange Occurs when the playback rate is increased or decreased.
	//     * @param ev The event.
	Onratechange() (onratechange func(Event))

	// Onratechange Occurs when the playback rate is increased or decreased.
	//     * @param ev The event.
	SetOnratechange(onratechange func(Event))

	// Onreadystatechange Fires when the state of the object has changed.
	//     * @param ev The event
	Onreadystatechange() (onreadystatechange func(Event))

	// Onreadystatechange Fires when the state of the object has changed.
	//     * @param ev The event
	SetOnreadystatechange(onreadystatechange func(Event))

	// Onreset Fires when the user resets a form.
	//     * @param ev The event.
	Onreset() (onreset func(Event))

	// Onreset Fires when the user resets a form.
	//     * @param ev The event.
	SetOnreset(onreset func(Event))

	// Onscroll Fires when the user repositions the scroll box in the scroll bar on the object.
	//     * @param ev The event.
	Onscroll() (onscroll func(Event))

	// Onscroll Fires when the user repositions the scroll box in the scroll bar on the object.
	//     * @param ev The event.
	SetOnscroll(onscroll func(Event))

	// Onseeked Occurs when the seek operation ends.
	//     * @param ev The event.
	Onseeked() (onseeked func(Event))

	// Onseeked Occurs when the seek operation ends.
	//     * @param ev The event.
	SetOnseeked(onseeked func(Event))

	// Onseeking Occurs when the current playback position is moved.
	//     * @param ev The event.
	Onseeking() (onseeking func(Event))

	// Onseeking Occurs when the current playback position is moved.
	//     * @param ev The event.
	SetOnseeking(onseeking func(Event))

	// Onselect Fires when the current selection changes.
	//     * @param ev The event.
	Onselect() (onselect func(Event))

	// Onselect Fires when the current selection changes.
	//     * @param ev The event.
	SetOnselect(onselect func(Event))

	// Onselectionchange Fires when the selection state of a document changes.
	//     * @param ev The event.
	Onselectionchange() (onselectionchange func(Event))

	// Onselectionchange Fires when the selection state of a document changes.
	//     * @param ev The event.
	SetOnselectionchange(onselectionchange func(Event))

	// Onselectstart
	Onselectstart() (onselectstart func(Event))

	// Onselectstart
	SetOnselectstart(onselectstart func(Event))

	// Onstalled Occurs when the download has stopped.
	//     * @param ev The event.
	Onstalled() (onstalled func(Event))

	// Onstalled Occurs when the download has stopped.
	//     * @param ev The event.
	SetOnstalled(onstalled func(Event))

	// Onstop Fires when the user clicks the Stop button or leaves the Web page.
	//     * @param ev The event.
	Onstop() (onstop func(Event))

	// Onstop Fires when the user clicks the Stop button or leaves the Web page.
	//     * @param ev The event.
	SetOnstop(onstop func(Event))

	// Onsubmit
	Onsubmit() (onsubmit func(Event))

	// Onsubmit
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend Occurs if the load operation has been intentionally halted.
	//     * @param ev The event.
	Onsuspend() (onsuspend func(Event))

	// Onsuspend Occurs if the load operation has been intentionally halted.
	//     * @param ev The event.
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate Occurs to indicate the current playback position.
	//     * @param ev The event.
	Ontimeupdate() (ontimeupdate func(Event))

	// Ontimeupdate Occurs to indicate the current playback position.
	//     * @param ev The event.
	SetOntimeupdate(ontimeupdate func(Event))

	// Ontouchcancel
	Ontouchcancel() (ontouchcancel func(Event))

	// Ontouchcancel
	SetOntouchcancel(ontouchcancel func(Event))

	// Ontouchend
	Ontouchend() (ontouchend func(Event))

	// Ontouchend
	SetOntouchend(ontouchend func(Event))

	// Ontouchmove
	Ontouchmove() (ontouchmove func(Event))

	// Ontouchmove
	SetOntouchmove(ontouchmove func(Event))

	// Ontouchstart
	Ontouchstart() (ontouchstart func(Event))

	// Ontouchstart
	SetOntouchstart(ontouchstart func(Event))

	// Onvolumechange Occurs when the volume is changed, or playback is muted or unmuted.
	//     * @param ev The event.
	Onvolumechange() (onvolumechange func(Event))

	// Onvolumechange Occurs when the volume is changed, or playback is muted or unmuted.
	//     * @param ev The event.
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting Occurs when playback stops because the next frame of a video resource is not available.
	//     * @param ev The event.
	Onwaiting() (onwaiting func(Event))

	// Onwaiting Occurs when playback stops because the next frame of a video resource is not available.
	//     * @param ev The event.
	SetOnwaiting(onwaiting func(Event))

	// Onwebkitfullscreenchange
	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenchange
	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenerror
	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// Onwebkitfullscreenerror
	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// Plugins
	Plugins() (plugins HTMLCollection)

	// PointerLockElement
	PointerLockElement() (pointerLockElement Element)

	// ReadyState Retrieves a value that indicates the current state of the object.
	ReadyState() (readyState string)

	// Referrer Gets the URL of the location that referred the user to the current page.
	Referrer() (referrer string)

	// RootElement Gets the root svg element in the document hierarchy.
	RootElement() (rootElement *SVGSVGElement)

	// Scripts Retrieves a collection of all script objects in the document.
	Scripts() (scripts HTMLCollection)

	// ScrollingElement
	ScrollingElement() (scrollingElement Element)

	// StyleSheets Retrieves a collection of styleSheet objects representing the style sheets that correspond to each instance of a link or style object in the document.
	StyleSheets() (styleSheets *StyleSheetList)

	// Title Contains the title of the document.
	Title() (title string)

	// Title Contains the title of the document.
	SetTitle(title string)

	// URL Sets or gets the URL for the current document.
	URL() (URL string)

	// URLUnencoded Gets the URL for the document, stripped of any character encoding.
	URLUnencoded() (URLUnencoded string)

	// VisibilityState
	VisibilityState() (visibilityState *visibilitystate.VisibilityState)

	// VlinkColor Sets or gets the color of the links that the user has visited.
	VlinkColor() (vlinkColor string)

	// VlinkColor Sets or gets the color of the links that the user has visited.
	SetVlinkColor(vlinkColor string)

	// WebkitCurrentFullScreenElement
	WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement Element)

	// WebkitFullscreenElement
	WebkitFullscreenElement() (webkitFullscreenElement Element)

	// WebkitFullscreenEnabled
	WebkitFullscreenEnabled() (webkitFullscreenEnabled bool)

	// WebkitIsFullScreen
	WebkitIsFullScreen() (webkitIsFullScreen bool)

	// XMLEncoding
	XMLEncoding() (xmlEncoding string)

	// XMLStandalone
	XMLStandalone() (xmlStandalone bool)

	// XMLStandalone
	SetXMLStandalone(xmlStandalone bool)

	// XMLVersion Gets or sets the version attribute specified in the declaration of an XML document.
	XMLVersion() (xmlVersion string)

	// XMLVersion Gets or sets the version attribute specified in the declaration of an XML document.
	SetXMLVersion(xmlVersion string)
}
