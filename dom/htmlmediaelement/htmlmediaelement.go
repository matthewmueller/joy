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
)

// js:"HTMLMediaElement,omit"
type HTMLMediaElement interface {
	window.HTMLElement

	// AddTextTrack
	// js:"addTextTrack"
	AddTextTrack(kind string, label *string, language *string) (t *texttrack.TextTrack)

	// CanPlayType Returns a string that specifies whether the client can play a given media resource type.
	// js:"canPlayType"
	CanPlayType(kind string) (s string)

	// Load Resets the audio or video object and loads a new media resource.
	// js:"load"
	Load()

	// MsClearEffects Clears all effects from the media pipeline.
	// js:"msClearEffects"
	MsClearEffects()

	// MsGetAsCastingSource
	// js:"msGetAsCastingSource"
	MsGetAsCastingSource() (i interface{})

	// MsInsertAudioEffect Inserts the specified audio effect into media pipeline.
	// js:"msInsertAudioEffect"
	MsInsertAudioEffect(activatableClassId string, effectRequired bool, config *interface{})

	// MsSetMediaKeys
	// js:"msSetMediaKeys"
	MsSetMediaKeys(mediaKeys *msmediakeys.MSMediaKeys)

	// MsSetMediaProtectionManager Specifies the media protection manager for a given media pipeline.
	// js:"msSetMediaProtectionManager"
	MsSetMediaProtectionManager(mediaProtectionManager *interface{})

	// Pause Pauses the current playback and sets paused to TRUE. This can be used to test whether the media is playing or paused. You can also use the pause or play events to tell whether the media is playing or not.
	// js:"pause"
	Pause()

	// Play Loads and starts playback of a media resource.
	// js:"play"
	Play()

	// SetMediaKeys
	// js:"setMediaKeys"
	SetMediaKeys(mediaKeys *window.MediaKeys)

	// AudioTracks prop Returns an AudioTrackList object with the audio tracks for a given video element.
	// js:"audioTracks"
	AudioTracks() (audioTracks *avtrack.AudioTrackList)

	// Autoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
	// js:"autoplay"
	Autoplay() (autoplay bool)

	// Autoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
	SetAutoplay(autoplay bool)

	// Buffered prop Gets a collection of buffered time ranges.
	// js:"buffered"
	Buffered() (buffered *timeranges.TimeRanges)

	// Controls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	// js:"controls"
	Controls() (controls bool)

	// Controls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	SetControls(controls bool)

	// CrossOrigin prop
	// js:"crossOrigin"
	CrossOrigin() (crossOrigin string)

	// CrossOrigin prop
	SetCrossOrigin(crossOrigin string)

	// CurrentSrc prop Gets the address or URL of the current media resource that is selected by IHTMLMediaElement.
	// js:"currentSrc"
	CurrentSrc() (currentSrc string)

	// CurrentTime prop Gets or sets the current playback position, in seconds.
	// js:"currentTime"
	CurrentTime() (currentTime float32)

	// CurrentTime prop Gets or sets the current playback position, in seconds.
	SetCurrentTime(currentTime float32)

	// DefaultMuted prop
	// js:"defaultMuted"
	DefaultMuted() (defaultMuted bool)

	// DefaultMuted prop
	SetDefaultMuted(defaultMuted bool)

	// DefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	// js:"defaultPlaybackRate"
	DefaultPlaybackRate() (defaultPlaybackRate float32)

	// DefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	SetDefaultPlaybackRate(defaultPlaybackRate float32)

	// Duration prop Returns the duration in seconds of the current media resource. A NaN value is returned if duration is not available, or Infinity if the media resource is streaming.
	// js:"duration"
	Duration() (duration float32)

	// Ended prop Gets information about whether the playback has ended or not.
	// js:"ended"
	Ended() (ended bool)

	// Error prop Returns an object representing the current error state of the audio or video element.
	// js:"error"
	Error() (err *mediaerror.MediaError)

	// Loop prop Gets or sets a flag to specify whether playback should restart after it completes.
	// js:"loop"
	Loop() (loop bool)

	// Loop prop Gets or sets a flag to specify whether playback should restart after it completes.
	SetLoop(loop bool)

	// MediaKeys prop
	// js:"mediaKeys"
	MediaKeys() (mediaKeys *window.MediaKeys)

	// MsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
	// js:"msAudioCategory"
	MsAudioCategory() (msAudioCategory string)

	// MsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
	SetMsAudioCategory(msAudioCategory string)

	// MsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
	// js:"msAudioDeviceType"
	MsAudioDeviceType() (msAudioDeviceType string)

	// MsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
	SetMsAudioDeviceType(msAudioDeviceType string)

	// MsGraphicsTrustStatus prop
	// js:"msGraphicsTrustStatus"
	MsGraphicsTrustStatus() (msGraphicsTrustStatus *msgraphicstrust.MSGraphicsTrust)

	// MsKeys prop Gets the MSMediaKeys object, which is used for decrypting media data, that is associated with this media element.
	// js:"msKeys"
	MsKeys() (msKeys *msmediakeys.MSMediaKeys)

	// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
	// js:"msPlayToDisabled"
	MsPlayToDisabled() (msPlayToDisabled bool)

	// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
	SetMsPlayToDisabled(msPlayToDisabled bool)

	// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	// js:"msPlayToPreferredSourceUri"
	MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string)

	// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string)

	// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
	// js:"msPlayToPrimary"
	MsPlayToPrimary() (msPlayToPrimary bool)

	// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
	SetMsPlayToPrimary(msPlayToPrimary bool)

	// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
	// js:"msPlayToSource"
	MsPlayToSource() (msPlayToSource interface{})

	// MsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
	// js:"msRealTime"
	MsRealTime() (msRealTime bool)

	// MsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
	SetMsRealTime(msRealTime bool)

	// Muted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	// js:"muted"
	Muted() (muted bool)

	// Muted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	SetMuted(muted bool)

	// NetworkState prop Gets the current network activity for the element.
	// js:"networkState"
	NetworkState() (networkState uint8)

	// Onencrypted prop
	// js:"onencrypted"
	Onencrypted() (onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// Onencrypted prop
	SetOnencrypted(onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// Onmsneedkey prop
	// js:"onmsneedkey"
	Onmsneedkey() (onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// Onmsneedkey prop
	SetOnmsneedkey(onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// Paused prop Gets a flag that specifies whether playback is paused.
	// js:"paused"
	Paused() (paused bool)

	// PlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	// js:"playbackRate"
	PlaybackRate() (playbackRate float32)

	// PlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	SetPlaybackRate(playbackRate float32)

	// Played prop Gets TimeRanges for the current media resource that has been played.
	// js:"played"
	Played() (played *timeranges.TimeRanges)

	// Preload prop Gets or sets the current playback position, in seconds.
	// js:"preload"
	Preload() (preload string)

	// Preload prop Gets or sets the current playback position, in seconds.
	SetPreload(preload string)

	// ReadyState prop
	// js:"readyState"
	ReadyState() (readyState interface{})

	// Seekable prop Returns a TimeRanges object that represents the ranges of the current media resource that can be seeked.
	// js:"seekable"
	Seekable() (seekable *timeranges.TimeRanges)

	// Seeking prop Gets a flag that indicates whether the the client is currently moving to a new playback position in the media resource.
	// js:"seeking"
	Seeking() (seeking bool)

	// Src prop The address or URL of the a media resource that is to be considered.
	// js:"src"
	Src() (src string)

	// Src prop The address or URL of the a media resource that is to be considered.
	SetSrc(src string)

	// SrcObject prop
	// js:"srcObject"
	SrcObject() (srcObject *window.MediaStream)

	// SrcObject prop
	SetSrcObject(srcObject *window.MediaStream)

	// TextTracks prop
	// js:"textTracks"
	TextTracks() (textTracks *texttracklist.TextTrackList)

	// VideoTracks prop
	// js:"videoTracks"
	VideoTracks() (videoTracks *avtrack.VideoTrackList)

	// Volume prop Gets or sets the volume level for audio portions of the media element.
	// js:"volume"
	Volume() (volume float32)

	// Volume prop Gets or sets the volume level for audio portions of the media element.
	SetVolume(volume float32)
}
