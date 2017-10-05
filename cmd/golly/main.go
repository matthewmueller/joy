package main

import (
	"fmt"
	"os"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly"
)

var (
	cli = kingpin.New("golly", "Go to Javascript compiler")

	build    = cli.Command("build", "build the packages")
	packages = build.Arg("packages", "packages to build").Required().Strings()

	serve = cli.Command("serve", "development server")

	// TODO: just have this be argv[1]
	// pkg   = cli.Flag("pkg", "package path").Required().String()
	// graph = cli.Flag("graph", "call graph").Bool()
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	// kingpin.MustParse(cli.Parse(os.Args[1:]))

	command, err := cli.Parse(os.Args[1:])
	if err != nil {
		cli.FatalUsage(err.Error())
	}

	switch command {
	case "build":
		builder()
	case "serve":
		server()
	}
}

func builder() {
	log.Infof("compiling...")
	start := time.Now()
	files, e := golly.Compile(*packages...)
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

func server() {

}
