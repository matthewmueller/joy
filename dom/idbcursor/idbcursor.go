package idbcursor

import (
	"github.com/matthewmueller/joy/dom/idbcursordirection"
	"github.com/matthewmueller/joy/dom/window"
)

// IDBCursor interface
// js:"IDBCursor"
type IDBCursor interface {

	// Advance
	// js:"advance"
	// jsrewrite:"$_.advance($1)"
	Advance(count uint)

	// Continue
	// js:"continue"
	// jsrewrite:"$_.continue($1)"
	Continue(key *interface{})

	// Delete
	// js:"delete"
	// jsrewrite:"$_.delete()"
	Delete() (w window.IDBRequest)

	// Update
	// js:"update"
	// jsrewrite:"$_.update($1)"
	Update(value interface{}) (w window.IDBRequest)

	// Direction prop
	// js:"direction"
	// jsrewrite:"$_.direction"
	Direction() (direction *idbcursordirection.IDBCursorDirection)

	// Key prop
	// js:"key"
	// jsrewrite:"$_.key"
	Key() (key interface{})

	// PrimaryKey prop
	// js:"primaryKey"
	// jsrewrite:"$_.primaryKey"
	PrimaryKey() (primaryKey interface{})

	// Source prop
	// js:"source"
	// jsrewrite:"$_.source"
	Source() (source interface{})
}
