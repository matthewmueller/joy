package msfilesaver

import "github.com/matthewmueller/golly/js"

type MSFileSaver struct {
}

func (*MSFileSaver) MsSaveBlob(blob interface{}, defaultName string) (b bool) {
	js.Rewrite("$<.msSaveBlob($1, $2)", blob, defaultName)
	return b
}

func (*MSFileSaver) MsSaveOrOpenBlob(blob interface{}, defaultName string) (b bool) {
	js.Rewrite("$<.msSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}
