package window

import (
	"github.com/matthewmueller/golly/dom/domstringmap"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLElement,omit"
type HTMLElement interface {
	Element

	// Blur
	// js:"blur,rewrite=blur"
	Blur()

	// Click
	// js:"click,rewrite=click"
	Click()

	// DragDrop
	// js:"dragDrop,rewrite=dragdrop"
	DragDrop() (b bool)

	// Focus
	// js:"focus,rewrite=focus"
	Focus()

	// GetElementsByClassName
	// js:"getElementsByClassName,rewrite=getelementsbyclassname"
	GetElementsByClassName(classNames string) (n *NodeList)

	// InsertAdjacentElement
	// js:"insertAdjacentElement,rewrite=insertadjacentelement"
	InsertAdjacentElement(position string, insertedElement Element) (e Element)

	// InsertAdjacentHTML
	// js:"insertAdjacentHTML,rewrite=insertadjacenthtml"
	InsertAdjacentHTML(where string, html string)

	// InsertAdjacentText
	// js:"insertAdjacentText,rewrite=insertadjacenttext"
	InsertAdjacentText(where string, text string)

	// MsGetInputContext
	// js:"msGetInputContext,rewrite=msgetinputcontext"
	MsGetInputContext() (m *MSInputMethodContext)

	// ScrollIntoView
	// js:"scrollIntoView,rewrite=scrollintoview"
	ScrollIntoView(top *bool)

	// AccessKey prop
	// js:"accessKey,rewrite=accesskey"
	AccessKey() (accessKey string)

	// AccessKey prop
	// js:"setaccessKey,rewrite=setaccesskey"
	SetAccessKey(accessKey string)

	// Children prop
	// js:"children,rewrite=children"
	Children() (children HTMLCollection)

	// ContentEditable prop
	// js:"contentEditable,rewrite=contenteditable"
	ContentEditable() (contentEditable string)

	// ContentEditable prop
	// js:"setcontentEditable,rewrite=setcontenteditable"
	SetContentEditable(contentEditable string)

	// Dataset prop
	// js:"dataset,rewrite=dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// Dir prop
	// js:"dir,rewrite=dir"
	Dir() (dir string)

	// Dir prop
	// js:"setdir,rewrite=setdir"
	SetDir(dir string)

	// Draggable prop
	// js:"draggable,rewrite=draggable"
	Draggable() (draggable bool)

	// Draggable prop
	// js:"setdraggable,rewrite=setdraggable"
	SetDraggable(draggable bool)

	// Hidden prop
	// js:"hidden,rewrite=hidden"
	Hidden() (hidden bool)

	// Hidden prop
	// js:"sethidden,rewrite=sethidden"
	SetHidden(hidden bool)

	// HideFocus prop
	// js:"hideFocus,rewrite=hidefocus"
	HideFocus() (hideFocus bool)

	// HideFocus prop
	// js:"sethideFocus,rewrite=sethidefocus"
	SetHideFocus(hideFocus bool)

	// InnerText prop
	// js:"innerText,rewrite=innertext"
	InnerText() (innerText string)

	// InnerText prop
	// js:"setinnerText,rewrite=setinnertext"
	SetInnerText(innerText string)

	// IsContentEditable prop
	// js:"isContentEditable,rewrite=iscontenteditable"
	IsContentEditable() (isContentEditable bool)

	// Lang prop
	// js:"lang,rewrite=lang"
	Lang() (lang string)

	// Lang prop
	// js:"setlang,rewrite=setlang"
	SetLang(lang string)

	// OffsetHeight prop
	// js:"offsetHeight,rewrite=offsetheight"
	OffsetHeight() (offsetHeight int)

	// OffsetLeft prop
	// js:"offsetLeft,rewrite=offsetleft"
	OffsetLeft() (offsetLeft int)

	// OffsetParent prop
	// js:"offsetParent,rewrite=offsetparent"
	OffsetParent() (offsetParent Element)

	// OffsetTop prop
	// js:"offsetTop,rewrite=offsettop"
	OffsetTop() (offsetTop int)

	// OffsetWidth prop
	// js:"offsetWidth,rewrite=offsetwidth"
	OffsetWidth() (offsetWidth int)

	// Onabort prop
	// js:"onabort,rewrite=onabort"
	Onabort() (onabort func(Event))

	// Onabort prop
	// js:"setonabort,rewrite=setonabort"
	SetOnabort(onabort func(Event))

	// Onactivate prop
	// js:"onactivate,rewrite=onactivate"
	Onactivate() (onactivate func(UIEvent))

	// Onactivate prop
	// js:"setonactivate,rewrite=setonactivate"
	SetOnactivate(onactivate func(UIEvent))

	// Onbeforeactivate prop
	// js:"onbeforeactivate,rewrite=onbeforeactivate"
	Onbeforeactivate() (onbeforeactivate func(UIEvent))

	// Onbeforeactivate prop
	// js:"setonbeforeactivate,rewrite=setonbeforeactivate"
	SetOnbeforeactivate(onbeforeactivate func(UIEvent))

	// Onbeforecopy prop
	// js:"onbeforecopy,rewrite=onbeforecopy"
	Onbeforecopy() (onbeforecopy func(*ClipboardEvent))

	// Onbeforecopy prop
	// js:"setonbeforecopy,rewrite=setonbeforecopy"
	SetOnbeforecopy(onbeforecopy func(*ClipboardEvent))

	// Onbeforecut prop
	// js:"onbeforecut,rewrite=onbeforecut"
	Onbeforecut() (onbeforecut func(*ClipboardEvent))

	// Onbeforecut prop
	// js:"setonbeforecut,rewrite=setonbeforecut"
	SetOnbeforecut(onbeforecut func(*ClipboardEvent))

	// Onbeforedeactivate prop
	// js:"onbeforedeactivate,rewrite=onbeforedeactivate"
	Onbeforedeactivate() (onbeforedeactivate func(UIEvent))

	// Onbeforedeactivate prop
	// js:"setonbeforedeactivate,rewrite=setonbeforedeactivate"
	SetOnbeforedeactivate(onbeforedeactivate func(UIEvent))

	// Onbeforepaste prop
	// js:"onbeforepaste,rewrite=onbeforepaste"
	Onbeforepaste() (onbeforepaste func(*ClipboardEvent))

	// Onbeforepaste prop
	// js:"setonbeforepaste,rewrite=setonbeforepaste"
	SetOnbeforepaste(onbeforepaste func(*ClipboardEvent))

	// Onblur prop
	// js:"onblur,rewrite=onblur"
	Onblur() (onblur func(*FocusEvent))

	// Onblur prop
	// js:"setonblur,rewrite=setonblur"
	SetOnblur(onblur func(*FocusEvent))

	// Oncanplay prop
	// js:"oncanplay,rewrite=oncanplay"
	Oncanplay() (oncanplay func(Event))

	// Oncanplay prop
	// js:"setoncanplay,rewrite=setoncanplay"
	SetOncanplay(oncanplay func(Event))

	// Oncanplaythrough prop
	// js:"oncanplaythrough,rewrite=oncanplaythrough"
	Oncanplaythrough() (oncanplaythrough func(Event))

	// Oncanplaythrough prop
	// js:"setoncanplaythrough,rewrite=setoncanplaythrough"
	SetOncanplaythrough(oncanplaythrough func(Event))

	// Onchange prop
	// js:"onchange,rewrite=onchange"
	Onchange() (onchange func(Event))

	// Onchange prop
	// js:"setonchange,rewrite=setonchange"
	SetOnchange(onchange func(Event))

	// Onclick prop
	// js:"onclick,rewrite=onclick"
	Onclick() (onclick func(MouseEvent))

	// Onclick prop
	// js:"setonclick,rewrite=setonclick"
	SetOnclick(onclick func(MouseEvent))

	// Oncontextmenu prop
	// js:"oncontextmenu,rewrite=oncontextmenu"
	Oncontextmenu() (oncontextmenu func(*PointerEvent))

	// Oncontextmenu prop
	// js:"setoncontextmenu,rewrite=setoncontextmenu"
	SetOncontextmenu(oncontextmenu func(*PointerEvent))

	// Oncopy prop
	// js:"oncopy,rewrite=oncopy"
	Oncopy() (oncopy func(*ClipboardEvent))

	// Oncopy prop
	// js:"setoncopy,rewrite=setoncopy"
	SetOncopy(oncopy func(*ClipboardEvent))

	// Oncuechange prop
	// js:"oncuechange,rewrite=oncuechange"
	Oncuechange() (oncuechange func(Event))

	// Oncuechange prop
	// js:"setoncuechange,rewrite=setoncuechange"
	SetOncuechange(oncuechange func(Event))

	// Oncut prop
	// js:"oncut,rewrite=oncut"
	Oncut() (oncut func(*ClipboardEvent))

	// Oncut prop
	// js:"setoncut,rewrite=setoncut"
	SetOncut(oncut func(*ClipboardEvent))

	// Ondblclick prop
	// js:"ondblclick,rewrite=ondblclick"
	Ondblclick() (ondblclick func(MouseEvent))

	// Ondblclick prop
	// js:"setondblclick,rewrite=setondblclick"
	SetOndblclick(ondblclick func(MouseEvent))

	// Ondeactivate prop
	// js:"ondeactivate,rewrite=ondeactivate"
	Ondeactivate() (ondeactivate func(UIEvent))

	// Ondeactivate prop
	// js:"setondeactivate,rewrite=setondeactivate"
	SetOndeactivate(ondeactivate func(UIEvent))

	// Ondrag prop
	// js:"ondrag,rewrite=ondrag"
	Ondrag() (ondrag func(*DragEvent))

	// Ondrag prop
	// js:"setondrag,rewrite=setondrag"
	SetOndrag(ondrag func(*DragEvent))

	// Ondragend prop
	// js:"ondragend,rewrite=ondragend"
	Ondragend() (ondragend func(*DragEvent))

	// Ondragend prop
	// js:"setondragend,rewrite=setondragend"
	SetOndragend(ondragend func(*DragEvent))

	// Ondragenter prop
	// js:"ondragenter,rewrite=ondragenter"
	Ondragenter() (ondragenter func(*DragEvent))

	// Ondragenter prop
	// js:"setondragenter,rewrite=setondragenter"
	SetOndragenter(ondragenter func(*DragEvent))

	// Ondragleave prop
	// js:"ondragleave,rewrite=ondragleave"
	Ondragleave() (ondragleave func(*DragEvent))

	// Ondragleave prop
	// js:"setondragleave,rewrite=setondragleave"
	SetOndragleave(ondragleave func(*DragEvent))

	// Ondragover prop
	// js:"ondragover,rewrite=ondragover"
	Ondragover() (ondragover func(*DragEvent))

	// Ondragover prop
	// js:"setondragover,rewrite=setondragover"
	SetOndragover(ondragover func(*DragEvent))

	// Ondragstart prop
	// js:"ondragstart,rewrite=ondragstart"
	Ondragstart() (ondragstart func(*DragEvent))

	// Ondragstart prop
	// js:"setondragstart,rewrite=setondragstart"
	SetOndragstart(ondragstart func(*DragEvent))

	// Ondrop prop
	// js:"ondrop,rewrite=ondrop"
	Ondrop() (ondrop func(*DragEvent))

	// Ondrop prop
	// js:"setondrop,rewrite=setondrop"
	SetOndrop(ondrop func(*DragEvent))

	// Ondurationchange prop
	// js:"ondurationchange,rewrite=ondurationchange"
	Ondurationchange() (ondurationchange func(Event))

	// Ondurationchange prop
	// js:"setondurationchange,rewrite=setondurationchange"
	SetOndurationchange(ondurationchange func(Event))

	// Onemptied prop
	// js:"onemptied,rewrite=onemptied"
	Onemptied() (onemptied func(Event))

	// Onemptied prop
	// js:"setonemptied,rewrite=setonemptied"
	SetOnemptied(onemptied func(Event))

	// Onended prop
	// js:"onended,rewrite=onended"
	Onended() (onended func(Event))

	// Onended prop
	// js:"setonended,rewrite=setonended"
	SetOnended(onended func(Event))

	// Onerror prop
	// js:"onerror,rewrite=onerror"
	Onerror() (onerror func(Event))

	// Onerror prop
	// js:"setonerror,rewrite=setonerror"
	SetOnerror(onerror func(Event))

	// Onfocus prop
	// js:"onfocus,rewrite=onfocus"
	Onfocus() (onfocus func(*FocusEvent))

	// Onfocus prop
	// js:"setonfocus,rewrite=setonfocus"
	SetOnfocus(onfocus func(*FocusEvent))

	// Oninput prop
	// js:"oninput,rewrite=oninput"
	Oninput() (oninput func(Event))

	// Oninput prop
	// js:"setoninput,rewrite=setoninput"
	SetOninput(oninput func(Event))

	// Oninvalid prop
	// js:"oninvalid,rewrite=oninvalid"
	Oninvalid() (oninvalid func(Event))

	// Oninvalid prop
	// js:"setoninvalid,rewrite=setoninvalid"
	SetOninvalid(oninvalid func(Event))

	// Onkeydown prop
	// js:"onkeydown,rewrite=onkeydown"
	Onkeydown() (onkeydown func(*KeyboardEvent))

	// Onkeydown prop
	// js:"setonkeydown,rewrite=setonkeydown"
	SetOnkeydown(onkeydown func(*KeyboardEvent))

	// Onkeypress prop
	// js:"onkeypress,rewrite=onkeypress"
	Onkeypress() (onkeypress func(*KeyboardEvent))

	// Onkeypress prop
	// js:"setonkeypress,rewrite=setonkeypress"
	SetOnkeypress(onkeypress func(*KeyboardEvent))

	// Onkeyup prop
	// js:"onkeyup,rewrite=onkeyup"
	Onkeyup() (onkeyup func(*KeyboardEvent))

	// Onkeyup prop
	// js:"setonkeyup,rewrite=setonkeyup"
	SetOnkeyup(onkeyup func(*KeyboardEvent))

	// Onload prop
	// js:"onload,rewrite=onload"
	Onload() (onload func(Event))

	// Onload prop
	// js:"setonload,rewrite=setonload"
	SetOnload(onload func(Event))

	// Onloadeddata prop
	// js:"onloadeddata,rewrite=onloadeddata"
	Onloadeddata() (onloadeddata func(Event))

	// Onloadeddata prop
	// js:"setonloadeddata,rewrite=setonloadeddata"
	SetOnloadeddata(onloadeddata func(Event))

	// Onloadedmetadata prop
	// js:"onloadedmetadata,rewrite=onloadedmetadata"
	Onloadedmetadata() (onloadedmetadata func(Event))

	// Onloadedmetadata prop
	// js:"setonloadedmetadata,rewrite=setonloadedmetadata"
	SetOnloadedmetadata(onloadedmetadata func(Event))

	// Onloadstart prop
	// js:"onloadstart,rewrite=onloadstart"
	Onloadstart() (onloadstart func(Event))

	// Onloadstart prop
	// js:"setonloadstart,rewrite=setonloadstart"
	SetOnloadstart(onloadstart func(Event))

	// Onmousedown prop
	// js:"onmousedown,rewrite=onmousedown"
	Onmousedown() (onmousedown func(MouseEvent))

	// Onmousedown prop
	// js:"setonmousedown,rewrite=setonmousedown"
	SetOnmousedown(onmousedown func(MouseEvent))

	// Onmouseenter prop
	// js:"onmouseenter,rewrite=onmouseenter"
	Onmouseenter() (onmouseenter func(MouseEvent))

	// Onmouseenter prop
	// js:"setonmouseenter,rewrite=setonmouseenter"
	SetOnmouseenter(onmouseenter func(MouseEvent))

	// Onmouseleave prop
	// js:"onmouseleave,rewrite=onmouseleave"
	Onmouseleave() (onmouseleave func(MouseEvent))

	// Onmouseleave prop
	// js:"setonmouseleave,rewrite=setonmouseleave"
	SetOnmouseleave(onmouseleave func(MouseEvent))

	// Onmousemove prop
	// js:"onmousemove,rewrite=onmousemove"
	Onmousemove() (onmousemove func(MouseEvent))

	// Onmousemove prop
	// js:"setonmousemove,rewrite=setonmousemove"
	SetOnmousemove(onmousemove func(MouseEvent))

	// Onmouseout prop
	// js:"onmouseout,rewrite=onmouseout"
	Onmouseout() (onmouseout func(MouseEvent))

	// Onmouseout prop
	// js:"setonmouseout,rewrite=setonmouseout"
	SetOnmouseout(onmouseout func(MouseEvent))

	// Onmouseover prop
	// js:"onmouseover,rewrite=onmouseover"
	Onmouseover() (onmouseover func(MouseEvent))

	// Onmouseover prop
	// js:"setonmouseover,rewrite=setonmouseover"
	SetOnmouseover(onmouseover func(MouseEvent))

	// Onmouseup prop
	// js:"onmouseup,rewrite=onmouseup"
	Onmouseup() (onmouseup func(MouseEvent))

	// Onmouseup prop
	// js:"setonmouseup,rewrite=setonmouseup"
	SetOnmouseup(onmouseup func(MouseEvent))

	// Onmousewheel prop
	// js:"onmousewheel,rewrite=onmousewheel"
	Onmousewheel() (onmousewheel func(*WheelEvent))

	// Onmousewheel prop
	// js:"setonmousewheel,rewrite=setonmousewheel"
	SetOnmousewheel(onmousewheel func(*WheelEvent))

	// Onmscontentzoom prop
	// js:"onmscontentzoom,rewrite=onmscontentzoom"
	Onmscontentzoom() (onmscontentzoom func(UIEvent))

	// Onmscontentzoom prop
	// js:"setonmscontentzoom,rewrite=setonmscontentzoom"
	SetOnmscontentzoom(onmscontentzoom func(UIEvent))

	// Onmsmanipulationstatechanged prop
	// js:"onmsmanipulationstatechanged,rewrite=onmsmanipulationstatechanged"
	Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onmsmanipulationstatechanged prop
	// js:"setonmsmanipulationstatechanged,rewrite=setonmsmanipulationstatechanged"
	SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent))

	// Onpaste prop
	// js:"onpaste,rewrite=onpaste"
	Onpaste() (onpaste func(*ClipboardEvent))

	// Onpaste prop
	// js:"setonpaste,rewrite=setonpaste"
	SetOnpaste(onpaste func(*ClipboardEvent))

	// Onpause prop
	// js:"onpause,rewrite=onpause"
	Onpause() (onpause func(Event))

	// Onpause prop
	// js:"setonpause,rewrite=setonpause"
	SetOnpause(onpause func(Event))

	// Onplay prop
	// js:"onplay,rewrite=onplay"
	Onplay() (onplay func(Event))

	// Onplay prop
	// js:"setonplay,rewrite=setonplay"
	SetOnplay(onplay func(Event))

	// Onplaying prop
	// js:"onplaying,rewrite=onplaying"
	Onplaying() (onplaying func(Event))

	// Onplaying prop
	// js:"setonplaying,rewrite=setonplaying"
	SetOnplaying(onplaying func(Event))

	// Onprogress prop
	// js:"onprogress,rewrite=onprogress"
	Onprogress() (onprogress func(Event))

	// Onprogress prop
	// js:"setonprogress,rewrite=setonprogress"
	SetOnprogress(onprogress func(Event))

	// Onratechange prop
	// js:"onratechange,rewrite=onratechange"
	Onratechange() (onratechange func(Event))

	// Onratechange prop
	// js:"setonratechange,rewrite=setonratechange"
	SetOnratechange(onratechange func(Event))

	// Onreset prop
	// js:"onreset,rewrite=onreset"
	Onreset() (onreset func(Event))

	// Onreset prop
	// js:"setonreset,rewrite=setonreset"
	SetOnreset(onreset func(Event))

	// Onscroll prop
	// js:"onscroll,rewrite=onscroll"
	Onscroll() (onscroll func(UIEvent))

	// Onscroll prop
	// js:"setonscroll,rewrite=setonscroll"
	SetOnscroll(onscroll func(UIEvent))

	// Onseeked prop
	// js:"onseeked,rewrite=onseeked"
	Onseeked() (onseeked func(Event))

	// Onseeked prop
	// js:"setonseeked,rewrite=setonseeked"
	SetOnseeked(onseeked func(Event))

	// Onseeking prop
	// js:"onseeking,rewrite=onseeking"
	Onseeking() (onseeking func(Event))

	// Onseeking prop
	// js:"setonseeking,rewrite=setonseeking"
	SetOnseeking(onseeking func(Event))

	// Onselect prop
	// js:"onselect,rewrite=onselect"
	Onselect() (onselect func(UIEvent))

	// Onselect prop
	// js:"setonselect,rewrite=setonselect"
	SetOnselect(onselect func(UIEvent))

	// Onselectstart prop
	// js:"onselectstart,rewrite=onselectstart"
	Onselectstart() (onselectstart func(Event))

	// Onselectstart prop
	// js:"setonselectstart,rewrite=setonselectstart"
	SetOnselectstart(onselectstart func(Event))

	// Onstalled prop
	// js:"onstalled,rewrite=onstalled"
	Onstalled() (onstalled func(Event))

	// Onstalled prop
	// js:"setonstalled,rewrite=setonstalled"
	SetOnstalled(onstalled func(Event))

	// Onsubmit prop
	// js:"onsubmit,rewrite=onsubmit"
	Onsubmit() (onsubmit func(Event))

	// Onsubmit prop
	// js:"setonsubmit,rewrite=setonsubmit"
	SetOnsubmit(onsubmit func(Event))

	// Onsuspend prop
	// js:"onsuspend,rewrite=onsuspend"
	Onsuspend() (onsuspend func(Event))

	// Onsuspend prop
	// js:"setonsuspend,rewrite=setonsuspend"
	SetOnsuspend(onsuspend func(Event))

	// Ontimeupdate prop
	// js:"ontimeupdate,rewrite=ontimeupdate"
	Ontimeupdate() (ontimeupdate func(Event))

	// Ontimeupdate prop
	// js:"setontimeupdate,rewrite=setontimeupdate"
	SetOntimeupdate(ontimeupdate func(Event))

	// Onvolumechange prop
	// js:"onvolumechange,rewrite=onvolumechange"
	Onvolumechange() (onvolumechange func(Event))

	// Onvolumechange prop
	// js:"setonvolumechange,rewrite=setonvolumechange"
	SetOnvolumechange(onvolumechange func(Event))

	// Onwaiting prop
	// js:"onwaiting,rewrite=onwaiting"
	Onwaiting() (onwaiting func(Event))

	// Onwaiting prop
	// js:"setonwaiting,rewrite=setonwaiting"
	SetOnwaiting(onwaiting func(Event))

	// OuterText prop
	// js:"outerText,rewrite=outertext"
	OuterText() (outerText string)

	// OuterText prop
	// js:"setouterText,rewrite=setoutertext"
	SetOuterText(outerText string)

	// Spellcheck prop
	// js:"spellcheck,rewrite=spellcheck"
	Spellcheck() (spellcheck bool)

	// Spellcheck prop
	// js:"setspellcheck,rewrite=setspellcheck"
	SetSpellcheck(spellcheck bool)

	// Style prop
	// js:"style,rewrite=style"
	Style() (style *CSSStyleDeclaration)

	// TabIndex prop
	// js:"tabIndex,rewrite=tabindex"
	TabIndex() (tabIndex int8)

	// TabIndex prop
	// js:"settabIndex,rewrite=settabindex"
	SetTabIndex(tabIndex int8)

	// Title prop
	// js:"title,rewrite=title"
	Title() (title string)

	// Title prop
	// js:"settitle,rewrite=settitle"
	SetTitle(title string)
}

