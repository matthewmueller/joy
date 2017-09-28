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
	pkg = cli.Flag("pkg", "package path").Required().String()
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	kingpin.MustParse(cli.Parse(os.Args[1:]))

	log.Infof("compiling... %s", pkg)
	start := time.Now()
	src, e := golly.Compile(*pkg)
	if e != nil {
		log.WithError(e).Fatalf("error loading go package")
	}
	log.Infof("compiled %s in %s", pkg, time.Since(start))
	fmt.Println(src)
}
