package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly/api"
	"github.com/matthewmueller/golly/golang"
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
	ctx := withContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
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
	if err := api.Run(ctx, *runFile); err != nil {
		return errors.Wrap(err, "error running file")
	}
	return nil
}

func build(ctx context.Context) error {
	var packages []string
	for _, pkg := range *buildPackages {
		packages = append(packages, path.Join(*root, pkg))
	}

	// start := time.Now()
	compiler := golang.New()
	files, _, e := compiler.Compile(packages...)
	if e != nil {
		return errors.Wrap(e, "error building packages")
	}

	for _, file := range files {
		fmt.Println("---")
		fmt.Println(file.Name)
		fmt.Println("---")
		fmt.Println(file.Source)
		fmt.Println("===")
	}
	return nil
}

func serve(ctx context.Context) error {
	port, e := strconv.Atoi(*servePort)
	if e != nil {
		return errors.Wrap(e, "invalid port")
	}

	var packages []string
	for _, pkg := range *servePackages {
		packages = append(packages, path.Join(*root, pkg))
	}

	return api.Serve(&api.ServeParams{
		Packages: packages,
		Port:     port,
	})
}

func withContext(ctx context.Context, sig ...os.Signal) context.Context {
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
