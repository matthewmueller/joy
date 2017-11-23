package msstream

import "github.com/matthewmueller/golly/js"

// MSStream struct
// js:"MSStream,omit"
type MSStream struct {
}

// MsClose fn
func (*MSStream) MsClose() {
	js.Rewrite("$<.msClose()")
}

// MsDetachStream fn
func (*MSStream) MsDetachStream() (i interface{}) {
	js.Rewrite("$<.msDetachStream()")
	return i
}

// Type prop
func (*MSStream) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
