package idbobjectstore

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domstringlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbcursordirection"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbindex"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbindexparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbrequest"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbtransaction"
	"github.com/matthewmueller/golly/js"
)

type IDBObjectStore struct {
}

func (*IDBObjectStore) Add(value interface{}, key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.add($1, $2)", value, key)
	return i
}

func (*IDBObjectStore) Clear() (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.clear()")
	return i
}

func (*IDBObjectStore) Count(key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.count($1)", key)
	return i
}

func (*IDBObjectStore) CreateIndex(name string, keyPath string, optionalParameters *idbindexparameters.IDBIndexParameters) (i *idbindex.IDBIndex) {
	js.Rewrite("$<.createIndex($1, $2, $3)", name, keyPath, optionalParameters)
	return i
}

func (*IDBObjectStore) Delete(key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.delete($1)", key)
	return i
}

func (*IDBObjectStore) DeleteIndex(indexName string) {
	js.Rewrite("$<.deleteIndex($1)", indexName)
}

func (*IDBObjectStore) Get(key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.get($1)", key)
	return i
}

func (*IDBObjectStore) Index(name string) (i *idbindex.IDBIndex) {
	js.Rewrite("$<.index($1)", name)
	return i
}

func (*IDBObjectStore) OpenCursor(rng interface{}, direction *idbcursordirection.IDBCursorDirection) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.openCursor($1, $2)", rng, direction)
	return i
}

func (*IDBObjectStore) Put(value interface{}, key interface{}) (i *idbrequest.IDBRequest) {
	js.Rewrite("$<.put($1, $2)", value, key)
	return i
}

func (*IDBObjectStore) GetIndexNames() (indexNames *domstringlist.DOMStringList) {
	js.Rewrite("$<.indexNames")
	return indexNames
}

func (*IDBObjectStore) GetKeyPath() (keyPath string) {
	js.Rewrite("$<.keyPath")
	return keyPath
}

func (*IDBObjectStore) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*IDBObjectStore) GetTransaction() (transaction *idbtransaction.IDBTransaction) {
	js.Rewrite("$<.transaction")
	return transaction
}
