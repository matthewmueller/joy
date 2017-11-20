package mediaencryptedeventinit

type MediaEncryptedEventInit struct {
	*EventInit

	initData     *[]byte
	initDataType *string
}
