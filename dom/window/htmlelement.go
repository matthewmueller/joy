package window

import "github.com/matthewmueller/golly/dom/domstringmap"

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
	// js:"setaccessKey"
	SetAccessKey(accessKey string)

	// Children prop
	// js:"children"
	Children() (children HTMLCollection)

	// ContentEditable prop
	// js:"contentEditable"
	ContentEditable() (contentEditable string)

	// ContentEditable prop
	// js:"setcontentEditable"
	SetContentEditable(contentEditable string)

	// Dataset prop
	// js:"dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// Dir prop
	// js:"dir"
	Dir() (dir string)

	// Dir prop
	// js:"setdir"
	SetDir(dir string)

	// Draggable prop
	// js:"draggable"
	Draggable() (draggable bool)

	// Draggable prop
	// js:"setdraggable"
	SetDraggable(draggable bool)

	// Hidden prop
	// js:"hidden"
	Hidden() (hidden bool)

	// Hidden prop
	// js:"sethidden"
	SetHidden(hidden bool)

	// HideFocus prop
	// js:"hideFocus"
	HideFocus() (hideFocus bool)

	// HideFocus prop
	// js:"sethideFocus"
	SetHideFocus(hideFocus bool)

	// InnerText prop
	// js:"innerText"
	InnerText() (innerText string)

	// InnerText prop
	// js:"setinnerText"
	SetInnerText(innerText string)

	// IsContentEditable prop
	// js:"isContentEditable"
	IsContentEditable() (isContentEditable bool)

	// Lang prop
	// js:"lang"
	Lang() (lang string)

	// Lang prop
	// js:"setlang"
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
	// js:"setonabort"
	SetOnabort(onabort func(Event))

	// Onactivate prop
	// js:"onactivate"
	Onactivate() (onactivate func(UIEvent))

	// Onactivate prop
	// js:"setonactivate"
	SetOnactivate(onactivate func(UIEvent))

	// Onbeforeactivate prop
	// js:"onbeforeactivate"
	Onbeforeactivate() (onbeforeactivate func(UIEvent))

	// Onbeforeactivate prop
	// js:"setonbeforeactivate"
	SetOnbeforeactivate(onbeforeactivate func(UIEvent))

	// Onbeforecopy prop
	// js:"onbeforecopy"
	Onbeforecopy() (onbeforecopy func(*ClipboardEvent))

	// Onbeforecopy prop
	// js:"setonbeforecopy"
	SetOnbeforecopy(onbeforecopy func(*ClipboardEvent))

	// Onbeforecut prop
	// js:"onbeforecut"
	Onbeforecut() (onbeforecut func(*ClipboardEvent))

	// Onbeforecut prop
	// js:"setonbeforecut"
	SetOnbeforecut(onbeforecut func(*ClipboardEvent))

	// Onbeforedeactivate prop
	// js:"onbeforedeactivate"
	Onbeforedeactivate() (onbeforedeactivate func(UIEvent))

	// Onbeforedeactivate prop
	// js:"setonbeforedeactivate"
	SetOnbeforedeactivate(onbeforedeactivate func(UIEvent))

	// Onbeforepaste prop
	// js:"onbeforepaste"
	Onbeforepaste() (onbeforepaste func(*ClipboardEvent))

	// Onbeforepaste prop
	// js:"setonbeforepaste"
	SetOnbeforepaste(onbeforepaste func(*ClipboardEvent))

	// Onblur prop
	// js:"onblur"
	Onblur() (onblur func(*FocusEvent))

	// Onblur prop
	// js:"setonblur"
	SetOnblur(onblur func(*FocusEvent))

	// Oncanplay prop
	// js:"oncanplay"
	Oncanplay() (oncanplay func(Event))

	// Oncanplay prop
	// js:"setoncanplay"
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough prop
	// js:"oncanplaythrough"
	Oncanplaythrough() (oncanplaythrough func(Event))

	// Oncanplaythrough prop
	// js:"setoncanplaythrough"
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange prop
	// js:"onchange"
	Onchange() (onchange func(Event))

	// Onchange prop
	// js:"setonchange"
	SetOnchange(onchange func(Event))

	// Onclick prop
	// js:"onclick"
	Onclick() (onclick func(MouseEvent))

	// Onclick prop
	// js:"setonclick"
	SetOnclick(onclick func(MouseEvent))

	// Oncontextmenu prop
	// js:"oncontextmenu"
	Oncontextmenu() (oncontextmenu func(*PointerEvent))

	// Oncontextmenu prop
	// js:"setoncontextmenu"
	SetOncontextmenu(oncontextmenu func(*PointerEvent))

	// Oncopy prop
	// js:"oncopy"
	Oncopy() (oncopy func(*ClipboardEvent))

	// Oncopy prop
	// js:"setoncopy"
	SetOncopy(oncopy func(*ClipboardEvent))

	// Oncuechange prop
	// js:"oncuechange"
	Oncuechange() (oncuechange func(Event))

	// Oncuechange prop
	// js:"setoncuechange"
	SetOncuechange(oncuechange func(Event))

	// Oncut prop
	// js:"oncut"
	Oncut() (oncut func(*ClipboardEvent))

	// Oncut prop
	// js:"setoncut"
	SetOncut(oncut func(*ClipboardEvent))

	// Ondblclick prop
	// js:"ondblclick"
	Ondblclick() (ondblclick func(MouseEvent))

	// Ondblclick prop
	// js:"setondblclick"
	SetOndblclick(ondblclick func(MouseEvent))

	// Ondeactivate prop
	// js:"ondeactivate"
	Ondeactivate() (ondeactivate func(UIEvent))

	// Ondeactivate prop
	// js:"setondeactivate"
	SetOndeactivate(ondeactivate func(UIEvent))

	// Ondrag prop
	// js:"ondrag"
	Ondrag() (ondrag func(*DragEvent))

	// Ondrag prop
	// js:"setondrag"
	SetOndrag(ondrag func(*DragEvent))

	// Ondragend prop
	// js:"ondragend"
	Ondragend() (ondragend func(*DragEvent))

	// Ondragend prop
	// js:"setondragend"
	SetOndragend(ondragend func(*DragEvent))

	// Ondragenter prop
	// js:"ondragenter"
	Ondragenter() (ondragenter func(*DragEvent))

	// Ondragenter prop
	// js:"setondragenter"
	SetOndragenter(ondragenter func(*DragEvent))

	// Ondragleave prop
	// js:"ondragleave"
	Ondragleave() (ondragleave func(*DragEvent))

	// Ondragleave prop
	// js:"setondragleave"
	SetOndragleave(ondragleave func(*DragEvent))

	// Ondragover prop
	// js:"ondragover"
	Ondragover() (ondragover func(*DragEvent))

	// Ondragover prop
	// js:"setondragover"
	SetOndragover(ondragover func(*DragEvent))

	// Ondragstart prop
	// js:"ondragstart"
	Ondragstart() (ondragstart func(*DragEvent))

	// Ondragstart prop
	// js:"setondragstart"
	SetOndragstart(ondragstart func(*DragEvent))

	// Ondrop prop
	// js:"ondrop"
	Ondrop() (ondrop func(*DragEvent))

	// Ondrop prop
	// js:"setondrop"
	SetOndrop(ondrop func(*DragEvent))

	// Ondurationchange prop
	// js:"ondurationchange"
	Ondurationchange() (ondurationchange func(Event))

	// Ondurationchange prop
	// js:"setondurationchange"
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied prop
	// js:"onemptied"
	Onemptied() (onemptied func(Event))

	// Onemptied prop
	// js:"setonemptied"
	SetOnemptied(onemptied func(Event))

	// Onended prop
	// js:"onended"
	Onended() (onended func(Event))

	// Onended prop
	// js:"setonended"
	SetOnended(onended func(Event))

	// Onerror prop
	// js:"onerror"
	Onerror() (onerror func(Event))

	// Onerror prop
	// js:"setonerror"
	SetOnerror(onerror func(Event))

	// Onfocus prop
	// js:"onfocus"
	Onfocus() (onfocus func(*FocusEvent))

	// Onfocus prop
	// js:"setonfocus"
	SetOnfocus(onfocus func(*FocusEvent))

	// Oninput prop
	// js:"oninput"
	Oninput() (oninput func(Event))

	// Oninput prop
	// js:"setoninput"
	SetOninput(oninput func(Event))

	// Oninvalid prop
	// js:"oninvalid"
	Oninvalid() (oninvalid func(Event))

	// Oninvalid prop
	// js:"setoninvalid"
	SetOninvalid(oninvalid func(Event))

	// Onkeydown prop
	// js:"onkeydown"
	Onkeydown() (onkeydown func(*KeyboardEvent))

	// Onkeydown prop
	// js:"setonkeydown"
	SetOnkeydown(onkeydown func(*KeyboardEvent))

	// Onkeypress prop
	// js:"onkeypress"
	Onkeypress() (onkeypress func(*KeyboardEvent))

	// Onkeypress prop
	// js:"setonkeypress"
	SetOnkeypress(onkeypress func(*KeyboardEvent))

	// Onkeyup prop
	// js:"onkeyup"
	Onkeyup() (onkeyup func(*KeyboardEvent))

	// Onkeyup prop
	// js:"setonkeyup"
	SetOnkeyup(onkeyup func(*KeyboardEvent))

	// Onload prop
	// js:"onload"
	Onload() (onload func(Event))

	// Onload prop
	// js:"setonload"
	SetOnload(onload func(Event))

	// Onloadeddata prop
	// js:"onloadeddata"
	Onloadeddata() (onloadeddata func(Event))

	// Onloadeddata prop
	// js:"setonloadeddata"
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata prop
	// js:"onloadedmetadata"
	Onloadedmetadata() (onloadedmetadata func(Event))

	// Onloadedmetadata prop
	// js:"setonloadedmetadata"
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart prop
	// js:"onloadstart"
	Onloadstart() (onloadstart func(Event))

	// Onloadstart prop
	// js:"setonloadstart"
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown prop
	// js:"onmousedown"
	Onmousedown() (onmousedown func(MouseEvent))

	// Onmousedown prop
	// js:"setonmousedown"
	SetOnmousedown(onmousedown func(MouseEvent))

	// Onmouseenter prop
	// js:"onmouseenter"
	Onmouseenter() (onmouseenter func(MouseEvent))

	// Onmouseenter prop
	// js:"setonmouseenter"
	SetOnmouseenter(onmouseenter func(MouseEvent))

	// Onmouseleave prop
	// js:"onmouseleave"
	Onmouseleave() (onmouseleave func(MouseEvent))

	// Onmouseleave prop
	// js:"setonmouseleave"
	SetOnmouseleave(onmouseleave func(MouseEvent))

	// Onmousemove prop
	// js:"onmousemove"
	Onmousemove() (onmousemove func(MouseEvent))

	// Onmousemove prop
	// js:"setonmousemove"
	SetOnmousemove(onmousemove func(MouseEvent))

	// Onmouseout prop
	// js:"onmouseout"
	Onmouseout() (onmouseout func(MouseEvent))

	// Onmouseout prop
	// js:"setonmouseout"
	SetOnmouseout(onmouseout func(MouseEvent))

	// Onmouseover prop
	// js:"onmouseover"
	Onmouseover() (onmouseover func(MouseEvent))

	// Onmouseover prop
	// js:"setonmouseover"
	SetOnmouseover(onmouseover func(MouseEvent))

	// Onmouseup prop
	// js:"onmouseup"
	Onmouseup() (onmouseup func(MouseEvent))

	// Onmouseup prop
	// js:"setonmouseup"
	SetOnmouseup(onmouseup func(MouseEvent))

	// Onmousewheel prop
	// js:"onmousewheel"
	Onmousewheel() (onmousewheel func(*WheelEvent))

	// Onmousewheel prop
	// js:"setonmousewheel"
	SetOnmousewheel(onmousewheel func(*WheelEvent))

	// Onmscontentzoom prop
	// js:"onmscontentzoom"
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// Onmscontentzoom prop
	// js:"setonmscontentzoom"
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmsmanipulationstatechanged prop
	// js:"setonmsmanipulationstatechanged"
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onpaste prop
	// js:"onpaste"
	Onpaste() (onpaste func(*ClipboardEvent))

	// Onpaste prop
	// js:"setonpaste"
	SetOnpaste(onpaste func(*ClipboardEvent))

	// Onpause prop
	// js:"onpause"
	Onpause() (onpause func(Event))

	// Onpause prop
	// js:"setonpause"
	SetOnpause(onpause func(Event))

	// Onplay prop
	// js:"onplay"
	Onplay() (onplay func(Event))

	// Onplay prop
	// js:"setonplay"
	SetOnplay(onplay func(Event))

	// Onplaying prop
	// js:"onplaying"
	Onplaying() (onplaying func(Event))

	// Onplaying prop
	// js:"setonplaying"
	SetOnplaying(onplaying func(Event))

	// Onprogress prop
	// js:"onprogress"
	Onprogress() (onprogress func(Event))

	// Onprogress prop
	// js:"setonprogress"
	SetOnprogress(onprogress func(Event))

	// Onratechange prop
	// js:"onratechange"
	Onratechange() (onratechange func(Event))

	// Onratechange prop
	// js:"setonratechange"
	SetOnratechange(onratechange func(Event))

	// Onreset prop
	// js:"onreset"
	Onreset() (onreset func(Event))

	// Onreset prop
	// js:"setonreset"
	SetOnreset(onreset func(Event))

	// Onscroll prop
	// js:"onscroll"
	Onscroll() (onscroll func(UIEvent))

	// Onscroll prop
	// js:"setonscroll"
	SetOnscroll(onscroll func(UIEvent))

	// Onseeked prop
	// js:"onseeked"
	Onseeked() (onseeked func(Event))

	// Onseeked prop
	// js:"setonseeked"
	SetOnseeked(onseeked func(Event))

	// Onseeking prop
	// js:"onseeking"
	Onseeking() (onseeking func(Event))

	// Onseeking prop
	// js:"setonseeking"
	SetOnseeking(onseeking func(Event))

	// Onselect prop
	// js:"onselect"
	Onselect() (onselect func(UIEvent))

	// Onselect prop
	// js:"setonselect"
	SetOnselect(onselect func(UIEvent))

	// Onselectstart prop
	// js:"onselectstart"
	Onselectstart() (onselectstart func(Event))

	// Onselectstart prop
	// js:"setonselectstart"
	SetOnselectstart(onselectstart func(Event))

	// Onstalled prop
	// js:"onstalled"
	Onstalled() (onstalled func(Event))

	// Onstalled prop
	// js:"setonstalled"
	SetOnstalled(onstalled func(Event))

	// Onsubmit prop
	// js:"onsubmit"
	Onsubmit() (onsubmit func(Event))

	// Onsubmit prop
	// js:"setonsubmit"
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend prop
	// js:"onsuspend"
	Onsuspend() (onsuspend func(Event))

	// Onsuspend prop
	// js:"setonsuspend"
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate prop
	// js:"ontimeupdate"
	Ontimeupdate() (ontimeupdate func(Event))

	// Ontimeupdate prop
	// js:"setontimeupdate"
	SetOntimeupdate(ontimeupdate func(Event))

	// Onvolumechange prop
	// js:"onvolumechange"
	Onvolumechange() (onvolumechange func(Event))

	// Onvolumechange prop
	// js:"setonvolumechange"
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting prop
	// js:"onwaiting"
	Onwaiting() (onwaiting func(Event))

	// Onwaiting prop
	// js:"setonwaiting"
	SetOnwaiting(onwaiting func(Event))

	// OuterText prop
	// js:"outerText"
	OuterText() (outerText string)

	// OuterText prop
	// js:"setouterText"
	SetOuterText(outerText string)

	// Spellcheck prop
	// js:"spellcheck"
	Spellcheck() (spellcheck bool)

	// Spellcheck prop
	// js:"setspellcheck"
	SetSpellcheck(spellcheck bool)

	// Style prop
	// js:"style"
	Style() (style *CSSStyleDeclaration)

	// TabIndex prop
	// js:"tabIndex"
	TabIndex() (tabIndex int8)

	// TabIndex prop
	// js:"settabIndex"
	SetTabIndex(tabIndex int8)

	// Title prop
	// js:"title"
	Title() (title string)

	// Title prop
	// js:"settitle"
	SetTitle(title string)
}
