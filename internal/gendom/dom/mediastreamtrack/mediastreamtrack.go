package mediastreamtrack

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamerrorevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrackstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediatrackcapabilities"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediatrackconstraints"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediatracksettings"
	"github.com/matthewmueller/golly/js"
)

type MediaStreamTrack struct {
	*eventtarget.EventTarget
}

func (*MediaStreamTrack) ApplyConstraints(constraints *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("await $<.applyConstraints($1)", constraints)
}

func (*MediaStreamTrack) Clone() (m *MediaStreamTrack) {
	js.Rewrite("$<.clone()")
	return m
}

func (*MediaStreamTrack) GetCapabilities() (m *mediatrackcapabilities.MediaTrackCapabilities) {
	js.Rewrite("$<.getCapabilities()")
	return m
}

func (*MediaStreamTrack) GetConstraints() (m *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("$<.getConstraints()")
	return m
}

func (*MediaStreamTrack) GetSettings() (m *mediatracksettings.MediaTrackSettings) {
	js.Rewrite("$<.getSettings()")
	return m
}

func (*MediaStreamTrack) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MediaStreamTrack) GetEnabled() (enabled bool) {
	js.Rewrite("$<.enabled")
	return enabled
}

func (*MediaStreamTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.enabled = $1", enabled)
}

func (*MediaStreamTrack) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*MediaStreamTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*MediaStreamTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*MediaStreamTrack) GetMuted() (muted bool) {
	js.Rewrite("$<.muted")
	return muted
}

func (*MediaStreamTrack) GetOnended() (ended *mediastreamerrorevent.MediaStreamErrorEvent) {
	js.Rewrite("$<.onended")
	return ended
}

func (*MediaStreamTrack) SetOnended(ended *mediastreamerrorevent.MediaStreamErrorEvent) {
	js.Rewrite("$<.onended = $1", ended)
}

func (*MediaStreamTrack) GetOnmute() (mute *event.Event) {
	js.Rewrite("$<.onmute")
	return mute
}

func (*MediaStreamTrack) SetOnmute(mute *event.Event) {
	js.Rewrite("$<.onmute = $1", mute)
}

func (*MediaStreamTrack) GetOnoverconstrained() (overconstrained *mediastreamerrorevent.MediaStreamErrorEvent) {
	js.Rewrite("$<.onoverconstrained")
	return overconstrained
}

func (*MediaStreamTrack) SetOnoverconstrained(overconstrained *mediastreamerrorevent.MediaStreamErrorEvent) {
	js.Rewrite("$<.onoverconstrained = $1", overconstrained)
}

func (*MediaStreamTrack) GetOnunmute() (unmute *event.Event) {
	js.Rewrite("$<.onunmute")
	return unmute
}

func (*MediaStreamTrack) SetOnunmute(unmute *event.Event) {
	js.Rewrite("$<.onunmute = $1", unmute)
}

func (*MediaStreamTrack) GetReadonly() (readonly bool) {
	js.Rewrite("$<.readonly")
	return readonly
}

func (*MediaStreamTrack) GetReadyState() (readyState *mediastreamtrackstate.MediaStreamTrackState) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MediaStreamTrack) GetRemote() (remote bool) {
	js.Rewrite("$<.remote")
	return remote
}
