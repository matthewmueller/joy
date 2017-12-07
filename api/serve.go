package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/julienschmidt/httprouter"
	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/compiler/util"
	"github.com/matthewmueller/joy/internal/livereload"
	"github.com/radovskyb/watcher"
)

// ServeSettings struct
type ServeSettings struct {
	Packages    []string
	Development bool
	Port        int
}

// html includes the client JavaScript
const html = `<!doctype html>
<html>
<head>
</head>
<body>
  <script src="/bundle.js"></script>
  <script src="/livereload.js"></script>
</body>
</html>`

// Serve fn
func Serve(ctx context.Context, settings *ServeSettings) error {
	log.SetHandler(text.New(os.Stderr))

	gosrc, e := util.GoSourcePath()
	if e != nil {
		return e
	}

	// compile
	c := compiler.New(&compiler.Params{
		Development: settings.Development,
	})
	index, graph, err := c.Parse(settings.Packages...)
	if err != nil {
		return err
	}

	scripts, err := c.Assemble(index, graph)
	if err != nil {
		return err
	}

	// Create file watcher
	// TODO: use fsnotify (this is polling)
	w := watcher.New()
	w.SetMaxEvents(1)
	defer w.Close()

	// Add dir to watcher
	for _, p := range index.Paths() {
		p = path.Join(gosrc, p)
		log.Debugf("watching: %s", p)
		if e := w.Add(p); e != nil {
			return e
		}
	}

	// Create and start LiveReload server
	lr := livereload.New(settings.Port)

	// TODO: multiple package support
	bundle := scripts[0].Source()
	bundleLock := sync.RWMutex{}

	// Start goroutine that requests reload upon watcher event
	go func() {
		for {
			select {
			case <-w.Event:
				log.Debugf("file changed")
				files, err := c.Compile(settings.Packages...)
				if err != nil {
					log.WithError(err).Errorf("error compiling joy")
					continue
				}
				bundleLock.Lock()
				bundle = files[0].Source()
				bundleLock.Unlock()
				log.Debugf("reloading bundle.js")
				lr.Reload("/bundle.js")
			case err := <-w.Error:
				log.WithError(err).Errorf("error while reloading")
			case <-w.Closed:
				return
			case <-ctx.Done():
				w.Close()
				return
			}
		}
	}()

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, html)
	})

	// attach livereload handlers
	router.Handler("GET", "/livereload.js", lr)
	router.Handler("GET", "/livereload", lr)

	router.GET("/favicon.ico", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(200)
	})

	router.GET("/bundle.js", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "application/javascript")
		bundleLock.RLock()
		src := bundle
		bundleLock.RUnlock()
		fmt.Fprintf(w, src)
	})

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(settings.Port),
		Handler: router,
	}

	eg := &errgroup.Group{}
	eg.Go(func() error {
		log.Infof("listening on http://localhost:" + strconv.Itoa(settings.Port))
		return server.ListenAndServe()
	})

	// start watching
	go w.Start(50 * time.Millisecond)

	<-ctx.Done()
	server.Shutdown(context.Background())
	return nil
}
