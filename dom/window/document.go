package window

import (
	"github.com/matthewmueller/golly/dom/htmlallcollection"
	"github.com/matthewmueller/golly/dom/xpathnsresolver"
	"github.com/matthewmueller/golly/js"
)

// js:"Document,omit"
type Document interface {
	Node

	// QuerySelector
	// js:"querySelector,rewrite=queryselector"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll,rewrite=queryselectorall"
	QuerySelectorAll(selectors string) (n *NodeList)

	// CreateEvent
	// js:"createEvent,rewrite=createevent"
	CreateEvent(eventInterface string) (e Event)

	// AdoptNode
	// js:"adoptNode,rewrite=adoptnode"
	AdoptNode(source Node) (n Node)

	// CaptureEvents
	// js:"captureEvents,rewrite=captureevents"
	CaptureEvents()

	// CaretRangeFromPoint
	// js:"caretRangeFromPoint,rewrite=caretrangefrompoint"
	CaretRangeFromPoint(x float32, y float32) (r *Range)

	// Clear
	// js:"clear,rewrite=clear"
	Clear()

	// Close Closes an output stream and forces the sent data to display.
	// js:"close,rewrite=cls"
	Close()

	// CreateAttribute Creates an attribute object with a specified name.
	//     * @param name String that sets the attribute object's name.
	// js:"createAttribute,rewrite=createattribute"
	CreateAttribute(name string) (a *Attr)

	// CreateAttributeNS
	// js:"createAttributeNS,rewrite=createattributens"
	CreateAttributeNS(namespaceURI string, qualifiedName string) (a *Attr)

	// CreateCDATASection
	// js:"createCDATASection,rewrite=createcdatasection"
	CreateCDATASection(data string) (c *CDATASection)

	// CreateComment Creates a comment object with the specified data.
	//     * @param data Sets the comment object's data.
	// js:"createComment,rewrite=createcomment"
	CreateComment(data string) (c *Comment)

	// CreateDocumentFragment Creates a new document.
	// js:"createDocumentFragment,rewrite=createdocumentfragment"
	CreateDocumentFragment() (d *DocumentFragment)

	// CreateElement Creates an instance of the element for the specified tag.
	//     * @param tagName The name of an element.
	// js:"createElement,rewrite=createelement"
	CreateElement(tagName string) (e Element)

	// CreateElementNS
	// js:"createElementNS,rewrite=createelementns"
	CreateElementNS(namespaceURI string, qualifiedName string) (e Element)

	// CreateExpression
	// js:"createExpression,rewrite=createexpression"
	CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (x *XPathExpression)

	// CreateNodeIterator Creates a NodeIterator object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list
	//     * @param filter A custom NodeFilter function to use. For more information, see filter. Use null for no filter.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	// js:"createNodeIterator,rewrite=createnodeiterator"
	CreateNodeIterator(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (n *NodeIterator)

	// CreateNSResolver
	// js:"createNSResolver,rewrite=creatensresolver"
	CreateNSResolver(nodeResolver Node) (x *xpathnsresolver.XPathNSResolver)

	// CreateProcessingInstruction
	// js:"createProcessingInstruction,rewrite=createprocessinginstruction"
	CreateProcessingInstruction(target string, data string) (p *ProcessingInstruction)

	// CreateRange Returns an empty range object that has both of its boundary points positioned at the beginning of the document.
	// js:"createRange,rewrite=createrange"
	CreateRange() (r *Range)

	// CreateTextNode Creates a text string from the specified value.
	//     * @param data String that specifies the nodeValue property of the text node.
	// js:"createTextNode,rewrite=createtextnode"
	CreateTextNode(data string) (t Text)

	// CreateTouch
	// js:"createTouch,rewrite=createtouch"
	CreateTouch(view *Window, target EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *Touch)

	// CreateTouchList
	// js:"createTouchList,rewrite=createtouchlist"
	CreateTouchList(touches *Touch) (t *TouchList)

	// CreateTreeWalker Creates a TreeWalker object that you can use to traverse filtered lists of nodes or elements in a document.
	//     * @param root The root element or node to start traversing on.
	//     * @param whatToShow The type of nodes or elements to appear in the node list. For more information, see whatToShow.
	//     * @param filter A custom NodeFilter function to use.
	//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
	// js:"createTreeWalker,rewrite=createtreewalker"
	CreateTreeWalker(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (t *TreeWalker)

	// ElementFromPoint Returns the element for the specified x coordinate and the specified y coordinate.
	//     * @param x The x-offset
	//     * @param y The y-offset
	// js:"elementFromPoint,rewrite=elementfrompoint"
	ElementFromPoint(x int, y int) (e Element)

	// Evaluate
	// js:"evaluate,rewrite=evaluate"
	Evaluate(expression string, contextNode Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *XPathResult) (x *XPathResult)

	// ExecCommand Executes a command on the current document, current selection, or the given range.
	//     * @param commandId String that specifies the command to execute. This command can be any of the command identifiers that can be executed in script.
	//     * @param showUI Display the user interface, defaults to false.
	//     * @param value Value to assign.
	// js:"execCommand,rewrite=execcommand"
	ExecCommand(commandId string, showUI *bool, value *interface{}) (b bool)

	// ExecCommandShowHelp Displays help information for the given command identifier.
	//     * @param commandId Displays help information for the given command identifier.
	// js:"execCommandShowHelp,rewrite=execcommandshowhelp"
	ExecCommandShowHelp(commandId string) (b bool)

	// ExitFullscreen
	// js:"exitFullscreen,rewrite=exitfullscreen"
	ExitFullscreen()

	// ExitPointerLock
	// js:"exitPointerLock,rewrite=exitpointerlock"
	ExitPointerLock()

	// Focus Causes the element to receive the focus and executes the code specified by the onfocus event.
	// js:"focus,rewrite=focus"
	Focus()

	// GetElementByID Returns a reference to the first object with the specified value of the ID or NAME attribute.
	//     * @param elementId String that specifies the ID value. Case-insensitive.
	// js:"getElementById,rewrite=getelementbyid"
	GetElementByID(elementId string) (e Element)

	// GetElementsByClassName
	// js:"getElementsByClassName,rewrite=getelementsbyclassname"
	GetElementsByClassName(classNames string) (n *NodeList)

	// GetElementsByName Gets a collection of objects based on the value of the NAME or ID attribute.
	//     * @param elementName Gets a collection of objects based on the value of the NAME or ID attribute.
	// js:"getElementsByName,rewrite=getelementsbyname"
	GetElementsByName(elementName string) (n *NodeList)

	// GetElementsByTagName Retrieves a collection of objects based on the specified element name.
	//     * @param name Specifies the name of an element.
	// js:"getElementsByTagName,rewrite=getelementsbytagname"
	GetElementsByTagName(tagname string) (n *NodeList)

	// GetElementsByTagNameNS
	// js:"getElementsByTagNameNS,rewrite=getelementsbytagnamens"
	GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList)

	// GetSelection Returns an object representing the current selection of the document that is loaded into the object displaying a webpage.
	// js:"getSelection,rewrite=getselection"
	GetSelection() (s *Selection)

	// HasFocus Gets a value indicating whether the object currently has focus.
	// js:"hasFocus,rewrite=hasfocus"
	HasFocus() (b bool)

	// ImportNode
	// js:"importNode,rewrite=importnode"
	ImportNode(importedNode Node, deep bool) (n Node)

	// MsElementsFromPoint
	// js:"msElementsFromPoint,rewrite=mselementsfrompoint"
	MsElementsFromPoint(x float32, y float32) (n *NodeList)

	// MsElementsFromRect
	// js:"msElementsFromRect,rewrite=mselementsfromrect"
	MsElementsFromRect(left float32, top float32, width float32, height float32) (n *NodeList)

	// Open Opens a new window and loads a document specified by a given URL. Also, opens a new window that uses the url parameter and the name parameter to collect the output of the write method and the writeln method.
	//     * @param url Specifies a MIME type for the document.
	//     * @param name Specifies the name of the window. This name is used as the value for the TARGET attribute on a form or an anchor element.
	//     * @param features Contains a list of items separated by commas. Each item consists of an option and a value, separated by an equals sign (for example, "fullscreen=yes, toolbar=yes"). The following values are supported.
	//     * @param replace Specifies whether the existing entry for the document is replaced in the history list.
	// js:"open,rewrite=open"
	Open(url *string, name *string, features *string, replace *bool) (i interface{})

	// QueryCommandEnabled Returns a Boolean value that indicates whether a specified command can be successfully executed using execCommand, given the current state of the document.
	//     * @param commandId Specifies a command identifier.
	// js:"queryCommandEnabled,rewrite=querycommandenabled"
	QueryCommandEnabled(commandId string) (b bool)

	// QueryCommandIndeterm Returns a Boolean value that indicates whether the specified command is in the indeterminate state.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandIndeterm,rewrite=querycommandindeterm"
	QueryCommandIndeterm(commandId string) (b bool)

	// QueryCommandState Returns a Boolean value that indicates the current state of the command.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandState,rewrite=querycommandstate"
	QueryCommandState(commandId string) (b bool)

	// QueryCommandSupported Returns a Boolean value that indicates whether the current command is supported on the current range.
	//     * @param commandId Specifies a command identifier.
	// js:"queryCommandSupported,rewrite=querycommandsupported"
	QueryCommandSupported(commandId string) (b bool)

	// QueryCommandText Retrieves the string associated with a command.
	//     * @param commandId String that contains the identifier of a command. This can be any command identifier given in the list of Command Identifiers.
	// js:"queryCommandText,rewrite=querycommandtext"
	QueryCommandText(commandId string) (s string)

	// QueryCommandValue Returns the current value of the document, range, or current selection for the given command.
	//     * @param commandId String that specifies a command identifier.
	// js:"queryCommandValue,rewrite=querycommandvalue"
	QueryCommandValue(commandId string) (s string)

	// ReleaseEvents
	// js:"releaseEvents,rewrite=releaseevents"
	ReleaseEvents()

	// UpdateSettings Allows updating the print settings for the page.
	// js:"updateSettings,rewrite=updatesettings"
	UpdateSettings()

	// WebkitCancelFullScreen
	// js:"webkitCancelFullScreen,rewrite=webkitcancelfullscreen"
	WebkitCancelFullScreen()

	// WebkitExitFullscreen
	// js:"webkitExitFullscreen,rewrite=webkitexitfullscreen"
	WebkitExitFullscreen()

	// Write Writes one or more HTML expressions to a document in the specified window.
	//     * @param content Specifies the text and HTML tags to write.
	// js:"write,rewrite=write"
	Write(content string)

	// Writeln Writes one or more HTML expressions, followed by a carriage return, to a document in the specified window.
	//     * @param content The text and HTML tags to write.
	// js:"writeln,rewrite=writeln"
	Writeln(content string)

	// Onpointercancel prop
	// js:"onpointercancel,rewrite=onpointercancel"
	Onpointercancel() (onpointercancel func(Event))

	// Onpointercancel prop
	// js:"setonpointercancel,rewrite=setonpointercancel"
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown prop
	// js:"onpointerdown,rewrite=onpointerdown"
	Onpointerdown() (onpointerdown func(Event))

	// Onpointerdown prop
	// js:"setonpointerdown,rewrite=setonpointerdown"
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter prop
	// js:"onpointerenter,rewrite=onpointerenter"
	Onpointerenter() (onpointerenter func(Event))

	// Onpointerenter prop
	// js:"setonpointerenter,rewrite=setonpointerenter"
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave prop
	// js:"onpointerleave,rewrite=onpointerleave"
	Onpointerleave() (onpointerleave func(Event))

	// Onpointerleave prop
	// js:"setonpointerleave,rewrite=setonpointerleave"
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove prop
	// js:"onpointermove,rewrite=onpointermove"
	Onpointermove() (onpointermove func(Event))

	// Onpointermove prop
	// js:"setonpointermove,rewrite=setonpointermove"
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout prop
	// js:"onpointerout,rewrite=onpointerout"
	Onpointerout() (onpointerout func(Event))

	// Onpointerout prop
	// js:"setonpointerout,rewrite=setonpointerout"
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover prop
	// js:"onpointerover,rewrite=onpointerover"
	Onpointerover() (onpointerover func(Event))

	// Onpointerover prop
	// js:"setonpointerover,rewrite=setonpointerover"
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup prop
	// js:"onpointerup,rewrite=onpointerup"
	Onpointerup() (onpointerup func(Event))

	// Onpointerup prop
	// js:"setonpointerup,rewrite=setonpointerup"
	SetOnpointerup(onpointerup func(Event))

	// Onwheel prop
	// js:"onwheel,rewrite=onwheel"
	Onwheel() (onwheel func(Event))

	// // Onwheel prop
	// // js:"setonwheel,rewrite=setonwheel"
	// SetOnwheel(onwheel func(Event))

	// // ActiveElement prop Gets the object that has the focus when the parent document has focus.
	// // js:"activeElement,rewrite=activeelement"
	// ActiveElement() (activeElement Element)

	// // AlinkColor prop Sets or gets the color of all active links in the document.
	// // js:"alinkColor,rewrite=alinkcolor"
	// AlinkColor() (alinkColor string)

	// // AlinkColor prop Sets or gets the color of all active links in the document.
	// // js:"setalinkColor,rewrite=setalinkcolor"
	// SetAlinkColor(alinkColor string)

	// // All prop Returns a reference to the collection of elements contained by the object.
	// // js:"all,rewrite=all"
	// All() (all *htmlallcollection.HTMLAllCollection)

	// // Anchors prop Retrieves a collection of all a objects that have a name and/or id property. Objects in this collection are in HTML source order.
	// // js:"anchors,rewrite=anchors"
	// Anchors() (anchors HTMLCollection)

	// // Applets prop Retrieves a collection of all applet objects in the document.
	// // js:"applets,rewrite=applets"
	// Applets() (applets HTMLCollection)

	// // BgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	// // js:"bgColor,rewrite=bgcolor"
	// BgColor() (bgColor string)

	// // BgColor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
	// // js:"setbgColor,rewrite=setbgcolor"
	// SetBgColor(bgColor string)

	// // Body prop Specifies the beginning and end of the document body.
	// // js:"body,rewrite=body"
	// Body() (body HTMLElement)

	// // Body prop Specifies the beginning and end of the document body.
	// // js:"setbody,rewrite=setbody"
	// SetBody(body HTMLElement)

	// // CharacterSet prop
	// // js:"characterSet,rewrite=characterset"
	// CharacterSet() (characterSet string)

	// // Charset prop Gets or sets the character set used to encode the object.
	// // js:"charset,rewrite=charset"
	// Charset() (charset string)

	// // Charset prop Gets or sets the character set used to encode the object.
	// // js:"setcharset,rewrite=setcharset"
	// SetCharset(charset string)

	// // CompatMode prop Gets a value that indicates whether standards-compliant mode is switched on for the object.
	// // js:"compatMode,rewrite=compatmode"
	// CompatMode() (compatMode string)

	// // Cookie prop
	// // js:"cookie,rewrite=cookie"
	// Cookie() (cookie string)

	// // Cookie prop
	// // js:"setcookie,rewrite=setcookie"
	// SetCookie(cookie string)

	// // CurrentScript prop
	// // js:"currentScript,rewrite=currentscript"
	// CurrentScript() (currentScript interface{})

	// // DefaultView prop
	// // js:"defaultView,rewrite=defaultview"
	// DefaultView() (defaultView *Window)

	// // DesignMode prop Sets or gets a value that indicates whether the document can be edited.
	// // js:"designMode,rewrite=designmode"
	// DesignMode() (designMode string)

	// // DesignMode prop Sets or gets a value that indicates whether the document can be edited.
	// // js:"setdesignMode,rewrite=setdesignmode"
	// SetDesignMode(designMode string)

	// // Dir prop Sets or retrieves a value that indicates the reading order of the object.
	// // js:"dir,rewrite=dir"
	// Dir() (dir string)

	// // Dir prop Sets or retrieves a value that indicates the reading order of the object.
	// // js:"setdir,rewrite=setdir"
	// SetDir(dir string)

	// // Doctype prop Gets an object representing the document type declaration associated with the current document.
	// // js:"doctype,rewrite=doctype"
	// Doctype() (doctype *DocumentType)

	// DocumentElement prop Gets a reference to the root node of the document.
	// js:"documentElement,rewrite=documentelement"
	DocumentElement() (documentElement Element)

	// // Domain prop Sets or gets the security domain of the document.
	// // js:"domain,rewrite=domain"
	// Domain() (domain string)

	// // Domain prop Sets or gets the security domain of the document.
	// // js:"setdomain,rewrite=setdomain"
	// SetDomain(domain string)

	// // Embeds prop Retrieves a collection of all embed objects in the document.
	// // js:"embeds,rewrite=embeds"
	// Embeds() (embeds HTMLCollection)

	// // FgColor prop Sets or gets the foreground (text) color of the document.
	// // js:"fgColor,rewrite=fgcolor"
	// FgColor() (fgColor string)

	// // FgColor prop Sets or gets the foreground (text) color of the document.
	// // js:"setfgColor,rewrite=setfgcolor"
	// SetFgColor(fgColor string)

	// // Forms prop Retrieves a collection, in source order, of all form objects in the document.
	// // js:"forms,rewrite=forms"
	// Forms() (forms HTMLCollection)

	// // FullscreenElement prop
	// // js:"fullscreenElement,rewrite=fullscreenelement"
	// FullscreenElement() (fullscreenElement Element)

	// // FullscreenEnabled prop
	// // js:"fullscreenEnabled,rewrite=fullscreenenabled"
	// FullscreenEnabled() (fullscreenEnabled bool)

	// // Head prop
	// // js:"head,rewrite=head"
	// Head() (head *HTMLHeadElement)

	// // Hidden prop
	// // js:"hidden,rewrite=hidden"
	// Hidden() (hidden bool)

	// // Images prop Retrieves a collection, in source order, of img objects in the document.
	// // js:"images,rewrite=images"
	// Images() (images HTMLCollection)

	// // Implementation prop Gets the implementation object of the current document.
	// // js:"implementation,rewrite=implementation"
	// Implementation() (implementation *DOMImplementation)

	// // InputEncoding prop Returns the character encoding used to create the webpage that is loaded into the document object.
	// // js:"inputEncoding,rewrite=inputencoding"
	// InputEncoding() (inputEncoding string)

	// // LastModified prop Gets the date that the page was last modified, if the page supplies one.
	// // js:"lastModified,rewrite=lastmodified"
	// LastModified() (lastModified string)

	// // LinkColor prop Sets or gets the color of the document links.
	// // js:"linkColor,rewrite=linkcolor"
	// LinkColor() (linkColor string)

	// // LinkColor prop Sets or gets the color of the document links.
	// // js:"setlinkColor,rewrite=setlinkcolor"
	// SetLinkColor(linkColor string)

	// // Links prop Retrieves a collection of all a objects that specify the href property and all area objects in the document.
	// // js:"links,rewrite=links"
	// Links() (links HTMLCollection)

	// // Location prop Contains information about the current URL.
	// // js:"location,rewrite=location"
	// // Location() (location *location.Location)

	// // MsCapsLockWarningOff prop
	// // js:"msCapsLockWarningOff,rewrite=mscapslockwarningoff"
	// MsCapsLockWarningOff() (msCapsLockWarningOff bool)

	// // MsCapsLockWarningOff prop
	// // js:"setmsCapsLockWarningOff,rewrite=setmscapslockwarningoff"
	// SetMsCapsLockWarningOff(msCapsLockWarningOff bool)

	// // MsCSSOMElementFloatMetrics prop
	// // js:"msCSSOMElementFloatMetrics,rewrite=mscssomelementfloatmetrics"
	// MsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool)

	// // MsCSSOMElementFloatMetrics prop
	// // js:"setmsCSSOMElementFloatMetrics,rewrite=setmscssomelementfloatmetrics"
	// SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool)

	// // Onabort prop Fires when the user aborts the download.
	// //     * @param ev The event.
	// // js:"onabort,rewrite=onabort"
	// Onabort() (onabort func(Event))

	// // Onabort prop Fires when the user aborts the download.
	// //     * @param ev The event.
	// // js:"setonabort,rewrite=setonabort"
	// SetOnabort(onabort func(Event))

	// // Onactivate prop Fires when the object is set as the active element.
	// //     * @param ev The event.
	// // js:"onactivate,rewrite=onactivate"
	// Onactivate() (onactivate func(Event))

	// // Onactivate prop Fires when the object is set as the active element.
	// //     * @param ev The event.
	// // js:"setonactivate,rewrite=setonactivate"
	// SetOnactivate(onactivate func(Event))

	// // Onbeforeactivate prop Fires immediately before the object is set as the active element.
	// //     * @param ev The event.
	// // js:"onbeforeactivate,rewrite=onbeforeactivate"
	// Onbeforeactivate() (onbeforeactivate func(Event))

	// // Onbeforeactivate prop Fires immediately before the object is set as the active element.
	// //     * @param ev The event.
	// // js:"setonbeforeactivate,rewrite=setonbeforeactivate"
	// SetOnbeforeactivate(onbeforeactivate func(Event))

	// // Onbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	// //     * @param ev The event.
	// // js:"onbeforedeactivate,rewrite=onbeforedeactivate"
	// Onbeforedeactivate() (onbeforedeactivate func(Event))

	// // Onbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
	// //     * @param ev The event.
	// // js:"setonbeforedeactivate,rewrite=setonbeforedeactivate"
	// SetOnbeforedeactivate(onbeforedeactivate func(Event))

	// // Onblur prop Fires when the object loses the input focus.
	// //     * @param ev The focus event.
	// // js:"onblur,rewrite=onblur"
	// Onblur() (onblur func(Event))

	// // Onblur prop Fires when the object loses the input focus.
	// //     * @param ev The focus event.
	// // js:"setonblur,rewrite=setonblur"
	// SetOnblur(onblur func(Event))

	// // Oncanplay prop Occurs when playback is possible, but would require further buffering.
	// //     * @param ev The event.
	// // js:"oncanplay,rewrite=oncanplay"
	// Oncanplay() (oncanplay func(Event))

	// // Oncanplay prop Occurs when playback is possible, but would require further buffering.
	// //     * @param ev The event.
	// // js:"setoncanplay,rewrite=setoncanplay"
	// SetOncanplay(oncanplay func(Event))

	// // Oncanplaythrough prop
	// // js:"oncanplaythrough,rewrite=oncanplaythrough"
	// Oncanplaythrough() (oncanplaythrough func(Event))

	// // Oncanplaythrough prop
	// // js:"setoncanplaythrough,rewrite=setoncanplaythrough"
	// SetOncanplaythrough(oncanplaythrough func(Event))

	// // Onchange prop Fires when the contents of the object or selection have changed.
	// //     * @param ev The event.
	// // js:"onchange,rewrite=onchange"
	// Onchange() (onchange func(Event))

	// // Onchange prop Fires when the contents of the object or selection have changed.
	// //     * @param ev The event.
	// // js:"setonchange,rewrite=setonchange"
	// SetOnchange(onchange func(Event))

	// // Onclick prop Fires when the user clicks the left mouse button on the object
	// //     * @param ev The mouse event.
	// // js:"onclick,rewrite=onclick"
	// Onclick() (onclick func(Event))

	// // Onclick prop Fires when the user clicks the left mouse button on the object
	// //     * @param ev The mouse event.
	// // js:"setonclick,rewrite=setonclick"
	// SetOnclick(onclick func(Event))

	// // Oncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
	// //     * @param ev The mouse event.
	// // js:"oncontextmenu,rewrite=oncontextmenu"
	// Oncontextmenu() (oncontextmenu func(Event))

	// // Oncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
	// //     * @param ev The mouse event.
	// // js:"setoncontextmenu,rewrite=setoncontextmenu"
	// SetOncontextmenu(oncontextmenu func(Event))

	// // Ondblclick prop Fires when the user double-clicks the object.
	// //     * @param ev The mouse event.
	// // js:"ondblclick,rewrite=ondblclick"
	// Ondblclick() (ondblclick func(Event))

	// // Ondblclick prop Fires when the user double-clicks the object.
	// //     * @param ev The mouse event.
	// // js:"setondblclick,rewrite=setondblclick"
	// SetOndblclick(ondblclick func(Event))

	// // Ondeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
	// //     * @param ev The UI Event
	// // js:"ondeactivate,rewrite=ondeactivate"
	// Ondeactivate() (ondeactivate func(Event))

	// // Ondeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
	// //     * @param ev The UI Event
	// // js:"setondeactivate,rewrite=setondeactivate"
	// SetOndeactivate(ondeactivate func(Event))

	// // Ondrag prop Fires on the source object continuously during a drag operation.
	// //     * @param ev The event.
	// // js:"ondrag,rewrite=ondrag"
	// Ondrag() (ondrag func(Event))

	// // Ondrag prop Fires on the source object continuously during a drag operation.
	// //     * @param ev The event.
	// // js:"setondrag,rewrite=setondrag"
	// SetOndrag(ondrag func(Event))

	// // Ondragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
	// //     * @param ev The event.
	// // js:"ondragend,rewrite=ondragend"
	// Ondragend() (ondragend func(Event))

	// // Ondragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
	// //     * @param ev The event.
	// // js:"setondragend,rewrite=setondragend"
	// SetOndragend(ondragend func(Event))

	// // Ondragenter prop Fires on the target element when the user drags the object to a valid drop target.
	// //     * @param ev The drag event.
	// // js:"ondragenter,rewrite=ondragenter"
	// Ondragenter() (ondragenter func(Event))

	// // Ondragenter prop Fires on the target element when the user drags the object to a valid drop target.
	// //     * @param ev The drag event.
	// // js:"setondragenter,rewrite=setondragenter"
	// SetOndragenter(ondragenter func(Event))

	// // Ondragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	// //     * @param ev The drag event.
	// // js:"ondragleave,rewrite=ondragleave"
	// Ondragleave() (ondragleave func(Event))

	// // Ondragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
	// //     * @param ev The drag event.
	// // js:"setondragleave,rewrite=setondragleave"
	// SetOndragleave(ondragleave func(Event))

	// // Ondragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
	// //     * @param ev The event.
	// // js:"ondragover,rewrite=ondragover"
	// Ondragover() (ondragover func(Event))

	// // Ondragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
	// //     * @param ev The event.
	// // js:"setondragover,rewrite=setondragover"
	// SetOndragover(ondragover func(Event))

	// // Ondragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
	// //     * @param ev The event.
	// // js:"ondragstart,rewrite=ondragstart"
	// Ondragstart() (ondragstart func(Event))

	// // Ondragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
	// //     * @param ev The event.
	// // js:"setondragstart,rewrite=setondragstart"
	// SetOndragstart(ondragstart func(Event))

	// // Ondrop prop
	// // js:"ondrop,rewrite=ondrop"
	// Ondrop() (ondrop func(Event))

	// // Ondrop prop
	// // js:"setondrop,rewrite=setondrop"
	// SetOndrop(ondrop func(Event))

	// // Ondurationchange prop Occurs when the duration attribute is updated.
	// //     * @param ev The event.
	// // js:"ondurationchange,rewrite=ondurationchange"
	// Ondurationchange() (ondurationchange func(Event))

	// // Ondurationchange prop Occurs when the duration attribute is updated.
	// //     * @param ev The event.
	// // js:"setondurationchange,rewrite=setondurationchange"
	// SetOndurationchange(ondurationchange func(Event))

	// // Onemptied prop Occurs when the media element is reset to its initial state.
	// //     * @param ev The event.
	// // js:"onemptied,rewrite=onemptied"
	// Onemptied() (onemptied func(Event))

	// // Onemptied prop Occurs when the media element is reset to its initial state.
	// //     * @param ev The event.
	// // js:"setonemptied,rewrite=setonemptied"
	// SetOnemptied(onemptied func(Event))

	// // Onended prop Occurs when the end of playback is reached.
	// //     * @param ev The event
	// // js:"onended,rewrite=onended"
	// Onended() (onended func(Event))

	// // Onended prop Occurs when the end of playback is reached.
	// //     * @param ev The event
	// // js:"setonended,rewrite=setonended"
	// SetOnended(onended func(Event))

	// // Onerror prop Fires when an error occurs during object loading.
	// //     * @param ev The event.
	// // js:"onerror,rewrite=onerror"
	// Onerror() (onerror func(Event))

	// // Onerror prop Fires when an error occurs during object loading.
	// //     * @param ev The event.
	// // js:"setonerror,rewrite=setonerror"
	// SetOnerror(onerror func(Event))

	// // Onfocus prop Fires when the object receives focus.
	// //     * @param ev The event.
	// // js:"onfocus,rewrite=onfocus"
	// Onfocus() (onfocus func(Event))

	// // Onfocus prop Fires when the object receives focus.
	// //     * @param ev The event.
	// // js:"setonfocus,rewrite=setonfocus"
	// SetOnfocus(onfocus func(Event))

	// // Onfullscreenchange prop
	// // js:"onfullscreenchange,rewrite=onfullscreenchange"
	// Onfullscreenchange() (onfullscreenchange func(Event))

	// // Onfullscreenchange prop
	// // js:"setonfullscreenchange,rewrite=setonfullscreenchange"
	// SetOnfullscreenchange(onfullscreenchange func(Event))

	// // Onfullscreenerror prop
	// // js:"onfullscreenerror,rewrite=onfullscreenerror"
	// Onfullscreenerror() (onfullscreenerror func(Event))

	// // Onfullscreenerror prop
	// // js:"setonfullscreenerror,rewrite=setonfullscreenerror"
	// SetOnfullscreenerror(onfullscreenerror func(Event))

	// // Oninput prop
	// // js:"oninput,rewrite=oninput"
	// Oninput() (oninput func(Event))

	// // Oninput prop
	// // js:"setoninput,rewrite=setoninput"
	// SetOninput(oninput func(Event))

	// // Oninvalid prop
	// // js:"oninvalid,rewrite=oninvalid"
	// Oninvalid() (oninvalid func(Event))

	// // Oninvalid prop
	// // js:"setoninvalid,rewrite=setoninvalid"
	// SetOninvalid(oninvalid func(Event))

	// // Onkeydown prop Fires when the user presses a key.
	// //     * @param ev The keyboard event
	// // js:"onkeydown,rewrite=onkeydown"
	// Onkeydown() (onkeydown func(Event))

	// // Onkeydown prop Fires when the user presses a key.
	// //     * @param ev The keyboard event
	// // js:"setonkeydown,rewrite=setonkeydown"
	// SetOnkeydown(onkeydown func(Event))

	// // Onkeypress prop Fires when the user presses an alphanumeric key.
	// //     * @param ev The event.
	// // js:"onkeypress,rewrite=onkeypress"
	// Onkeypress() (onkeypress func(Event))

	// // Onkeypress prop Fires when the user presses an alphanumeric key.
	// //     * @param ev The event.
	// // js:"setonkeypress,rewrite=setonkeypress"
	// SetOnkeypress(onkeypress func(Event))

	// // Onkeyup prop Fires when the user releases a key.
	// //     * @param ev The keyboard event
	// // js:"onkeyup,rewrite=onkeyup"
	// Onkeyup() (onkeyup func(Event))

	// // Onkeyup prop Fires when the user releases a key.
	// //     * @param ev The keyboard event
	// // js:"setonkeyup,rewrite=setonkeyup"
	// SetOnkeyup(onkeyup func(Event))

	// // Onload prop Fires immediately after the browser loads the object.
	// //     * @param ev The event.
	// // js:"onload,rewrite=onload"
	// Onload() (onload func(Event))

	// // Onload prop Fires immediately after the browser loads the object.
	// //     * @param ev The event.
	// // js:"setonload,rewrite=setonload"
	// SetOnload(onload func(Event))

	// // Onloadeddata prop Occurs when media data is loaded at the current playback position.
	// //     * @param ev The event.
	// // js:"onloadeddata,rewrite=onloadeddata"
	// Onloadeddata() (onloadeddata func(Event))

	// // Onloadeddata prop Occurs when media data is loaded at the current playback position.
	// //     * @param ev The event.
	// // js:"setonloadeddata,rewrite=setonloadeddata"
	// SetOnloadeddata(onloadeddata func(Event))

	// // Onloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
	// //     * @param ev The event.
	// // js:"onloadedmetadata,rewrite=onloadedmetadata"
	// Onloadedmetadata() (onloadedmetadata func(Event))

	// // Onloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
	// //     * @param ev The event.
	// // js:"setonloadedmetadata,rewrite=setonloadedmetadata"
	// SetOnloadedmetadata(onloadedmetadata func(Event))

	// // Onloadstart prop Occurs when Internet Explorer begins looking for media data.
	// //     * @param ev The event.
	// // js:"onloadstart,rewrite=onloadstart"
	// Onloadstart() (onloadstart func(Event))

	// // Onloadstart prop Occurs when Internet Explorer begins looking for media data.
	// //     * @param ev The event.
	// // js:"setonloadstart,rewrite=setonloadstart"
	// SetOnloadstart(onloadstart func(Event))

	// // Onmousedown prop Fires when the user clicks the object with either mouse button.
	// //     * @param ev The mouse event.
	// // js:"onmousedown,rewrite=onmousedown"
	// Onmousedown() (onmousedown func(Event))

	// // Onmousedown prop Fires when the user clicks the object with either mouse button.
	// //     * @param ev The mouse event.
	// // js:"setonmousedown,rewrite=setonmousedown"
	// SetOnmousedown(onmousedown func(Event))

	// // Onmousemove prop Fires when the user moves the mouse over the object.
	// //     * @param ev The mouse event.
	// // js:"onmousemove,rewrite=onmousemove"
	// Onmousemove() (onmousemove func(Event))

	// // Onmousemove prop Fires when the user moves the mouse over the object.
	// //     * @param ev The mouse event.
	// // js:"setonmousemove,rewrite=setonmousemove"
	// SetOnmousemove(onmousemove func(Event))

	// // Onmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
	// //     * @param ev The mouse event.
	// // js:"onmouseout,rewrite=onmouseout"
	// Onmouseout() (onmouseout func(Event))

	// // Onmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
	// //     * @param ev The mouse event.
	// // js:"setonmouseout,rewrite=setonmouseout"
	// SetOnmouseout(onmouseout func(Event))

	// // Onmouseover prop Fires when the user moves the mouse pointer into the object.
	// //     * @param ev The mouse event.
	// // js:"onmouseover,rewrite=onmouseover"
	// Onmouseover() (onmouseover func(Event))

	// // Onmouseover prop Fires when the user moves the mouse pointer into the object.
	// //     * @param ev The mouse event.
	// // js:"setonmouseover,rewrite=setonmouseover"
	// SetOnmouseover(onmouseover func(Event))

	// // Onmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
	// //     * @param ev The mouse event.
	// // js:"onmouseup,rewrite=onmouseup"
	// Onmouseup() (onmouseup func(Event))

	// // Onmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
	// //     * @param ev The mouse event.
	// // js:"setonmouseup,rewrite=setonmouseup"
	// SetOnmouseup(onmouseup func(Event))

	// // Onmousewheel prop Fires when the wheel button is rotated.
	// //     * @param ev The mouse event
	// // js:"onmousewheel,rewrite=onmousewheel"
	// Onmousewheel() (onmousewheel func(Event))

	// // Onmousewheel prop Fires when the wheel button is rotated.
	// //     * @param ev The mouse event
	// // js:"setonmousewheel,rewrite=setonmousewheel"
	// SetOnmousewheel(onmousewheel func(Event))

	// // Onmscontentzoom prop
	// // js:"onmscontentzoom,rewrite=onmscontentzoom"
	// Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// // Onmscontentzoom prop
	// // js:"setonmscontentzoom,rewrite=setonmscontentzoom"
	// SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// // Onmsgesturechange prop
	// // js:"onmsgesturechange,rewrite=onmsgesturechange"
	// Onmsgesturechange() (onmsgesturechange func(Event))

	// // Onmsgesturechange prop
	// // js:"setonmsgesturechange,rewrite=setonmsgesturechange"
	// SetOnmsgesturechange(onmsgesturechange func(Event))

	// // Onmsgesturedoubletap prop
	// // js:"onmsgesturedoubletap,rewrite=onmsgesturedoubletap"
	// Onmsgesturedoubletap() (onmsgesturedoubletap func(Event))

	// // Onmsgesturedoubletap prop
	// // js:"setonmsgesturedoubletap,rewrite=setonmsgesturedoubletap"
	// SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event))

	// // Onmsgestureend prop
	// // js:"onmsgestureend,rewrite=onmsgestureend"
	// Onmsgestureend() (onmsgestureend func(Event))

	// // Onmsgestureend prop
	// // js:"setonmsgestureend,rewrite=setonmsgestureend"
	// SetOnmsgestureend(onmsgestureend func(Event))

	// // Onmsgesturehold prop
	// // js:"onmsgesturehold,rewrite=onmsgesturehold"
	// Onmsgesturehold() (onmsgesturehold func(Event))

	// // Onmsgesturehold prop
	// // js:"setonmsgesturehold,rewrite=setonmsgesturehold"
	// SetOnmsgesturehold(onmsgesturehold func(Event))

	// // Onmsgesturestart prop
	// // js:"onmsgesturestart,rewrite=onmsgesturestart"
	// Onmsgesturestart() (onmsgesturestart func(Event))

	// // Onmsgesturestart prop
	// // js:"setonmsgesturestart,rewrite=setonmsgesturestart"
	// SetOnmsgesturestart(onmsgesturestart func(Event))

	// // Onmsgesturetap prop
	// // js:"onmsgesturetap,rewrite=onmsgesturetap"
	// Onmsgesturetap() (onmsgesturetap func(Event))

	// // Onmsgesturetap prop
	// // js:"setonmsgesturetap,rewrite=setonmsgesturetap"
	// SetOnmsgesturetap(onmsgesturetap func(Event))

	// // Onmsinertiastart prop
	// // js:"onmsinertiastart,rewrite=onmsinertiastart"
	// Onmsinertiastart() (onmsinertiastart func(Event))

	// // Onmsinertiastart prop
	// // js:"setonmsinertiastart,rewrite=setonmsinertiastart"
	// SetOnmsinertiastart(onmsinertiastart func(Event))

	// // Onmsmanipulationstatechanged prop
	// // js:"onmsmanipulationstatechanged,rewrite=onmsmanipulationstatechanged"
	// Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// // Onmsmanipulationstatechanged prop
	// // js:"setonmsmanipulationstatechanged,rewrite=setonmsmanipulationstatechanged"
	// SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// // Onmspointercancel prop
	// // js:"onmspointercancel,rewrite=onmspointercancel"
	// Onmspointercancel() (onmspointercancel func(Event))

	// // Onmspointercancel prop
	// // js:"setonmspointercancel,rewrite=setonmspointercancel"
	// SetOnmspointercancel(onmspointercancel func(Event))

	// // Onmspointerdown prop
	// // js:"onmspointerdown,rewrite=onmspointerdown"
	// Onmspointerdown() (onmspointerdown func(Event))

	// // Onmspointerdown prop
	// // js:"setonmspointerdown,rewrite=setonmspointerdown"
	// SetOnmspointerdown(onmspointerdown func(Event))

	// // Onmspointerenter prop
	// // js:"onmspointerenter,rewrite=onmspointerenter"
	// Onmspointerenter() (onmspointerenter func(Event))

	// // Onmspointerenter prop
	// // js:"setonmspointerenter,rewrite=setonmspointerenter"
	// SetOnmspointerenter(onmspointerenter func(Event))

	// // Onmspointerleave prop
	// // js:"onmspointerleave,rewrite=onmspointerleave"
	// Onmspointerleave() (onmspointerleave func(Event))

	// // Onmspointerleave prop
	// // js:"setonmspointerleave,rewrite=setonmspointerleave"
	// SetOnmspointerleave(onmspointerleave func(Event))

	// // Onmspointermove prop
	// // js:"onmspointermove,rewrite=onmspointermove"
	// Onmspointermove() (onmspointermove func(Event))

	// // Onmspointermove prop
	// // js:"setonmspointermove,rewrite=setonmspointermove"
	// SetOnmspointermove(onmspointermove func(Event))

	// // Onmspointerout prop
	// // js:"onmspointerout,rewrite=onmspointerout"
	// Onmspointerout() (onmspointerout func(Event))

	// // Onmspointerout prop
	// // js:"setonmspointerout,rewrite=setonmspointerout"
	// SetOnmspointerout(onmspointerout func(Event))

	// // Onmspointerover prop
	// // js:"onmspointerover,rewrite=onmspointerover"
	// Onmspointerover() (onmspointerover func(Event))

	// // Onmspointerover prop
	// // js:"setonmspointerover,rewrite=setonmspointerover"
	// SetOnmspointerover(onmspointerover func(Event))

	// // Onmspointerup prop
	// // js:"onmspointerup,rewrite=onmspointerup"
	// Onmspointerup() (onmspointerup func(Event))

	// // Onmspointerup prop
	// // js:"setonmspointerup,rewrite=setonmspointerup"
	// SetOnmspointerup(onmspointerup func(Event))

	// // Onmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	// //     * @param ev The event.
	// // js:"onmssitemodejumplistitemremoved,rewrite=onmssitemodejumplistitemremoved"
	// Onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// // Onmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
	// //     * @param ev The event.
	// // js:"setonmssitemodejumplistitemremoved,rewrite=setonmssitemodejumplistitemremoved"
	// SetOnmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*MSSiteModeEvent))

	// // Onmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	// //     * @param ev The event.
	// // js:"onmsthumbnailclick,rewrite=onmsthumbnailclick"
	// Onmsthumbnailclick() (onmsthumbnailclick func(*MSSiteModeEvent))

	// // Onmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
	// //     * @param ev The event.
	// // js:"setonmsthumbnailclick,rewrite=setonmsthumbnailclick"
	// SetOnmsthumbnailclick(onmsthumbnailclick func(*MSSiteModeEvent))

	// // Onpause prop Occurs when playback is paused.
	// //     * @param ev The event.
	// // js:"onpause,rewrite=onpause"
	// Onpause() (onpause func(Event))

	// // Onpause prop Occurs when playback is paused.
	// //     * @param ev The event.
	// // js:"setonpause,rewrite=setonpause"
	// SetOnpause(onpause func(Event))

	// // Onplay prop Occurs when the play method is requested.
	// //     * @param ev The event.
	// // js:"onplay,rewrite=onplay"
	// Onplay() (onplay func(Event))

	// // Onplay prop Occurs when the play method is requested.
	// //     * @param ev The event.
	// // js:"setonplay,rewrite=setonplay"
	// SetOnplay(onplay func(Event))

	// // Onplaying prop Occurs when the audio or video has started playing.
	// //     * @param ev The event.
	// // js:"onplaying,rewrite=onplaying"
	// Onplaying() (onplaying func(Event))

	// // Onplaying prop Occurs when the audio or video has started playing.
	// //     * @param ev The event.
	// // js:"setonplaying,rewrite=setonplaying"
	// SetOnplaying(onplaying func(Event))

	// // Onpointerlockchange prop
	// // js:"onpointerlockchange,rewrite=onpointerlockchange"
	// Onpointerlockchange() (onpointerlockchange func(Event))

	// // Onpointerlockchange prop
	// // js:"setonpointerlockchange,rewrite=setonpointerlockchange"
	// SetOnpointerlockchange(onpointerlockchange func(Event))

	// // Onpointerlockerror prop
	// // js:"onpointerlockerror,rewrite=onpointerlockerror"
	// Onpointerlockerror() (onpointerlockerror func(Event))

	// // Onpointerlockerror prop
	// // js:"setonpointerlockerror,rewrite=setonpointerlockerror"
	// SetOnpointerlockerror(onpointerlockerror func(Event))

	// // Onprogress prop Occurs to indicate progress while downloading media data.
	// //     * @param ev The event.
	// // js:"onprogress,rewrite=onprogress"
	// Onprogress() (onprogress func(Event))

	// // Onprogress prop Occurs to indicate progress while downloading media data.
	// //     * @param ev The event.
	// // js:"setonprogress,rewrite=setonprogress"
	// SetOnprogress(onprogress func(Event))

	// // Onratechange prop Occurs when the playback rate is increased or decreased.
	// //     * @param ev The event.
	// // js:"onratechange,rewrite=onratechange"
	// Onratechange() (onratechange func(Event))

	// // Onratechange prop Occurs when the playback rate is increased or decreased.
	// //     * @param ev The event.
	// // js:"setonratechange,rewrite=setonratechange"
	// SetOnratechange(onratechange func(Event))

	// // Onreadystatechange prop Fires when the state of the object has changed.
	// //     * @param ev The event
	// // js:"onreadystatechange,rewrite=onreadystatechange"
	// Onreadystatechange() (onreadystatechange func(Event))

	// // Onreadystatechange prop Fires when the state of the object has changed.
	// //     * @param ev The event
	// // js:"setonreadystatechange,rewrite=setonreadystatechange"
	// SetOnreadystatechange(onreadystatechange func(Event))

	// // Onreset prop Fires when the user resets a form.
	// //     * @param ev The event.
	// // js:"onreset,rewrite=onreset"
	// Onreset() (onreset func(Event))

	// // Onreset prop Fires when the user resets a form.
	// //     * @param ev The event.
	// // js:"setonreset,rewrite=setonreset"
	// SetOnreset(onreset func(Event))

	// // Onscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
	// //     * @param ev The event.
	// // js:"onscroll,rewrite=onscroll"
	// Onscroll() (onscroll func(Event))

	// // Onscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
	// //     * @param ev The event.
	// // js:"setonscroll,rewrite=setonscroll"
	// SetOnscroll(onscroll func(Event))

	// // Onseeked prop Occurs when the seek operation ends.
	// //     * @param ev The event.
	// // js:"onseeked,rewrite=onseeked"
	// Onseeked() (onseeked func(Event))

	// // Onseeked prop Occurs when the seek operation ends.
	// //     * @param ev The event.
	// // js:"setonseeked,rewrite=setonseeked"
	// SetOnseeked(onseeked func(Event))

	// // Onseeking prop Occurs when the current playback position is moved.
	// //     * @param ev The event.
	// // js:"onseeking,rewrite=onseeking"
	// Onseeking() (onseeking func(Event))

	// // Onseeking prop Occurs when the current playback position is moved.
	// //     * @param ev The event.
	// // js:"setonseeking,rewrite=setonseeking"
	// SetOnseeking(onseeking func(Event))

	// // Onselect prop Fires when the current selection changes.
	// //     * @param ev The event.
	// // js:"onselect,rewrite=onselect"
	// Onselect() (onselect func(Event))

	// // Onselect prop Fires when the current selection changes.
	// //     * @param ev The event.
	// // js:"setonselect,rewrite=setonselect"
	// SetOnselect(onselect func(Event))

	// // Onselectionchange prop Fires when the selection state of a document changes.
	// //     * @param ev The event.
	// // js:"onselectionchange,rewrite=onselectionchange"
	// Onselectionchange() (onselectionchange func(Event))

	// // Onselectionchange prop Fires when the selection state of a document changes.
	// //     * @param ev The event.
	// // js:"setonselectionchange,rewrite=setonselectionchange"
	// SetOnselectionchange(onselectionchange func(Event))

	// // Onselectstart prop
	// // js:"onselectstart,rewrite=onselectstart"
	// Onselectstart() (onselectstart func(Event))

	// // Onselectstart prop
	// // js:"setonselectstart,rewrite=setonselectstart"
	// SetOnselectstart(onselectstart func(Event))

	// // Onstalled prop Occurs when the download has stopped.
	// //     * @param ev The event.
	// // js:"onstalled,rewrite=onstalled"
	// Onstalled() (onstalled func(Event))

	// // Onstalled prop Occurs when the download has stopped.
	// //     * @param ev The event.
	// // js:"setonstalled,rewrite=setonstalled"
	// SetOnstalled(onstalled func(Event))

	// // Onstop prop Fires when the user clicks the Stop button or leaves the Web page.
	// //     * @param ev The event.
	// // js:"onstop,rewrite=onstop"
	// Onstop() (onstop func(Event))

	// // Onstop prop Fires when the user clicks the Stop button or leaves the Web page.
	// //     * @param ev The event.
	// // js:"setonstop,rewrite=setonstop"
	// SetOnstop(onstop func(Event))

	// // Onsubmit prop
	// // js:"onsubmit,rewrite=onsubmit"
	// Onsubmit() (onsubmit func(Event))

	// // Onsubmit prop
	// // js:"setonsubmit,rewrite=setonsubmit"
	// SetOnsubmit(onsubmit func(Event))

	// // Onsuspend prop Occurs if the load operation has been intentionally halted.
	// //     * @param ev The event.
	// // js:"onsuspend,rewrite=onsuspend"
	// Onsuspend() (onsuspend func(Event))

	// // Onsuspend prop Occurs if the load operation has been intentionally halted.
	// //     * @param ev The event.
	// // js:"setonsuspend,rewrite=setonsuspend"
	// SetOnsuspend(onsuspend func(Event))

	// // Ontimeupdate prop Occurs to indicate the current playback position.
	// //     * @param ev The event.
	// // js:"ontimeupdate,rewrite=ontimeupdate"
	// Ontimeupdate() (ontimeupdate func(Event))

	// // Ontimeupdate prop Occurs to indicate the current playback position.
	// //     * @param ev The event.
	// // js:"setontimeupdate,rewrite=setontimeupdate"
	// SetOntimeupdate(ontimeupdate func(Event))

	// // Ontouchcancel prop
	// // js:"ontouchcancel,rewrite=ontouchcancel"
	// Ontouchcancel() (ontouchcancel func(Event))

	// // Ontouchcancel prop
	// // js:"setontouchcancel,rewrite=setontouchcancel"
	// SetOntouchcancel(ontouchcancel func(Event))

	// // Ontouchend prop
	// // js:"ontouchend,rewrite=ontouchend"
	// Ontouchend() (ontouchend func(Event))

	// // Ontouchend prop
	// // js:"setontouchend,rewrite=setontouchend"
	// SetOntouchend(ontouchend func(Event))

	// // Ontouchmove prop
	// // js:"ontouchmove,rewrite=ontouchmove"
	// Ontouchmove() (ontouchmove func(Event))

	// // Ontouchmove prop
	// // js:"setontouchmove,rewrite=setontouchmove"
	// SetOntouchmove(ontouchmove func(Event))

	// // Ontouchstart prop
	// // js:"ontouchstart,rewrite=ontouchstart"
	// Ontouchstart() (ontouchstart func(Event))

	// // Ontouchstart prop
	// // js:"setontouchstart,rewrite=setontouchstart"
	// SetOntouchstart(ontouchstart func(Event))

	// // Onvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
	// //     * @param ev The event.
	// // js:"onvolumechange,rewrite=onvolumechange"
	// Onvolumechange() (onvolumechange func(Event))

	// // Onvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
	// //     * @param ev The event.
	// // js:"setonvolumechange,rewrite=setonvolumechange"
	// SetOnvolumechange(onvolumechange func(Event))

	// // Onwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
	// //     * @param ev The event.
	// // js:"onwaiting,rewrite=onwaiting"
	// Onwaiting() (onwaiting func(Event))

	// // Onwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
	// //     * @param ev The event.
	// // js:"setonwaiting,rewrite=setonwaiting"
	// SetOnwaiting(onwaiting func(Event))

	// // Onwebkitfullscreenchange prop
	// // js:"onwebkitfullscreenchange,rewrite=onwebkitfullscreenchange"
	// Onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event))

	// // Onwebkitfullscreenchange prop
	// // js:"setonwebkitfullscreenchange,rewrite=setonwebkitfullscreenchange"
	// SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(Event))

	// // Onwebkitfullscreenerror prop
	// // js:"onwebkitfullscreenerror,rewrite=onwebkitfullscreenerror"
	// Onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event))

	// // Onwebkitfullscreenerror prop
	// // js:"setonwebkitfullscreenerror,rewrite=setonwebkitfullscreenerror"
	// SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(Event))

	// // Plugins prop
	// // js:"plugins,rewrite=plugins"
	// Plugins() (plugins HTMLCollection)

	// // PointerLockElement prop
	// // js:"pointerLockElement,rewrite=pointerlockelement"
	// PointerLockElement() (pointerLockElement Element)

	// // ReadyState prop Retrieves a value that indicates the current state of the object.
	// // js:"readyState,rewrite=readystate"
	// ReadyState() (readyState string)

	// // Referrer prop Gets the URL of the location that referred the user to the current page.
	// // js:"referrer,rewrite=referrer"
	// Referrer() (referrer string)

	// // RootElement prop Gets the root svg element in the document hierarchy.
	// // js:"rootElement,rewrite=rootelement"
	// RootElement() (rootElement *SVGSVGElement)

	// // Scripts prop Retrieves a collection of all script objects in the document.
	// // js:"scripts,rewrite=scripts"
	// Scripts() (scripts HTMLCollection)

	// // ScrollingElement prop
	// // js:"scrollingElement,rewrite=scrollingelement"
	// ScrollingElement() (scrollingElement Element)

	// // StyleSheets prop Retrieves a collection of styleSheet objects representing the style sheets that correspond to each instance of a link or style object in the document.
	// // js:"styleSheets,rewrite=stylesheets"
	// StyleSheets() (styleSheets *StyleSheetList)

	// // Title prop Contains the title of the document.
	// // js:"title,rewrite=title"
	// Title() (title string)

	// // Title prop Contains the title of the document.
	// // js:"settitle,rewrite=settitle"
	// SetTitle(title string)

	// // URL prop Sets or gets the URL for the current document.
	// // js:"URL,rewrite=url"
	// URL() (URL string)

	// // URLUnencoded prop Gets the URL for the document, stripped of any character encoding.
	// // js:"URLUnencoded,rewrite=urlunencoded"
	// URLUnencoded() (URLUnencoded string)

	// // VisibilityState prop
	// // js:"visibilityState,rewrite=visibilitystate"
	// // VisibilityState() (visibilityState *visibilitystate.VisibilityState)

	// // VlinkColor prop Sets or gets the color of the links that the user has visited.
	// // js:"vlinkColor,rewrite=vlinkcolor"
	// VlinkColor() (vlinkColor string)

	// // VlinkColor prop Sets or gets the color of the links that the user has visited.
	// // js:"setvlinkColor,rewrite=setvlinkcolor"
	// SetVlinkColor(vlinkColor string)

	// // WebkitCurrentFullScreenElement prop
	// // js:"webkitCurrentFullScreenElement,rewrite=webkitcurrentfullscreenelement"
	// WebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement Element)

	// // WebkitFullscreenElement prop
	// // js:"webkitFullscreenElement,rewrite=webkitfullscreenelement"
	// WebkitFullscreenElement() (webkitFullscreenElement Element)

	// // WebkitFullscreenEnabled prop
	// // js:"webkitFullscreenEnabled,rewrite=webkitfullscreenenabled"
	// WebkitFullscreenEnabled() (webkitFullscreenEnabled bool)

	// // WebkitIsFullScreen prop
	// // js:"webkitIsFullScreen,rewrite=webkitisfullscreen"
	// WebkitIsFullScreen() (webkitIsFullScreen bool)

	// // XMLEncoding prop
	// // js:"xmlEncoding,rewrite=xmlencoding"
	// XMLEncoding() (xmlEncoding string)

	// // XMLStandalone prop
	// // js:"xmlStandalone,rewrite=xmlstandalone"
	// XMLStandalone() (xmlStandalone bool)

	// // XMLStandalone prop
	// // js:"setxmlStandalone,rewrite=setxmlstandalone"
	// SetXMLStandalone(xmlStandalone bool)

	// // XMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
	// // js:"xmlVersion,rewrite=xmlversion"
	// XMLVersion() (xmlVersion string)

	// // XMLVersion prop Gets or sets the version attribute specified in the declaration of an XML document.
	// // js:"setxmlVersion,rewrite=setxmlversion"
	// SetXMLVersion(xmlVersion string)
}

func queryselector(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func queryselectorall(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func createevent(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointercancel(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointercancel(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointerdown(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointerdown(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointerenter(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointerenter(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointermove(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointermove(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointerleave(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointerleave(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointerout(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointerout(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointerover(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointerover(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onpointerup(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func setonpointerup(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

func onwheel(selector string) (el Element) {
	js.Rewrite("$<.querySelector($1)", selector)
	return el
}

// adoptnode fn
func adoptnode(source Node) (n Node) {
	js.Rewrite("$<.adoptNode($1)", source)
	return n
}

// captureevents fn
func captureevents() {
	js.Rewrite("$<.captureEvents()")
}

// caretrangefrompoint fn
func caretrangefrompoint(x float32, y float32) (r *Range) {
	js.Rewrite("$<.caretRangeFromPoint($1, $2)", x, y)
	return r
}

// clear fn
func clear() {
	js.Rewrite("$<.clear()")
}

// cls fn Closes an output stream and forces the sent data to display.
func cls() {
	js.Rewrite("$<.close()")
}

// createattribute fn Creates an attribute object with a specified name.
//     * @param name String that sets the attribute object's name.
func createattribute(name string) (a *Attr) {
	js.Rewrite("$<.createAttribute($1)", name)
	return a
}

// createattributens fn
func createattributens(namespaceURI string, qualifiedName string) (a *Attr) {
	js.Rewrite("$<.createAttributeNS($1, $2)", namespaceURI, qualifiedName)
	return a
}

// createcdatasection fn
func createcdatasection(data string) (c *CDATASection) {
	js.Rewrite("$<.createCDATASection($1)", data)
	return c
}

// createcomment fn Creates a comment object with the specified data.
//     * @param data Sets the comment object's data.
func createcomment(data string) (c *Comment) {
	js.Rewrite("$<.createComment($1)", data)
	return c
}

// createdocumentfragment fn Creates a new document.
func createdocumentfragment() (d *DocumentFragment) {
	js.Rewrite("$<.createDocumentFragment()")
	return d
}

// createelement fn Creates an instance of the element for the specified tag.
//     * @param tagName The name of an element.
func createelement(tagName string) (e Element) {
	js.Rewrite("$<.createElement($1)", tagName)
	return e
}

// createelementns fn
func createelementns(namespaceURI string, qualifiedName string) (e Element) {
	js.Rewrite("$<.createElementNS($1, $2)", namespaceURI, qualifiedName)
	return e
}

// createexpression fn
func createexpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (x *XPathExpression) {
	js.Rewrite("$<.createExpression($1, $2)", expression, resolver)
	return x
}

// createnodeiterator fn Creates a NodeIterator object that you can use to traverse filtered lists of nodes or elements in a document.
//     * @param root The root element or node to start traversing on.
//     * @param whatToShow The type of nodes or elements to appear in the node list
//     * @param filter A custom NodeFilter function to use. For more information, see filter. Use null for no filter.
//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
func createnodeiterator(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (n *NodeIterator) {
	js.Rewrite("$<.createNodeIterator($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return n
}

// creatensresolver fn
func creatensresolver(nodeResolver Node) (x *xpathnsresolver.XPathNSResolver) {
	js.Rewrite("$<.createNSResolver($1)", nodeResolver)
	return x
}

// createprocessinginstruction fn
func createprocessinginstruction(target string, data string) (p *ProcessingInstruction) {
	js.Rewrite("$<.createProcessingInstruction($1, $2)", target, data)
	return p
}

// createrange fn Returns an empty range object that has both of its boundary points positioned at the beginning of the document.
func createrange() (r *Range) {
	js.Rewrite("$<.createRange()")
	return r
}

// createtextnode fn Creates a text string from the specified value.
//     * @param data String that specifies the nodeValue property of the text node.
func createtextnode(data string) (t Text) {
	js.Rewrite("$<.createTextNode($1)", data)
	return t
}

// createtouch fn
func createtouch(view *Window, target EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *Touch) {
	js.Rewrite("$<.createTouch($1, $2, $3, $4, $5, $6, $7)", view, target, identifier, pageX, pageY, screenX, screenY)
	return t
}

// createtouchlist fn
func createtouchlist(touches *Touch) (t *TouchList) {
	js.Rewrite("$<.createTouchList($1)", touches)
	return t
}

// createtreewalker fn Creates a TreeWalker object that you can use to traverse filtered lists of nodes or elements in a document.
//     * @param root The root element or node to start traversing on.
//     * @param whatToShow The type of nodes or elements to appear in the node list. For more information, see whatToShow.
//     * @param filter A custom NodeFilter function to use.
//     * @param entityReferenceExpansion A flag that specifies whether entity reference nodes are expanded.
func createtreewalker(root Node, whatToShow *uint, filter *NodeFilter, entityReferenceExpansion *bool) (t *TreeWalker) {
	js.Rewrite("$<.createTreeWalker($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return t
}

// elementfrompoint fn Returns the element for the specified x coordinate and the specified y coordinate.
//     * @param x The x-offset
//     * @param y The y-offset
func elementfrompoint(x int, y int) (e Element) {
	js.Rewrite("$<.elementFromPoint($1, $2)", x, y)
	return e
}

// evaluate fn
func evaluate(expression string, contextNode Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *XPathResult) (x *XPathResult) {
	js.Rewrite("$<.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return x
}

// execcommand fn Executes a command on the current document, current selection, or the given range.
//     * @param commandId String that specifies the command to execute. This command can be any of the command identifiers that can be executed in script.
//     * @param showUI Display the user interface, defaults to false.
//     * @param value Value to assign.
func execcommand(commandId string, showUI *bool, value *interface{}) (b bool) {
	js.Rewrite("$<.execCommand($1, $2, $3)", commandId, showUI, value)
	return b
}

// execcommandshowhelp fn Displays help information for the given command identifier.
//     * @param commandId Displays help information for the given command identifier.
func execcommandshowhelp(commandId string) (b bool) {
	js.Rewrite("$<.execCommandShowHelp($1)", commandId)
	return b
}

// exitfullscreen fn
func exitfullscreen() {
	js.Rewrite("$<.exitFullscreen()")
}

// exitpointerlock fn
func exitpointerlock() {
	js.Rewrite("$<.exitPointerLock()")
}

// focus fn Causes the element to receive the focus and executes the code specified by the onfocus event.
func focus() {
	js.Rewrite("$<.focus()")
}

// getelementbyid fn Returns a reference to the first object with the specified value of the ID or NAME attribute.
//     * @param elementId String that specifies the ID value. Case-insensitive.
func getelementbyid(elementId string) (e Element) {
	js.Rewrite("$<.getElementById($1)", elementId)
	return e
}

// getelementsbyclassname fn
func getelementsbyclassname(classNames string) (n *NodeList) {
	js.Rewrite("$<.getElementsByClassName($1)", classNames)
	return n
}

// getelementsbyname fn Gets a collection of objects based on the value of the NAME or ID attribute.
//     * @param elementName Gets a collection of objects based on the value of the NAME or ID attribute.
func getelementsbyname(elementName string) (n *NodeList) {
	js.Rewrite("$<.getElementsByName($1)", elementName)
	return n
}

// getelementsbytagname fn Retrieves a collection of objects based on the specified element name.
//     * @param name Specifies the name of an element.
func getelementsbytagname(tagname string) (n *NodeList) {
	js.Rewrite("$<.getElementsByTagName($1)", tagname)
	return n
}

// getelementsbytagnamens fn
func getelementsbytagnamens(namespaceURI string, localName string) (n *NodeList) {
	js.Rewrite("$<.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

// getselection fn Returns an object representing the current selection of the document that is loaded into the object displaying a webpage.
func getselection() (s *Selection) {
	js.Rewrite("$<.getSelection()")
	return s
}

// hasfocus fn Gets a value indicating whether the object currently has focus.
func hasfocus() (b bool) {
	js.Rewrite("$<.hasFocus()")
	return b
}

// importnode fn
func importnode(importedNode Node, deep bool) (n Node) {
	js.Rewrite("$<.importNode($1, $2)", importedNode, deep)
	return n
}

// mselementsfrompoint fn
func mselementsfrompoint(x float32, y float32) (n *NodeList) {
	js.Rewrite("$<.msElementsFromPoint($1, $2)", x, y)
	return n
}

// mselementsfromrect fn
func mselementsfromrect(left float32, top float32, width float32, height float32) (n *NodeList) {
	js.Rewrite("$<.msElementsFromRect($1, $2, $3, $4)", left, top, width, height)
	return n
}

// open fn Opens a new window and loads a document specified by a given URL. Also, opens a new window that uses the url parameter and the name parameter to collect the output of the write method and the writeln method.
//     * @param url Specifies a MIME type for the document.
//     * @param name Specifies the name of the window. This name is used as the value for the TARGET attribute on a form or an anchor element.
//     * @param features Contains a list of items separated by commas. Each item consists of an option and a value, separated by an equals sign (for example, "fullscreen=yes, toolbar=yes"). The following values are supported.
//     * @param replace Specifies whether the existing entry for the document is replaced in the history list.
func open(url *string, name *string, features *string, replace *bool) (i interface{}) {
	js.Rewrite("$<.open($1, $2, $3, $4)", url, name, features, replace)
	return i
}

// querycommandenabled fn Returns a Boolean value that indicates whether a specified command can be successfully executed using execCommand, given the current state of the document.
//     * @param commandId Specifies a command identifier.
func querycommandenabled(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandEnabled($1)", commandId)
	return b
}

// querycommandindeterm fn Returns a Boolean value that indicates whether the specified command is in the indeterminate state.
//     * @param commandId String that specifies a command identifier.
func querycommandindeterm(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandIndeterm($1)", commandId)
	return b
}

// querycommandstate fn Returns a Boolean value that indicates the current state of the command.
//     * @param commandId String that specifies a command identifier.
func querycommandstate(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandState($1)", commandId)
	return b
}

// querycommandsupported fn Returns a Boolean value that indicates whether the current command is supported on the current range.
//     * @param commandId Specifies a command identifier.
func querycommandsupported(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandSupported($1)", commandId)
	return b
}

// querycommandtext fn Retrieves the string associated with a command.
//     * @param commandId String that contains the identifier of a command. This can be any command identifier given in the list of Command Identifiers.
func querycommandtext(commandId string) (s string) {
	js.Rewrite("$<.queryCommandText($1)", commandId)
	return s
}

// querycommandvalue fn Returns the current value of the document, range, or current selection for the given command.
//     * @param commandId String that specifies a command identifier.
func querycommandvalue(commandId string) (s string) {
	js.Rewrite("$<.queryCommandValue($1)", commandId)
	return s
}

// releaseevents fn
func releaseevents() {
	js.Rewrite("$<.releaseEvents()")
}

// updatesettings fn Allows updating the print settings for the page.
func updatesettings() {
	js.Rewrite("$<.updateSettings()")
}

// webkitcancelfullscreen fn
func webkitcancelfullscreen() {
	js.Rewrite("$<.webkitCancelFullScreen()")
}

// webkitexitfullscreen fn
func webkitexitfullscreen() {
	js.Rewrite("$<.webkitExitFullscreen()")
}

// write fn Writes one or more HTML expressions to a document in the specified window.
//     * @param content Specifies the text and HTML tags to write.
func write(content string) {
	js.Rewrite("$<.write($1)", content)
}

// writeln fn Writes one or more HTML expressions, followed by a carriage return, to a document in the specified window.
//     * @param content The text and HTML tags to write.
func writeln(content string) {
	js.Rewrite("$<.writeln($1)", content)
}

// activeelement prop Gets the object that has the focus when the parent document has focus.
func activeelement() (activeElement Element) {
	js.Rewrite("$<.activeElement")
	return activeElement
}

// alinkcolor prop Sets or gets the color of all active links in the document.
func alinkcolor() (alinkColor string) {
	js.Rewrite("$<.alinkColor")
	return alinkColor
}

// setalinkcolor prop Sets or gets the color of all active links in the document.
func setalinkcolor(alinkColor string) {
	js.Rewrite("$<.alinkColor = alinkColor")
}

// all prop Returns a reference to the collection of elements contained by the object.
func all() (all *htmlallcollection.HTMLAllCollection) {
	js.Rewrite("$<.all")
	return all
}

// anchors prop Retrieves a collection of all a objects that have a name and/or id property. Objects in this collection are in HTML source order.
func anchors() (anchors HTMLCollection) {
	js.Rewrite("$<.anchors")
	return anchors
}

// applets prop Retrieves a collection of all applet objects in the document.
func applets() (applets HTMLCollection) {
	js.Rewrite("$<.applets")
	return applets
}

// bgcolor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
func bgcolor() (bgColor string) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

// setbgcolor prop Deprecated. Sets or retrieves a value that indicates the background color behind the object.
func setbgcolor(bgColor string) {
	js.Rewrite("$<.bgColor = bgColor")
}

// body prop Specifies the beginning and end of the document body.
func body() (body HTMLElement) {
	js.Rewrite("$<.body")
	return body
}

// setbody prop Specifies the beginning and end of the document body.
func setbody(body HTMLElement) {
	js.Rewrite("$<.body = body")
}

// characterset prop
func characterset() (characterSet string) {
	js.Rewrite("$<.characterSet")
	return characterSet
}

// charset prop Gets or sets the character set used to encode the object.
func charset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

// setcharset prop Gets or sets the character set used to encode the object.
func setcharset(charset string) {
	js.Rewrite("$<.charset = charset")
}

// compatmode prop Gets a value that indicates whether standards-compliant mode is switched on for the object.
func compatmode() (compatMode string) {
	js.Rewrite("$<.compatMode")
	return compatMode
}

// cookie prop
func cookie() (cookie string) {
	js.Rewrite("$<.cookie")
	return cookie
}

// setcookie prop
func setcookie(cookie string) {
	js.Rewrite("$<.cookie = cookie")
}

// currentscript prop
func currentscript() (currentScript interface{}) {
	js.Rewrite("$<.currentScript")
	return currentScript
}

// defaultview prop
func defaultview() (defaultView *Window) {
	js.Rewrite("$<.defaultView")
	return defaultView
}

// designmode prop Sets or gets a value that indicates whether the document can be edited.
func designmode() (designMode string) {
	js.Rewrite("$<.designMode")
	return designMode
}

// setdesignmode prop Sets or gets a value that indicates whether the document can be edited.
func setdesignmode(designMode string) {
	js.Rewrite("$<.designMode = designMode")
}

// dir prop Sets or retrieves a value that indicates the reading order of the object.
func dir() (dir string) {
	js.Rewrite("$<.dir")
	return dir
}

// setdir prop Sets or retrieves a value that indicates the reading order of the object.
func setdir(dir string) {
	js.Rewrite("$<.dir = dir")
}

// doctype prop Gets an object representing the document type declaration associated with the current document.
func doctype() (doctype *DocumentType) {
	js.Rewrite("$<.doctype")
	return doctype
}

// documentelement prop Gets a reference to the root node of the document.
func documentelement() (documentElement Element) {
	js.Rewrite("$<.documentElement")
	return documentElement
}

// domain prop Sets or gets the security domain of the document.
func domain() (domain string) {
	js.Rewrite("$<.domain")
	return domain
}

// setdomain prop Sets or gets the security domain of the document.
func setdomain(domain string) {
	js.Rewrite("$<.domain = domain")
}

// embeds prop Retrieves a collection of all embed objects in the document.
func embeds() (embeds HTMLCollection) {
	js.Rewrite("$<.embeds")
	return embeds
}

// fgcolor prop Sets or gets the foreground (text) color of the document.
func fgcolor() (fgColor string) {
	js.Rewrite("$<.fgColor")
	return fgColor
}

// setfgcolor prop Sets or gets the foreground (text) color of the document.
func setfgcolor(fgColor string) {
	js.Rewrite("$<.fgColor = fgColor")
}

// forms prop Retrieves a collection, in source order, of all form objects in the document.
func forms() (forms HTMLCollection) {
	js.Rewrite("$<.forms")
	return forms
}

// fullscreenelement prop
func fullscreenelement() (fullscreenElement Element) {
	js.Rewrite("$<.fullscreenElement")
	return fullscreenElement
}

// fullscreenenabled prop
func fullscreenenabled() (fullscreenEnabled bool) {
	js.Rewrite("$<.fullscreenEnabled")
	return fullscreenEnabled
}

// head prop
func head() (head *HTMLHeadElement) {
	js.Rewrite("$<.head")
	return head
}

// hidden prop
func hidden() (hidden bool) {
	js.Rewrite("$<.hidden")
	return hidden
}

// images prop Retrieves a collection, in source order, of img objects in the document.
func images() (images HTMLCollection) {
	js.Rewrite("$<.images")
	return images
}

// implementation prop Gets the implementation object of the current document.
func implementation() (implementation *DOMImplementation) {
	js.Rewrite("$<.implementation")
	return implementation
}

// inputencoding prop Returns the character encoding used to create the webpage that is loaded into the document object.
func inputencoding() (inputEncoding string) {
	js.Rewrite("$<.inputEncoding")
	return inputEncoding
}

// lastmodified prop Gets the date that the page was last modified, if the page supplies one.
func lastmodified() (lastModified string) {
	js.Rewrite("$<.lastModified")
	return lastModified
}

// linkcolor prop Sets or gets the color of the document links.
func linkcolor() (linkColor string) {
	js.Rewrite("$<.linkColor")
	return linkColor
}

// setlinkcolor prop Sets or gets the color of the document links.
func setlinkcolor(linkColor string) {
	js.Rewrite("$<.linkColor = linkColor")
}

// links prop Retrieves a collection of all a objects that specify the href property and all area objects in the document.
func links() (links HTMLCollection) {
	js.Rewrite("$<.links")
	return links
}

// // location prop Contains information about the current URL.
// func location() (location *location.Location) {
// 	js.Rewrite("$<.location")
// 	return location
// }

// mscapslockwarningoff prop
func mscapslockwarningoff() (msCapsLockWarningOff bool) {
	js.Rewrite("$<.msCapsLockWarningOff")
	return msCapsLockWarningOff
}

// setmscapslockwarningoff prop
func setmscapslockwarningoff(msCapsLockWarningOff bool) {
	js.Rewrite("$<.msCapsLockWarningOff = msCapsLockWarningOff")
}

// mscssomelementfloatmetrics prop
func mscssomelementfloatmetrics() (msCSSOMElementFloatMetrics bool) {
	js.Rewrite("$<.msCSSOMElementFloatMetrics")
	return msCSSOMElementFloatMetrics
}

// setmscssomelementfloatmetrics prop
func setmscssomelementfloatmetrics(msCSSOMElementFloatMetrics bool) {
	js.Rewrite("$<.msCSSOMElementFloatMetrics = msCSSOMElementFloatMetrics")
}

// onabort prop Fires when the user aborts the download.
//     * @param ev The event.
func onabort() (onabort func(Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// setonabort prop Fires when the user aborts the download.
//     * @param ev The event.
func setonabort(onabort func(Event)) {
	js.Rewrite("$<.onabort = onabort")
}

// onactivate prop Fires when the object is set as the active element.
//     * @param ev The event.
func onactivate() (onactivate func(Event)) {
	js.Rewrite("$<.onactivate")
	return onactivate
}

// setonactivate prop Fires when the object is set as the active element.
//     * @param ev The event.
func setonactivate(onactivate func(Event)) {
	js.Rewrite("$<.onactivate = onactivate")
}

// onbeforeactivate prop Fires immediately before the object is set as the active element.
//     * @param ev The event.
func onbeforeactivate() (onbeforeactivate func(Event)) {
	js.Rewrite("$<.onbeforeactivate")
	return onbeforeactivate
}

// setonbeforeactivate prop Fires immediately before the object is set as the active element.
//     * @param ev The event.
func setonbeforeactivate(onbeforeactivate func(Event)) {
	js.Rewrite("$<.onbeforeactivate = onbeforeactivate")
}

// onbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The event.
func onbeforedeactivate() (onbeforedeactivate func(Event)) {
	js.Rewrite("$<.onbeforedeactivate")
	return onbeforedeactivate
}

// setonbeforedeactivate prop Fires immediately before the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The event.
func setonbeforedeactivate(onbeforedeactivate func(Event)) {
	js.Rewrite("$<.onbeforedeactivate = onbeforedeactivate")
}

// onblur prop Fires when the object loses the input focus.
//     * @param ev The focus event.
func onblur() (onblur func(Event)) {
	js.Rewrite("$<.onblur")
	return onblur
}

// setonblur prop Fires when the object loses the input focus.
//     * @param ev The focus event.
func setonblur(onblur func(Event)) {
	js.Rewrite("$<.onblur = onblur")
}

// oncanplay prop Occurs when playback is possible, but would require further buffering.
//     * @param ev The event.
func oncanplay() (oncanplay func(Event)) {
	js.Rewrite("$<.oncanplay")
	return oncanplay
}

// setoncanplay prop Occurs when playback is possible, but would require further buffering.
//     * @param ev The event.
func setoncanplay(oncanplay func(Event)) {
	js.Rewrite("$<.oncanplay = oncanplay")
}

// oncanplaythrough prop
func oncanplaythrough() (oncanplaythrough func(Event)) {
	js.Rewrite("$<.oncanplaythrough")
	return oncanplaythrough
}

// setoncanplaythrough prop
func setoncanplaythrough(oncanplaythrough func(Event)) {
	js.Rewrite("$<.oncanplaythrough = oncanplaythrough")
}

// onchange prop Fires when the contents of the object or selection have changed.
//     * @param ev The event.
func onchange() (onchange func(Event)) {
	js.Rewrite("$<.onchange")
	return onchange
}

// setonchange prop Fires when the contents of the object or selection have changed.
//     * @param ev The event.
func setonchange(onchange func(Event)) {
	js.Rewrite("$<.onchange = onchange")
}

// onclick prop Fires when the user clicks the left mouse button on the object
//     * @param ev The mouse event.
func onclick() (onclick func(Event)) {
	js.Rewrite("$<.onclick")
	return onclick
}

// setonclick prop Fires when the user clicks the left mouse button on the object
//     * @param ev The mouse event.
func setonclick(onclick func(Event)) {
	js.Rewrite("$<.onclick = onclick")
}

// oncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
//     * @param ev The mouse event.
func oncontextmenu() (oncontextmenu func(Event)) {
	js.Rewrite("$<.oncontextmenu")
	return oncontextmenu
}

// setoncontextmenu prop Fires when the user clicks the right mouse button in the client area, opening the context menu.
//     * @param ev The mouse event.
func setoncontextmenu(oncontextmenu func(Event)) {
	js.Rewrite("$<.oncontextmenu = oncontextmenu")
}

// ondblclick prop Fires when the user double-clicks the object.
//     * @param ev The mouse event.
func ondblclick() (ondblclick func(Event)) {
	js.Rewrite("$<.ondblclick")
	return ondblclick
}

// setondblclick prop Fires when the user double-clicks the object.
//     * @param ev The mouse event.
func setondblclick(ondblclick func(Event)) {
	js.Rewrite("$<.ondblclick = ondblclick")
}

// ondeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The UI Event
func ondeactivate() (ondeactivate func(Event)) {
	js.Rewrite("$<.ondeactivate")
	return ondeactivate
}

// setondeactivate prop Fires when the activeElement is changed from the current object to another object in the parent document.
//     * @param ev The UI Event
func setondeactivate(ondeactivate func(Event)) {
	js.Rewrite("$<.ondeactivate = ondeactivate")
}

// ondrag prop Fires on the source object continuously during a drag operation.
//     * @param ev The event.
func ondrag() (ondrag func(Event)) {
	js.Rewrite("$<.ondrag")
	return ondrag
}

// setondrag prop Fires on the source object continuously during a drag operation.
//     * @param ev The event.
func setondrag(ondrag func(Event)) {
	js.Rewrite("$<.ondrag = ondrag")
}

// ondragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
//     * @param ev The event.
func ondragend() (ondragend func(Event)) {
	js.Rewrite("$<.ondragend")
	return ondragend
}

// setondragend prop Fires on the source object when the user releases the mouse at the close of a drag operation.
//     * @param ev The event.
func setondragend(ondragend func(Event)) {
	js.Rewrite("$<.ondragend = ondragend")
}

// ondragenter prop Fires on the target element when the user drags the object to a valid drop target.
//     * @param ev The drag event.
func ondragenter() (ondragenter func(Event)) {
	js.Rewrite("$<.ondragenter")
	return ondragenter
}

// setondragenter prop Fires on the target element when the user drags the object to a valid drop target.
//     * @param ev The drag event.
func setondragenter(ondragenter func(Event)) {
	js.Rewrite("$<.ondragenter = ondragenter")
}

// ondragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
//     * @param ev The drag event.
func ondragleave() (ondragleave func(Event)) {
	js.Rewrite("$<.ondragleave")
	return ondragleave
}

// setondragleave prop Fires on the target object when the user moves the mouse out of a valid drop target during a drag operation.
//     * @param ev The drag event.
func setondragleave(ondragleave func(Event)) {
	js.Rewrite("$<.ondragleave = ondragleave")
}

// ondragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
//     * @param ev The event.
func ondragover() (ondragover func(Event)) {
	js.Rewrite("$<.ondragover")
	return ondragover
}

// setondragover prop Fires on the target element continuously while the user drags the object over a valid drop target.
//     * @param ev The event.
func setondragover(ondragover func(Event)) {
	js.Rewrite("$<.ondragover = ondragover")
}

// ondragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
//     * @param ev The event.
func ondragstart() (ondragstart func(Event)) {
	js.Rewrite("$<.ondragstart")
	return ondragstart
}

// setondragstart prop Fires on the source object when the user starts to drag a text selection or selected object.
//     * @param ev The event.
func setondragstart(ondragstart func(Event)) {
	js.Rewrite("$<.ondragstart = ondragstart")
}

// ondrop prop
func ondrop() (ondrop func(Event)) {
	js.Rewrite("$<.ondrop")
	return ondrop
}

// setondrop prop
func setondrop(ondrop func(Event)) {
	js.Rewrite("$<.ondrop = ondrop")
}

// ondurationchange prop Occurs when the duration attribute is updated.
//     * @param ev The event.
func ondurationchange() (ondurationchange func(Event)) {
	js.Rewrite("$<.ondurationchange")
	return ondurationchange
}

// setondurationchange prop Occurs when the duration attribute is updated.
//     * @param ev The event.
func setondurationchange(ondurationchange func(Event)) {
	js.Rewrite("$<.ondurationchange = ondurationchange")
}

// onemptied prop Occurs when the media element is reset to its initial state.
//     * @param ev The event.
func onemptied() (onemptied func(Event)) {
	js.Rewrite("$<.onemptied")
	return onemptied
}

// setonemptied prop Occurs when the media element is reset to its initial state.
//     * @param ev The event.
func setonemptied(onemptied func(Event)) {
	js.Rewrite("$<.onemptied = onemptied")
}

// onended prop Occurs when the end of playback is reached.
//     * @param ev The event
func onended() (onended func(Event)) {
	js.Rewrite("$<.onended")
	return onended
}

// setonended prop Occurs when the end of playback is reached.
//     * @param ev The event
func setonended(onended func(Event)) {
	js.Rewrite("$<.onended = onended")
}

// onerror prop Fires when an error occurs during object loading.
//     * @param ev The event.
func onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// setonerror prop Fires when an error occurs during object loading.
//     * @param ev The event.
func setonerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = onerror")
}

// onfocus prop Fires when the object receives focus.
//     * @param ev The event.
func onfocus() (onfocus func(Event)) {
	js.Rewrite("$<.onfocus")
	return onfocus
}

// setonfocus prop Fires when the object receives focus.
//     * @param ev The event.
func setonfocus(onfocus func(Event)) {
	js.Rewrite("$<.onfocus = onfocus")
}

// onfullscreenchange prop
func onfullscreenchange() (onfullscreenchange func(Event)) {
	js.Rewrite("$<.onfullscreenchange")
	return onfullscreenchange
}

// setonfullscreenchange prop
func setonfullscreenchange(onfullscreenchange func(Event)) {
	js.Rewrite("$<.onfullscreenchange = onfullscreenchange")
}

// onfullscreenerror prop
func onfullscreenerror() (onfullscreenerror func(Event)) {
	js.Rewrite("$<.onfullscreenerror")
	return onfullscreenerror
}

// setonfullscreenerror prop
func setonfullscreenerror(onfullscreenerror func(Event)) {
	js.Rewrite("$<.onfullscreenerror = onfullscreenerror")
}

// oninput prop
func oninput() (oninput func(Event)) {
	js.Rewrite("$<.oninput")
	return oninput
}

// setoninput prop
func setoninput(oninput func(Event)) {
	js.Rewrite("$<.oninput = oninput")
}

// oninvalid prop
func oninvalid() (oninvalid func(Event)) {
	js.Rewrite("$<.oninvalid")
	return oninvalid
}

// setoninvalid prop
func setoninvalid(oninvalid func(Event)) {
	js.Rewrite("$<.oninvalid = oninvalid")
}

// onkeydown prop Fires when the user presses a key.
//     * @param ev The keyboard event
func onkeydown() (onkeydown func(Event)) {
	js.Rewrite("$<.onkeydown")
	return onkeydown
}

// setonkeydown prop Fires when the user presses a key.
//     * @param ev The keyboard event
func setonkeydown(onkeydown func(Event)) {
	js.Rewrite("$<.onkeydown = onkeydown")
}

// onkeypress prop Fires when the user presses an alphanumeric key.
//     * @param ev The event.
func onkeypress() (onkeypress func(Event)) {
	js.Rewrite("$<.onkeypress")
	return onkeypress
}

// setonkeypress prop Fires when the user presses an alphanumeric key.
//     * @param ev The event.
func setonkeypress(onkeypress func(Event)) {
	js.Rewrite("$<.onkeypress = onkeypress")
}

// onkeyup prop Fires when the user releases a key.
//     * @param ev The keyboard event
func onkeyup() (onkeyup func(Event)) {
	js.Rewrite("$<.onkeyup")
	return onkeyup
}

// setonkeyup prop Fires when the user releases a key.
//     * @param ev The keyboard event
func setonkeyup(onkeyup func(Event)) {
	js.Rewrite("$<.onkeyup = onkeyup")
}

// onload prop Fires immediately after the browser loads the object.
//     * @param ev The event.
func onload() (onload func(Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// setonload prop Fires immediately after the browser loads the object.
//     * @param ev The event.
func setonload(onload func(Event)) {
	js.Rewrite("$<.onload = onload")
}

// onloadeddata prop Occurs when media data is loaded at the current playback position.
//     * @param ev The event.
func onloadeddata() (onloadeddata func(Event)) {
	js.Rewrite("$<.onloadeddata")
	return onloadeddata
}

// setonloadeddata prop Occurs when media data is loaded at the current playback position.
//     * @param ev The event.
func setonloadeddata(onloadeddata func(Event)) {
	js.Rewrite("$<.onloadeddata = onloadeddata")
}

// onloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
//     * @param ev The event.
func onloadedmetadata() (onloadedmetadata func(Event)) {
	js.Rewrite("$<.onloadedmetadata")
	return onloadedmetadata
}

// setonloadedmetadata prop Occurs when the duration and dimensions of the media have been determined.
//     * @param ev The event.
func setonloadedmetadata(onloadedmetadata func(Event)) {
	js.Rewrite("$<.onloadedmetadata = onloadedmetadata")
}

// onloadstart prop Occurs when Internet Explorer begins looking for media data.
//     * @param ev The event.
func onloadstart() (onloadstart func(Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// setonloadstart prop Occurs when Internet Explorer begins looking for media data.
//     * @param ev The event.
func setonloadstart(onloadstart func(Event)) {
	js.Rewrite("$<.onloadstart = onloadstart")
}

// onmousedown prop Fires when the user clicks the object with either mouse button.
//     * @param ev The mouse event.
func onmousedown() (onmousedown func(Event)) {
	js.Rewrite("$<.onmousedown")
	return onmousedown
}

// setonmousedown prop Fires when the user clicks the object with either mouse button.
//     * @param ev The mouse event.
func setonmousedown(onmousedown func(Event)) {
	js.Rewrite("$<.onmousedown = onmousedown")
}

// onmousemove prop Fires when the user moves the mouse over the object.
//     * @param ev The mouse event.
func onmousemove() (onmousemove func(Event)) {
	js.Rewrite("$<.onmousemove")
	return onmousemove
}

// setonmousemove prop Fires when the user moves the mouse over the object.
//     * @param ev The mouse event.
func setonmousemove(onmousemove func(Event)) {
	js.Rewrite("$<.onmousemove = onmousemove")
}

// onmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
//     * @param ev The mouse event.
func onmouseout() (onmouseout func(Event)) {
	js.Rewrite("$<.onmouseout")
	return onmouseout
}

// setonmouseout prop Fires when the user moves the mouse pointer outside the boundaries of the object.
//     * @param ev The mouse event.
func setonmouseout(onmouseout func(Event)) {
	js.Rewrite("$<.onmouseout = onmouseout")
}

// onmouseover prop Fires when the user moves the mouse pointer into the object.
//     * @param ev The mouse event.
func onmouseover() (onmouseover func(Event)) {
	js.Rewrite("$<.onmouseover")
	return onmouseover
}

// setonmouseover prop Fires when the user moves the mouse pointer into the object.
//     * @param ev The mouse event.
func setonmouseover(onmouseover func(Event)) {
	js.Rewrite("$<.onmouseover = onmouseover")
}

// onmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
//     * @param ev The mouse event.
func onmouseup() (onmouseup func(Event)) {
	js.Rewrite("$<.onmouseup")
	return onmouseup
}

// setonmouseup prop Fires when the user releases a mouse button while the mouse is over the object.
//     * @param ev The mouse event.
func setonmouseup(onmouseup func(Event)) {
	js.Rewrite("$<.onmouseup = onmouseup")
}

// onmousewheel prop Fires when the wheel button is rotated.
//     * @param ev The mouse event
func onmousewheel() (onmousewheel func(Event)) {
	js.Rewrite("$<.onmousewheel")
	return onmousewheel
}

// setonmousewheel prop Fires when the wheel button is rotated.
//     * @param ev The mouse event
func setonmousewheel(onmousewheel func(Event)) {
	js.Rewrite("$<.onmousewheel = onmousewheel")
}

// onmscontentzoom prop
func onmscontentzoom() (onmscontentzoom func(UIEvent)) {
	js.Rewrite("$<.onmscontentzoom")
	return onmscontentzoom
}

// setonmscontentzoom prop
func setonmscontentzoom(onmscontentzoom func(UIEvent)) {
	js.Rewrite("$<.onmscontentzoom = onmscontentzoom")
}

// onmsgesturechange prop
func onmsgesturechange() (onmsgesturechange func(Event)) {
	js.Rewrite("$<.onmsgesturechange")
	return onmsgesturechange
}

// setonmsgesturechange prop
func setonmsgesturechange(onmsgesturechange func(Event)) {
	js.Rewrite("$<.onmsgesturechange = onmsgesturechange")
}

// onmsgesturedoubletap prop
func onmsgesturedoubletap() (onmsgesturedoubletap func(Event)) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// setonmsgesturedoubletap prop
func setonmsgesturedoubletap(onmsgesturedoubletap func(Event)) {
	js.Rewrite("$<.onmsgesturedoubletap = onmsgesturedoubletap")
}

// onmsgestureend prop
func onmsgestureend() (onmsgestureend func(Event)) {
	js.Rewrite("$<.onmsgestureend")
	return onmsgestureend
}

// setonmsgestureend prop
func setonmsgestureend(onmsgestureend func(Event)) {
	js.Rewrite("$<.onmsgestureend = onmsgestureend")
}

// onmsgesturehold prop
func onmsgesturehold() (onmsgesturehold func(Event)) {
	js.Rewrite("$<.onmsgesturehold")
	return onmsgesturehold
}

// setonmsgesturehold prop
func setonmsgesturehold(onmsgesturehold func(Event)) {
	js.Rewrite("$<.onmsgesturehold = onmsgesturehold")
}

// onmsgesturestart prop
func onmsgesturestart() (onmsgesturestart func(Event)) {
	js.Rewrite("$<.onmsgesturestart")
	return onmsgesturestart
}

// setonmsgesturestart prop
func setonmsgesturestart(onmsgesturestart func(Event)) {
	js.Rewrite("$<.onmsgesturestart = onmsgesturestart")
}

// onmsgesturetap prop
func onmsgesturetap() (onmsgesturetap func(Event)) {
	js.Rewrite("$<.onmsgesturetap")
	return onmsgesturetap
}

// setonmsgesturetap prop
func setonmsgesturetap(onmsgesturetap func(Event)) {
	js.Rewrite("$<.onmsgesturetap = onmsgesturetap")
}

// onmsinertiastart prop
func onmsinertiastart() (onmsinertiastart func(Event)) {
	js.Rewrite("$<.onmsinertiastart")
	return onmsinertiastart
}

// setonmsinertiastart prop
func setonmsinertiastart(onmsinertiastart func(Event)) {
	js.Rewrite("$<.onmsinertiastart = onmsinertiastart")
}

// onmsmanipulationstatechanged prop
func onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent)) {
	js.Rewrite("$<.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// setonmsmanipulationstatechanged prop
func setonmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent)) {
	js.Rewrite("$<.onmsmanipulationstatechanged = onmsmanipulationstatechanged")
}

// onmspointercancel prop
func onmspointercancel() (onmspointercancel func(Event)) {
	js.Rewrite("$<.onmspointercancel")
	return onmspointercancel
}

// setonmspointercancel prop
func setonmspointercancel(onmspointercancel func(Event)) {
	js.Rewrite("$<.onmspointercancel = onmspointercancel")
}

// onmspointerdown prop
func onmspointerdown() (onmspointerdown func(Event)) {
	js.Rewrite("$<.onmspointerdown")
	return onmspointerdown
}

// setonmspointerdown prop
func setonmspointerdown(onmspointerdown func(Event)) {
	js.Rewrite("$<.onmspointerdown = onmspointerdown")
}

// onmspointerenter prop
func onmspointerenter() (onmspointerenter func(Event)) {
	js.Rewrite("$<.onmspointerenter")
	return onmspointerenter
}

// setonmspointerenter prop
func setonmspointerenter(onmspointerenter func(Event)) {
	js.Rewrite("$<.onmspointerenter = onmspointerenter")
}

// onmspointerleave prop
func onmspointerleave() (onmspointerleave func(Event)) {
	js.Rewrite("$<.onmspointerleave")
	return onmspointerleave
}

// setonmspointerleave prop
func setonmspointerleave(onmspointerleave func(Event)) {
	js.Rewrite("$<.onmspointerleave = onmspointerleave")
}

// onmspointermove prop
func onmspointermove() (onmspointermove func(Event)) {
	js.Rewrite("$<.onmspointermove")
	return onmspointermove
}

// setonmspointermove prop
func setonmspointermove(onmspointermove func(Event)) {
	js.Rewrite("$<.onmspointermove = onmspointermove")
}

// onmspointerout prop
func onmspointerout() (onmspointerout func(Event)) {
	js.Rewrite("$<.onmspointerout")
	return onmspointerout
}

// setonmspointerout prop
func setonmspointerout(onmspointerout func(Event)) {
	js.Rewrite("$<.onmspointerout = onmspointerout")
}

// onmspointerover prop
func onmspointerover() (onmspointerover func(Event)) {
	js.Rewrite("$<.onmspointerover")
	return onmspointerover
}

// setonmspointerover prop
func setonmspointerover(onmspointerover func(Event)) {
	js.Rewrite("$<.onmspointerover = onmspointerover")
}

// onmspointerup prop
func onmspointerup() (onmspointerup func(Event)) {
	js.Rewrite("$<.onmspointerup")
	return onmspointerup
}

// setonmspointerup prop
func setonmspointerup(onmspointerup func(Event)) {
	js.Rewrite("$<.onmspointerup = onmspointerup")
}

// onmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
//     * @param ev The event.
func onmssitemodejumplistitemremoved() (onmssitemodejumplistitemremoved func(*MSSiteModeEvent)) {
	js.Rewrite("$<.onmssitemodejumplistitemremoved")
	return onmssitemodejumplistitemremoved
}

// setonmssitemodejumplistitemremoved prop Occurs when an item is removed from a Jump List of a webpage running in Site Mode.
//     * @param ev The event.
func setonmssitemodejumplistitemremoved(onmssitemodejumplistitemremoved func(*MSSiteModeEvent)) {
	js.Rewrite("$<.onmssitemodejumplistitemremoved = onmssitemodejumplistitemremoved")
}

// onmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
//     * @param ev The event.
func onmsthumbnailclick() (onmsthumbnailclick func(*MSSiteModeEvent)) {
	js.Rewrite("$<.onmsthumbnailclick")
	return onmsthumbnailclick
}

// setonmsthumbnailclick prop Occurs when a user clicks a button in a Thumbnail Toolbar of a webpage running in Site Mode.
//     * @param ev The event.
func setonmsthumbnailclick(onmsthumbnailclick func(*MSSiteModeEvent)) {
	js.Rewrite("$<.onmsthumbnailclick = onmsthumbnailclick")
}

// onpause prop Occurs when playback is paused.
//     * @param ev The event.
func onpause() (onpause func(Event)) {
	js.Rewrite("$<.onpause")
	return onpause
}

// setonpause prop Occurs when playback is paused.
//     * @param ev The event.
func setonpause(onpause func(Event)) {
	js.Rewrite("$<.onpause = onpause")
}

// onplay prop Occurs when the play method is requested.
//     * @param ev The event.
func onplay() (onplay func(Event)) {
	js.Rewrite("$<.onplay")
	return onplay
}

// setonplay prop Occurs when the play method is requested.
//     * @param ev The event.
func setonplay(onplay func(Event)) {
	js.Rewrite("$<.onplay = onplay")
}

// onplaying prop Occurs when the audio or video has started playing.
//     * @param ev The event.
func onplaying() (onplaying func(Event)) {
	js.Rewrite("$<.onplaying")
	return onplaying
}

// setonplaying prop Occurs when the audio or video has started playing.
//     * @param ev The event.
func setonplaying(onplaying func(Event)) {
	js.Rewrite("$<.onplaying = onplaying")
}

// onpointerlockchange prop
func onpointerlockchange() (onpointerlockchange func(Event)) {
	js.Rewrite("$<.onpointerlockchange")
	return onpointerlockchange
}

// setonpointerlockchange prop
func setonpointerlockchange(onpointerlockchange func(Event)) {
	js.Rewrite("$<.onpointerlockchange = onpointerlockchange")
}

// onpointerlockerror prop
func onpointerlockerror() (onpointerlockerror func(Event)) {
	js.Rewrite("$<.onpointerlockerror")
	return onpointerlockerror
}

// setonpointerlockerror prop
func setonpointerlockerror(onpointerlockerror func(Event)) {
	js.Rewrite("$<.onpointerlockerror = onpointerlockerror")
}

// onprogress prop Occurs to indicate progress while downloading media data.
//     * @param ev The event.
func onprogress() (onprogress func(Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// setonprogress prop Occurs to indicate progress while downloading media data.
//     * @param ev The event.
func setonprogress(onprogress func(Event)) {
	js.Rewrite("$<.onprogress = onprogress")
}

// onratechange prop Occurs when the playback rate is increased or decreased.
//     * @param ev The event.
func onratechange() (onratechange func(Event)) {
	js.Rewrite("$<.onratechange")
	return onratechange
}

// setonratechange prop Occurs when the playback rate is increased or decreased.
//     * @param ev The event.
func setonratechange(onratechange func(Event)) {
	js.Rewrite("$<.onratechange = onratechange")
}

// onreadystatechange prop Fires when the state of the object has changed.
//     * @param ev The event
func onreadystatechange() (onreadystatechange func(Event)) {
	js.Rewrite("$<.onreadystatechange")
	return onreadystatechange
}

// setonreadystatechange prop Fires when the state of the object has changed.
//     * @param ev The event
func setonreadystatechange(onreadystatechange func(Event)) {
	js.Rewrite("$<.onreadystatechange = onreadystatechange")
}

// onreset prop Fires when the user resets a form.
//     * @param ev The event.
func onreset() (onreset func(Event)) {
	js.Rewrite("$<.onreset")
	return onreset
}

// setonreset prop Fires when the user resets a form.
//     * @param ev The event.
func setonreset(onreset func(Event)) {
	js.Rewrite("$<.onreset = onreset")
}

// onscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
//     * @param ev The event.
func onscroll() (onscroll func(Event)) {
	js.Rewrite("$<.onscroll")
	return onscroll
}

// setonscroll prop Fires when the user repositions the scroll box in the scroll bar on the object.
//     * @param ev The event.
func setonscroll(onscroll func(Event)) {
	js.Rewrite("$<.onscroll = onscroll")
}

// onseeked prop Occurs when the seek operation ends.
//     * @param ev The event.
func onseeked() (onseeked func(Event)) {
	js.Rewrite("$<.onseeked")
	return onseeked
}

// setonseeked prop Occurs when the seek operation ends.
//     * @param ev The event.
func setonseeked(onseeked func(Event)) {
	js.Rewrite("$<.onseeked = onseeked")
}

// onseeking prop Occurs when the current playback position is moved.
//     * @param ev The event.
func onseeking() (onseeking func(Event)) {
	js.Rewrite("$<.onseeking")
	return onseeking
}

// setonseeking prop Occurs when the current playback position is moved.
//     * @param ev The event.
func setonseeking(onseeking func(Event)) {
	js.Rewrite("$<.onseeking = onseeking")
}

// onselect prop Fires when the current selection changes.
//     * @param ev The event.
func onselect() (onselect func(Event)) {
	js.Rewrite("$<.onselect")
	return onselect
}

// setonselect prop Fires when the current selection changes.
//     * @param ev The event.
func setonselect(onselect func(Event)) {
	js.Rewrite("$<.onselect = onselect")
}

// onselectionchange prop Fires when the selection state of a document changes.
//     * @param ev The event.
func onselectionchange() (onselectionchange func(Event)) {
	js.Rewrite("$<.onselectionchange")
	return onselectionchange
}

// setonselectionchange prop Fires when the selection state of a document changes.
//     * @param ev The event.
func setonselectionchange(onselectionchange func(Event)) {
	js.Rewrite("$<.onselectionchange = onselectionchange")
}

// onselectstart prop
func onselectstart() (onselectstart func(Event)) {
	js.Rewrite("$<.onselectstart")
	return onselectstart
}

// setonselectstart prop
func setonselectstart(onselectstart func(Event)) {
	js.Rewrite("$<.onselectstart = onselectstart")
}

// onstalled prop Occurs when the download has stopped.
//     * @param ev The event.
func onstalled() (onstalled func(Event)) {
	js.Rewrite("$<.onstalled")
	return onstalled
}

// setonstalled prop Occurs when the download has stopped.
//     * @param ev The event.
func setonstalled(onstalled func(Event)) {
	js.Rewrite("$<.onstalled = onstalled")
}

// onstop prop Fires when the user clicks the Stop button or leaves the Web page.
//     * @param ev The event.
func onstop() (onstop func(Event)) {
	js.Rewrite("$<.onstop")
	return onstop
}

// setonstop prop Fires when the user clicks the Stop button or leaves the Web page.
//     * @param ev The event.
func setonstop(onstop func(Event)) {
	js.Rewrite("$<.onstop = onstop")
}

// onsubmit prop
func onsubmit() (onsubmit func(Event)) {
	js.Rewrite("$<.onsubmit")
	return onsubmit
}

// setonsubmit prop
func setonsubmit(onsubmit func(Event)) {
	js.Rewrite("$<.onsubmit = onsubmit")
}

// onsuspend prop Occurs if the load operation has been intentionally halted.
//     * @param ev The event.
func onsuspend() (onsuspend func(Event)) {
	js.Rewrite("$<.onsuspend")
	return onsuspend
}

// setonsuspend prop Occurs if the load operation has been intentionally halted.
//     * @param ev The event.
func setonsuspend(onsuspend func(Event)) {
	js.Rewrite("$<.onsuspend = onsuspend")
}

// ontimeupdate prop Occurs to indicate the current playback position.
//     * @param ev The event.
func ontimeupdate() (ontimeupdate func(Event)) {
	js.Rewrite("$<.ontimeupdate")
	return ontimeupdate
}

// setontimeupdate prop Occurs to indicate the current playback position.
//     * @param ev The event.
func setontimeupdate(ontimeupdate func(Event)) {
	js.Rewrite("$<.ontimeupdate = ontimeupdate")
}

// ontouchcancel prop
func ontouchcancel() (ontouchcancel func(Event)) {
	js.Rewrite("$<.ontouchcancel")
	return ontouchcancel
}

// setontouchcancel prop
func setontouchcancel(ontouchcancel func(Event)) {
	js.Rewrite("$<.ontouchcancel = ontouchcancel")
}

// ontouchend prop
func ontouchend() (ontouchend func(Event)) {
	js.Rewrite("$<.ontouchend")
	return ontouchend
}

// setontouchend prop
func setontouchend(ontouchend func(Event)) {
	js.Rewrite("$<.ontouchend = ontouchend")
}

// ontouchmove prop
func ontouchmove() (ontouchmove func(Event)) {
	js.Rewrite("$<.ontouchmove")
	return ontouchmove
}

// setontouchmove prop
func setontouchmove(ontouchmove func(Event)) {
	js.Rewrite("$<.ontouchmove = ontouchmove")
}

// ontouchstart prop
func ontouchstart() (ontouchstart func(Event)) {
	js.Rewrite("$<.ontouchstart")
	return ontouchstart
}

// setontouchstart prop
func setontouchstart(ontouchstart func(Event)) {
	js.Rewrite("$<.ontouchstart = ontouchstart")
}

// onvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
//     * @param ev The event.
func onvolumechange() (onvolumechange func(Event)) {
	js.Rewrite("$<.onvolumechange")
	return onvolumechange
}

// setonvolumechange prop Occurs when the volume is changed, or playback is muted or unmuted.
//     * @param ev The event.
func setonvolumechange(onvolumechange func(Event)) {
	js.Rewrite("$<.onvolumechange = onvolumechange")
}

// onwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
//     * @param ev The event.
func onwaiting() (onwaiting func(Event)) {
	js.Rewrite("$<.onwaiting")
	return onwaiting
}

// setonwaiting prop Occurs when playback stops because the next frame of a video resource is not available.
//     * @param ev The event.
func setonwaiting(onwaiting func(Event)) {
	js.Rewrite("$<.onwaiting = onwaiting")
}

// onwebkitfullscreenchange prop
func onwebkitfullscreenchange() (onwebkitfullscreenchange func(Event)) {
	js.Rewrite("$<.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// setonwebkitfullscreenchange prop
func setonwebkitfullscreenchange(onwebkitfullscreenchange func(Event)) {
	js.Rewrite("$<.onwebkitfullscreenchange = onwebkitfullscreenchange")
}

// onwebkitfullscreenerror prop
func onwebkitfullscreenerror() (onwebkitfullscreenerror func(Event)) {
	js.Rewrite("$<.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// setonwebkitfullscreenerror prop
func setonwebkitfullscreenerror(onwebkitfullscreenerror func(Event)) {
	js.Rewrite("$<.onwebkitfullscreenerror = onwebkitfullscreenerror")
}

// plugins prop
func plugins() (plugins HTMLCollection) {
	js.Rewrite("$<.plugins")
	return plugins
}

// pointerlockelement prop
func pointerlockelement() (pointerLockElement Element) {
	js.Rewrite("$<.pointerLockElement")
	return pointerLockElement
}

// readystate prop Retrieves a value that indicates the current state of the object.
func readystate() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

// referrer prop Gets the URL of the location that referred the user to the current page.
func referrer() (referrer string) {
	js.Rewrite("$<.referrer")
	return referrer
}

// rootelement prop Gets the root svg element in the document hierarchy.
func rootelement() (rootElement *SVGSVGElement) {
	js.Rewrite("$<.rootElement")
	return rootElement
}

// scripts prop Retrieves a collection of all script objects in the document.
func scripts() (scripts HTMLCollection) {
	js.Rewrite("$<.scripts")
	return scripts
}

// scrollingelement prop
func scrollingelement() (scrollingElement Element) {
	js.Rewrite("$<.scrollingElement")
	return scrollingElement
}

// stylesheets prop Retrieves a collection of styleSheet objects representing the style sheets that correspond to each instance of a link or style object in the document.
func stylesheets() (styleSheets *StyleSheetList) {
	js.Rewrite("$<.styleSheets")
	return styleSheets
}

// title prop Contains the title of the document.
func title() (title string) {
	js.Rewrite("$<.title")
	return title
}

// settitle prop Contains the title of the document.
func settitle(title string) {
	js.Rewrite("$<.title = title")
}

// url prop Sets or gets the URL for the current document.
func url() (URL string) {
	js.Rewrite("$<.URL")
	return URL
}

// urlunencoded prop Gets the URL for the document, stripped of any character encoding.
func urlunencoded() (URLUnencoded string) {
	js.Rewrite("$<.URLUnencoded")
	return URLUnencoded
}

// // visibilitystate prop
// func visibilitystate() (visibilityState *visibilitystate.VisibilityState) {
// 	js.Rewrite("$<.visibilityState")
// 	return visibilityState
// }

// vlinkcolor prop Sets or gets the color of the links that the user has visited.
func vlinkcolor() (vlinkColor string) {
	js.Rewrite("$<.vlinkColor")
	return vlinkColor
}

// setvlinkcolor prop Sets or gets the color of the links that the user has visited.
func setvlinkcolor(vlinkColor string) {
	js.Rewrite("$<.vlinkColor = vlinkColor")
}

// webkitcurrentfullscreenelement prop
func webkitcurrentfullscreenelement() (webkitCurrentFullScreenElement Element) {
	js.Rewrite("$<.webkitCurrentFullScreenElement")
	return webkitCurrentFullScreenElement
}

// webkitfullscreenelement prop
func webkitfullscreenelement() (webkitFullscreenElement Element) {
	js.Rewrite("$<.webkitFullscreenElement")
	return webkitFullscreenElement
}

// webkitfullscreenenabled prop
func webkitfullscreenenabled() (webkitFullscreenEnabled bool) {
	js.Rewrite("$<.webkitFullscreenEnabled")
	return webkitFullscreenEnabled
}

// webkitisfullscreen prop
func webkitisfullscreen() (webkitIsFullScreen bool) {
	js.Rewrite("$<.webkitIsFullScreen")
	return webkitIsFullScreen
}

// xmlencoding prop
func xmlencoding() (xmlEncoding string) {
	js.Rewrite("$<.xmlEncoding")
	return xmlEncoding
}

// xmlstandalone prop
func xmlstandalone() (xmlStandalone bool) {
	js.Rewrite("$<.xmlStandalone")
	return xmlStandalone
}

// setxmlstandalone prop
func setxmlstandalone(xmlStandalone bool) {
	js.Rewrite("$<.xmlStandalone = xmlStandalone")
}

// xmlversion prop Gets or sets the version attribute specified in the declaration of an XML document.
func xmlversion() (xmlVersion string) {
	js.Rewrite("$<.xmlVersion")
	return xmlVersion
}

// setxmlversion prop Gets or sets the version attribute specified in the declaration of an XML document.
func setxmlversion(xmlVersion string) {
	js.Rewrite("$<.xmlVersion = xmlVersion")
}
