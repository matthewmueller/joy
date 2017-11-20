package msstream

import "github.com/matthewmueller/golly/js"

type MSStream struct {
}

func (*MSStream) MsClose() {
	js.Rewrite("$<.msClose()")
}

func (*MSStream) MsDetachStream() (i interface{}) {
	js.Rewrite("$<.msDetachStream()")
	return i
}

func (*MSStream) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
