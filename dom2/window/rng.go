package window

import (
	"github.com/matthewmueller/golly/dom2/clientrect"
	"github.com/matthewmueller/golly/dom2/clientrectlist"
	"github.com/matthewmueller/golly/dom2/expandgranularity"
	"github.com/matthewmueller/golly/js"
)

// Range struct
// js:"Range,omit"
type Range struct {
}

// CloneContents
func (*Range) CloneContents() (d *DocumentFragment) {
	js.Rewrite("$<.CloneContents()")
	return d
}

// CloneRange
func (*Range) CloneRange() (r *Range) {
	js.Rewrite("$<.CloneRange()")
	return r
}

// Collapse
func (*Range) Collapse(toStart bool) {
	js.Rewrite("$<.Collapse($1)", toStart)
}

// CompareBoundaryPoints
func (*Range) CompareBoundaryPoints(how uint8, sourceRange *Range) (i int8) {
	js.Rewrite("$<.CompareBoundaryPoints($1, $2)", how, sourceRange)
	return i
}

// CreateContextualFragment
func (*Range) CreateContextualFragment(fragment string) (d *DocumentFragment) {
	js.Rewrite("$<.CreateContextualFragment($1)", fragment)
	return d
}

// DeleteContents
func (*Range) DeleteContents() {
	js.Rewrite("$<.DeleteContents()")
}

// Detach
func (*Range) Detach() {
	js.Rewrite("$<.Detach()")
}

// Expand
func (*Range) Expand(Unit *expandgranularity.ExpandGranularity) (b bool) {
	js.Rewrite("$<.Expand($1)", Unit)
	return b
}

// ExtractContents
func (*Range) ExtractContents() (d *DocumentFragment) {
	js.Rewrite("$<.ExtractContents()")
	return d
}

// GetBoundingClientRect
func (*Range) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.GetBoundingClientRect()")
	return c
}

// GetClientRects
func (*Range) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$<.GetClientRects()")
	return c
}

// InsertNode
func (*Range) InsertNode(newNode Node) {
	js.Rewrite("$<.InsertNode($1)", newNode)
}

// SelectNode
func (*Range) SelectNode(refNode Node) {
	js.Rewrite("$<.SelectNode($1)", refNode)
}

// SelectNodeContents
func (*Range) SelectNodeContents(refNode Node) {
	js.Rewrite("$<.SelectNodeContents($1)", refNode)
}

// SetEnd
func (*Range) SetEnd(refNode Node, offset int) {
	js.Rewrite("$<.SetEnd($1, $2)", refNode, offset)
}

// SetEndAfter
func (*Range) SetEndAfter(refNode Node) {
	js.Rewrite("$<.SetEndAfter($1)", refNode)
}

// SetEndBefore
func (*Range) SetEndBefore(refNode Node) {
	js.Rewrite("$<.SetEndBefore($1)", refNode)
}

// SetStart
func (*Range) SetStart(refNode Node, offset int) {
	js.Rewrite("$<.SetStart($1, $2)", refNode, offset)
}

// SetStartAfter
func (*Range) SetStartAfter(refNode Node) {
	js.Rewrite("$<.SetStartAfter($1)", refNode)
}

// SetStartBefore
func (*Range) SetStartBefore(refNode Node) {
	js.Rewrite("$<.SetStartBefore($1)", refNode)
}

// SurroundContents
func (*Range) SurroundContents(newParent Node) {
	js.Rewrite("$<.SurroundContents($1)", newParent)
}

// ToString
func (*Range) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Collapsed
func (*Range) Collapsed() (collapsed bool) {
	js.Rewrite("$<.Collapsed")
	return collapsed
}

// CommonAncestorContainer
func (*Range) CommonAncestorContainer() (commonAncestorContainer Node) {
	js.Rewrite("$<.CommonAncestorContainer")
	return commonAncestorContainer
}

// EndContainer
func (*Range) EndContainer() (endContainer Node) {
	js.Rewrite("$<.EndContainer")
	return endContainer
}

// EndOffset
func (*Range) EndOffset() (endOffset int) {
	js.Rewrite("$<.EndOffset")
	return endOffset
}

// StartContainer
func (*Range) StartContainer() (startContainer Node) {
	js.Rewrite("$<.StartContainer")
	return startContainer
}

// StartOffset
func (*Range) StartOffset() (startOffset int) {
	js.Rewrite("$<.StartOffset")
	return startOffset
}
