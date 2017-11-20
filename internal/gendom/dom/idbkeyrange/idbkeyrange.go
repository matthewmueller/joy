package idbkeyrange

import "github.com/matthewmueller/golly/js"

type IDBKeyRange struct {
}

func (*IDBKeyRange) Bound(lower interface{}, upper interface{}, lowerOpen bool, upperOpen bool) (i *IDBKeyRange) {
	js.Rewrite("$<.bound($1, $2, $3, $4)", lower, upper, lowerOpen, upperOpen)
	return i
}

func (*IDBKeyRange) LowerBound(lower interface{}, open bool) (i *IDBKeyRange) {
	js.Rewrite("$<.lowerBound($1, $2)", lower, open)
	return i
}

func (*IDBKeyRange) Only(value interface{}) (i *IDBKeyRange) {
	js.Rewrite("$<.only($1)", value)
	return i
}

func (*IDBKeyRange) UpperBound(upper interface{}, open bool) (i *IDBKeyRange) {
	js.Rewrite("$<.upperBound($1, $2)", upper, open)
	return i
}

func (*IDBKeyRange) GetLower() (lower interface{}) {
	js.Rewrite("$<.lower")
	return lower
}

func (*IDBKeyRange) GetLowerOpen() (lowerOpen bool) {
	js.Rewrite("$<.lowerOpen")
	return lowerOpen
}

func (*IDBKeyRange) GetUpper() (upper interface{}) {
	js.Rewrite("$<.upper")
	return upper
}

func (*IDBKeyRange) GetUpperOpen() (upperOpen bool) {
	js.Rewrite("$<.upperOpen")
	return upperOpen
}
