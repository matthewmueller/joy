package htmlelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/clipboardevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/internal/gendom/dom/domstringmap"
	"github.com/matthewmueller/golly/internal/gendom/dom/dragevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/focusevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/keyboardevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mouseevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/msinputmethodcontext"
	"github.com/matthewmueller/golly/internal/gendom/dom/msmanipulationevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/pointerevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/wheelevent"
	"github.com/matthewmueller/golly/js"
)

type HTMLElement struct {
	*element.Element
}

func (*HTMLElement) Blur() {
	js.Rewrite("$<.blur()")
}

func (*HTMLElement) Click() {
	js.Rewrite("$<.click()")
}

func (*HTMLElement) DragDrop() (b bool) {
	js.Rewrite("$<.dragDrop()")
	return b
}

func (*HTMLElement) Focus() {
	js.Rewrite("$<.focus()")
}

func (*HTMLElement) GetElementsByClassName(classNames string) (n *nodelist.NodeList) {
	js.Rewrite("$<.getElementsByClassName($1)", classNames)
	return n
}

func (*HTMLElement) InsertAdjacentElement(position string, insertedElement *element.Element) (e *element.Element) {
	js.Rewrite("$<.insertAdjacentElement($1, $2)", position, insertedElement)
	return e
}

func (*HTMLElement) InsertAdjacentHTML(where string, html string) {
	js.Rewrite("$<.insertAdjacentHTML($1, $2)", where, html)
}

func (*HTMLElement) InsertAdjacentText(where string, text string) {
	js.Rewrite("$<.insertAdjacentText($1, $2)", where, text)
}

func (*HTMLElement) MsGetInputContext() (m *msinputmethodcontext.MSInputMethodContext) {
	js.Rewrite("$<.msGetInputContext()")
	return m
}

func (*HTMLElement) ScrollIntoView(top bool) {
	js.Rewrite("$<.scrollIntoView($1)", top)
}

func (*HTMLElement) GetAccessKey() (accessKey string) {
	js.Rewrite("$<.accessKey")
	return accessKey
}

func (*HTMLElement) SetAccessKey(accessKey string) {
	js.Rewrite("$<.accessKey = $1", accessKey)
}

func (*HTMLElement) GetChildren() (children *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.children")
	return children
}

func (*HTMLElement) GetContentEditable() (contentEditable string) {
	js.Rewrite("$<.contentEditable")
	return contentEditable
}

func (*HTMLElement) SetContentEditable(contentEditable string) {
	js.Rewrite("$<.contentEditable = $1", contentEditable)
}

func (*HTMLElement) GetDataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$<.dataset")
	return dataset
}

func (*HTMLElement) GetDir() (dir string) {
	js.Rewrite("$<.dir")
	return dir
}

func (*HTMLElement) SetDir(dir string) {
	js.Rewrite("$<.dir = $1", dir)
}

func (*HTMLElement) GetDraggable() (draggable bool) {
	js.Rewrite("$<.draggable")
	return draggable
}

func (*HTMLElement) SetDraggable(draggable bool) {
	js.Rewrite("$<.draggable = $1", draggable)
}

func (*HTMLElement) GetHidden() (hidden bool) {
	js.Rewrite("$<.hidden")
	return hidden
}

func (*HTMLElement) SetHidden(hidden bool) {
	js.Rewrite("$<.hidden = $1", hidden)
}

func (*HTMLElement) GetHideFocus() (hideFocus bool) {
	js.Rewrite("$<.hideFocus")
	return hideFocus
}

func (*HTMLElement) SetHideFocus(hideFocus bool) {
	js.Rewrite("$<.hideFocus = $1", hideFocus)
}

func (*HTMLElement) GetInnerText() (innerText string) {
	js.Rewrite("$<.innerText")
	return innerText
}

func (*HTMLElement) SetInnerText(innerText string) {
	js.Rewrite("$<.innerText = $1", innerText)
}

func (*HTMLElement) GetIsContentEditable() (isContentEditable bool) {
	js.Rewrite("$<.isContentEditable")
	return isContentEditable
}

func (*HTMLElement) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*HTMLElement) SetLang(lang string) {
	js.Rewrite("$<.lang = $1", lang)
}

func (*HTMLElement) GetOffsetHeight() (offsetHeight int) {
	js.Rewrite("$<.offsetHeight")
	return offsetHeight
}

func (*HTMLElement) GetOffsetLeft() (offsetLeft int) {
	js.Rewrite("$<.offsetLeft")
	return offsetLeft
}

func (*HTMLElement) GetOffsetParent() (offsetParent *element.Element) {
	js.Rewrite("$<.offsetParent")
	return offsetParent
}

