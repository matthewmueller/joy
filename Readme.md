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

## High-level:

- [x] implement if statements
- [ ] implement loops
- [ ] implement structs
- [ ] implement classes
- [ ] implement switch statements
- [ ] implement scope
- [ ] implement packaging
- [ ] write basic document package
- [ ] write preact package
- [ ] write async 
- [ ] write fetch package

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