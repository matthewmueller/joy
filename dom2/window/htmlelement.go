package window

// js:"HTMLElement,omit"
type HTMLElement interface {
	Element

	// Blur
	Blur()

	// Click
	Click()

	// DragDrop
	DragDrop() (b bool)

	// Focus
	Focus()

	// GetElementsByClassName
	GetElementsByClassName(classNames string) (n *NodeList)

	// InsertAdjacentElement
	InsertAdjacentElement(position string, insertedElement Element) (e Element)

	// InsertAdjacentHTML
	InsertAdjacentHTML(where string, html string)

	// InsertAdjacentText
	InsertAdjacentText(where string, text string)

	// MsGetInputContext
	MsGetInputContext() (m *MSInputMethodContext)

	// ScrollIntoView
	ScrollIntoView(top *bool)

	// AccessKey
	AccessKey() (accessKey string)

	// AccessKey
	SetAccessKey(accessKey string)

	// Children
	Children() (children HTMLCollection)

	// ContentEditable
	ContentEditable() (contentEditable string)

	// ContentEditable
	SetContentEditable(contentEditable string)

	// Dataset
	Dataset() (dataset *domstringmap.DOMStringMap)

	// Dir
	Dir() (dir string)

	// Dir
	SetDir(dir string)

	// Draggable
	Draggable() (draggable bool)

	// Draggable
	SetDraggable(draggable bool)

	// Hidden
	Hidden() (hidden bool)

	// Hidden
	SetHidden(hidden bool)

	// HideFocus
	HideFocus() (hideFocus bool)

	// HideFocus
	SetHideFocus(hideFocus bool)

	// InnerText
	InnerText() (innerText string)

	// InnerText
	SetInnerText(innerText string)

	// IsContentEditable
	IsContentEditable() (isContentEditable bool)

	// Lang
	Lang() (lang string)

	// Lang
	SetLang(lang string)

	// OffsetHeight
	OffsetHeight() (offsetHeight int)

	// OffsetLeft
	OffsetLeft() (offsetLeft int)

	// OffsetParent
	OffsetParent() (offsetParent Element)

	// OffsetTop
	OffsetTop() (offsetTop int)

	// OffsetWidth
	OffsetWidth() (offsetWidth int)

	// Onabort
	Onabort() (onabort func(Event))

	// Onabort
	SetOnabort(onabort func(Event))

	// Onactivate
	Onactivate() (onactivate func(UIEvent))

	// Onactivate
	SetOnactivate(onactivate func(UIEvent))

	// Onbeforeactivate
	Onbeforeactivate() (onbeforeactivate func(UIEvent))

	// Onbeforeactivate
	SetOnbeforeactivate(onbeforeactivate func(UIEvent))

	// Onbeforecopy
	Onbeforecopy() (onbeforecopy func(*ClipboardEvent))

	// Onbeforecopy
	SetOnbeforecopy(onbeforecopy func(*ClipboardEvent))

	// Onbeforecut
	Onbeforecut() (onbeforecut func(*ClipboardEvent))

	// Onbeforecut
	SetOnbeforecut(onbeforecut func(*ClipboardEvent))

	// Onbeforedeactivate
	Onbeforedeactivate() (onbeforedeactivate func(UIEvent))

	// Onbeforedeactivate
	SetOnbeforedeactivate(onbeforedeactivate func(UIEvent))

	// Onbeforepaste
	Onbeforepaste() (onbeforepaste func(*ClipboardEvent))

	// Onbeforepaste
	SetOnbeforepaste(onbeforepaste func(*ClipboardEvent))

	// Onblur
	Onblur() (onblur func(*FocusEvent))

	// Onblur
	SetOnblur(onblur func(*FocusEvent))

	// Oncanplay
	Oncanplay() (oncanplay func(Event))

	// Oncanplay
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough
	Oncanplaythrough() (oncanplaythrough func(Event))

	// Oncanplaythrough
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange
	Onchange() (onchange func(Event))

	// Onchange
	SetOnchange(onchange func(Event))

	// Onclick
	Onclick() (onclick func(MouseEvent))

	// Onclick
	SetOnclick(onclick func(MouseEvent))

	// Oncontextmenu
	Oncontextmenu() (oncontextmenu func(*PointerEvent))

	// Oncontextmenu
	SetOncontextmenu(oncontextmenu func(*PointerEvent))

	// Oncopy
	Oncopy() (oncopy func(*ClipboardEvent))

	// Oncopy
	SetOncopy(oncopy func(*ClipboardEvent))

	// Oncuechange
	Oncuechange() (oncuechange func(Event))

	// Oncuechange
	SetOncuechange(oncuechange func(Event))

	// Oncut
	Oncut() (oncut func(*ClipboardEvent))

	// Oncut
	SetOncut(oncut func(*ClipboardEvent))

	// Ondblclick
	Ondblclick() (ondblclick func(MouseEvent))

	// Ondblclick
	SetOndblclick(ondblclick func(MouseEvent))

	// Ondeactivate
	Ondeactivate() (ondeactivate func(UIEvent))

	// Ondeactivate
	SetOndeactivate(ondeactivate func(UIEvent))

	// Ondrag
	Ondrag() (ondrag func(*DragEvent))

	// Ondrag
	SetOndrag(ondrag func(*DragEvent))

	// Ondragend
	Ondragend() (ondragend func(*DragEvent))

	// Ondragend
	SetOndragend(ondragend func(*DragEvent))

	// Ondragenter
	Ondragenter() (ondragenter func(*DragEvent))

	// Ondragenter
	SetOndragenter(ondragenter func(*DragEvent))

	// Ondragleave
	Ondragleave() (ondragleave func(*DragEvent))

	// Ondragleave
	SetOndragleave(ondragleave func(*DragEvent))

	// Ondragover
	Ondragover() (ondragover func(*DragEvent))

	// Ondragover
	SetOndragover(ondragover func(*DragEvent))

	// Ondragstart
	Ondragstart() (ondragstart func(*DragEvent))

	// Ondragstart
	SetOndragstart(ondragstart func(*DragEvent))

	// Ondrop
	Ondrop() (ondrop func(*DragEvent))

	// Ondrop
	SetOndrop(ondrop func(*DragEvent))

	// Ondurationchange
	Ondurationchange() (ondurationchange func(Event))

	// Ondurationchange
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied
	Onemptied() (onemptied func(Event))

	// Onemptied
	SetOnemptied(onemptied func(Event))

	// Onended
	Onended() (onended func(Event))

	// Onended
	SetOnended(onended func(Event))

	// Onerror
	Onerror() (onerror func(Event))

	// Onerror
	SetOnerror(onerror func(Event))

	// Onfocus
	Onfocus() (onfocus func(*FocusEvent))

	// Onfocus
	SetOnfocus(onfocus func(*FocusEvent))

	// Oninput
	Oninput() (oninput func(Event))

	// Oninput
	SetOninput(oninput func(Event))

	// Oninvalid
	Oninvalid() (oninvalid func(Event))

	// Oninvalid
	SetOninvalid(oninvalid func(Event))

	// Onkeydown
	Onkeydown() (onkeydown func(*KeyboardEvent))

	// Onkeydown
	SetOnkeydown(onkeydown func(*KeyboardEvent))

	// Onkeypress
	Onkeypress() (onkeypress func(*KeyboardEvent))

	// Onkeypress
	SetOnkeypress(onkeypress func(*KeyboardEvent))

	// Onkeyup
	Onkeyup() (onkeyup func(*KeyboardEvent))

	// Onkeyup
	SetOnkeyup(onkeyup func(*KeyboardEvent))

	// Onload
	Onload() (onload func(Event))

	// Onload
	SetOnload(onload func(Event))

	// Onloadeddata
	Onloadeddata() (onloadeddata func(Event))

	// Onloadeddata
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata
	Onloadedmetadata() (onloadedmetadata func(Event))

	// Onloadedmetadata
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart
	Onloadstart() (onloadstart func(Event))

	// Onloadstart
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown
	Onmousedown() (onmousedown func(MouseEvent))

	// Onmousedown
	SetOnmousedown(onmousedown func(MouseEvent))

	// Onmouseenter
	Onmouseenter() (onmouseenter func(MouseEvent))

	// Onmouseenter
	SetOnmouseenter(onmouseenter func(MouseEvent))

	// Onmouseleave
	Onmouseleave() (onmouseleave func(MouseEvent))

	// Onmouseleave
	SetOnmouseleave(onmouseleave func(MouseEvent))

	// Onmousemove
	Onmousemove() (onmousemove func(MouseEvent))

	// Onmousemove
	SetOnmousemove(onmousemove func(MouseEvent))

	// Onmouseout
	Onmouseout() (onmouseout func(MouseEvent))

	// Onmouseout
	SetOnmouseout(onmouseout func(MouseEvent))

	// Onmouseover
	Onmouseover() (onmouseover func(MouseEvent))

	// Onmouseover
	SetOnmouseover(onmouseover func(MouseEvent))

	// Onmouseup
	Onmouseup() (onmouseup func(MouseEvent))

	// Onmouseup
	SetOnmouseup(onmouseup func(MouseEvent))

	// Onmousewheel
	Onmousewheel() (onmousewheel func(*WheelEvent))

	// Onmousewheel
	SetOnmousewheel(onmousewheel func(*WheelEvent))

	// Onmscontentzoom
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// Onmscontentzoom
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsmanipulationstatechanged
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmsmanipulationstatechanged
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onpaste
	Onpaste() (onpaste func(*ClipboardEvent))

	// Onpaste
	SetOnpaste(onpaste func(*ClipboardEvent))

	// Onpause
	Onpause() (onpause func(Event))

	// Onpause
	SetOnpause(onpause func(Event))

	// Onplay
	Onplay() (onplay func(Event))

	// Onplay
	SetOnplay(onplay func(Event))

	// Onplaying
	Onplaying() (onplaying func(Event))

	// Onplaying
	SetOnplaying(onplaying func(Event))

	// Onprogress
	Onprogress() (onprogress func(Event))

	// Onprogress
	SetOnprogress(onprogress func(Event))

	// Onratechange
	Onratechange() (onratechange func(Event))

	// Onratechange
	SetOnratechange(onratechange func(Event))

	// Onreset
	Onreset() (onreset func(Event))

	// Onreset
	SetOnreset(onreset func(Event))

	// Onscroll
	Onscroll() (onscroll func(UIEvent))

	// Onscroll
	SetOnscroll(onscroll func(UIEvent))

	// Onseeked
	Onseeked() (onseeked func(Event))

	// Onseeked
	SetOnseeked(onseeked func(Event))

	// Onseeking
	Onseeking() (onseeking func(Event))

	// Onseeking
	SetOnseeking(onseeking func(Event))

	// Onselect
	Onselect() (onselect func(UIEvent))

	// Onselect
	SetOnselect(onselect func(UIEvent))

	// Onselectstart
	Onselectstart() (onselectstart func(Event))

	// Onselectstart
	SetOnselectstart(onselectstart func(Event))

	// Onstalled
	Onstalled() (onstalled func(Event))

	// Onstalled
	SetOnstalled(onstalled func(Event))

	// Onsubmit
	Onsubmit() (onsubmit func(Event))

	// Onsubmit
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend
	Onsuspend() (onsuspend func(Event))

	// Onsuspend
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate
	Ontimeupdate() (ontimeupdate func(Event))

	// Ontimeupdate
	SetOntimeupdate(ontimeupdate func(Event))

	// Onvolumechange
	Onvolumechange() (onvolumechange func(Event))

	// Onvolumechange
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting
	Onwaiting() (onwaiting func(Event))

	// Onwaiting
	SetOnwaiting(onwaiting func(Event))

	// OuterText
	OuterText() (outerText string)

	// OuterText
	SetOuterText(outerText string)

	// Spellcheck
	Spellcheck() (spellcheck bool)

	// Spellcheck
	SetSpellcheck(spellcheck bool)

	// Style
	Style() (style *CSSStyleDeclaration)

	// TabIndex
	TabIndex() (tabIndex int8)

	// TabIndex
	SetTabIndex(tabIndex int8)

	// Title
	Title() (title string)

	// Title
	SetTitle(title string)
}
