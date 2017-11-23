package medialist

import "github.com/matthewmueller/golly/js"

// MediaList struct
// js:"MediaList,omit"
type MediaList struct {
}

// AppendMedium
func (*MediaList) AppendMedium(newMedium string) {
	js.Rewrite("$<.AppendMedium($1)", newMedium)
}

// DeleteMedium
func (*MediaList) DeleteMedium(oldMedium string) {
	js.Rewrite("$<.DeleteMedium($1)", oldMedium)
}

// Item
func (*MediaList) Item(index uint) (s string) {
	js.Rewrite("$<.Item($1)", index)
	return s
}

// ToString
func (*MediaList) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Length
func (*MediaList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// MediaText
func (*MediaList) MediaText() (mediaText string) {
	js.Rewrite("$<.MediaText")
	return mediaText
}

// MediaText
func (*MediaList) SetMediaText(mediaText string) {
	js.Rewrite("$<.MediaText = $1", mediaText)
}
