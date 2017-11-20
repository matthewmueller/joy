package idbcursor

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/idbcursordirection"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbrequest"
	"github.com/matthewmueller/golly/js"
)

type IDBCursor struct {
}

func (*IDBCursor) Advance(count uint) {
	js.Rewrite("$<.advance($1)", count)
}

func (*IDBCursor) Continue(key interface{}) {
	js.Rewrite("$<.continue($1)", key)
}

func (*IDBCursor) Delete() (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.delete()")
	return i
}

func (*IDBCursor) Update(value interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.update($1)", value)
	return i
}

func (*IDBCursor) GetDirection() (direction *idbcursordirection.IDBCursorDirection) {
	js.Rewrite("$<.direction")
	return direction
}

func (*IDBCursor) GetKey() (key interface{}) {
	js.Rewrite("$<.key")
	return key
}

func (*IDBCursor) GetPrimaryKey() (primaryKey interface{}) {
	js.Rewrite("$<.primaryKey")
	return primaryKey
}

func (*IDBCursor) GetSource() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}
