package idbcursorwithvalue

import (
	"github.com/matthewmueller/golly/dom/idbcursor"
	"github.com/matthewmueller/golly/js"
)

// IDBCursorWithValue struct
// js:"IDBCursorWithValue,omit"
type IDBCursorWithValue struct {
	idbcursor.IDBCursor
}

// Value prop
func (*IDBCursorWithValue) Value() (value interface{}) {
	js.Rewrite("$<.value")
	return value
}
