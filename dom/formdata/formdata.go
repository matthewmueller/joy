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

// Append fn
// js:"append"
func (*FormData) Append(name interface{}, value interface{}, blobName *string) {
	js.Rewrite("$_.append($1, $2, $3)", name, value, blobName)
}
