package childnode

import "github.com/matthewmueller/golly/js"

type ChildNode struct {
}

func (*ChildNode) Remove() {
	js.Rewrite("$<.remove()")
}
