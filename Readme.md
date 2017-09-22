Golly gee, it's another Go to JS compiler.

```go
program := js.CreateProgram(
  js.CreateEmptyStatement(),
  js.CreateExpressionStatement(
    js.CreateCallExpression(
      js.CreateFunctionExpression(nil, []js.IPattern{},
        js.CreateFunctionBody(
          js.CreateExpressionStatement(
            js.CreateCallExpression(
              js.CreateMemberExpression(
                js.CreateIdentifier("console"),
                js.CreateIdentifier("log"),
                false,
              ),
              []js.IExpression{
                js.CreateString("hi world!"),
              },
            ),
          ),
        ),
      ),
      []js.IExpression{},
    ),
  ),
)
```

## Directory Structure

```
js/
  syntax.go: this is the AST format for Javascript (ES3) (TODO: move to internal/)
  assemble.go: this will turn JS AST nodes into Javascript source code (TODO: move to internal/)
  api.go: this is the public interface for creating JS AST nodes
translator/
  translate.go: this translates our Go AST into a JS AST
nodejs/
  node: the oldest node binary I could compile (0.6.0). no promises, generators, async/await. I've just been using this for testing
regenerator/: this just is some test files for trying to wrap my head around how this transform works
golly.go: public API for golly
cmd/golly/main.go: CLI for golly
```

## Open Questions

- How does the JS scope differ from the Go scope and what do I need to account for?
- How will we deal with existing Go standard libraries?
- How will the different base Go types map to javascript types? (like maps)
  - This is where Golly may have to be more strict than Go (e.g. map[boolean]string)
- How will we make DOM typesafe, but not include it in the build?
  - DOM will be by default be an server-side API (like jsdom), but include interfaces to JS AST Nodes that Golly will call into over transpiling the Go code into JS
- Do Go's interfaces show up in the emitted Javascript in any form?
- Is it possible to "upgrade" an existing Go package to use Golly when you don't control the package's source, e.g. "fmt"?
- Should we rename this project to Jolly? :-D

## High-level TODO:

File-Level:

- [x] implement if statements
- [x] implement loops
- [x] implement nested structs
- [x] implement classes
- [x] implement arrays
- [x] implement maps
- [ ] implement functions with multiple return values
- [ ] anonymous functions
- [ ] implement built-in functions (append, copy, len, make)
- [ ] implement spreads (users...)
- [ ] implement switch statements
- [ ] handle User{"a"}
- [ ] breaks, continues, panics
- [ ] custom types (type Blah = string)
- [ ] global variables, constants, etc.

Package-Level:

- [ ] support multiple files
  - files are wrapped in closures and the package wraps the file closures in another closure calling main()
  - we can be clever and do topological dependency sorting to eliminate this later (look to rollup for inspiration here)
- [ ] write basic DOM package using Typescript's libdefs
  - Source: https://github.com/Microsoft/TypeScript/blob/master/lib/lib.dom.d.ts
  - Typescript has a script to generate these files, so maybe we can use that 

Hard but probably necessary:

- [ ] support goroutines and channels (using async/await)
  - we need some async functionality for things like AJAX
  - probably makes sense to prototype this as a babel transform first 
  - for cross-browser we can run it through this: https://github.com/facebook/regenerator
    - this process relies on a Promise polyfill, but if written
    to use generators, we no longer need promises.
- [ ] handle differences in variable scope between JS <-> Golang

Before I can use:

- [ ] write preact package (~1000LOC)
- [ ] write fetch package (https://github.com/developit/unfetch ~50LOC)
- [ ] basic development server (livereload is fine for now)

Later:

- [ ] add spacing (space, tab, newline) tokens into the JS AST to improve our JS output format
- [ ] implement defer function () { ... }
- [ ] error out on maps without key strings
- [ ] recover statement
- [ ] goto + label statements
- [ ] support go's select statement

## Design

- Small single-file builds that run in all browsers
- Compiles to ES3, the most popular JS dialect
- The "bucklescript" for Go
- Fast builds

## Wait, why?

- Javascript has become increasingly difficult to package up and distribute.
- Even with WebAssembly, Javascript will be supported for years to come
- Go has a minimal syntax and clear future priorities
- Go is both easy to read and easy to maintain
- Golly code runs in Go (not necessarily the other way around, see below)

## Tradeoffs

**Golly is learn once, write everywhere. Not a write once, run everywhere**

Just because your code is running in Go does not mean it will run on Golly without modification. For everything that browserify & webpack brought to the community,they also brought with it mysteriously large builds.

Golly aims to generate clean, minimal & performant ES3 code that works in every browser. It will not jump hoops to make your Go code work inside the browser. GopherJS is much better suited for that task.

**Golly doesn't include Go's standard library**

All 3rd party packages are opt-in. This is to make sure that our builds remain purposeful and small.