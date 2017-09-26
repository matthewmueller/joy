package golang

import (
	"go/build"
	"sort"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/js"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// CompilePackage compiles a package by it's path
func CompilePackage(packagePath string) error {
	var conf loader.Config
	conf.Import(packagePath)

	// TODO: clean up, this is ugly
	order := []string{}
	imports := map[string][]string{}

	// load each of the imports
	conf.FindPackage = func(ctx *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
		if imports[fromDir] == nil {
			order = append(order, fromDir)
		}
		imports[fromDir] = append(imports[fromDir], importPath)
		return ctx.Import(importPath, fromDir, mode)
	}

	// load the package
	pkgs, err := conf.Load()
	if err != nil {
		return errors.Wrap(err, "unable to load the go package")
	}

	// get a deterministic toposort of the imports
	// for _, dep := range deps {
	// 	log.Infof("dep: %s", dep)
	// }

	// get the topological sort of the dependencies
	var deps []string
	for _, o := range order {
		lvl := imports[o]
		sort.Strings(lvl)
		deps = append(deps, lvl...)
	}
	deps = reverse(deps)
	// sort.Reverse(deps)
	// deps = strings.Re

	for _, dep := range deps {
		log.Infof("dep: %s", dep)
	}

	var body []interface{}
	var stmts []interface{}

	// translate each file to their Javascript counterpart
	for _, info := range pkgs.AllPackages {
		// log.Infof("building... %s", pkg.Name())
		for _, file := range info.Files {
			// log.Infof("translating... %s", info.Pkg.Path())
			stmt, e := Translate(info, file)
			if e != nil {
				return errors.Wrapf(e, "error translating %s", file.Name.Name)
			}

			filepath := info.Pkg.Path()
			log.Infof("filepath %s", filepath)
			stmts = append(stmts, stmt)
			// log.Infof("translated: %s\n%s", file.Name.Name, ast)
		}

		// TODO: bundle each file by package

	}

	init := js.CreateVariableDeclaration(
		"var",
		js.CreateVariableDeclarator(
			js.CreateIdentifier("pkg"),
			js.CreateObjectExpression([]js.Property{}),
		),
	)

	callmain := js.CreateExpressionStatement(
		js.CreateCallExpression(
			js.CreateMemberExpression(
				js.CreateMemberExpression(
					js.CreateIdentifier("pkg"),
					js.CreateString(`"`+packagePath+`"`),
					true,
				),
				js.CreateIdentifier("main"),
				false,
			),
			[]js.IExpression{},
		),
	)

	body = append(body, init)
	body = append(body, stmts...)
	body = append(body, callmain)

	prog := js.CreateProgram(
		js.CreateExpressionStatement(
			js.CreateCallExpression(
				js.CreateFunctionExpression(nil, []js.IPattern{},
					js.CreateFunctionBody(body...),
				),
				[]js.IExpression{},
			),
		),
	)
	_ = prog
	// log.Infof("%s", prog)

	return nil
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
