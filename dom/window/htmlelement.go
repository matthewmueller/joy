package window

import "github.com/matthewmueller/golly/dom2/domstringmap"

// js:"HTMLElement,omit"
type HTMLElement interface {
	Element

	// Blur
	// js:"blur"
	Blur()

	// Click
	// js:"click"
	Click()

	// DragDrop
	// js:"dragDrop"
	DragDrop() (b bool)

	// Focus
	// js:"focus"
	Focus()

	// GetElementsByClassName
	// js:"getElementsByClassName"
	GetElementsByClassName(classNames string) (n *NodeList)

	// InsertAdjacentElement
	// js:"insertAdjacentElement"
	InsertAdjacentElement(position string, insertedElement Element) (e Element)

	// InsertAdjacentHTML
	// js:"insertAdjacentHTML"
	InsertAdjacentHTML(where string, html string)

	// InsertAdjacentText
	// js:"insertAdjacentText"
	InsertAdjacentText(where string, text string)

	// MsGetInputContext
	// js:"msGetInputContext"
	MsGetInputContext() (m *MSInputMethodContext)

	// ScrollIntoView
	// js:"scrollIntoView"
	ScrollIntoView(top *bool)

	// AccessKey prop
	// js:"accessKey"
	AccessKey() (accessKey string)

	// AccessKey prop
	SetAccessKey(accessKey string)

	// Children prop
	// js:"children"
	Children() (children HTMLCollection)

	// ContentEditable prop
	// js:"contentEditable"
	ContentEditable() (contentEditable string)

	// ContentEditable prop
	SetContentEditable(contentEditable string)

	// Dataset prop
	// js:"dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// Dir prop
	// js:"dir"
	Dir() (dir string)

	// Dir prop
	SetDir(dir string)

	// Draggable prop
	// js:"draggable"
	Draggable() (draggable bool)

	// Draggable prop
	SetDraggable(draggable bool)

	// Hidden prop
	// js:"hidden"
	Hidden() (hidden bool)

	// Hidden prop
	SetHidden(hidden bool)

	// HideFocus prop
	// js:"hideFocus"
	HideFocus() (hideFocus bool)

	// HideFocus prop
	SetHideFocus(hideFocus bool)

	// InnerText prop
	// js:"innerText"
	InnerText() (innerText string)

	// InnerText prop
	SetInnerText(innerText string)

	// IsContentEditable prop
	// js:"isContentEditable"
	IsContentEditable() (isContentEditable bool)

	// Lang prop
	// js:"lang"
	Lang() (lang string)

	// Lang prop
	SetLang(lang string)

	// OffsetHeight prop
	// js:"offsetHeight"
	OffsetHeight() (offsetHeight int)

	// OffsetLeft prop
	// js:"offsetLeft"
	OffsetLeft() (offsetLeft int)

	// OffsetParent prop
	// js:"offsetParent"
	OffsetParent() (offsetParent Element)

	// OffsetTop prop
	// js:"offsetTop"
	OffsetTop() (offsetTop int)

	// OffsetWidth prop
	// js:"offsetWidth"
	OffsetWidth() (offsetWidth int)

	// Onabort prop
	// js:"onabort"
	Onabort() (onabort func(Event))

	// Onabort prop
	SetOnabort(onabort func(Event))

	// Onactivate prop
	// js:"onactivate"
	Onactivate() (onactivate func(UIEvent))

	// Onactivate prop
	SetOnactivate(onactivate func(UIEvent))

	// Onbeforeactivate prop
	// js:"onbeforeactivate"
	Onbeforeactivate() (onbeforeactivate func(UIEvent))

	// Onbeforeactivate prop
	SetOnbeforeactivate(onbeforeactivate func(UIEvent))

	// Onbeforecopy prop
	// js:"onbeforecopy"
	Onbeforecopy() (onbeforecopy func(*ClipboardEvent))

	// Onbeforecopy prop
	SetOnbeforecopy(onbeforecopy func(*ClipboardEvent))

	// Onbeforecut prop
	// js:"onbeforecut"
	Onbeforecut() (onbeforecut func(*ClipboardEvent))

	// Onbeforecut prop
	SetOnbeforecut(onbeforecut func(*ClipboardEvent))

	// Onbeforedeactivate prop
	// js:"onbeforedeactivate"
	Onbeforedeactivate() (onbeforedeactivate func(UIEvent))

	// Onbeforedeactivate prop
	SetOnbeforedeactivate(onbeforedeactivate func(UIEvent))

	// Onbeforepaste prop
	// js:"onbeforepaste"
	Onbeforepaste() (onbeforepaste func(*ClipboardEvent))

	// Onbeforepaste prop
	SetOnbeforepaste(onbeforepaste func(*ClipboardEvent))

	// Onblur prop
	// js:"onblur"
	Onblur() (onblur func(*FocusEvent))

	// Onblur prop
	SetOnblur(onblur func(*FocusEvent))

	// Oncanplay prop
	// js:"oncanplay"
	Oncanplay() (oncanplay func(Event))

	// Oncanplay prop
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough prop
	// js:"oncanplaythrough"
	Oncanplaythrough() (oncanplaythrough func(Event))

	// Oncanplaythrough prop
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange prop
	// js:"onchange"
	Onchange() (onchange func(Event))

	// Onchange prop
	SetOnchange(onchange func(Event))

	// Onclick prop
	// js:"onclick"
	Onclick() (onclick func(MouseEvent))

	// Onclick prop
	SetOnclick(onclick func(MouseEvent))

	// Oncontextmenu prop
	// js:"oncontextmenu"
	Oncontextmenu() (oncontextmenu func(*PointerEvent))

	// Oncontextmenu prop
	SetOncontextmenu(oncontextmenu func(*PointerEvent))

	// Oncopy prop
	// js:"oncopy"
	Oncopy() (oncopy func(*ClipboardEvent))

	// Oncopy prop
	SetOncopy(oncopy func(*ClipboardEvent))

	// Oncuechange prop
	// js:"oncuechange"
	Oncuechange() (oncuechange func(Event))

	// Oncuechange prop
	SetOncuechange(oncuechange func(Event))

	// Oncut prop
	// js:"oncut"
	Oncut() (oncut func(*ClipboardEvent))

	// Oncut prop
	SetOncut(oncut func(*ClipboardEvent))

	// Ondblclick prop
	// js:"ondblclick"
	Ondblclick() (ondblclick func(MouseEvent))

	// Ondblclick prop
	SetOndblclick(ondblclick func(MouseEvent))

	// Ondeactivate prop
	// js:"ondeactivate"
	Ondeactivate() (ondeactivate func(UIEvent))

	// Ondeactivate prop
	SetOndeactivate(ondeactivate func(UIEvent))

	// Ondrag prop
	// js:"ondrag"
	Ondrag() (ondrag func(*DragEvent))

	// Ondrag prop
	SetOndrag(ondrag func(*DragEvent))

	// Ondragend prop
	// js:"ondragend"
	Ondragend() (ondragend func(*DragEvent))

	// Ondragend prop
	SetOndragend(ondragend func(*DragEvent))

	// Ondragenter prop
	// js:"ondragenter"
	Ondragenter() (ondragenter func(*DragEvent))

	// Ondragenter prop
	SetOndragenter(ondragenter func(*DragEvent))

	// Ondragleave prop
	// js:"ondragleave"
	Ondragleave() (ondragleave func(*DragEvent))

	// Ondragleave prop
	SetOndragleave(ondragleave func(*DragEvent))

	// Ondragover prop
	// js:"ondragover"
	Ondragover() (ondragover func(*DragEvent))

	// Ondragover prop
	SetOndragover(ondragover func(*DragEvent))

	// Ondragstart prop
	// js:"ondragstart"
	Ondragstart() (ondragstart func(*DragEvent))

	// Ondragstart prop
	SetOndragstart(ondragstart func(*DragEvent))

	// Ondrop prop
	// js:"ondrop"
	Ondrop() (ondrop func(*DragEvent))

	// Ondrop prop
	SetOndrop(ondrop func(*DragEvent))

	// Ondurationchange prop
	// js:"ondurationchange"
	Ondurationchange() (ondurationchange func(Event))

	// Ondurationchange prop
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied prop
	// js:"onemptied"
	Onemptied() (onemptied func(Event))

	// Onemptied prop
	SetOnemptied(onemptied func(Event))

	// Onended prop
	// js:"onended"
	Onended() (onended func(Event))

	// Onended prop
	SetOnended(onended func(Event))

	// Onerror prop
	// js:"onerror"
	Onerror() (onerror func(Event))

	// Onerror prop
	SetOnerror(onerror func(Event))

	// Onfocus prop
	// js:"onfocus"
	Onfocus() (onfocus func(*FocusEvent))

	// Onfocus prop
	SetOnfocus(onfocus func(*FocusEvent))

	// Oninput prop
	// js:"oninput"
	Oninput() (oninput func(Event))

	// Oninput prop
	SetOninput(oninput func(Event))

	// Oninvalid prop
	// js:"oninvalid"
	Oninvalid() (oninvalid func(Event))

	// Oninvalid prop
	SetOninvalid(oninvalid func(Event))

	// Onkeydown prop
	// js:"onkeydown"
	Onkeydown() (onkeydown func(*KeyboardEvent))

	// Onkeydown prop
	SetOnkeydown(onkeydown func(*KeyboardEvent))

	// Onkeypress prop
	// js:"onkeypress"
	Onkeypress() (onkeypress func(*KeyboardEvent))

	// Onkeypress prop
	SetOnkeypress(onkeypress func(*KeyboardEvent))

	// Onkeyup prop
	// js:"onkeyup"
	Onkeyup() (onkeyup func(*KeyboardEvent))

	// Onkeyup prop
	SetOnkeyup(onkeyup func(*KeyboardEvent))

	// Onload prop
	// js:"onload"
	Onload() (onload func(Event))

	// Onload prop
	SetOnload(onload func(Event))

	// Onloadeddata prop
	// js:"onloadeddata"
	Onloadeddata() (onloadeddata func(Event))

	// Onloadeddata prop
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata prop
	// js:"onloadedmetadata"
	Onloadedmetadata() (onloadedmetadata func(Event))

	// Onloadedmetadata prop
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart prop
	// js:"onloadstart"
	Onloadstart() (onloadstart func(Event))

	// Onloadstart prop
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown prop
	// js:"onmousedown"
	Onmousedown() (onmousedown func(MouseEvent))

	// Onmousedown prop
	SetOnmousedown(onmousedown func(MouseEvent))

	// Onmouseenter prop
	// js:"onmouseenter"
	Onmouseenter() (onmouseenter func(MouseEvent))

	// Onmouseenter prop
	SetOnmouseenter(onmouseenter func(MouseEvent))

	// Onmouseleave prop
	// js:"onmouseleave"
	Onmouseleave() (onmouseleave func(MouseEvent))

	// Onmouseleave prop
	SetOnmouseleave(onmouseleave func(MouseEvent))

	// Onmousemove prop
	// js:"onmousemove"
	Onmousemove() (onmousemove func(MouseEvent))

	// Onmousemove prop
	SetOnmousemove(onmousemove func(MouseEvent))

	// Onmouseout prop
	// js:"onmouseout"
	Onmouseout() (onmouseout func(MouseEvent))

	// Onmouseout prop
	SetOnmouseout(onmouseout func(MouseEvent))

	// Onmouseover prop
	// js:"onmouseover"
	Onmouseover() (onmouseover func(MouseEvent))

	// Onmouseover prop
	SetOnmouseover(onmouseover func(MouseEvent))

	// Onmouseup prop
	// js:"onmouseup"
	Onmouseup() (onmouseup func(MouseEvent))

	// Onmouseup prop
	SetOnmouseup(onmouseup func(MouseEvent))

	// Onmousewheel prop
	// js:"onmousewheel"
	Onmousewheel() (onmousewheel func(*WheelEvent))

	// Onmousewheel prop
	SetOnmousewheel(onmousewheel func(*WheelEvent))

	// Onmscontentzoom prop
	// js:"onmscontentzoom"
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// Onmscontentzoom prop
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmsmanipulationstatechanged prop
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onpaste prop
	// js:"onpaste"
	Onpaste() (onpaste func(*ClipboardEvent))

	// Onpaste prop
	SetOnpaste(onpaste func(*ClipboardEvent))

	// Onpause prop
	// js:"onpause"
	Onpause() (onpause func(Event))

	// Onpause prop
	SetOnpause(onpause func(Event))

	// Onplay prop
	// js:"onplay"
	Onplay() (onplay func(Event))

	// Onplay prop
	SetOnplay(onplay func(Event))

	// Onplaying prop
	// js:"onplaying"
	Onplaying() (onplaying func(Event))

	// Onplaying prop
	SetOnplaying(onplaying func(Event))

	// Onprogress prop
	// js:"onprogress"
	Onprogress() (onprogress func(Event))

	// Onprogress prop
	SetOnprogress(onprogress func(Event))

	// Onratechange prop
	// js:"onratechange"
	Onratechange() (onratechange func(Event))

	// Onratechange prop
	SetOnratechange(onratechange func(Event))

	// Onreset prop
	// js:"onreset"
	Onreset() (onreset func(Event))

	// Onreset prop
	SetOnreset(onreset func(Event))

	// Onscroll prop
	// js:"onscroll"
	Onscroll() (onscroll func(UIEvent))

	// Onscroll prop
	SetOnscroll(onscroll func(UIEvent))

	// Onseeked prop
	// js:"onseeked"
	Onseeked() (onseeked func(Event))

	// Onseeked prop
	SetOnseeked(onseeked func(Event))

	// Onseeking prop
	// js:"onseeking"
	Onseeking() (onseeking func(Event))

	// Onseeking prop
	SetOnseeking(onseeking func(Event))

	// Onselect prop
	// js:"onselect"
	Onselect() (onselect func(UIEvent))

	// Onselect prop
	SetOnselect(onselect func(UIEvent))

	// Onselectstart prop
	// js:"onselectstart"
	Onselectstart() (onselectstart func(Event))

	// Onselectstart prop
	SetOnselectstart(onselectstart func(Event))

	// Onstalled prop
	// js:"onstalled"
	Onstalled() (onstalled func(Event))

	// Onstalled prop
	SetOnstalled(onstalled func(Event))

	// Onsubmit prop
	// js:"onsubmit"
	Onsubmit() (onsubmit func(Event))

	// Onsubmit prop
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend prop
	// js:"onsuspend"
	Onsuspend() (onsuspend func(Event))

	// Onsuspend prop
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate prop
	// js:"ontimeupdate"
	Ontimeupdate() (ontimeupdate func(Event))

	// Ontimeupdate prop
	SetOntimeupdate(ontimeupdate func(Event))

	// Onvolumechange prop
	// js:"onvolumechange"
	Onvolumechange() (onvolumechange func(Event))

	// Onvolumechange prop
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting prop
	// js:"onwaiting"
	Onwaiting() (onwaiting func(Event))

	// Onwaiting prop
	SetOnwaiting(onwaiting func(Event))

	// OuterText prop
	// js:"outerText"
	OuterText() (outerText string)

	// OuterText prop
	SetOuterText(outerText string)

	// Spellcheck prop
	// js:"spellcheck"
	Spellcheck() (spellcheck bool)

	// Spellcheck prop
	SetSpellcheck(spellcheck bool)

	// Style prop
	// js:"style"
	Style() (style *CSSStyleDeclaration)

	// TabIndex prop
	// js:"tabIndex"
	TabIndex() (tabIndex int8)

	// TabIndex prop
	SetTabIndex(tabIndex int8)

	// Title prop
	// js:"title"
	Title() (title string)

	// Title prop
	SetTitle(title string)
}
