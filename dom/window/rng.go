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
func (*Range) CloneContents() (d *DocumentFragment) {
	js.Rewrite("$<.cloneContents()")
	return d
}

// CloneRange fn
func (*Range) CloneRange() (r *Range) {
	js.Rewrite("$<.cloneRange()")
	return r
}

// Collapse fn
func (*Range) Collapse(toStart bool) {
	js.Rewrite("$<.collapse($1)", toStart)
}

// CompareBoundaryPoints fn
func (*Range) CompareBoundaryPoints(how uint8, sourceRange *Range) (i int8) {
	js.Rewrite("$<.compareBoundaryPoints($1, $2)", how, sourceRange)
	return i
}

// CreateContextualFragment fn
func (*Range) CreateContextualFragment(fragment string) (d *DocumentFragment) {
	js.Rewrite("$<.createContextualFragment($1)", fragment)
	return d
}

// DeleteContents fn
func (*Range) DeleteContents() {
	js.Rewrite("$<.deleteContents()")
}

// Detach fn
func (*Range) Detach() {
	js.Rewrite("$<.detach()")
}

// Expand fn
func (*Range) Expand(Unit *expandgranularity.ExpandGranularity) (b bool) {
	js.Rewrite("$<.expand($1)", Unit)
	return b
}

// ExtractContents fn
func (*Range) ExtractContents() (d *DocumentFragment) {
	js.Rewrite("$<.extractContents()")
	return d
}

// GetBoundingClientRect fn
func (*Range) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.getBoundingClientRect()")
	return c
}

// GetClientRects fn
func (*Range) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$<.getClientRects()")
	return c
}

// InsertNode fn
func (*Range) InsertNode(newNode Node) {
	js.Rewrite("$<.insertNode($1)", newNode)
}

// SelectNode fn
func (*Range) SelectNode(refNode Node) {
	js.Rewrite("$<.selectNode($1)", refNode)
}

// SelectNodeContents fn
func (*Range) SelectNodeContents(refNode Node) {
	js.Rewrite("$<.selectNodeContents($1)", refNode)
}

// SetEnd fn
func (*Range) SetEnd(refNode Node, offset int) {
	js.Rewrite("$<.setEnd($1, $2)", refNode, offset)
}

// SetEndAfter fn
func (*Range) SetEndAfter(refNode Node) {
	js.Rewrite("$<.setEndAfter($1)", refNode)
}

// SetEndBefore fn
func (*Range) SetEndBefore(refNode Node) {
	js.Rewrite("$<.setEndBefore($1)", refNode)
}

// SetStart fn
func (*Range) SetStart(refNode Node, offset int) {
	js.Rewrite("$<.setStart($1, $2)", refNode, offset)
}

// SetStartAfter fn
func (*Range) SetStartAfter(refNode Node) {
	js.Rewrite("$<.setStartAfter($1)", refNode)
}

// SetStartBefore fn
func (*Range) SetStartBefore(refNode Node) {
	js.Rewrite("$<.setStartBefore($1)", refNode)
}

// SurroundContents fn
func (*Range) SurroundContents(newParent Node) {
	js.Rewrite("$<.surroundContents($1)", newParent)
}

// ToString fn
func (*Range) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Collapsed prop
func (*Range) Collapsed() (collapsed bool) {
	js.Rewrite("$<.collapsed")
	return collapsed
}

// CommonAncestorContainer prop
func (*Range) CommonAncestorContainer() (commonAncestorContainer Node) {
	js.Rewrite("$<.commonAncestorContainer")
	return commonAncestorContainer
}

// EndContainer prop
func (*Range) EndContainer() (endContainer Node) {
	js.Rewrite("$<.endContainer")
	return endContainer
}

// EndOffset prop
func (*Range) EndOffset() (endOffset int) {
	js.Rewrite("$<.endOffset")
	return endOffset
}

// StartContainer prop
func (*Range) StartContainer() (startContainer Node) {
	js.Rewrite("$<.startContainer")
	return startContainer
}

// StartOffset prop
func (*Range) StartOffset() (startOffset int) {
	js.Rewrite("$<.startOffset")
	return startOffset
}
