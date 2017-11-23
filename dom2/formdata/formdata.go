package formdata

import "github.com/matthewmueller/golly/js"

// New fn
func New() *FormData {
	js.Rewrite("FormData")
	return &FormData{}
}

// FormData struct
// js:"FormData,omit"
type FormData struct {
}

// Append
func (*FormData) Append(name interface{}, value interface{}, blobName *string) {
	js.Rewrite("$<.Append($1, $2, $3)", name, value, blobName)
}
