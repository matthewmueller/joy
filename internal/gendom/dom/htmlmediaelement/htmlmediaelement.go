package htmlmediaelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiotracklist"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediaencryptedevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediaerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeys"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastream"
	"github.com/matthewmueller/golly/internal/gendom/dom/msgraphicstrust"
	"github.com/matthewmueller/golly/internal/gendom/dom/msmediakeyneededevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/msmediakeys"
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/texttracklist"
	"github.com/matthewmueller/golly/internal/gendom/dom/timeranges"
	"github.com/matthewmueller/golly/internal/gendom/dom/videotracklist"
	"github.com/matthewmueller/golly/js"
)

type HTMLMediaElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLMediaElement) AddTextTrack(kind string, label string, language string) (t *texttrack.TextTrack) {
	js.Rewrite("$<.addTextTrack($1, $2, $3)", kind, label, language)
	return t
}

func (*HTMLMediaElement) CanPlayType(kind string) (s string) {
	js.Rewrite("$<.canPlayType($1)", kind)
	return s
}

func (*HTMLMediaElement) Load() {
	js.Rewrite("$<.load()")
}

func (*HTMLMediaElement) MsClearEffects() {
	js.Rewrite("$<.msClearEffects()")
}

func (*HTMLMediaElement) MsGetAsCastingSource() (i interface{}) {
	js.Rewrite("$<.msGetAsCastingSource()")
	return i
}