func (*HTMLElement) GetOffsetTop() (offsetTop int) {
	js.Rewrite("$<.offsetTop")
	return offsetTop
}

func (*HTMLElement) GetOffsetWidth() (offsetWidth int) {
	js.Rewrite("$<.offsetWidth")
	return offsetWidth
}

func (*HTMLElement) GetOnabort() (e *event.Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*HTMLElement) SetOnabort(e *event.Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*HTMLElement) GetOnactivate() (activate *uievent.UIEvent) {
	js.Rewrite("$<.onactivate")
	return activate
}

func (*HTMLElement) SetOnactivate(activate *uievent.UIEvent) {
	js.Rewrite("$<.onactivate = $1", activate)
}

func (*HTMLElement) GetOnbeforeactivate() (beforeactivate *uievent.UIEvent) {
	js.Rewrite("$<.onbeforeactivate")
	return beforeactivate
}

func (*HTMLElement) SetOnbeforeactivate(beforeactivate *uievent.UIEvent) {
	js.Rewrite("$<.onbeforeactivate = $1", beforeactivate)
}

func (*HTMLElement) GetOnbeforecopy() (beforecopy *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onbeforecopy")
	return beforecopy
}

func (*HTMLElement) SetOnbeforecopy(beforecopy *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onbeforecopy = $1", beforecopy)
}

func (*HTMLElement) GetOnbeforecut() (beforecut *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onbeforecut")
	return beforecut
}

func (*HTMLElement) SetOnbeforecut(beforecut *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onbeforecut = $1", beforecut)
}

func (*HTMLElement) GetOnbeforedeactivate() (beforedeactivate *uievent.UIEvent) {
	js.Rewrite("$<.onbeforedeactivate")
	return beforedeactivate
}

func (*HTMLElement) SetOnbeforedeactivate(beforedeactivate *uievent.UIEvent) {
	js.Rewrite("$<.onbeforedeactivate = $1", beforedeactivate)
}

func (*HTMLElement) GetOnbeforepaste() (beforepaste *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onbeforepaste")
	return beforepaste
}

func (*HTMLElement) SetOnbeforepaste(beforepaste *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onbeforepaste = $1", beforepaste)
}

func (*HTMLElement) GetOnblur() (blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*HTMLElement) SetOnblur(blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*HTMLElement) GetOncanplay() (e *event.Event) {
	js.Rewrite("$<.oncanplay")
	return e
}

func (*HTMLElement) SetOncanplay(e *event.Event) {
	js.Rewrite("$<.oncanplay = $1", e)
}

func (*HTMLElement) GetOncanplaythrough() (e *event.Event) {
	js.Rewrite("$<.oncanplaythrough")
	return e
}

func (*HTMLElement) SetOncanplaythrough(e *event.Event) {
	js.Rewrite("$<.oncanplaythrough = $1", e)
}

func (*HTMLElement) GetOnchange() (e *event.Event) {
	js.Rewrite("$<.onchange")
	return e
}

func (*HTMLElement) SetOnchange(e *event.Event) {
	js.Rewrite("$<.onchange = $1", e)
}

func (*HTMLElement) GetOnclick() (click *mouseevent.MouseEvent) {
	js.Rewrite("$<.onclick")
	return click
}

func (*HTMLElement) SetOnclick(click *mouseevent.MouseEvent) {
	js.Rewrite("$<.onclick = $1", click)
}

func (*HTMLElement) GetOncontextmenu() (contextmenu *pointerevent.PointerEvent) {
	js.Rewrite("$<.oncontextmenu")
	return contextmenu
}

func (*HTMLElement) SetOncontextmenu(contextmenu *pointerevent.PointerEvent) {
	js.Rewrite("$<.oncontextmenu = $1", contextmenu)
}

func (*HTMLElement) GetOncopy() (cpy *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.oncopy")
	return cpy
}

func (*HTMLElement) SetOncopy(cpy *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.oncopy = $1", cpy)
}

func (*HTMLElement) GetOncuechange() (e *event.Event) {
	js.Rewrite("$<.oncuechange")
	return e
}

func (*HTMLElement) SetOncuechange(e *event.Event) {
	js.Rewrite("$<.oncuechange = $1", e)
}

func (*HTMLElement) GetOncut() (cut *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.oncut")
	return cut
}

func (*HTMLElement) SetOncut(cut *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.oncut = $1", cut)
}

func (*HTMLElement) GetOndblclick() (dblclick *mouseevent.MouseEvent) {
	js.Rewrite("$<.ondblclick")
	return dblclick
}

