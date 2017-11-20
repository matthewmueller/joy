package rng

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/clientrect"
	"github.com/matthewmueller/golly/internal/gendom/dom/clientrectlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/documentfragment"
	"github.com/matthewmueller/golly/internal/gendom/dom/expandgranularity"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type Range struct {
}

func (*Range) CloneContents() (d *documentfragment.DocumentFragment) {
	js.Rewrite("$<.cloneContents()")
	return d
}

func (*Range) CloneRange() (r *Range) {
	js.Rewrite("$<.cloneRange()")
	return r
}

func (*Range) Collapse(toStart bool) {
	js.Rewrite("$<.collapse($1)", toStart)
}

func (*Range) CompareBoundaryPoints(how uint8, sourceRange *Range) (i int8) {
	js.Rewrite("$<.compareBoundaryPoints($1, $2)", how, sourceRange)
	return i
}

func (*Range) CreateContextualFragment(fragment string) (d *documentfragment.DocumentFragment) {
	js.Rewrite("$<.createContextualFragment($1)", fragment)
	return d
}

func (*Range) DeleteContents() {
	js.Rewrite("$<.deleteContents()")
}

func (*Range) Detach() {
	js.Rewrite("$<.detach()")
}

func (*Range) Expand(Unit *expandgranularity.ExpandGranularity) (b bool) {
	js.Rewrite("$<.expand($1)", Unit)
	return b
}

func (*Range) ExtractContents() (d *documentfragment.DocumentFragment) {
	js.Rewrite("$<.extractContents()")
	return d
}

func (*Range) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$<.getBoundingClientRect()")
	return c
}

func (*Range) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$<.getClientRects()")
	return c
}

func (*Range) InsertNode(newNode *node.Node) {
	js.Rewrite("$<.insertNode($1)", newNode)
}

func (*Range) SelectNode(refNode *node.Node) {
	js.Rewrite("$<.selectNode($1)", refNode)
}

func (*Range) SelectNodeContents(refNode *node.Node) {
	js.Rewrite("$<.selectNodeContents($1)", refNode)
}

func (*Range) SetEnd(refNode *node.Node, offset int) {
	js.Rewrite("$<.setEnd($1, $2)", refNode, offset)
}

func (*Range) SetEndAfter(refNode *node.Node) {
	js.Rewrite("$<.setEndAfter($1)", refNode)
}

func (*Range) SetEndBefore(refNode *node.Node) {
	js.Rewrite("$<.setEndBefore($1)", refNode)
}

func (*Range) SetStart(refNode *node.Node, offset int) {
	js.Rewrite("$<.setStart($1, $2)", refNode, offset)
}

func (*Range) SetStartAfter(refNode *node.Node) {
	js.Rewrite("$<.setStartAfter($1)", refNode)
}

func (*Range) SetStartBefore(refNode *node.Node) {
	js.Rewrite("$<.setStartBefore($1)", refNode)
}

func (*Range) SurroundContents(newParent *node.Node) {
	js.Rewrite("$<.surroundContents($1)", newParent)
}

func (*Range) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*Range) GetCollapsed() (collapsed bool) {
	js.Rewrite("$<.collapsed")
	return collapsed
}

func (*Range) GetCommonAncestorContainer() (commonAncestorContainer *node.Node) {
	js.Rewrite("$<.commonAncestorContainer")
	return commonAncestorContainer
}

func (*Range) GetEndContainer() (endContainer *node.Node) {
	js.Rewrite("$<.endContainer")
	return endContainer
}

func (*Range) GetEndOffset() (endOffset int) {
	js.Rewrite("$<.endOffset")
	return endOffset
}

func (*Range) GetStartContainer() (startContainer *node.Node) {
	js.Rewrite("$<.startContainer")
	return startContainer
}

func (*Range) GetStartOffset() (startOffset int) {
	js.Rewrite("$<.startOffset")
	return startOffset
}
