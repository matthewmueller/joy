package window

import (
	"github.com/matthewmueller/joy/dom/domstringlist"
	"github.com/matthewmueller/joy/dom/idbcursordirection"
	"github.com/matthewmueller/joy/dom/idbindexparameters"
	"github.com/matthewmueller/joy/macro"
)

// IDBObjectStore struct
// js:"IDBObjectStore,omit"
type IDBObjectStore struct {
}

// Add fn
// js:"add"
func (*IDBObjectStore) Add(value interface{}, key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.add($1, $2)", value, key)
	return i
}

// Clear fn
// js:"clear"
func (*IDBObjectStore) Clear() (i IDBRequest) {
	macro.Rewrite("$_.clear()")
	return i
}

// Count fn
// js:"count"
func (*IDBObjectStore) Count(key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.count($1)", key)
	return i
}

// CreateIndex fn
// js:"createIndex"
func (*IDBObjectStore) CreateIndex(name string, keyPath string, optionalParameters *idbindexparameters.IDBIndexParameters) (i *IDBIndex) {
	macro.Rewrite("$_.createIndex($1, $2, $3)", name, keyPath, optionalParameters)
	return i
}

// Delete fn
// js:"delete"
func (*IDBObjectStore) Delete(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.delete($1)", key)
	return i
}

// DeleteIndex fn
// js:"deleteIndex"
func (*IDBObjectStore) DeleteIndex(indexName string) {
	macro.Rewrite("$_.deleteIndex($1)", indexName)
}

// Get fn
// js:"get"
func (*IDBObjectStore) Get(key interface{}) (i IDBRequest) {
	macro.Rewrite("$_.get($1)", key)
	return i
}

// Index fn
// js:"index"
func (*IDBObjectStore) Index(name string) (i *IDBIndex) {
	macro.Rewrite("$_.index($1)", name)
	return i
}

// OpenCursor fn
// js:"openCursor"
func (*IDBObjectStore) OpenCursor(rng *interface{}, direction *idbcursordirection.IDBCursorDirection) (i IDBRequest) {
	macro.Rewrite("$_.openCursor($1, $2)", rng, direction)
	return i
}

// Put fn
// js:"put"
func (*IDBObjectStore) Put(value interface{}, key *interface{}) (i IDBRequest) {
	macro.Rewrite("$_.put($1, $2)", value, key)
	return i
}

// IndexNames prop
// js:"indexNames"
func (*IDBObjectStore) IndexNames() (indexNames *domstringlist.DOMStringList) {
	macro.Rewrite("$_.indexNames")
	return indexNames
}

// KeyPath prop
// js:"keyPath"
func (*IDBObjectStore) KeyPath() (keyPath string) {
	macro.Rewrite("$_.keyPath")
	return keyPath
}

// Name prop
// js:"name"
func (*IDBObjectStore) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// Transaction prop
// js:"transaction"
func (*IDBObjectStore) Transaction() (transaction *IDBTransaction) {
	macro.Rewrite("$_.transaction")
	return transaction
}