func (*HTMLElement) SetOndblclick(dblclick *mouseevent.MouseEvent) {
	js.Rewrite("$<.ondblclick = $1", dblclick)
}

func (*HTMLElement) GetOndeactivate() (deactivate *uievent.UIEvent) {
	js.Rewrite("$<.ondeactivate")
	return deactivate
}

func (*HTMLElement) SetOndeactivate(deactivate *uievent.UIEvent) {
	js.Rewrite("$<.ondeactivate = $1", deactivate)
}

func (*HTMLElement) GetOndrag() (drag *dragevent.DragEvent) {
	js.Rewrite("$<.ondrag")
	return drag
}

func (*HTMLElement) SetOndrag(drag *dragevent.DragEvent) {
	js.Rewrite("$<.ondrag = $1", drag)
}

func (*HTMLElement) GetOndragend() (dragend *dragevent.DragEvent) {
	js.Rewrite("$<.ondragend")
	return dragend
}

func (*HTMLElement) SetOndragend(dragend *dragevent.DragEvent) {
	js.Rewrite("$<.ondragend = $1", dragend)
}

func (*HTMLElement) GetOndragenter() (dragenter *dragevent.DragEvent) {
	js.Rewrite("$<.ondragenter")
	return dragenter
}

func (*HTMLElement) SetOndragenter(dragenter *dragevent.DragEvent) {
	js.Rewrite("$<.ondragenter = $1", dragenter)
}

func (*HTMLElement) GetOndragleave() (dragleave *dragevent.DragEvent) {
	js.Rewrite("$<.ondragleave")
	return dragleave
}

func (*HTMLElement) SetOndragleave(dragleave *dragevent.DragEvent) {
	js.Rewrite("$<.ondragleave = $1", dragleave)
}

func (*HTMLElement) GetOndragover() (dragover *dragevent.DragEvent) {
	js.Rewrite("$<.ondragover")
	return dragover
}

func (*HTMLElement) SetOndragover(dragover *dragevent.DragEvent) {
	js.Rewrite("$<.ondragover = $1", dragover)
}

func (*HTMLElement) GetOndragstart() (dragstart *dragevent.DragEvent) {
	js.Rewrite("$<.ondragstart")
	return dragstart
}

func (*HTMLElement) SetOndragstart(dragstart *dragevent.DragEvent) {
	js.Rewrite("$<.ondragstart = $1", dragstart)
}

func (*HTMLElement) GetOndrop() (drop *dragevent.DragEvent) {
	js.Rewrite("$<.ondrop")
	return drop
}

func (*HTMLElement) SetOndrop(drop *dragevent.DragEvent) {
	js.Rewrite("$<.ondrop = $1", drop)
}

func (*HTMLElement) GetOndurationchange() (e *event.Event) {
	js.Rewrite("$<.ondurationchange")
	return e
}

func (*HTMLElement) SetOndurationchange(e *event.Event) {
	js.Rewrite("$<.ondurationchange = $1", e)
}

func (*HTMLElement) GetOnemptied() (e *event.Event) {
	js.Rewrite("$<.onemptied")
	return e
}

func (*HTMLElement) SetOnemptied(e *event.Event) {
	js.Rewrite("$<.onemptied = $1", e)
}

func (*HTMLElement) GetOnended() (e *event.Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*HTMLElement) SetOnended(e *event.Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*HTMLElement) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*HTMLElement) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*HTMLElement) GetOnfocus() (focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*HTMLElement) SetOnfocus(focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*HTMLElement) GetOninput() (e *event.Event) {
	js.Rewrite("$<.oninput")
	return e
}

func (*HTMLElement) SetOninput(e *event.Event) {
	js.Rewrite("$<.oninput = $1", e)
}

func (*HTMLElement) GetOninvalid() (e *event.Event) {
	js.Rewrite("$<.oninvalid")
	return e
}

func (*HTMLElement) SetOninvalid(e *event.Event) {
	js.Rewrite("$<.oninvalid = $1", e)
}

func (*HTMLElement) GetOnkeydown() (keydown *keyboardevent.KeyboardEvent) {
	js.Rewrite("$<.onkeydown")
	return keydown
}

func (*HTMLElement) SetOnkeydown(keydown *keyboardevent.KeyboardEvent) {
	js.Rewrite("$<.onkeydown = $1", keydown)
}

func (*HTMLElement) GetOnkeypress() (keypress *keyboardevent.KeyboardEvent) {
	js.Rewrite("$<.onkeypress")
	return keypress
}

func (*HTMLElement) SetOnkeypress(keypress *keyboardevent.KeyboardEvent) {
	js.Rewrite("$<.onkeypress = $1", keypress)
}

