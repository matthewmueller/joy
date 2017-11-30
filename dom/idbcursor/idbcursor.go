package idbcursor

import (
	"github.com/matthewmueller/golly/dom/idbcursordirection"
	"github.com/matthewmueller/golly/dom/window"
)

// IDBCursor interface
// js:"IDBCursor"
type IDBCursor interface {

	// Advance
	// js:"advance"
	Advance(count uint)

	// Continue
	// js:"continue"
	Continue(key *interface{})

	// Delete
	// js:"delete"
	Delete() (w window.IDBRequest)

	// Update
	// js:"update"
	Update(value interface{}) (w window.IDBRequest)

	// Direction prop
	// js:"direction"
	Direction() (direction *idbcursordirection.IDBCursorDirection)

	// Key prop
	// js:"key"
	Key() (key interface{})

	// PrimaryKey prop
	// js:"primaryKey"
	PrimaryKey() (primaryKey interface{})

	// Source prop
	// js:"source"
	Source() (source interface{})
}
