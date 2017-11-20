package mediakeystatusmap

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeystatus"
	"github.com/matthewmueller/golly/js"
)

type MediaKeyStatusMap struct {
}

func (*MediaKeyStatusMap) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$<.forEach($1)", callback)
}

func (*MediaKeyStatusMap) Get(keyId []byte) (m *mediakeystatus.MediaKeyStatus) {
	js.Rewrite("$<.get($1)", keyId)
	return m
}

func (*MediaKeyStatusMap) Has(keyId []byte) (b bool) {
	js.Rewrite("$<.has($1)", keyId)
	return b
}

func (*MediaKeyStatusMap) GetSize() (size uint) {
	js.Rewrite("$<.size")
	return size
}
