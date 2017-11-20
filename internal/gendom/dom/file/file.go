package file

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/blob"
	"github.com/matthewmueller/golly/js"
)

type File struct {
	*blob.Blob
}

func (*File) GetLastModifiedDate() (lastModifiedDate interface{}) {
	js.Rewrite("$<.lastModifiedDate")
	return lastModifiedDate
}

func (*File) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*File) GetWebkitRelativePath() (webkitRelativePath string) {
	js.Rewrite("$<.webkitRelativePath")
	return webkitRelativePath
}
