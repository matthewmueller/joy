package idb

import "github.com/matthewmueller/golly/js"

// IDBEnvironment struct
// js:"IDBEnvironment,omit"
type IDBEnvironment struct {
}

// IndexedDB
func (*IDBEnvironment) IndexedDB() (indexedDB *IDBFactory) {
	js.Rewrite("$<.IndexedDB")
	return indexedDB
}
