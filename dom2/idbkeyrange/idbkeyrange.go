package idbkeyrange

import "github.com/matthewmueller/golly/js"

// js:"IDBKeyRange,omit"
type IDBKeyRange struct {
}

// Bound
func (*IDBKeyRange) Bound(lower interface{}, upper interface{}, lowerOpen *bool, upperOpen *bool) (i *IDBKeyRange) {
	js.Rewrite("$<.Bound($1, $2, $3, $4)", lower, upper, lowerOpen, upperOpen)
	return i
}

// LowerBound
func (*IDBKeyRange) LowerBound(lower interface{}, open *bool) (i *IDBKeyRange) {
	js.Rewrite("$<.LowerBound($1, $2)", lower, open)
	return i
}

// Only
func (*IDBKeyRange) Only(value interface{}) (i *IDBKeyRange) {
	js.Rewrite("$<.Only($1)", value)
	return i
}

// UpperBound
func (*IDBKeyRange) UpperBound(upper interface{}, open *bool) (i *IDBKeyRange) {
	js.Rewrite("$<.UpperBound($1, $2)", upper, open)
	return i
}

// Lower
func (*IDBKeyRange) Lower() (lower interface{}) {
	js.Rewrite("$<.Lower")
	return lower
}

// LowerOpen
func (*IDBKeyRange) LowerOpen() (lowerOpen bool) {
	js.Rewrite("$<.LowerOpen")
	return lowerOpen
}

// Upper
func (*IDBKeyRange) Upper() (upper interface{}) {
	js.Rewrite("$<.Upper")
	return upper
}

// UpperOpen
func (*IDBKeyRange) UpperOpen() (upperOpen bool) {
	js.Rewrite("$<.UpperOpen")
	return upperOpen
}