// blur fn
func blur() {
	js.Rewrite("$<.blur()")
}

// click fn
func click() {
	js.Rewrite("$<.click()")
}

// dragdrop fn
func dragdrop() (b bool) {
	js.Rewrite("$<.dragDrop()")
	return b
}

// // focus fn
// func focus() {
// 	js.Rewrite("$<.focus()")
// }

// // getelementsbyclassname fn
// func getelementsbyclassname(classNames string) (n *NodeList) {
// 	js.Rewrite("$<.getElementsByClassName($1)", classNames)
// 	return n
// }

// insertadjacentelement fn
func insertadjacentelement(position string, insertedElement Element) (e Element) {
	js.Rewrite("$<.insertAdjacentElement($1, $2)", position, insertedElement)
	return e
}

// insertadjacenthtml fn
func insertadjacenthtml(where string, html string) {
	js.Rewrite("$<.insertAdjacentHTML($1, $2)", where, html)
}

// insertadjacenttext fn
func insertadjacenttext(where string, text string) {
	js.Rewrite("$<.insertAdjacentText($1, $2)", where, text)
}

// msgetinputcontext fn
func msgetinputcontext() (m *MSInputMethodContext) {
	js.Rewrite("$<.msGetInputContext()")
	return m
}

// scrollintoview fn
func scrollintoview(top *bool) {
	js.Rewrite("$<.scrollIntoView($1)", top)
}

