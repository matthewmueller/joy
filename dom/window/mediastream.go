package window

import "github.com/matthewmueller/joy/macro"

var _ EventTarget = (*MediaStream)(nil)

// NewMediaStream fn
func NewMediaStream(streamortracks *interface{}) *MediaStream {
	macro.Rewrite("MediaStream")
	return &MediaStream{}
}

// MediaStream struct
// js:"MediaStream,omit"
type MediaStream struct {
}

// AddTrack fn
// js:"addTrack"
func (*MediaStream) AddTrack(track *MediaStreamTrack) {
	macro.Rewrite("$_.addTrack($1)", track)
}

// Clone fn
// js:"clone"
func (*MediaStream) Clone() (m *MediaStream) {
	macro.Rewrite("$_.clone()")
	return m
}

// GetAudioTracks fn
// js:"getAudioTracks"
func (*MediaStream) GetAudioTracks() (m []*MediaStreamTrack) {
	macro.Rewrite("$_.getAudioTracks()")
	return m
}

// GetTrackByID fn
// js:"getTrackById"
func (*MediaStream) GetTrackByID(trackId string) (m *MediaStreamTrack) {
	macro.Rewrite("$_.getTrackById($1)", trackId)
	return m
}

// GetTracks fn
// js:"getTracks"
func (*MediaStream) GetTracks() (m []*MediaStreamTrack) {
	macro.Rewrite("$_.getTracks()")
	return m
}

// GetVideoTracks fn
// js:"getVideoTracks"
func (*MediaStream) GetVideoTracks() (m []*MediaStreamTrack) {
	macro.Rewrite("$_.getVideoTracks()")
	return m
}

// RemoveTrack fn
// js:"removeTrack"
func (*MediaStream) RemoveTrack(track *MediaStreamTrack) {
	macro.Rewrite("$_.removeTrack($1)", track)
}

// Stop fn
// js:"stop"
func (*MediaStream) Stop() {
	macro.Rewrite("$_.stop()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaStream) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaStream) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaStream) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Active prop
// js:"active"
func (*MediaStream) Active() (active bool) {
	macro.Rewrite("$_.active")
	return active
}

// ID prop
// js:"id"
func (*MediaStream) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// Onactive prop
// js:"onactive"
func (*MediaStream) Onactive() (onactive func(Event)) {
	macro.Rewrite("$_.onactive")
	return onactive
}

// SetOnactive prop
// js:"onactive"
func (*MediaStream) SetOnactive(onactive func(Event)) {
	macro.Rewrite("$_.onactive = $1", onactive)
}

// Onaddtrack prop
// js:"onaddtrack"
func (*MediaStream) Onaddtrack() (onaddtrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onaddtrack")
	return onaddtrack
}

// SetOnaddtrack prop
// js:"onaddtrack"
func (*MediaStream) SetOnaddtrack(onaddtrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onaddtrack = $1", onaddtrack)
}

// Oninactive prop
// js:"oninactive"
func (*MediaStream) Oninactive() (oninactive func(Event)) {
	macro.Rewrite("$_.oninactive")
	return oninactive
}

// SetOninactive prop
// js:"oninactive"
func (*MediaStream) SetOninactive(oninactive func(Event)) {
	macro.Rewrite("$_.oninactive = $1", oninactive)
}

// Onremovetrack prop
// js:"onremovetrack"
func (*MediaStream) Onremovetrack() (onremovetrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onremovetrack")
	return onremovetrack
}

// SetOnremovetrack prop
// js:"onremovetrack"
func (*MediaStream) SetOnremovetrack(onremovetrack func(*MediaStreamTrackEvent)) {
	macro.Rewrite("$_.onremovetrack = $1", onremovetrack)
}