func (*HTMLElement) GetOnkeyup() (keyup *keyboardevent.KeyboardEvent) {
	js.Rewrite("$<.onkeyup")
	return keyup
}

func (*HTMLElement) SetOnkeyup(keyup *keyboardevent.KeyboardEvent) {
	js.Rewrite("$<.onkeyup = $1", keyup)
}

func (*HTMLElement) GetOnload() (e *event.Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*HTMLElement) SetOnload(e *event.Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*HTMLElement) GetOnloadeddata() (e *event.Event) {
	js.Rewrite("$<.onloadeddata")
	return e
}

func (*HTMLElement) SetOnloadeddata(e *event.Event) {
	js.Rewrite("$<.onloadeddata = $1", e)
}

func (*HTMLElement) GetOnloadedmetadata() (e *event.Event) {
	js.Rewrite("$<.onloadedmetadata")
	return e
}

func (*HTMLElement) SetOnloadedmetadata(e *event.Event) {
	js.Rewrite("$<.onloadedmetadata = $1", e)
}

func (*HTMLElement) GetOnloadstart() (e *event.Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*HTMLElement) SetOnloadstart(e *event.Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*HTMLElement) GetOnmousedown() (mousedown *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousedown")
	return mousedown
}

func (*HTMLElement) SetOnmousedown(mousedown *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousedown = $1", mousedown)
}

func (*HTMLElement) GetOnmouseenter() (mouseenter *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseenter")
	return mouseenter
}

func (*HTMLElement) SetOnmouseenter(mouseenter *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseenter = $1", mouseenter)
}

func (*HTMLElement) GetOnmouseleave() (mouseleave *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseleave")
	return mouseleave
}

func (*HTMLElement) SetOnmouseleave(mouseleave *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseleave = $1", mouseleave)
}

func (*HTMLElement) GetOnmousemove() (mousemove *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousemove")
	return mousemove
}

func (*HTMLElement) SetOnmousemove(mousemove *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmousemove = $1", mousemove)
}

func (*HTMLElement) GetOnmouseout() (mouseout *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseout")
	return mouseout
}

func (*HTMLElement) SetOnmouseout(mouseout *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseout = $1", mouseout)
}

func (*HTMLElement) GetOnmouseover() (mouseover *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseover")
	return mouseover
}

func (*HTMLElement) SetOnmouseover(mouseover *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseover = $1", mouseover)
}

func (*HTMLElement) GetOnmouseup() (mouseup *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseup")
	return mouseup
}

func (*HTMLElement) SetOnmouseup(mouseup *mouseevent.MouseEvent) {
	js.Rewrite("$<.onmouseup = $1", mouseup)
}

func (*HTMLElement) GetOnmousewheel() (mousewheel *wheelevent.WheelEvent) {
	js.Rewrite("$<.onmousewheel")
	return mousewheel
}

func (*HTMLElement) SetOnmousewheel(mousewheel *wheelevent.WheelEvent) {
	js.Rewrite("$<.onmousewheel = $1", mousewheel)
}

func (*HTMLElement) GetOnmscontentzoom() (MSContentZoom *uievent.UIEvent) {
	js.Rewrite("$<.onmscontentzoom")
	return MSContentZoom
}

func (*HTMLElement) SetOnmscontentzoom(MSContentZoom *uievent.UIEvent) {
	js.Rewrite("$<.onmscontentzoom = $1", MSContentZoom)
}

func (*HTMLElement) GetOnmsmanipulationstatechanged() (MSManipulationStateChanged *msmanipulationevent.MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged")
	return MSManipulationStateChanged
}

func (*HTMLElement) SetOnmsmanipulationstatechanged(MSManipulationStateChanged *msmanipulationevent.MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged = $1", MSManipulationStateChanged)
}

func (*HTMLElement) GetOnpaste() (paste *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onpaste")
	return paste
}

func (*HTMLElement) SetOnpaste(paste *clipboardevent.ClipboardEvent) {
	js.Rewrite("$<.onpaste = $1", paste)
}

func (*HTMLElement) GetOnpause() (e *event.Event) {
	js.Rewrite("$<.onpause")
	return e
}

func (*HTMLElement) SetOnpause(e *event.Event) {
	js.Rewrite("$<.onpause = $1", e)
}

func (*HTMLElement) GetOnplay() (e *event.Event) {
	js.Rewrite("$<.onplay")
	return e
}

func (*HTMLElement) SetOnplay(e *event.Event) {
	js.Rewrite("$<.onplay = $1", e)
}

func (*HTMLElement) GetOnplaying() (e *event.Event) {
	js.Rewrite("$<.onplaying")
	return e
}

