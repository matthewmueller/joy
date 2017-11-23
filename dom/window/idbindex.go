package window

import (
	"github.com/matthewmueller/golly/dom/idbcursordirection"
	"github.com/matthewmueller/golly/js"
)

// IDBIndex struct
// js:"IDBIndex,omit"
type IDBIndex struct {
}

// Count fn
func (*IDBIndex) Count(key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.count($1)", key)
	return i
}

// Get fn
func (*IDBIndex) Get(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.get($1)", key)
	return i
}

// GetKey fn
func (*IDBIndex) GetKey(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.getKey($1)", key)
	return i
}

// OpenCursor fn
func (*IDBIndex) OpenCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	js.Rewrite("$<.openCursor($1, $2)", rng, direction)
	return i
}

// OpenKeyCursor fn
func (*IDBIndex) OpenKeyCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	js.Rewrite("$<.openKeyCursor($1, $2)", rng, direction)
	return i
}

// KeyPath prop
func (*IDBIndex) KeyPath() (keyPath string) {
	js.Rewrite("$<.keyPath")
	return keyPath
}

// Name prop
func (*IDBIndex) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// ObjectStore prop
func (*IDBIndex) ObjectStore() (objectStore *IDBObjectStore) {
	js.Rewrite("$<.objectStore")
	return objectStore
}

// Unique prop
func (*IDBIndex) Unique() (unique bool) {
	js.Rewrite("$<.unique")
	return unique
}
