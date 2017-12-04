package hr

import (
	"encoding/json"
	"strings"

	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/vdom"
)

// Hr struct
// js:"hr,omit"
type Hr struct {
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

// New hr element
//
// The hr element represents a paragraph-level thematic break, e.g. a scene change in a story, or a transition to another topic within a section of a reference book.
func New(props *Props, children ...vdom.Child) *Hr {
	macro.Rewrite("$1('hr', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
	if props == nil {
		props = &Props{attrs: map[string]interface{}{}}
	}
	return &Hr{
		attrs:    props.attrs,
		children: children,
	}
}

// Render fn
func (s *Hr) Render() vdom.Node {
	return s
}

func (s *Hr) String() string {
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
		return "<hr " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</hr>"
	}

	return "<hr>" + strings.Join(children, "") + "</hr>"
}

// AriaActivedescendant fn
func AriaActivedescendant(ariaActivedescendant string) *Props {
	macro.Rewrite("$1().Set('aria-activedescendant', $2)", macro.Runtime("Map", "Set", "JSON"), ariaActivedescendant)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaActivedescendant(ariaActivedescendant)
}

// AriaActivedescendant fn
func (p *Props) AriaActivedescendant(ariaActivedescendant string) *Props {
	macro.Rewrite("$_.Set('aria-activedescendant', $1)", ariaActivedescendant)
	p.attrs["aria-activedescendant"] = ariaActivedescendant
	return p
}

// AriaAtomic fn
func AriaAtomic(ariaAtomic string) *Props {
	macro.Rewrite("$1().Set('aria-atomic', $2)", macro.Runtime("Map", "Set", "JSON"), ariaAtomic)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaAtomic(ariaAtomic)
}

// AriaAtomic fn
func (p *Props) AriaAtomic(ariaAtomic string) *Props {
	macro.Rewrite("$_.Set('aria-atomic', $1)", ariaAtomic)
	p.attrs["aria-atomic"] = ariaAtomic
	return p
}

// AriaAutocomplete fn
func AriaAutocomplete(ariaAutocomplete string) *Props {
	macro.Rewrite("$1().Set('aria-autocomplete', $2)", macro.Runtime("Map", "Set", "JSON"), ariaAutocomplete)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaAutocomplete(ariaAutocomplete)
}

// AriaAutocomplete fn
func (p *Props) AriaAutocomplete(ariaAutocomplete string) *Props {
	macro.Rewrite("$_.Set('aria-autocomplete', $1)", ariaAutocomplete)
	p.attrs["aria-autocomplete"] = ariaAutocomplete
	return p
}

// AriaBusy fn
func AriaBusy(ariaBusy string) *Props {
	macro.Rewrite("$1().Set('aria-busy', $2)", macro.Runtime("Map", "Set", "JSON"), ariaBusy)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaBusy(ariaBusy)
}

// AriaBusy fn
func (p *Props) AriaBusy(ariaBusy string) *Props {
	macro.Rewrite("$_.Set('aria-busy', $1)", ariaBusy)
	p.attrs["aria-busy"] = ariaBusy
	return p
}

// AriaChecked fn
func AriaChecked(ariaChecked string) *Props {
	macro.Rewrite("$1().Set('aria-checked', $2)", macro.Runtime("Map", "Set", "JSON"), ariaChecked)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaChecked(ariaChecked)
}

// AriaChecked fn
func (p *Props) AriaChecked(ariaChecked string) *Props {
	macro.Rewrite("$_.Set('aria-checked', $1)", ariaChecked)
	p.attrs["aria-checked"] = ariaChecked
	return p
}

// AriaColcount fn
func AriaColcount(ariaColcount string) *Props {
	macro.Rewrite("$1().Set('aria-colcount', $2)", macro.Runtime("Map", "Set", "JSON"), ariaColcount)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaColcount(ariaColcount)
}

// AriaColcount fn
func (p *Props) AriaColcount(ariaColcount string) *Props {
	macro.Rewrite("$_.Set('aria-colcount', $1)", ariaColcount)
	p.attrs["aria-colcount"] = ariaColcount
	return p
}

// AriaColindex fn
func AriaColindex(ariaColindex string) *Props {
	macro.Rewrite("$1().Set('aria-colindex', $2)", macro.Runtime("Map", "Set", "JSON"), ariaColindex)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaColindex(ariaColindex)
}

// AriaColindex fn
func (p *Props) AriaColindex(ariaColindex string) *Props {
	macro.Rewrite("$_.Set('aria-colindex', $1)", ariaColindex)
	p.attrs["aria-colindex"] = ariaColindex
	return p
}

// AriaColspan fn
func AriaColspan(ariaColspan string) *Props {
	macro.Rewrite("$1().Set('aria-colspan', $2)", macro.Runtime("Map", "Set", "JSON"), ariaColspan)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaColspan(ariaColspan)
}

// AriaColspan fn
func (p *Props) AriaColspan(ariaColspan string) *Props {
	macro.Rewrite("$_.Set('aria-colspan', $1)", ariaColspan)
	p.attrs["aria-colspan"] = ariaColspan
	return p
}

// AriaControls fn
func AriaControls(ariaControls string) *Props {
	macro.Rewrite("$1().Set('aria-controls', $2)", macro.Runtime("Map", "Set", "JSON"), ariaControls)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaControls(ariaControls)
}

// AriaControls fn
func (p *Props) AriaControls(ariaControls string) *Props {
	macro.Rewrite("$_.Set('aria-controls', $1)", ariaControls)
	p.attrs["aria-controls"] = ariaControls
	return p
}

// AriaCurrent fn
func AriaCurrent(ariaCurrent string) *Props {
	macro.Rewrite("$1().Set('aria-current', $2)", macro.Runtime("Map", "Set", "JSON"), ariaCurrent)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaCurrent(ariaCurrent)
}

// AriaCurrent fn
func (p *Props) AriaCurrent(ariaCurrent string) *Props {
	macro.Rewrite("$_.Set('aria-current', $1)", ariaCurrent)
	p.attrs["aria-current"] = ariaCurrent
	return p
}

// AriaDescribedat fn
func AriaDescribedat(ariaDescribedat string) *Props {
	macro.Rewrite("$1().Set('aria-describedat', $2)", macro.Runtime("Map", "Set", "JSON"), ariaDescribedat)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaDescribedat(ariaDescribedat)
}

// AriaDescribedat fn
func (p *Props) AriaDescribedat(ariaDescribedat string) *Props {
	macro.Rewrite("$_.Set('aria-describedat', $1)", ariaDescribedat)
	p.attrs["aria-describedat"] = ariaDescribedat
	return p
}

// AriaDescribedby fn
func AriaDescribedby(ariaDescribedby string) *Props {
	macro.Rewrite("$1().Set('aria-describedby', $2)", macro.Runtime("Map", "Set", "JSON"), ariaDescribedby)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaDescribedby(ariaDescribedby)
}

// AriaDescribedby fn
func (p *Props) AriaDescribedby(ariaDescribedby string) *Props {
	macro.Rewrite("$_.Set('aria-describedby', $1)", ariaDescribedby)
	p.attrs["aria-describedby"] = ariaDescribedby
	return p
}

// AriaDisabled fn
func AriaDisabled(ariaDisabled string) *Props {
	macro.Rewrite("$1().Set('aria-disabled', $2)", macro.Runtime("Map", "Set", "JSON"), ariaDisabled)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaDisabled(ariaDisabled)
}

// AriaDisabled fn
func (p *Props) AriaDisabled(ariaDisabled string) *Props {
	macro.Rewrite("$_.Set('aria-disabled', $1)", ariaDisabled)
	p.attrs["aria-disabled"] = ariaDisabled
	return p
}

// AriaDropeffect fn
func AriaDropeffect(ariaDropeffect string) *Props {
	macro.Rewrite("$1().Set('aria-dropeffect', $2)", macro.Runtime("Map", "Set", "JSON"), ariaDropeffect)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaDropeffect(ariaDropeffect)
}

// AriaDropeffect fn
func (p *Props) AriaDropeffect(ariaDropeffect string) *Props {
	macro.Rewrite("$_.Set('aria-dropeffect', $1)", ariaDropeffect)
	p.attrs["aria-dropeffect"] = ariaDropeffect
	return p
}

// AriaErrormessage fn
func AriaErrormessage(ariaErrormessage string) *Props {
	macro.Rewrite("$1().Set('aria-errormessage', $2)", macro.Runtime("Map", "Set", "JSON"), ariaErrormessage)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaErrormessage(ariaErrormessage)
}

// AriaErrormessage fn
func (p *Props) AriaErrormessage(ariaErrormessage string) *Props {
	macro.Rewrite("$_.Set('aria-errormessage', $1)", ariaErrormessage)
	p.attrs["aria-errormessage"] = ariaErrormessage
	return p
}

// AriaExpanded fn
func AriaExpanded(ariaExpanded string) *Props {
	macro.Rewrite("$1().Set('aria-expanded', $2)", macro.Runtime("Map", "Set", "JSON"), ariaExpanded)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaExpanded(ariaExpanded)
}

// AriaExpanded fn
func (p *Props) AriaExpanded(ariaExpanded string) *Props {
	macro.Rewrite("$_.Set('aria-expanded', $1)", ariaExpanded)
	p.attrs["aria-expanded"] = ariaExpanded
	return p
}

// AriaFlowto fn
func AriaFlowto(ariaFlowto string) *Props {
	macro.Rewrite("$1().Set('aria-flowto', $2)", macro.Runtime("Map", "Set", "JSON"), ariaFlowto)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaFlowto(ariaFlowto)
}

// AriaFlowto fn
func (p *Props) AriaFlowto(ariaFlowto string) *Props {
	macro.Rewrite("$_.Set('aria-flowto', $1)", ariaFlowto)
	p.attrs["aria-flowto"] = ariaFlowto
	return p
}

// AriaGrabbed fn
func AriaGrabbed(ariaGrabbed string) *Props {
	macro.Rewrite("$1().Set('aria-grabbed', $2)", macro.Runtime("Map", "Set", "JSON"), ariaGrabbed)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaGrabbed(ariaGrabbed)
}

// AriaGrabbed fn
func (p *Props) AriaGrabbed(ariaGrabbed string) *Props {
	macro.Rewrite("$_.Set('aria-grabbed', $1)", ariaGrabbed)
	p.attrs["aria-grabbed"] = ariaGrabbed
	return p
}

// AriaHaspopup fn
func AriaHaspopup(ariaHaspopup string) *Props {
	macro.Rewrite("$1().Set('aria-haspopup', $2)", macro.Runtime("Map", "Set", "JSON"), ariaHaspopup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaHaspopup(ariaHaspopup)
}

// AriaHaspopup fn
func (p *Props) AriaHaspopup(ariaHaspopup string) *Props {
	macro.Rewrite("$_.Set('aria-haspopup', $1)", ariaHaspopup)
	p.attrs["aria-haspopup"] = ariaHaspopup
	return p
}

// AriaHidden fn
func AriaHidden(ariaHidden string) *Props {
	macro.Rewrite("$1().Set('aria-hidden', $2)", macro.Runtime("Map", "Set", "JSON"), ariaHidden)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaHidden(ariaHidden)
}

// AriaHidden fn
func (p *Props) AriaHidden(ariaHidden string) *Props {
	macro.Rewrite("$_.Set('aria-hidden', $1)", ariaHidden)
	p.attrs["aria-hidden"] = ariaHidden
	return p
}

// AriaInvalid fn
func AriaInvalid(ariaInvalid string) *Props {
	macro.Rewrite("$1().Set('aria-invalid', $2)", macro.Runtime("Map", "Set", "JSON"), ariaInvalid)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaInvalid(ariaInvalid)
}

// AriaInvalid fn
func (p *Props) AriaInvalid(ariaInvalid string) *Props {
	macro.Rewrite("$_.Set('aria-invalid', $1)", ariaInvalid)
	p.attrs["aria-invalid"] = ariaInvalid
	return p
}

// AriaKbdshortcuts fn
func AriaKbdshortcuts(ariaKbdshortcuts string) *Props {
	macro.Rewrite("$1().Set('aria-kbdshortcuts', $2)", macro.Runtime("Map", "Set", "JSON"), ariaKbdshortcuts)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaKbdshortcuts(ariaKbdshortcuts)
}

// AriaKbdshortcuts fn
func (p *Props) AriaKbdshortcuts(ariaKbdshortcuts string) *Props {
	macro.Rewrite("$_.Set('aria-kbdshortcuts', $1)", ariaKbdshortcuts)
	p.attrs["aria-kbdshortcuts"] = ariaKbdshortcuts
	return p
}

// AriaLabel fn
func AriaLabel(ariaLabel string) *Props {
	macro.Rewrite("$1().Set('aria-label', $2)", macro.Runtime("Map", "Set", "JSON"), ariaLabel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaLabel(ariaLabel)
}

// AriaLabel fn
func (p *Props) AriaLabel(ariaLabel string) *Props {
	macro.Rewrite("$_.Set('aria-label', $1)", ariaLabel)
	p.attrs["aria-label"] = ariaLabel
	return p
}

// AriaLabelledby fn
func AriaLabelledby(ariaLabelledby string) *Props {
	macro.Rewrite("$1().Set('aria-labelledby', $2)", macro.Runtime("Map", "Set", "JSON"), ariaLabelledby)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaLabelledby(ariaLabelledby)
}

// AriaLabelledby fn
func (p *Props) AriaLabelledby(ariaLabelledby string) *Props {
	macro.Rewrite("$_.Set('aria-labelledby', $1)", ariaLabelledby)
	p.attrs["aria-labelledby"] = ariaLabelledby
	return p
}

// AriaLevel fn
func AriaLevel(ariaLevel string) *Props {
	macro.Rewrite("$1().Set('aria-level', $2)", macro.Runtime("Map", "Set", "JSON"), ariaLevel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaLevel(ariaLevel)
}

// AriaLevel fn
func (p *Props) AriaLevel(ariaLevel string) *Props {
	macro.Rewrite("$_.Set('aria-level', $1)", ariaLevel)
	p.attrs["aria-level"] = ariaLevel
	return p
}

// AriaLive fn
func AriaLive(ariaLive string) *Props {
	macro.Rewrite("$1().Set('aria-live', $2)", macro.Runtime("Map", "Set", "JSON"), ariaLive)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaLive(ariaLive)
}

// AriaLive fn
func (p *Props) AriaLive(ariaLive string) *Props {
	macro.Rewrite("$_.Set('aria-live', $1)", ariaLive)
	p.attrs["aria-live"] = ariaLive
	return p
}

// AriaModal fn
func AriaModal(ariaModal string) *Props {
	macro.Rewrite("$1().Set('aria-modal', $2)", macro.Runtime("Map", "Set", "JSON"), ariaModal)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaModal(ariaModal)
}

// AriaModal fn
func (p *Props) AriaModal(ariaModal string) *Props {
	macro.Rewrite("$_.Set('aria-modal', $1)", ariaModal)
	p.attrs["aria-modal"] = ariaModal
	return p
}

// AriaMultiline fn
func AriaMultiline(ariaMultiline string) *Props {
	macro.Rewrite("$1().Set('aria-multiline', $2)", macro.Runtime("Map", "Set", "JSON"), ariaMultiline)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaMultiline(ariaMultiline)
}

// AriaMultiline fn
func (p *Props) AriaMultiline(ariaMultiline string) *Props {
	macro.Rewrite("$_.Set('aria-multiline', $1)", ariaMultiline)
	p.attrs["aria-multiline"] = ariaMultiline
	return p
}

// AriaMultiselectable fn
func AriaMultiselectable(ariaMultiselectable string) *Props {
	macro.Rewrite("$1().Set('aria-multiselectable', $2)", macro.Runtime("Map", "Set", "JSON"), ariaMultiselectable)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaMultiselectable(ariaMultiselectable)
}

// AriaMultiselectable fn
func (p *Props) AriaMultiselectable(ariaMultiselectable string) *Props {
	macro.Rewrite("$_.Set('aria-multiselectable', $1)", ariaMultiselectable)
	p.attrs["aria-multiselectable"] = ariaMultiselectable
	return p
}

// AriaOrientation fn
func AriaOrientation(ariaOrientation string) *Props {
	macro.Rewrite("$1().Set('aria-orientation', $2)", macro.Runtime("Map", "Set", "JSON"), ariaOrientation)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaOrientation(ariaOrientation)
}

// AriaOrientation fn
func (p *Props) AriaOrientation(ariaOrientation string) *Props {
	macro.Rewrite("$_.Set('aria-orientation', $1)", ariaOrientation)
	p.attrs["aria-orientation"] = ariaOrientation
	return p
}

// AriaOwns fn
func AriaOwns(ariaOwns string) *Props {
	macro.Rewrite("$1().Set('aria-owns', $2)", macro.Runtime("Map", "Set", "JSON"), ariaOwns)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaOwns(ariaOwns)
}

// AriaOwns fn
func (p *Props) AriaOwns(ariaOwns string) *Props {
	macro.Rewrite("$_.Set('aria-owns', $1)", ariaOwns)
	p.attrs["aria-owns"] = ariaOwns
	return p
}

// AriaPlaceholder fn
func AriaPlaceholder(ariaPlaceholder string) *Props {
	macro.Rewrite("$1().Set('aria-placeholder', $2)", macro.Runtime("Map", "Set", "JSON"), ariaPlaceholder)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaPlaceholder(ariaPlaceholder)
}

// AriaPlaceholder fn
func (p *Props) AriaPlaceholder(ariaPlaceholder string) *Props {
	macro.Rewrite("$_.Set('aria-placeholder', $1)", ariaPlaceholder)
	p.attrs["aria-placeholder"] = ariaPlaceholder
	return p
}

// AriaPosinset fn
func AriaPosinset(ariaPosinset string) *Props {
	macro.Rewrite("$1().Set('aria-posinset', $2)", macro.Runtime("Map", "Set", "JSON"), ariaPosinset)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaPosinset(ariaPosinset)
}

// AriaPosinset fn
func (p *Props) AriaPosinset(ariaPosinset string) *Props {
	macro.Rewrite("$_.Set('aria-posinset', $1)", ariaPosinset)
	p.attrs["aria-posinset"] = ariaPosinset
	return p
}

// AriaPressed fn
func AriaPressed(ariaPressed string) *Props {
	macro.Rewrite("$1().Set('aria-pressed', $2)", macro.Runtime("Map", "Set", "JSON"), ariaPressed)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaPressed(ariaPressed)
}

// AriaPressed fn
func (p *Props) AriaPressed(ariaPressed string) *Props {
	macro.Rewrite("$_.Set('aria-pressed', $1)", ariaPressed)
	p.attrs["aria-pressed"] = ariaPressed
	return p
}

// AriaReadonly fn
func AriaReadonly(ariaReadonly string) *Props {
	macro.Rewrite("$1().Set('aria-readonly', $2)", macro.Runtime("Map", "Set", "JSON"), ariaReadonly)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaReadonly(ariaReadonly)
}

// AriaReadonly fn
func (p *Props) AriaReadonly(ariaReadonly string) *Props {
	macro.Rewrite("$_.Set('aria-readonly', $1)", ariaReadonly)
	p.attrs["aria-readonly"] = ariaReadonly
	return p
}

// AriaRelevant fn
func AriaRelevant(ariaRelevant string) *Props {
	macro.Rewrite("$1().Set('aria-relevant', $2)", macro.Runtime("Map", "Set", "JSON"), ariaRelevant)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaRelevant(ariaRelevant)
}

// AriaRelevant fn
func (p *Props) AriaRelevant(ariaRelevant string) *Props {
	macro.Rewrite("$_.Set('aria-relevant', $1)", ariaRelevant)
	p.attrs["aria-relevant"] = ariaRelevant
	return p
}

// AriaRequired fn
func AriaRequired(ariaRequired string) *Props {
	macro.Rewrite("$1().Set('aria-required', $2)", macro.Runtime("Map", "Set", "JSON"), ariaRequired)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaRequired(ariaRequired)
}

// AriaRequired fn
func (p *Props) AriaRequired(ariaRequired string) *Props {
	macro.Rewrite("$_.Set('aria-required', $1)", ariaRequired)
	p.attrs["aria-required"] = ariaRequired
	return p
}

// AriaRoledescription fn
func AriaRoledescription(ariaRoledescription string) *Props {
	macro.Rewrite("$1().Set('aria-roledescription', $2)", macro.Runtime("Map", "Set", "JSON"), ariaRoledescription)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaRoledescription(ariaRoledescription)
}

// AriaRoledescription fn
func (p *Props) AriaRoledescription(ariaRoledescription string) *Props {
	macro.Rewrite("$_.Set('aria-roledescription', $1)", ariaRoledescription)
	p.attrs["aria-roledescription"] = ariaRoledescription
	return p
}

// AriaRowcount fn
func AriaRowcount(ariaRowcount string) *Props {
	macro.Rewrite("$1().Set('aria-rowcount', $2)", macro.Runtime("Map", "Set", "JSON"), ariaRowcount)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaRowcount(ariaRowcount)
}

// AriaRowcount fn
func (p *Props) AriaRowcount(ariaRowcount string) *Props {
	macro.Rewrite("$_.Set('aria-rowcount', $1)", ariaRowcount)
	p.attrs["aria-rowcount"] = ariaRowcount
	return p
}

// AriaRowindex fn
func AriaRowindex(ariaRowindex string) *Props {
	macro.Rewrite("$1().Set('aria-rowindex', $2)", macro.Runtime("Map", "Set", "JSON"), ariaRowindex)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaRowindex(ariaRowindex)
}

// AriaRowindex fn
func (p *Props) AriaRowindex(ariaRowindex string) *Props {
	macro.Rewrite("$_.Set('aria-rowindex', $1)", ariaRowindex)
	p.attrs["aria-rowindex"] = ariaRowindex
	return p
}

// AriaRowspan fn
func AriaRowspan(ariaRowspan string) *Props {
	macro.Rewrite("$1().Set('aria-rowspan', $2)", macro.Runtime("Map", "Set", "JSON"), ariaRowspan)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaRowspan(ariaRowspan)
}

// AriaRowspan fn
func (p *Props) AriaRowspan(ariaRowspan string) *Props {
	macro.Rewrite("$_.Set('aria-rowspan', $1)", ariaRowspan)
	p.attrs["aria-rowspan"] = ariaRowspan
	return p
}

// AriaSelected fn
func AriaSelected(ariaSelected string) *Props {
	macro.Rewrite("$1().Set('aria-selected', $2)", macro.Runtime("Map", "Set", "JSON"), ariaSelected)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaSelected(ariaSelected)
}

// AriaSelected fn
func (p *Props) AriaSelected(ariaSelected string) *Props {
	macro.Rewrite("$_.Set('aria-selected', $1)", ariaSelected)
	p.attrs["aria-selected"] = ariaSelected
	return p
}

// AriaSetsize fn
func AriaSetsize(ariaSetsize string) *Props {
	macro.Rewrite("$1().Set('aria-setsize', $2)", macro.Runtime("Map", "Set", "JSON"), ariaSetsize)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaSetsize(ariaSetsize)
}

// AriaSetsize fn
func (p *Props) AriaSetsize(ariaSetsize string) *Props {
	macro.Rewrite("$_.Set('aria-setsize', $1)", ariaSetsize)
	p.attrs["aria-setsize"] = ariaSetsize
	return p
}

// AriaSort fn
func AriaSort(ariaSort string) *Props {
	macro.Rewrite("$1().Set('aria-sort', $2)", macro.Runtime("Map", "Set", "JSON"), ariaSort)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaSort(ariaSort)
}

// AriaSort fn
func (p *Props) AriaSort(ariaSort string) *Props {
	macro.Rewrite("$_.Set('aria-sort', $1)", ariaSort)
	p.attrs["aria-sort"] = ariaSort
	return p
}

// AriaValuemax fn
func AriaValuemax(ariaValuemax string) *Props {
	macro.Rewrite("$1().Set('aria-valuemax', $2)", macro.Runtime("Map", "Set", "JSON"), ariaValuemax)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaValuemax(ariaValuemax)
}

// AriaValuemax fn
func (p *Props) AriaValuemax(ariaValuemax string) *Props {
	macro.Rewrite("$_.Set('aria-valuemax', $1)", ariaValuemax)
	p.attrs["aria-valuemax"] = ariaValuemax
	return p
}

// AriaValuemin fn
func AriaValuemin(ariaValuemin string) *Props {
	macro.Rewrite("$1().Set('aria-valuemin', $2)", macro.Runtime("Map", "Set", "JSON"), ariaValuemin)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaValuemin(ariaValuemin)
}

// AriaValuemin fn
func (p *Props) AriaValuemin(ariaValuemin string) *Props {
	macro.Rewrite("$_.Set('aria-valuemin', $1)", ariaValuemin)
	p.attrs["aria-valuemin"] = ariaValuemin
	return p
}

// AriaValuenow fn
func AriaValuenow(ariaValuenow string) *Props {
	macro.Rewrite("$1().Set('aria-valuenow', $2)", macro.Runtime("Map", "Set", "JSON"), ariaValuenow)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaValuenow(ariaValuenow)
}

// AriaValuenow fn
func (p *Props) AriaValuenow(ariaValuenow string) *Props {
	macro.Rewrite("$_.Set('aria-valuenow', $1)", ariaValuenow)
	p.attrs["aria-valuenow"] = ariaValuenow
	return p
}

// AriaValuetext fn
func AriaValuetext(ariaValuetext string) *Props {
	macro.Rewrite("$1().Set('aria-valuetext', $2)", macro.Runtime("Map", "Set", "JSON"), ariaValuetext)
	p := &Props{attrs: map[string]interface{}{}}
	return p.AriaValuetext(ariaValuetext)
}

// AriaValuetext fn
func (p *Props) AriaValuetext(ariaValuetext string) *Props {
	macro.Rewrite("$_.Set('aria-valuetext', $1)", ariaValuetext)
	p.attrs["aria-valuetext"] = ariaValuetext
	return p
}

// Accesskey fn
func Accesskey(accesskey string) *Props {
	macro.Rewrite("$1().Set('accesskey', $2)", macro.Runtime("Map", "Set", "JSON"), accesskey)
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
	macro.Rewrite("$1().Set('class', $2)", macro.Runtime("Map", "Set", "JSON"), class)
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
	macro.Rewrite("$1().Set('contenteditable', $2)", macro.Runtime("Map", "Set", "JSON"), contenteditable)
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
	macro.Rewrite("$1().Set('contextmenu', $2)", macro.Runtime("Map", "Set", "JSON"), contextmenu)
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
	macro.Rewrite("$1().Set('dir', $2)", macro.Runtime("Map", "Set", "JSON"), dir)
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
	macro.Rewrite("$1().Set('draggable', $2)", macro.Runtime("Map", "Set", "JSON"), draggable)
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
	macro.Rewrite("$1().Set('dropzone', $2)", macro.Runtime("Map", "Set", "JSON"), dropzone)
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
	macro.Rewrite("$1().Set('hidden', $2)", macro.Runtime("Map", "Set", "JSON"), hidden)
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
	macro.Rewrite("$1().Set('id', $2)", macro.Runtime("Map", "Set", "JSON"), id)
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
	macro.Rewrite("$1().Set('itemid', $2)", macro.Runtime("Map", "Set", "JSON"), itemid)
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
	macro.Rewrite("$1().Set('itemprop', $2)", macro.Runtime("Map", "Set", "JSON"), itemprop)
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
	macro.Rewrite("$1().Set('itemref', $2)", macro.Runtime("Map", "Set", "JSON"), itemref)
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
	macro.Rewrite("$1().Set('itemscope', $2)", macro.Runtime("Map", "Set", "JSON"), itemscope)
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
	macro.Rewrite("$1().Set('itemtype', $2)", macro.Runtime("Map", "Set", "JSON"), itemtype)
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
	macro.Rewrite("$1().Set('lang', $2)", macro.Runtime("Map", "Set", "JSON"), lang)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Lang(lang)
}

// Lang fn
func (p *Props) Lang(lang string) *Props {
	macro.Rewrite("$_.Set('lang', $1)", lang)
	p.attrs["lang"] = lang
	return p
}

// Role fn
func Role(role string) *Props {
	macro.Rewrite("$1().Set('role', $2)", macro.Runtime("Map", "Set", "JSON"), role)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Role(role)
}

// Role fn
func (p *Props) Role(role string) *Props {
	macro.Rewrite("$_.Set('role', $1)", role)
	p.attrs["role"] = role
	return p
}

// Spellcheck fn
func Spellcheck(spellcheck string) *Props {
	macro.Rewrite("$1().Set('spellcheck', $2)", macro.Runtime("Map", "Set", "JSON"), spellcheck)
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
	macro.Rewrite("$1().Set('style', $2)", macro.Runtime("Map", "Set", "JSON"), style)
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
	macro.Rewrite("$1().Set('tabindex', $2)", macro.Runtime("Map", "Set", "JSON"), tabindex)
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
	macro.Rewrite("$1().Set('title', $2)", macro.Runtime("Map", "Set", "JSON"), title)
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
	macro.Rewrite("$1().Set('translate', $2)", macro.Runtime("Map", "Set", "JSON"), translate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Translate(translate)
}

// Translate fn
func (p *Props) Translate(translate string) *Props {
	macro.Rewrite("$_.Set('translate', $1)", translate)
	p.attrs["translate"] = translate
	return p
}

// Onabort fn
func Onabort(onabort func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onabort', $2)", macro.Runtime("Map", "Set", "JSON"), onabort)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onabort(onabort)
}

// Onabort fn
func (p *Props) Onabort(onabort func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onabort', $1)", onabort)
	p.attrs["onabort"] = onabort
	return p
}

// Onblur fn
func Onblur(onblur func(e window.FocusEvent)) *Props {
	macro.Rewrite("$1().Set('onblur', $2)", macro.Runtime("Map", "Set", "JSON"), onblur)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onblur(onblur)
}

// Onblur fn
func (p *Props) Onblur(onblur func(e window.FocusEvent)) *Props {
	macro.Rewrite("$_.Set('onblur', $1)", onblur)
	p.attrs["onblur"] = onblur
	return p
}

// Oncanplay fn
func Oncanplay(oncanplay func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oncanplay', $2)", macro.Runtime("Map", "Set", "JSON"), oncanplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncanplay(oncanplay)
}

// Oncanplay fn
func (p *Props) Oncanplay(oncanplay func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oncanplay', $1)", oncanplay)
	p.attrs["oncanplay"] = oncanplay
	return p
}

// Oncanplaythrough fn
func Oncanplaythrough(oncanplaythrough func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oncanplaythrough', $2)", macro.Runtime("Map", "Set", "JSON"), oncanplaythrough)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncanplaythrough(oncanplaythrough)
}

// Oncanplaythrough fn
func (p *Props) Oncanplaythrough(oncanplaythrough func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oncanplaythrough', $1)", oncanplaythrough)
	p.attrs["oncanplaythrough"] = oncanplaythrough
	return p
}

// Onchange fn
func Onchange(onchange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onchange', $2)", macro.Runtime("Map", "Set", "JSON"), onchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onchange(onchange)
}

// Onchange fn
func (p *Props) Onchange(onchange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onchange', $1)", onchange)
	p.attrs["onchange"] = onchange
	return p
}

// Onclick fn
func Onclick(onclick func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onclick', $2)", macro.Runtime("Map", "Set", "JSON"), onclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onclick(onclick)
}

// Onclick fn
func (p *Props) Onclick(onclick func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onclick', $1)", onclick)
	p.attrs["onclick"] = onclick
	return p
}

// Oncontextmenu fn
func Oncontextmenu(oncontextmenu func(e window.PointerEvent)) *Props {
	macro.Rewrite("$1().Set('oncontextmenu', $2)", macro.Runtime("Map", "Set", "JSON"), oncontextmenu)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncontextmenu(oncontextmenu)
}

// Oncontextmenu fn
func (p *Props) Oncontextmenu(oncontextmenu func(e window.PointerEvent)) *Props {
	macro.Rewrite("$_.Set('oncontextmenu', $1)", oncontextmenu)
	p.attrs["oncontextmenu"] = oncontextmenu
	return p
}

// Ondblclick fn
func Ondblclick(ondblclick func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('ondblclick', $2)", macro.Runtime("Map", "Set", "JSON"), ondblclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondblclick(ondblclick)
}

// Ondblclick fn
func (p *Props) Ondblclick(ondblclick func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('ondblclick', $1)", ondblclick)
	p.attrs["ondblclick"] = ondblclick
	return p
}

// Ondrag fn
func Ondrag(ondrag func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondrag', $2)", macro.Runtime("Map", "Set", "JSON"), ondrag)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondrag(ondrag)
}

// Ondrag fn
func (p *Props) Ondrag(ondrag func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondrag', $1)", ondrag)
	p.attrs["ondrag"] = ondrag
	return p
}

// Ondragend fn
func Ondragend(ondragend func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragend', $2)", macro.Runtime("Map", "Set", "JSON"), ondragend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragend(ondragend)
}

// Ondragend fn
func (p *Props) Ondragend(ondragend func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragend', $1)", ondragend)
	p.attrs["ondragend"] = ondragend
	return p
}

// Ondragenter fn
func Ondragenter(ondragenter func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragenter', $2)", macro.Runtime("Map", "Set", "JSON"), ondragenter)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragenter(ondragenter)
}

// Ondragenter fn
func (p *Props) Ondragenter(ondragenter func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragenter', $1)", ondragenter)
	p.attrs["ondragenter"] = ondragenter
	return p
}

// Ondragleave fn
func Ondragleave(ondragleave func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragleave', $2)", macro.Runtime("Map", "Set", "JSON"), ondragleave)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragleave(ondragleave)
}

// Ondragleave fn
func (p *Props) Ondragleave(ondragleave func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragleave', $1)", ondragleave)
	p.attrs["ondragleave"] = ondragleave
	return p
}

// Ondragover fn
func Ondragover(ondragover func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragover', $2)", macro.Runtime("Map", "Set", "JSON"), ondragover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragover(ondragover)
}

// Ondragover fn
func (p *Props) Ondragover(ondragover func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragover', $1)", ondragover)
	p.attrs["ondragover"] = ondragover
	return p
}

// Ondragstart fn
func Ondragstart(ondragstart func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragstart', $2)", macro.Runtime("Map", "Set", "JSON"), ondragstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragstart(ondragstart)
}

// Ondragstart fn
func (p *Props) Ondragstart(ondragstart func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragstart', $1)", ondragstart)
	p.attrs["ondragstart"] = ondragstart
	return p
}

// Ondrop fn
func Ondrop(ondrop func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondrop', $2)", macro.Runtime("Map", "Set", "JSON"), ondrop)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondrop(ondrop)
}

// Ondrop fn
func (p *Props) Ondrop(ondrop func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondrop', $1)", ondrop)
	p.attrs["ondrop"] = ondrop
	return p
}

// Ondurationchange fn
func Ondurationchange(ondurationchange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('ondurationchange', $2)", macro.Runtime("Map", "Set", "JSON"), ondurationchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondurationchange(ondurationchange)
}

// Ondurationchange fn
func (p *Props) Ondurationchange(ondurationchange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('ondurationchange', $1)", ondurationchange)
	p.attrs["ondurationchange"] = ondurationchange
	return p
}

// Onemptied fn
func Onemptied(onemptied func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onemptied', $2)", macro.Runtime("Map", "Set", "JSON"), onemptied)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onemptied(onemptied)
}

// Onemptied fn
func (p *Props) Onemptied(onemptied func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onemptied', $1)", onemptied)
	p.attrs["onemptied"] = onemptied
	return p
}

// Onended fn
func Onended(onended func(e window.MediaStreamErrorEvent)) *Props {
	macro.Rewrite("$1().Set('onended', $2)", macro.Runtime("Map", "Set", "JSON"), onended)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onended(onended)
}

// Onended fn
func (p *Props) Onended(onended func(e window.MediaStreamErrorEvent)) *Props {
	macro.Rewrite("$_.Set('onended', $1)", onended)
	p.attrs["onended"] = onended
	return p
}

// Onerror fn
func Onerror(onerror func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onerror', $2)", macro.Runtime("Map", "Set", "JSON"), onerror)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onerror(onerror)
}

// Onerror fn
func (p *Props) Onerror(onerror func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onerror', $1)", onerror)
	p.attrs["onerror"] = onerror
	return p
}

// Onfocus fn
func Onfocus(onfocus func(e window.FocusEvent)) *Props {
	macro.Rewrite("$1().Set('onfocus', $2)", macro.Runtime("Map", "Set", "JSON"), onfocus)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onfocus(onfocus)
}

// Onfocus fn
func (p *Props) Onfocus(onfocus func(e window.FocusEvent)) *Props {
	macro.Rewrite("$_.Set('onfocus', $1)", onfocus)
	p.attrs["onfocus"] = onfocus
	return p
}

// Oninput fn
func Oninput(oninput func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oninput', $2)", macro.Runtime("Map", "Set", "JSON"), oninput)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oninput(oninput)
}

// Oninput fn
func (p *Props) Oninput(oninput func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oninput', $1)", oninput)
	p.attrs["oninput"] = oninput
	return p
}

// Oninvalid fn
func Oninvalid(oninvalid func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oninvalid', $2)", macro.Runtime("Map", "Set", "JSON"), oninvalid)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oninvalid(oninvalid)
}

// Oninvalid fn
func (p *Props) Oninvalid(oninvalid func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oninvalid', $1)", oninvalid)
	p.attrs["oninvalid"] = oninvalid
	return p
}

// Onkeydown fn
func Onkeydown(onkeydown func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('onkeydown', $2)", macro.Runtime("Map", "Set", "JSON"), onkeydown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeydown(onkeydown)
}

// Onkeydown fn
func (p *Props) Onkeydown(onkeydown func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('onkeydown', $1)", onkeydown)
	p.attrs["onkeydown"] = onkeydown
	return p
}

// Onkeypress fn
func Onkeypress(onkeypress func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('onkeypress', $2)", macro.Runtime("Map", "Set", "JSON"), onkeypress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeypress(onkeypress)
}

// Onkeypress fn
func (p *Props) Onkeypress(onkeypress func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('onkeypress', $1)", onkeypress)
	p.attrs["onkeypress"] = onkeypress
	return p
}

// Onkeyup fn
func Onkeyup(onkeyup func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('onkeyup', $2)", macro.Runtime("Map", "Set", "JSON"), onkeyup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeyup(onkeyup)
}

// Onkeyup fn
func (p *Props) Onkeyup(onkeyup func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('onkeyup', $1)", onkeyup)
	p.attrs["onkeyup"] = onkeyup
	return p
}

// Onload fn
func Onload(onload func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onload', $2)", macro.Runtime("Map", "Set", "JSON"), onload)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onload(onload)
}

// Onload fn
func (p *Props) Onload(onload func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onload', $1)", onload)
	p.attrs["onload"] = onload
	return p
}

// Onloadeddata fn
func Onloadeddata(onloadeddata func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onloadeddata', $2)", macro.Runtime("Map", "Set", "JSON"), onloadeddata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadeddata(onloadeddata)
}

// Onloadeddata fn
func (p *Props) Onloadeddata(onloadeddata func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onloadeddata', $1)", onloadeddata)
	p.attrs["onloadeddata"] = onloadeddata
	return p
}

// Onloadedmetadata fn
func Onloadedmetadata(onloadedmetadata func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onloadedmetadata', $2)", macro.Runtime("Map", "Set", "JSON"), onloadedmetadata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadedmetadata(onloadedmetadata)
}

// Onloadedmetadata fn
func (p *Props) Onloadedmetadata(onloadedmetadata func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onloadedmetadata', $1)", onloadedmetadata)
	p.attrs["onloadedmetadata"] = onloadedmetadata
	return p
}

// Onloadstart fn
func Onloadstart(onloadstart func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onloadstart', $2)", macro.Runtime("Map", "Set", "JSON"), onloadstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadstart(onloadstart)
}

// Onloadstart fn
func (p *Props) Onloadstart(onloadstart func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onloadstart', $1)", onloadstart)
	p.attrs["onloadstart"] = onloadstart
	return p
}

// Onmousedown fn
func Onmousedown(onmousedown func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmousedown', $2)", macro.Runtime("Map", "Set", "JSON"), onmousedown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousedown(onmousedown)
}

// Onmousedown fn
func (p *Props) Onmousedown(onmousedown func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmousedown', $1)", onmousedown)
	p.attrs["onmousedown"] = onmousedown
	return p
}

// Onmousemove fn
func Onmousemove(onmousemove func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmousemove', $2)", macro.Runtime("Map", "Set", "JSON"), onmousemove)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousemove(onmousemove)
}

// Onmousemove fn
func (p *Props) Onmousemove(onmousemove func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmousemove', $1)", onmousemove)
	p.attrs["onmousemove"] = onmousemove
	return p
}

// Onmouseout fn
func Onmouseout(onmouseout func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmouseout', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseout)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseout(onmouseout)
}

// Onmouseout fn
func (p *Props) Onmouseout(onmouseout func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmouseout', $1)", onmouseout)
	p.attrs["onmouseout"] = onmouseout
	return p
}

// Onmouseover fn
func Onmouseover(onmouseover func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmouseover', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseover(onmouseover)
}

// Onmouseover fn
func (p *Props) Onmouseover(onmouseover func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmouseover', $1)", onmouseover)
	p.attrs["onmouseover"] = onmouseover
	return p
}

// Onmouseup fn
func Onmouseup(onmouseup func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmouseup', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseup(onmouseup)
}

// Onmouseup fn
func (p *Props) Onmouseup(onmouseup func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmouseup', $1)", onmouseup)
	p.attrs["onmouseup"] = onmouseup
	return p
}

// Onmousewheel fn
func Onmousewheel(onmousewheel func(e window.WheelEvent)) *Props {
	macro.Rewrite("$1().Set('onmousewheel', $2)", macro.Runtime("Map", "Set", "JSON"), onmousewheel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousewheel(onmousewheel)
}

// Onmousewheel fn
func (p *Props) Onmousewheel(onmousewheel func(e window.WheelEvent)) *Props {
	macro.Rewrite("$_.Set('onmousewheel', $1)", onmousewheel)
	p.attrs["onmousewheel"] = onmousewheel
	return p
}

// Onpause fn
func Onpause(onpause func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onpause', $2)", macro.Runtime("Map", "Set", "JSON"), onpause)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onpause(onpause)
}

// Onpause fn
func (p *Props) Onpause(onpause func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onpause', $1)", onpause)
	p.attrs["onpause"] = onpause
	return p
}

// Onplay fn
func Onplay(onplay func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onplay', $2)", macro.Runtime("Map", "Set", "JSON"), onplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onplay(onplay)
}

// Onplay fn
func (p *Props) Onplay(onplay func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onplay', $1)", onplay)
	p.attrs["onplay"] = onplay
	return p
}

// Onplaying fn
func Onplaying(onplaying func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onplaying', $2)", macro.Runtime("Map", "Set", "JSON"), onplaying)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onplaying(onplaying)
}

// Onplaying fn
func (p *Props) Onplaying(onplaying func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onplaying', $1)", onplaying)
	p.attrs["onplaying"] = onplaying
	return p
}

// Onprogress fn
func Onprogress(onprogress func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onprogress', $2)", macro.Runtime("Map", "Set", "JSON"), onprogress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onprogress(onprogress)
}

// Onprogress fn
func (p *Props) Onprogress(onprogress func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onprogress', $1)", onprogress)
	p.attrs["onprogress"] = onprogress
	return p
}

// Onratechange fn
func Onratechange(onratechange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onratechange', $2)", macro.Runtime("Map", "Set", "JSON"), onratechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onratechange(onratechange)
}

// Onratechange fn
func (p *Props) Onratechange(onratechange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onratechange', $1)", onratechange)
	p.attrs["onratechange"] = onratechange
	return p
}

// Onreset fn
func Onreset(onreset func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onreset', $2)", macro.Runtime("Map", "Set", "JSON"), onreset)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onreset(onreset)
}

// Onreset fn
func (p *Props) Onreset(onreset func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onreset', $1)", onreset)
	p.attrs["onreset"] = onreset
	return p
}

// Onresize fn
func Onresize(onresize func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('onresize', $2)", macro.Runtime("Map", "Set", "JSON"), onresize)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onresize(onresize)
}

// Onresize fn
func (p *Props) Onresize(onresize func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('onresize', $1)", onresize)
	p.attrs["onresize"] = onresize
	return p
}

// Onreadystatechange fn
func Onreadystatechange(onreadystatechange func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onreadystatechange', $2)", macro.Runtime("Map", "Set", "JSON"), onreadystatechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onreadystatechange(onreadystatechange)
}

// Onreadystatechange fn
func (p *Props) Onreadystatechange(onreadystatechange func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onreadystatechange', $1)", onreadystatechange)
	p.attrs["onreadystatechange"] = onreadystatechange
	return p
}

// Onscroll fn
func Onscroll(onscroll func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('onscroll', $2)", macro.Runtime("Map", "Set", "JSON"), onscroll)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onscroll(onscroll)
}

// Onscroll fn
func (p *Props) Onscroll(onscroll func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('onscroll', $1)", onscroll)
	p.attrs["onscroll"] = onscroll
	return p
}

// Onseeked fn
func Onseeked(onseeked func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onseeked', $2)", macro.Runtime("Map", "Set", "JSON"), onseeked)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onseeked(onseeked)
}

// Onseeked fn
func (p *Props) Onseeked(onseeked func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onseeked', $1)", onseeked)
	p.attrs["onseeked"] = onseeked
	return p
}

// Onseeking fn
func Onseeking(onseeking func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onseeking', $2)", macro.Runtime("Map", "Set", "JSON"), onseeking)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onseeking(onseeking)
}

// Onseeking fn
func (p *Props) Onseeking(onseeking func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onseeking', $1)", onseeking)
	p.attrs["onseeking"] = onseeking
	return p
}

// Onselect fn
func Onselect(onselect func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('onselect', $2)", macro.Runtime("Map", "Set", "JSON"), onselect)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onselect(onselect)
}

// Onselect fn
func (p *Props) Onselect(onselect func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('onselect', $1)", onselect)
	p.attrs["onselect"] = onselect
	return p
}

// Onshow fn
func Onshow(onshow func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onshow', $2)", macro.Runtime("Map", "Set", "JSON"), onshow)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onshow(onshow)
}

// Onshow fn
func (p *Props) Onshow(onshow func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onshow', $1)", onshow)
	p.attrs["onshow"] = onshow
	return p
}

// Onstalled fn
func Onstalled(onstalled func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onstalled', $2)", macro.Runtime("Map", "Set", "JSON"), onstalled)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onstalled(onstalled)
}

// Onstalled fn
func (p *Props) Onstalled(onstalled func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onstalled', $1)", onstalled)
	p.attrs["onstalled"] = onstalled
	return p
}

// Onsubmit fn
func Onsubmit(onsubmit func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onsubmit', $2)", macro.Runtime("Map", "Set", "JSON"), onsubmit)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onsubmit(onsubmit)
}

// Onsubmit fn
func (p *Props) Onsubmit(onsubmit func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onsubmit', $1)", onsubmit)
	p.attrs["onsubmit"] = onsubmit
	return p
}

// Onsuspend fn
func Onsuspend(onsuspend func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onsuspend', $2)", macro.Runtime("Map", "Set", "JSON"), onsuspend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onsuspend(onsuspend)
}

// Onsuspend fn
func (p *Props) Onsuspend(onsuspend func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onsuspend', $1)", onsuspend)
	p.attrs["onsuspend"] = onsuspend
	return p
}

// Ontimeupdate fn
func Ontimeupdate(ontimeupdate func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('ontimeupdate', $2)", macro.Runtime("Map", "Set", "JSON"), ontimeupdate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ontimeupdate(ontimeupdate)
}

// Ontimeupdate fn
func (p *Props) Ontimeupdate(ontimeupdate func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('ontimeupdate', $1)", ontimeupdate)
	p.attrs["ontimeupdate"] = ontimeupdate
	return p
}

// Onvolumechange fn
func Onvolumechange(onvolumechange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onvolumechange', $2)", macro.Runtime("Map", "Set", "JSON"), onvolumechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onvolumechange(onvolumechange)
}

// Onvolumechange fn
func (p *Props) Onvolumechange(onvolumechange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onvolumechange', $1)", onvolumechange)
	p.attrs["onvolumechange"] = onvolumechange
	return p
}

// Onwaiting fn
func Onwaiting(onwaiting func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onwaiting', $2)", macro.Runtime("Map", "Set", "JSON"), onwaiting)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onwaiting(onwaiting)
}

// Onwaiting fn
func (p *Props) Onwaiting(onwaiting func(e window.Event)) *Props {
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
