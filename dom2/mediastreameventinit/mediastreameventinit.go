package mediastreameventinit

type MediaStreamEventInit struct {
	*eventinit.EventInit

	stream *window.MediaStream
}
