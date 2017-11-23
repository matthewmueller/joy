package window

import "github.com/matthewmueller/golly/js"

// IDBFactory struct
// js:"IDBFactory,omit"
type IDBFactory struct {
}

// Cmp fn
func (*IDBFactory) Cmp(first interface{}, second interface{}) (i int8) {
	js.Rewrite("$<.cmp($1, $2)", first, second)
	return i
}

// DeleteDatabase fn
func (*IDBFactory) DeleteDatabase(name string) (i *IDBOpenDBRequest) {
	js.Rewrite("$<.deleteDatabase($1)", name)
	return i
}

// Open fn
func (*IDBFactory) Open(name string, version *uint64) (i *IDBOpenDBRequest) {
	js.Rewrite("$<.open($1, $2)", name, version)
	return i
}
