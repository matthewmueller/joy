package idbkeyrange

import "github.com/matthewmueller/golly/js"

// IDBKeyRange struct
// js:"IDBKeyRange,omit"
type IDBKeyRange struct {
}

// Bound fn
func (*IDBKeyRange) Bound(lower interface{}, upper interface{}, lowerOpen *bool, upperOpen *bool) (i *IDBKeyRange) {
	js.Rewrite("$<.bound($1, $2, $3, $4)", lower, upper, lowerOpen, upperOpen)
	return i
}

// LowerBound fn
func (*IDBKeyRange) LowerBound(lower interface{}, open *bool) (i *IDBKeyRange) {
	js.Rewrite("$<.lowerBound($1, $2)", lower, open)
	return i
}

// Only fn
func (*IDBKeyRange) Only(value interface{}) (i *IDBKeyRange) {
	js.Rewrite("$<.only($1)", value)
	return i
}

// UpperBound fn
func (*IDBKeyRange) UpperBound(upper interface{}, open *bool) (i *IDBKeyRange) {
	js.Rewrite("$<.upperBound($1, $2)", upper, open)
	return i
}

// Lower prop
func (*IDBKeyRange) Lower() (lower interface{}) {
	js.Rewrite("$<.lower")
	return lower
}

// LowerOpen prop
func (*IDBKeyRange) LowerOpen() (lowerOpen bool) {
	js.Rewrite("$<.lowerOpen")
	return lowerOpen
}

// Upper prop
func (*IDBKeyRange) Upper() (upper interface{}) {
	js.Rewrite("$<.upper")
	return upper
}

// UpperOpen prop
func (*IDBKeyRange) UpperOpen() (upperOpen bool) {
	js.Rewrite("$<.upperOpen")
	return upperOpen
}
