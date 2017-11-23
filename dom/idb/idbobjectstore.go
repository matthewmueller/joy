package idb

import (
	"github.com/matthewmueller/golly/dom2/domstringlist"
	"github.com/matthewmueller/golly/dom2/idbcursordirection"
	"github.com/matthewmueller/golly/dom2/idbindexparameters"
	"github.com/matthewmueller/golly/js"
)

// IDBObjectStore struct
// js:"IDBObjectStore,omit"
type IDBObjectStore struct {
}

// Add
func (*IDBObjectStore) Add(value interface{}, key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.Add($1, $2)", value, key)
	return i
}

// Clear
func (*IDBObjectStore) Clear() (i IDBRequest) {
	js.Rewrite("$<.Clear()")
	return i
}

// Count
func (*IDBObjectStore) Count(key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.Count($1)", key)
	return i
}

// CreateIndex
func (*IDBObjectStore) CreateIndex(name string, keyPath string, optionalParameters *idbindexparameters.IDBIndexParameters) (i *IDBIndex) {
	js.Rewrite("$<.CreateIndex($1, $2, $3)", name, keyPath, optionalParameters)
	return i
}

// Delete
func (*IDBObjectStore) Delete(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.Delete($1)", key)
	return i
}

// DeleteIndex
func (*IDBObjectStore) DeleteIndex(indexName string) {
	js.Rewrite("$<.DeleteIndex($1)", indexName)
}

// Get
func (*IDBObjectStore) Get(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.Get($1)", key)
	return i
}

// Index
func (*IDBObjectStore) Index(name string) (i *IDBIndex) {
	js.Rewrite("$<.Index($1)", name)
	return i
}

// OpenCursor
func (*IDBObjectStore) OpenCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	js.Rewrite("$<.OpenCursor($1, $2)", rng, direction)
	return i
}

// Put
func (*IDBObjectStore) Put(value interface{}, key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.Put($1, $2)", value, key)
	return i
}

// IndexNames
func (*IDBObjectStore) IndexNames() (indexNames *domstringlist.DOMStringList) {
	js.Rewrite("$<.IndexNames")
	return indexNames
}

// KeyPath
func (*IDBObjectStore) KeyPath() (keyPath string) {
	js.Rewrite("$<.KeyPath")
	return keyPath
}

// Name
func (*IDBObjectStore) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Transaction
func (*IDBObjectStore) Transaction() (transaction *IDBTransaction) {
	js.Rewrite("$<.Transaction")
	return transaction
}
