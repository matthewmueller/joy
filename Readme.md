Golly gee, it's another Go to JS compiler.

![img](https://cldup.com/uQb67D_DJT.png)

## Helpful Resources

- https://astexplorer.net/: I've been using this to figure out how to build the JS tree
- http://goast.yuroyoro.net/ : Simple Go AST viewer
- https://github.com/estree/estree/blob/master/es5.md: Link to the ES3 AST format, this is implemented in syntax.go
- https://golang.org/ref/spec: You may need to refer to this to see what types are possible in Go's AST
- To run all tests: `go test -v`
- To run individual tests: `go test -v -run Test/08`
- To run examples: `go run cmd/golly/main.go --pkg $(PWD)/_examples/dom`
- To see a callgraph: `go run cmd/golly/main.go --pkg $(PWD)/_examples/dom --graph`
- `pretty.Println(ast)` will pretty print the JS AST
- `ast.Print(nil, node)` will pretty print the Go AST

## Design

- Dead-code elimination across all your dependencies
- Simple interface between untyped JS and typed Go
- Compiles to ES3, the most universal JS dialect
- Idiomatic Go to readable Javascript
- Fast builds

## Wait, why?

- Javascript has become increasingly difficult to package up and distribute
- Despite all the hard work that's gone into Javascript tooling, build systems and transpilers remain slow
- Even when there's widespread WebAssembly adoption, Javascript will be supported for years to come
- Go is a minimal language with a simple type system and clear future priorities around better tools
- Go's type system and tooling make Go programs easy to understand & maintain

## Tradeoffs

**Go will compile Golly code, but Golly will probably not compile existing Go code**

Just because your code is running in Go does not mean it will run on Golly without modification. For everything that browserify & webpack brought to the community, they also brought with it mysteriously large builds.

Golly aims to generate clean, minimal & performant ES3 code that works in every browser. It will not jump hoops to make your Go code work inside the browser.

If you need browser compatibility with existing Go code, GopherJS is much better suited for that task.

**Golly doesn't include Go's wonderful standard library**

All 3rd party packages are opt-in. This is to make sure that our builds remain purposeful and small. Over time, Golly will add compatibility with certain standard library packages where it makes sense.

**Golly aims to be a browser solution, not a node.js solution**

While there's a good chance Golly programs will work in node.js, it's not the focus of this project. The stack we're going for here is Go on the backend and Go on the frontend compiled with Golly.

## Additional Motivation

The source code to metadata ratio is out of wack nowadays.

![img](https://cldup.com/4EtJ3jqzdw.png)

Let's change this.