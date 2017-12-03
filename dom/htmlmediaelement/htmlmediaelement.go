package htmlmediaelement

import (
	"github.com/matthewmueller/joy/dom/avtrack"
	"github.com/matthewmueller/joy/dom/mediaencryptedevent"
	"github.com/matthewmueller/joy/dom/mediaerror"
	"github.com/matthewmueller/joy/dom/msgraphicstrust"
	"github.com/matthewmueller/joy/dom/msmediakeyneededevent"
	"github.com/matthewmueller/joy/dom/msmediakeys"
	"github.com/matthewmueller/joy/dom/texttrack"
	"github.com/matthewmueller/joy/dom/texttracklist"
	"github.com/matthewmueller/joy/dom/timeranges"
	"github.com/matthewmueller/joy/dom/window"
)

// HTMLMediaElement interface
// js:"HTMLMediaElement"
type HTMLMediaElement interface {
	window.HTMLElement

	// AddTextTrack
	// js:"addTextTrack"
	// jsrewrite:"$_.addTextTrack($1, $2, $3)"
	AddTextTrack(kind string, label *string, language *string) (t *texttrack.TextTrack)

	// CanPlayType Returns a string that specifies whether the client can play a given media resource type.
	// js:"canPlayType"
	// jsrewrite:"$_.canPlayType($1)"
	CanPlayType(kind string) (s string)

	// Load Resets the audio or video object and loads a new media resource.
	// js:"load"
	// jsrewrite:"$_.load()"
	Load()

	// MsClearEffects Clears all effects from the media pipeline.
	// js:"msClearEffects"
	// jsrewrite:"$_.msClearEffects()"
	MsClearEffects()

	// MsGetAsCastingSource
	// js:"msGetAsCastingSource"
	// jsrewrite:"$_.msGetAsCastingSource()"
	MsGetAsCastingSource() (i interface{})

	// MsInsertAudioEffect Inserts the specified audio effect into media pipeline.
	// js:"msInsertAudioEffect"
	// jsrewrite:"$_.msInsertAudioEffect($1, $2, $3)"
	MsInsertAudioEffect(activatableClassId string, effectRequired bool, config *interface{})

	// MsSetMediaKeys
	// js:"msSetMediaKeys"
	// jsrewrite:"$_.msSetMediaKeys($1)"
	MsSetMediaKeys(mediaKeys *msmediakeys.MSMediaKeys)

	// MsSetMediaProtectionManager Specifies the media protection manager for a given media pipeline.
	// js:"msSetMediaProtectionManager"
	// jsrewrite:"$_.msSetMediaProtectionManager($1)"
	MsSetMediaProtectionManager(mediaProtectionManager *interface{})

	// Pause Pauses the current playback and sets paused to TRUE. This can be used to test whether the media is playing or paused. You can also use the pause or play events to tell whether the media is playing or not.
	// js:"pause"
	// jsrewrite:"$_.pause()"
	Pause()

	// Play Loads and starts playback of a media resource.
	// js:"play"
	// jsrewrite:"$_.play()"
	Play()

	// SetMediaKeys
	// js:"setMediaKeys"
	// jsrewrite:"await $_.setMediaKeys($1)"
	SetMediaKeys(mediaKeys *window.MediaKeys)

	// AudioTracks prop Returns an AudioTrackList object with the audio tracks for a given video element.
	// js:"audioTracks"
	// jsrewrite:"$_.audioTracks"
	AudioTracks() (audioTracks *avtrack.AudioTrackList)

	// Autoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
	// js:"autoplay"
	// jsrewrite:"$_.autoplay"
	Autoplay() (autoplay bool)

	// SetAutoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
	// js:"autoplay"
	// jsrewrite:"$_.autoplay = $1"
	SetAutoplay(autoplay bool)

	// Buffered prop Gets a collection of buffered time ranges.
	// js:"buffered"
	// jsrewrite:"$_.buffered"
	Buffered() (buffered *timeranges.TimeRanges)

	// Controls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	// js:"controls"
	// jsrewrite:"$_.controls"
	Controls() (controls bool)

	// SetControls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	// js:"controls"
	// jsrewrite:"$_.controls = $1"
	SetControls(controls bool)

	// CrossOrigin prop
	// js:"crossOrigin"
	// jsrewrite:"$_.crossOrigin"
	CrossOrigin() (crossOrigin string)

	// SetCrossOrigin prop
	// js:"crossOrigin"
	// jsrewrite:"$_.crossOrigin = $1"
	SetCrossOrigin(crossOrigin string)

	// CurrentSrc prop Gets the address or URL of the current media resource that is selected by IHTMLMediaElement.
	// js:"currentSrc"
	// jsrewrite:"$_.currentSrc"
	CurrentSrc() (currentSrc string)

	// CurrentTime prop Gets or sets the current playback position, in seconds.
	// js:"currentTime"
	// jsrewrite:"$_.currentTime"
	CurrentTime() (currentTime float32)

	// SetCurrentTime prop Gets or sets the current playback position, in seconds.
	// js:"currentTime"
	// jsrewrite:"$_.currentTime = $1"
	SetCurrentTime(currentTime float32)

	// DefaultMuted prop
	// js:"defaultMuted"
	// jsrewrite:"$_.defaultMuted"
	DefaultMuted() (defaultMuted bool)

	// SetDefaultMuted prop
	// js:"defaultMuted"
	// jsrewrite:"$_.defaultMuted = $1"
	SetDefaultMuted(defaultMuted bool)

	// DefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	// js:"defaultPlaybackRate"
	// jsrewrite:"$_.defaultPlaybackRate"
	DefaultPlaybackRate() (defaultPlaybackRate float32)

	// SetDefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	// js:"defaultPlaybackRate"
	// jsrewrite:"$_.defaultPlaybackRate = $1"
	SetDefaultPlaybackRate(defaultPlaybackRate float32)

	// Duration prop Returns the duration in seconds of the current media resource. A NaN value is returned if duration is not available, or Infinity if the media resource is streaming.
	// js:"duration"
	// jsrewrite:"$_.duration"
	Duration() (duration float32)

	// Ended prop Gets information about whether the playback has ended or not.
	// js:"ended"
	// jsrewrite:"$_.ended"
	Ended() (ended bool)

	// Error prop Returns an object representing the current error state of the audio or video element.
	// js:"error"
	// jsrewrite:"$_.error"
	Error() (err *mediaerror.MediaError)

	// Loop prop Gets or sets a flag to specify whether playback should restart after it completes.
	// js:"loop"
	// jsrewrite:"$_.loop"
	Loop() (loop bool)

	// SetLoop prop Gets or sets a flag to specify whether playback should restart after it completes.
	// js:"loop"
	// jsrewrite:"$_.loop = $1"
	SetLoop(loop bool)

	// MediaKeys prop
	// js:"mediaKeys"
	// jsrewrite:"$_.mediaKeys"
	MediaKeys() (mediaKeys *window.MediaKeys)

	// MsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
	// js:"msAudioCategory"
	// jsrewrite:"$_.msAudioCategory"
	MsAudioCategory() (msAudioCategory string)

	// SetMsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
	// js:"msAudioCategory"
	// jsrewrite:"$_.msAudioCategory = $1"
	SetMsAudioCategory(msAudioCategory string)

	// MsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
	// js:"msAudioDeviceType"
	// jsrewrite:"$_.msAudioDeviceType"
	MsAudioDeviceType() (msAudioDeviceType string)

	// SetMsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
	// js:"msAudioDeviceType"
	// jsrewrite:"$_.msAudioDeviceType = $1"
	SetMsAudioDeviceType(msAudioDeviceType string)

	// MsGraphicsTrustStatus prop
	// js:"msGraphicsTrustStatus"
	// jsrewrite:"$_.msGraphicsTrustStatus"
	MsGraphicsTrustStatus() (msGraphicsTrustStatus *msgraphicstrust.MSGraphicsTrust)

	// MsKeys prop Gets the MSMediaKeys object, which is used for decrypting media data, that is associated with this media element.
	// js:"msKeys"
	// jsrewrite:"$_.msKeys"
	MsKeys() (msKeys *msmediakeys.MSMediaKeys)

	// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
	// js:"msPlayToDisabled"
	// jsrewrite:"$_.msPlayToDisabled"
	MsPlayToDisabled() (msPlayToDisabled bool)

	// SetMsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
	// js:"msPlayToDisabled"
	// jsrewrite:"$_.msPlayToDisabled = $1"
	SetMsPlayToDisabled(msPlayToDisabled bool)

	// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	// js:"msPlayToPreferredSourceUri"
	// jsrewrite:"$_.msPlayToPreferredSourceUri"
	MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string)

	// SetMsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	// js:"msPlayToPreferredSourceUri"
	// jsrewrite:"$_.msPlayToPreferredSourceUri = $1"
	SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string)

	// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
	// js:"msPlayToPrimary"
	// jsrewrite:"$_.msPlayToPrimary"
	MsPlayToPrimary() (msPlayToPrimary bool)

	// SetMsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
	// js:"msPlayToPrimary"
	// jsrewrite:"$_.msPlayToPrimary = $1"
	SetMsPlayToPrimary(msPlayToPrimary bool)

	// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
	// js:"msPlayToSource"
	// jsrewrite:"$_.msPlayToSource"
	MsPlayToSource() (msPlayToSource interface{})

	// MsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
	// js:"msRealTime"
	// jsrewrite:"$_.msRealTime"
	MsRealTime() (msRealTime bool)

	// SetMsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
	// js:"msRealTime"
	// jsrewrite:"$_.msRealTime = $1"
	SetMsRealTime(msRealTime bool)

	// Muted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	// js:"muted"
	// jsrewrite:"$_.muted"
	Muted() (muted bool)

	// SetMuted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	// js:"muted"
	// jsrewrite:"$_.muted = $1"
	SetMuted(muted bool)

	// NetworkState prop Gets the current network activity for the element.
	// js:"networkState"
	// jsrewrite:"$_.networkState"
	NetworkState() (networkState uint8)

	// Onencrypted prop
	// js:"onencrypted"
	// jsrewrite:"$_.onencrypted"
	Onencrypted() (onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// SetOnencrypted prop
	// js:"onencrypted"
	// jsrewrite:"$_.onencrypted = $1"
	SetOnencrypted(onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// Onmsneedkey prop
	// js:"onmsneedkey"
	// jsrewrite:"$_.onmsneedkey"
	Onmsneedkey() (onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// SetOnmsneedkey prop
	// js:"onmsneedkey"
	// jsrewrite:"$_.onmsneedkey = $1"
	SetOnmsneedkey(onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// Paused prop Gets a flag that specifies whether playback is paused.
	// js:"paused"
	// jsrewrite:"$_.paused"
	Paused() (paused bool)

	// PlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	// js:"playbackRate"
	// jsrewrite:"$_.playbackRate"
	PlaybackRate() (playbackRate float32)

	// SetPlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	// js:"playbackRate"
	// jsrewrite:"$_.playbackRate = $1"
	SetPlaybackRate(playbackRate float32)

	// Played prop Gets TimeRanges for the current media resource that has been played.
	// js:"played"
	// jsrewrite:"$_.played"
	Played() (played *timeranges.TimeRanges)

	// Preload prop Gets or sets the current playback position, in seconds.
	// js:"preload"
	// jsrewrite:"$_.preload"
	Preload() (preload string)

	// SetPreload prop Gets or sets the current playback position, in seconds.
	// js:"preload"
	// jsrewrite:"$_.preload = $1"
	SetPreload(preload string)

	// ReadyState prop
	// js:"readyState"
	// jsrewrite:"$_.readyState"
	ReadyState() (readyState interface{})

	// Seekable prop Returns a TimeRanges object that represents the ranges of the current media resource that can be seeked.
	// js:"seekable"
	// jsrewrite:"$_.seekable"
	Seekable() (seekable *timeranges.TimeRanges)

	// Seeking prop Gets a flag that indicates whether the the client is currently moving to a new playback position in the media resource.
	// js:"seeking"
	// jsrewrite:"$_.seeking"
	Seeking() (seeking bool)

	// Src prop The address or URL of the a media resource that is to be considered.
	// js:"src"
	// jsrewrite:"$_.src"
	Src() (src string)

	// SetSrc prop The address or URL of the a media resource that is to be considered.
	// js:"src"
	// jsrewrite:"$_.src = $1"
	SetSrc(src string)

	// SrcObject prop
	// js:"srcObject"
	// jsrewrite:"$_.srcObject"
	SrcObject() (srcObject *window.MediaStream)

	// SetSrcObject prop
	// js:"srcObject"
	// jsrewrite:"$_.srcObject = $1"
	SetSrcObject(srcObject *window.MediaStream)

	// TextTracks prop
	// js:"textTracks"
	// jsrewrite:"$_.textTracks"
	TextTracks() (textTracks *texttracklist.TextTrackList)

	// VideoTracks prop
	// js:"videoTracks"
	// jsrewrite:"$_.videoTracks"
	VideoTracks() (videoTracks *avtrack.VideoTrackList)

	// Volume prop Gets or sets the volume level for audio portions of the media element.
	// js:"volume"
	// jsrewrite:"$_.volume"
	Volume() (volume float32)

	// SetVolume prop Gets or sets the volume level for audio portions of the media element.
	// js:"volume"
	// jsrewrite:"$_.volume = $1"
	SetVolume(volume float32)
}
