package idbenvironment

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/idbfactory"
	"github.com/matthewmueller/golly/js"
)

type IDBEnvironment struct {
}

func (*IDBEnvironment) GetIndexedDB() (indexedDB *idbfactory.IDBFactory) {
	js.Rewrite("$<.indexedDB")
	return indexedDB
}
