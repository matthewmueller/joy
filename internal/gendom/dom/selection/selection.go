package selection

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/rng"
	"github.com/matthewmueller/golly/js"
)

type Selection struct {
}

func (*Selection) AddRange(rng *rng.Range) {
	js.Rewrite("$<.addRange($1)", rng)
}

func (*Selection) Collapse(parentNode *node.Node, offset int) {
	js.Rewrite("$<.collapse($1, $2)", parentNode, offset)
}

func (*Selection) CollapseToEnd() {
	js.Rewrite("$<.collapseToEnd()")
}

func (*Selection) CollapseToStart() {
	js.Rewrite("$<.collapseToStart()")
}

func (*Selection) ContainsNode(node *node.Node, partlyContained bool) (b bool) {
	js.Rewrite("$<.containsNode($1, $2)", node, partlyContained)
	return b
}

func (*Selection) DeleteFromDocument() {
	js.Rewrite("$<.deleteFromDocument()")
}

func (*Selection) Empty() {
	js.Rewrite("$<.empty()")
}

func (*Selection) Extend(newNode *node.Node, offset int) {
	js.Rewrite("$<.extend($1, $2)", newNode, offset)
}

func (*Selection) GetRangeAt(index int) (r *rng.Range) {
	js.Rewrite("$<.getRangeAt($1)", index)
	return r
}

func (*Selection) RemoveAllRanges() {
	js.Rewrite("$<.removeAllRanges()")
}

func (*Selection) RemoveRange(rng *rng.Range) {
	js.Rewrite("$<.removeRange($1)", rng)
}

func (*Selection) SelectAllChildren(parentNode *node.Node) {
	js.Rewrite("$<.selectAllChildren($1)", parentNode)
}

func (*Selection) SetBaseAndExtent(baseNode *node.Node, baseOffset int, extentNode *node.Node, extentOffset int) {
	js.Rewrite("$<.setBaseAndExtent($1, $2, $3, $4)", baseNode, baseOffset, extentNode, extentOffset)
}

func (*Selection) SetPosition(parentNode *node.Node, offset int) {
	js.Rewrite("$<.setPosition($1, $2)", parentNode, offset)
}

func (*Selection) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*Selection) GetAnchorNode() (anchorNode *node.Node) {
	js.Rewrite("$<.anchorNode")
	return anchorNode
}

func (*Selection) GetAnchorOffset() (anchorOffset int) {
	js.Rewrite("$<.anchorOffset")
	return anchorOffset
}

func (*Selection) GetBaseNode() (baseNode *node.Node) {
	js.Rewrite("$<.baseNode")
	return baseNode
}

func (*Selection) GetBaseOffset() (baseOffset int) {
	js.Rewrite("$<.baseOffset")
	return baseOffset
}

func (*Selection) GetExtentNode() (extentNode *node.Node) {
	js.Rewrite("$<.extentNode")
	return extentNode
}

func (*Selection) GetExtentOffset() (extentOffset int) {
	js.Rewrite("$<.extentOffset")
	return extentOffset
}

func (*Selection) GetFocusNode() (focusNode *node.Node) {
	js.Rewrite("$<.focusNode")
	return focusNode
}

func (*Selection) GetFocusOffset() (focusOffset int) {
	js.Rewrite("$<.focusOffset")
	return focusOffset
}

func (*Selection) GetIsCollapsed() (isCollapsed bool) {
	js.Rewrite("$<.isCollapsed")
	return isCollapsed
}

func (*Selection) GetRangeCount() (rangeCount int) {
	js.Rewrite("$<.rangeCount")
	return rangeCount
}

func (*Selection) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
