package msstream

import "github.com/matthewmueller/golly/js"

// js:"MSStream,omit"
type MSStream struct {
}

// MsClose
func (*MSStream) MsClose() {
	js.Rewrite("$<.MsClose()")
}

// MsDetachStream
func (*MSStream) MsDetachStream() (i interface{}) {
	js.Rewrite("$<.MsDetachStream()")
	return i
}

// Type
func (*MSStream) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}
