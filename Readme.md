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

## Open Questions

- How does the JS scope differ from the Go scope and what do I need to account for?
- How will we deal with existing Go standard libraries?
- How will the different base Go types map to javascript types? (like maps)
  - This is where Golly may have to be more strict than Go (e.g. map[boolean]string)
- How will we make DOM typesafe, but not include it in the build?
  - DOM will be an API like js/api, but emit JS AST nodes!
  - Javascripter interface to know if Golly should turn into an AST or not!
- Do Go's interfaces show up in the emitted Javascript in any form?

## High-level TODO:

Files:

- [x] implement if statements
- [x] implement loops
- [x] implement nested structs
- [x] implement classes
- [ ] implement arrays
- [ ] implement maps
- [ ] implement built-in functions (append, copy, len, make)
- [ ] implement spreads (...)
- [ ] implement functions with multiple return values
- [ ] implement switch statements
- [ ] breaks, continues, panics, 

Packages:

- [ ] support multiple files
- [ ] write basic DOM package using Typescript's libdefs

Hard:

- [ ] support goroutines and channels (using async/await)
- [ ] implement scope for variable rewrites

Before I can use

- [ ] write preact package
- [ ] write fetch package
- [ ] basic development server

Later:

- [ ] implement defer function () { ... }
- [ ] recover statement
- [ ] goto statements
- [ ] support select statements

## Design

- Small browser builds that run in all browsers
- Compiles to ES3, the most supported JS dialect
- The "bucklescript" for Go
- Fast builds

## Wait, why?

- Even with WebAssembly, Javascript will be supported for years to come
- Go has a minimal syntax and clear future priorities
- Go is both easy to read and easy to maintain
- Existing Go implementations 
- Golly code runs in Go (not necessarily the other way around, see below)

## Tradeoffs

**Golly is learn once, write everywhere. Not a write once, run everywhere**

Just because your code is running in Go does not mean it will run on Golly without modification. For everything that browserify & webpack brought to the community,they also brought with it mysteriously large builds.

Golly aims to generate clean, minimal & performant ES3 code that works in every browser. It will not jump hoops to make your Go code work inside the browser. GopherJS is much better suited for that task.