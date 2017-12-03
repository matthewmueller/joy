package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/matthewmueller/golly/internal/chrome"
	"github.com/matthewmueller/golly/internal/compiler/util"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly/api"
	"github.com/matthewmueller/golly/internal/compiler"
	"github.com/pkg/errors"
)

var (
	cli  = kingpin.New("golly", "Go to Javascript compiler")
	root = cli.Flag("root", "package root").String()

	buildCmd      = cli.Command("build", "build the packages")
	buildPackages = buildCmd.Arg("packages", "packages to build").Required().Strings()

	serveCmd      = cli.Command("serve", "development server")
	servePackages = serveCmd.Arg("packages", "packages to bundle").Required().Strings()
	servePort     = serveCmd.Flag("port", "port to serve from").Default("8080").String()

	runCmd  = cli.Command("run", "run a package")
	runFile = runCmd.Arg("file", "file to run").Required().String()

	// TODO: just have this be argv[1]
	// pkg   = cli.Flag("pkg", "package path").Required().String()
	// graph = cli.Flag("graph", "call graph").Bool()
)

func main() {
	ctx := signalContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	log.SetHandler(text.New(os.Stderr))

	command, err := cli.Parse(os.Args[1:])
	if err != nil {
		cli.FatalUsage(err.Error())
	}

	if *root == "" {
		dir, e := os.Getwd()
		if e != nil {
			log.Fatal(e.Error())
		}
		*root = dir
	}

	var e error
	switch command {
	case "build":
		e = build(ctx)
	case "serve":
		e = serve(ctx)
	case "run":
		e = run(ctx)
	}
	if e != nil {
		log.Fatal(e.Error())
	}
}

func run(ctx context.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := path.Join(cwd, *runFile)

	root, err := util.GollyPath()
	if err != nil {
		return errors.Wrapf(err, "error getting root path")
	}

	chromePath, err := chrome.Find(path.Join(root, "chrome"))
	if err != nil {
		return errors.Wrapf(err, "unable to get chrome path")
	}

	result, err := api.Run(ctx, &api.RunSettings{
		ChromePath: chromePath,
		FilePath:   filePath,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(result)
	return nil
}

func build(ctx context.Context) error {
	packages, err := getMains(*buildPackages)
	if err != nil {
		return err
	}

	// start := time.Now()
	c := compiler.New()
	files, e := c.Compile(packages...)
	if e != nil {
		return errors.Wrap(e, "error building packages")
	}

	for _, file := range files {
		// fmt.Println("---")
		// fmt.Println(file.Name())
		// fmt.Println("---")
		fmt.Println(file.Source())
		// fmt.Println("===")
	}
	return nil
}

func serve(ctx context.Context) error {
	packages, err := getMains(*servePackages)
	if err != nil {
		return err
	}

	port, e := strconv.Atoi(*servePort)
	if e != nil {
		return errors.Wrap(e, "invalid port")
	}

	return api.Serve(ctx, &api.ServeSettings{
		Packages: packages,
		Port:     port,
	})
}

func signalContext(ctx context.Context, sig ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, sig...)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
		}
	}()

	return ctx
}

// GoMainDirs returns absolute file paths to the packages that are "main"
// packages, from the list of packages given. The list of packages can
// include relative paths, the special "..." Go keyword, etc.
//
// Sourced from:
// https://github.com/mitchellh/gox/blob/c9740af9c6574448fd48eb30a71f964014c7a837/go.go#L123
func getMains(packages []string) ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	goSrc, err := util.GoSourcePath()
	if err != nil {
		return nil, err
	}

	for i, pkg := range packages {
		if filepath.Ext(pkg) != "" {
			pkg = path.Dir(pkg)
		}

		// absolute => relative
		if pkg[0] == filepath.Separator {
			rel, err := filepath.Rel(pwd, pkg)
			if err != nil {
				return nil, errors.Wrapf(err, "couldn't get relative path")
			}
			pkg = rel
		}

		// relative
		if pkg[0] == '.' && pkg[1] == filepath.Separator {
			packages[i] = pkg
			continue
		}

		pkg = "." + string(filepath.Separator) + pkg
		packages[i] = pkg
	}

	goCmd, err := exec.LookPath("go")
	if err != nil {
		return nil, err
	}

	args := make([]string, 0, len(packages)+3)
	args = append(args, "list", "-f", "{{.Name}}|{{.ImportPath}}")
	args = append(args, packages...)
	output, err := execGo(goCmd, nil, "", args...)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0, len(output))
	for _, line := range strings.Split(output, "\n") {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			log.Warnf("Bad line reading packages: %s", line)
			continue
		}

		if parts[0] == "main" {
			// TODO: not sure if this is reliable
			// but it's for when you pass a filepath
			// to go list, it returns command-line-arguments
			if parts[1] == "command-line-arguments" {
				// TODO: i don't think this will work for multiple files
				for _, pkg := range packages {
					results = append(results, pkg)
				}
				return results, nil
			}

			fullpath := strings.TrimPrefix(parts[1], "_")
			if fullpath[0] != '/' {
				fullpath = path.Join(goSrc, fullpath)
			}

			results = append(results, fullpath)
		}
	}

	return results, nil
}

func execGo(GoCmd string, env []string, dir string, args ...string) (string, error) {
	var stderr, stdout bytes.Buffer
	cmd := exec.Command(GoCmd, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if env != nil {
		cmd.Env = env
	}
	if dir != "" {
		cmd.Dir = dir
	}
	if err := cmd.Run(); err != nil {
		err = fmt.Errorf("%s\nStderr: %s", err, stderr.String())
		return "", err
	}

	return stdout.String(), nil
}
