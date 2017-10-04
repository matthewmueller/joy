package scope_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parse(t *testing.T, src string) *ast.File {
	fset := token.NewFileSet()
	f, e := parser.ParseFile(fset, "", src, 0)
	if e != nil {
		t.Fatal(e)
	}
	return f
}

// func TestFunctionScope(t *testing.T) {
// 	f := parse(t, `
// 		package main
// 		func main() {
// 			a := 3
// 		}
// 		`)

// 	filescope := scope.File(f.Scope)
// 	main := f.Decls[0].(*ast.FuncDecl)
// 	mainscope := filescope.New(main)

// 	a := main.Body.List[0].(*ast.AssignStmt).Lhs[0].(*ast.Ident)
// 	mainscope.Insert(a.Obj)

// 	assert.NotNil(t, filescope.Lookup("main"))
// 	assert.Nil(t, filescope.Lookup("a"))

// 	assert.NotNil(t, mainscope.Lookup("main"))
// 	assert.NotNil(t, mainscope.Lookup("a"))
// 	assert.Nil(t, mainscope.Within("main"))
// 	assert.NotNil(t, mainscope.Within("a"))

// 	assert.NotNil(t, mainscope.Outer.Lookup("main"))
// 	assert.Nil(t, mainscope.Outer.Lookup("a"))
// }

// func TestFunctionLit(t *testing.T) {
// 	f := parse(t, `
// 		package main
// 		func main() {
// 			b := 0
// 			func() {
// 				a := 3
// 			}()
// 		}
// 		`)

// 	filescope := scope.File(f.Scope)

// 	main := f.Decls[0].(*ast.FuncDecl)
// 	mainscope := filescope.New(main)

// 	b := main.Body.List[0].(*ast.AssignStmt).Lhs[0].(*ast.Ident)
// 	mainscope.Insert(b.Obj)

// 	anon := main.Body.List[1].(*ast.ExprStmt).X.(*ast.CallExpr).Fun.(*ast.FuncLit)
// 	anonscope := mainscope.New(anon)

// 	a := anon.Body.List[0].(*ast.AssignStmt).Lhs[0].(*ast.Ident)
// 	anonscope.Insert(a.Obj)

// 	assert.NotNil(t, filescope.Lookup("main"))
// 	assert.Nil(t, filescope.Lookup("a"))

// 	assert.NotNil(t, mainscope.Lookup("main"))
// 	assert.NotNil(t, mainscope.Lookup("b"))
// 	assert.Nil(t, mainscope.Lookup("a"))
// 	assert.NotNil(t, mainscope.Owner)

// 	assert.NotNil(t, anonscope.Lookup("main"))
// 	assert.NotNil(t, anonscope.Lookup("b"))
// 	assert.NotNil(t, anonscope.Lookup("a"))
// 	assert.NotNil(t, anonscope.Owner)

// 	assert.Nil(t, anonscope.Within("main"))
// 	assert.Nil(t, anonscope.Within("b"))
// 	assert.NotNil(t, anonscope.Lookup("a"))

// 	assert.NotNil(t, anonscope.Outer.Lookup("main"))
// 	assert.Nil(t, anonscope.Outer.Lookup("a"))

// 	assert.NotNil(t, anonscope.Outer.Outer.Lookup("main"))
// 	assert.Nil(t, anonscope.Outer.Outer.Lookup("a"))

// 	// assert.NotNil(t, mainscope.Outer.Lookup("main"))
// 	// assert.Nil(t, mainscope.Outer.Lookup("a"))
// }
