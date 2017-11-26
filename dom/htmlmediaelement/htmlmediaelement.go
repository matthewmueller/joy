package htmlmediaelement

import (
	"github.com/matthewmueller/golly/dom/avtrack"
	"github.com/matthewmueller/golly/dom/mediaencryptedevent"
	"github.com/matthewmueller/golly/dom/mediaerror"
	"github.com/matthewmueller/golly/dom/msgraphicstrust"
	"github.com/matthewmueller/golly/dom/msmediakeyneededevent"
	"github.com/matthewmueller/golly/dom/msmediakeys"
	"github.com/matthewmueller/golly/dom/texttrack"
	"github.com/matthewmueller/golly/dom/texttracklist"
	"github.com/matthewmueller/golly/dom/timeranges"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLMediaElement,omit"
type HTMLMediaElement interface {
	window.HTMLElement

	// AddTextTrack
	// js:"addTextTrack,rewrite=addtexttrack"
	AddTextTrack(kind string, label *string, language *string) (t *texttrack.TextTrack)

	// CanPlayType Returns a string that specifies whether the client can play a given media resource type.
	// js:"canPlayType,rewrite=canplaytype"
	CanPlayType(kind string) (s string)

	// Load Resets the audio or video object and loads a new media resource.
	// js:"load,rewrite=load"
	Load()

	// MsClearEffects Clears all effects from the media pipeline.
	// js:"msClearEffects,rewrite=mscleareffects"
	MsClearEffects()

	// MsGetAsCastingSource
	// js:"msGetAsCastingSource,rewrite=msgetascastingsource"
	MsGetAsCastingSource() (i interface{})

	// MsInsertAudioEffect Inserts the specified audio effect into media pipeline.
	// js:"msInsertAudioEffect,rewrite=msinsertaudioeffect"
	MsInsertAudioEffect(activatableClassId string, effectRequired bool, config *interface{})

	// MsSetMediaKeys
	// js:"msSetMediaKeys,rewrite=mssetmediakeys"
	MsSetMediaKeys(mediaKeys *msmediakeys.MSMediaKeys)

	// MsSetMediaProtectionManager Specifies the media protection manager for a given media pipeline.
	// js:"msSetMediaProtectionManager,rewrite=mssetmediaprotectionmanager"
	MsSetMediaProtectionManager(mediaProtectionManager *interface{})

	// Pause Pauses the current playback and sets paused to TRUE. This can be used to test whether the media is playing or paused. You can also use the pause or play events to tell whether the media is playing or not.
	// js:"pause,rewrite=pause"
	Pause()

	// Play Loads and starts playback of a media resource.
	// js:"play,rewrite=play"
	Play()

	// SetMediaKeys
	// js:"setMediaKeys,rewrite=setmediakeys"
	SetMediaKeys(mediaKeys *window.MediaKeys)

	// AudioTracks prop Returns an AudioTrackList object with the audio tracks for a given video element.
	// js:"audioTracks,rewrite=audiotracks"
	AudioTracks() (audioTracks *avtrack.AudioTrackList)

	// Autoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
	// js:"autoplay,rewrite=autoplay"
	Autoplay() (autoplay bool)

	// Autoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
	// js:"setautoplay,rewrite=setautoplay"
	SetAutoplay(autoplay bool)

	// Buffered prop Gets a collection of buffered time ranges.
	// js:"buffered,rewrite=buffered"
	Buffered() (buffered *timeranges.TimeRanges)

	// Controls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	// js:"controls,rewrite=controls"
	Controls() (controls bool)

	// Controls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	// js:"setcontrols,rewrite=setcontrols"
	SetControls(controls bool)

	// CrossOrigin prop
	// js:"crossOrigin,rewrite=crossorigin"
	CrossOrigin() (crossOrigin string)

	// CrossOrigin prop
	// js:"setcrossOrigin,rewrite=setcrossorigin"
	SetCrossOrigin(crossOrigin string)

	// CurrentSrc prop Gets the address or URL of the current media resource that is selected by IHTMLMediaElement.
	// js:"currentSrc,rewrite=currentsrc"
	CurrentSrc() (currentSrc string)

	// CurrentTime prop Gets or sets the current playback position, in seconds.
	// js:"currentTime,rewrite=currenttime"
	CurrentTime() (currentTime float32)

	// CurrentTime prop Gets or sets the current playback position, in seconds.
	// js:"setcurrentTime,rewrite=setcurrenttime"
	SetCurrentTime(currentTime float32)

	// DefaultMuted prop
	// js:"defaultMuted,rewrite=defaultmuted"
	DefaultMuted() (defaultMuted bool)

	// DefaultMuted prop
	// js:"setdefaultMuted,rewrite=setdefaultmuted"
	SetDefaultMuted(defaultMuted bool)

	// DefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	// js:"defaultPlaybackRate,rewrite=defaultplaybackrate"
	DefaultPlaybackRate() (defaultPlaybackRate float32)

	// DefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	// js:"setdefaultPlaybackRate,rewrite=setdefaultplaybackrate"
	SetDefaultPlaybackRate(defaultPlaybackRate float32)

	// Duration prop Returns the duration in seconds of the current media resource. A NaN value is returned if duration is not available, or Infinity if the media resource is streaming.
	// js:"duration,rewrite=duration"
	Duration() (duration float32)

	// Ended prop Gets information about whether the playback has ended or not.
	// js:"ended,rewrite=ended"
	Ended() (ended bool)

	// Error prop Returns an object representing the current error state of the audio or video element.
	// js:"error,rewrite=err"
	Error() (err *mediaerror.MediaError)

	// Loop prop Gets or sets a flag to specify whether playback should restart after it completes.
	// js:"loop,rewrite=loop"
	Loop() (loop bool)

	// Loop prop Gets or sets a flag to specify whether playback should restart after it completes.
	// js:"setloop,rewrite=setloop"
	SetLoop(loop bool)

	// MediaKeys prop
	// js:"mediaKeys,rewrite=mediakeys"
	MediaKeys() (mediaKeys *window.MediaKeys)

	// MsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
	// js:"msAudioCategory,rewrite=msaudiocategory"
	MsAudioCategory() (msAudioCategory string)

	// MsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
	// js:"setmsAudioCategory,rewrite=setmsaudiocategory"
	SetMsAudioCategory(msAudioCategory string)

	// MsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
	// js:"msAudioDeviceType,rewrite=msaudiodevicetype"
	MsAudioDeviceType() (msAudioDeviceType string)

	// MsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
	// js:"setmsAudioDeviceType,rewrite=setmsaudiodevicetype"
	SetMsAudioDeviceType(msAudioDeviceType string)

	// MsGraphicsTrustStatus prop
	// js:"msGraphicsTrustStatus,rewrite=msgraphicstruststatus"
	MsGraphicsTrustStatus() (msGraphicsTrustStatus *msgraphicstrust.MSGraphicsTrust)

	// MsKeys prop Gets the MSMediaKeys object, which is used for decrypting media data, that is associated with this media element.
	// js:"msKeys,rewrite=mskeys"
	MsKeys() (msKeys *msmediakeys.MSMediaKeys)

	// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
	// js:"msPlayToDisabled,rewrite=msplaytodisabled"
	MsPlayToDisabled() (msPlayToDisabled bool)

	// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
	// js:"setmsPlayToDisabled,rewrite=setmsplaytodisabled"
	SetMsPlayToDisabled(msPlayToDisabled bool)

	// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	// js:"msPlayToPreferredSourceUri,rewrite=msplaytopreferredsourceuri"
	MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string)

	// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	// js:"setmsPlayToPreferredSourceUri,rewrite=setmsplaytopreferredsourceuri"
	SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string)

	// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
	// js:"msPlayToPrimary,rewrite=msplaytoprimary"
	MsPlayToPrimary() (msPlayToPrimary bool)

	// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
	// js:"setmsPlayToPrimary,rewrite=setmsplaytoprimary"
	SetMsPlayToPrimary(msPlayToPrimary bool)

	// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
	// js:"msPlayToSource,rewrite=msplaytosource"
	MsPlayToSource() (msPlayToSource interface{})

	// MsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
	// js:"msRealTime,rewrite=msrealtime"
	MsRealTime() (msRealTime bool)

	// MsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
	// js:"setmsRealTime,rewrite=setmsrealtime"
	SetMsRealTime(msRealTime bool)

	// Muted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	// js:"muted,rewrite=muted"
	Muted() (muted bool)

	// Muted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	// js:"setmuted,rewrite=setmuted"
	SetMuted(muted bool)

	// NetworkState prop Gets the current network activity for the element.
	// js:"networkState,rewrite=networkstate"
	NetworkState() (networkState uint8)

	// Onencrypted prop
	// js:"onencrypted,rewrite=onencrypted"
	Onencrypted() (onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// Onencrypted prop
	// js:"setonencrypted,rewrite=setonencrypted"
	SetOnencrypted(onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// Onmsneedkey prop
	// js:"onmsneedkey,rewrite=onmsneedkey"
	Onmsneedkey() (onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// Onmsneedkey prop
	// js:"setonmsneedkey,rewrite=setonmsneedkey"
	SetOnmsneedkey(onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// Paused prop Gets a flag that specifies whether playback is paused.
	// js:"paused,rewrite=paused"
	Paused() (paused bool)

	// PlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	// js:"playbackRate,rewrite=playbackrate"
	PlaybackRate() (playbackRate float32)

	// PlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	// js:"setplaybackRate,rewrite=setplaybackrate"
	SetPlaybackRate(playbackRate float32)

	// Played prop Gets TimeRanges for the current media resource that has been played.
	// js:"played,rewrite=played"
	Played() (played *timeranges.TimeRanges)

	// Preload prop Gets or sets the current playback position, in seconds.
	// js:"preload,rewrite=preload"
	Preload() (preload string)

	// Preload prop Gets or sets the current playback position, in seconds.
	// js:"setpreload,rewrite=setpreload"
	SetPreload(preload string)

	// ReadyState prop
	// js:"readyState,rewrite=readystate"
	ReadyState() (readyState interface{})

	// Seekable prop Returns a TimeRanges object that represents the ranges of the current media resource that can be seeked.
	// js:"seekable,rewrite=seekable"
	Seekable() (seekable *timeranges.TimeRanges)

	// Seeking prop Gets a flag that indicates whether the the client is currently moving to a new playback position in the media resource.
	// js:"seeking,rewrite=seeking"
	Seeking() (seeking bool)

	// Src prop The address or URL of the a media resource that is to be considered.
	// js:"src,rewrite=src"
	Src() (src string)

	// Src prop The address or URL of the a media resource that is to be considered.
	// js:"setsrc,rewrite=setsrc"
	SetSrc(src string)

	// SrcObject prop
	// js:"srcObject,rewrite=srcobject"
	SrcObject() (srcObject *window.MediaStream)

	// SrcObject prop
	// js:"setsrcObject,rewrite=setsrcobject"
	SetSrcObject(srcObject *window.MediaStream)

	// TextTracks prop
	// js:"textTracks,rewrite=texttracks"
	TextTracks() (textTracks *texttracklist.TextTrackList)

	// VideoTracks prop
	// js:"videoTracks,rewrite=videotracks"
	VideoTracks() (videoTracks *avtrack.VideoTrackList)

	// Volume prop Gets or sets the volume level for audio portions of the media element.
	// js:"volume,rewrite=volume"
	Volume() (volume float32)

	// Volume prop Gets or sets the volume level for audio portions of the media element.
	// js:"setvolume,rewrite=setvolume"
	SetVolume(volume float32)
}

// addtexttrack fn
func addtexttrack(kind string, label *string, language *string) (t *texttrack.TextTrack) {
	js.Rewrite("$<.addTextTrack($1, $2, $3)", kind, label, language)
	return t
}

// canplaytype fn Returns a string that specifies whether the client can play a given media resource type.
func canplaytype(kind string) (s string) {
	js.Rewrite("$<.canPlayType($1)", kind)
	return s
}

// load fn Resets the audio or video object and loads a new media resource.
func load() {
	js.Rewrite("$<.load()")
}

// mscleareffects fn Clears all effects from the media pipeline.
func mscleareffects() {
	js.Rewrite("$<.msClearEffects()")
}

// msgetascastingsource fn
func msgetascastingsource() (i interface{}) {
	js.Rewrite("$<.msGetAsCastingSource()")
	return i
}

// msinsertaudioeffect fn Inserts the specified audio effect into media pipeline.
func msinsertaudioeffect(activatableClassId string, effectRequired bool, config *interface{}) {
	js.Rewrite("$<.msInsertAudioEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

// mssetmediakeys fn
func mssetmediakeys(mediaKeys *msmediakeys.MSMediaKeys) {
	js.Rewrite("$<.msSetMediaKeys($1)", mediaKeys)
}

// mssetmediaprotectionmanager fn Specifies the media protection manager for a given media pipeline.
func mssetmediaprotectionmanager(mediaProtectionManager *interface{}) {
	js.Rewrite("$<.msSetMediaProtectionManager($1)", mediaProtectionManager)
}

// pause fn Pauses the current playback and sets paused to TRUE. This can be used to test whether the media is playing or paused. You can also use the pause or play events to tell whether the media is playing or not.
func pause() {
	js.Rewrite("$<.pause()")
}

// play fn Loads and starts playback of a media resource.
func play() {
	js.Rewrite("$<.play()")
}

// setmediakeys fn
func setmediakeys(mediaKeys *window.MediaKeys) {
	js.Rewrite("await $<.setMediaKeys($1)", mediaKeys)
}

// audiotracks prop Returns an AudioTrackList object with the audio tracks for a given video element.
func audiotracks() (audioTracks *avtrack.AudioTrackList) {
	js.Rewrite("$<.audioTracks")
	return audioTracks
}

// autoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
func autoplay() (autoplay bool) {
	js.Rewrite("$<.autoplay")
	return autoplay
}

// setautoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
func setautoplay(autoplay bool) {
	js.Rewrite("$<.autoplay = autoplay")
}

// buffered prop Gets a collection of buffered time ranges.
func buffered() (buffered *timeranges.TimeRanges) {
	js.Rewrite("$<.buffered")
	return buffered
}

// controls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
func controls() (controls bool) {
	js.Rewrite("$<.controls")
	return controls
}

// setcontrols prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
func setcontrols(controls bool) {
	js.Rewrite("$<.controls = controls")
}

// crossorigin prop
func crossorigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

// setcrossorigin prop
func setcrossorigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = crossOrigin")
}

// currentsrc prop Gets the address or URL of the current media resource that is selected by IHTMLMediaElement.
func currentsrc() (currentSrc string) {
	js.Rewrite("$<.currentSrc")
	return currentSrc
}

// currenttime prop Gets or sets the current playback position, in seconds.
func currenttime() (currentTime float32) {
	js.Rewrite("$<.currentTime")
	return currentTime
}

// setcurrenttime prop Gets or sets the current playback position, in seconds.
func setcurrenttime(currentTime float32) {
	js.Rewrite("$<.currentTime = currentTime")
}

// defaultmuted prop
func defaultmuted() (defaultMuted bool) {
	js.Rewrite("$<.defaultMuted")
	return defaultMuted
}

// setdefaultmuted prop
func setdefaultmuted(defaultMuted bool) {
	js.Rewrite("$<.defaultMuted = defaultMuted")
}

// defaultplaybackrate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
func defaultplaybackrate() (defaultPlaybackRate float32) {
	js.Rewrite("$<.defaultPlaybackRate")
	return defaultPlaybackRate
}

// setdefaultplaybackrate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
func setdefaultplaybackrate(defaultPlaybackRate float32) {
	js.Rewrite("$<.defaultPlaybackRate = defaultPlaybackRate")
}

// duration prop Returns the duration in seconds of the current media resource. A NaN value is returned if duration is not available, or Infinity if the media resource is streaming.
func duration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

// ended prop Gets information about whether the playback has ended or not.
func ended() (ended bool) {
	js.Rewrite("$<.ended")
	return ended
}

// err prop Returns an object representing the current error state of the audio or video element.
func err() (err *mediaerror.MediaError) {
	js.Rewrite("$<.error")
	return err
}

// loop prop Gets or sets a flag to specify whether playback should restart after it completes.
func loop() (loop bool) {
	js.Rewrite("$<.loop")
	return loop
}

// setloop prop Gets or sets a flag to specify whether playback should restart after it completes.
func setloop(loop bool) {
	js.Rewrite("$<.loop = loop")
}

// mediakeys prop
func mediakeys() (mediaKeys *window.MediaKeys) {
	js.Rewrite("$<.mediaKeys")
	return mediaKeys
}

// msaudiocategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
func msaudiocategory() (msAudioCategory string) {
	js.Rewrite("$<.msAudioCategory")
	return msAudioCategory
}

// setmsaudiocategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
func setmsaudiocategory(msAudioCategory string) {
	js.Rewrite("$<.msAudioCategory = msAudioCategory")
}

// msaudiodevicetype prop Specifies the output device id that the audio will be sent to.
func msaudiodevicetype() (msAudioDeviceType string) {
	js.Rewrite("$<.msAudioDeviceType")
	return msAudioDeviceType
}

// setmsaudiodevicetype prop Specifies the output device id that the audio will be sent to.
func setmsaudiodevicetype(msAudioDeviceType string) {
	js.Rewrite("$<.msAudioDeviceType = msAudioDeviceType")
}

// msgraphicstruststatus prop
func msgraphicstruststatus() (msGraphicsTrustStatus *msgraphicstrust.MSGraphicsTrust) {
	js.Rewrite("$<.msGraphicsTrustStatus")
	return msGraphicsTrustStatus
}

// mskeys prop Gets the MSMediaKeys object, which is used for decrypting media data, that is associated with this media element.
func mskeys() (msKeys *msmediakeys.MSMediaKeys) {
	js.Rewrite("$<.msKeys")
	return msKeys
}

// msplaytodisabled prop Gets or sets whether the DLNA PlayTo device is available.
func msplaytodisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

// setmsplaytodisabled prop Gets or sets whether the DLNA PlayTo device is available.
func setmsplaytodisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = msPlayToDisabled")
}

// msplaytopreferredsourceuri prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func msplaytopreferredsourceuri() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

// setmsplaytopreferredsourceuri prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func setmsplaytopreferredsourceuri(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = msPlayToPreferredSourceUri")
}

// msplaytoprimary prop Gets or sets the primary DLNA PlayTo device.
func msplaytoprimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

// setmsplaytoprimary prop Gets or sets the primary DLNA PlayTo device.
func setmsplaytoprimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = msPlayToPrimary")
}

// msplaytosource prop Gets the source associated with the media element for use by the PlayToManager.
func msplaytosource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

// msrealtime prop Specifies whether or not to enable low-latency playback on the media element.
func msrealtime() (msRealTime bool) {
	js.Rewrite("$<.msRealTime")
	return msRealTime
}

// setmsrealtime prop Specifies whether or not to enable low-latency playback on the media element.
func setmsrealtime(msRealTime bool) {
	js.Rewrite("$<.msRealTime = msRealTime")
}

// muted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
func muted() (muted bool) {
	js.Rewrite("$<.muted")
	return muted
}

// setmuted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
func setmuted(muted bool) {
	js.Rewrite("$<.muted = muted")
}

// networkstate prop Gets the current network activity for the element.
func networkstate() (networkState uint8) {
	js.Rewrite("$<.networkState")
	return networkState
}

// onencrypted prop
func onencrypted() (onencrypted func(*mediaencryptedevent.MediaEncryptedEvent)) {
	js.Rewrite("$<.onencrypted")
	return onencrypted
}

// setonencrypted prop
func setonencrypted(onencrypted func(*mediaencryptedevent.MediaEncryptedEvent)) {
	js.Rewrite("$<.onencrypted = onencrypted")
}

// onmsneedkey prop
func onmsneedkey() (onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent)) {
	js.Rewrite("$<.onmsneedkey")
	return onmsneedkey
}

