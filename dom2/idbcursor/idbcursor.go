package idbcursor

import "github.com/matthewmueller/golly/dom2/idbcursordirection"

// js:"IDBCursor,omit"
type IDBCursor interface {

	// Advance
	Advance(count uint)

	// Continue
	Continue(key *interface{})

	// Delete
	Delete() (w window.IDBRequest)

	// Update
	Update(value interface{}) (w window.IDBRequest)

	// Direction
	Direction() (direction *idbcursordirection.IDBCursorDirection)

	// Key
	Key() (key interface{})

	// PrimaryKey
	PrimaryKey() (primaryKey interface{})

	// Source
	Source() (source interface{})
}