// accesskey prop
func accesskey() (accessKey string) {
	js.Rewrite("$<.accessKey")
	return accessKey
}

// setaccesskey prop
func setaccesskey(accessKey string) {
	js.Rewrite("$<.accessKey = accessKey")
}

// children prop
func children() (children HTMLCollection) {
	js.Rewrite("$<.children")
	return children
}

// contenteditable prop
func contenteditable() (contentEditable string) {
	js.Rewrite("$<.contentEditable")
	return contentEditable
}

// setcontenteditable prop
func setcontenteditable(contentEditable string) {
	js.Rewrite("$<.contentEditable = contentEditable")
}

// dataset prop
// func dataset() (dataset *domstringmap.DOMStringMap) {
// 	js.Rewrite("$<.dataset")
// 	return dataset
// }

// dir prop
// func dir() (dir string) {
// 	js.Rewrite("$<.dir")
// 	return dir
// }

// // setdir prop
// func setdir(dir string) {
// 	js.Rewrite("$<.dir = dir")
// }

// draggable prop
func draggable() (draggable bool) {
	js.Rewrite("$<.draggable")
	return draggable
}

// setdraggable prop
func setdraggable(draggable bool) {
	js.Rewrite("$<.draggable = draggable")
}

// hidden prop
// func hidden() (hidden bool) {
// 	js.Rewrite("$<.hidden")
// 	return hidden
// }