// setonmsneedkey prop
func setonmsneedkey(onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent)) {
	js.Rewrite("$<.onmsneedkey = onmsneedkey")
}

// paused prop Gets a flag that specifies whether playback is paused.
func paused() (paused bool) {
	js.Rewrite("$<.paused")
	return paused
}

// playbackrate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
func playbackrate() (playbackRate float32) {
	js.Rewrite("$<.playbackRate")
	return playbackRate
}

// setplaybackrate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
func setplaybackrate(playbackRate float32) {
	js.Rewrite("$<.playbackRate = playbackRate")
}

// played prop Gets TimeRanges for the current media resource that has been played.
func played() (played *timeranges.TimeRanges) {
	js.Rewrite("$<.played")
	return played
}

// preload prop Gets or sets the current playback position, in seconds.
func preload() (preload string) {
	js.Rewrite("$<.preload")
	return preload
}

// setpreload prop Gets or sets the current playback position, in seconds.
func setpreload(preload string) {
	js.Rewrite("$<.preload = preload")
}

// readystate prop
func readystate() (readyState interface{}) {
	js.Rewrite("$<.readyState")
	return readyState
}

// seekable prop Returns a TimeRanges object that represents the ranges of the current media resource that can be seeked.
func seekable() (seekable *timeranges.TimeRanges) {
	js.Rewrite("$<.seekable")
	return seekable
}

