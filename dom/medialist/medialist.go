package medialist

import "github.com/matthewmueller/golly/js"

// MediaList struct
// js:"MediaList,omit"
type MediaList struct {
}

// AppendMedium fn
func (*MediaList) AppendMedium(newMedium string) {
	js.Rewrite("$<.appendMedium($1)", newMedium)
}

// DeleteMedium fn
func (*MediaList) DeleteMedium(oldMedium string) {
	js.Rewrite("$<.deleteMedium($1)", oldMedium)
}

// Item fn
func (*MediaList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

// ToString fn
func (*MediaList) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Length prop
func (*MediaList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// MediaText prop
func (*MediaList) MediaText() (mediaText string) {
	js.Rewrite("$<.mediaText")
	return mediaText
}

// MediaText prop
func (*MediaList) SetMediaText(mediaText string) {
	js.Rewrite("$<.mediaText = $1", mediaText)
}
