package formdata

import "github.com/matthewmueller/golly/js"

type FormData struct {
}

func (*FormData) Append(name interface{}, value interface{}, blobName string) {
	js.Rewrite("$<.append($1, $2, $3)", name, value, blobName)
}
