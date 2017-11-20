package idbfactory

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/idbopendbrequest"
	"github.com/matthewmueller/golly/js"
)

type IDBFactory struct {
}

func (*IDBFactory) Cmp(first interface{}, second interface{}) (i int8) {
	js.Rewrite("$<.cmp($1, $2)", first, second)
	return i
}

func (*IDBFactory) DeleteDatabase(name string) (i *idbopendbrequest.IDBOpenDBRequest) {
	js.Rewrite("$<.deleteDatabase($1)", name)
	return i
}

func (*IDBFactory) Open(name string, version uint64) (i *idbopendbrequest.IDBOpenDBRequest) {
	js.Rewrite("$<.open($1, $2)", name, version)
	return i
}
