package main

import (
	"fmt"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly"
)

// var (
// 	cli = kingpin.New("golly", "golang to javascript compiler")

// 	// TODO: just have this be argv[1]
// 	pkg   = cli.Flag("pkg", "package path").Required().String()
// 	graph = cli.Flag("graph", "call graph").Bool()
// )

func main() {
	log.SetHandler(text.New(os.Stderr))
	// kingpin.MustParse(cli.Parse(os.Args[1:]))

	// if *graph {
	// 	start := time.Now()
	// 	golly.CallGraph(*pkg)
	// 	log.Infof("computed callgraph of %s in %s", *pkg, time.Since(start))
	// 	return
	// }

	// log.Infof("args %+v", os.Args[1:])

	log.Infof("compiling...")
	start := time.Now()
	files, e := golly.Compile(os.Args[1:]...)
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
