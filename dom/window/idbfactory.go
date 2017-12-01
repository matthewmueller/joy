package window

import "github.com/matthewmueller/golly/js"

// IDBFactory struct
// js:"IDBFactory,omit"
type IDBFactory struct {
}

// Cmp fn
// js:"cmp"
func (*IDBFactory) Cmp(first interface{}, second interface{}) (i int8) {
	js.Rewrite("$_.cmp($1, $2)", first, second)
	return i
}

// DeleteDatabase fn
// js:"deleteDatabase"
func (*IDBFactory) DeleteDatabase(name string) (i *IDBOpenDBRequest) {
	js.Rewrite("$_.deleteDatabase($1)", name)
	return i
}

// Open fn
// js:"open"
func (*IDBFactory) Open(name string, version *uint64) (i *IDBOpenDBRequest) {
	js.Rewrite("$_.open($1, $2)", name, version)
	return i
}
