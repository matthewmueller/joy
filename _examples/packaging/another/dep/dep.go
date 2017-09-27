package dep

import (
	dep2 "github.com/matthewmueller/golly/_examples/packaging/another/more"
)

// AnotherGet fn
func AnotherGet() string {
	return "another world " + dep2.More()
}
