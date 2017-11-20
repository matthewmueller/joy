package closeeventinit

type CloseEventInit struct {
	*EventInit

	code     *uint8
	reason   *string
	wasClean *bool
}
