package mutationobserver

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mutationobserverinit"
	"github.com/matthewmueller/golly/internal/gendom/dom/mutationrecord"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type MutationObserver struct {
}

func (*MutationObserver) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

func (*MutationObserver) Observe(target *node.Node, options *mutationobserverinit.MutationObserverInit) {
	js.Rewrite("$<.observe($1, $2)", target, options)
}

func (*MutationObserver) TakeRecords() (m []*mutationrecord.MutationRecord) {
	js.Rewrite("$<.takeRecords()")
	return m
}
