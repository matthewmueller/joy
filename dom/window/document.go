package window

import (
	"github.com/matthewmueller/joy/dom/htmlallcollection"
	"github.com/matthewmueller/joy/dom/location"
	"github.com/matthewmueller/joy/dom/visibilitystate"
	"github.com/matthewmueller/joy/dom/xpathnsresolver"
)

// Document interface
// js:"Document"
type Document interface {
	Node

	// QuerySelector
	// js:"querySelector"
	// jsrewrite:"$_.querySelector($1)"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll"
	// jsrewrite:"$_.querySelectorAll($1)"
	QuerySelectorAll(selectors string) (n *NodeList)

	// CreateEvent
	// js:"createEvent"
	// jsrewrite:"$_.createEvent($1)"
	CreateEvent(eventInterface string) (e Event)

	// AdoptNode
	// js:"adoptNode"
	// jsrewrite:"$_.adoptNode($1)"
	AdoptNode(source Node) (n Node)

	// CaptureEvents
	// js:"captureEvents"
	// jsrewrite:"$_.captureEvents()"
	CaptureEvents()

	// CaretRangeFromPoint
	// js:"caretRangeFromPoint"
	// jsrewrite:"$_.caretRangeFromPoint($1, $2)"
	CaretRangeFromPoint(x float32, y float32) (r *Range)

	// Clear
	// js:"clear"
	// jsrewrite:"$_.clear()"
	Clear()

	// Close Closes an output stream and forces the sent data to display.
	// js:"close"
	// jsrewrite:"$_.close()"
	Close()

	// CreateAttribute Creates an attribute object with a specified name.
	//     * @param name String that sets the attribute object's name.
	// js:"createAttribute"
	// jsrewrite:"$_.createAttribute($1)"
	CreateAttribute(name string) (a *Attr)

	// CreateAttributeNS
	// js:"createAttributeNS"
	// jsrewrite:"$_.createAttributeNS($1, $2)"
	CreateAttributeNS(namespaceURI string, qualifiedName string) (a *Attr)

	// CreateCDATASection
	// js:"createCDATASection"
	// jsrewrite:"$_.createCDATASection($1)"
	CreateCDATASection(data string) (c *CDATASection)

	// CreateComment Creates a comment object with the specified data.
	//     * @param data Sets the comment object's data.
	// js:"createComment"
	// jsrewrite:"$_.createComment($1)"
	CreateComment(data string) (c *Comment)

	// CreateDocumentFragment Creates a new document.
	// js:"createDocumentFragment"
	// jsrewrite:"$_.createDocumentFragment()"
	CreateDocumentFragment() (d *DocumentFragment)

	// CreateElement Creates an instance of the element for the specified tag.
	//     * @param tagName The name of an element.
	// js:"createElement"
	// jsrewrite:"$_.createElement($1)"
	CreateElement(tagName string) (e Element)

	// CreateElementNS
	// js:"createElementNS"
	// jsrewrite:"$_.createElementNS($1, $2)"
	CreateElementNS(namespaceURI string, qualifiedName string) (e Element)

	// CreateExpression
	// js:"createExpression"
	// jsrewrite:"$_.createExpression($1, $2)"
	CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (x *XPathExpression)

	// CreateNodeIterator Creates a NodeIterator object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list
	//     * @param filter A custom NodeFilter function to use. For more information, see filter. Use null for no filter.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	// js:"createNodeIterator"
	// jsrewrite:"$_.createNodeIterator($1, $2, $3, $4)"
	CreateNodeIterator(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (n *NodeIterator)

	// CreateNSResolver
	// js:"createNSResolver"
	// jsrewrite:"$_.createNSResolver($1)"
	CreateNSResolver(nodeResolver Node) (x *xpathnsresolver.XPathNSResolver)

	// CreateProcessingInstruction
	// js:"createProcessingInstruction"
	// jsrewrite:"$_.createProcessingInstruction($1, $2)"
	CreateProcessingInstruction(target string, data string) (p *ProcessingInstruction)

	// CreateRange Returns an empty range object that has both of its boundary points positioned at the beginning of the document.
	// js:"createRange"
	// jsrewrite:"$_.createRange()"
	CreateRange() (r *Range)

	// CreateTextNode Creates a text string from the specified value.
	//     * @param data String that specifies the nodeValue property of the text node.
	// js:"createTextNode"
	// jsrewrite:"$_.createTextNode($1)"
	CreateTextNode(data string) (t Text)

	// CreateTouch
	// js:"createTouch"
	// jsrewrite:"$_.createTouch($1, $2, $3, $4, $5, $6, $7)"
	CreateTouch(view *Window, target EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *Touch)

	// CreateTouchList
	// js:"createTouchList"
	// jsrewrite:"$_.createTouchList($1)"
	CreateTouchList(touches *Touch) (t *TouchList)

	// CreateTreeWalker Creates a TreeWalker object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list. For more information, see whatToShow.
	//     * @param filter A custom NodeFilter function to use.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	// js:"createTreeWalker"
	// jsrewrite:"$_.createTreeWalker($1, $2, $3, $4)"
	CreateTreeWalker(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (t *TreeWalker)

	// ElementFromPoint Returns the element for the specified x coordinate and the specified y coordinate.
	//     * @param x The x-offset
	//     * @param y The y-offset
	// js:"elementFromPoint"
	// jsrewrite:"$_.elementFromPoint($1, $2)"
	ElementFromPoint(x int, y int) (e Element)

	// Evaluate
	// js:"evaluate"
	// jsrewrite:"$_.evaluate($1, $2, $3, $4, $5)"
	Evaluate(expression string, contextNode Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *XPathResult) (x *XPathResult)

	// ExecCommand Executes a command on the current document, current selection, or the given range.
	//     * @param commandId String that specifies the command to execute. This command can be any of the command identifiers that can be executed in script.
	//     * @param showUI Display the user interface, defaults to false.
	//     * @param value Value to assign.
	// js:"execCommand"
	// jsrewrite:"$_.execCommand($1, $2, $3)"
	ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool)

	// ExecCommandShowHelp Displays help information for the given command identifier.
	//     * @param commandId Displays help information for the given command identifier.
	// js:"execCommandShowHelp"
	// jsrewrite:"$_.execCommandShowHelp($1)"
	ExecCommandShowHelp(commandId string) (b bool)

	// ExitFullscreen
	// js:"exitFullscreen"
	// jsrewrite:"$_.exitFullscreen()"
	ExitFullscreen()

	// ExitPointerLock
	// js:"exitPointerLock"
	// jsrewrite:"$_.exitPointerLock()"
	ExitPointerLock()

	// Focus Causes the element to receive the focus and executes the code specified by the onfocus event.
	// js:"focus"
	// jsrewrite:"$_.focus()"
	Focus()

	// GetElementByID Returns a reference to the first object with the specified value of the ID or NAME attribute.
	//     * @param elementId String that specifies the ID value. Case-insensitive.
	// js:"getElementById"
	// jsrewrite:"$_.getElementById($1)"
	GetElementByID(elementId string) (e Element)

	// GetElementsByClassName
	// js:"getElementsByClassName"
	// jsrewrite:"$_.getElementsByClassName($1)"
	GetElementsByClassName(classNames string) (n *NodeList)

	// GetElementsByName Gets a collection of objects based on the value of the NAME or ID attribute.
	//     * @param elementName Gets a collection of objects based on the value of the NAME or ID attribute.
	// js:"getElementsByName"
	// jsrewrite:"$_.getElementsByName($1)"
	GetElementsByName(elementName string) (n *NodeList)

	// GetElementsByTagName Retrieves a collection of objects based on the specified element name.
	//     * @param name Specifies the name of an element.
	// js:"getElementsByTagName"
	// jsrewrite:"$_.getElementsByTagName($1)"
	GetElementsByTagName(tagname string) (n *NodeList)

	// GetElementsByTagNameNS
	// js:"getElementsByTagNameNS"
	// jsrewrite:"$_.getElementsByTagNameNS($1, $2)"
	GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// GetSelection Returns an object representing the current selection of the document that is loaded into the object displaying a webpage.
	// js:"getSelection"
	// jsrewrite:"$_.getSelection()"
	GetSelection() (s *Selection)

	// HasFocus Gets a value indicating whether the object currently has focus.
	// js:"hasFocus"
	// jsrewrite:"$_.hasFocus()"
	HasFocus() (b bool)

	// ImportNode
	// js:"importNode"
	// jsrewrite:"$_.importNode($1, $2)"
	ImportNode(importedNode Node, deep bool) (n Node)

	// MsElementsFromPoint
	// js:"msElementsFromPoint"
	// jsrewrite:"$_.msElementsFromPoint($1, $2)"
	MsElementsFromPoint(x float32, y float32) (n *NodeList)

	// MsElementsFromRect
	// js:"msElementsFromRect"
	// jsrewrite:"$_.msElementsFromRect($1, $2, $3, $4)"
	MsElementsFromRect(left float32, top float32, width float32, height float32) (n *NodeList)

	// Open Opens a new window and loads a document specified by a given URL. Also, opens a new window that uses the url parameter and the name parameter to collect the output of the write method and the writeln method.
	//     * @param url Specifies a MIME type for the document.
	//     * @param name Specifies the name of the window. This name is used as the value for the TARGET attribute on a form or an anchor element.
	//     * @param features Contains a list of items separated by commas. Each item consists of an option and a value, separated by an equals sign (for example, "fullscreen=yes, toolbar=yes"). The following values are supported.
	//     * @param replace Specifies whether the existing entry for the document is replaced in the history list.
	// js:"open"
	// jsrewrite:"$_.open($1, $2, $3, $4)"
	Open(url *string, name *string, features *string, replace *bool) (i interface{})

	// QueryCommandEnabled Returns a Boolean value that indicates whether a specified command can be successfully executed using execCommand, given the current state of the document.
	//     * @param commandId Specifies a command identifier.
	// js:"queryCommandEnabled"
	// jsrewrite:"$_.queryCommandEnabled($1)"
	QueryCommandEnabled(commandId string) (b bool)

	// QueryCommandIndeterm Returns a Boolean value that indicates whether the specified command is in the indeterminate state.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandIndeterm"
	// jsrewrite:"$_.queryCommandIndeterm($1)"
	QueryCommandIndeterm(commandId string) (b bool)

	// QueryCommandState Returns a Boolean value that indicates the current state of the command.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandState"
	// jsrewrite:"$_.queryCommandState($1)"
	QueryCommandState(commandId string) (b bool)

	// QueryCommandSupported Returns a Boolean value that indicates whether the current command is supported on the current range.
	//     * @param commandId Specifies a command identifier.
	// js:"queryCommandSupported"
	// jsrewrite:"$_.queryCommandSupported($1)"
	QueryCommandSupported(commandId string) (b bool)

	// QueryCommandText Retrieves the string associated with a command.
	//     * @param commandId String that contains the identifier of a command. This can be any command identifier given in the list of Command Identifiers.
	// js:"queryCommandText"
	// jsrewrite:"$_.queryCommandText($1)"
	QueryCommandText(commandId string) (s string)

	// QueryCommandValue Returns the current value of the document, range, or current selection for the given command.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandValue"
	// jsrewrite:"$_.queryCommandValue($1)"
	QueryCommandValue(commandId string) (s string)

	// ReleaseEvents
	// js:"releaseEvents"
	// jsrewrite:"$_.releaseEvents()"
	ReleaseEvents()

	// UpdateSettings Allows updating the print settings for the page.
	// js:"updateSettings"
	// jsrewrite:"$_.updateSettings()"
	UpdateSettings()

	// WebkitCancelFullScreen
	// js:"webkitCancelFullScreen"
	// jsrewrite:"$_.webkitCancelFullScreen()"
	WebkitCancelFullScreen()

	// WebkitExitFullscreen
	// js:"webkitExitFullscreen"
	// jsrewrite:"$_.webkitExitFullscreen()"
	WebkitExitFullscreen()

	// Write Writes one or more HTML expressions to a document in the specified window.
	//     * @param content Specifies the text and HTML tags to write.
	// js:"write"
	// jsrewrite:"$_.write($1)"
	Write(content string)

	// Writeln Writes one or more HTML expressions, followed by a carriage return, to a document in the specified window.
	//     * @param content The text and HTML tags to write.
	// js:"writeln"
	// jsrewrite:"$_.writeln($1)"
	Writeln(content string)

	// Onpointercancel prop
	// js:"onpointercancel"
	// jsrewrite:"$_.onpointercancel"
	Onpointercancel() (onpointercancel func(Event))

	// SetOnpointercancel prop
	// js:"onpointercancel"
	// jsrewrite:"$_.onpointercancel = $1"
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown prop
	// js:"onpointerdown"
	// jsrewrite:"$_.onpointerdown"
	Onpointerdown() (onpointerdown func(Event))

	// SetOnpointerdown prop
	// js:"onpointerdown"
	// jsrewrite:"$_.onpointerdown = $1"
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter prop
	// js:"onpointerenter"
	// jsrewrite:"$_.onpointerenter"
	Onpointerenter() (onpointerenter func(Event))

	// SetOnpointerenter prop
	// js:"onpointerenter"
	// jsrewrite:"$_.onpointerenter = $1"
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave prop
	// js:"onpointerleave"
	// jsrewrite:"$_.onpointerleave"
	Onpointerleave() (onpointerleave func(Event))

	// SetOnpointerleave prop
	// js:"onpointerleave"
	// jsrewrite:"$_.onpointerleave = $1"
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove prop
	// js:"onpointermove"
	// jsrewrite:"$_.onpointermove"
	Onpointermove() (onpointermove func(Event))

	// SetOnpointermove prop
	// js:"onpointermove"
	// jsrewrite:"$_.onpointermove = $1"
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout prop
	// js:"onpointerout"
	// jsrewrite:"$_.onpointerout"
	Onpointerout() (onpointerout func(Event))

	// SetOnpointerout prop
	// js:"onpointerout"
	// jsrewrite:"$_.onpointerout = $1"
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover prop
	// js:"onpointerover"
	// jsrewrite:"$_.onpointerover"
	Onpointerover() (onpointerover func(Event))

	// SetOnpointerover prop
	// js:"onpointerover"
	// jsrewrite:"$_.onpointerover = $1"
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup prop
	// js:"onpointerup"
	// jsrewrite:"$_.onpointerup"
	Onpointerup() (onpointerup func(Event))

	// SetOnpointerup prop
	// js:"onpointerup"
	// jsrewrite:"$_.onpointerup = $1"
	SetOnpointerup(onpointerup func(Event))

	// Onwheel prop
	// js:"onwheel"
	// jsrewrite:"$_.onwheel"
	Onwheel() (onwheel func(Event))

	// SetOnwheel prop
	// js:"onwheel"
	// jsrewrite:"$_.onwheel = $1"
	SetOnwheel(onwheel func(Event))

	// ActiveElement prop Gets the object that has the focus when the parent document has focus.
	// js:"activeElement"
	// jsrewrite:"$_.activeElement"
	ActiveElement() (activeElement Element)

	// AlinkColor prop Sets or gets the color of all active links in the document.
	// js:"alinkColor"
	// jsrewrite:"$_.alinkColor"
	AlinkColor() (alinkColor string)

	// SetAlinkColor prop Sets or gets the color of all active links in the document.
	// js:"alinkColor"
	// jsrewrite:"$_.alinkColor = $1"
	SetAlinkColor(alinkColor string)

	// All prop Returns a reference to the collection of elements contained by the object.
	// js:"all"
	// jsrewrite:"$_.all"
	All() (all *htmlallcollection.HTMLAllCollection)

	// Anchors prop Retrieves a collection of all a objects that have a name and/or id property. Objects in this collection are in HTML source order.
	// js:"anchors"
	// jsrewrite:"$_.anchors"
	Anchors() (anchors HTMLCollection)

	// Applets prop Retrieves a collection of all applet objects in the document.
	// js:"applets"
	// jsrewrite:"$_.applets"
	Applets() (applets HTMLCollection)

	// BgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	// js:"bgColor"
	// jsrewrite:"$_.bgColor"
	BgColor() (bgColor string)

	// SetBgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	// js:"bgColor"
	// jsrewrite:"$_.bgColor = $1"
	SetBgColor(bgColor string)

	// Body prop Specifies the beginning and end of the document body.
	// js:"body"
	// jsrewrite:"$_.body"
	Body() (body HTMLElement)

	// SetBody prop Specifies the beginning and end of the document body.
	// js:"body"
	// jsrewrite:"$_.body = $1"
	SetBody(body HTMLElement)

	// CharacterSet prop
	// js:"characterSet"
	// jsrewrite:"$_.characterSet"
	CharacterSet() (characterSet string)

	// Charset prop Gets or sets the character set used to encode the object.
	// js:"charset"
	// jsrewrite:"$_.charset"
	Charset() (charset string)

	// SetCharset prop Gets or sets the character set used to encode the object.
	// js:"charset"
	// jsrewrite:"$_.charset = $1"
	SetCharset(charset string)

	// CompatMode prop Gets a value that indicates whether standards-compliant mode is switched on for the object.
	// js:"compatMode"
	// jsrewrite:"$_.compatMode"
	CompatMode() (compatMode string)

	// Cookie prop
	// js:"cookie"
	// jsrewrite:"$_.cookie"
	Cookie() (cookie string)

	// SetCookie prop
	// js:"cookie"
	// jsrewrite:"$_.cookie = $1"
	SetCookie(cookie string)

	// CurrentScript prop
	// js:"currentScript"
	// jsrewrite:"$_.currentScript"
	CurrentScript() (currentScript interface{})

	// DefaultView prop
	// js:"defaultView"
	// jsrewrite:"$_.defaultView"
	DefaultView() (defaultView *Window)

	// DesignMode prop Sets or gets a value that indicates whether the document can be edited.
	// js:"designMode"
	// jsrewrite:"$_.designMode"
	DesignMode() (designMode string)

	// SetDesignMode prop Sets or gets a value that indicates whether the document can be edited.
	// js:"designMode"
	// jsrewrite:"$_.designMode = $1"
	SetDesignMode(designMode string)

	// Dir prop Sets or retrieves a value that indicates the reading order of the object.
	// js:"dir"
	// jsrewrite:"$_.dir"
	Dir() (dir string)

	// SetDir prop Sets or retrieves a value that indicates the reading order of the object.
	// js:"dir"
	// jsrewrite:"$_.dir = $1"
	SetDir(dir string)

	// Doctype prop Gets an object representing the document type declaration associated with the current document.
	// js:"doctype"
	// jsrewrite:"$_.doctype"
	Doctype() (doctype *DocumentType)

	// DocumentElement prop Gets a reference to the root node of the document.
	// js:"documentElement"
	// jsrewrite:"$_.documentElement"
	DocumentElement() (documentElement Element)

	// Domain prop Sets or gets the security domain of the document.
	// js:"domain"
	// jsrewrite:"$_.domain"
	Domain() (domain string)

	// SetDomain prop Sets or gets the security domain of the document.
	// js:"domain"
	// jsrewrite:"$_.domain = $1"
	SetDomain(domain string)

	// Embeds prop Retrieves a collection of all embed objects in the document.
	// js:"embeds"
	// jsrewrite:"$_.embeds"
	Embeds() (embeds HTMLCollection)

	// FgColor prop Sets or gets the foreground (text) color of the document.
	// js:"fgColor"
	// jsrewrite:"$_.fgColor"
	FgColor() (fgColor string)

	// SetFgColor prop Sets or gets the foreground (text) color of the document.
	// js:"fgColor"
	// jsrewrite:"$_.fgColor = $1"
	SetFgColor(fgColor string)

	// Forms prop Retrieves a collection, in source order, of all form objects in the document.
	// js:"forms"
	// jsrewrite:"$_.forms"
	Forms() (forms HTMLCollection)

	// FullscreenElement prop
	// js:"fullscreenElement"
	// jsrewrite:"$_.fullscreenElement"
	FullscreenElement() (fullscreenElement Element)

	// FullscreenEnabled prop
	// js:"fullscreenEnabled"
	// jsrewrite:"$_.fullscreenEnabled"
	FullscreenEnabled() (fullscreenEnabled bool)

	// Head prop
	// js:"head"
	// jsrewrite:"$_.head"
	Head() (head *HTMLHeadElement)

	// Hidden prop
	// js:"hidden"
	// jsrewrite:"$_.hidden"
	Hidden() (hidden bool)

	// Images prop Retrieves a collection, in source order, of img objects in the document.
	// js:"images"
	// jsrewrite:"$_.images"
	Images() (images HTMLCollection)

	// Implementation prop Gets the implementation object of the current document.
	// js:"implementation"
	// jsrewrite:"$_.implementation"
	Implementation() (implementation *DOMImplementation)

	// InputEncoding prop Returns the character encoding used to create the webpage that is loaded into the document object.
	// js:"inputEncoding"
	// jsrewrite:"$_.inputEncoding"
	InputEncoding() (inputEncoding string)

	// LastModified prop Gets the date that the page was last modified, if the page supplies one.
	// js:"lastModified"
	// jsrewrite:"$_.lastModified"
	LastModified() (lastModified string)

	// LinkColor prop Sets or gets the color of the document links.
	// js:"linkColor"
	// jsrewrite:"$_.linkColor"
	LinkColor() (linkColor string)

	// SetLinkColor prop Sets or gets the color of the document links.
	// js:"linkColor"
	// jsrewrite:"$_.linkColor = $1"
	SetLinkColor(linkColor string)

	// Links prop Retrieves a collection of all a objects that specify the href property and all area objects in the document.
	// js:"links"
	// jsrewrite:"$_.links"
	Links() (links HTMLCollection)

	// Location prop Contains information about the current URL.
	// js:"location"
	// jsrewrite:"$_.location"
	Location() (location *location.Location)

	// MsCapsLockWarningOff prop
	// js:"msCapsLockWarningOff"
	// jsrewrite:"$_.msCapsLockWarningOff"
	MsCapsLockWarningOff() (msCapsLockWarningOff bool)

	// SetMsCapsLockWarningOff prop
	// js:"msCapsLockWarningOff"
	// jsrewrite:"$_.msCapsLockWarningOff = $1"
	SetMsCapsLockWarningOff(msCapsLockWarningOff bool)

	// MsCSSOMElementFloatMetrics prop
	// js:"msCSSOMElementFloatMetrics"
	// jsrewrite:"$_.msCSSOMElementFloatMetrics"
	MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool)

	// SetMsCSSOMElementFloatMetrics prop
	// js:"msCSSOMElementFloatMetrics"
	// jsrewrite:"$_.msCSSOMElementFloatMetrics = $1"
	SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool)

	// Onabort prop Fires when the user aborts the download.
	//     * @param ev The event.
	// js:"onabort"
	// jsrewrite:"$_.onabort"
	Onabort() (onabort func(Event))

	// SetOnabort prop Fires when the user aborts the download.
	//     * @param ev The event.
	// js:"onabort"
	// jsrewrite:"$_.onabort = $1"
	SetOnabort(onabort func(Event))

	// Onactivate prop Fires when the object is set as the active element.
	//     * @param ev The event.
	// js:"onactivate"
	// jsrewrite:"$_.onactivate"
	Onactivate() (onactivate func(Event))

	// SetOnactivate prop Fires when the object is set as the active element.
	//     * @param ev The event.
	// js:"onactivate"
	// jsrewrite:"$_.onactivate = $1"
	SetOnactivate(onactivate func(Event))

	// Onbeforeactivate prop Fires immediately before the object is set as the active element.
	//     * @param ev The event.
	// js:"onbeforeactivate"
	// jsrewrite:"$_.onbeforeactivate"
	Onbeforeactivate() (onbeforeactivate func(Event))

	// SetOnbeforeactivate prop Fires immediately before the object is set as the active element.
	//     * @param ev The event.
	// js:"onbeforeactivate"
	// jsrewrite:"$_.onbeforeactivate = $1"
	SetOnbeforeactivate(onbeforeactivate func(Event))

	// Onbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The event.
	// js:"onbeforedeactivate"
	// jsrewrite:"$_.onbeforedeactivate"
	Onbeforedeactivate() (onbeforedeactivate func(Event))

	// SetOnbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The event.
	// js:"onbeforedeactivate"
	// jsrewrite:"$_.onbeforedeactivate = $1"
	SetOnbeforedeactivate(onbeforedeactivate func(Event))

	// Onblur prop Fires when the object loses the input focus.
	//     * @param ev The focus event.
	// js:"onblur"
	// jsrewrite:"$_.onblur"
	Onblur() (onblur func(Event))

	// SetOnblur prop Fires when the object loses the input focus.
	//     * @param ev The focus event.
	// js:"onblur"
	// jsrewrite:"$_.onblur = $1"
	SetOnblur(onblur func(Event))

	// Oncanplay prop Occurs when playback is possible, but would require further buffering.
	//     * @param ev The event.
	// js:"oncanplay"
	// jsrewrite:"$_.oncanplay"
	Oncanplay() (oncanplay func(Event))

	// SetOncanplay prop Occurs when playback is possible, but would require further buffering.
	//     * @param ev The event.
	// js:"oncanplay"
	// jsrewrite:"$_.oncanplay = $1"
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough prop
	// js:"oncanplaythrough"
	// jsrewrite:"$_.oncanplaythrough"
	Oncanplaythrough() (oncanplaythrough func(Event))

	// SetOncanplaythrough prop
	// js:"oncanplaythrough"
	// jsrewrite:"$_.oncanplaythrough = $1"
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange prop Fires when the contents of the object or selection have changed.
	//     * @param ev The event.
	// js:"onchange"
	// jsrewrite:"$_.onchange"
	Onchange() (onchange func(Event))

	// SetOnchange prop Fires when the contents of the object or selection have changed.
	//     * @param ev The event.
	// js:"onchange"
	// jsrewrite:"$_.onchange = $1"
	SetOnchange(onchange func(Event))

	// Onclick prop Fires when the user clicks the left mouse button on the object
	//     * @param ev The mouse event.
	// js:"onclick"
	// jsrewrite:"$_.onclick"
	Onclick() (onclick func(Event))

	// SetOnclick prop Fires when the user clicks the left mouse button on the object
	//     * @param ev The mouse event.
	// js:"onclick"
	// jsrewrite:"$_.onclick = $1"
	SetOnclick(onclick func(Event))

	// Oncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
	//     * @param ev The mouse event.
	// js:"oncontextmenu"
	// jsrewrite:"$_.oncontextmenu"
	Oncontextmenu() (oncontextmenu func(Event))

	// SetOncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
	//     * @param ev The mouse event.
	// js:"oncontextmenu"
	// jsrewrite:"$_.oncontextmenu = $1"
	SetOncontextmenu(oncontextmenu func(Event))

	// Ondblclick prop Fires when the user double-clicks the object.
	//     * @param ev The mouse event.
	// js:"ondblclick"
	// jsrewrite:"$_.ondblclick"
	Ondblclick() (ondblclick func(Event))

	// SetOndblclick prop Fires when the user double-clicks the object.
	//     * @param ev The mouse event.
	// js:"ondblclick"
	// jsrewrite:"$_.ondblclick = $1"
	SetOndblclick(ondblclick func(Event))

	// Ondeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The UI Event
	// js:"ondeactivate"
	// jsrewrite:"$_.ondeactivate"
	Ondeactivate() (ondeactivate func(Event))

	// SetOndeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
	//     * @param ev The UI Event
	// js:"ondeactivate"
	// jsrewrite:"$_.ondeactivate = $1"
	SetOndeactivate(ondeactivate func(Event))

	// Ondrag prop Fires on the source object continuously during a drag operation.
	//     * @param ev The event.
	// js:"ondrag"
	// jsrewrite:"$_.ondrag"
	Ondrag() (ondrag func(Event))

	// SetOndrag prop Fires on the source object continuously during a drag operation.
	//     * @param ev The event.
	// js:"ondrag"
	// jsrewrite:"$_.ondrag = $1"
	SetOndrag(ondrag func(Event))

	// Ondragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
	//     * @param ev The event.
	// js:"ondragend"
	// jsrewrite:"$_.ondragend"
	Ondragend() (ondragend func(Event))

	// SetOndragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
	//     * @param ev The event.
	// js:"ondragend"
	// jsrewrite:"$_.ondragend = $1"
	SetOndragend(ondragend func(Event))

	// Ondragenter prop Fires on the target element when the user drags the object to a valid drop target.
	//     * @param ev The drag event.
	// js:"ondragenter"
	// jsrewrite:"$_.ondragenter"
	Ondragenter() (ondragenter func(Event))

	// SetOndragenter prop Fires on the target element when the user drags the object to a valid drop target.
	//     * @param ev The drag event.
	// js:"ondragenter"
	// jsrewrite:"$_.ondragenter = $1"
	SetOndragenter(ondragenter func(Event))

	// Ondragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	//     * @param ev The drag event.
	// js:"ondragleave"
	// jsrewrite:"$_.ondragleave"
	Ondragleave() (ondragleave func(Event))

	// SetOndragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	//     * @param ev The drag event.
	// js:"ondragleave"
	// jsrewrite:"$_.ondragleave = $1"
	SetOndragleave(ondragleave func(Event))

	// Ondragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
	//     * @param ev The event.
	// js:"ondragover"
	// jsrewrite:"$_.ondragover"
	Ondragover() (ondragover func(Event))

	// SetOndragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
	//     * @param ev The event.
	// js:"ondragover"
	// jsrewrite:"$_.ondragover = $1"
	SetOndragover(ondragover func(Event))

	// Ondragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
	//     * @param ev The event.
	// js:"ondragstart"
	// jsrewrite:"$_.ondragstart"
	Ondragstart() (ondragstart func(Event))

	// SetOndragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
	//     * @param ev The event.
	// js:"ondragstart"
	// jsrewrite:"$_.ondragstart = $1"
	SetOndragstart(ondragstart func(Event))

	// Ondrop prop
	// js:"ondrop"
	// jsrewrite:"$_.ondrop"
	Ondrop() (ondrop func(Event))

	// SetOndrop prop
	// js:"ondrop"
	// jsrewrite:"$_.ondrop = $1"
	SetOndrop(ondrop func(Event))

	// Ondurationchange prop Occurs when the duration attribute is updated.
	//     * @param ev The event.
	// js:"ondurationchange"
	// jsrewrite:"$_.ondurationchange"
	Ondurationchange() (ondurationchange func(Event))

	// SetOndurationchange prop Occurs when the duration attribute is updated.
	//     * @param ev The event.
	// js:"ondurationchange"
	// jsrewrite:"$_.ondurationchange = $1"
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied prop Occurs when the media element is reset to its initial state.
	//     * @param ev The event.
	// js:"onemptied"
	// jsrewrite:"$_.onemptied"
	Onemptied() (onemptied func(Event))

	// SetOnemptied prop Occurs when the media element is reset to its initial state.
	//     * @param ev The event.
	// js:"onemptied"
	// jsrewrite:"$_.onemptied = $1"
	SetOnemptied(onemptied func(Event))

	// Onended prop Occurs when the end of playback is reached.
	//     * @param ev The event
	// js:"onended"
	// jsrewrite:"$_.onended"
	Onended() (onended func(Event))

	// SetOnended prop Occurs when the end of playback is reached.
	//     * @param ev The event
	// js:"onended"
	// jsrewrite:"$_.onended = $1"
	SetOnended(onended func(Event))

	// Onerror prop Fires when an error occurs during object loading.
	//     * @param ev The event.
	// js:"onerror"
	// jsrewrite:"$_.onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop Fires when an error occurs during object loading.
	//     * @param ev The event.
	// js:"onerror"
	// jsrewrite:"$_.onerror = $1"
	SetOnerror(onerror func(Event))

	// Onfocus prop Fires when the object receives focus.
	//     * @param ev The event.
	// js:"onfocus"
	// jsrewrite:"$_.onfocus"
	Onfocus() (onfocus func(Event))

	// SetOnfocus prop Fires when the object receives focus.
	//     * @param ev The event.
	// js:"onfocus"
	// jsrewrite:"$_.onfocus = $1"
	SetOnfocus(onfocus func(Event))

	// Onfullscreenchange prop
	// js:"onfullscreenchange"
	// jsrewrite:"$_.onfullscreenchange"
	Onfullscreenchange() (onfullscreenchange func(Event))

	// SetOnfullscreenchange prop
	// js:"onfullscreenchange"
	// jsrewrite:"$_.onfullscreenchange = $1"
	SetOnfullscreenchange(onfullscreenchange func(Event))

	// Onfullscreenerror prop
	// js:"onfullscreenerror"
	// jsrewrite:"$_.onfullscreenerror"
	Onfullscreenerror() (onfullscreenerror func(Event))

	// SetOnfullscreenerror prop
	// js:"onfullscreenerror"
	// jsrewrite:"$_.onfullscreenerror = $1"
	SetOnfullscreenerror(onfullscreenerror func(Event))

	// Oninput prop
	// js:"oninput"
	// jsrewrite:"$_.oninput"
	Oninput() (oninput func(Event))

	// SetOninput prop
	// js:"oninput"
	// jsrewrite:"$_.oninput = $1"
	SetOninput(oninput func(Event))

	// Oninvalid prop
	// js:"oninvalid"
	// jsrewrite:"$_.oninvalid"
	Oninvalid() (oninvalid func(Event))

	// SetOninvalid prop
	// js:"oninvalid"
	// jsrewrite:"$_.oninvalid = $1"
	SetOninvalid(oninvalid func(Event))

	// Onkeydown prop Fires when the user presses a key.
	//     * @param ev The keyboard event
	// js:"onkeydown"
	// jsrewrite:"$_.onkeydown"
	Onkeydown() (onkeydown func(Event))

	// SetOnkeydown prop Fires when the user presses a key.
	//     * @param ev The keyboard event
	// js:"onkeydown"
	// jsrewrite:"$_.onkeydown = $1"
	SetOnkeydown(onkeydown func(Event))

	// Onkeypress prop Fires when the user presses an alphanumeric key.
	//     * @param ev The event.
	// js:"onkeypress"
	// jsrewrite:"$_.onkeypress"
	Onkeypress() (onkeypress func(Event))

	// SetOnkeypress prop Fires when the user presses an alphanumeric key.
	//     * @param ev The event.
	// js:"onkeypress"
	// jsrewrite:"$_.onkeypress = $1"
	SetOnkeypress(onkeypress func(Event))

	// Onkeyup prop Fires when the user releases a key.
	//     * @param ev The keyboard event
	// js:"onkeyup"
	// jsrewrite:"$_.onkeyup"
	Onkeyup() (onkeyup func(Event))

	// SetOnkeyup prop Fires when the user releases a key.
	//     * @param ev The keyboard event
	// js:"onkeyup"
	// jsrewrite:"$_.onkeyup = $1"
	SetOnkeyup(onkeyup func(Event))

	// Onload prop Fires immediately after the browser loads the object.
	//     * @param ev The event.
	// js:"onload"
	// jsrewrite:"$_.onload"
	Onload() (onload func(Event))

	// SetOnload prop Fires immediately after the browser loads the object.
	//     * @param ev The event.
	// js:"onload"
	// jsrewrite:"$_.onload = $1"
	SetOnload(onload func(Event))

	// Onloadeddata prop Occurs when media data is loaded at the current playback position.
	//     * @param ev The event.
	// js:"onloadeddata"
	// jsrewrite:"$_.onloadeddata"
	Onloadeddata() (onloadeddata func(Event))

	// SetOnloadeddata prop Occurs when media data is loaded at the current playback position.
	//     * @param ev The event.
	// js:"onloadeddata"
	// jsrewrite:"$_.onloadeddata = $1"
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
	//     * @param ev The event.
	// js:"onloadedmetadata"
	// jsrewrite:"$_.onloadedmetadata"
	Onloadedmetadata() (onloadedmetadata func(Event))

	// SetOnloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
	//     * @param ev The event.
	// js:"onloadedmetadata"
	// jsrewrite:"$_.onloadedmetadata = $1"
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart prop Occurs when Internet Explorer begins looking for media data.
	//     * @param ev The event.
	// js:"onloadstart"
	// jsrewrite:"$_.onloadstart"
	Onloadstart() (onloadstart func(Event))

	// SetOnloadstart prop Occurs when Internet Explorer begins looking for media data.
	//     * @param ev The event.
	// js:"onloadstart"
	// jsrewrite:"$_.onloadstart = $1"
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown prop Fires when the user clicks the object with either mouse button.
	//     * @param ev The mouse event.
	// js:"onmousedown"
	// jsrewrite:"$_.onmousedown"
	Onmousedown() (onmousedown func(Event))

	// SetOnmousedown prop Fires when the user clicks the object with either mouse button.
	//     * @param ev The mouse event.
	// js:"onmousedown"
	// jsrewrite:"$_.onmousedown = $1"
	SetOnmousedown(onmousedown func(Event))

	// Onmousemove prop Fires when the user moves the mouse over the object.
	//     * @param ev The mouse event.
	// js:"onmousemove"
	// jsrewrite:"$_.onmousemove"
	Onmousemove() (onmousemove func(Event))

	// SetOnmousemove prop Fires when the user moves the mouse over the object.
	//     * @param ev The mouse event.
	// js:"onmousemove"
	// jsrewrite:"$_.onmousemove = $1"
	SetOnmousemove(onmousemove func(Event))

	// Onmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
	//     * @param ev The mouse event.
	// js:"onmouseout"
	// jsrewrite:"$_.onmouseout"
	Onmouseout() (onmouseout func(Event))

	// SetOnmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
	//     * @param ev The mouse event.
	// js:"onmouseout"
	// jsrewrite:"$_.onmouseout = $1"
	SetOnmouseout(onmouseout func(Event))

	// Onmouseover prop Fires when the user moves the mouse pointer into the object.
	//     * @param ev The mouse event.
	// js:"onmouseover"
	// jsrewrite:"$_.onmouseover"
	Onmouseover() (onmouseover func(Event))

	// SetOnmouseover prop Fires when the user moves the mouse pointer into the object.
	//     * @param ev The mouse event.
	// js:"onmouseover"
	// jsrewrite:"$_.onmouseover = $1"
	SetOnmouseover(onmouseover func(Event))

	// Onmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
	//     * @param ev The mouse event.
	// js:"onmouseup"
	// jsrewrite:"$_.onmouseup"
	Onmouseup() (onmouseup func(Event))

	// SetOnmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
	//     * @param ev The mouse event.
	// js:"onmouseup"
	// jsrewrite:"$_.onmouseup = $1"
	SetOnmouseup(onmouseup func(Event))

	// Onmousewheel prop Fires when the wheel button is rotated.
	//     * @param ev The mouse event
	// js:"onmousewheel"
	// jsrewrite:"$_.onmousewheel"
	Onmousewheel() (onmousewheel func(Event))

	// SetOnmousewheel prop Fires when the wheel button is rotated.
	//     * @param ev The mouse event
	// js:"onmousewheel"
	// jsrewrite:"$_.onmousewheel = $1"
	SetOnmousewheel(onmousewheel func(Event))

	// Onmscontentzoom prop
	// js:"onmscontentzoom"
	// jsrewrite:"$_.onmscontentzoom"
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// SetOnmscontentzoom prop
	// js:"onmscontentzoom"
	// jsrewrite:"$_.onmscontentzoom = $1"
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsgesturechange prop
	// js:"onmsgesturechange"
	// jsrewrite:"$_.onmsgesturechange"
	Onmsgesturechange() (onmsgesturechange func(Event))

	// SetOnmsgesturechange prop
	// js:"onmsgesturechange"
	// jsrewrite:"$_.onmsgesturechange = $1"
	SetOnmsgesturechange(onmsgesturechange func(Event))

	// Onmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	// jsrewrite:"$_.onmsgesturedoubletap"
	Onmsgesturedoubletap() (onmsgesturedoubletap func(Event))

	// SetOnmsgesturedoubletap prop
	// js:"onmsgesturedoubletap"
	// jsrewrite:"$_.onmsgesturedoubletap = $1"
	SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event))

	// Onmsgestureend prop
	// js:"onmsgestureend"
	// jsrewrite:"$_.onmsgestureend"
	Onmsgestureend() (onmsgestureend func(Event))

	// SetOnmsgestureend prop
	// js:"onmsgestureend"
	// jsrewrite:"$_.onmsgestureend = $1"
	SetOnmsgestureend(onmsgestureend func(Event))

	// Onmsgesturehold prop
	// js:"onmsgesturehold"
	// jsrewrite:"$_.onmsgesturehold"
	Onmsgesturehold() (onmsgesturehold func(Event))

	// SetOnmsgesturehold prop
	// js:"onmsgesturehold"
	// jsrewrite:"$_.onmsgesturehold = $1"
	SetOnmsgesturehold(onmsgesturehold func(Event))

	// Onmsgesturestart prop
	// js:"onmsgesturestart"
	// jsrewrite:"$_.onmsgesturestart"
	Onmsgesturestart() (onmsgesturestart func(Event))

	// SetOnmsgesturestart prop
	// js:"onmsgesturestart"
	// jsrewrite:"$_.onmsgesturestart = $1"
	SetOnmsgesturestart(onmsgesturestart func(Event))

	// Onmsgesturetap prop
	// js:"onmsgesturetap"
	// jsrewrite:"$_.onmsgesturetap"
	Onmsgesturetap() (onmsgesturetap func(Event))

	// SetOnmsgesturetap prop
	// js:"onmsgesturetap"
	// jsrewrite:"$_.onmsgesturetap = $1"
	SetOnmsgesturetap(onmsgesturetap func(Event))

	// Onmsinertiastart prop
	// js:"onmsinertiastart"
	// jsrewrite:"$_.onmsinertiastart"
	Onmsinertiastart() (onmsinertiastart func(Event))

	// SetOnmsinertiastart prop
	// js:"onmsinertiastart"
	// jsrewrite:"$_.onmsinertiastart = $1"
	SetOnmsinertiastart(onmsinertiastart func(Event))

	// Onmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	// jsrewrite:"$_.onmsmanipulationstatechanged"
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// SetOnmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	// jsrewrite:"$_.onmsmanipulationstatechanged = $1"
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmspointercancel prop
	// js:"onmspointercancel"
	// jsrewrite:"$_.onmspointercancel"
	Onmspointercancel() (onmspointercancel func(Event))

	// SetOnmspointercancel prop
	// js:"onmspointercancel"
	// jsrewrite:"$_.onmspointercancel = $1"
	SetOnmspointercancel(onmspointercancel func(Event))

	// Onmspointerdown prop
	// js:"onmspointerdown"
	// jsrewrite:"$_.onmspointerdown"
	Onmspointerdown() (onmspointerdown func(Event))

	// SetOnmspointerdown prop
	// js:"onmspointerdown"
	// jsrewrite:"$_.onmspointerdown = $1"
	SetOnmspointerdown(onmspointerdown func(Event))

	// Onmspointerenter prop
	// js:"onmspointerenter"
	// jsrewrite:"$_.onmspointerenter"
	Onmspointerenter() (onmspointerenter func(Event))

	// SetOnmspointerenter prop
	// js:"onmspointerenter"
	// jsrewrite:"$_.onmspointerenter = $1"
	SetOnmspointerenter(onmspointerenter func(Event))

	// Onmspointerleave prop
	// js:"onmspointerleave"
	// jsrewrite:"$_.onmspointerleave"
	Onmspointerleave() (onmspointerleave func(Event))

	// SetOnmspointerleave prop
	// js:"onmspointerleave"
	// jsrewrite:"$_.onmspointerleave = $1"
	SetOnmspointerleave(onmspointerleave func(Event))

	// Onmspointermove prop
	// js:"onmspointermove"
	// jsrewrite:"$_.onmspointermove"
	Onmspointermove() (onmspointermove func(Event))

	// SetOnmspointermove prop
	// js:"onmspointermove"
	// jsrewrite:"$_.onmspointermove = $1"
	SetOnmspointermove(onmspointermove func(Event))

	// Onmspointerout prop
	// js:"onmspointerout"
	// jsrewrite:"$_.onmspointerout"
	Onmspointerout() (onmspointerout func(Event))

	// SetOnmspointerout prop
	// js:"onmspointerout"
	// jsrewrite:"$_.onmspointerout = $1"
	SetOnmspointerout(onmspointerout func(Event))

	// Onmspointerover prop
	// js:"onmspointerover"
	// jsrewrite:"$_.onmspointerover"
	Onmspointerover() (onmspointerover func(Event))

	// SetOnmspointerover prop
	// js:"onmspointerover"
	// jsrewrite:"$_.onmspointerover = $1"
	SetOnmspointerover(onmspointerover func(Event))

	// Onmspointerup prop
	// js:"onmspointerup"
	// jsrewrite:"$_.onmspointerup"
	Onmspointerup() (onmspointerup func(Event))

	// SetOnmspointerup prop
	// js:"onmspointerup"
	// jsrewrite:"$_.onmspointerup = $1"
	SetOnmspointerup(onmspointerup func(Event))

	// Onmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmssitemodejumplistitemremoved"
	// jsrewrite:"$_.onmssitemodejumplistitemremoved"
	Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// SetOnmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmssitemodejumplistitemremoved"
	// jsrewrite:"$_.onmssitemodejumplistitemremoved = $1"
	SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// Onmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmsthumbnailclick"
	// jsrewrite:"$_.onmsthumbnailclick"
	Onmsthumbnailclick() (onmsthumbnailclick func(*MSSiteModeEvent))

	// SetOnmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	//     * @param ev The event.
	// js:"onmsthumbnailclick"
	// jsrewrite:"$_.onmsthumbnailclick = $1"
	SetOnmsthumbnailclick(onmsthumbnailclick func(*MSSiteModeEvent))

	// Onpause prop Occurs when playback is paused.
	//     * @param ev The event.
	// js:"onpause"
	// jsrewrite:"$_.onpause"
	Onpause() (onpause func(Event))

	// SetOnpause prop Occurs when playback is paused.
	//     * @param ev The event.
	// js:"onpause"
	// jsrewrite:"$_.onpause = $1"
	SetOnpause(onpause func(Event))

	// Onplay prop Occurs when the play method is requested.
	//     * @param ev The event.
	// js:"onplay"
	// jsrewrite:"$_.onplay"
	Onplay() (onplay func(Event))

	// SetOnplay prop Occurs when the play method is requested.
	//     * @param ev The event.
	// js:"onplay"
	// jsrewrite:"$_.onplay = $1"
	SetOnplay(onplay func(Event))

	// Onplaying prop Occurs when the audio or video has started playing.
	//     * @param ev The event.
	// js:"onplaying"
	// jsrewrite:"$_.onplaying"
	Onplaying() (onplaying func(Event))

	// SetOnplaying prop Occurs when the audio or video has started playing.
	//     * @param ev The event.
	// js:"onplaying"
	// jsrewrite:"$_.onplaying = $1"
	SetOnplaying(onplaying func(Event))

	// Onpointerlockchange prop
	// js:"onpointerlockchange"
	// jsrewrite:"$_.onpointerlockchange"
	Onpointerlockchange() (onpointerlockchange func(Event))

	// SetOnpointerlockchange prop
	// js:"onpointerlockchange"
	// jsrewrite:"$_.onpointerlockchange = $1"
	SetOnpointerlockchange(onpointerlockchange func(Event))

	// Onpointerlockerror prop
	// js:"onpointerlockerror"
	// jsrewrite:"$_.onpointerlockerror"
	Onpointerlockerror() (onpointerlockerror func(Event))

	// SetOnpointerlockerror prop
	// js:"onpointerlockerror"
	// jsrewrite:"$_.onpointerlockerror = $1"
	SetOnpointerlockerror(onpointerlockerror func(Event))

	// Onprogress prop Occurs to indicate progress while downloading media data.
	//     * @param ev The event.
	// js:"onprogress"
	// jsrewrite:"$_.onprogress"
	Onprogress() (onprogress func(Event))

	// SetOnprogress prop Occurs to indicate progress while downloading media data.
	//     * @param ev The event.
	// js:"onprogress"
	// jsrewrite:"$_.onprogress = $1"
	SetOnprogress(onprogress func(Event))

	// Onratechange prop Occurs when the playback rate is increased or decreased.
	//     * @param ev The event.
	// js:"onratechange"
	// jsrewrite:"$_.onratechange"
	Onratechange() (onratechange func(Event))

	// SetOnratechange prop Occurs when the playback rate is increased or decreased.
	//     * @param ev The event.
	// js:"onratechange"
	// jsrewrite:"$_.onratechange = $1"
	SetOnratechange(onratechange func(Event))

	// Onreadystatechange prop Fires when the state of the object has changed.
	//     * @param ev The event
	// js:"onreadystatechange"
	// jsrewrite:"$_.onreadystatechange"
	Onreadystatechange() (onreadystatechange func(Event))

	// SetOnreadystatechange prop Fires when the state of the object has changed.
	//     * @param ev The event
	// js:"onreadystatechange"
	// jsrewrite:"$_.onreadystatechange = $1"
	SetOnreadystatechange(onreadystatechange func(Event))

	// Onreset prop Fires when the user resets a form.
	//     * @param ev The event.
	// js:"onreset"
	// jsrewrite:"$_.onreset"
	Onreset() (onreset func(Event))

	// SetOnreset prop Fires when the user resets a form.
	//     * @param ev The event.
	// js:"onreset"
	// jsrewrite:"$_.onreset = $1"
	SetOnreset(onreset func(Event))

	// Onscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
	//     * @param ev The event.
	// js:"onscroll"
	// jsrewrite:"$_.onscroll"
	Onscroll() (onscroll func(Event))

	// SetOnscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
	//     * @param ev The event.
	// js:"onscroll"
	// jsrewrite:"$_.onscroll = $1"
	SetOnscroll(onscroll func(Event))

	// Onseeked prop Occurs when the seek operation ends.
	//     * @param ev The event.
	// js:"onseeked"
	// jsrewrite:"$_.onseeked"
	Onseeked() (onseeked func(Event))

	// SetOnseeked prop Occurs when the seek operation ends.
	//     * @param ev The event.
	// js:"onseeked"
	// jsrewrite:"$_.onseeked = $1"
	SetOnseeked(onseeked func(Event))

	// Onseeking prop Occurs when the current playback position is moved.
	//     * @param ev The event.
	// js:"onseeking"
	// jsrewrite:"$_.onseeking"
	Onseeking() (onseeking func(Event))

	// SetOnseeking prop Occurs when the current playback position is moved.
	//     * @param ev The event.
	// js:"onseeking"
	// jsrewrite:"$_.onseeking = $1"
	SetOnseeking(onseeking func(Event))

	// Onselect prop Fires when the current selection changes.
	//     * @param ev The event.
	// js:"onselect"
	// jsrewrite:"$_.onselect"
	Onselect() (onselect func(Event))

	// SetOnselect prop Fires when the current selection changes.
	//     * @param ev The event.
	// js:"onselect"
	// jsrewrite:"$_.onselect = $1"
	SetOnselect(onselect func(Event))

	// Onselectionchange prop Fires when the selection state of a document changes.
	//     * @param ev The event.
	// js:"onselectionchange"
	// jsrewrite:"$_.onselectionchange"
	Onselectionchange() (onselectionchange func(Event))

	// SetOnselectionchange prop Fires when the selection state of a document changes.
	//     * @param ev The event.
	// js:"onselectionchange"
	// jsrewrite:"$_.onselectionchange = $1"
	SetOnselectionchange(onselectionchange func(Event))

	// Onselectstart prop
	// js:"onselectstart"
	// jsrewrite:"$_.onselectstart"
	Onselectstart() (onselectstart func(Event))

	// SetOnselectstart prop
	// js:"onselectstart"
	// jsrewrite:"$_.onselectstart = $1"
	SetOnselectstart(onselectstart func(Event))

	// Onstalled prop Occurs when the download has stopped.
	//     * @param ev The event.
	// js:"onstalled"
	// jsrewrite:"$_.onstalled"
	Onstalled() (onstalled func(Event))

	// SetOnstalled prop Occurs when the download has stopped.
	//     * @param ev The event.
	// js:"onstalled"
	// jsrewrite:"$_.onstalled = $1"
	SetOnstalled(onstalled func(Event))

	// Onstop prop Fires when the user clicks the Stop button or leaves the Web page.
	//     * @param ev The event.
	// js:"onstop"
	// jsrewrite:"$_.onstop"
	Onstop() (onstop func(Event))

	// SetOnstop prop Fires when the user clicks the Stop button or leaves the Web page.
	//     * @param ev The event.
	// js:"onstop"
	// jsrewrite:"$_.onstop = $1"
	SetOnstop(onstop func(Event))

	// Onsubmit prop
	// js:"onsubmit"
	// jsrewrite:"$_.onsubmit"
	Onsubmit() (onsubmit func(Event))

	// SetOnsubmit prop
	// js:"onsubmit"
	// jsrewrite:"$_.onsubmit = $1"
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend prop Occurs if the load operation has been intentionally halted.
	//     * @param ev The event.
	// js:"onsuspend"
	// jsrewrite:"$_.onsuspend"
	Onsuspend() (onsuspend func(Event))

	// SetOnsuspend prop Occurs if the load operation has been intentionally halted.
	//     * @param ev The event.
	// js:"onsuspend"
	// jsrewrite:"$_.onsuspend = $1"
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate prop Occurs to indicate the current playback position.
	//     * @param ev The event.
	// js:"ontimeupdate"
	// jsrewrite:"$_.ontimeupdate"
	Ontimeupdate() (ontimeupdate func(Event))

	// SetOntimeupdate prop Occurs to indicate the current playback position.
	//     * @param ev The event.
	// js:"ontimeupdate"
	// jsrewrite:"$_.ontimeupdate = $1"
	SetOntimeupdate(ontimeupdate func(Event))

	// Ontouchcancel prop
	// js:"ontouchcancel"
	// jsrewrite:"$_.ontouchcancel"
	Ontouchcancel() (ontouchcancel func(Event))

	// SetOntouchcancel prop
	// js:"ontouchcancel"
	// jsrewrite:"$_.ontouchcancel = $1"
	SetOntouchcancel(ontouchcancel func(Event))

	// Ontouchend prop
	// js:"ontouchend"
	// jsrewrite:"$_.ontouchend"
	Ontouchend() (ontouchend func(Event))

	// SetOntouchend prop
	// js:"ontouchend"
	// jsrewrite:"$_.ontouchend = $1"
	SetOntouchend(ontouchend func(Event))

	// Ontouchmove prop
	// js:"ontouchmove"
	// jsrewrite:"$_.ontouchmove"
	Ontouchmove() (ontouchmove func(Event))

	// SetOntouchmove prop
	// js:"ontouchmove"
	// jsrewrite:"$_.ontouchmove = $1"
	SetOntouchmove(ontouchmove func(Event))

	// Ontouchstart prop
	// js:"ontouchstart"
	// jsrewrite:"$_.ontouchstart"
	Ontouchstart() (ontouchstart func(Event))

	// SetOntouchstart prop
	// js:"ontouchstart"
	// jsrewrite:"$_.ontouchstart = $1"
	SetOntouchstart(ontouchstart func(Event))

	// Onvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
	//     * @param ev The event.
	// js:"onvolumechange"
	// jsrewrite:"$_.onvolumechange"
	Onvolumechange() (onvolumechange func(Event))

	// SetOnvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
	//     * @param ev The event.
	// js:"onvolumechange"
	// jsrewrite:"$_.onvolumechange = $1"
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
	//     * @param ev The event.
	// js:"onwaiting"
	// jsrewrite:"$_.onwaiting"
	Onwaiting() (onwaiting func(Event))

	// SetOnwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
	//     * @param ev The event.
	// js:"onwaiting"
	// jsrewrite:"$_.onwaiting = $1"
	SetOnwaiting(onwaiting func(Event))

	// Onwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	// jsrewrite:"$_.onwebkitfullscreenchange"
	Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// SetOnwebkitfullscreenchange prop
	// js:"onwebkitfullscreenchange"
	// jsrewrite:"$_.onwebkitfullscreenchange = $1"
	SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// Onwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	// jsrewrite:"$_.onwebkitfullscreenerror"
	Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// SetOnwebkitfullscreenerror prop
	// js:"onwebkitfullscreenerror"
	// jsrewrite:"$_.onwebkitfullscreenerror = $1"
	SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// Plugins prop
	// js:"plugins"
	// jsrewrite:"$_.plugins"
	Plugins() (plugins HTMLCollection)

	// PointerLockElement prop
	// js:"pointerLockElement"
	// jsrewrite:"$_.pointerLockElement"
	PointerLockElement() (pointerLockElement Element)

	// ReadyState prop Retrieves a value that indicates the current state of the object.
	// js:"readyState"
	// jsrewrite:"$_.readyState"
	ReadyState() (readyState string)

	// Referrer prop Gets the URL of the location that referred the user to the current page.
	// js:"referrer"
	// jsrewrite:"$_.referrer"
	Referrer() (referrer string)

	// RootElement prop Gets the root svg element in the document hierarchy.
	// js:"rootElement"
	// jsrewrite:"$_.rootElement"
	RootElement() (rootElement *SVGSVGElement)

	// Scripts prop Retrieves a collection of all script objects in the document.
	// js:"scripts"
	// jsrewrite:"$_.scripts"
	Scripts() (scripts HTMLCollection)

	// ScrollingElement prop
	// js:"scrollingElement"
	// jsrewrite:"$_.scrollingElement"
	ScrollingElement() (scrollingElement Element)

	// StyleSheets prop Retrieves a collection of styleSheet objects representing the style sheets that correspond to each instance of a link or style object in the document.
	// js:"styleSheets"
	// jsrewrite:"$_.styleSheets"
	StyleSheets() (styleSheets *StyleSheetList)

	// Title prop Contains the title of the document.
	// js:"title"
	// jsrewrite:"$_.title"
	Title() (title string)

	// SetTitle prop Contains the title of the document.
	// js:"title"
	// jsrewrite:"$_.title = $1"
	SetTitle(title string)

	// URL prop Sets or gets the URL for the current document.
	// js:"URL"
	// jsrewrite:"$_.URL"
	URL() (URL string)

	// URLUnencoded prop Gets the URL for the document, stripped of any character encoding.
	// js:"URLUnencoded"
	// jsrewrite:"$_.URLUnencoded"
	URLUnencoded() (URLUnencoded string)

	// VisibilityState prop
	// js:"visibilityState"
	// jsrewrite:"$_.visibilityState"
	VisibilityState() (visibilityState *visibilitystate.VisibilityState)

	// VlinkColor prop Sets or gets the color of the links that the user has visited.
	// js:"vlinkColor"
	// jsrewrite:"$_.vlinkColor"
	VlinkColor() (vlinkColor string)

	// SetVlinkColor prop Sets or gets the color of the links that the user has visited.
	// js:"vlinkColor"
	// jsrewrite:"$_.vlinkColor = $1"
	SetVlinkColor(vlinkColor string)

	// WebkitCurrentFullScreenElement prop
	// js:"webkitCurrentFullScreenElement"
	// jsrewrite:"$_.webkitCurrentFullScreenElement"
	WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement Element)

	// WebkitFullscreenElement prop
	// js:"webkitFullscreenElement"
	// jsrewrite:"$_.webkitFullscreenElement"
	WebkitFullscreenElement() (webkitFullscreenElement Element)

	// WebkitFullscreenEnabled prop
	// js:"webkitFullscreenEnabled"
	// jsrewrite:"$_.webkitFullscreenEnabled"
	WebkitFullscreenEnabled() (webkitFullscreenEnabled bool)

	// WebkitIsFullScreen prop
	// js:"webkitIsFullScreen"
	// jsrewrite:"$_.webkitIsFullScreen"
	WebkitIsFullScreen() (webkitIsFullScreen bool)

	// XMLEncoding prop
	// js:"xmlEncoding"
	// jsrewrite:"$_.xmlEncoding"
	XMLEncoding() (xmlEncoding string)

	// XMLStandalone prop
	// js:"xmlStandalone"
	// jsrewrite:"$_.xmlStandalone"
	XMLStandalone() (xmlStandalone bool)

	// SetXMLStandalone prop
	// js:"xmlStandalone"
	// jsrewrite:"$_.xmlStandalone = $1"
	SetXMLStandalone(xmlStandalone bool)

	// XMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
	// js:"xmlVersion"
	// jsrewrite:"$_.xmlVersion"
	XMLVersion() (xmlVersion string)

	// SetXMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
	// js:"xmlVersion"
	// jsrewrite:"$_.xmlVersion = $1"
	SetXMLVersion(xmlVersion string)
}