// seeking prop Gets a flag that indicates whether the the client is currently moving to a new playback position in the media resource.
func seeking() (seeking bool) {
	js.Rewrite("$<.seeking")
	return seeking
}

// src prop The address or URL of the a media resource that is to be considered.
func src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// setsrc prop The address or URL of the a media resource that is to be considered.
func setsrc(src string) {
	js.Rewrite("$<.src = src")
}

// srcobject prop
func srcobject() (srcObject *window.MediaStream) {
	js.Rewrite("$<.srcObject")
	return srcObject
}

// setsrcobject prop
func setsrcobject(srcObject *window.MediaStream) {
	js.Rewrite("$<.srcObject = srcObject")
}

// texttracks prop
func texttracks() (textTracks *texttracklist.TextTrackList) {
	js.Rewrite("$<.textTracks")
	return textTracks
}

// videotracks prop
func videotracks() (videoTracks *avtrack.VideoTrackList) {
	js.Rewrite("$<.videoTracks")
	return videoTracks
}

// volume prop Gets or sets the volume level for audio portions of the media element.
func volume() (volume float32) {
	js.Rewrite("$<.volume")
	return volume
}

// setvolume prop Gets or sets the volume level for audio portions of the media element.
func setvolume(volume float32) {
	js.Rewrite("$<.volume = volume")
}