// sethidden prop
func sethidden(hidden bool) {
	js.Rewrite("$<.hidden = hidden")
}

// hidefocus prop
func hidefocus() (hideFocus bool) {
	js.Rewrite("$<.hideFocus")
	return hideFocus
}

// sethidefocus prop
func sethidefocus(hideFocus bool) {
	js.Rewrite("$<.hideFocus = hideFocus")
}

// innertext prop
func innertext() (innerText string) {
	js.Rewrite("$<.innerText")
	return innerText
}

// setinnertext prop
func setinnertext(innerText string) {
	js.Rewrite("$<.innerText = innerText")
}

// iscontenteditable prop
func iscontenteditable() (isContentEditable bool) {
	js.Rewrite("$<.isContentEditable")
	return isContentEditable
}

// lang prop
func lang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

// setlang prop
func setlang(lang string) {
	js.Rewrite("$<.lang = lang")
}

// offsetheight prop
func offsetheight() (offsetHeight int) {
	js.Rewrite("$<.offsetHeight")
	return offsetHeight
}

// offsetleft prop
func offsetleft() (offsetLeft int) {
	js.Rewrite("$<.offsetLeft")
	return offsetLeft
}

// offsetparent prop
func offsetparent() (offsetParent Element) {
	js.Rewrite("$<.offsetParent")
	return offsetParent
}

