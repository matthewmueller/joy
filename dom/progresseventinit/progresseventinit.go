package progresseventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type ProgressEventInit struct {
	*eventinit.EventInit

	lengthComputable *bool
	loaded           *uint64
	total            *uint64
}
