package webglcontexteventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type WebGLContextEventInit struct {
	*eventinit.EventInit

	statusMessage *string
}
