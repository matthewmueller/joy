package progresseventinit

type ProgressEventInit struct {
	*EventInit

	lengthComputable *bool
	loaded           *uint64
	total            *uint64
}
