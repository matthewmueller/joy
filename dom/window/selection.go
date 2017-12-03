package window

import "github.com/matthewmueller/joy/macro"

// Selection struct
// js:"Selection,omit"
type Selection struct {
}

// AddRange fn
// js:"addRange"
func (*Selection) AddRange(rng *Range) {
	macro.Rewrite("$_.addRange($1)", rng)
}

// Collapse fn
// js:"collapse"
func (*Selection) Collapse(parentNode Node, offset int) {
	macro.Rewrite("$_.collapse($1, $2)", parentNode, offset)
}

// CollapseToEnd fn
// js:"collapseToEnd"
func (*Selection) CollapseToEnd() {
	macro.Rewrite("$_.collapseToEnd()")
}

// CollapseToStart fn
// js:"collapseToStart"
func (*Selection) CollapseToStart() {
	macro.Rewrite("$_.collapseToStart()")
}

// ContainsNode fn
// js:"containsNode"
func (*Selection) ContainsNode(node Node, partlyContained bool) (b bool) {
	macro.Rewrite("$_.containsNode($1, $2)", node, partlyContained)
	return b
}

// DeleteFromDocument fn
// js:"deleteFromDocument"
func (*Selection) DeleteFromDocument() {
	macro.Rewrite("$_.deleteFromDocument()")
}

// Empty fn
// js:"empty"
func (*Selection) Empty() {
	macro.Rewrite("$_.empty()")
}

// Extend fn
// js:"extend"
func (*Selection) Extend(newNode Node, offset int) {
	macro.Rewrite("$_.extend($1, $2)", newNode, offset)
}

// GetRangeAt fn
// js:"getRangeAt"
func (*Selection) GetRangeAt(index int) (r *Range) {
	macro.Rewrite("$_.getRangeAt($1)", index)
	return r
}

// RemoveAllRanges fn
// js:"removeAllRanges"
func (*Selection) RemoveAllRanges() {
	macro.Rewrite("$_.removeAllRanges()")
}

// RemoveRange fn
// js:"removeRange"
func (*Selection) RemoveRange(rng *Range) {
	macro.Rewrite("$_.removeRange($1)", rng)
}

// SelectAllChildren fn
// js:"selectAllChildren"
func (*Selection) SelectAllChildren(parentNode Node) {
	macro.Rewrite("$_.selectAllChildren($1)", parentNode)
}

// SetBaseAndExtent fn
// js:"setBaseAndExtent"
func (*Selection) SetBaseAndExtent(baseNode Node, baseOffset int, extentNode Node, extentOffset int) {
	macro.Rewrite("$_.setBaseAndExtent($1, $2, $3, $4)", baseNode, baseOffset, extentNode, extentOffset)
}

// SetPosition fn
// js:"setPosition"
func (*Selection) SetPosition(parentNode Node, offset int) {
	macro.Rewrite("$_.setPosition($1, $2)", parentNode, offset)
}

// ToString fn
// js:"toString"
func (*Selection) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// AnchorNode prop
// js:"anchorNode"
func (*Selection) AnchorNode() (anchorNode Node) {
	macro.Rewrite("$_.anchorNode")
	return anchorNode
}

// AnchorOffset prop
// js:"anchorOffset"
func (*Selection) AnchorOffset() (anchorOffset int) {
	macro.Rewrite("$_.anchorOffset")
	return anchorOffset
}

// BaseNode prop
// js:"baseNode"
func (*Selection) BaseNode() (baseNode Node) {
	macro.Rewrite("$_.baseNode")
	return baseNode
}

// BaseOffset prop
// js:"baseOffset"
func (*Selection) BaseOffset() (baseOffset int) {
	macro.Rewrite("$_.baseOffset")
	return baseOffset
}

// ExtentNode prop
// js:"extentNode"
func (*Selection) ExtentNode() (extentNode Node) {
	macro.Rewrite("$_.extentNode")
	return extentNode
}

// ExtentOffset prop
// js:"extentOffset"
func (*Selection) ExtentOffset() (extentOffset int) {
	macro.Rewrite("$_.extentOffset")
	return extentOffset
}

// FocusNode prop
// js:"focusNode"
func (*Selection) FocusNode() (focusNode Node) {
	macro.Rewrite("$_.focusNode")
	return focusNode
}

// FocusOffset prop
// js:"focusOffset"
func (*Selection) FocusOffset() (focusOffset int) {
	macro.Rewrite("$_.focusOffset")
	return focusOffset
}

// IsCollapsed prop
// js:"isCollapsed"
func (*Selection) IsCollapsed() (isCollapsed bool) {
	macro.Rewrite("$_.isCollapsed")
	return isCollapsed
}

// RangeCount prop
// js:"rangeCount"
func (*Selection) RangeCount() (rangeCount int) {
	macro.Rewrite("$_.rangeCount")
	return rangeCount
}

// Type prop
// js:"type"
func (*Selection) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
