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
	cli = kingpin.New("golly", "golang to javascript compiler")

	// TODO: just have this be argv[1]
	pkg   = cli.Flag("pkg", "package path").Required().String()
	graph = cli.Flag("graph", "call graph").Bool()
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	kingpin.MustParse(cli.Parse(os.Args[1:]))

	if *graph {
		start := time.Now()
		golly.CallGraph(*pkg)
		log.Infof("computed callgraph of %s in %s", *pkg, time.Since(start))
		return
	}

	log.Infof("compiling... %s", *pkg)
	start := time.Now()
	src, e := golly.Compile(*pkg)
	if e != nil {
		log.WithError(e).Fatalf("error loading go package")
	}
	log.Infof("compiled %s in %s", *pkg, time.Since(start))
	fmt.Println(src)
}
