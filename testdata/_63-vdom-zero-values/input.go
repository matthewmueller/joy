package main

import (
	"strconv"
	"strings"

	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/vdom"
	"github.com/matthewmueller/golly/vdom/h/strong"
)

// User struct
type User struct {
	vdom.Component

	props *props
}

type props struct {
	FirstName string
	LastName  string
	Age       int
	Phone     Phone
}

// Phone struct
type Phone struct {
	Type   string
	Number string
}

// Render fn
func (u *User) Render() vdom.Node {
	if u.props.Age == 0 {
		u.props.Age = 28
	}

	if u.props.FirstName == "" {
		u.props.FirstName = "matt"
	}

	if u.props.Phone.Number == "" {
		u.props.Phone.Number = "555-5555"
	}

	props := []string{u.props.FirstName, strconv.Itoa(u.props.Age), u.props.Phone.Number, u.props.Phone.Type}
	return strong.New(nil, vdom.S(strings.Join(props, " ")))
}

func main() {
	vdom.Use("preact.h", "./preact.js")
	user := &User{props: &props{}}
	vdom.Render(user, document.Body, nil)
	println(document.Body.InnerHTML())
	user = &User{}
	vdom.Render(user, document.Body, document.Body.FirstElementChild())
	println(document.Body.InnerHTML())
}
