package window

import "github.com/matthewmueller/golly/js"

// js:"IDBFactory,omit"
type IDBFactory struct {
}

// Cmp
func (*IDBFactory) Cmp(first interface{}, second interface{}) (i int8) {
	js.Rewrite("$<.Cmp($1, $2)", first, second)
	return i
}

// DeleteDatabase
func (*IDBFactory) DeleteDatabase(name string) (i *IDBOpenDBRequest) {
	js.Rewrite("$<.DeleteDatabase($1)", name)
	return i
}

// Open
func (*IDBFactory) Open(name string, version *uint64) (i *IDBOpenDBRequest) {
	js.Rewrite("$<.Open($1, $2)", name, version)
	return i
}
