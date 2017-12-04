package mp

import (
	"encoding/json"
	"strings"

	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/vdom"
)

// Map struct
// js:"mp,omit"
type Map struct {
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

// New mp element
//
// The map element, in conjunction with an img element and any area element descendants, defines an image map. The element represents its children.
func New(props *Props, children ...vdom.Child) *Map {
	macro.Rewrite("$1('map', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
	if props == nil {
		props = &Props{attrs: map[string]interface{}{}}
	}
	return &Map{
		attrs:    props.attrs,
		children: children,
	}
}

// Render fn
func (s *Map) Render() vdom.Node {
	return s
}

func (s *Map) String() string {
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
		return "<mp " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</mp>"
	}

	return "<mp>" + strings.Join(children, "") + "</mp>"
}

// Name fn
func Name(name string) *Props {
	macro.Rewrite("$1().Set('name', $2)", macro.Runtime("Map", "Set", "JSON"), name)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Name(name)
}

// Name fn
func (p *Props) Name(name string) *Props {
	macro.Rewrite("$_.Set('name', $1)", name)
	p.attrs["name"] = name
	return p
}

// Onabort fn
func Onabort(onabort string) *Props {
	macro.Rewrite("$1().Set('onabort', $2)", macro.Runtime("Map", "Set", "JSON"), onabort)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onabort(onabort)
}

// Onabort fn
func (p *Props) Onabort(onabort string) *Props {
	macro.Rewrite("$_.Set('onabort', $1)", onabort)
	p.attrs["onabort"] = onabort
	return p
}

// Onblur fn
func Onblur(onblur string) *Props {
	macro.Rewrite("$1().Set('onblur', $2)", macro.Runtime("Map", "Set", "JSON"), onblur)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onblur(onblur)
}

// Onblur fn
func (p *Props) Onblur(onblur string) *Props {
	macro.Rewrite("$_.Set('onblur', $1)", onblur)
	p.attrs["onblur"] = onblur
	return p
}

// Oncanplay fn
func Oncanplay(oncanplay string) *Props {
	macro.Rewrite("$1().Set('oncanplay', $2)", macro.Runtime("Map", "Set", "JSON"), oncanplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncanplay(oncanplay)
}

// Oncanplay fn
func (p *Props) Oncanplay(oncanplay string) *Props {
	macro.Rewrite("$_.Set('oncanplay', $1)", oncanplay)
	p.attrs["oncanplay"] = oncanplay
	return p
}

// Oncanplaythrough fn
func Oncanplaythrough(oncanplaythrough string) *Props {
	macro.Rewrite("$1().Set('oncanplaythrough', $2)", macro.Runtime("Map", "Set", "JSON"), oncanplaythrough)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncanplaythrough(oncanplaythrough)
}

// Oncanplaythrough fn
func (p *Props) Oncanplaythrough(oncanplaythrough string) *Props {
	macro.Rewrite("$_.Set('oncanplaythrough', $1)", oncanplaythrough)
	p.attrs["oncanplaythrough"] = oncanplaythrough
	return p
}

// Onchange fn
func Onchange(onchange string) *Props {
	macro.Rewrite("$1().Set('onchange', $2)", macro.Runtime("Map", "Set", "JSON"), onchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onchange(onchange)
}

// Onchange fn
func (p *Props) Onchange(onchange string) *Props {
	macro.Rewrite("$_.Set('onchange', $1)", onchange)
	p.attrs["onchange"] = onchange
	return p
}

// Onclick fn
func Onclick(onclick string) *Props {
	macro.Rewrite("$1().Set('onclick', $2)", macro.Runtime("Map", "Set", "JSON"), onclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onclick(onclick)
}

// Onclick fn
func (p *Props) Onclick(onclick string) *Props {
	macro.Rewrite("$_.Set('onclick', $1)", onclick)
	p.attrs["onclick"] = onclick
	return p
}

// Oncontextmenu fn
func Oncontextmenu(oncontextmenu string) *Props {
	macro.Rewrite("$1().Set('oncontextmenu', $2)", macro.Runtime("Map", "Set", "JSON"), oncontextmenu)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncontextmenu(oncontextmenu)
}

// Oncontextmenu fn
func (p *Props) Oncontextmenu(oncontextmenu string) *Props {
	macro.Rewrite("$_.Set('oncontextmenu', $1)", oncontextmenu)
	p.attrs["oncontextmenu"] = oncontextmenu
	return p
}

// Ondblclick fn
func Ondblclick(ondblclick string) *Props {
	macro.Rewrite("$1().Set('ondblclick', $2)", macro.Runtime("Map", "Set", "JSON"), ondblclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondblclick(ondblclick)
}

// Ondblclick fn
func (p *Props) Ondblclick(ondblclick string) *Props {
	macro.Rewrite("$_.Set('ondblclick', $1)", ondblclick)
	p.attrs["ondblclick"] = ondblclick
	return p
}

// Ondrag fn
func Ondrag(ondrag string) *Props {
	macro.Rewrite("$1().Set('ondrag', $2)", macro.Runtime("Map", "Set", "JSON"), ondrag)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondrag(ondrag)
}

// Ondrag fn
func (p *Props) Ondrag(ondrag string) *Props {
	macro.Rewrite("$_.Set('ondrag', $1)", ondrag)
	p.attrs["ondrag"] = ondrag
	return p
}

// Ondragend fn
func Ondragend(ondragend string) *Props {
	macro.Rewrite("$1().Set('ondragend', $2)", macro.Runtime("Map", "Set", "JSON"), ondragend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragend(ondragend)
}

// Ondragend fn
func (p *Props) Ondragend(ondragend string) *Props {
	macro.Rewrite("$_.Set('ondragend', $1)", ondragend)
	p.attrs["ondragend"] = ondragend
	return p
}

// Ondragenter fn
func Ondragenter(ondragenter string) *Props {
	macro.Rewrite("$1().Set('ondragenter', $2)", macro.Runtime("Map", "Set", "JSON"), ondragenter)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragenter(ondragenter)
}

// Ondragenter fn
func (p *Props) Ondragenter(ondragenter string) *Props {
	macro.Rewrite("$_.Set('ondragenter', $1)", ondragenter)
	p.attrs["ondragenter"] = ondragenter
	return p
}

// Ondragleave fn
func Ondragleave(ondragleave string) *Props {
	macro.Rewrite("$1().Set('ondragleave', $2)", macro.Runtime("Map", "Set", "JSON"), ondragleave)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragleave(ondragleave)
}

// Ondragleave fn
func (p *Props) Ondragleave(ondragleave string) *Props {
	macro.Rewrite("$_.Set('ondragleave', $1)", ondragleave)
	p.attrs["ondragleave"] = ondragleave
	return p
}

// Ondragover fn
func Ondragover(ondragover string) *Props {
	macro.Rewrite("$1().Set('ondragover', $2)", macro.Runtime("Map", "Set", "JSON"), ondragover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragover(ondragover)
}

// Ondragover fn
func (p *Props) Ondragover(ondragover string) *Props {
	macro.Rewrite("$_.Set('ondragover', $1)", ondragover)
	p.attrs["ondragover"] = ondragover
	return p
}

// Ondragstart fn
func Ondragstart(ondragstart string) *Props {
	macro.Rewrite("$1().Set('ondragstart', $2)", macro.Runtime("Map", "Set", "JSON"), ondragstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragstart(ondragstart)
}

// Ondragstart fn
func (p *Props) Ondragstart(ondragstart string) *Props {
	macro.Rewrite("$_.Set('ondragstart', $1)", ondragstart)
	p.attrs["ondragstart"] = ondragstart
	return p
}

// Ondrop fn
func Ondrop(ondrop string) *Props {
	macro.Rewrite("$1().Set('ondrop', $2)", macro.Runtime("Map", "Set", "JSON"), ondrop)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondrop(ondrop)
}

// Ondrop fn
func (p *Props) Ondrop(ondrop string) *Props {
	macro.Rewrite("$_.Set('ondrop', $1)", ondrop)
	p.attrs["ondrop"] = ondrop
	return p
}

// Ondurationchange fn
func Ondurationchange(ondurationchange string) *Props {
	macro.Rewrite("$1().Set('ondurationchange', $2)", macro.Runtime("Map", "Set", "JSON"), ondurationchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondurationchange(ondurationchange)
}

// Ondurationchange fn
func (p *Props) Ondurationchange(ondurationchange string) *Props {
	macro.Rewrite("$_.Set('ondurationchange', $1)", ondurationchange)
	p.attrs["ondurationchange"] = ondurationchange
	return p
}

// Onemptied fn
func Onemptied(onemptied string) *Props {
	macro.Rewrite("$1().Set('onemptied', $2)", macro.Runtime("Map", "Set", "JSON"), onemptied)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onemptied(onemptied)
}

// Onemptied fn
func (p *Props) Onemptied(onemptied string) *Props {
	macro.Rewrite("$_.Set('onemptied', $1)", onemptied)
	p.attrs["onemptied"] = onemptied
	return p
}

// Onended fn
func Onended(onended string) *Props {
	macro.Rewrite("$1().Set('onended', $2)", macro.Runtime("Map", "Set", "JSON"), onended)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onended(onended)
}

// Onended fn
func (p *Props) Onended(onended string) *Props {
	macro.Rewrite("$_.Set('onended', $1)", onended)
	p.attrs["onended"] = onended
	return p
}

// Onerror fn
func Onerror(onerror string) *Props {
	macro.Rewrite("$1().Set('onerror', $2)", macro.Runtime("Map", "Set", "JSON"), onerror)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onerror(onerror)
}

// Onerror fn
func (p *Props) Onerror(onerror string) *Props {
	macro.Rewrite("$_.Set('onerror', $1)", onerror)
	p.attrs["onerror"] = onerror
	return p
}

// Onfocus fn
func Onfocus(onfocus string) *Props {
	macro.Rewrite("$1().Set('onfocus', $2)", macro.Runtime("Map", "Set", "JSON"), onfocus)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onfocus(onfocus)
}

// Onfocus fn
func (p *Props) Onfocus(onfocus string) *Props {
	macro.Rewrite("$_.Set('onfocus', $1)", onfocus)
	p.attrs["onfocus"] = onfocus
	return p
}

// Onformchange fn
func Onformchange(onformchange string) *Props {
	macro.Rewrite("$1().Set('onformchange', $2)", macro.Runtime("Map", "Set", "JSON"), onformchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onformchange(onformchange)
}

// Onformchange fn
func (p *Props) Onformchange(onformchange string) *Props {
	macro.Rewrite("$_.Set('onformchange', $1)", onformchange)
	p.attrs["onformchange"] = onformchange
	return p
}

// Onforminput fn
func Onforminput(onforminput string) *Props {
	macro.Rewrite("$1().Set('onforminput', $2)", macro.Runtime("Map", "Set", "JSON"), onforminput)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onforminput(onforminput)
}

// Onforminput fn
func (p *Props) Onforminput(onforminput string) *Props {
	macro.Rewrite("$_.Set('onforminput', $1)", onforminput)
	p.attrs["onforminput"] = onforminput
	return p
}

// Oninput fn
func Oninput(oninput string) *Props {
	macro.Rewrite("$1().Set('oninput', $2)", macro.Runtime("Map", "Set", "JSON"), oninput)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oninput(oninput)
}

// Oninput fn
func (p *Props) Oninput(oninput string) *Props {
	macro.Rewrite("$_.Set('oninput', $1)", oninput)
	p.attrs["oninput"] = oninput
	return p
}

// Oninvalid fn
func Oninvalid(oninvalid string) *Props {
	macro.Rewrite("$1().Set('oninvalid', $2)", macro.Runtime("Map", "Set", "JSON"), oninvalid)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oninvalid(oninvalid)
}

// Oninvalid fn
func (p *Props) Oninvalid(oninvalid string) *Props {
	macro.Rewrite("$_.Set('oninvalid', $1)", oninvalid)
	p.attrs["oninvalid"] = oninvalid
	return p
}

// Onkeydown fn
func Onkeydown(onkeydown string) *Props {
	macro.Rewrite("$1().Set('onkeydown', $2)", macro.Runtime("Map", "Set", "JSON"), onkeydown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeydown(onkeydown)
}

// Onkeydown fn
func (p *Props) Onkeydown(onkeydown string) *Props {
	macro.Rewrite("$_.Set('onkeydown', $1)", onkeydown)
	p.attrs["onkeydown"] = onkeydown
	return p
}

// Onkeypress fn
func Onkeypress(onkeypress string) *Props {
	macro.Rewrite("$1().Set('onkeypress', $2)", macro.Runtime("Map", "Set", "JSON"), onkeypress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeypress(onkeypress)
}

// Onkeypress fn
func (p *Props) Onkeypress(onkeypress string) *Props {
	macro.Rewrite("$_.Set('onkeypress', $1)", onkeypress)
	p.attrs["onkeypress"] = onkeypress
	return p
}

// Onkeyup fn
func Onkeyup(onkeyup string) *Props {
	macro.Rewrite("$1().Set('onkeyup', $2)", macro.Runtime("Map", "Set", "JSON"), onkeyup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeyup(onkeyup)
}

// Onkeyup fn
func (p *Props) Onkeyup(onkeyup string) *Props {
	macro.Rewrite("$_.Set('onkeyup', $1)", onkeyup)
	p.attrs["onkeyup"] = onkeyup
	return p
}

// Onload fn
func Onload(onload string) *Props {
	macro.Rewrite("$1().Set('onload', $2)", macro.Runtime("Map", "Set", "JSON"), onload)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onload(onload)
}

// Onload fn
func (p *Props) Onload(onload string) *Props {
	macro.Rewrite("$_.Set('onload', $1)", onload)
	p.attrs["onload"] = onload
	return p
}

// Onloadeddata fn
func Onloadeddata(onloadeddata string) *Props {
	macro.Rewrite("$1().Set('onloadeddata', $2)", macro.Runtime("Map", "Set", "JSON"), onloadeddata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadeddata(onloadeddata)
}

// Onloadeddata fn
func (p *Props) Onloadeddata(onloadeddata string) *Props {
	macro.Rewrite("$_.Set('onloadeddata', $1)", onloadeddata)
	p.attrs["onloadeddata"] = onloadeddata
	return p
}

// Onloadedmetadata fn
func Onloadedmetadata(onloadedmetadata string) *Props {
	macro.Rewrite("$1().Set('onloadedmetadata', $2)", macro.Runtime("Map", "Set", "JSON"), onloadedmetadata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadedmetadata(onloadedmetadata)
}

// Onloadedmetadata fn
func (p *Props) Onloadedmetadata(onloadedmetadata string) *Props {
	macro.Rewrite("$_.Set('onloadedmetadata', $1)", onloadedmetadata)
	p.attrs["onloadedmetadata"] = onloadedmetadata
	return p
}

// Onloadstart fn
func Onloadstart(onloadstart string) *Props {
	macro.Rewrite("$1().Set('onloadstart', $2)", macro.Runtime("Map", "Set", "JSON"), onloadstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadstart(onloadstart)
}

// Onloadstart fn
func (p *Props) Onloadstart(onloadstart string) *Props {
	macro.Rewrite("$_.Set('onloadstart', $1)", onloadstart)
	p.attrs["onloadstart"] = onloadstart
	return p
}

// Onmousedown fn
func Onmousedown(onmousedown string) *Props {
	macro.Rewrite("$1().Set('onmousedown', $2)", macro.Runtime("Map", "Set", "JSON"), onmousedown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousedown(onmousedown)
}

// Onmousedown fn
func (p *Props) Onmousedown(onmousedown string) *Props {
	macro.Rewrite("$_.Set('onmousedown', $1)", onmousedown)
	p.attrs["onmousedown"] = onmousedown
	return p
}

// Onmousemove fn
func Onmousemove(onmousemove string) *Props {
	macro.Rewrite("$1().Set('onmousemove', $2)", macro.Runtime("Map", "Set", "JSON"), onmousemove)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousemove(onmousemove)
}

// Onmousemove fn
func (p *Props) Onmousemove(onmousemove string) *Props {
	macro.Rewrite("$_.Set('onmousemove', $1)", onmousemove)
	p.attrs["onmousemove"] = onmousemove
	return p
}

// Onmouseout fn
func Onmouseout(onmouseout string) *Props {
	macro.Rewrite("$1().Set('onmouseout', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseout)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseout(onmouseout)
}

// Onmouseout fn
func (p *Props) Onmouseout(onmouseout string) *Props {
	macro.Rewrite("$_.Set('onmouseout', $1)", onmouseout)
	p.attrs["onmouseout"] = onmouseout
	return p
}

// Onmouseover fn
func Onmouseover(onmouseover string) *Props {
	macro.Rewrite("$1().Set('onmouseover', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseover(onmouseover)
}

// Onmouseover fn
func (p *Props) Onmouseover(onmouseover string) *Props {
	macro.Rewrite("$_.Set('onmouseover', $1)", onmouseover)
	p.attrs["onmouseover"] = onmouseover
	return p
}

// Onmouseup fn
func Onmouseup(onmouseup string) *Props {
	macro.Rewrite("$1().Set('onmouseup', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseup(onmouseup)
}

// Onmouseup fn
func (p *Props) Onmouseup(onmouseup string) *Props {
	macro.Rewrite("$_.Set('onmouseup', $1)", onmouseup)
	p.attrs["onmouseup"] = onmouseup
	return p
}

// Onmousewheel fn
func Onmousewheel(onmousewheel string) *Props {
	macro.Rewrite("$1().Set('onmousewheel', $2)", macro.Runtime("Map", "Set", "JSON"), onmousewheel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousewheel(onmousewheel)
}

// Onmousewheel fn
func (p *Props) Onmousewheel(onmousewheel string) *Props {
	macro.Rewrite("$_.Set('onmousewheel', $1)", onmousewheel)
	p.attrs["onmousewheel"] = onmousewheel
	return p
}

// Onpause fn
func Onpause(onpause string) *Props {
	macro.Rewrite("$1().Set('onpause', $2)", macro.Runtime("Map", "Set", "JSON"), onpause)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onpause(onpause)
}

// Onpause fn
func (p *Props) Onpause(onpause string) *Props {
	macro.Rewrite("$_.Set('onpause', $1)", onpause)
	p.attrs["onpause"] = onpause
	return p
}

// Onplay fn
func Onplay(onplay string) *Props {
	macro.Rewrite("$1().Set('onplay', $2)", macro.Runtime("Map", "Set", "JSON"), onplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onplay(onplay)
}

// Onplay fn
func (p *Props) Onplay(onplay string) *Props {
	macro.Rewrite("$_.Set('onplay', $1)", onplay)
	p.attrs["onplay"] = onplay
	return p
}

// Onplaying fn
func Onplaying(onplaying string) *Props {
	macro.Rewrite("$1().Set('onplaying', $2)", macro.Runtime("Map", "Set", "JSON"), onplaying)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onplaying(onplaying)
}

// Onplaying fn
func (p *Props) Onplaying(onplaying string) *Props {
	macro.Rewrite("$_.Set('onplaying', $1)", onplaying)
	p.attrs["onplaying"] = onplaying
	return p
}

// Onprogress fn
func Onprogress(onprogress string) *Props {
	macro.Rewrite("$1().Set('onprogress', $2)", macro.Runtime("Map", "Set", "JSON"), onprogress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onprogress(onprogress)
}

// Onprogress fn
func (p *Props) Onprogress(onprogress string) *Props {
	macro.Rewrite("$_.Set('onprogress', $1)", onprogress)
	p.attrs["onprogress"] = onprogress
	return p
}

// Onratechange fn
func Onratechange(onratechange string) *Props {
	macro.Rewrite("$1().Set('onratechange', $2)", macro.Runtime("Map", "Set", "JSON"), onratechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onratechange(onratechange)
}

// Onratechange fn
func (p *Props) Onratechange(onratechange string) *Props {
	macro.Rewrite("$_.Set('onratechange', $1)", onratechange)
	p.attrs["onratechange"] = onratechange
	return p
}

// Onreset fn
func Onreset(onreset string) *Props {
	macro.Rewrite("$1().Set('onreset', $2)", macro.Runtime("Map", "Set", "JSON"), onreset)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onreset(onreset)
}

// Onreset fn
func (p *Props) Onreset(onreset string) *Props {
	macro.Rewrite("$_.Set('onreset', $1)", onreset)
	p.attrs["onreset"] = onreset
	return p
}

// Onresize fn
func Onresize(onresize string) *Props {
	macro.Rewrite("$1().Set('onresize', $2)", macro.Runtime("Map", "Set", "JSON"), onresize)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onresize(onresize)
}

// Onresize fn
func (p *Props) Onresize(onresize string) *Props {
	macro.Rewrite("$_.Set('onresize', $1)", onresize)
	p.attrs["onresize"] = onresize
	return p
}

// Onreadystatechange fn
func Onreadystatechange(onreadystatechange string) *Props {
	macro.Rewrite("$1().Set('onreadystatechange', $2)", macro.Runtime("Map", "Set", "JSON"), onreadystatechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onreadystatechange(onreadystatechange)
}

// Onreadystatechange fn
func (p *Props) Onreadystatechange(onreadystatechange string) *Props {
	macro.Rewrite("$_.Set('onreadystatechange', $1)", onreadystatechange)
	p.attrs["onreadystatechange"] = onreadystatechange
	return p
}

// Onscroll fn
func Onscroll(onscroll string) *Props {
	macro.Rewrite("$1().Set('onscroll', $2)", macro.Runtime("Map", "Set", "JSON"), onscroll)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onscroll(onscroll)
}

// Onscroll fn
func (p *Props) Onscroll(onscroll string) *Props {
	macro.Rewrite("$_.Set('onscroll', $1)", onscroll)
	p.attrs["onscroll"] = onscroll
	return p
}

// Onseeked fn
func Onseeked(onseeked string) *Props {
	macro.Rewrite("$1().Set('onseeked', $2)", macro.Runtime("Map", "Set", "JSON"), onseeked)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onseeked(onseeked)
}

// Onseeked fn
func (p *Props) Onseeked(onseeked string) *Props {
	macro.Rewrite("$_.Set('onseeked', $1)", onseeked)
	p.attrs["onseeked"] = onseeked
	return p
}

// Onseeking fn
func Onseeking(onseeking string) *Props {
	macro.Rewrite("$1().Set('onseeking', $2)", macro.Runtime("Map", "Set", "JSON"), onseeking)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onseeking(onseeking)
}

// Onseeking fn
func (p *Props) Onseeking(onseeking string) *Props {
	macro.Rewrite("$_.Set('onseeking', $1)", onseeking)
	p.attrs["onseeking"] = onseeking
	return p
}

// Onselect fn
func Onselect(onselect string) *Props {
	macro.Rewrite("$1().Set('onselect', $2)", macro.Runtime("Map", "Set", "JSON"), onselect)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onselect(onselect)
}

// Onselect fn
func (p *Props) Onselect(onselect string) *Props {
	macro.Rewrite("$_.Set('onselect', $1)", onselect)
	p.attrs["onselect"] = onselect
	return p
}

// Onshow fn
func Onshow(onshow string) *Props {
	macro.Rewrite("$1().Set('onshow', $2)", macro.Runtime("Map", "Set", "JSON"), onshow)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onshow(onshow)
}

// Onshow fn
func (p *Props) Onshow(onshow string) *Props {
	macro.Rewrite("$_.Set('onshow', $1)", onshow)
	p.attrs["onshow"] = onshow
	return p
}

// Onstalled fn
func Onstalled(onstalled string) *Props {
	macro.Rewrite("$1().Set('onstalled', $2)", macro.Runtime("Map", "Set", "JSON"), onstalled)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onstalled(onstalled)
}

// Onstalled fn
func (p *Props) Onstalled(onstalled string) *Props {
	macro.Rewrite("$_.Set('onstalled', $1)", onstalled)
	p.attrs["onstalled"] = onstalled
	return p
}

// Onsubmit fn
func Onsubmit(onsubmit string) *Props {
	macro.Rewrite("$1().Set('onsubmit', $2)", macro.Runtime("Map", "Set", "JSON"), onsubmit)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onsubmit(onsubmit)
}

// Onsubmit fn
func (p *Props) Onsubmit(onsubmit string) *Props {
	macro.Rewrite("$_.Set('onsubmit', $1)", onsubmit)
	p.attrs["onsubmit"] = onsubmit
	return p
}

// Onsuspend fn
func Onsuspend(onsuspend string) *Props {
	macro.Rewrite("$1().Set('onsuspend', $2)", macro.Runtime("Map", "Set", "JSON"), onsuspend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onsuspend(onsuspend)
}

// Onsuspend fn
func (p *Props) Onsuspend(onsuspend string) *Props {
	macro.Rewrite("$_.Set('onsuspend', $1)", onsuspend)
	p.attrs["onsuspend"] = onsuspend
	return p
}

// Ontimeupdate fn
func Ontimeupdate(ontimeupdate string) *Props {
	macro.Rewrite("$1().Set('ontimeupdate', $2)", macro.Runtime("Map", "Set", "JSON"), ontimeupdate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ontimeupdate(ontimeupdate)
}

// Ontimeupdate fn
func (p *Props) Ontimeupdate(ontimeupdate string) *Props {
	macro.Rewrite("$_.Set('ontimeupdate', $1)", ontimeupdate)
	p.attrs["ontimeupdate"] = ontimeupdate
	return p
}

// Onvolumechange fn
func Onvolumechange(onvolumechange string) *Props {
	macro.Rewrite("$1().Set('onvolumechange', $2)", macro.Runtime("Map", "Set", "JSON"), onvolumechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onvolumechange(onvolumechange)
}

// Onvolumechange fn
func (p *Props) Onvolumechange(onvolumechange string) *Props {
	macro.Rewrite("$_.Set('onvolumechange', $1)", onvolumechange)
	p.attrs["onvolumechange"] = onvolumechange
	return p
}

// Onwaiting fn
func Onwaiting(onwaiting string) *Props {
	macro.Rewrite("$1().Set('onwaiting', $2)", macro.Runtime("Map", "Set", "JSON"), onwaiting)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onwaiting(onwaiting)
}

// Onwaiting fn
func (p *Props) Onwaiting(onwaiting string) *Props {
	macro.Rewrite("$_.Set('onwaiting', $1)", onwaiting)
	p.attrs["onwaiting"] = onwaiting
	return p
}

// Attr fn
func Attr(key string, value interface{}) *Props {
	macro.Rewrite("$1().Set($2, $3)", macro.Runtime("Map", "Set", "JSON"), key, value)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Attr(key, value)
}

// Attr fn
func (p *Props) Attr(key string, value interface{}) *Props {
	macro.Rewrite("$_.Set($1, $2)", key, value)
	p.attrs[key] = value
	return p
}
