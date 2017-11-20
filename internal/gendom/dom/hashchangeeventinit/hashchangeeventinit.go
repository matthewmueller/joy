package hashchangeeventinit

type HashChangeEventInit struct {
	*EventInit

	newURL *string
	oldURL *string
}