// offsettop prop
func offsettop() (offsetTop int) {
	js.Rewrite("$<.offsetTop")
	return offsetTop
}

// offsetwidth prop
func offsetwidth() (offsetWidth int) {
	js.Rewrite("$<.offsetWidth")
	return offsetWidth
}

// // onabort prop
// func onabort() (onabort func(Event)) {
// 	js.Rewrite("$<.onabort")
// 	return onabort
// }

// // setonabort prop
// func setonabort(onabort func(Event)) {
// 	js.Rewrite("$<.onabort = onabort")
// }

// // onactivate prop
// func onactivate() (onactivate func(UIEvent)) {
// 	js.Rewrite("$<.onactivate")
// 	return onactivate
// }

// // setonactivate prop
// func setonactivate(onactivate func(UIEvent)) {
// 	js.Rewrite("$<.onactivate = onactivate")
// }

// // onbeforeactivate prop
// func onbeforeactivate() (onbeforeactivate func(UIEvent)) {
// 	js.Rewrite("$<.onbeforeactivate")
// 	return onbeforeactivate
// }

// // setonbeforeactivate prop
// func setonbeforeactivate(onbeforeactivate func(UIEvent)) {
// 	js.Rewrite("$<.onbeforeactivate = onbeforeactivate")
// }

// onbeforecopy prop
func onbeforecopy() (onbeforecopy func(*ClipboardEvent)) {
	js.Rewrite("$<.onbeforecopy")
	return onbeforecopy
}

// setonbeforecopy prop
func setonbeforecopy(onbeforecopy func(*ClipboardEvent)) {
	js.Rewrite("$<.onbeforecopy = onbeforecopy")
}

// onbeforecut prop
func onbeforecut() (onbeforecut func(*ClipboardEvent)) {
	js.Rewrite("$<.onbeforecut")
	return onbeforecut
}

// setonbeforecut prop
func setonbeforecut(onbeforecut func(*ClipboardEvent)) {
	js.Rewrite("$<.onbeforecut = onbeforecut")
}

// // onbeforedeactivate prop
// func onbeforedeactivate() (onbeforedeactivate func(UIEvent)) {
// 	js.Rewrite("$<.onbeforedeactivate")
// 	return onbeforedeactivate
// }

// // setonbeforedeactivate prop
// func setonbeforedeactivate(onbeforedeactivate func(UIEvent)) {
// 	js.Rewrite("$<.onbeforedeactivate = onbeforedeactivate")
// }

// onbeforepaste prop
func onbeforepaste() (onbeforepaste func(*ClipboardEvent)) {
	js.Rewrite("$<.onbeforepaste")
	return onbeforepaste
}

// setonbeforepaste prop
func setonbeforepaste(onbeforepaste func(*ClipboardEvent)) {
	js.Rewrite("$<.onbeforepaste = onbeforepaste")
}

// // onblur prop
// func onblur() (onblur func(*FocusEvent)) {
// 	js.Rewrite("$<.onblur")
// 	return onblur
// }

// // setonblur prop
// func setonblur(onblur func(*FocusEvent)) {
// 	js.Rewrite("$<.onblur = onblur")
// }

// // oncanplay prop
// func oncanplay() (oncanplay func(Event)) {
// 	js.Rewrite("$<.oncanplay")
// 	return oncanplay
// }

// // setoncanplay prop
// func setoncanplay(oncanplay func(Event)) {
// 	js.Rewrite("$<.oncanplay = oncanplay")
// }

// // oncanplaythrough prop
// func oncanplaythrough() (oncanplaythrough func(Event)) {
// 	js.Rewrite("$<.oncanplaythrough")
// 	return oncanplaythrough
// }

// // setoncanplaythrough prop
// func setoncanplaythrough(oncanplaythrough func(Event)) {
// 	js.Rewrite("$<.oncanplaythrough = oncanplaythrough")
// }

// // onchange prop
// func onchange() (onchange func(Event)) {
// 	js.Rewrite("$<.onchange")
// 	return onchange
// }

