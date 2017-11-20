package headers

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeystatus"
	"github.com/matthewmueller/golly/js"
)

type Headers struct {
}

func (*Headers) Append(name string, value string) {
	js.Rewrite("$<.append($1, $2)", name, value)
}

func (*Headers) Delete(name string) {
	js.Rewrite("$<.delete($1)", name)
}

func (*Headers) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$<.forEach($1)", callback)
}

func (*Headers) Get(name string) (s string) {
	js.Rewrite("$<.get($1)", name)
	return s
}

func (*Headers) Has(name string) (b bool) {
	js.Rewrite("$<.has($1)", name)
	return b
}

func (*Headers) Set(name string, value string) {
	js.Rewrite("$<.set($1, $2)", name, value)
}
