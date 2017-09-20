Golly gee, it's a minimal Go to JS (ES3) compiler.

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