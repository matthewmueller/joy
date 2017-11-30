package window

import (
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/expandgranularity"
	"github.com/matthewmueller/golly/js"
)

// Range struct
// js:"Range,omit"
type Range struct {
}

// CloneContents fn
// js:"cloneContents"
func (*Range) CloneContents() (d *DocumentFragment) {
	js.Rewrite("$_.cloneContents()")
	return d
}

// CloneRange fn
// js:"cloneRange"
func (*Range) CloneRange() (r *Range) {
	js.Rewrite("$_.cloneRange()")
	return r
}

// Collapse fn
// js:"collapse"
func (*Range) Collapse(toStart bool) {
	js.Rewrite("$_.collapse($1)", toStart)
}

// CompareBoundaryPoints fn
// js:"compareBoundaryPoints"
func (*Range) CompareBoundaryPoints(how uint8, sourceRange *Range) (i int8) {
	js.Rewrite("$_.compareBoundaryPoints($1, $2)", how, sourceRange)
	return i
}

// CreateContextualFragment fn
// js:"createContextualFragment"
func (*Range) CreateContextualFragment(fragment string) (d *DocumentFragment) {
	js.Rewrite("$_.createContextualFragment($1)", fragment)
	return d
}

// DeleteContents fn
// js:"deleteContents"
func (*Range) DeleteContents() {
	js.Rewrite("$_.deleteContents()")
}

// Detach fn
// js:"detach"
func (*Range) Detach() {
	js.Rewrite("$_.detach()")
}

// Expand fn
// js:"expand"
func (*Range) Expand(Unit *expandgranularity.ExpandGranularity) (b bool) {
	js.Rewrite("$_.expand($1)", Unit)
	return b
}

// ExtractContents fn
// js:"extractContents"
func (*Range) ExtractContents() (d *DocumentFragment) {
	js.Rewrite("$_.extractContents()")
	return d
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*Range) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*Range) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// InsertNode fn
// js:"insertNode"
func (*Range) InsertNode(newNode Node) {
	js.Rewrite("$_.insertNode($1)", newNode)
}

// SelectNode fn
// js:"selectNode"
func (*Range) SelectNode(refNode Node) {
	js.Rewrite("$_.selectNode($1)", refNode)
}

// SelectNodeContents fn
// js:"selectNodeContents"
func (*Range) SelectNodeContents(refNode Node) {
	js.Rewrite("$_.selectNodeContents($1)", refNode)
}

// SetEnd fn
// js:"setEnd"
func (*Range) SetEnd(refNode Node, offset int) {
	js.Rewrite("$_.setEnd($1, $2)", refNode, offset)
}

// SetEndAfter fn
// js:"setEndAfter"
func (*Range) SetEndAfter(refNode Node) {
	js.Rewrite("$_.setEndAfter($1)", refNode)
}

// SetEndBefore fn
// js:"setEndBefore"
func (*Range) SetEndBefore(refNode Node) {
	js.Rewrite("$_.setEndBefore($1)", refNode)
}

// SetStart fn
// js:"setStart"
func (*Range) SetStart(refNode Node, offset int) {
	js.Rewrite("$_.setStart($1, $2)", refNode, offset)
}

// SetStartAfter fn
// js:"setStartAfter"
func (*Range) SetStartAfter(refNode Node) {
	js.Rewrite("$_.setStartAfter($1)", refNode)
}

// SetStartBefore fn
// js:"setStartBefore"
func (*Range) SetStartBefore(refNode Node) {
	js.Rewrite("$_.setStartBefore($1)", refNode)
}

// SurroundContents fn
// js:"surroundContents"
func (*Range) SurroundContents(newParent Node) {
	js.Rewrite("$_.surroundContents($1)", newParent)
}

// ToString fn
// js:"toString"
func (*Range) ToString() (s string) {
	js.Rewrite("$_.toString()")
	return s
}

// Collapsed prop
// js:"collapsed"
func (*Range) Collapsed() (collapsed bool) {
	js.Rewrite("$_.collapsed")
	return collapsed
}

// CommonAncestorContainer prop
// js:"commonAncestorContainer"
func (*Range) CommonAncestorContainer() (commonAncestorContainer Node) {
	js.Rewrite("$_.commonAncestorContainer")
	return commonAncestorContainer
}

// EndContainer prop
// js:"endContainer"
func (*Range) EndContainer() (endContainer Node) {
	js.Rewrite("$_.endContainer")
	return endContainer
}

// EndOffset prop
// js:"endOffset"
func (*Range) EndOffset() (endOffset int) {
	js.Rewrite("$_.endOffset")
	return endOffset
}

// StartContainer prop
// js:"startContainer"
func (*Range) StartContainer() (startContainer Node) {
	js.Rewrite("$_.startContainer")
	return startContainer
}

// StartOffset prop
// js:"startOffset"
func (*Range) StartOffset() (startOffset int) {
	js.Rewrite("$_.startOffset")
	return startOffset
}
