package main

import (
	"os"
	"path/filepath"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly"
)

var (
	cli = kingpin.New("golly", "golang to javascript compiler")

	pkg = cli.Flag("pkg", "package path").String()
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	kingpin.MustParse(cli.Parse(os.Args[1:]))

	// support relative and absolute paths
	// TODO: fix crappy code
	pkgpath := *pkg
	if !filepath.IsAbs(pkgpath) {
		cwd, e := os.Getwd()
		if e != nil {
			log.WithError(e).Fatalf("error getting current working directory")
		}
		pkgpath = filepath.Join(cwd, pkgpath)
	}
	pkgpath, _ = filepath.Rel(os.Getenv("GOPATH")+"/src", pkgpath)

	log.Infof("compiling... %s", pkgpath)
	start := time.Now()
	src, e := golly.Compile(pkgpath, "")
	if e != nil {
		log.WithError(e).Fatalf("error loading go package")
	}
	log.Infof("src %s", src)
	log.Infof("compiled %s in %s", pkgpath, time.Since(start))
}
