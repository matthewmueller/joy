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

func (a *app) loadFirst(nth int) ([]item.Item, error) {
	res, err := fetch.Get(api + "/v0/topstories.json")
	if err != nil {
		return nil, err
	}

	var itemIDs []int
	if err := res.JSON(&itemIDs); err != nil {
		return nil, err
	}

	// get each item's details
	// TODO: waitgroup
	var items []item.Item
	for _, id := range itemIDs[:nth] {
		res, err := fetch.Get(api + "/v0/item/" + strconv.Itoa(id) + ".json")
		if err != nil {
			return nil, err
		}

		var it item.Item
		if err := res.JSON(&it); err != nil {
			return nil, err
		}
		items = append(items, it)
	}

	return items, nil
}

// js:"componentDidMount"
func (a *app) ComponentDidMount() {
	println("loading first...")
	items, err := a.loadFirst(2)
	if err != nil {
		panic(err)
	}

	println("setting state...")
	a.SetState(state{
		items: items,
	})
}

// Render
// js:"render"
func (a *app) Render() jsx.JSX {
	var items []jsx.Node
	// for _, item := range a.state.items {
	// 	// items = append(items, jsx.T(`
	// 	// 	<div>{item.ID}</div>
	// 	// `))
	// 	items = append(items, jsx.H("div", map[string]interface{}{
	// 		"key": item.ID,
	// 	}, &jsx.Text{Value: item.Title}))
	// }

	// println(items)
	// // return jsx.T(`
	// // 	<div>{}</div>
	// // `)
	children := append([]jsx.Node{&jsx.Text{Value: "Hi"}}, items...)
	return jsx.H("div", nil, children...)
}
