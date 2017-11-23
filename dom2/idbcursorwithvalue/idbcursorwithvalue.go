package idbcursorwithvalue

import (
	"github.com/matthewmueller/golly/dom2/idbcursor"
	"github.com/matthewmueller/golly/js"
)

// js:"IDBCursorWithValue,omit"
type IDBCursorWithValue struct {
	idbcursor.IDBCursor
}

// Value
func (*IDBCursorWithValue) Value() (value interface{}) {
	js.Rewrite("$<.Value")
	return value
}
