package window

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

// Add fn
func (*IDBObjectStore) Add(value interface{}, key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.add($1, $2)", value, key)
	return i
}

// Clear fn
func (*IDBObjectStore) Clear() (i IDBRequest) {
	js.Rewrite("$<.clear()")
	return i
}

// Count fn
func (*IDBObjectStore) Count(key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.count($1)", key)
	return i
}

// CreateIndex fn
func (*IDBObjectStore) CreateIndex(name string, keyPath string, optionalParameters *idbindexparameters.IDBIndexParameters) (i *IDBIndex) {
	js.Rewrite("$<.createIndex($1, $2, $3)", name, keyPath, optionalParameters)
	return i
}

// Delete fn
func (*IDBObjectStore) Delete(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.delete($1)", key)
	return i
}

// DeleteIndex fn
func (*IDBObjectStore) DeleteIndex(indexName string) {
	js.Rewrite("$<.deleteIndex($1)", indexName)
}

// Get fn
func (*IDBObjectStore) Get(key interface{}) (i IDBRequest) {
	js.Rewrite("$<.get($1)", key)
	return i
}

// Index fn
func (*IDBObjectStore) Index(name string) (i *IDBIndex) {
	js.Rewrite("$<.index($1)", name)
	return i
}

// OpenCursor fn
func (*IDBObjectStore) OpenCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	js.Rewrite("$<.openCursor($1, $2)", rng, direction)
	return i
}

// Put fn
func (*IDBObjectStore) Put(value interface{}, key *interface{}) (i IDBRequest) {
	js.Rewrite("$<.put($1, $2)", value, key)
	return i
}

// IndexNames prop
func (*IDBObjectStore) IndexNames() (indexNames *domstringlist.DOMStringList) {
	js.Rewrite("$<.indexNames")
	return indexNames
}

// KeyPath prop
func (*IDBObjectStore) KeyPath() (keyPath string) {
	js.Rewrite("$<.keyPath")
	return keyPath
}

// Name prop
func (*IDBObjectStore) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Transaction prop
func (*IDBObjectStore) Transaction() (transaction *IDBTransaction) {
	js.Rewrite("$<.transaction")
	return transaction
}
