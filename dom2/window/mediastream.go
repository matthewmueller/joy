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

// AddTrack
func (*MediaStream) AddTrack(track *MediaStreamTrack) {
	js.Rewrite("$<.AddTrack($1)", track)
}

// Clone
func (*MediaStream) Clone() (m *MediaStream) {
	js.Rewrite("$<.Clone()")
	return m
}

// GetAudioTracks
func (*MediaStream) GetAudioTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.GetAudioTracks()")
	return m
}

// GetTrackByID
func (*MediaStream) GetTrackByID(trackId string) (m *MediaStreamTrack) {
	js.Rewrite("$<.GetTrackByID($1)", trackId)
	return m
}

// GetTracks
func (*MediaStream) GetTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.GetTracks()")
	return m
}

// GetVideoTracks
func (*MediaStream) GetVideoTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.GetVideoTracks()")
	return m
}

// RemoveTrack
func (*MediaStream) RemoveTrack(track *MediaStreamTrack) {
	js.Rewrite("$<.RemoveTrack($1)", track)
}

// Stop
func (*MediaStream) Stop() {
	js.Rewrite("$<.Stop()")
}

// Active
func (*MediaStream) Active() (active bool) {
	js.Rewrite("$<.Active")
	return active
}

// ID
func (*MediaStream) ID() (id string) {
	js.Rewrite("$<.ID")
	return id
}

// Onactive
func (*MediaStream) Onactive() (onactive func(Event)) {
	js.Rewrite("$<.Onactive")
	return onactive
}

// Onactive
func (*MediaStream) SetOnactive(onactive func(Event)) {
	js.Rewrite("$<.Onactive = $1", onactive)
}

// Onaddtrack
func (*MediaStream) Onaddtrack() (onaddtrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.Onaddtrack")
	return onaddtrack
}

// Onaddtrack
func (*MediaStream) SetOnaddtrack(onaddtrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.Onaddtrack = $1", onaddtrack)
}

// Oninactive
func (*MediaStream) Oninactive() (oninactive func(Event)) {
	js.Rewrite("$<.Oninactive")
	return oninactive
}

// Oninactive
func (*MediaStream) SetOninactive(oninactive func(Event)) {
	js.Rewrite("$<.Oninactive = $1", oninactive)
}

// Onremovetrack
func (*MediaStream) Onremovetrack() (onremovetrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.Onremovetrack")
	return onremovetrack
}

// Onremovetrack
func (*MediaStream) SetOnremovetrack(onremovetrack func(*MediaStreamTrackEvent)) {
	js.Rewrite("$<.Onremovetrack = $1", onremovetrack)
}
