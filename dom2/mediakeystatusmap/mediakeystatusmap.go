package mediakeystatusmap

import (
	"github.com/matthewmueller/golly/dom2/mediakeystatus"
	"github.com/matthewmueller/golly/js"
)

// js:"MediaKeyStatusMap,omit"
type MediaKeyStatusMap struct {
}

// ForEach
func (*MediaKeyStatusMap) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$<.ForEach($1)", callback)
}

// Get
func (*MediaKeyStatusMap) Get(keyId []byte) (m *mediakeystatus.MediaKeyStatus) {
	js.Rewrite("$<.Get($1)", keyId)
	return m
}

// Has
func (*MediaKeyStatusMap) Has(keyId []byte) (b bool) {
	js.Rewrite("$<.Has($1)", keyId)
	return b
}

// Size
func (*MediaKeyStatusMap) Size() (size uint) {
	js.Rewrite("$<.Size")
	return size
}
