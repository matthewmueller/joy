package b

import (
	"encoding/json"
	"strings"

	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/vdom"
)

// B struct
// js:"b,omit"
type B struct {
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

// New b element
//
// The b element represents a span of text to which attention is being drawn for utilitarian purposes without conveying any extra importance and with no implication of an alternate voice or mood, such as key words in a document abstract, product names in a review, actionable words in interactive text-driven software, or an article lede.
func New(props *Props, children ...vdom.Child) *B {
	macro.Rewrite("$1('b', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
	if props == nil {
		props = &Props{attrs: map[string]interface{}{}}
	}
	return &B{
		attrs:    props.attrs,
		children: children,
	}
}

// Render fn
func (s *B) Render() vdom.Node {
	return s
}

func (s *B) String() string {
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
		return "<b " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</b>"
	}

	return "<b>" + strings.Join(children, "") + "</b>"
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

// OnAbort fn
func OnAbort(onAbort func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('on-abort', $2)", macro.Runtime("Map", "Set", "JSON"), onAbort)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnAbort(onAbort)
}

// OnAbort fn
func (p *Props) OnAbort(onAbort func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('on-abort', $1)", onAbort)
	p.attrs["on-abort"] = onAbort
	return p
}

// OnBlur fn
func OnBlur(onBlur func(e window.FocusEvent)) *Props {
	macro.Rewrite("$1().Set('on-blur', $2)", macro.Runtime("Map", "Set", "JSON"), onBlur)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnBlur(onBlur)
}

// OnBlur fn
func (p *Props) OnBlur(onBlur func(e window.FocusEvent)) *Props {
	macro.Rewrite("$_.Set('on-blur', $1)", onBlur)
	p.attrs["on-blur"] = onBlur
	return p
}

// OnCanplay fn
func OnCanplay(onCanplay func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-canplay', $2)", macro.Runtime("Map", "Set", "JSON"), onCanplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCanplay(onCanplay)
}

// OnCanplay fn
func (p *Props) OnCanplay(onCanplay func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-canplay', $1)", onCanplay)
	p.attrs["on-canplay"] = onCanplay
	return p
}

// OnCanplaythrough fn
func OnCanplaythrough(onCanplaythrough func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-canplaythrough', $2)", macro.Runtime("Map", "Set", "JSON"), onCanplaythrough)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnCanplaythrough(onCanplaythrough)
}

// OnCanplaythrough fn
func (p *Props) OnCanplaythrough(onCanplaythrough func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-canplaythrough', $1)", onCanplaythrough)
	p.attrs["on-canplaythrough"] = onCanplaythrough
	return p
}

// OnChange fn
func OnChange(onChange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-change', $2)", macro.Runtime("Map", "Set", "JSON"), onChange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnChange(onChange)
}

// OnChange fn
func (p *Props) OnChange(onChange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-change', $1)", onChange)
	p.attrs["on-change"] = onChange
	return p
}

// OnClick fn
func OnClick(onClick func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-click', $2)", macro.Runtime("Map", "Set", "JSON"), onClick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnClick(onClick)
}

// OnClick fn
func (p *Props) OnClick(onClick func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-click', $1)", onClick)
	p.attrs["on-click"] = onClick
	return p
}

// OnContextmenu fn
func OnContextmenu(onContextmenu func(e window.PointerEvent)) *Props {
	macro.Rewrite("$1().Set('on-contextmenu', $2)", macro.Runtime("Map", "Set", "JSON"), onContextmenu)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnContextmenu(onContextmenu)
}

// OnContextmenu fn
func (p *Props) OnContextmenu(onContextmenu func(e window.PointerEvent)) *Props {
	macro.Rewrite("$_.Set('on-contextmenu', $1)", onContextmenu)
	p.attrs["on-contextmenu"] = onContextmenu
	return p
}

// OnDblclick fn
func OnDblclick(onDblclick func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('on-dblclick', $2)", macro.Runtime("Map", "Set", "JSON"), onDblclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDblclick(onDblclick)
}

// OnDblclick fn
func (p *Props) OnDblclick(onDblclick func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('on-dblclick', $1)", onDblclick)
	p.attrs["on-dblclick"] = onDblclick
	return p
}

// OnDrag fn
func OnDrag(onDrag func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('on-drag', $2)", macro.Runtime("Map", "Set", "JSON"), onDrag)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDrag(onDrag)
}

// OnDrag fn
func (p *Props) OnDrag(onDrag func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('on-drag', $1)", onDrag)
	p.attrs["on-drag"] = onDrag
	return p
}

// OnDragend fn
func OnDragend(onDragend func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('on-dragend', $2)", macro.Runtime("Map", "Set", "JSON"), onDragend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragend(onDragend)
}

// OnDragend fn
func (p *Props) OnDragend(onDragend func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('on-dragend', $1)", onDragend)
	p.attrs["on-dragend"] = onDragend
	return p
}

// OnDragenter fn
func OnDragenter(onDragenter func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('on-dragenter', $2)", macro.Runtime("Map", "Set", "JSON"), onDragenter)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragenter(onDragenter)
}

// OnDragenter fn
func (p *Props) OnDragenter(onDragenter func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('on-dragenter', $1)", onDragenter)
	p.attrs["on-dragenter"] = onDragenter
	return p
}

// OnDragleave fn
func OnDragleave(onDragleave func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('on-dragleave', $2)", macro.Runtime("Map", "Set", "JSON"), onDragleave)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragleave(onDragleave)
}

// OnDragleave fn
func (p *Props) OnDragleave(onDragleave func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('on-dragleave', $1)", onDragleave)
	p.attrs["on-dragleave"] = onDragleave
	return p
}

// OnDragover fn
func OnDragover(onDragover func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('on-dragover', $2)", macro.Runtime("Map", "Set", "JSON"), onDragover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragover(onDragover)
}

// OnDragover fn
func (p *Props) OnDragover(onDragover func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('on-dragover', $1)", onDragover)
	p.attrs["on-dragover"] = onDragover
	return p
}

// OnDragstart fn
func OnDragstart(onDragstart func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('on-dragstart', $2)", macro.Runtime("Map", "Set", "JSON"), onDragstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDragstart(onDragstart)
}

// OnDragstart fn
func (p *Props) OnDragstart(onDragstart func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('on-dragstart', $1)", onDragstart)
	p.attrs["on-dragstart"] = onDragstart
	return p
}

// OnDrop fn
func OnDrop(onDrop func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('on-drop', $2)", macro.Runtime("Map", "Set", "JSON"), onDrop)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDrop(onDrop)
}

// OnDrop fn
func (p *Props) OnDrop(onDrop func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('on-drop', $1)", onDrop)
	p.attrs["on-drop"] = onDrop
	return p
}

// OnDurationchange fn
func OnDurationchange(onDurationchange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-durationchange', $2)", macro.Runtime("Map", "Set", "JSON"), onDurationchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnDurationchange(onDurationchange)
}

// OnDurationchange fn
func (p *Props) OnDurationchange(onDurationchange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-durationchange', $1)", onDurationchange)
	p.attrs["on-durationchange"] = onDurationchange
	return p
}

// OnEmptied fn
func OnEmptied(onEmptied func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-emptied', $2)", macro.Runtime("Map", "Set", "JSON"), onEmptied)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnEmptied(onEmptied)
}

// OnEmptied fn
func (p *Props) OnEmptied(onEmptied func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-emptied', $1)", onEmptied)
	p.attrs["on-emptied"] = onEmptied
	return p
}

// OnEnded fn
func OnEnded(onEnded func(e window.MediaStreamErrorEvent)) *Props {
	macro.Rewrite("$1().Set('on-ended', $2)", macro.Runtime("Map", "Set", "JSON"), onEnded)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnEnded(onEnded)
}

// OnEnded fn
func (p *Props) OnEnded(onEnded func(e window.MediaStreamErrorEvent)) *Props {
	macro.Rewrite("$_.Set('on-ended', $1)", onEnded)
	p.attrs["on-ended"] = onEnded
	return p
}

// OnError fn
func OnError(onError func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('on-error', $2)", macro.Runtime("Map", "Set", "JSON"), onError)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnError(onError)
}

// OnError fn
func (p *Props) OnError(onError func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('on-error', $1)", onError)
	p.attrs["on-error"] = onError
	return p
}

// OnFocus fn
func OnFocus(onFocus func(e window.FocusEvent)) *Props {
	macro.Rewrite("$1().Set('on-focus', $2)", macro.Runtime("Map", "Set", "JSON"), onFocus)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnFocus(onFocus)
}

// OnFocus fn
func (p *Props) OnFocus(onFocus func(e window.FocusEvent)) *Props {
	macro.Rewrite("$_.Set('on-focus', $1)", onFocus)
	p.attrs["on-focus"] = onFocus
	return p
}

// OnInput fn
func OnInput(onInput func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-input', $2)", macro.Runtime("Map", "Set", "JSON"), onInput)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnInput(onInput)
}

// OnInput fn
func (p *Props) OnInput(onInput func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-input', $1)", onInput)
	p.attrs["on-input"] = onInput
	return p
}

// OnInvalid fn
func OnInvalid(onInvalid func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-invalid', $2)", macro.Runtime("Map", "Set", "JSON"), onInvalid)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnInvalid(onInvalid)
}

// OnInvalid fn
func (p *Props) OnInvalid(onInvalid func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-invalid', $1)", onInvalid)
	p.attrs["on-invalid"] = onInvalid
	return p
}

// OnKeydown fn
func OnKeydown(onKeydown func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('on-keydown', $2)", macro.Runtime("Map", "Set", "JSON"), onKeydown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnKeydown(onKeydown)
}

// OnKeydown fn
func (p *Props) OnKeydown(onKeydown func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('on-keydown', $1)", onKeydown)
	p.attrs["on-keydown"] = onKeydown
	return p
}

// OnKeypress fn
func OnKeypress(onKeypress func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('on-keypress', $2)", macro.Runtime("Map", "Set", "JSON"), onKeypress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnKeypress(onKeypress)
}

// OnKeypress fn
func (p *Props) OnKeypress(onKeypress func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('on-keypress', $1)", onKeypress)
	p.attrs["on-keypress"] = onKeypress
	return p
}

// OnKeyup fn
func OnKeyup(onKeyup func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('on-keyup', $2)", macro.Runtime("Map", "Set", "JSON"), onKeyup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnKeyup(onKeyup)
}

// OnKeyup fn
func (p *Props) OnKeyup(onKeyup func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('on-keyup', $1)", onKeyup)
	p.attrs["on-keyup"] = onKeyup
	return p
}

// OnLoad fn
func OnLoad(onLoad func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('on-load', $2)", macro.Runtime("Map", "Set", "JSON"), onLoad)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoad(onLoad)
}

// OnLoad fn
func (p *Props) OnLoad(onLoad func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('on-load', $1)", onLoad)
	p.attrs["on-load"] = onLoad
	return p
}

// OnLoadeddata fn
func OnLoadeddata(onLoadeddata func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-loadeddata', $2)", macro.Runtime("Map", "Set", "JSON"), onLoadeddata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoadeddata(onLoadeddata)
}

// OnLoadeddata fn
func (p *Props) OnLoadeddata(onLoadeddata func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-loadeddata', $1)", onLoadeddata)
	p.attrs["on-loadeddata"] = onLoadeddata
	return p
}

// OnLoadedmetadata fn
func OnLoadedmetadata(onLoadedmetadata func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-loadedmetadata', $2)", macro.Runtime("Map", "Set", "JSON"), onLoadedmetadata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoadedmetadata(onLoadedmetadata)
}

// OnLoadedmetadata fn
func (p *Props) OnLoadedmetadata(onLoadedmetadata func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-loadedmetadata', $1)", onLoadedmetadata)
	p.attrs["on-loadedmetadata"] = onLoadedmetadata
	return p
}

// OnLoadstart fn
func OnLoadstart(onLoadstart func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('on-loadstart', $2)", macro.Runtime("Map", "Set", "JSON"), onLoadstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnLoadstart(onLoadstart)
}

// OnLoadstart fn
func (p *Props) OnLoadstart(onLoadstart func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('on-loadstart', $1)", onLoadstart)
	p.attrs["on-loadstart"] = onLoadstart
	return p
}

// OnMousedown fn
func OnMousedown(onMousedown func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('on-mousedown', $2)", macro.Runtime("Map", "Set", "JSON"), onMousedown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMousedown(onMousedown)
}

// OnMousedown fn
func (p *Props) OnMousedown(onMousedown func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('on-mousedown', $1)", onMousedown)
	p.attrs["on-mousedown"] = onMousedown
	return p
}

// OnMousemove fn
func OnMousemove(onMousemove func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('on-mousemove', $2)", macro.Runtime("Map", "Set", "JSON"), onMousemove)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMousemove(onMousemove)
}

// OnMousemove fn
func (p *Props) OnMousemove(onMousemove func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('on-mousemove', $1)", onMousemove)
	p.attrs["on-mousemove"] = onMousemove
	return p
}

// OnMouseout fn
func OnMouseout(onMouseout func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('on-mouseout', $2)", macro.Runtime("Map", "Set", "JSON"), onMouseout)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseout(onMouseout)
}

// OnMouseout fn
func (p *Props) OnMouseout(onMouseout func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('on-mouseout', $1)", onMouseout)
	p.attrs["on-mouseout"] = onMouseout
	return p
}

// OnMouseover fn
func OnMouseover(onMouseover func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('on-mouseover', $2)", macro.Runtime("Map", "Set", "JSON"), onMouseover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseover(onMouseover)
}

// OnMouseover fn
func (p *Props) OnMouseover(onMouseover func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('on-mouseover', $1)", onMouseover)
	p.attrs["on-mouseover"] = onMouseover
	return p
}

// OnMouseup fn
func OnMouseup(onMouseup func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('on-mouseup', $2)", macro.Runtime("Map", "Set", "JSON"), onMouseup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMouseup(onMouseup)
}

// OnMouseup fn
func (p *Props) OnMouseup(onMouseup func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('on-mouseup', $1)", onMouseup)
	p.attrs["on-mouseup"] = onMouseup
	return p
}

// OnMousewheel fn
func OnMousewheel(onMousewheel func(e window.WheelEvent)) *Props {
	macro.Rewrite("$1().Set('on-mousewheel', $2)", macro.Runtime("Map", "Set", "JSON"), onMousewheel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnMousewheel(onMousewheel)
}

// OnMousewheel fn
func (p *Props) OnMousewheel(onMousewheel func(e window.WheelEvent)) *Props {
	macro.Rewrite("$_.Set('on-mousewheel', $1)", onMousewheel)
	p.attrs["on-mousewheel"] = onMousewheel
	return p
}

// OnPause fn
func OnPause(onPause func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-pause', $2)", macro.Runtime("Map", "Set", "JSON"), onPause)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnPause(onPause)
}

// OnPause fn
func (p *Props) OnPause(onPause func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-pause', $1)", onPause)
	p.attrs["on-pause"] = onPause
	return p
}

// OnPlay fn
func OnPlay(onPlay func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-play', $2)", macro.Runtime("Map", "Set", "JSON"), onPlay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnPlay(onPlay)
}

// OnPlay fn
func (p *Props) OnPlay(onPlay func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-play', $1)", onPlay)
	p.attrs["on-play"] = onPlay
	return p
}

// OnPlaying fn
func OnPlaying(onPlaying func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-playing', $2)", macro.Runtime("Map", "Set", "JSON"), onPlaying)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnPlaying(onPlaying)
}

// OnPlaying fn
func (p *Props) OnPlaying(onPlaying func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-playing', $1)", onPlaying)
	p.attrs["on-playing"] = onPlaying
	return p
}

// OnProgress fn
func OnProgress(onProgress func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('on-progress', $2)", macro.Runtime("Map", "Set", "JSON"), onProgress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnProgress(onProgress)
}

// OnProgress fn
func (p *Props) OnProgress(onProgress func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('on-progress', $1)", onProgress)
	p.attrs["on-progress"] = onProgress
	return p
}

// OnRatechange fn
func OnRatechange(onRatechange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-ratechange', $2)", macro.Runtime("Map", "Set", "JSON"), onRatechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnRatechange(onRatechange)
}

// OnRatechange fn
func (p *Props) OnRatechange(onRatechange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-ratechange', $1)", onRatechange)
	p.attrs["on-ratechange"] = onRatechange
	return p
}

// OnReset fn
func OnReset(onReset func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-reset', $2)", macro.Runtime("Map", "Set", "JSON"), onReset)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnReset(onReset)
}

// OnReset fn
func (p *Props) OnReset(onReset func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-reset', $1)", onReset)
	p.attrs["on-reset"] = onReset
	return p
}

// OnResize fn
func OnResize(onResize func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('on-resize', $2)", macro.Runtime("Map", "Set", "JSON"), onResize)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnResize(onResize)
}

// OnResize fn
func (p *Props) OnResize(onResize func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('on-resize', $1)", onResize)
	p.attrs["on-resize"] = onResize
	return p
}

// OnReadystatechange fn
func OnReadystatechange(onReadystatechange func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('on-readystatechange', $2)", macro.Runtime("Map", "Set", "JSON"), onReadystatechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnReadystatechange(onReadystatechange)
}

// OnReadystatechange fn
func (p *Props) OnReadystatechange(onReadystatechange func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('on-readystatechange', $1)", onReadystatechange)
	p.attrs["on-readystatechange"] = onReadystatechange
	return p
}

// OnScroll fn
func OnScroll(onScroll func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('on-scroll', $2)", macro.Runtime("Map", "Set", "JSON"), onScroll)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnScroll(onScroll)
}

// OnScroll fn
func (p *Props) OnScroll(onScroll func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('on-scroll', $1)", onScroll)
	p.attrs["on-scroll"] = onScroll
	return p
}

// OnSeeked fn
func OnSeeked(onSeeked func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-seeked', $2)", macro.Runtime("Map", "Set", "JSON"), onSeeked)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSeeked(onSeeked)
}

// OnSeeked fn
func (p *Props) OnSeeked(onSeeked func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-seeked', $1)", onSeeked)
	p.attrs["on-seeked"] = onSeeked
	return p
}

// OnSeeking fn
func OnSeeking(onSeeking func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-seeking', $2)", macro.Runtime("Map", "Set", "JSON"), onSeeking)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSeeking(onSeeking)
}

// OnSeeking fn
func (p *Props) OnSeeking(onSeeking func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-seeking', $1)", onSeeking)
	p.attrs["on-seeking"] = onSeeking
	return p
}

// OnSelect fn
func OnSelect(onSelect func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('on-select', $2)", macro.Runtime("Map", "Set", "JSON"), onSelect)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSelect(onSelect)
}

// OnSelect fn
func (p *Props) OnSelect(onSelect func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('on-select', $1)", onSelect)
	p.attrs["on-select"] = onSelect
	return p
}

// OnShow fn
func OnShow(onShow func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-show', $2)", macro.Runtime("Map", "Set", "JSON"), onShow)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnShow(onShow)
}

// OnShow fn
func (p *Props) OnShow(onShow func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-show', $1)", onShow)
	p.attrs["on-show"] = onShow
	return p
}

// OnStalled fn
func OnStalled(onStalled func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-stalled', $2)", macro.Runtime("Map", "Set", "JSON"), onStalled)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnStalled(onStalled)
}

// OnStalled fn
func (p *Props) OnStalled(onStalled func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-stalled', $1)", onStalled)
	p.attrs["on-stalled"] = onStalled
	return p
}

// OnSubmit fn
func OnSubmit(onSubmit func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-submit', $2)", macro.Runtime("Map", "Set", "JSON"), onSubmit)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSubmit(onSubmit)
}

// OnSubmit fn
func (p *Props) OnSubmit(onSubmit func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-submit', $1)", onSubmit)
	p.attrs["on-submit"] = onSubmit
	return p
}

// OnSuspend fn
func OnSuspend(onSuspend func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-suspend', $2)", macro.Runtime("Map", "Set", "JSON"), onSuspend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnSuspend(onSuspend)
}

// OnSuspend fn
func (p *Props) OnSuspend(onSuspend func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-suspend', $1)", onSuspend)
	p.attrs["on-suspend"] = onSuspend
	return p
}

// OnTimeupdate fn
func OnTimeupdate(onTimeupdate func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-timeupdate', $2)", macro.Runtime("Map", "Set", "JSON"), onTimeupdate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnTimeupdate(onTimeupdate)
}

// OnTimeupdate fn
func (p *Props) OnTimeupdate(onTimeupdate func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-timeupdate', $1)", onTimeupdate)
	p.attrs["on-timeupdate"] = onTimeupdate
	return p
}

// OnVolumechange fn
func OnVolumechange(onVolumechange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-volumechange', $2)", macro.Runtime("Map", "Set", "JSON"), onVolumechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnVolumechange(onVolumechange)
}

// OnVolumechange fn
func (p *Props) OnVolumechange(onVolumechange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-volumechange', $1)", onVolumechange)
	p.attrs["on-volumechange"] = onVolumechange
	return p
}

// OnWaiting fn
func OnWaiting(onWaiting func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('on-waiting', $2)", macro.Runtime("Map", "Set", "JSON"), onWaiting)
	p := &Props{attrs: map[string]interface{}{}}
	return p.OnWaiting(onWaiting)
}

// OnWaiting fn
func (p *Props) OnWaiting(onWaiting func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('on-waiting', $1)", onWaiting)
	p.attrs["on-waiting"] = onWaiting
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
