package window

import "github.com/matthewmueller/joy/dom/domstringmap"

// HTMLElement interface
// js:"HTMLElement"
type HTMLElement interface {
	Element

	// Blur
	// js:"blur"
	// jsrewrite:"$_.blur()"
	Blur()

	// Click
	// js:"click"
	// jsrewrite:"$_.click()"
	Click()

	// DragDrop
	// js:"dragDrop"
	// jsrewrite:"$_.dragDrop()"
	DragDrop() (b bool)

	// Focus
	// js:"focus"
	// jsrewrite:"$_.focus()"
	Focus()

	// GetElementsByClassName
	// js:"getElementsByClassName"
	// jsrewrite:"$_.getElementsByClassName($1)"
	GetElementsByClassName(classNames string) (n *NodeList)

	// InsertAdjacentElement
	// js:"insertAdjacentElement"
	// jsrewrite:"$_.insertAdjacentElement($1, $2)"
	InsertAdjacentElement(position string, insertedElement Element) (e Element)

	// InsertAdjacentHTML
	// js:"insertAdjacentHTML"
	// jsrewrite:"$_.insertAdjacentHTML($1, $2)"
	InsertAdjacentHTML(where string, html string)

	// InsertAdjacentText
	// js:"insertAdjacentText"
	// jsrewrite:"$_.insertAdjacentText($1, $2)"
	InsertAdjacentText(where string, text string)

	// MsGetInputContext
	// js:"msGetInputContext"
	// jsrewrite:"$_.msGetInputContext()"
	MsGetInputContext() (m *MSInputMethodContext)

	// ScrollIntoView
	// js:"scrollIntoView"
	// jsrewrite:"$_.scrollIntoView($1)"
	ScrollIntoView(top *bool)

	// AccessKey prop
	// js:"accessKey"
	// jsrewrite:"$_.accessKey"
	AccessKey() (accessKey string)

	// SetAccessKey prop
	// js:"accessKey"
	// jsrewrite:"$_.accessKey = $1"
	SetAccessKey(accessKey string)

	// Children prop
	// js:"children"
	// jsrewrite:"$_.children"
	Children() (children HTMLCollection)

	// ContentEditable prop
	// js:"contentEditable"
	// jsrewrite:"$_.contentEditable"
	ContentEditable() (contentEditable string)

	// SetContentEditable prop
	// js:"contentEditable"
	// jsrewrite:"$_.contentEditable = $1"
	SetContentEditable(contentEditable string)

	// Dataset prop
	// js:"dataset"
	// jsrewrite:"$_.dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// Dir prop
	// js:"dir"
	// jsrewrite:"$_.dir"
	Dir() (dir string)

	// SetDir prop
	// js:"dir"
	// jsrewrite:"$_.dir = $1"
	SetDir(dir string)

	// Draggable prop
	// js:"draggable"
	// jsrewrite:"$_.draggable"
	Draggable() (draggable bool)

	// SetDraggable prop
	// js:"draggable"
	// jsrewrite:"$_.draggable = $1"
	SetDraggable(draggable bool)

	// Hidden prop
	// js:"hidden"
	// jsrewrite:"$_.hidden"
	Hidden() (hidden bool)

	// SetHidden prop
	// js:"hidden"
	// jsrewrite:"$_.hidden = $1"
	SetHidden(hidden bool)

	// HideFocus prop
	// js:"hideFocus"
	// jsrewrite:"$_.hideFocus"
	HideFocus() (hideFocus bool)

	// SetHideFocus prop
	// js:"hideFocus"
	// jsrewrite:"$_.hideFocus = $1"
	SetHideFocus(hideFocus bool)

	// InnerText prop
	// js:"innerText"
	// jsrewrite:"$_.innerText"
	InnerText() (innerText string)

	// SetInnerText prop
	// js:"innerText"
	// jsrewrite:"$_.innerText = $1"
	SetInnerText(innerText string)

	// IsContentEditable prop
	// js:"isContentEditable"
	// jsrewrite:"$_.isContentEditable"
	IsContentEditable() (isContentEditable bool)

	// Lang prop
	// js:"lang"
	// jsrewrite:"$_.lang"
	Lang() (lang string)

	// SetLang prop
	// js:"lang"
	// jsrewrite:"$_.lang = $1"
	SetLang(lang string)

	// OffsetHeight prop
	// js:"offsetHeight"
	// jsrewrite:"$_.offsetHeight"
	OffsetHeight() (offsetHeight int)

	// OffsetLeft prop
	// js:"offsetLeft"
	// jsrewrite:"$_.offsetLeft"
	OffsetLeft() (offsetLeft int)

	// OffsetParent prop
	// js:"offsetParent"
	// jsrewrite:"$_.offsetParent"
	OffsetParent() (offsetParent Element)

	// OffsetTop prop
	// js:"offsetTop"
	// jsrewrite:"$_.offsetTop"
	OffsetTop() (offsetTop int)

	// OffsetWidth prop
	// js:"offsetWidth"
	// jsrewrite:"$_.offsetWidth"
	OffsetWidth() (offsetWidth int)

	// Onabort prop
	// js:"onabort"
	// jsrewrite:"$_.onabort"
	Onabort() (onabort func(Event))

	// SetOnabort prop
	// js:"onabort"
	// jsrewrite:"$_.onabort = $1"
	SetOnabort(onabort func(Event))

	// Onactivate prop
	// js:"onactivate"
	// jsrewrite:"$_.onactivate"
	Onactivate() (onactivate func(UIEvent))

	// SetOnactivate prop
	// js:"onactivate"
	// jsrewrite:"$_.onactivate = $1"
	SetOnactivate(onactivate func(UIEvent))

	// Onbeforeactivate prop
	// js:"onbeforeactivate"
	// jsrewrite:"$_.onbeforeactivate"
	Onbeforeactivate() (onbeforeactivate func(UIEvent))

	// SetOnbeforeactivate prop
	// js:"onbeforeactivate"
	// jsrewrite:"$_.onbeforeactivate = $1"
	SetOnbeforeactivate(onbeforeactivate func(UIEvent))

	// Onbeforecopy prop
	// js:"onbeforecopy"
	// jsrewrite:"$_.onbeforecopy"
	Onbeforecopy() (onbeforecopy func(*ClipboardEvent))

	// SetOnbeforecopy prop
	// js:"onbeforecopy"
	// jsrewrite:"$_.onbeforecopy = $1"
	SetOnbeforecopy(onbeforecopy func(*ClipboardEvent))

	// Onbeforecut prop
	// js:"onbeforecut"
	// jsrewrite:"$_.onbeforecut"
	Onbeforecut() (onbeforecut func(*ClipboardEvent))

	// SetOnbeforecut prop
	// js:"onbeforecut"
	// jsrewrite:"$_.onbeforecut = $1"
	SetOnbeforecut(onbeforecut func(*ClipboardEvent))

	// Onbeforedeactivate prop
	// js:"onbeforedeactivate"
	// jsrewrite:"$_.onbeforedeactivate"
	Onbeforedeactivate() (onbeforedeactivate func(UIEvent))

	// SetOnbeforedeactivate prop
	// js:"onbeforedeactivate"
	// jsrewrite:"$_.onbeforedeactivate = $1"
	SetOnbeforedeactivate(onbeforedeactivate func(UIEvent))

	// Onbeforepaste prop
	// js:"onbeforepaste"
	// jsrewrite:"$_.onbeforepaste"
	Onbeforepaste() (onbeforepaste func(*ClipboardEvent))

	// SetOnbeforepaste prop
	// js:"onbeforepaste"
	// jsrewrite:"$_.onbeforepaste = $1"
	SetOnbeforepaste(onbeforepaste func(*ClipboardEvent))

	// Onblur prop
	// js:"onblur"
	// jsrewrite:"$_.onblur"
	Onblur() (onblur func(*FocusEvent))

	// SetOnblur prop
	// js:"onblur"
	// jsrewrite:"$_.onblur = $1"
	SetOnblur(onblur func(*FocusEvent))

	// Oncanplay prop
	// js:"oncanplay"
	// jsrewrite:"$_.oncanplay"
	Oncanplay() (oncanplay func(Event))

	// SetOncanplay prop
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

	// Onchange prop
	// js:"onchange"
	// jsrewrite:"$_.onchange"
	Onchange() (onchange func(Event))

	// SetOnchange prop
	// js:"onchange"
	// jsrewrite:"$_.onchange = $1"
	SetOnchange(onchange func(Event))

	// Onclick prop
	// js:"onclick"
	// jsrewrite:"$_.onclick"
	Onclick() (onclick func(MouseEvent))

	// SetOnclick prop
	// js:"onclick"
	// jsrewrite:"$_.onclick = $1"
	SetOnclick(onclick func(MouseEvent))

	// Oncontextmenu prop
	// js:"oncontextmenu"
	// jsrewrite:"$_.oncontextmenu"
	Oncontextmenu() (oncontextmenu func(*PointerEvent))

	// SetOncontextmenu prop
	// js:"oncontextmenu"
	// jsrewrite:"$_.oncontextmenu = $1"
	SetOncontextmenu(oncontextmenu func(*PointerEvent))

	// Oncopy prop
	// js:"oncopy"
	// jsrewrite:"$_.oncopy"
	Oncopy() (oncopy func(*ClipboardEvent))

	// SetOncopy prop
	// js:"oncopy"
	// jsrewrite:"$_.oncopy = $1"
	SetOncopy(oncopy func(*ClipboardEvent))

	// Oncuechange prop
	// js:"oncuechange"
	// jsrewrite:"$_.oncuechange"
	Oncuechange() (oncuechange func(Event))

	// SetOncuechange prop
	// js:"oncuechange"
	// jsrewrite:"$_.oncuechange = $1"
	SetOncuechange(oncuechange func(Event))

	// Oncut prop
	// js:"oncut"
	// jsrewrite:"$_.oncut"
	Oncut() (oncut func(*ClipboardEvent))

	// SetOncut prop
	// js:"oncut"
	// jsrewrite:"$_.oncut = $1"
	SetOncut(oncut func(*ClipboardEvent))

	// Ondblclick prop
	// js:"ondblclick"
	// jsrewrite:"$_.ondblclick"
	Ondblclick() (ondblclick func(MouseEvent))

	// SetOndblclick prop
	// js:"ondblclick"
	// jsrewrite:"$_.ondblclick = $1"
	SetOndblclick(ondblclick func(MouseEvent))

	// Ondeactivate prop
	// js:"ondeactivate"
	// jsrewrite:"$_.ondeactivate"
	Ondeactivate() (ondeactivate func(UIEvent))

	// SetOndeactivate prop
	// js:"ondeactivate"
	// jsrewrite:"$_.ondeactivate = $1"
	SetOndeactivate(ondeactivate func(UIEvent))

	// Ondrag prop
	// js:"ondrag"
	// jsrewrite:"$_.ondrag"
	Ondrag() (ondrag func(*DragEvent))

	// SetOndrag prop
	// js:"ondrag"
	// jsrewrite:"$_.ondrag = $1"
	SetOndrag(ondrag func(*DragEvent))

	// Ondragend prop
	// js:"ondragend"
	// jsrewrite:"$_.ondragend"
	Ondragend() (ondragend func(*DragEvent))

	// SetOndragend prop
	// js:"ondragend"
	// jsrewrite:"$_.ondragend = $1"
	SetOndragend(ondragend func(*DragEvent))

	// Ondragenter prop
	// js:"ondragenter"
	// jsrewrite:"$_.ondragenter"
	Ondragenter() (ondragenter func(*DragEvent))

	// SetOndragenter prop
	// js:"ondragenter"
	// jsrewrite:"$_.ondragenter = $1"
	SetOndragenter(ondragenter func(*DragEvent))

	// Ondragleave prop
	// js:"ondragleave"
	// jsrewrite:"$_.ondragleave"
	Ondragleave() (ondragleave func(*DragEvent))

	// SetOndragleave prop
	// js:"ondragleave"
	// jsrewrite:"$_.ondragleave = $1"
	SetOndragleave(ondragleave func(*DragEvent))

	// Ondragover prop
	// js:"ondragover"
	// jsrewrite:"$_.ondragover"
	Ondragover() (ondragover func(*DragEvent))

	// SetOndragover prop
	// js:"ondragover"
	// jsrewrite:"$_.ondragover = $1"
	SetOndragover(ondragover func(*DragEvent))

	// Ondragstart prop
	// js:"ondragstart"
	// jsrewrite:"$_.ondragstart"
	Ondragstart() (ondragstart func(*DragEvent))

	// SetOndragstart prop
	// js:"ondragstart"
	// jsrewrite:"$_.ondragstart = $1"
	SetOndragstart(ondragstart func(*DragEvent))

	// Ondrop prop
	// js:"ondrop"
	// jsrewrite:"$_.ondrop"
	Ondrop() (ondrop func(*DragEvent))

	// SetOndrop prop
	// js:"ondrop"
	// jsrewrite:"$_.ondrop = $1"
	SetOndrop(ondrop func(*DragEvent))

	// Ondurationchange prop
	// js:"ondurationchange"
	// jsrewrite:"$_.ondurationchange"
	Ondurationchange() (ondurationchange func(Event))

	// SetOndurationchange prop
	// js:"ondurationchange"
	// jsrewrite:"$_.ondurationchange = $1"
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied prop
	// js:"onemptied"
	// jsrewrite:"$_.onemptied"
	Onemptied() (onemptied func(Event))

	// SetOnemptied prop
	// js:"onemptied"
	// jsrewrite:"$_.onemptied = $1"
	SetOnemptied(onemptied func(Event))

	// Onended prop
	// js:"onended"
	// jsrewrite:"$_.onended"
	Onended() (onended func(Event))

	// SetOnended prop
	// js:"onended"
	// jsrewrite:"$_.onended = $1"
	SetOnended(onended func(Event))

	// Onerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror = $1"
	SetOnerror(onerror func(Event))

	// Onfocus prop
	// js:"onfocus"
	// jsrewrite:"$_.onfocus"
	Onfocus() (onfocus func(*FocusEvent))

	// SetOnfocus prop
	// js:"onfocus"
	// jsrewrite:"$_.onfocus = $1"
	SetOnfocus(onfocus func(*FocusEvent))

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

	// Onkeydown prop
	// js:"onkeydown"
	// jsrewrite:"$_.onkeydown"
	Onkeydown() (onkeydown func(*KeyboardEvent))

	// SetOnkeydown prop
	// js:"onkeydown"
	// jsrewrite:"$_.onkeydown = $1"
	SetOnkeydown(onkeydown func(*KeyboardEvent))

	// Onkeypress prop
	// js:"onkeypress"
	// jsrewrite:"$_.onkeypress"
	Onkeypress() (onkeypress func(*KeyboardEvent))

	// SetOnkeypress prop
	// js:"onkeypress"
	// jsrewrite:"$_.onkeypress = $1"
	SetOnkeypress(onkeypress func(*KeyboardEvent))

	// Onkeyup prop
	// js:"onkeyup"
	// jsrewrite:"$_.onkeyup"
	Onkeyup() (onkeyup func(*KeyboardEvent))

	// SetOnkeyup prop
	// js:"onkeyup"
	// jsrewrite:"$_.onkeyup = $1"
	SetOnkeyup(onkeyup func(*KeyboardEvent))

	// Onload prop
	// js:"onload"
	// jsrewrite:"$_.onload"
	Onload() (onload func(Event))

	// SetOnload prop
	// js:"onload"
	// jsrewrite:"$_.onload = $1"
	SetOnload(onload func(Event))

	// Onloadeddata prop
	// js:"onloadeddata"
	// jsrewrite:"$_.onloadeddata"
	Onloadeddata() (onloadeddata func(Event))

	// SetOnloadeddata prop
	// js:"onloadeddata"
	// jsrewrite:"$_.onloadeddata = $1"
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata prop
	// js:"onloadedmetadata"
	// jsrewrite:"$_.onloadedmetadata"
	Onloadedmetadata() (onloadedmetadata func(Event))

	// SetOnloadedmetadata prop
	// js:"onloadedmetadata"
	// jsrewrite:"$_.onloadedmetadata = $1"
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart prop
	// js:"onloadstart"
	// jsrewrite:"$_.onloadstart"
	Onloadstart() (onloadstart func(Event))

	// SetOnloadstart prop
	// js:"onloadstart"
	// jsrewrite:"$_.onloadstart = $1"
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown prop
	// js:"onmousedown"
	// jsrewrite:"$_.onmousedown"
	Onmousedown() (onmousedown func(MouseEvent))

	// SetOnmousedown prop
	// js:"onmousedown"
	// jsrewrite:"$_.onmousedown = $1"
	SetOnmousedown(onmousedown func(MouseEvent))

	// Onmouseenter prop
	// js:"onmouseenter"
	// jsrewrite:"$_.onmouseenter"
	Onmouseenter() (onmouseenter func(MouseEvent))

	// SetOnmouseenter prop
	// js:"onmouseenter"
	// jsrewrite:"$_.onmouseenter = $1"
	SetOnmouseenter(onmouseenter func(MouseEvent))

	// Onmouseleave prop
	// js:"onmouseleave"
	// jsrewrite:"$_.onmouseleave"
	Onmouseleave() (onmouseleave func(MouseEvent))

	// SetOnmouseleave prop
	// js:"onmouseleave"
	// jsrewrite:"$_.onmouseleave = $1"
	SetOnmouseleave(onmouseleave func(MouseEvent))

	// Onmousemove prop
	// js:"onmousemove"
	// jsrewrite:"$_.onmousemove"
	Onmousemove() (onmousemove func(MouseEvent))

	// SetOnmousemove prop
	// js:"onmousemove"
	// jsrewrite:"$_.onmousemove = $1"
	SetOnmousemove(onmousemove func(MouseEvent))

	// Onmouseout prop
	// js:"onmouseout"
	// jsrewrite:"$_.onmouseout"
	Onmouseout() (onmouseout func(MouseEvent))

	// SetOnmouseout prop
	// js:"onmouseout"
	// jsrewrite:"$_.onmouseout = $1"
	SetOnmouseout(onmouseout func(MouseEvent))

	// Onmouseover prop
	// js:"onmouseover"
	// jsrewrite:"$_.onmouseover"
	Onmouseover() (onmouseover func(MouseEvent))

	// SetOnmouseover prop
	// js:"onmouseover"
	// jsrewrite:"$_.onmouseover = $1"
	SetOnmouseover(onmouseover func(MouseEvent))

	// Onmouseup prop
	// js:"onmouseup"
	// jsrewrite:"$_.onmouseup"
	Onmouseup() (onmouseup func(MouseEvent))

	// SetOnmouseup prop
	// js:"onmouseup"
	// jsrewrite:"$_.onmouseup = $1"
	SetOnmouseup(onmouseup func(MouseEvent))

	// Onmousewheel prop
	// js:"onmousewheel"
	// jsrewrite:"$_.onmousewheel"
	Onmousewheel() (onmousewheel func(*WheelEvent))

	// SetOnmousewheel prop
	// js:"onmousewheel"
	// jsrewrite:"$_.onmousewheel = $1"
	SetOnmousewheel(onmousewheel func(*WheelEvent))

	// Onmscontentzoom prop
	// js:"onmscontentzoom"
	// jsrewrite:"$_.onmscontentzoom"
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// SetOnmscontentzoom prop
	// js:"onmscontentzoom"
	// jsrewrite:"$_.onmscontentzoom = $1"
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	// jsrewrite:"$_.onmsmanipulationstatechanged"
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// SetOnmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	// jsrewrite:"$_.onmsmanipulationstatechanged = $1"
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onpaste prop
	// js:"onpaste"
	// jsrewrite:"$_.onpaste"
	Onpaste() (onpaste func(*ClipboardEvent))

	// SetOnpaste prop
	// js:"onpaste"
	// jsrewrite:"$_.onpaste = $1"
	SetOnpaste(onpaste func(*ClipboardEvent))

	// Onpause prop
	// js:"onpause"
	// jsrewrite:"$_.onpause"
	Onpause() (onpause func(Event))

	// SetOnpause prop
	// js:"onpause"
	// jsrewrite:"$_.onpause = $1"
	SetOnpause(onpause func(Event))

	// Onplay prop
	// js:"onplay"
	// jsrewrite:"$_.onplay"
	Onplay() (onplay func(Event))

	// SetOnplay prop
	// js:"onplay"
	// jsrewrite:"$_.onplay = $1"
	SetOnplay(onplay func(Event))

	// Onplaying prop
	// js:"onplaying"
	// jsrewrite:"$_.onplaying"
	Onplaying() (onplaying func(Event))

	// SetOnplaying prop
	// js:"onplaying"
	// jsrewrite:"$_.onplaying = $1"
	SetOnplaying(onplaying func(Event))

	// Onprogress prop
	// js:"onprogress"
	// jsrewrite:"$_.onprogress"
	Onprogress() (onprogress func(Event))

	// SetOnprogress prop
	// js:"onprogress"
	// jsrewrite:"$_.onprogress = $1"
	SetOnprogress(onprogress func(Event))

	// Onratechange prop
	// js:"onratechange"
	// jsrewrite:"$_.onratechange"
	Onratechange() (onratechange func(Event))

	// SetOnratechange prop
	// js:"onratechange"
	// jsrewrite:"$_.onratechange = $1"
	SetOnratechange(onratechange func(Event))

	// Onreset prop
	// js:"onreset"
	// jsrewrite:"$_.onreset"
	Onreset() (onreset func(Event))

	// SetOnreset prop
	// js:"onreset"
	// jsrewrite:"$_.onreset = $1"
	SetOnreset(onreset func(Event))

	// Onscroll prop
	// js:"onscroll"
	// jsrewrite:"$_.onscroll"
	Onscroll() (onscroll func(UIEvent))

	// SetOnscroll prop
	// js:"onscroll"
	// jsrewrite:"$_.onscroll = $1"
	SetOnscroll(onscroll func(UIEvent))

	// Onseeked prop
	// js:"onseeked"
	// jsrewrite:"$_.onseeked"
	Onseeked() (onseeked func(Event))

	// SetOnseeked prop
	// js:"onseeked"
	// jsrewrite:"$_.onseeked = $1"
	SetOnseeked(onseeked func(Event))

	// Onseeking prop
	// js:"onseeking"
	// jsrewrite:"$_.onseeking"
	Onseeking() (onseeking func(Event))

	// SetOnseeking prop
	// js:"onseeking"
	// jsrewrite:"$_.onseeking = $1"
	SetOnseeking(onseeking func(Event))

	// Onselect prop
	// js:"onselect"
	// jsrewrite:"$_.onselect"
	Onselect() (onselect func(UIEvent))

	// SetOnselect prop
	// js:"onselect"
	// jsrewrite:"$_.onselect = $1"
	SetOnselect(onselect func(UIEvent))

	// Onselectstart prop
	// js:"onselectstart"
	// jsrewrite:"$_.onselectstart"
	Onselectstart() (onselectstart func(Event))

	// SetOnselectstart prop
	// js:"onselectstart"
	// jsrewrite:"$_.onselectstart = $1"
	SetOnselectstart(onselectstart func(Event))

	// Onstalled prop
	// js:"onstalled"
	// jsrewrite:"$_.onstalled"
	Onstalled() (onstalled func(Event))

	// SetOnstalled prop
	// js:"onstalled"
	// jsrewrite:"$_.onstalled = $1"
	SetOnstalled(onstalled func(Event))

	// Onsubmit prop
	// js:"onsubmit"
	// jsrewrite:"$_.onsubmit"
	Onsubmit() (onsubmit func(Event))

	// SetOnsubmit prop
	// js:"onsubmit"
	// jsrewrite:"$_.onsubmit = $1"
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend prop
	// js:"onsuspend"
	// jsrewrite:"$_.onsuspend"
	Onsuspend() (onsuspend func(Event))

	// SetOnsuspend prop
	// js:"onsuspend"
	// jsrewrite:"$_.onsuspend = $1"
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate prop
	// js:"ontimeupdate"
	// jsrewrite:"$_.ontimeupdate"
	Ontimeupdate() (ontimeupdate func(Event))

	// SetOntimeupdate prop
	// js:"ontimeupdate"
	// jsrewrite:"$_.ontimeupdate = $1"
	SetOntimeupdate(ontimeupdate func(Event))

	// Onvolumechange prop
	// js:"onvolumechange"
	// jsrewrite:"$_.onvolumechange"
	Onvolumechange() (onvolumechange func(Event))

	// SetOnvolumechange prop
	// js:"onvolumechange"
	// jsrewrite:"$_.onvolumechange = $1"
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting prop
	// js:"onwaiting"
	// jsrewrite:"$_.onwaiting"
	Onwaiting() (onwaiting func(Event))

	// SetOnwaiting prop
	// js:"onwaiting"
	// jsrewrite:"$_.onwaiting = $1"
	SetOnwaiting(onwaiting func(Event))

	// OuterText prop
	// js:"outerText"
	// jsrewrite:"$_.outerText"
	OuterText() (outerText string)

	// SetOuterText prop
	// js:"outerText"
	// jsrewrite:"$_.outerText = $1"
	SetOuterText(outerText string)

	// Spellcheck prop
	// js:"spellcheck"
	// jsrewrite:"$_.spellcheck"
	Spellcheck() (spellcheck bool)

	// SetSpellcheck prop
	// js:"spellcheck"
	// jsrewrite:"$_.spellcheck = $1"
	SetSpellcheck(spellcheck bool)

	// Style prop
	// js:"style"
	// jsrewrite:"$_.style"
	Style() (style *CSSStyleDeclaration)

	// TabIndex prop
	// js:"tabIndex"
	// jsrewrite:"$_.tabIndex"
	TabIndex() (tabIndex int8)

	// SetTabIndex prop
	// js:"tabIndex"
	// jsrewrite:"$_.tabIndex = $1"
	SetTabIndex(tabIndex int8)

	// Title prop
	// js:"title"
	// jsrewrite:"$_.title"
	Title() (title string)

	// SetTitle prop
	// js:"title"
	// jsrewrite:"$_.title = $1"
	SetTitle(title string)
}
