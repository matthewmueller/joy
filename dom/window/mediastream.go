package window

import "github.com/matthewmueller/golly/js"

// NewMediaStream fn
func NewMediaStream(streamortracks *interface{}) *MediaStream {
	js.Rewrite("MediaStream")
	return &MediaStream{}
}

// MediaStream struct
// js:"MediaStream,omit"
type MediaStream struct {
	EventTarget
}

// AddTrack fn
func (*MediaStream) AddTrack(track *MediaStreamTrack) {
	js.Rewrite("$<.addTrack($1)", track)
}

// Clone fn
func (*MediaStream) Clone() (m *MediaStream) {
	js.Rewrite("$<.clone()")
	return m
}

// GetAudioTracks fn
func (*MediaStream) GetAudioTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.getAudioTracks()")
	return m
}

// GetTrackByID fn
func (*MediaStream) GetTrackByID(trackId string) (m *MediaStreamTrack) {
	js.Rewrite("$<.getTrackById($1)", trackId)
	return m
}

// GetTracks fn
func (*MediaStream) GetTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.getTracks()")
	return m
}

// GetVideoTracks fn
func (*MediaStream) GetVideoTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.getVideoTracks()")
	return m
}

// RemoveTrack fn
func (*MediaStream) RemoveTrack(track *MediaStreamTrack) {
	js.Rewrite("$<.removeTrack($1)", track)
}

// Stop fn
func (*MediaStream) Stop() {
	js.Rewrite("$<.stop()")
}

// Active prop
func (*MediaStream) Active() (active bool) {
	js.Rewrite("$<.active")
	return active
}

// ID prop
func (*MediaStream) ID() (id string) {
	js.Rewrite("$<.id")
	return id
}

// Onactive prop
func (*MediaStream) Onactive() (onactive func(Event)) {
	js.Rewrite("$<.onactive")
	return onactive
}

// Onactive prop
func (*MediaStream) SetOnactive(onactive func(Event)) {
	js.Rewrite("$<.onactive = $1", onactive)
}

// Onaddtrack prop
func (*MediaStream) Onaddtrack() (onaddtrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.onaddtrack")
	return onaddtrack
}

// Onaddtrack prop
func (*MediaStream) SetOnaddtrack(onaddtrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.onaddtrack = $1", onaddtrack)
}

// Oninactive prop
func (*MediaStream) Oninactive() (oninactive func(Event)) {
	js.Rewrite("$<.oninactive")
	return oninactive
}

// Oninactive prop
func (*MediaStream) SetOninactive(oninactive func(Event)) {
	js.Rewrite("$<.oninactive = $1", oninactive)
}

// Onremovetrack prop
func (*MediaStream) Onremovetrack() (onremovetrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.onremovetrack")
	return onremovetrack
}

// Onremovetrack prop
func (*MediaStream) SetOnremovetrack(onremovetrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.onremovetrack = $1", onremovetrack)
}
