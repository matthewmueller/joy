package webglcontexteventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type WebGLContextEventInit struct {
	*eventinit.EventInit

	statusMessage *string
}
