package golang_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
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

func TestVarSingle(t *testing.T) {
	f := parse(t, `
		package main
		var a = 3
	`)

	ast.Print(nil, f)
}

func TestVarZero(t *testing.T) {
	f := parse(t, `
		package main
		var a int
	`)

	ast.Print(nil, f)
}

func TestVarTypeSingle(t *testing.T) {
	f := parse(t, `
		package main
		var a int = 3
	`)

	ast.Print(nil, f)
}

func TestVarMultiple(t *testing.T) {
	f := parse(t, `
		package main
		var (
			a int = 3
			b int = 6
		)
	`)

	ast.Print(nil, f)
}

func TestVarMultipleZero(t *testing.T) {
	f := parse(t, `
		package main
		var (
			a int
			b []string
		)
	`)

	ast.Print(nil, f)
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

	ast.Print(nil, f)
}

func TestShortDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			a := 3
		}
	`)

	ast.Print(nil, f)
}

func TestBalancedDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j := 0, 10
		}
	`)

	ast.Print(nil, f)
}

func TestUnbalancedDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j := pipe(fd)
		}
	`)

	ast.Print(nil, f)
}

func TestUnderscoreDecl(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			_, y, _ := coord(p)
		}
	`)

	ast.Print(nil, f)
}

func TestShortAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			a = 3
		}
	`)

	ast.Print(nil, f)
}

func TestBalancedAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j = 0, 10
		}
	`)

	ast.Print(nil, f)
}

func TestUnbalancedAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i, j = pipe(fd)
		}
	`)

	ast.Print(nil, f)
}

func TestUnderscoreAssign(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			_, y, _ = coord(p)
		}
	`)

	ast.Print(nil, f)
}

func TestShortSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			a.b = 3
		}
	`)

	ast.Print(nil, f)
}

func TestBalancedSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i.a, j.b = 0, 10
		}
	`)

	ast.Print(nil, f)
}

func TestUnbalancedSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			i.a, j.b = pipe(fd)
		}
	`)

	ast.Print(nil, f)
}

func TestUnderscoreSel(t *testing.T) {
	f := parse(t, `
		package main
		func main() {
			_, y.b, _ = coord(p)
		}
	`)

	ast.Print(nil, f)
}