func (*HTMLMediaElement) MsInsertAudioEffect(activatableClassId string, effectRequired bool, config interface{}) {
	js.Rewrite("$<.msInsertAudioEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

func (*HTMLMediaElement) MsSetMediaKeys(mediaKeys *msmediakeys.MSMediaKeys) {
	js.Rewrite("$<.msSetMediaKeys($1)", mediaKeys)
}

func (*HTMLMediaElement) MsSetMediaProtectionManager(mediaProtectionManager interface{}) {
	js.Rewrite("$<.msSetMediaProtectionManager($1)", mediaProtectionManager)
}

func (*HTMLMediaElement) Pause() {
	js.Rewrite("$<.pause()")
}

func (*HTMLMediaElement) Play() {
	js.Rewrite("$<.play()")
}

func (*HTMLMediaElement) SetMediaKeys(mediaKeys *mediakeys.MediaKeys) {
	js.Rewrite("await $<.setMediaKeys($1)", mediaKeys)
}

func (*HTMLMediaElement) GetAudioTracks() (audioTracks *audiotracklist.AudioTrackList) {
	js.Rewrite("$<.audioTracks")
	return audioTracks
}

func (*HTMLMediaElement) GetAutoplay() (autoplay bool) {
	js.Rewrite("$<.autoplay")
	return autoplay
}

func (*HTMLMediaElement) SetAutoplay(autoplay bool) {
	js.Rewrite("$<.autoplay = $1", autoplay)
}

func (*HTMLMediaElement) GetBuffered() (buffered *timeranges.TimeRanges) {
	js.Rewrite("$<.buffered")
	return buffered
}

func (*HTMLMediaElement) GetControls() (controls bool) {
	js.Rewrite("$<.controls")
	return controls
}

func (*HTMLMediaElement) SetControls(controls bool) {
	js.Rewrite("$<.controls = $1", controls)
}

func (*HTMLMediaElement) GetCrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

func (*HTMLMediaElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

func (*HTMLMediaElement) GetCurrentSrc() (currentSrc string) {
	js.Rewrite("$<.currentSrc")
	return currentSrc
}

func (*HTMLMediaElement) GetCurrentTime() (currentTime float32) {
	js.Rewrite("$<.currentTime")
	return currentTime
}

func (*HTMLMediaElement) SetCurrentTime(currentTime float32) {
	js.Rewrite("$<.currentTime = $1", currentTime)
}

func (*HTMLMediaElement) GetDefaultMuted() (defaultMuted bool) {
	js.Rewrite("$<.defaultMuted")
	return defaultMuted
}

func (*HTMLMediaElement) SetDefaultMuted(defaultMuted bool) {
	js.Rewrite("$<.defaultMuted = $1", defaultMuted)
}

func (*HTMLMediaElement) GetDefaultPlaybackRate() (defaultPlaybackRate float32) {
	js.Rewrite("$<.defaultPlaybackRate")
	return defaultPlaybackRate
}

func (*HTMLMediaElement) SetDefaultPlaybackRate(defaultPlaybackRate float32) {
	js.Rewrite("$<.defaultPlaybackRate = $1", defaultPlaybackRate)
}

func (*HTMLMediaElement) GetDuration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

func (*HTMLMediaElement) GetEnded() (ended bool) {
	js.Rewrite("$<.ended")
	return ended
}

func (*HTMLMediaElement) GetError() (err *mediaerror.MediaError) {
	js.Rewrite("$<.error")
	return err
}

func (*HTMLMediaElement) GetLoop() (loop bool) {
	js.Rewrite("$<.loop")
	return loop
}

func (*HTMLMediaElement) SetLoop(loop bool) {
	js.Rewrite("$<.loop = $1", loop)
}

func (*HTMLMediaElement) GetMediaKeys() (mediaKeys *mediakeys.MediaKeys) {
	js.Rewrite("$<.mediaKeys")
	return mediaKeys
}

func (*HTMLMediaElement) GetMsAudioCategory() (msAudioCategory string) {
	js.Rewrite("$<.msAudioCategory")
	return msAudioCategory
}

func (*HTMLMediaElement) SetMsAudioCategory(msAudioCategory string) {
	js.Rewrite("$<.msAudioCategory = $1", msAudioCategory)
}

func (*HTMLMediaElement) GetMsAudioDeviceType() (msAudioDeviceType string) {
	js.Rewrite("$<.msAudioDeviceType")
	return msAudioDeviceType
}

func (*HTMLMediaElement) SetMsAudioDeviceType(msAudioDeviceType string) {
	js.Rewrite("$<.msAudioDeviceType = $1", msAudioDeviceType)
}

func (*HTMLMediaElement) GetMsGraphicsTrustStatus() (msGraphicsTrustStatus *msgraphicstrust.MSGraphicsTrust) {
	js.Rewrite("$<.msGraphicsTrustStatus")
	return msGraphicsTrustStatus
}

func (*HTMLMediaElement) GetMsKeys() (msKeys *msmediakeys.MSMediaKeys) {
	js.Rewrite("$<.msKeys")
	return msKeys
}

func (*HTMLMediaElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLMediaElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLMediaElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLMediaElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLMediaElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLMediaElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLMediaElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLMediaElement) GetMsRealTime() (msRealTime bool) {
	js.Rewrite("$<.msRealTime")
	return msRealTime
}

func (*HTMLMediaElement) SetMsRealTime(msRealTime bool) {
	js.Rewrite("$<.msRealTime = $1", msRealTime)
}

func (*HTMLMediaElement) GetMuted() (muted bool) {
	js.Rewrite("$<.muted")
	return muted
}

func (*HTMLMediaElement) SetMuted(muted bool) {
	js.Rewrite("$<.muted = $1", muted)
}

func (*HTMLMediaElement) GetNetworkState() (networkState uint8) {
	js.Rewrite("$<.networkState")
	return networkState
}

func (*HTMLMediaElement) GetOnencrypted() (encrypted *mediaencryptedevent.MediaEncryptedEvent) {
	js.Rewrite("$<.onencrypted")
	return encrypted
}

func (*HTMLMediaElement) SetOnencrypted(encrypted *mediaencryptedevent.MediaEncryptedEvent) {
	js.Rewrite("$<.onencrypted = $1", encrypted)
}

func (*HTMLMediaElement) GetOnmsneedkey() (msneedkey *msmediakeyneededevent.MSMediaKeyNeededEvent) {
	js.Rewrite("$<.onmsneedkey")
	return msneedkey
}

func (*HTMLMediaElement) SetOnmsneedkey(msneedkey *msmediakeyneededevent.MSMediaKeyNeededEvent) {
	js.Rewrite("$<.onmsneedkey = $1", msneedkey)
}

func (*HTMLMediaElement) GetPaused() (paused bool) {
	js.Rewrite("$<.paused")
	return paused
}

func (*HTMLMediaElement) GetPlaybackRate() (playbackRate float32) {
	js.Rewrite("$<.playbackRate")
	return playbackRate
}

func (*HTMLMediaElement) SetPlaybackRate(playbackRate float32) {
	js.Rewrite("$<.playbackRate = $1", playbackRate)
}

func (*HTMLMediaElement) GetPlayed() (played *timeranges.TimeRanges) {
	js.Rewrite("$<.played")
	return played
}

func (*HTMLMediaElement) GetPreload() (preload string) {
	js.Rewrite("$<.preload")
	return preload
}

func (*HTMLMediaElement) SetPreload(preload string) {
	js.Rewrite("$<.preload = $1", preload)
}

func (*HTMLMediaElement) GetReadyState() (readyState interface{}) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLMediaElement) GetSeekable() (seekable *timeranges.TimeRanges) {
	js.Rewrite("$<.seekable")
	return seekable
}

func (*HTMLMediaElement) GetSeeking() (seeking bool) {
	js.Rewrite("$<.seeking")
	return seeking
}

func (*HTMLMediaElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLMediaElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLMediaElement) GetSrcObject() (srcObject *mediastream.MediaStream) {
	js.Rewrite("$<.srcObject")
	return srcObject
}

func (*HTMLMediaElement) SetSrcObject(srcObject *mediastream.MediaStream) {
	js.Rewrite("$<.srcObject = $1", srcObject)
}

func (*HTMLMediaElement) GetTextTracks() (textTracks *texttracklist.TextTrackList) {
	js.Rewrite("$<.textTracks")
	return textTracks
}

func (*HTMLMediaElement) GetVideoTracks() (videoTracks *videotracklist.VideoTrackList) {
	js.Rewrite("$<.videoTracks")
	return videoTracks
}

func (*HTMLMediaElement) GetVolume() (volume float32) {
	js.Rewrite("$<.volume")
	return volume
}

func (*HTMLMediaElement) SetVolume(volume float32) {
	js.Rewrite("$<.volume = $1", volume)
}
