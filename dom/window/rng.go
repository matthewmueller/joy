package window

import (
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/dom/clientrectlist"
	"github.com/matthewmueller/joy/dom/expandgranularity"
	"github.com/matthewmueller/joy/macro"
)

// Range struct
// js:"Range,omit"
type Range struct {
}

// CloneContents fn
// js:"cloneContents"
func (*Range) CloneContents() (d *DocumentFragment) {
	macro.Rewrite("$_.cloneContents()")
	return d
}

// CloneRange fn
// js:"cloneRange"
func (*Range) CloneRange() (r *Range) {
	macro.Rewrite("$_.cloneRange()")
	return r
}

// Collapse fn
// js:"collapse"
func (*Range) Collapse(toStart bool) {
	macro.Rewrite("$_.collapse($1)", toStart)
}

// CompareBoundaryPoints fn
// js:"compareBoundaryPoints"
func (*Range) CompareBoundaryPoints(how uint8, sourceRange *Range) (i int8) {
	macro.Rewrite("$_.compareBoundaryPoints($1, $2)", how, sourceRange)
	return i
}

// CreateContextualFragment fn
// js:"createContextualFragment"
func (*Range) CreateContextualFragment(fragment string) (d *DocumentFragment) {
	macro.Rewrite("$_.createContextualFragment($1)", fragment)
	return d
}

// DeleteContents fn
// js:"deleteContents"
func (*Range) DeleteContents() {
	macro.Rewrite("$_.deleteContents()")
}

// Detach fn
// js:"detach"
func (*Range) Detach() {
	macro.Rewrite("$_.detach()")
}

// Expand fn
// js:"expand"
func (*Range) Expand(Unit *expandgranularity.ExpandGranularity) (b bool) {
	macro.Rewrite("$_.expand($1)", Unit)
	return b
}

// ExtractContents fn
// js:"extractContents"
func (*Range) ExtractContents() (d *DocumentFragment) {
	macro.Rewrite("$_.extractContents()")
	return d
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*Range) GetBoundingClientRect() (c *clientrect.ClientRect) {
	macro.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*Range) GetClientRects() (c *clientrectlist.ClientRectList) {
	macro.Rewrite("$_.getClientRects()")
	return c
}

// InsertNode fn
// js:"insertNode"
func (*Range) InsertNode(newNode Node) {
	macro.Rewrite("$_.insertNode($1)", newNode)
}

// SelectNode fn
// js:"selectNode"
func (*Range) SelectNode(refNode Node) {
	macro.Rewrite("$_.selectNode($1)", refNode)
}

// SelectNodeContents fn
// js:"selectNodeContents"
func (*Range) SelectNodeContents(refNode Node) {
	macro.Rewrite("$_.selectNodeContents($1)", refNode)
}

// SetEnd fn
// js:"setEnd"
func (*Range) SetEnd(refNode Node, offset int) {
	macro.Rewrite("$_.setEnd($1, $2)", refNode, offset)
}

// SetEndAfter fn
// js:"setEndAfter"
func (*Range) SetEndAfter(refNode Node) {
	macro.Rewrite("$_.setEndAfter($1)", refNode)
}

// SetEndBefore fn
// js:"setEndBefore"
func (*Range) SetEndBefore(refNode Node) {
	macro.Rewrite("$_.setEndBefore($1)", refNode)
}

// SetStart fn
// js:"setStart"
func (*Range) SetStart(refNode Node, offset int) {
	macro.Rewrite("$_.setStart($1, $2)", refNode, offset)
}

// SetStartAfter fn
// js:"setStartAfter"
func (*Range) SetStartAfter(refNode Node) {
	macro.Rewrite("$_.setStartAfter($1)", refNode)
}

// SetStartBefore fn
// js:"setStartBefore"
func (*Range) SetStartBefore(refNode Node) {
	macro.Rewrite("$_.setStartBefore($1)", refNode)
}

// SurroundContents fn
// js:"surroundContents"
func (*Range) SurroundContents(newParent Node) {
	macro.Rewrite("$_.surroundContents($1)", newParent)
}

// ToString fn
// js:"toString"
func (*Range) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Collapsed prop
// js:"collapsed"
func (*Range) Collapsed() (collapsed bool) {
	macro.Rewrite("$_.collapsed")
	return collapsed
}

// CommonAncestorContainer prop
// js:"commonAncestorContainer"
func (*Range) CommonAncestorContainer() (commonAncestorContainer Node) {
	macro.Rewrite("$_.commonAncestorContainer")
	return commonAncestorContainer
}

// EndContainer prop
// js:"endContainer"
func (*Range) EndContainer() (endContainer Node) {
	macro.Rewrite("$_.endContainer")
	return endContainer
}

// EndOffset prop
// js:"endOffset"
func (*Range) EndOffset() (endOffset int) {
	macro.Rewrite("$_.endOffset")
	return endOffset
}

// StartContainer prop
// js:"startContainer"
func (*Range) StartContainer() (startContainer Node) {
	macro.Rewrite("$_.startContainer")
	return startContainer
}

// StartOffset prop
// js:"startOffset"
func (*Range) StartOffset() (startOffset int) {
	macro.Rewrite("$_.startOffset")
	return startOffset
}
