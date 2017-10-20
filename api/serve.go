package api

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path"
// 	"sync"
// 	"time"

// 	"github.com/apex/log"
// 	"github.com/apex/log/handlers/text"
// 	"github.com/jaschaephraim/lrserver"
// 	"github.com/julienschmidt/httprouter"
// 	"github.com/matthewmueller/golly/golang"
// 	"github.com/radovskyb/watcher"
// )

// // ServeSettings struct
// type ServeSettings struct {
// 	Packages []string
// 	Port     int
// }

// // html includes the client JavaScript
// const html = `<!doctype html>
// <html>
// <head>
// </head>
// <body>
//   <script src="/bundle.js"></script>
//   <script src="http://localhost:35729/livereload.js"></script>
// </body>
// </html>`

// // Serve fn
// func Serve(ctx context.Context, settings *ServeSettings) error {
// 	log.SetHandler(text.New(os.Stderr))

// 	gosrc := path.Join(os.Getenv("GOPATH"), "src")

// 	// build
// 	compiler := golang.New()
// 	files, err := compiler.Compile(settings.Packages...)
// 	if err != nil {
// 		return err
// 	}
// 	_ = files

// 	// Create file watcher
// 	// TODO: use fsnotify (this is polling)
// 	w := watcher.New()
// 	w.SetMaxEvents(1)
// 	defer w.Close()

// 	// Add dir to watcher
// 	for _, main := range graph {
// 		// add packages
// 		for _, pkg := range main.Packages {
// 			pkgpath := path.Join(gosrc, pkg.Path)
// 			if e := w.Add(pkgpath); e != nil {
// 				return e
// 			}
// 		}

// 		// add raw files
// 		for _, file := range main.RawFiles {
// 			pkgpath := path.Join(gosrc, file.Name)
// 			if e := w.Add(pkgpath); e != nil {
// 				return e
// 			}
// 		}
// 	}

// 	// Create and start LiveReload server
// 	lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
// 	go lr.ListenAndServe()

// 	// TODO: multiple package support
// 	bundle := files[0].Source
// 	bundleLock := sync.RWMutex{}

// 	// Start goroutine that requests reload upon watcher event
// 	go func() {
// 		for {
// 			select {
// 			case <-w.Event:
// 				files, err := compiler.Compile(settings.Packages...)
// 				if err != nil {
// 					log.WithError(err).Errorf("error compiling golly")
// 					continue
// 				}
// 				bundleLock.Lock()
// 				bundle = files[0].Source()
// 				bundleLock.Unlock()
// 				lr.Reload(bundle)
// 			case err := <-w.Error:
// 				log.WithError(err).Errorf("error while reloading")
// 			case <-w.Closed:
// 				return
// 			}
// 		}
// 	}()

// 	go w.Start(100 * time.Millisecond)

// 	router := httprouter.New()
// 	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 		fmt.Fprint(w, html)
// 	})

// 	router.GET("/bundle.js", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 		w.Header().Add("Content-Type", "application/javascript")
// 		bundleLock.RLock()
// 		src := bundle
// 		bundleLock.RUnlock()
// 		fmt.Fprintf(w, src)
// 	})

// 	log.Infof("listening on http://localhost:8080")
// 	return http.ListenAndServe(":8080", router)
// }