// // setonchange prop
// func setonchange(onchange func(Event)) {
// 	js.Rewrite("$<.onchange = onchange")
// }

// // onclick prop
// func onclick() (onclick func(MouseEvent)) {
// 	js.Rewrite("$<.onclick")
// 	return onclick
// }

// // setonclick prop
// func setonclick(onclick func(MouseEvent)) {
// 	js.Rewrite("$<.onclick = onclick")
// }

// // oncontextmenu prop
// func oncontextmenu() (oncontextmenu func(*PointerEvent)) {
// 	js.Rewrite("$<.oncontextmenu")
// 	return oncontextmenu
// }

// // setoncontextmenu prop
// func setoncontextmenu(oncontextmenu func(*PointerEvent)) {
// 	js.Rewrite("$<.oncontextmenu = oncontextmenu")
// }

// oncopy prop
func oncopy() (oncopy func(*ClipboardEvent)) {
	js.Rewrite("$<.oncopy")
	return oncopy
}

// setoncopy prop
func setoncopy(oncopy func(*ClipboardEvent)) {
	js.Rewrite("$<.oncopy = oncopy")
}

// oncuechange prop
func oncuechange() (oncuechange func(Event)) {
	js.Rewrite("$<.oncuechange")
	return oncuechange
}

// setoncuechange prop
func setoncuechange(oncuechange func(Event)) {
	js.Rewrite("$<.oncuechange = oncuechange")
}

// oncut prop
func oncut() (oncut func(*ClipboardEvent)) {
	js.Rewrite("$<.oncut")
	return oncut
}

// setoncut prop
func setoncut(oncut func(*ClipboardEvent)) {
	js.Rewrite("$<.oncut = oncut")
}

// // ondblclick prop
// func ondblclick() (ondblclick func(MouseEvent)) {
// 	js.Rewrite("$<.ondblclick")
// 	return ondblclick
// }

// // setondblclick prop
// func setondblclick(ondblclick func(MouseEvent)) {
// 	js.Rewrite("$<.ondblclick = ondblclick")
// }

// // ondeactivate prop
// func ondeactivate() (ondeactivate func(UIEvent)) {
// 	js.Rewrite("$<.ondeactivate")
// 	return ondeactivate
// }

// // setondeactivate prop
// func setondeactivate(ondeactivate func(UIEvent)) {
// 	js.Rewrite("$<.ondeactivate = ondeactivate")
// }

// // ondrag prop
// func ondrag() (ondrag func(*DragEvent)) {
// 	js.Rewrite("$<.ondrag")
// 	return ondrag
// }

// // setondrag prop
// func setondrag(ondrag func(*DragEvent)) {
// 	js.Rewrite("$<.ondrag = ondrag")
// }

// // ondragend prop
// func ondragend() (ondragend func(*DragEvent)) {
// 	js.Rewrite("$<.ondragend")
// 	return ondragend
// }

// // setondragend prop
// func setondragend(ondragend func(*DragEvent)) {
// 	js.Rewrite("$<.ondragend = ondragend")
// }

// // ondragenter prop
// func ondragenter() (ondragenter func(*DragEvent)) {
// 	js.Rewrite("$<.ondragenter")
// 	return ondragenter
// }

// // setondragenter prop
// func setondragenter(ondragenter func(*DragEvent)) {
// 	js.Rewrite("$<.ondragenter = ondragenter")
// }

// // ondragleave prop
// func ondragleave() (ondragleave func(*DragEvent)) {
// 	js.Rewrite("$<.ondragleave")
// 	return ondragleave
// }

// // setondragleave prop
// func setondragleave(ondragleave func(*DragEvent)) {
// 	js.Rewrite("$<.ondragleave = ondragleave")
// }

// // ondragover prop
// func ondragover() (ondragover func(*DragEvent)) {
// 	js.Rewrite("$<.ondragover")
// 	return ondragover
// }

// // setondragover prop
// func setondragover(ondragover func(*DragEvent)) {
// 	js.Rewrite("$<.ondragover = ondragover")
// }

// // ondragstart prop
// func ondragstart() (ondragstart func(*DragEvent)) {
// 	js.Rewrite("$<.ondragstart")
// 	return ondragstart
// }

// // setondragstart prop
// func setondragstart(ondragstart func(*DragEvent)) {
// 	js.Rewrite("$<.ondragstart = ondragstart")
// }

// // ondrop prop
// func ondrop() (ondrop func(*DragEvent)) {
// 	js.Rewrite("$<.ondrop")
// 	return ondrop
// }

// // setondrop prop
// func setondrop(ondrop func(*DragEvent)) {
// 	js.Rewrite("$<.ondrop = ondrop")
// }

// // ondurationchange prop
// func ondurationchange() (ondurationchange func(Event)) {
// 	js.Rewrite("$<.ondurationchange")
// 	return ondurationchange
// }

// // setondurationchange prop
// func setondurationchange(ondurationchange func(Event)) {
// 	js.Rewrite("$<.ondurationchange = ondurationchange")
// }

// // onemptied prop
// func onemptied() (onemptied func(Event)) {
// 	js.Rewrite("$<.onemptied")
// 	return onemptied
// }

// // setonemptied prop
// func setonemptied(onemptied func(Event)) {
// 	js.Rewrite("$<.onemptied = onemptied")
// }

// // onended prop
// func onended() (onended func(Event)) {
// 	js.Rewrite("$<.onended")
// 	return onended
// }

// // setonended prop
// func setonended(onended func(Event)) {
// 	js.Rewrite("$<.onended = onended")
// }

// // onerror prop
// func onerror() (onerror func(Event)) {
// 	js.Rewrite("$<.onerror")
// 	return onerror
// }

// // setonerror prop
// func setonerror(onerror func(Event)) {
// 	js.Rewrite("$<.onerror = onerror")
// }

// // onfocus prop
// func onfocus() (onfocus func(*FocusEvent)) {
// 	js.Rewrite("$<.onfocus")
// 	return onfocus
// }

// // setonfocus prop
// func setonfocus(onfocus func(*FocusEvent)) {
// 	js.Rewrite("$<.onfocus = onfocus")
// }

// // oninput prop
// func oninput() (oninput func(Event)) {
// 	js.Rewrite("$<.oninput")
// 	return oninput
// }

// // setoninput prop
// func setoninput(oninput func(Event)) {
// 	js.Rewrite("$<.oninput = oninput")
// }

