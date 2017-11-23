package msfilesaver

import "github.com/matthewmueller/golly/js"

// MSFileSaver struct
// js:"MSFileSaver,omit"
type MSFileSaver struct {
}

// MsSaveBlob
func (*MSFileSaver) MsSaveBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.MsSaveBlob($1, $2)", blob, defaultName)
	return b
}

// MsSaveOrOpenBlob
func (*MSFileSaver) MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.MsSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}
