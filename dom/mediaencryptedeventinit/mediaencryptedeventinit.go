package mediaencryptedeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type MediaEncryptedEventInit struct {
	*eventinit.EventInit

	initData     *[]byte
	initDataType *string
}