// // oninvalid prop
// func oninvalid() (oninvalid func(Event)) {
// 	js.Rewrite("$<.oninvalid")
// 	return oninvalid
// }

// // setoninvalid prop
// func setoninvalid(oninvalid func(Event)) {
// 	js.Rewrite("$<.oninvalid = oninvalid")
// }

// // onkeydown prop
// func onkeydown() (onkeydown func(*KeyboardEvent)) {
// 	js.Rewrite("$<.onkeydown")
// 	return onkeydown
// }

// // setonkeydown prop
// func setonkeydown(onkeydown func(*KeyboardEvent)) {
// 	js.Rewrite("$<.onkeydown = onkeydown")
// }

// // onkeypress prop
// func onkeypress() (onkeypress func(*KeyboardEvent)) {
// 	js.Rewrite("$<.onkeypress")
// 	return onkeypress
// }

// // setonkeypress prop
// func setonkeypress(onkeypress func(*KeyboardEvent)) {
// 	js.Rewrite("$<.onkeypress = onkeypress")
// }

// // onkeyup prop
// func onkeyup() (onkeyup func(*KeyboardEvent)) {
// 	js.Rewrite("$<.onkeyup")
// 	return onkeyup
// }

// // setonkeyup prop
// func setonkeyup(onkeyup func(*KeyboardEvent)) {
// 	js.Rewrite("$<.onkeyup = onkeyup")
// }

// // onload prop
// func onload() (onload func(Event)) {
// 	js.Rewrite("$<.onload")
// 	return onload
// }

// // setonload prop
// func setonload(onload func(Event)) {
// 	js.Rewrite("$<.onload = onload")
// }

// // onloadeddata prop
// func onloadeddata() (onloadeddata func(Event)) {
// 	js.Rewrite("$<.onloadeddata")
// 	return onloadeddata
// }

// // setonloadeddata prop
// func setonloadeddata(onloadeddata func(Event)) {
// 	js.Rewrite("$<.onloadeddata = onloadeddata")
// }

// // onloadedmetadata prop
// func onloadedmetadata() (onloadedmetadata func(Event)) {
// 	js.Rewrite("$<.onloadedmetadata")
// 	return onloadedmetadata
// }

// // setonloadedmetadata prop
// func setonloadedmetadata(onloadedmetadata func(Event)) {
// 	js.Rewrite("$<.onloadedmetadata = onloadedmetadata")
// }

// // onloadstart prop
// func onloadstart() (onloadstart func(Event)) {
// 	js.Rewrite("$<.onloadstart")
// 	return onloadstart
// }

// // setonloadstart prop
// func setonloadstart(onloadstart func(Event)) {
// 	js.Rewrite("$<.onloadstart = onloadstart")
// }

// // onmousedown prop
// func onmousedown() (onmousedown func(MouseEvent)) {
// 	js.Rewrite("$<.onmousedown")
// 	return onmousedown
// }

// // setonmousedown prop
// func setonmousedown(onmousedown func(MouseEvent)) {
// 	js.Rewrite("$<.onmousedown = onmousedown")
// }

// onmouseenter prop
func onmouseenter() (onmouseenter func(MouseEvent)) {
	js.Rewrite("$<.onmouseenter")
	return onmouseenter
}

// setonmouseenter prop
func setonmouseenter(onmouseenter func(MouseEvent)) {
	js.Rewrite("$<.onmouseenter = onmouseenter")
}

// onmouseleave prop
func onmouseleave() (onmouseleave func(MouseEvent)) {
	js.Rewrite("$<.onmouseleave")
	return onmouseleave
}

// setonmouseleave prop
func setonmouseleave(onmouseleave func(MouseEvent)) {
	js.Rewrite("$<.onmouseleave = onmouseleave")
}

// // onmousemove prop
// func onmousemove() (onmousemove func(MouseEvent)) {
// 	js.Rewrite("$<.onmousemove")
// 	return onmousemove
// }

// // setonmousemove prop
// func setonmousemove(onmousemove func(MouseEvent)) {
// 	js.Rewrite("$<.onmousemove = onmousemove")
// }

// // onmouseout prop
// func onmouseout() (onmouseout func(MouseEvent)) {
// 	js.Rewrite("$<.onmouseout")
// 	return onmouseout
// }

// // setonmouseout prop
// func setonmouseout(onmouseout func(MouseEvent)) {
// 	js.Rewrite("$<.onmouseout = onmouseout")
// }

// // onmouseover prop
// func onmouseover() (onmouseover func(MouseEvent)) {
// 	js.Rewrite("$<.onmouseover")
// 	return onmouseover
// }

// // setonmouseover prop
// func setonmouseover(onmouseover func(MouseEvent)) {
// 	js.Rewrite("$<.onmouseover = onmouseover")
// }

// // onmouseup prop
// func onmouseup() (onmouseup func(MouseEvent)) {
// 	js.Rewrite("$<.onmouseup")
// 	return onmouseup
// }

// // setonmouseup prop
// func setonmouseup(onmouseup func(MouseEvent)) {
// 	js.Rewrite("$<.onmouseup = onmouseup")
// }

// // onmousewheel prop
// func onmousewheel() (onmousewheel func(*WheelEvent)) {
// 	js.Rewrite("$<.onmousewheel")
// 	return onmousewheel
// }

// // setonmousewheel prop
// func setonmousewheel(onmousewheel func(*WheelEvent)) {
// 	js.Rewrite("$<.onmousewheel = onmousewheel")
// }

// // onmscontentzoom prop
// func onmscontentzoom() (onmscontentzoom func(UIEvent)) {
// 	js.Rewrite("$<.onmscontentzoom")
// 	return onmscontentzoom
// }

// // setonmscontentzoom prop
// func setonmscontentzoom(onmscontentzoom func(UIEvent)) {
// 	js.Rewrite("$<.onmscontentzoom = onmscontentzoom")
// }

// // onmsmanipulationstatechanged prop
// func onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*MSManipulationEvent)) {
// 	js.Rewrite("$<.onmsmanipulationstatechanged")
// 	return onmsmanipulationstatechanged
// }

// // setonmsmanipulationstatechanged prop
// func setonmsmanipulationstatechanged(onmsmanipulationstatechanged func(*MSManipulationEvent)) {
// 	js.Rewrite("$<.onmsmanipulationstatechanged = onmsmanipulationstatechanged")
// }

