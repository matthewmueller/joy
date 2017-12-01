package mediakeystatusmap

import (
	"github.com/matthewmueller/golly/dom/mediakeystatus"
	"github.com/matthewmueller/golly/js"
)

// MediaKeyStatusMap struct
// js:"MediaKeyStatusMap,omit"
type MediaKeyStatusMap struct {
}

// ForEach fn
// js:"forEach"
func (*MediaKeyStatusMap) ForEach(callback func(keyId []byte, status *mediakeystatus.MediaKeyStatus)) {
	js.Rewrite("$_.forEach($1)", callback)
}

// Get fn
// js:"get"
func (*MediaKeyStatusMap) Get(keyId []byte) (m *mediakeystatus.MediaKeyStatus) {
	js.Rewrite("$_.get($1)", keyId)
	return m
}

// Has fn
// js:"has"
func (*MediaKeyStatusMap) Has(keyId []byte) (b bool) {
	js.Rewrite("$_.has($1)", keyId)
	return b
}

// Size prop
// js:"size"
func (*MediaKeyStatusMap) Size() (size uint) {
	js.Rewrite("$_.size")
	return size
}
