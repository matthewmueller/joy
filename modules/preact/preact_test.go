package preact_test

import (
	"testing"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/dom"
	"github.com/matthewmueller/golly/modules/preact"
)

func TestUsage(t *testing.T) {
	document := dom.Document{}
	page := &Page{}
	preact.Render(preact.C(page, nil), document.Body, nil)
}

func NewPage(props preact.IProps) preact.IComponent {
	return &Page{}
}

type Page struct {
	state preact.IState
}

func (p *Page) Render(props preact.IProps, state preact.IState) *preact.VNode {
	log.Infof("rendering page...")
	return nil
}
