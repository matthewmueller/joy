package idbkeyrange

import "github.com/matthewmueller/joy/js"

// IDBKeyRange struct
// js:"IDBKeyRange,omit"
type IDBKeyRange struct {
}

// Bound fn
// js:"bound"
func (*IDBKeyRange) Bound(lower interface{}, upper interface{}, lowerOpen *bool, upperOpen *bool) (i *IDBKeyRange) {
	js.Rewrite("$_.bound($1, $2, $3, $4)", lower, upper, lowerOpen, upperOpen)
	return i
}

// LowerBound fn
// js:"lowerBound"
func (*IDBKeyRange) LowerBound(lower interface{}, open *bool) (i *IDBKeyRange) {
	js.Rewrite("$_.lowerBound($1, $2)", lower, open)
	return i
}

// Only fn
// js:"only"
func (*IDBKeyRange) Only(value interface{}) (i *IDBKeyRange) {
	js.Rewrite("$_.only($1)", value)
	return i
}

// UpperBound fn
// js:"upperBound"
func (*IDBKeyRange) UpperBound(upper interface{}, open *bool) (i *IDBKeyRange) {
	js.Rewrite("$_.upperBound($1, $2)", upper, open)
	return i
}

// Lower prop
// js:"lower"
func (*IDBKeyRange) Lower() (lower interface{}) {
	js.Rewrite("$_.lower")
	return lower
}

// LowerOpen prop
// js:"lowerOpen"
func (*IDBKeyRange) LowerOpen() (lowerOpen bool) {
	js.Rewrite("$_.lowerOpen")
	return lowerOpen
}

// Upper prop
// js:"upper"
func (*IDBKeyRange) Upper() (upper interface{}) {
	js.Rewrite("$_.upper")
	return upper
}

// UpperOpen prop
// js:"upperOpen"
func (*IDBKeyRange) UpperOpen() (upperOpen bool) {
	js.Rewrite("$_.upperOpen")
	return upperOpen
}