// onpaste prop
func onpaste() (onpaste func(*ClipboardEvent)) {
	js.Rewrite("$<.onpaste")
	return onpaste
}

// setonpaste prop
func setonpaste(onpaste func(*ClipboardEvent)) {
	js.Rewrite("$<.onpaste = onpaste")
}

// // onpause prop
// func onpause() (onpause func(Event)) {
// 	js.Rewrite("$<.onpause")
// 	return onpause
// }

// // setonpause prop
// func setonpause(onpause func(Event)) {
// 	js.Rewrite("$<.onpause = onpause")
// }

// // onplay prop
// func onplay() (onplay func(Event)) {
// 	js.Rewrite("$<.onplay")
// 	return onplay
// }

// // setonplay prop
// func setonplay(onplay func(Event)) {
// 	js.Rewrite("$<.onplay = onplay")
// }

// // onplaying prop
// func onplaying() (onplaying func(Event)) {
// 	js.Rewrite("$<.onplaying")
// 	return onplaying
// }

// // setonplaying prop
// func setonplaying(onplaying func(Event)) {
// 	js.Rewrite("$<.onplaying = onplaying")
// }

// // onprogress prop
// func onprogress() (onprogress func(Event)) {
// 	js.Rewrite("$<.onprogress")
// 	return onprogress
// }

// // setonprogress prop
// func setonprogress(onprogress func(Event)) {
// 	js.Rewrite("$<.onprogress = onprogress")
// }

// // onratechange prop
// func onratechange() (onratechange func(Event)) {
// 	js.Rewrite("$<.onratechange")
// 	return onratechange
// }

// // setonratechange prop
// func setonratechange(onratechange func(Event)) {
// 	js.Rewrite("$<.onratechange = onratechange")
// }

// // onreset prop
// func onreset() (onreset func(Event)) {
// 	js.Rewrite("$<.onreset")
// 	return onreset
// }

// // setonreset prop
// func setonreset(onreset func(Event)) {
// 	js.Rewrite("$<.onreset = onreset")
// }

// // onscroll prop
// func onscroll() (onscroll func(UIEvent)) {
// 	js.Rewrite("$<.onscroll")
// 	return onscroll
// }

// // setonscroll prop
// func setonscroll(onscroll func(UIEvent)) {
// 	js.Rewrite("$<.onscroll = onscroll")
// }

// // onseeked prop
// func onseeked() (onseeked func(Event)) {
// 	js.Rewrite("$<.onseeked")
// 	return onseeked
// }

// // setonseeked prop
// func setonseeked(onseeked func(Event)) {
// 	js.Rewrite("$<.onseeked = onseeked")
// }

// // onseeking prop
// func onseeking() (onseeking func(Event)) {
// 	js.Rewrite("$<.onseeking")
// 	return onseeking
// }

// // setonseeking prop
// func setonseeking(onseeking func(Event)) {
// 	js.Rewrite("$<.onseeking = onseeking")
// }

// // onselect prop
// func onselect() (onselect func(UIEvent)) {
// 	js.Rewrite("$<.onselect")
// 	return onselect
// }

// // setonselect prop
// func setonselect(onselect func(UIEvent)) {
// 	js.Rewrite("$<.onselect = onselect")
// }

// // onselectstart prop
// func onselectstart() (onselectstart func(Event)) {
// 	js.Rewrite("$<.onselectstart")
// 	return onselectstart
// }

// // setonselectstart prop
// func setonselectstart(onselectstart func(Event)) {
// 	js.Rewrite("$<.onselectstart = onselectstart")
// }

// // onstalled prop
// func onstalled() (onstalled func(Event)) {
// 	js.Rewrite("$<.onstalled")
// 	return onstalled
// }

// // setonstalled prop
// func setonstalled(onstalled func(Event)) {
// 	js.Rewrite("$<.onstalled = onstalled")
// }

// // onsubmit prop
// func onsubmit() (onsubmit func(Event)) {
// 	js.Rewrite("$<.onsubmit")
// 	return onsubmit
// }

// // setonsubmit prop
// func setonsubmit(onsubmit func(Event)) {
// 	js.Rewrite("$<.onsubmit = onsubmit")
// }

// // onsuspend prop
// func onsuspend() (onsuspend func(Event)) {
// 	js.Rewrite("$<.onsuspend")
// 	return onsuspend
// }

// // setonsuspend prop
// func setonsuspend(onsuspend func(Event)) {
// 	js.Rewrite("$<.onsuspend = onsuspend")
// }

// // ontimeupdate prop
// func ontimeupdate() (ontimeupdate func(Event)) {
// 	js.Rewrite("$<.ontimeupdate")
// 	return ontimeupdate
// }

// // setontimeupdate prop
// func setontimeupdate(ontimeupdate func(Event)) {
// 	js.Rewrite("$<.ontimeupdate = ontimeupdate")
// }

// // onvolumechange prop
// func onvolumechange() (onvolumechange func(Event)) {
// 	js.Rewrite("$<.onvolumechange")
// 	return onvolumechange
// }

// // setonvolumechange prop
// func setonvolumechange(onvolumechange func(Event)) {
// 	js.Rewrite("$<.onvolumechange = onvolumechange")
// }

// // onwaiting prop
// func onwaiting() (onwaiting func(Event)) {
// 	js.Rewrite("$<.onwaiting")
// 	return onwaiting
// }

// // setonwaiting prop
// func setonwaiting(onwaiting func(Event)) {
// 	js.Rewrite("$<.onwaiting = onwaiting")
// }

// outertext prop
func outertext() (outerText string) {
	js.Rewrite("$<.outerText")
	return outerText
}

// setoutertext prop
func setoutertext(outerText string) {
	js.Rewrite("$<.outerText = outerText")
}

// spellcheck prop
func spellcheck() (spellcheck bool) {
	js.Rewrite("$<.spellcheck")
	return spellcheck
}

// setspellcheck prop
func setspellcheck(spellcheck bool) {
	js.Rewrite("$<.spellcheck = spellcheck")
}

// style prop
func style() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

// tabindex prop
func tabindex() (tabIndex int8) {
	js.Rewrite("$<.tabIndex")
	return tabIndex
}

// settabindex prop
func settabindex(tabIndex int8) {
	js.Rewrite("$<.tabIndex = tabIndex")
}

// title prop
// func title() (title string) {
// 	js.Rewrite("$<.title")
// 	return title
// }

// // settitle prop
// func settitle(title string) {
// 	js.Rewrite("$<.title = title")
// }
