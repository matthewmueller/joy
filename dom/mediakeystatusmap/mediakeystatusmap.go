package mediakeystatusmap

import (
	"github.com/matthewmueller/golly/dom2/mediakeystatus"
	"github.com/matthewmueller/golly/js"
)

// MediaKeyStatusMap struct
// js:"MediaKeyStatusMap,omit"
type MediaKeyStatusMap struct {
}

// ForEach fn
func (*MediaKeyStatusMap) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$<.forEach($1)", callback)
}

// Get fn
func (*MediaKeyStatusMap) Get(keyId []byte) (m *mediakeystatus.MediaKeyStatus) {
	js.Rewrite("$<.get($1)", keyId)
	return m
}

// Has fn
func (*MediaKeyStatusMap) Has(keyId []byte) (b bool) {
	js.Rewrite("$<.has($1)", keyId)
	return b
}

// Size prop
func (*MediaKeyStatusMap) Size() (size uint) {
	js.Rewrite("$<.size")
	return size
}
