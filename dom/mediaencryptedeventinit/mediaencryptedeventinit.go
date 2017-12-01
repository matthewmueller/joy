package mediaencryptedeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type MediaEncryptedEventInit struct {
	*eventinit.EventInit

	initData     *[]byte
	initDataType *string
}
