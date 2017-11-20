package webkitfileentry

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/file"
	"github.com/matthewmueller/golly/internal/gendom/dom/webkitentry"
	"github.com/matthewmueller/golly/js"
)

type WebKitFileEntry struct {
	*webkitentry.WebKitEntry
}

func (*WebKitFileEntry) File(successCallback func(file *file.File), errorCallback func(err *domerror.DOMError)) {
	js.Rewrite("$<.file($1, $2)", successCallback, errorCallback)
}
