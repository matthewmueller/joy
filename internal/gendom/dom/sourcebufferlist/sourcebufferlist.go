package sourcebufferlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/sourcebuffer"
	"github.com/matthewmueller/golly/js"
)

type SourceBufferList struct {
	*eventtarget.EventTarget
}

func (*SourceBufferList) Item(index uint) (s *sourcebuffer.SourceBuffer) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*SourceBufferList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
