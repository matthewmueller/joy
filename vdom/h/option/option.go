package option

import (
	"encoding/json"
	"strings"

	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/vdom"
)

// Option struct
// js:"option,omit"
type Option struct {
	vdom.Child
	vdom.Node

	attrs    map[string]interface{}
	children []vdom.Child
}

// Props struct
// js:"props,omit"
type Props struct {
	attrs map[string]interface{}
}

// New fn
func New(props *Props, children ...vdom.Child) *Option {
	macro.Rewrite("$1('option', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
	if props == nil {
		props = &Props{attrs: map[string]interface{}{}}
	}
	return &Option{
		attrs:    props.attrs,
		children: children,
	}
}

// Render fn
func (s *Option) Render() vdom.Node {
	return s
}

func (s *Option) String() string {
	// props
	var props []string
	for key, val := range s.attrs {
		bytes, e := json.Marshal(val)
		// TODO: skip over errors?
		if e != nil {
			continue
		}
		props = append(props, key+"="+string(bytes))
	}

	// children
	var children []string
	for _, child := range s.children {
		children = append(children, child.Render().String())
	}

	if len(props) > 0 {
		return "<option " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</option>"
	}

	return "<option>" + strings.Join(children, "") + "</option>"
}

// Accesskey fn
func Accesskey(accesskey string) *Props {
	macro.Rewrite("$1().Set('accesskey', $2)", js.Runtime("Map", "Set", "JSON"), accesskey)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Accesskey(accesskey)
}

// Accesskey fn
func (p *Props) Accesskey(accesskey string) *Props {
	macro.Rewrite("$_.Set('accesskey', $1)", accesskey)
	p.attrs["accesskey"] = accesskey
	return p
}

// Class fn
func Class(class string) *Props {
	macro.Rewrite("$1().Set('class', $2)", js.Runtime("Map", "Set", "JSON"), class)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Class(class)
}

// Class fn
func (p *Props) Class(class string) *Props {
	macro.Rewrite("$_.Set('class', $1)", class)
	p.attrs["class"] = class
	return p
}

// Contenteditable fn
func Contenteditable(contenteditable string) *Props {
	macro.Rewrite("$1().Set('contenteditable', $2)", js.Runtime("Map", "Set", "JSON"), contenteditable)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Contenteditable(contenteditable)
}

// Contenteditable fn
func (p *Props) Contenteditable(contenteditable string) *Props {
	macro.Rewrite("$_.Set('contenteditable', $1)", contenteditable)
	p.attrs["contenteditable"] = contenteditable
	return p
}

// Contextmenu fn
func Contextmenu(contextmenu string) *Props {
	macro.Rewrite("$1().Set('contextmenu', $2)", js.Runtime("Map", "Set", "JSON"), contextmenu)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Contextmenu(contextmenu)
}

// Contextmenu fn
func (p *Props) Contextmenu(contextmenu string) *Props {
	macro.Rewrite("$_.Set('contextmenu', $1)", contextmenu)
	p.attrs["contextmenu"] = contextmenu
	return p
}

// Dir fn
func Dir(dir string) *Props {
	macro.Rewrite("$1().Set('dir', $2)", js.Runtime("Map", "Set", "JSON"), dir)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Dir(dir)
}

// Dir fn
func (p *Props) Dir(dir string) *Props {
	macro.Rewrite("$_.Set('dir', $1)", dir)
	p.attrs["dir"] = dir
	return p
}

// Draggable fn
func Draggable(draggable string) *Props {
	macro.Rewrite("$1().Set('draggable', $2)", js.Runtime("Map", "Set", "JSON"), draggable)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Draggable(draggable)
}

// Draggable fn
func (p *Props) Draggable(draggable string) *Props {
	macro.Rewrite("$_.Set('draggable', $1)", draggable)
	p.attrs["draggable"] = draggable
	return p
}

// Dropzone fn
func Dropzone(dropzone string) *Props {
	macro.Rewrite("$1().Set('dropzone', $2)", js.Runtime("Map", "Set", "JSON"), dropzone)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Dropzone(dropzone)
}

// Dropzone fn
func (p *Props) Dropzone(dropzone string) *Props {
	macro.Rewrite("$_.Set('dropzone', $1)", dropzone)
	p.attrs["dropzone"] = dropzone
	return p
}

// Hidden fn
func Hidden(hidden string) *Props {
	macro.Rewrite("$1().Set('hidden', $2)", js.Runtime("Map", "Set", "JSON"), hidden)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Hidden(hidden)
}

// Hidden fn
func (p *Props) Hidden(hidden string) *Props {
	macro.Rewrite("$_.Set('hidden', $1)", hidden)
	p.attrs["hidden"] = hidden
	return p
}

// ID fn
func ID(id string) *Props {
	macro.Rewrite("$1().Set('id', $2)", js.Runtime("Map", "Set", "JSON"), id)
	p := &Props{attrs: map[string]interface{}{}}
	return p.ID(id)
}

// ID fn
func (p *Props) ID(id string) *Props {
	macro.Rewrite("$_.Set('id', $1)", id)
	p.attrs["id"] = id
	return p
}

// Itemid fn
func Itemid(itemid string) *Props {
	macro.Rewrite("$1().Set('itemid', $2)", js.Runtime("Map", "Set", "JSON"), itemid)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Itemid(itemid)
}

// Itemid fn
func (p *Props) Itemid(itemid string) *Props {
	macro.Rewrite("$_.Set('itemid', $1)", itemid)
	p.attrs["itemid"] = itemid
	return p
}

// Itemprop fn
func Itemprop(itemprop string) *Props {
	macro.Rewrite("$1().Set('itemprop', $2)", js.Runtime("Map", "Set", "JSON"), itemprop)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Itemprop(itemprop)
}

// Itemprop fn
func (p *Props) Itemprop(itemprop string) *Props {
	macro.Rewrite("$_.Set('itemprop', $1)", itemprop)
	p.attrs["itemprop"] = itemprop
	return p
}

// Itemref fn
func Itemref(itemref string) *Props {
	macro.Rewrite("$1().Set('itemref', $2)", js.Runtime("Map", "Set", "JSON"), itemref)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Itemref(itemref)
}

// Itemref fn
func (p *Props) Itemref(itemref string) *Props {
	macro.Rewrite("$_.Set('itemref', $1)", itemref)
	p.attrs["itemref"] = itemref
	return p
}

// Itemscope fn
func Itemscope(itemscope string) *Props {
	macro.Rewrite("$1().Set('itemscope', $2)", js.Runtime("Map", "Set", "JSON"), itemscope)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Itemscope(itemscope)
}

// Itemscope fn
func (p *Props) Itemscope(itemscope string) *Props {
	macro.Rewrite("$_.Set('itemscope', $1)", itemscope)
	p.attrs["itemscope"] = itemscope
	return p
}

// Itemtype fn
func Itemtype(itemtype string) *Props {
	macro.Rewrite("$1().Set('itemtype', $2)", js.Runtime("Map", "Set", "JSON"), itemtype)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Itemtype(itemtype)
}

// Itemtype fn
func (p *Props) Itemtype(itemtype string) *Props {
	macro.Rewrite("$_.Set('itemtype', $1)", itemtype)
	p.attrs["itemtype"] = itemtype
	return p
}

// Lang fn
func Lang(lang string) *Props {
	macro.Rewrite("$1().Set('lang', $2)", js.Runtime("Map", "Set", "JSON"), lang)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Lang(lang)
}

// Lang fn
func (p *Props) Lang(lang string) *Props {
	macro.Rewrite("$_.Set('lang', $1)", lang)
	p.attrs["lang"] = lang
	return p
}

// Spellcheck fn
func Spellcheck(spellcheck string) *Props {
	macro.Rewrite("$1().Set('spellcheck', $2)", js.Runtime("Map", "Set", "JSON"), spellcheck)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Spellcheck(spellcheck)
}

// Spellcheck fn
func (p *Props) Spellcheck(spellcheck string) *Props {
	macro.Rewrite("$_.Set('spellcheck', $1)", spellcheck)
	p.attrs["spellcheck"] = spellcheck
	return p
}

// Style fn
func Style(style string) *Props {
	macro.Rewrite("$1().Set('style', $2)", js.Runtime("Map", "Set", "JSON"), style)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Style(style)
}

// Style fn
func (p *Props) Style(style string) *Props {
	macro.Rewrite("$_.Set('style', $1)", style)
	p.attrs["style"] = style
	return p
}

// Tabindex fn
func Tabindex(tabindex string) *Props {
	macro.Rewrite("$1().Set('tabindex', $2)", js.Runtime("Map", "Set", "JSON"), tabindex)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Tabindex(tabindex)
}

// Tabindex fn
func (p *Props) Tabindex(tabindex string) *Props {
	macro.Rewrite("$_.Set('tabindex', $1)", tabindex)
	p.attrs["tabindex"] = tabindex
	return p
}

// Title fn
func Title(title string) *Props {
	macro.Rewrite("$1().Set('title', $2)", js.Runtime("Map", "Set", "JSON"), title)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Title(title)
}

// Title fn
func (p *Props) Title(title string) *Props {
	macro.Rewrite("$_.Set('title', $1)", title)
	p.attrs["title"] = title
	return p
}

// Translate fn
func Translate(translate string) *Props {
	macro.Rewrite("$1().Set('translate', $2)", js.Runtime("Map", "Set", "JSON"), translate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Translate(translate)
}

// Translate fn
func (p *Props) Translate(translate string) *Props {
	macro.Rewrite("$_.Set('translate', $1)", translate)
	p.attrs["translate"] = translate
	return p
}

// Key fn
func Key(key string) *Props {
	macro.Rewrite("$1().Set('key', $2)", js.Runtime("Map", "Set", "JSON"), key)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Key(key)
}

// Key fn
func (p *Props) Key(key string) *Props {
	macro.Rewrite("$_.Set('key', $1)", key)
	p.attrs["key"] = key
	return p
}

// OnMount fn
func OnMount(onmount string) *Props {
	macro.Rewrite("$1().Set('onmount', $2)", js.Runtime("Map", "Set", "JSON"), onmount)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMount(onmount)
}

// OnMount fn
func (p *Props) OnMount(onmount string) *Props {
	macro.Rewrite("$_.Set('onmount', $1)", onmount)
	p.attrs["onMount"] = onmount
	return p
}

// OnUnmount fn
func OnUnmount(onunmount string) *Props {
	macro.Rewrite("$1().Set('onunmount', $2)", js.Runtime("Map", "Set", "JSON"), onunmount)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnUnmount(onunmount)
}

// OnUnmount fn
func (p *Props) OnUnmount(onunmount string) *Props {
	macro.Rewrite("$_.Set('onunmount', $1)", onunmount)
	p.attrs["onUnmount"] = onunmount
	return p
}

// OnCopy fn
func OnCopy(oncopy string) *Props {
	macro.Rewrite("$1().Set('oncopy', $2)", js.Runtime("Map", "Set", "JSON"), oncopy)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCopy(oncopy)
}

// OnCopy fn
func (p *Props) OnCopy(oncopy string) *Props {
	macro.Rewrite("$_.Set('oncopy', $1)", oncopy)
	p.attrs["onCopy"] = oncopy
	return p
}

// OnCut fn
func OnCut(oncut string) *Props {
	macro.Rewrite("$1().Set('oncut', $2)", js.Runtime("Map", "Set", "JSON"), oncut)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCut(oncut)
}

// OnCut fn
func (p *Props) OnCut(oncut string) *Props {
	macro.Rewrite("$_.Set('oncut', $1)", oncut)
	p.attrs["onCut"] = oncut
	return p
}

// OnPaste fn
func OnPaste(onpaste string) *Props {
	macro.Rewrite("$1().Set('onpaste', $2)", js.Runtime("Map", "Set", "JSON"), onpaste)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnPaste(onpaste)
}

// OnPaste fn
func (p *Props) OnPaste(onpaste string) *Props {
	macro.Rewrite("$_.Set('onpaste', $1)", onpaste)
	p.attrs["onPaste"] = onpaste
	return p
}

// OnCompositionEnd fn
func OnCompositionEnd(oncompositionend string) *Props {
	macro.Rewrite("$1().Set('oncompositionend', $2)", js.Runtime("Map", "Set", "JSON"), oncompositionend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCompositionEnd(oncompositionend)
}

// OnCompositionEnd fn
func (p *Props) OnCompositionEnd(oncompositionend string) *Props {
	macro.Rewrite("$_.Set('oncompositionend', $1)", oncompositionend)
	p.attrs["onCompositionEnd"] = oncompositionend
	return p
}

// OnCompositionStart fn
func OnCompositionStart(oncompositionstart string) *Props {
	macro.Rewrite("$1().Set('oncompositionstart', $2)", js.Runtime("Map", "Set", "JSON"), oncompositionstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCompositionStart(oncompositionstart)
}

// OnCompositionStart fn
func (p *Props) OnCompositionStart(oncompositionstart string) *Props {
	macro.Rewrite("$_.Set('oncompositionstart', $1)", oncompositionstart)
	p.attrs["onCompositionStart"] = oncompositionstart
	return p
}

// OnCompositionUpdate fn
func OnCompositionUpdate(oncompositionupdate string) *Props {
	macro.Rewrite("$1().Set('oncompositionupdate', $2)", js.Runtime("Map", "Set", "JSON"), oncompositionupdate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCompositionUpdate(oncompositionupdate)
}

// OnCompositionUpdate fn
func (p *Props) OnCompositionUpdate(oncompositionupdate string) *Props {
	macro.Rewrite("$_.Set('oncompositionupdate', $1)", oncompositionupdate)
	p.attrs["onCompositionUpdate"] = oncompositionupdate
	return p
}

// OnKeyDown fn
func OnKeyDown(onkeydown string) *Props {
	macro.Rewrite("$1().Set('onkeydown', $2)", js.Runtime("Map", "Set", "JSON"), onkeydown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnKeyDown(onkeydown)
}

// OnKeyDown fn
func (p *Props) OnKeyDown(onkeydown string) *Props {
	macro.Rewrite("$_.Set('onkeydown', $1)", onkeydown)
	p.attrs["onKeyDown"] = onkeydown
	return p
}

// OnKeyPress fn
func OnKeyPress(onkeypress string) *Props {
	macro.Rewrite("$1().Set('onkeypress', $2)", js.Runtime("Map", "Set", "JSON"), onkeypress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnKeyPress(onkeypress)
}

// OnKeyPress fn
func (p *Props) OnKeyPress(onkeypress string) *Props {
	macro.Rewrite("$_.Set('onkeypress', $1)", onkeypress)
	p.attrs["onKeyPress"] = onkeypress
	return p
}

// OnKeyUp fn
func OnKeyUp(onkeyup string) *Props {
	macro.Rewrite("$1().Set('onkeyup', $2)", js.Runtime("Map", "Set", "JSON"), onkeyup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnKeyUp(onkeyup)
}

// OnKeyUp fn
func (p *Props) OnKeyUp(onkeyup string) *Props {
	macro.Rewrite("$_.Set('onkeyup', $1)", onkeyup)
	p.attrs["onKeyUp"] = onkeyup
	return p
}

// OnFocus fn
func OnFocus(onfocus string) *Props {
	macro.Rewrite("$1().Set('onfocus', $2)", js.Runtime("Map", "Set", "JSON"), onfocus)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnFocus(onfocus)
}

// OnFocus fn
func (p *Props) OnFocus(onfocus string) *Props {
	macro.Rewrite("$_.Set('onfocus', $1)", onfocus)
	p.attrs["onFocus"] = onfocus
	return p
}

// OnBlur fn
func OnBlur(onblur string) *Props {
	macro.Rewrite("$1().Set('onblur', $2)", js.Runtime("Map", "Set", "JSON"), onblur)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnBlur(onblur)
}

// OnBlur fn
func (p *Props) OnBlur(onblur string) *Props {
	macro.Rewrite("$_.Set('onblur', $1)", onblur)
	p.attrs["onBlur"] = onblur
	return p
}

// OnChange fn
func OnChange(onchange string) *Props {
	macro.Rewrite("$1().Set('onchange', $2)", js.Runtime("Map", "Set", "JSON"), onchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnChange(onchange)
}

// OnChange fn
func (p *Props) OnChange(onchange string) *Props {
	macro.Rewrite("$_.Set('onchange', $1)", onchange)
	p.attrs["onChange"] = onchange
	return p
}

// OnInput fn
func OnInput(oninput string) *Props {
	macro.Rewrite("$1().Set('oninput', $2)", js.Runtime("Map", "Set", "JSON"), oninput)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnInput(oninput)
}

// OnInput fn
func (p *Props) OnInput(oninput string) *Props {
	macro.Rewrite("$_.Set('oninput', $1)", oninput)
	p.attrs["onInput"] = oninput
	return p
}

// OnSubmit fn
func OnSubmit(onsubmit string) *Props {
	macro.Rewrite("$1().Set('onsubmit', $2)", js.Runtime("Map", "Set", "JSON"), onsubmit)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSubmit(onsubmit)
}

// OnSubmit fn
func (p *Props) OnSubmit(onsubmit string) *Props {
	macro.Rewrite("$_.Set('onsubmit', $1)", onsubmit)
	p.attrs["onSubmit"] = onsubmit
	return p
}

// OnClick fn
func OnClick(onclick string) *Props {
	macro.Rewrite("$1().Set('onclick', $2)", js.Runtime("Map", "Set", "JSON"), onclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnClick(onclick)
}

// OnClick fn
func (p *Props) OnClick(onclick string) *Props {
	macro.Rewrite("$_.Set('onclick', $1)", onclick)
	p.attrs["onClick"] = onclick
	return p
}

// OnContextMenu fn
func OnContextMenu(oncontextmenu string) *Props {
	macro.Rewrite("$1().Set('oncontextmenu', $2)", js.Runtime("Map", "Set", "JSON"), oncontextmenu)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnContextMenu(oncontextmenu)
}

// OnContextMenu fn
func (p *Props) OnContextMenu(oncontextmenu string) *Props {
	macro.Rewrite("$_.Set('oncontextmenu', $1)", oncontextmenu)
	p.attrs["onContextMenu"] = oncontextmenu
	return p
}

// OnDoubleClick fn
func OnDoubleClick(ondoubleclick string) *Props {
	macro.Rewrite("$1().Set('ondoubleclick', $2)", js.Runtime("Map", "Set", "JSON"), ondoubleclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDoubleClick(ondoubleclick)
}

// OnDoubleClick fn
func (p *Props) OnDoubleClick(ondoubleclick string) *Props {
	macro.Rewrite("$_.Set('ondoubleclick', $1)", ondoubleclick)
	p.attrs["onDoubleClick"] = ondoubleclick
	return p
}

// OnDrag fn
func OnDrag(ondrag string) *Props {
	macro.Rewrite("$1().Set('ondrag', $2)", js.Runtime("Map", "Set", "JSON"), ondrag)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDrag(ondrag)
}

// OnDrag fn
func (p *Props) OnDrag(ondrag string) *Props {
	macro.Rewrite("$_.Set('ondrag', $1)", ondrag)
	p.attrs["onDrag"] = ondrag
	return p
}

// OnDragEnd fn
func OnDragEnd(ondragend string) *Props {
	macro.Rewrite("$1().Set('ondragend', $2)", js.Runtime("Map", "Set", "JSON"), ondragend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragEnd(ondragend)
}

// OnDragEnd fn
func (p *Props) OnDragEnd(ondragend string) *Props {
	macro.Rewrite("$_.Set('ondragend', $1)", ondragend)
	p.attrs["onDragEnd"] = ondragend
	return p
}

// OnDragEnter fn
func OnDragEnter(ondragenter string) *Props {
	macro.Rewrite("$1().Set('ondragenter', $2)", js.Runtime("Map", "Set", "JSON"), ondragenter)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragEnter(ondragenter)
}

// OnDragEnter fn
func (p *Props) OnDragEnter(ondragenter string) *Props {
	macro.Rewrite("$_.Set('ondragenter', $1)", ondragenter)
	p.attrs["onDragEnter"] = ondragenter
	return p
}

// OnDragExit fn
func OnDragExit(ondragexit string) *Props {
	macro.Rewrite("$1().Set('ondragexit', $2)", js.Runtime("Map", "Set", "JSON"), ondragexit)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragExit(ondragexit)
}

// OnDragExit fn
func (p *Props) OnDragExit(ondragexit string) *Props {
	macro.Rewrite("$_.Set('ondragexit', $1)", ondragexit)
	p.attrs["onDragExit"] = ondragexit
	return p
}

// OnDragLeave fn
func OnDragLeave(ondragleave string) *Props {
	macro.Rewrite("$1().Set('ondragleave', $2)", js.Runtime("Map", "Set", "JSON"), ondragleave)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragLeave(ondragleave)
}

// OnDragLeave fn
func (p *Props) OnDragLeave(ondragleave string) *Props {
	macro.Rewrite("$_.Set('ondragleave', $1)", ondragleave)
	p.attrs["onDragLeave"] = ondragleave
	return p
}

// OnDragOver fn
func OnDragOver(ondragover string) *Props {
	macro.Rewrite("$1().Set('ondragover', $2)", js.Runtime("Map", "Set", "JSON"), ondragover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragOver(ondragover)
}

// OnDragOver fn
func (p *Props) OnDragOver(ondragover string) *Props {
	macro.Rewrite("$_.Set('ondragover', $1)", ondragover)
	p.attrs["onDragOver"] = ondragover
	return p
}

// OnDragStart fn
func OnDragStart(ondragstart string) *Props {
	macro.Rewrite("$1().Set('ondragstart', $2)", js.Runtime("Map", "Set", "JSON"), ondragstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragStart(ondragstart)
}

// OnDragStart fn
func (p *Props) OnDragStart(ondragstart string) *Props {
	macro.Rewrite("$_.Set('ondragstart', $1)", ondragstart)
	p.attrs["onDragStart"] = ondragstart
	return p
}

// OnDrop fn
func OnDrop(ondrop string) *Props {
	macro.Rewrite("$1().Set('ondrop', $2)", js.Runtime("Map", "Set", "JSON"), ondrop)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDrop(ondrop)
}

// OnDrop fn
func (p *Props) OnDrop(ondrop string) *Props {
	macro.Rewrite("$_.Set('ondrop', $1)", ondrop)
	p.attrs["onDrop"] = ondrop
	return p
}

// OnMouseDown fn
func OnMouseDown(onmousedown string) *Props {
	macro.Rewrite("$1().Set('onmousedown', $2)", js.Runtime("Map", "Set", "JSON"), onmousedown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseDown(onmousedown)
}

// OnMouseDown fn
func (p *Props) OnMouseDown(onmousedown string) *Props {
	macro.Rewrite("$_.Set('onmousedown', $1)", onmousedown)
	p.attrs["onMouseDown"] = onmousedown
	return p
}

// OnMouseEnter fn
func OnMouseEnter(onmouseenter string) *Props {
	macro.Rewrite("$1().Set('onmouseenter', $2)", js.Runtime("Map", "Set", "JSON"), onmouseenter)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseEnter(onmouseenter)
}

// OnMouseEnter fn
func (p *Props) OnMouseEnter(onmouseenter string) *Props {
	macro.Rewrite("$_.Set('onmouseenter', $1)", onmouseenter)
	p.attrs["onMouseEnter"] = onmouseenter
	return p
}

// OnMouseLeave fn
func OnMouseLeave(onmouseleave string) *Props {
	macro.Rewrite("$1().Set('onmouseleave', $2)", js.Runtime("Map", "Set", "JSON"), onmouseleave)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseLeave(onmouseleave)
}

// OnMouseLeave fn
func (p *Props) OnMouseLeave(onmouseleave string) *Props {
	macro.Rewrite("$_.Set('onmouseleave', $1)", onmouseleave)
	p.attrs["onMouseLeave"] = onmouseleave
	return p
}

// OnMouseMove fn
func OnMouseMove(onmousemove string) *Props {
	macro.Rewrite("$1().Set('onmousemove', $2)", js.Runtime("Map", "Set", "JSON"), onmousemove)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseMove(onmousemove)
}

// OnMouseMove fn
func (p *Props) OnMouseMove(onmousemove string) *Props {
	macro.Rewrite("$_.Set('onmousemove', $1)", onmousemove)
	p.attrs["onMouseMove"] = onmousemove
	return p
}

// OnMouseOut fn
func OnMouseOut(onmouseout string) *Props {
	macro.Rewrite("$1().Set('onmouseout', $2)", js.Runtime("Map", "Set", "JSON"), onmouseout)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseOut(onmouseout)
}

// OnMouseOut fn
func (p *Props) OnMouseOut(onmouseout string) *Props {
	macro.Rewrite("$_.Set('onmouseout', $1)", onmouseout)
	p.attrs["onMouseOut"] = onmouseout
	return p
}

// OnMouseOver fn
func OnMouseOver(onmouseover string) *Props {
	macro.Rewrite("$1().Set('onmouseover', $2)", js.Runtime("Map", "Set", "JSON"), onmouseover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseOver(onmouseover)
}

// OnMouseOver fn
func (p *Props) OnMouseOver(onmouseover string) *Props {
	macro.Rewrite("$_.Set('onmouseover', $1)", onmouseover)
	p.attrs["onMouseOver"] = onmouseover
	return p
}

// OnMouseUp fn
func OnMouseUp(onmouseup string) *Props {
	macro.Rewrite("$1().Set('onmouseup', $2)", js.Runtime("Map", "Set", "JSON"), onmouseup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseUp(onmouseup)
}

// OnMouseUp fn
func (p *Props) OnMouseUp(onmouseup string) *Props {
	macro.Rewrite("$_.Set('onmouseup', $1)", onmouseup)
	p.attrs["onMouseUp"] = onmouseup
	return p
}

// OnSelect fn
func OnSelect(onselect string) *Props {
	macro.Rewrite("$1().Set('onselect', $2)", js.Runtime("Map", "Set", "JSON"), onselect)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSelect(onselect)
}

// OnSelect fn
func (p *Props) OnSelect(onselect string) *Props {
	macro.Rewrite("$_.Set('onselect', $1)", onselect)
	p.attrs["onSelect"] = onselect
	return p
}

// OnTouchCancel fn
func OnTouchCancel(ontouchcancel string) *Props {
	macro.Rewrite("$1().Set('ontouchcancel', $2)", js.Runtime("Map", "Set", "JSON"), ontouchcancel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnTouchCancel(ontouchcancel)
}

// OnTouchCancel fn
func (p *Props) OnTouchCancel(ontouchcancel string) *Props {
	macro.Rewrite("$_.Set('ontouchcancel', $1)", ontouchcancel)
	p.attrs["onTouchCancel"] = ontouchcancel
	return p
}

// OnTouchEnd fn
func OnTouchEnd(ontouchend string) *Props {
	macro.Rewrite("$1().Set('ontouchend', $2)", js.Runtime("Map", "Set", "JSON"), ontouchend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnTouchEnd(ontouchend)
}

// OnTouchEnd fn
func (p *Props) OnTouchEnd(ontouchend string) *Props {
	macro.Rewrite("$_.Set('ontouchend', $1)", ontouchend)
	p.attrs["onTouchEnd"] = ontouchend
	return p
}

// OnTouchMove fn
func OnTouchMove(ontouchmove string) *Props {
	macro.Rewrite("$1().Set('ontouchmove', $2)", js.Runtime("Map", "Set", "JSON"), ontouchmove)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnTouchMove(ontouchmove)
}

// OnTouchMove fn
func (p *Props) OnTouchMove(ontouchmove string) *Props {
	macro.Rewrite("$_.Set('ontouchmove', $1)", ontouchmove)
	p.attrs["onTouchMove"] = ontouchmove
	return p
}

// OnTouchStart fn
func OnTouchStart(ontouchstart string) *Props {
	macro.Rewrite("$1().Set('ontouchstart', $2)", js.Runtime("Map", "Set", "JSON"), ontouchstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnTouchStart(ontouchstart)
}

// OnTouchStart fn
func (p *Props) OnTouchStart(ontouchstart string) *Props {
	macro.Rewrite("$_.Set('ontouchstart', $1)", ontouchstart)
	p.attrs["onTouchStart"] = ontouchstart
	return p
}

// OnScroll fn
func OnScroll(onscroll string) *Props {
	macro.Rewrite("$1().Set('onscroll', $2)", js.Runtime("Map", "Set", "JSON"), onscroll)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnScroll(onscroll)
}

// OnScroll fn
func (p *Props) OnScroll(onscroll string) *Props {
	macro.Rewrite("$_.Set('onscroll', $1)", onscroll)
	p.attrs["onScroll"] = onscroll
	return p
}

// OnWheel fn
func OnWheel(onwheel string) *Props {
	macro.Rewrite("$1().Set('onwheel', $2)", js.Runtime("Map", "Set", "JSON"), onwheel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnWheel(onwheel)
}

// OnWheel fn
func (p *Props) OnWheel(onwheel string) *Props {
	macro.Rewrite("$_.Set('onwheel', $1)", onwheel)
	p.attrs["onWheel"] = onwheel
	return p
}

// OnAbort fn
func OnAbort(onabort string) *Props {
	macro.Rewrite("$1().Set('onabort', $2)", js.Runtime("Map", "Set", "JSON"), onabort)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnAbort(onabort)
}

// OnAbort fn
func (p *Props) OnAbort(onabort string) *Props {
	macro.Rewrite("$_.Set('onabort', $1)", onabort)
	p.attrs["onAbort"] = onabort
	return p
}

// OnCanPlay fn
func OnCanPlay(oncanplay string) *Props {
	macro.Rewrite("$1().Set('oncanplay', $2)", js.Runtime("Map", "Set", "JSON"), oncanplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCanPlay(oncanplay)
}

// OnCanPlay fn
func (p *Props) OnCanPlay(oncanplay string) *Props {
	macro.Rewrite("$_.Set('oncanplay', $1)", oncanplay)
	p.attrs["onCanPlay"] = oncanplay
	return p
}

// OnCanPlayThrough fn
func OnCanPlayThrough(oncanplaythrough string) *Props {
	macro.Rewrite("$1().Set('oncanplaythrough', $2)", js.Runtime("Map", "Set", "JSON"), oncanplaythrough)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCanPlayThrough(oncanplaythrough)
}

// OnCanPlayThrough fn
func (p *Props) OnCanPlayThrough(oncanplaythrough string) *Props {
	macro.Rewrite("$_.Set('oncanplaythrough', $1)", oncanplaythrough)
	p.attrs["onCanPlayThrough"] = oncanplaythrough
	return p
}

// OnDurationChange fn
func OnDurationChange(ondurationchange string) *Props {
	macro.Rewrite("$1().Set('ondurationchange', $2)", js.Runtime("Map", "Set", "JSON"), ondurationchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDurationChange(ondurationchange)
}

// OnDurationChange fn
func (p *Props) OnDurationChange(ondurationchange string) *Props {
	macro.Rewrite("$_.Set('ondurationchange', $1)", ondurationchange)
	p.attrs["onDurationChange"] = ondurationchange
	return p
}

// OnEmptied fn
func OnEmptied(onemptied string) *Props {
	macro.Rewrite("$1().Set('onemptied', $2)", js.Runtime("Map", "Set", "JSON"), onemptied)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnEmptied(onemptied)
}

// OnEmptied fn
func (p *Props) OnEmptied(onemptied string) *Props {
	macro.Rewrite("$_.Set('onemptied', $1)", onemptied)
	p.attrs["onEmptied"] = onemptied
	return p
}

// OnEncrypted fn
func OnEncrypted(onencrypted string) *Props {
	macro.Rewrite("$1().Set('onencrypted', $2)", js.Runtime("Map", "Set", "JSON"), onencrypted)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnEncrypted(onencrypted)
}

// OnEncrypted fn
func (p *Props) OnEncrypted(onencrypted string) *Props {
	macro.Rewrite("$_.Set('onencrypted', $1)", onencrypted)
	p.attrs["onEncrypted"] = onencrypted
	return p
}

// OnEnded fn
func OnEnded(onended string) *Props {
	macro.Rewrite("$1().Set('onended', $2)", js.Runtime("Map", "Set", "JSON"), onended)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnEnded(onended)
}

// OnEnded fn
func (p *Props) OnEnded(onended string) *Props {
	macro.Rewrite("$_.Set('onended', $1)", onended)
	p.attrs["onEnded"] = onended
	return p
}

// OnError fn
func OnError(onerror string) *Props {
	macro.Rewrite("$1().Set('onerror', $2)", js.Runtime("Map", "Set", "JSON"), onerror)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnError(onerror)
}

// OnError fn
func (p *Props) OnError(onerror string) *Props {
	macro.Rewrite("$_.Set('onerror', $1)", onerror)
	p.attrs["onError"] = onerror
	return p
}

// OnLoadedData fn
func OnLoadedData(onloadeddata string) *Props {
	macro.Rewrite("$1().Set('onloadeddata', $2)", js.Runtime("Map", "Set", "JSON"), onloadeddata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoadedData(onloadeddata)
}

// OnLoadedData fn
func (p *Props) OnLoadedData(onloadeddata string) *Props {
	macro.Rewrite("$_.Set('onloadeddata', $1)", onloadeddata)
	p.attrs["onLoadedData"] = onloadeddata
	return p
}

// OnLoadedMetadata fn
func OnLoadedMetadata(onloadedmetadata string) *Props {
	macro.Rewrite("$1().Set('onloadedmetadata', $2)", js.Runtime("Map", "Set", "JSON"), onloadedmetadata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoadedMetadata(onloadedmetadata)
}

// OnLoadedMetadata fn
func (p *Props) OnLoadedMetadata(onloadedmetadata string) *Props {
	macro.Rewrite("$_.Set('onloadedmetadata', $1)", onloadedmetadata)
	p.attrs["onLoadedMetadata"] = onloadedmetadata
	return p
}

// OnLoadStart fn
func OnLoadStart(onloadstart string) *Props {
	macro.Rewrite("$1().Set('onloadstart', $2)", js.Runtime("Map", "Set", "JSON"), onloadstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoadStart(onloadstart)
}

// OnLoadStart fn
func (p *Props) OnLoadStart(onloadstart string) *Props {
	macro.Rewrite("$_.Set('onloadstart', $1)", onloadstart)
	p.attrs["onLoadStart"] = onloadstart
	return p
}

// OnPause fn
func OnPause(onpause string) *Props {
	macro.Rewrite("$1().Set('onpause', $2)", js.Runtime("Map", "Set", "JSON"), onpause)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnPause(onpause)
}

// OnPause fn
func (p *Props) OnPause(onpause string) *Props {
	macro.Rewrite("$_.Set('onpause', $1)", onpause)
	p.attrs["onPause"] = onpause
	return p
}

// OnPlay fn
func OnPlay(onplay string) *Props {
	macro.Rewrite("$1().Set('onplay', $2)", js.Runtime("Map", "Set", "JSON"), onplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnPlay(onplay)
}

// OnPlay fn
func (p *Props) OnPlay(onplay string) *Props {
	macro.Rewrite("$_.Set('onplay', $1)", onplay)
	p.attrs["onPlay"] = onplay
	return p
}

// OnPlaying fn
func OnPlaying(onplaying string) *Props {
	macro.Rewrite("$1().Set('onplaying', $2)", js.Runtime("Map", "Set", "JSON"), onplaying)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnPlaying(onplaying)
}

// OnPlaying fn
func (p *Props) OnPlaying(onplaying string) *Props {
	macro.Rewrite("$_.Set('onplaying', $1)", onplaying)
	p.attrs["onPlaying"] = onplaying
	return p
}

// OnProgress fn
func OnProgress(onprogress string) *Props {
	macro.Rewrite("$1().Set('onprogress', $2)", js.Runtime("Map", "Set", "JSON"), onprogress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnProgress(onprogress)
}

// OnProgress fn
func (p *Props) OnProgress(onprogress string) *Props {
	macro.Rewrite("$_.Set('onprogress', $1)", onprogress)
	p.attrs["onProgress"] = onprogress
	return p
}

// OnRateChange fn
func OnRateChange(onratechange string) *Props {
	macro.Rewrite("$1().Set('onratechange', $2)", js.Runtime("Map", "Set", "JSON"), onratechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnRateChange(onratechange)
}

// OnRateChange fn
func (p *Props) OnRateChange(onratechange string) *Props {
	macro.Rewrite("$_.Set('onratechange', $1)", onratechange)
	p.attrs["onRateChange"] = onratechange
	return p
}

// OnSeeked fn
func OnSeeked(onseeked string) *Props {
	macro.Rewrite("$1().Set('onseeked', $2)", js.Runtime("Map", "Set", "JSON"), onseeked)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSeeked(onseeked)
}

// OnSeeked fn
func (p *Props) OnSeeked(onseeked string) *Props {
	macro.Rewrite("$_.Set('onseeked', $1)", onseeked)
	p.attrs["onSeeked"] = onseeked
	return p
}

// OnSeeking fn
func OnSeeking(onseeking string) *Props {
	macro.Rewrite("$1().Set('onseeking', $2)", js.Runtime("Map", "Set", "JSON"), onseeking)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSeeking(onseeking)
}

// OnSeeking fn
func (p *Props) OnSeeking(onseeking string) *Props {
	macro.Rewrite("$_.Set('onseeking', $1)", onseeking)
	p.attrs["onSeeking"] = onseeking
	return p
}

// OnStalled fn
func OnStalled(onstalled string) *Props {
	macro.Rewrite("$1().Set('onstalled', $2)", js.Runtime("Map", "Set", "JSON"), onstalled)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnStalled(onstalled)
}

// OnStalled fn
func (p *Props) OnStalled(onstalled string) *Props {
	macro.Rewrite("$_.Set('onstalled', $1)", onstalled)
	p.attrs["onStalled"] = onstalled
	return p
}

// OnSuspend fn
func OnSuspend(onsuspend string) *Props {
	macro.Rewrite("$1().Set('onsuspend', $2)", js.Runtime("Map", "Set", "JSON"), onsuspend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSuspend(onsuspend)
}

// OnSuspend fn
func (p *Props) OnSuspend(onsuspend string) *Props {
	macro.Rewrite("$_.Set('onsuspend', $1)", onsuspend)
	p.attrs["onSuspend"] = onsuspend
	return p
}

// OnTimeUpdate fn
func OnTimeUpdate(ontimeupdate string) *Props {
	macro.Rewrite("$1().Set('ontimeupdate', $2)", js.Runtime("Map", "Set", "JSON"), ontimeupdate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnTimeUpdate(ontimeupdate)
}

// OnTimeUpdate fn
func (p *Props) OnTimeUpdate(ontimeupdate string) *Props {
	macro.Rewrite("$_.Set('ontimeupdate', $1)", ontimeupdate)
	p.attrs["onTimeUpdate"] = ontimeupdate
	return p
}

// OnVolumeChange fn
func OnVolumeChange(onvolumechange string) *Props {
	macro.Rewrite("$1().Set('onvolumechange', $2)", js.Runtime("Map", "Set", "JSON"), onvolumechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnVolumeChange(onvolumechange)
}

// OnVolumeChange fn
func (p *Props) OnVolumeChange(onvolumechange string) *Props {
	macro.Rewrite("$_.Set('onvolumechange', $1)", onvolumechange)
	p.attrs["onVolumeChange"] = onvolumechange
	return p
}

// OnWaiting fn
func OnWaiting(onwaiting string) *Props {
	macro.Rewrite("$1().Set('onwaiting', $2)", js.Runtime("Map", "Set", "JSON"), onwaiting)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnWaiting(onwaiting)
}

// OnWaiting fn
func (p *Props) OnWaiting(onwaiting string) *Props {
	macro.Rewrite("$_.Set('onwaiting', $1)", onwaiting)
	p.attrs["onWaiting"] = onwaiting
	return p
}

// OnLoad fn
func OnLoad(onload string) *Props {
	macro.Rewrite("$1().Set('onload', $2)", js.Runtime("Map", "Set", "JSON"), onload)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoad(onload)
}

// OnLoad fn
func (p *Props) OnLoad(onload string) *Props {
	macro.Rewrite("$_.Set('onload', $1)", onload)
	p.attrs["onLoad"] = onload
	return p
}

// OnAnimationStart fn
func OnAnimationStart(onanimationstart string) *Props {
	macro.Rewrite("$1().Set('onanimationstart', $2)", js.Runtime("Map", "Set", "JSON"), onanimationstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnAnimationStart(onanimationstart)
}

// OnAnimationStart fn
func (p *Props) OnAnimationStart(onanimationstart string) *Props {
	macro.Rewrite("$_.Set('onanimationstart', $1)", onanimationstart)
	p.attrs["onAnimationStart"] = onanimationstart
	return p
}

// OnAnimationEnd fn
func OnAnimationEnd(onanimationend string) *Props {
	macro.Rewrite("$1().Set('onanimationend', $2)", js.Runtime("Map", "Set", "JSON"), onanimationend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnAnimationEnd(onanimationend)
}

// OnAnimationEnd fn
func (p *Props) OnAnimationEnd(onanimationend string) *Props {
	macro.Rewrite("$_.Set('onanimationend', $1)", onanimationend)
	p.attrs["onAnimationEnd"] = onanimationend
	return p
}

// OnAnimationIteration fn
func OnAnimationIteration(onanimationiteration string) *Props {
	macro.Rewrite("$1().Set('onanimationiteration', $2)", js.Runtime("Map", "Set", "JSON"), onanimationiteration)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnAnimationIteration(onanimationiteration)
}

// OnAnimationIteration fn
func (p *Props) OnAnimationIteration(onanimationiteration string) *Props {
	macro.Rewrite("$_.Set('onanimationiteration', $1)", onanimationiteration)
	p.attrs["onAnimationIteration"] = onanimationiteration
	return p
}

// OnTransitionEnd fn
func OnTransitionEnd(ontransitionend string) *Props {
	macro.Rewrite("$1().Set('ontransitionend', $2)", js.Runtime("Map", "Set", "JSON"), ontransitionend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnTransitionEnd(ontransitionend)
}

// OnTransitionEnd fn
func (p *Props) OnTransitionEnd(ontransitionend string) *Props {
	macro.Rewrite("$_.Set('ontransitionend', $1)", ontransitionend)
	p.attrs["onTransitionEnd"] = ontransitionend
	return p
}

// Disabled fn
func Disabled(disabled string) *Props {
	macro.Rewrite("$1().Set('disabled', $2)", js.Runtime("Map", "Set", "JSON"), disabled)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Disabled(disabled)
}

// Disabled fn
func (p *Props) Disabled(disabled string) *Props {
	macro.Rewrite("$_.Set('disabled', $1)", disabled)
	p.attrs["disabled"] = disabled
	return p
}

// Label fn
func Label(label string) *Props {
	macro.Rewrite("$1().Set('label', $2)", js.Runtime("Map", "Set", "JSON"), label)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Label(label)
}

// Label fn
func (p *Props) Label(label string) *Props {
	macro.Rewrite("$_.Set('label', $1)", label)
	p.attrs["label"] = label
	return p
}

// Selected fn
func Selected(selected string) *Props {
	macro.Rewrite("$1().Set('selected', $2)", js.Runtime("Map", "Set", "JSON"), selected)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Selected(selected)
}

// Selected fn
func (p *Props) Selected(selected string) *Props {
	macro.Rewrite("$_.Set('selected', $1)", selected)
	p.attrs["selected"] = selected
	return p
}

// Value fn
func Value(value string) *Props {
	macro.Rewrite("$1().Set('value', $2)", js.Runtime("Map", "Set", "JSON"), value)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Value(value)
}

// Value fn
func (p *Props) Value(value string) *Props {
	macro.Rewrite("$_.Set('value', $1)", value)
	p.attrs["value"] = value
	return p
}

// Attr fn
func Attr(key string, value interface{}) *Props {
	macro.Rewrite("$1().Set($2, $3)", js.Runtime("Map", "Set", "JSON"), key, value)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Attr(key, value)
}

// Attr fn
func (p *Props) Attr(key string, value interface{}) *Props {
	macro.Rewrite("$_.Set($1, $2)", key, value)
	p.attrs[key] = value
	return p
}