func (*HTMLElement) SetOnplaying(e *event.Event) {
	js.Rewrite("$<.onplaying = $1", e)
}

func (*HTMLElement) GetOnprogress() (e *event.Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*HTMLElement) SetOnprogress(e *event.Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*HTMLElement) GetOnratechange() (e *event.Event) {
	js.Rewrite("$<.onratechange")
	return e
}

func (*HTMLElement) SetOnratechange(e *event.Event) {
	js.Rewrite("$<.onratechange = $1", e)
}

func (*HTMLElement) GetOnreset() (e *event.Event) {
	js.Rewrite("$<.onreset")
	return e
}

func (*HTMLElement) SetOnreset(e *event.Event) {
	js.Rewrite("$<.onreset = $1", e)
}

func (*HTMLElement) GetOnscroll() (scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*HTMLElement) SetOnscroll(scroll *uievent.UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*HTMLElement) GetOnseeked() (e *event.Event) {
	js.Rewrite("$<.onseeked")
	return e
}

func (*HTMLElement) SetOnseeked(e *event.Event) {
	js.Rewrite("$<.onseeked = $1", e)
}

func (*HTMLElement) GetOnseeking() (e *event.Event) {
	js.Rewrite("$<.onseeking")
	return e
}

func (*HTMLElement) SetOnseeking(e *event.Event) {
	js.Rewrite("$<.onseeking = $1", e)
}

func (*HTMLElement) GetOnselect() (sel *uievent.UIEvent) {
	js.Rewrite("$<.onselect")
	return sel
}

func (*HTMLElement) SetOnselect(sel *uievent.UIEvent) {
	js.Rewrite("$<.onselect = $1", sel)
}

func (*HTMLElement) GetOnselectstart() (selectstart *event.Event) {
	js.Rewrite("$<.onselectstart")
	return selectstart
}

func (*HTMLElement) SetOnselectstart(selectstart *event.Event) {
	js.Rewrite("$<.onselectstart = $1", selectstart)
}

func (*HTMLElement) GetOnstalled() (e *event.Event) {
	js.Rewrite("$<.onstalled")
	return e
}

func (*HTMLElement) SetOnstalled(e *event.Event) {
	js.Rewrite("$<.onstalled = $1", e)
}

func (*HTMLElement) GetOnsubmit() (e *event.Event) {
	js.Rewrite("$<.onsubmit")
	return e
}

func (*HTMLElement) SetOnsubmit(e *event.Event) {
	js.Rewrite("$<.onsubmit = $1", e)
}

func (*HTMLElement) GetOnsuspend() (e *event.Event) {
	js.Rewrite("$<.onsuspend")
	return e
}

func (*HTMLElement) SetOnsuspend(e *event.Event) {
	js.Rewrite("$<.onsuspend = $1", e)
}

func (*HTMLElement) GetOntimeupdate() (e *event.Event) {
	js.Rewrite("$<.ontimeupdate")
	return e
}

func (*HTMLElement) SetOntimeupdate(e *event.Event) {
	js.Rewrite("$<.ontimeupdate = $1", e)
}

func (*HTMLElement) GetOnvolumechange() (e *event.Event) {
	js.Rewrite("$<.onvolumechange")
	return e
}

func (*HTMLElement) SetOnvolumechange(e *event.Event) {
	js.Rewrite("$<.onvolumechange = $1", e)
}

func (*HTMLElement) GetOnwaiting() (e *event.Event) {
	js.Rewrite("$<.onwaiting")
	return e
}

func (*HTMLElement) SetOnwaiting(e *event.Event) {
	js.Rewrite("$<.onwaiting = $1", e)
}

func (*HTMLElement) GetOuterText() (outerText string) {
	js.Rewrite("$<.outerText")
	return outerText
}

func (*HTMLElement) SetOuterText(outerText string) {
	js.Rewrite("$<.outerText = $1", outerText)
}

func (*HTMLElement) GetSpellcheck() (spellcheck bool) {
	js.Rewrite("$<.spellcheck")
	return spellcheck
}

func (*HTMLElement) SetSpellcheck(spellcheck bool) {
	js.Rewrite("$<.spellcheck = $1", spellcheck)
}

func (*HTMLElement) GetStyle() (style *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

func (*HTMLElement) GetTabIndex() (tabIndex int8) {
	js.Rewrite("$<.tabIndex")
	return tabIndex
}

func (*HTMLElement) SetTabIndex(tabIndex int8) {
	js.Rewrite("$<.tabIndex = $1", tabIndex)
}

func (*HTMLElement) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*HTMLElement) SetTitle(title string) {
	js.Rewrite("$<.title = $1", title)
}
