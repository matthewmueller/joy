package window

import "github.com/matthewmueller/joy/macro"

// XPathResult struct
// js:"XPathResult,omit"
type XPathResult struct {
}

// IterateNext fn
// js:"iterateNext"
func (*XPathResult) IterateNext() (n Node) {
	macro.Rewrite("$_.iterateNext()")
	return n
}

// SnapshotItem fn
// js:"snapshotItem"
func (*XPathResult) SnapshotItem(index uint) (n Node) {
	macro.Rewrite("$_.snapshotItem($1)", index)
	return n
}

// BooleanValue prop
// js:"booleanValue"
func (*XPathResult) BooleanValue() (booleanValue bool) {
	macro.Rewrite("$_.booleanValue")
	return booleanValue
}

// InvalidIteratorState prop
// js:"invalidIteratorState"
func (*XPathResult) InvalidIteratorState() (invalidIteratorState bool) {
	macro.Rewrite("$_.invalidIteratorState")
	return invalidIteratorState
}

// NumberValue prop
// js:"numberValue"
func (*XPathResult) NumberValue() (numberValue float32) {
	macro.Rewrite("$_.numberValue")
	return numberValue
}

// ResultType prop
// js:"resultType"
func (*XPathResult) ResultType() (resultType uint8) {
	macro.Rewrite("$_.resultType")
	return resultType
}

// SingleNodeValue prop
// js:"singleNodeValue"
func (*XPathResult) SingleNodeValue() (singleNodeValue Node) {
	macro.Rewrite("$_.singleNodeValue")
	return singleNodeValue
}

// SnapshotLength prop
// js:"snapshotLength"
func (*XPathResult) SnapshotLength() (snapshotLength uint) {
	macro.Rewrite("$_.snapshotLength")
	return snapshotLength
}

// StringValue prop
// js:"stringValue"
func (*XPathResult) StringValue() (stringValue string) {
	macro.Rewrite("$_.stringValue")
	return stringValue
}
