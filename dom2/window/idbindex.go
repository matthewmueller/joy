package window

import "github.com/matthewmueller/golly/js"

// js:"IDBIndex,omit"
type IDBIndex struct {
}

// Count
func (*IDBIndex) Count(key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.Count($1)", key)
	return i
}

// Get
func (*IDBIndex) Get(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.Get($1)", key)
	return i
}

// GetKey
func (*IDBIndex) GetKey(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.GetKey($1)", key)
	return i
}

// OpenCursor
func (*IDBIndex) OpenCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	js.Rewrite("$<.OpenCursor($1, $2)", rng, direction)
	return i
}

// OpenKeyCursor
func (*IDBIndex) OpenKeyCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	js.Rewrite("$<.OpenKeyCursor($1, $2)", rng, direction)
	return i
}

// KeyPath
func (*IDBIndex) KeyPath() (keyPath string) {
	js.Rewrite("$<.KeyPath")
	return keyPath
}

// Name
func (*IDBIndex) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// ObjectStore
func (*IDBIndex) ObjectStore() (objectStore *IDBObjectStore) {
	js.Rewrite("$<.ObjectStore")
	return objectStore
}

// Unique
func (*IDBIndex) Unique() (unique bool) {
	js.Rewrite("$<.Unique")
	return unique
}
