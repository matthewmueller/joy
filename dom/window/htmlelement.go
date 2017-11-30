package window

import "github.com/matthewmueller/golly/dom/domstringmap"

// HTMLElement interface
// js:"HTMLElement"
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

	// SetAccessKey prop
	// js:"accessKey"
	SetAccessKey(accessKey string)

	// Children prop
	// js:"children"
	Children() (children HTMLCollection)

	// ContentEditable prop
	// js:"contentEditable"
	ContentEditable() (contentEditable string)

	// SetContentEditable prop
	// js:"contentEditable"
	SetContentEditable(contentEditable string)

	// Dataset prop
	// js:"dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// Dir prop
	// js:"dir"
	Dir() (dir string)

	// SetDir prop
	// js:"dir"
	SetDir(dir string)

	// Draggable prop
	// js:"draggable"
	Draggable() (draggable bool)

	// SetDraggable prop
	// js:"draggable"
	SetDraggable(draggable bool)

	// Hidden prop
	// js:"hidden"
	Hidden() (hidden bool)

	// SetHidden prop
	// js:"hidden"
	SetHidden(hidden bool)

	// HideFocus prop
	// js:"hideFocus"
	HideFocus() (hideFocus bool)

	// SetHideFocus prop
	// js:"hideFocus"
	SetHideFocus(hideFocus bool)

	// InnerText prop
	// js:"innerText"
	InnerText() (innerText string)

	// SetInnerText prop
	// js:"innerText"
	SetInnerText(innerText string)

	// IsContentEditable prop
	// js:"isContentEditable"
	IsContentEditable() (isContentEditable bool)

	// Lang prop
	// js:"lang"
	Lang() (lang string)

	// SetLang prop
	// js:"lang"
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

	// SetOnabort prop
	// js:"onabort"
	SetOnabort(onabort func(Event))

	// Onactivate prop
	// js:"onactivate"
	Onactivate() (onactivate func(UIEvent))

	// SetOnactivate prop
	// js:"onactivate"
	SetOnactivate(onactivate func(UIEvent))

	// Onbeforeactivate prop
	// js:"onbeforeactivate"
	Onbeforeactivate() (onbeforeactivate func(UIEvent))

	// SetOnbeforeactivate prop
	// js:"onbeforeactivate"
	SetOnbeforeactivate(onbeforeactivate func(UIEvent))

	// Onbeforecopy prop
	// js:"onbeforecopy"
	Onbeforecopy() (onbeforecopy func(*ClipboardEvent))

	// SetOnbeforecopy prop
	// js:"onbeforecopy"
	SetOnbeforecopy(onbeforecopy func(*ClipboardEvent))

	// Onbeforecut prop
	// js:"onbeforecut"
	Onbeforecut() (onbeforecut func(*ClipboardEvent))

	// SetOnbeforecut prop
	// js:"onbeforecut"
	SetOnbeforecut(onbeforecut func(*ClipboardEvent))

	// Onbeforedeactivate prop
	// js:"onbeforedeactivate"
	Onbeforedeactivate() (onbeforedeactivate func(UIEvent))

	// SetOnbeforedeactivate prop
	// js:"onbeforedeactivate"
	SetOnbeforedeactivate(onbeforedeactivate func(UIEvent))

	// Onbeforepaste prop
	// js:"onbeforepaste"
	Onbeforepaste() (onbeforepaste func(*ClipboardEvent))

	// SetOnbeforepaste prop
	// js:"onbeforepaste"
	SetOnbeforepaste(onbeforepaste func(*ClipboardEvent))

	// Onblur prop
	// js:"onblur"
	Onblur() (onblur func(*FocusEvent))

	// SetOnblur prop
	// js:"onblur"
	SetOnblur(onblur func(*FocusEvent))

	// Oncanplay prop
	// js:"oncanplay"
	Oncanplay() (oncanplay func(Event))

	// SetOncanplay prop
	// js:"oncanplay"
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough prop
	// js:"oncanplaythrough"
	Oncanplaythrough() (oncanplaythrough func(Event))

	// SetOncanplaythrough prop
	// js:"oncanplaythrough"
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange prop
	// js:"onchange"
	Onchange() (onchange func(Event))

	// SetOnchange prop
	// js:"onchange"
	SetOnchange(onchange func(Event))

	// Onclick prop
	// js:"onclick"
	Onclick() (onclick func(MouseEvent))

	// SetOnclick prop
	// js:"onclick"
	SetOnclick(onclick func(MouseEvent))

	// Oncontextmenu prop
	// js:"oncontextmenu"
	Oncontextmenu() (oncontextmenu func(*PointerEvent))

	// SetOncontextmenu prop
	// js:"oncontextmenu"
	SetOncontextmenu(oncontextmenu func(*PointerEvent))

	// Oncopy prop
	// js:"oncopy"
	Oncopy() (oncopy func(*ClipboardEvent))

	// SetOncopy prop
	// js:"oncopy"
	SetOncopy(oncopy func(*ClipboardEvent))

	// Oncuechange prop
	// js:"oncuechange"
	Oncuechange() (oncuechange func(Event))

	// SetOncuechange prop
	// js:"oncuechange"
	SetOncuechange(oncuechange func(Event))

	// Oncut prop
	// js:"oncut"
	Oncut() (oncut func(*ClipboardEvent))

	// SetOncut prop
	// js:"oncut"
	SetOncut(oncut func(*ClipboardEvent))

	// Ondblclick prop
	// js:"ondblclick"
	Ondblclick() (ondblclick func(MouseEvent))

	// SetOndblclick prop
	// js:"ondblclick"
	SetOndblclick(ondblclick func(MouseEvent))

	// Ondeactivate prop
	// js:"ondeactivate"
	Ondeactivate() (ondeactivate func(UIEvent))

	// SetOndeactivate prop
	// js:"ondeactivate"
	SetOndeactivate(ondeactivate func(UIEvent))

	// Ondrag prop
	// js:"ondrag"
	Ondrag() (ondrag func(*DragEvent))

	// SetOndrag prop
	// js:"ondrag"
	SetOndrag(ondrag func(*DragEvent))

	// Ondragend prop
	// js:"ondragend"
	Ondragend() (ondragend func(*DragEvent))

	// SetOndragend prop
	// js:"ondragend"
	SetOndragend(ondragend func(*DragEvent))

	// Ondragenter prop
	// js:"ondragenter"
	Ondragenter() (ondragenter func(*DragEvent))

	// SetOndragenter prop
	// js:"ondragenter"
	SetOndragenter(ondragenter func(*DragEvent))

	// Ondragleave prop
	// js:"ondragleave"
	Ondragleave() (ondragleave func(*DragEvent))

	// SetOndragleave prop
	// js:"ondragleave"
	SetOndragleave(ondragleave func(*DragEvent))

	// Ondragover prop
	// js:"ondragover"
	Ondragover() (ondragover func(*DragEvent))

	// SetOndragover prop
	// js:"ondragover"
	SetOndragover(ondragover func(*DragEvent))

	// Ondragstart prop
	// js:"ondragstart"
	Ondragstart() (ondragstart func(*DragEvent))

	// SetOndragstart prop
	// js:"ondragstart"
	SetOndragstart(ondragstart func(*DragEvent))

	// Ondrop prop
	// js:"ondrop"
	Ondrop() (ondrop func(*DragEvent))

	// SetOndrop prop
	// js:"ondrop"
	SetOndrop(ondrop func(*DragEvent))

	// Ondurationchange prop
	// js:"ondurationchange"
	Ondurationchange() (ondurationchange func(Event))

	// SetOndurationchange prop
	// js:"ondurationchange"
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied prop
	// js:"onemptied"
	Onemptied() (onemptied func(Event))

	// SetOnemptied prop
	// js:"onemptied"
	SetOnemptied(onemptied func(Event))

	// Onended prop
	// js:"onended"
	Onended() (onended func(Event))

	// SetOnended prop
	// js:"onended"
	SetOnended(onended func(Event))

	// Onerror prop
	// js:"onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop
	// js:"onerror"
	SetOnerror(onerror func(Event))

	// Onfocus prop
	// js:"onfocus"
	Onfocus() (onfocus func(*FocusEvent))

	// SetOnfocus prop
	// js:"onfocus"
	SetOnfocus(onfocus func(*FocusEvent))

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

	// Onkeydown prop
	// js:"onkeydown"
	Onkeydown() (onkeydown func(*KeyboardEvent))

	// SetOnkeydown prop
	// js:"onkeydown"
	SetOnkeydown(onkeydown func(*KeyboardEvent))

	// Onkeypress prop
	// js:"onkeypress"
	Onkeypress() (onkeypress func(*KeyboardEvent))

	// SetOnkeypress prop
	// js:"onkeypress"
	SetOnkeypress(onkeypress func(*KeyboardEvent))

	// Onkeyup prop
	// js:"onkeyup"
	Onkeyup() (onkeyup func(*KeyboardEvent))

	// SetOnkeyup prop
	// js:"onkeyup"
	SetOnkeyup(onkeyup func(*KeyboardEvent))

	// Onload prop
	// js:"onload"
	Onload() (onload func(Event))

	// SetOnload prop
	// js:"onload"
	SetOnload(onload func(Event))

	// Onloadeddata prop
	// js:"onloadeddata"
	Onloadeddata() (onloadeddata func(Event))

	// SetOnloadeddata prop
	// js:"onloadeddata"
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata prop
	// js:"onloadedmetadata"
	Onloadedmetadata() (onloadedmetadata func(Event))

	// SetOnloadedmetadata prop
	// js:"onloadedmetadata"
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart prop
	// js:"onloadstart"
	Onloadstart() (onloadstart func(Event))

	// SetOnloadstart prop
	// js:"onloadstart"
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown prop
	// js:"onmousedown"
	Onmousedown() (onmousedown func(MouseEvent))

	// SetOnmousedown prop
	// js:"onmousedown"
	SetOnmousedown(onmousedown func(MouseEvent))

	// Onmouseenter prop
	// js:"onmouseenter"
	Onmouseenter() (onmouseenter func(MouseEvent))

	// SetOnmouseenter prop
	// js:"onmouseenter"
	SetOnmouseenter(onmouseenter func(MouseEvent))

	// Onmouseleave prop
	// js:"onmouseleave"
	Onmouseleave() (onmouseleave func(MouseEvent))

	// SetOnmouseleave prop
	// js:"onmouseleave"
	SetOnmouseleave(onmouseleave func(MouseEvent))

	// Onmousemove prop
	// js:"onmousemove"
	Onmousemove() (onmousemove func(MouseEvent))

	// SetOnmousemove prop
	// js:"onmousemove"
	SetOnmousemove(onmousemove func(MouseEvent))

	// Onmouseout prop
	// js:"onmouseout"
	Onmouseout() (onmouseout func(MouseEvent))

	// SetOnmouseout prop
	// js:"onmouseout"
	SetOnmouseout(onmouseout func(MouseEvent))

	// Onmouseover prop
	// js:"onmouseover"
	Onmouseover() (onmouseover func(MouseEvent))

	// SetOnmouseover prop
	// js:"onmouseover"
	SetOnmouseover(onmouseover func(MouseEvent))

	// Onmouseup prop
	// js:"onmouseup"
	Onmouseup() (onmouseup func(MouseEvent))

	// SetOnmouseup prop
	// js:"onmouseup"
	SetOnmouseup(onmouseup func(MouseEvent))

	// Onmousewheel prop
	// js:"onmousewheel"
	Onmousewheel() (onmousewheel func(*WheelEvent))

	// SetOnmousewheel prop
	// js:"onmousewheel"
	SetOnmousewheel(onmousewheel func(*WheelEvent))

	// Onmscontentzoom prop
	// js:"onmscontentzoom"
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// SetOnmscontentzoom prop
	// js:"onmscontentzoom"
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// SetOnmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged"
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onpaste prop
	// js:"onpaste"
	Onpaste() (onpaste func(*ClipboardEvent))

	// SetOnpaste prop
	// js:"onpaste"
	SetOnpaste(onpaste func(*ClipboardEvent))

	// Onpause prop
	// js:"onpause"
	Onpause() (onpause func(Event))

	// SetOnpause prop
	// js:"onpause"
	SetOnpause(onpause func(Event))

	// Onplay prop
	// js:"onplay"
	Onplay() (onplay func(Event))

	// SetOnplay prop
	// js:"onplay"
	SetOnplay(onplay func(Event))

	// Onplaying prop
	// js:"onplaying"
	Onplaying() (onplaying func(Event))

	// SetOnplaying prop
	// js:"onplaying"
	SetOnplaying(onplaying func(Event))

	// Onprogress prop
	// js:"onprogress"
	Onprogress() (onprogress func(Event))

	// SetOnprogress prop
	// js:"onprogress"
	SetOnprogress(onprogress func(Event))

	// Onratechange prop
	// js:"onratechange"
	Onratechange() (onratechange func(Event))

	// SetOnratechange prop
	// js:"onratechange"
	SetOnratechange(onratechange func(Event))

	// Onreset prop
	// js:"onreset"
	Onreset() (onreset func(Event))

	// SetOnreset prop
	// js:"onreset"
	SetOnreset(onreset func(Event))

	// Onscroll prop
	// js:"onscroll"
	Onscroll() (onscroll func(UIEvent))

	// SetOnscroll prop
	// js:"onscroll"
	SetOnscroll(onscroll func(UIEvent))

	// Onseeked prop
	// js:"onseeked"
	Onseeked() (onseeked func(Event))

	// SetOnseeked prop
	// js:"onseeked"
	SetOnseeked(onseeked func(Event))

	// Onseeking prop
	// js:"onseeking"
	Onseeking() (onseeking func(Event))

	// SetOnseeking prop
	// js:"onseeking"
	SetOnseeking(onseeking func(Event))

	// Onselect prop
	// js:"onselect"
	Onselect() (onselect func(UIEvent))

	// SetOnselect prop
	// js:"onselect"
	SetOnselect(onselect func(UIEvent))

	// Onselectstart prop
	// js:"onselectstart"
	Onselectstart() (onselectstart func(Event))

	// SetOnselectstart prop
	// js:"onselectstart"
	SetOnselectstart(onselectstart func(Event))

	// Onstalled prop
	// js:"onstalled"
	Onstalled() (onstalled func(Event))

	// SetOnstalled prop
	// js:"onstalled"
	SetOnstalled(onstalled func(Event))

	// Onsubmit prop
	// js:"onsubmit"
	Onsubmit() (onsubmit func(Event))

	// SetOnsubmit prop
	// js:"onsubmit"
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend prop
	// js:"onsuspend"
	Onsuspend() (onsuspend func(Event))

	// SetOnsuspend prop
	// js:"onsuspend"
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate prop
	// js:"ontimeupdate"
	Ontimeupdate() (ontimeupdate func(Event))

	// SetOntimeupdate prop
	// js:"ontimeupdate"
	SetOntimeupdate(ontimeupdate func(Event))

	// Onvolumechange prop
	// js:"onvolumechange"
	Onvolumechange() (onvolumechange func(Event))

	// SetOnvolumechange prop
	// js:"onvolumechange"
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting prop
	// js:"onwaiting"
	Onwaiting() (onwaiting func(Event))

	// SetOnwaiting prop
	// js:"onwaiting"
	SetOnwaiting(onwaiting func(Event))

	// OuterText prop
	// js:"outerText"
	OuterText() (outerText string)

	// SetOuterText prop
	// js:"outerText"
	SetOuterText(outerText string)

	// Spellcheck prop
	// js:"spellcheck"
	Spellcheck() (spellcheck bool)

	// SetSpellcheck prop
	// js:"spellcheck"
	SetSpellcheck(spellcheck bool)

	// Style prop
	// js:"style"
	Style() (style *CSSStyleDeclaration)

	// TabIndex prop
	// js:"tabIndex"
	TabIndex() (tabIndex int8)

	// SetTabIndex prop
	// js:"tabIndex"
	SetTabIndex(tabIndex int8)

	// Title prop
	// js:"title"
	Title() (title string)

	// SetTitle prop
	// js:"title"
	SetTitle(title string)
}
