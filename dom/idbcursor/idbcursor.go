package idbcursor

import (
	"github.com/matthewmueller/golly/dom/idbcursordirection"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"IDBCursor,omit"
type IDBCursor interface {

	// Advance
	// js:"advance,rewrite=advance"
	Advance(count uint)

	// Continue
	// js:"continue,rewrite=cont"
	Continue(key *interface{})

	// Delete
	// js:"delete,rewrite=del"
	Delete() (w window.IDBRequest)

	// Update
	// js:"update,rewrite=update"
	Update(value interface{}) (w window.IDBRequest)

	// Direction prop
	// js:"direction,rewrite=direction"
	Direction() (direction *idbcursordirection.IDBCursorDirection)

	// Key prop
	// js:"key,rewrite=key"
	Key() (key interface{})

	// PrimaryKey prop
	// js:"primaryKey,rewrite=primarykey"
	PrimaryKey() (primaryKey interface{})

	// Source prop
	// js:"source,rewrite=source"
	Source() (source interface{})
}

// advance fn
func advance(count uint) {
	js.Rewrite("$<.advance($1)", count)
}

// cont fn
func cont(key *interface{}) {
	js.Rewrite("$<.continue($1)", key)
}

// del fn
func del() (w window.IDBRequest) {
	js.Rewrite("$<.delete()")
	return w
}

// update fn
func update(value interface{}) (w window.IDBRequest) {
	js.Rewrite("$<.update($1)", value)
	return w
}

// direction prop
func direction() (direction *idbcursordirection.IDBCursorDirection) {
	js.Rewrite("$<.direction")
	return direction
}

// key prop
func key() (key interface{}) {
	js.Rewrite("$<.key")
	return key
}

// primarykey prop
func primarykey() (primaryKey interface{}) {
	js.Rewrite("$<.primaryKey")
	return primaryKey
}

// source prop
func source() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}
