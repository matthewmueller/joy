package idbcursorwithvalue

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/idbcursor"
	"github.com/matthewmueller/golly/js"
)

type IDBCursorWithValue struct {
	*idbcursor.IDBCursor
}

func (*IDBCursorWithValue) GetValue() (value interface{}) {
	js.Rewrite("$<.value")
	return value
}
