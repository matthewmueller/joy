package golang_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/matthewmueller/golly/golang"
	"github.com/stretchr/testify/assert"
)

// TODO: consts
// TODO: iotas

func parse(t *testing.T, src string) *ast.File {
	fset := token.NewFileSet()
	f, e := parser.ParseFile(fset, "", src, 0)
	if e != nil {
		t.Fatal(e)
	}
	return f
}

func toString(t *testing.T, n interface{}) string {
	s, ok := n.(fmt.Stringer)
	if !ok {
		t.Fatal("not a stringer: ", n)
	}
	return s.String()
}

func TestVarSingle(t *testing.T) {
	f := parse(t, `
		package main
		var a = 3
	`)

	stmt, expr, err := golang.VariableHandler(f.Decls[0])
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, "var a = 3", toString(t, stmt))
}

func TestVarZero(t *testing.T) {
	f := parse(t, `
		package main
		var a int
	`)

	stmt, expr, err := golang.VariableHandler(f.Decls[0])
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, "var a = 0", toString(t, stmt))
}

func TestVarTypeSingle(t *testing.T) {
	f := parse(t, `
		package main
		var a int = 3
	`)

	stmt, expr, err := golang.VariableHandler(f.Decls[0])
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, "var a = 3", toString(t, stmt))
}

func TestVarMultipleSimple(t *testing.T) {
	f := parse(t, `
		package main
		var (
			a int = 3
			b int = 6
		)
	`)

	stmt, expr, err := golang.VariableHandler(f.Decls[0])
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, "var a = 3, b = 6", toString(t, stmt))
}

func TestVarMultipleZero(t *testing.T) {
	f := parse(t, `
		package main
		var (
			a int
			b []string
			c string
		)
	`)

	stmt, expr, err := golang.VariableHandler(f.Decls[0])
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, `var a = 0, b = [], c = ""`, toString(t, stmt))
}

func TestVarMultipleMixed(t *testing.T) {
	f := parse(t, `
		package main
		var (
			a int = 5
			b []string
			c = "c"
		)
	`)

	stmt, expr, err := golang.VariableHandler(f.Decls[0])
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, `var a = 5, b = [], c = "c"`, toString(t, stmt))
}

func TestShortDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			a := 3
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, `var a = 3`, toString(t, stmt))
}

func TestBalancedDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j := 0, 10
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, `var i = 0, j = 10`, toString(t, stmt))
}

func TestUnbalancedDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j := pipe
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, `var $i = pipe, i = $i[0], j = $i[1]`, toString(t, stmt))
}

func TestUnderscoreDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			_, y, _ := coord
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if expr != nil {
		t.Fatal("expr should be nil")
	}

	assert.Equal(t, `var $y = coord, y = $y[1]`, toString(t, stmt))
}

func TestShortAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			a = 3
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if stmt != nil {
		t.Fatal("stmt should be nil")
	}

	assert.Equal(t, `a = 3`, toString(t, expr))
}

func TestBalancedAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j = 0, 10
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if stmt != nil {
		t.Fatal("stmt should be nil")
	}

	assert.Equal(t, `i = 0, j = 10`, toString(t, expr))
}

func TestUnbalancedAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j = pipe
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `var $i = pipe`, toString(t, stmt))
	assert.Equal(t, `i = $i[0], j = $i[1]`, toString(t, expr))
}

func TestUnderscoreAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			_, y, _ = coord
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `var $y = coord`, toString(t, stmt))
	assert.Equal(t, `y = $y[1]`, toString(t, expr))
}

func TestAllBlank(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			_, _ = coord, tester
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if stmt != nil {
		t.Fatal("stmt should be nil")
	}

	assert.Equal(t, `coord, tester`, toString(t, expr))
}

func TestShortSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			a.b = 3
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if stmt != nil {
		t.Fatal("stmt should be nil")
	}

	assert.Equal(t, `a.b = 3`, toString(t, expr))
}

func TestBalancedSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i.a, j.b = 0, 10
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	} else if stmt != nil {
		t.Fatal("stmt should be nil")
	}

	assert.Equal(t, `i.a = 0, j.b = 10`, toString(t, expr))
}

func TestUnbalancedSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i.a, j.b = pipe
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `i.a = $i[0], j.b = $i[1]`, toString(t, expr))
	assert.Equal(t, `var $i = pipe`, toString(t, stmt))
}

func TestUnderscoreSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			_, y.b, _ = coord
		}
	`)

	main := f.Decls[0].(*ast.FuncDecl).Body.List[0]
	stmt, expr, err := golang.VariableHandler(main)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `var $y = coord`, toString(t, stmt))
	assert.Equal(t, `y.b = $y[1]`, toString(t, expr))
}
