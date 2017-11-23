package webglcontexteventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type WebGLContextEventInit struct {
	*eventinit.EventInit

	statusMessage *string
}
