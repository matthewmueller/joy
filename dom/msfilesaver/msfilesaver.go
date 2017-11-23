package msfilesaver

import "github.com/matthewmueller/golly/js"

// MSFileSaver struct
// js:"MSFileSaver,omit"
type MSFileSaver struct {
}

// MsSaveBlob fn
func (*MSFileSaver) MsSaveBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.msSaveBlob($1, $2)", blob, defaultName)
	return b
}

// MsSaveOrOpenBlob fn
func (*MSFileSaver) MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.msSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}
