package idbindex

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/idbcursordirection"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbobjectstore"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbrequest"
	"github.com/matthewmueller/golly/js"
)

type IDBIndex struct {
}

func (*IDBIndex) Count(key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.count($1)", key)
	return i
}

func (*IDBIndex) Get(key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.get($1)", key)
	return i
}

func (*IDBIndex) GetKey(key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.getKey($1)", key)
	return i
}

func (*IDBIndex) OpenCursor(rng interface{}, direction *idbcursordirection.IDBCursorDirection) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.openCursor($1, $2)", rng, direction)
	return i
}

func (*IDBIndex) OpenKeyCursor(rng interface{}, direction *idbcursordirection.IDBCursorDirection) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.openKeyCursor($1, $2)", rng, direction)
	return i
}

func (*IDBIndex) GetKeyPath() (keyPath string) {
	js.Rewrite("$<.keyPath")
	return keyPath
}

func (*IDBIndex) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*IDBIndex) GetObjectStore() (objectStore *idbobjectstore.IDBObjectStore) {
	js.Rewrite("$<.objectStore")
	return objectStore
}

func (*IDBIndex) GetUnique() (unique bool) {
	js.Rewrite("$<.unique")
	return unique
}
