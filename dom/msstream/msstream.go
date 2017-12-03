package msstream

import "github.com/matthewmueller/joy/macro"

// MSStream struct
// js:"MSStream,omit"
type MSStream struct {
}

// MsClose fn
// js:"msClose"
func (*MSStream) MsClose() {
	macro.Rewrite("$_.msClose()")
}

// MsDetachStream fn
// js:"msDetachStream"
func (*MSStream) MsDetachStream() (i interface{}) {
	macro.Rewrite("$_.msDetachStream()")
	return i
}

// Type prop
// js:"type"
func (*MSStream) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
