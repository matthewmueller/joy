package msstream

import "github.com/matthewmueller/golly/js"

// MSStream struct
// js:"MSStream,omit"
type MSStream struct {
}

// MsClose fn
// js:"msClose"
func (*MSStream) MsClose() {
	js.Rewrite("$_.msClose()")
}

// MsDetachStream fn
// js:"msDetachStream"
func (*MSStream) MsDetachStream() (i interface{}) {
	js.Rewrite("$_.msDetachStream()")
	return i
}

// Type prop
// js:"type"
func (*MSStream) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
