package xmlhttprequestupload

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/dom2/xmlhttprequesteventtarget"
)

// js:"XMLHTTPRequestUpload,omit"
type XMLHTTPRequestUpload struct {
	window.EventTarget
	xmlhttprequesteventtarget.XMLHTTPRequestEventTarget
}
