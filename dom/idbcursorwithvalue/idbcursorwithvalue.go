package idbcursorwithvalue

import (
	"github.com/matthewmueller/joy/dom/idbcursor"
	"github.com/matthewmueller/joy/dom/idbcursordirection"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ idbcursor.IDBCursor = (*IDBCursorWithValue)(nil)

// IDBCursorWithValue struct
// js:"IDBCursorWithValue,omit"
type IDBCursorWithValue struct {
}

// Advance fn
// js:"advance"
func (*IDBCursorWithValue) Advance(count uint) {
	macro.Rewrite("$_.advance($1)", count)
}

// Continue fn
// js:"continue"
func (*IDBCursorWithValue) Continue(key *interface{}) {
	macro.Rewrite("$_.continue($1)", key)
}

// Delete fn
// js:"delete"
func (*IDBCursorWithValue) Delete() (w window.IDBRequest) {
	macro.Rewrite("$_.delete()")
	return w
}

// Update fn
// js:"update"
func (*IDBCursorWithValue) Update(value interface{}) (w window.IDBRequest) {
	macro.Rewrite("$_.update($1)", value)
	return w
}

// Value prop
// js:"value"
func (*IDBCursorWithValue) Value() (value interface{}) {
	macro.Rewrite("$_.value")
	return value
}

// Direction prop
// js:"direction"
func (*IDBCursorWithValue) Direction() (direction *idbcursordirection.IDBCursorDirection) {
	macro.Rewrite("$_.direction")
	return direction
}

// Key prop
// js:"key"
func (*IDBCursorWithValue) Key() (key interface{}) {
	macro.Rewrite("$_.key")
	return key
}

// PrimaryKey prop
// js:"primaryKey"
func (*IDBCursorWithValue) PrimaryKey() (primaryKey interface{}) {
	macro.Rewrite("$_.primaryKey")
	return primaryKey
}

// Source prop
// js:"source"
func (*IDBCursorWithValue) Source() (source interface{}) {
	macro.Rewrite("$_.source")
	return source
}
