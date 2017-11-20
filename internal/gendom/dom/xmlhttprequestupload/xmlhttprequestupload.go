package xmlhttprequestupload

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/xmlhttprequesteventtarget"
)

type XMLHttpRequestUpload struct {
	*eventtarget.EventTarget
	*xmlhttprequesteventtarget.XMLHttpRequestEventTarget
}
