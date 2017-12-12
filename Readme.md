![joy compiler](https://user-images.githubusercontent.com/170299/33872480-b10be348-df49-11e7-80b1-06736c8298ae.png)

Translate idiomatic Go into concise Javascript that works in every browser. Use Go's type system and world-class tooling to build large web applications with confidence.

Visit [mat.tm/joy](https://mat.tm/joy) to learn more about Joy.

---

[Getting Started](#getting-started)
&nbsp;&nbsp;&nbsp;&nbsp;&#183;&nbsp;&nbsp;&nbsp;&nbsp;[Examples](#examples)
&nbsp;&nbsp;&nbsp;&nbsp;&#183;&nbsp;&nbsp;&nbsp;&nbsp;[Using the CLI](#using-the-cli)

[Contributing](#contributing)&nbsp;&nbsp;&nbsp;&nbsp;&#183;&nbsp;&nbsp;&nbsp;&nbsp;[FAQ](#faq)&nbsp;&nbsp;&nbsp;&nbsp;&#183;&nbsp;&nbsp;&nbsp;&nbsp;[More Links](#more-links)

**GoDoc:**

[Joy API](https://godoc.org/github.com/matthewmueller/joy)&nbsp;&nbsp;&nbsp;&nbsp;&#183;&nbsp;&nbsp;&nbsp;&nbsp;[DOM API](https://godoc.org/github.com/matthewmueller/joy/dom)&nbsp;&nbsp;&nbsp;&nbsp;&#183;&nbsp;&nbsp;&nbsp;&nbsp;[Virtual DOM API](https://godoc.org/github.com/matthewmueller/joy/vdom)

---

## Getting Started

**1. Install Joy:**

```sh
curl -sfL https://raw.githubusercontent.com/matthewmueller/joy/master/install.sh | sh
```

> Note: you can also download from the [releases tab](https://github.com/matthewmueller/joy/releases).

**2. Create a `main.go` file with the following code:**

```go
package main

func main() {
  println("hi world!")
}
```

**3. Run the code in a real browser:**

```sh
joy run main.go
```

**4. See the compiled Javascript:**

```sh
joy main.go
```

## Examples

Visit https://mat.tm/joy/#examples or peruse the [test suite](https://github.com/matthewmueller/joy/tree/master/testdata).

## Using the CLI

**Compile Go into Javascript:**

```
joy <main.go>
```

**Compile and run the Go code in headless chrome:**

```
joy run <main.go>
```

**Build a development version of the code:**

```
joy build --dev <main.go>...
```

**Build a production version of the code (coming soon!):**

```
joy build <main.go>...
```

**Start a development server with livereload:**

```
joy serve <main.go>...
```

> Run `joy help` for additional details.

## Contributing

So happy to hear you're interested in contributing! Here's a quick rundown of how to get setup:

**Setup**

1. Make sure you have the Go environment setup on your computer. There are quite a few better resources online on how to do that

2. `go get -u -v github.com/matthewmueller/joy/...` to install the compiler from source

3. `go test -v` to run all the tests

**Links and tips:**

- https://astexplorer.net/: I've been using this to figure out how to build the JS tree
- http://goast.yuroyoro.net/ : Simple Go AST viewer
- https://github.com/estree/estree/blob/master/es5.md: Link to the ES3 AST format, this is implemented in [syntax.go](internal/jsast/syntax.go)
- https://golang.org/ref/spec: You may need to refer to this to see what types are possible in Go's AST
- To run all tests: `go test -v`
- To run individual tests: `go test -v -run Test/08`
- `pretty.Println(ast)` will pretty print the JS AST (requires [this package](https://github.com/kr/pretty))
- `ast.Print(nil, node)` will pretty print the Go AST

If you have any further questions, [open an issue](github.com/matthewmueller/joy/issues) or reach out to me on [twitter](https://twitter.com/mattmueller).

## FAQ

Visit https://mat.tm/joy/#faq to view the FAQ.

## More Links

- Run `joy help` to see what else Joy can do for you
- Visit mat.tm/joy to read more about Joy's origins
- Chat with us in **#joy-compiler** on Slack at gophers.slack.com
- Star github.com/matthewmueller/joy to follow the development
- Follow twitter.com/mattmueller for project updates
