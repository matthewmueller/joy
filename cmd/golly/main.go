package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly"
	"github.com/matthewmueller/golly/api"
)

var (
	cli = kingpin.New("golly", "Go to Javascript compiler")

	buildCmd      = cli.Command("build", "build the packages")
	buildPackages = buildCmd.Arg("packages", "packages to build").Required().Strings()

	serveCmd = cli.Command("serve", "development server")

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

	switch command {
	case "build":
		build()
	case "serve":
		serve()
	case "run":
		run(ctx, *runFile)
	}
}

func run(ctx context.Context, file string) {
	if err := api.Run(ctx, *runFile); err != nil {
		log.WithError(err).Fatal("error running file")
	}
}

func build() {
	log.Infof("compiling...")
	start := time.Now()
	files, e := golly.Compile(*buildPackages...)
	if e != nil {
		log.WithError(e).Fatalf("error compiling go package")
	}
	log.Infof("compiled in %s", time.Since(start))
	for _, file := range files {
		fmt.Println("---")
		fmt.Println(file.Name)
		fmt.Println("---")
		fmt.Println(file.Source)
		fmt.Println("===")
	}
}

func serve() {

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
