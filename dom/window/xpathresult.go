package window

import "github.com/matthewmueller/golly/js"

// XPathResult struct
// js:"XPathResult,omit"
type XPathResult struct {
}

// IterateNext fn
// js:"iterateNext"
func (*XPathResult) IterateNext() (n Node) {
	js.Rewrite("$_.iterateNext()")
	return n
}

// SnapshotItem fn
// js:"snapshotItem"
func (*XPathResult) SnapshotItem(index uint) (n Node) {
	js.Rewrite("$_.snapshotItem($1)", index)
	return n
}

// BooleanValue prop
// js:"booleanValue"
func (*XPathResult) BooleanValue() (booleanValue bool) {
	js.Rewrite("$_.booleanValue")
	return booleanValue
}

// InvalidIteratorState prop
// js:"invalidIteratorState"
func (*XPathResult) InvalidIteratorState() (invalidIteratorState bool) {
	js.Rewrite("$_.invalidIteratorState")
	return invalidIteratorState
}

// NumberValue prop
// js:"numberValue"
func (*XPathResult) NumberValue() (numberValue float32) {
	js.Rewrite("$_.numberValue")
	return numberValue
}

// ResultType prop
// js:"resultType"
func (*XPathResult) ResultType() (resultType uint8) {
	js.Rewrite("$_.resultType")
	return resultType
}

// SingleNodeValue prop
// js:"singleNodeValue"
func (*XPathResult) SingleNodeValue() (singleNodeValue Node) {
	js.Rewrite("$_.singleNodeValue")
	return singleNodeValue
}

// SnapshotLength prop
// js:"snapshotLength"
func (*XPathResult) SnapshotLength() (snapshotLength uint) {
	js.Rewrite("$_.snapshotLength")
	return snapshotLength
}

// StringValue prop
// js:"stringValue"
func (*XPathResult) StringValue() (stringValue string) {
	js.Rewrite("$_.stringValue")
	return stringValue
}
