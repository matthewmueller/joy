package htmlmediaelement

import (
	"github.com/matthewmueller/golly/dom2/mediaencryptedevent"
	"github.com/matthewmueller/golly/dom2/mediaerror"
	"github.com/matthewmueller/golly/dom2/texttrack"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"HTMLMediaElement,omit"
type HTMLMediaElement interface {
	window.HTMLElement

	// AddTextTrack
	AddTextTrack(kind string, label *string, language *string) (t *texttrack.TextTrack)

	// CanPlayType Returns a string that specifies whether the client can play a given media resource type.
	CanPlayType(kind string) (s string)

	// Load Resets the audio or video object and loads a new media resource.
	Load()

	// MsClearEffects Clears all effects from the media pipeline.
	MsClearEffects()

	// MsGetAsCastingSource
	MsGetAsCastingSource() (i interface{})

	// MsInsertAudioEffect Inserts the specified audio effect into media pipeline.
	MsInsertAudioEffect(activatableClassId string, effectRequired bool, config *interface{})

	// MsSetMediaKeys
	MsSetMediaKeys(mediaKeys *msmediakeys.MSMediaKeys)

	// MsSetMediaProtectionManager Specifies the media protection manager for a given media pipeline.
	MsSetMediaProtectionManager(mediaProtectionManager *interface{})

	// Pause Pauses the current playback and sets paused to TRUE. This can be used to test whether the media is playing or paused. You can also use the pause or play events to tell whether the media is playing or not.
	Pause()

	// Play Loads and starts playback of a media resource.
	Play()

	// SetMediaKeys
	SetMediaKeys(mediaKeys *window.MediaKeys)

	// AudioTracks Returns an AudioTrackList object with the audio tracks for a given video element.
	AudioTracks() (audioTracks *avtrack.AudioTrackList)

	// Autoplay Gets or sets a value that indicates whether to start playing the media automatically.
	Autoplay() (autoplay bool)

	// Autoplay Gets or sets a value that indicates whether to start playing the media automatically.
	SetAutoplay(autoplay bool)

	// Buffered Gets a collection of buffered time ranges.
	Buffered() (buffered *timeranges.TimeRanges)

	// Controls Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	Controls() (controls bool)

	// Controls Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
	SetControls(controls bool)

	// CrossOrigin
	CrossOrigin() (crossOrigin string)

	// CrossOrigin
	SetCrossOrigin(crossOrigin string)

	// CurrentSrc Gets the address or URL of the current media resource that is selected by IHTMLMediaElement.
	CurrentSrc() (currentSrc string)

	// CurrentTime Gets or sets the current playback position, in seconds.
	CurrentTime() (currentTime float32)

	// CurrentTime Gets or sets the current playback position, in seconds.
	SetCurrentTime(currentTime float32)

	// DefaultMuted
	DefaultMuted() (defaultMuted bool)

	// DefaultMuted
	SetDefaultMuted(defaultMuted bool)

	// DefaultPlaybackRate Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	DefaultPlaybackRate() (defaultPlaybackRate float32)

	// DefaultPlaybackRate Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
	SetDefaultPlaybackRate(defaultPlaybackRate float32)

	// Duration Returns the duration in seconds of the current media resource. A NaN value is returned if duration is not available, or Infinity if the media resource is streaming.
	Duration() (duration float32)

	// Ended Gets information about whether the playback has ended or not.
	Ended() (ended bool)

	// Error Returns an object representing the current error state of the audio or video element.
	Error() (err *mediaerror.MediaError)

	// Loop Gets or sets a flag to specify whether playback should restart after it completes.
	Loop() (loop bool)

	// Loop Gets or sets a flag to specify whether playback should restart after it completes.
	SetLoop(loop bool)

	// MediaKeys
	MediaKeys() (mediaKeys *window.MediaKeys)

	// MsAudioCategory Specifies the purpose of the audio or video media, such as background audio or alerts.
	MsAudioCategory() (msAudioCategory string)

	// MsAudioCategory Specifies the purpose of the audio or video media, such as background audio or alerts.
	SetMsAudioCategory(msAudioCategory string)

	// MsAudioDeviceType Specifies the output device id that the audio will be sent to.
	MsAudioDeviceType() (msAudioDeviceType string)

	// MsAudioDeviceType Specifies the output device id that the audio will be sent to.
	SetMsAudioDeviceType(msAudioDeviceType string)

	// MsGraphicsTrustStatus
	MsGraphicsTrustStatus() (msGraphicsTrustStatus *msgraphicstrust.MSGraphicsTrust)

	// MsKeys Gets the MSMediaKeys object, which is used for decrypting media data, that is associated with this media element.
	MsKeys() (msKeys *msmediakeys.MSMediaKeys)

	// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
	MsPlayToDisabled() (msPlayToDisabled bool)

	// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
	SetMsPlayToDisabled(msPlayToDisabled bool)

	// MsPlayToPreferredSourceURI Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string)

	// MsPlayToPreferredSourceURI Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
	SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string)

	// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
	MsPlayToPrimary() (msPlayToPrimary bool)

	// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
	SetMsPlayToPrimary(msPlayToPrimary bool)

	// MsPlayToSource Gets the source associated with the media element for use by the PlayToManager.
	MsPlayToSource() (msPlayToSource interface{})

	// MsRealTime Specifies whether or not to enable low-latency playback on the media element.
	MsRealTime() (msRealTime bool)

	// MsRealTime Specifies whether or not to enable low-latency playback on the media element.
	SetMsRealTime(msRealTime bool)

	// Muted Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	Muted() (muted bool)

	// Muted Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
	SetMuted(muted bool)

	// NetworkState Gets the current network activity for the element.
	NetworkState() (networkState uint8)

	// Onencrypted
	Onencrypted() (onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// Onencrypted
	SetOnencrypted(onencrypted func(*mediaencryptedevent.MediaEncryptedEvent))

	// Onmsneedkey
	Onmsneedkey() (onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// Onmsneedkey
	SetOnmsneedkey(onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent))

	// Paused Gets a flag that specifies whether playback is paused.
	Paused() (paused bool)

	// PlaybackRate Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	PlaybackRate() (playbackRate float32)

	// PlaybackRate Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
	SetPlaybackRate(playbackRate float32)

	// Played Gets TimeRanges for the current media resource that has been played.
	Played() (played *timeranges.TimeRanges)

	// Preload Gets or sets the current playback position, in seconds.
	Preload() (preload string)

	// Preload Gets or sets the current playback position, in seconds.
	SetPreload(preload string)

	// ReadyState
	ReadyState() (readyState interface{})

	// Seekable Returns a TimeRanges object that represents the ranges of the current media resource that can be seeked.
	Seekable() (seekable *timeranges.TimeRanges)

	// Seeking Gets a flag that indicates whether the the client is currently moving to a new playback position in the media resource.
	Seeking() (seeking bool)

	// Src The address or URL of the a media resource that is to be considered.
	Src() (src string)

	// Src The address or URL of the a media resource that is to be considered.
	SetSrc(src string)

	// SrcObject
	SrcObject() (srcObject *window.MediaStream)

	// SrcObject
	SetSrcObject(srcObject *window.MediaStream)

	// TextTracks
	TextTracks() (textTracks *texttracklist.TextTrackList)

	// VideoTracks
	VideoTracks() (videoTracks *avtrack.VideoTrackList)

	// Volume Gets or sets the volume level for audio portions of the media element.
	Volume() (volume float32)

	// Volume Gets or sets the volume level for audio portions of the media element.
	SetVolume(volume float32)
}
