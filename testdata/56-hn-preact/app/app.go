package app

import (
	"strconv"

	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/56-hn-preact/fetch"
	"github.com/matthewmueller/golly/testdata/56-hn-preact/item"
	"github.com/matthewmueller/golly/testdata/56-hn-preact/preact"
)

var api = "https://hacker-news.firebaseio.com"

type app struct {
	preact.Component

	props props
	state state
}

type props struct {
}

type state struct {
	items []item.Item
}

// New app
func New() jsx.Node {
	a := &app{}
	// TODO: mark these are present in the build
	// or figure out how to include these in preact
	_ = a.Render
	_ = a.ComponentDidMount
	return a
}

func (a *app) loadItems() {
	res, err := fetch.Get(api + "/v0/topstories.json")
	if err != nil {
		println(err)
		return
	}

	var itemIDs []int
	if err := res.JSON(&itemIDs); err != nil {
		return
	}

	// get each item's details
	var items []item.Item
	for _, id := range itemIDs {
		res, err := fetch.Get(api + "/v0/items/" + strconv.Itoa(id) + ".json")
		if err != nil {
			println(err)
			return
		}

		var item item.Item
		if err := res.JSON(&item); err != nil {
			return
		}
		items = append(items, item)
	}

	a.SetState(&state{
		items: items,
	})
}

// js:"componentDidMount"
func (a *app) ComponentDidMount() {
	println("component did mount...")
	a.loadItems()
}

// Render
// js:"render"
func (a *app) Render() jsx.JSX {
	println("here")
	return jsx.H("div", nil, &jsx.Text{Value: "hi..."})
}
