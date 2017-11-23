package mediaencryptedeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type MediaEncryptedEventInit struct {
	*eventinit.EventInit

	initData     *[]byte
	initDataType *string
}
