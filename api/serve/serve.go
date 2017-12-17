package serve

import (
	"context"
	"fmt"
	"go/build"
	"net/http"
	"path"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/apex/log"
	"github.com/julienschmidt/httprouter"
	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/livereload"
	"github.com/matthewmueller/joy/internal/mains"
	"github.com/matthewmueller/joy/internal/paths"
	"github.com/pkg/errors"
	"github.com/radovskyb/watcher"
)

// Config struct
type Config struct {
	Context     context.Context
	Packages    []string
	Development bool
	JoyPath     string
	Port        int
	Log         log.Interface // Log (optional)
}

func (c *Config) defaults() error {
	if c.Context == nil {
		c.Context = context.Background()
	}

	if c.Log == nil {
		c.Log = log.Log
	}

	if c.JoyPath == "" {
		p, err := paths.Joy()
		if err != nil {
			return errors.Wrapf(err, "error getting joy's root path")
		}
		c.JoyPath = p
	}

	return nil
}

// Serve fn
func Serve(cfg *Config) error {
	if err := cfg.defaults(); err != nil {
		return errors.Wrapf(err, "error setting defaults")
	}

	packages, err := mains.Find(cfg.Packages...)
	if err != nil {
		return errors.Wrapf(err, "error finding mains")
	}

	index, graph, err := compiler.Parse(&compiler.Config{
		Packages:    packages,
		Development: cfg.Development,
		JoyPath:     cfg.JoyPath,
	})
	if err != nil {
		return err
	}

	scripts, err := compiler.Assemble(index, graph)
	if err != nil {
		return err
	}

	// Create file watcher
	// TODO: use fsnotify (this is polling)
	w := watcher.New()
	w.SetMaxEvents(1)
	defer w.Close()

	gopath := build.Default.GOPATH

	// Add dir to watcher
	for _, p := range index.Paths() {
		p = path.Join(gopath, "src", p)
		cfg.Log.Debugf("watching: %s", p)
		if e := w.Add(p); e != nil {
			return e
		}
	}

	// Create and start LiveReload server
	lr := livereload.New(cfg.Port)

	// TODO: multiple package support
	bundle := scripts[0].Source()
	bundleLock := sync.RWMutex{}

	// Start goroutine that requests reload upon watcher event
	go func() {
		for {
			select {
			case <-w.Event:
				cfg.Log.Debugf("file changed")
				files, err := compiler.Compile(&compiler.Config{
					Packages:    cfg.Packages,
					Development: cfg.Development,
					JoyPath:     cfg.JoyPath,
				})
				if err != nil {
					cfg.Log.WithError(err).Errorf("error compiling joy")
					continue
				}
				bundleLock.Lock()
				bundle = files[0].Source()
				bundleLock.Unlock()
				cfg.Log.Debugf("reloading bundle.js")
				lr.Reload("/bundle.js")
			case err := <-w.Error:
				cfg.Log.WithError(err).Errorf("error while reloading")
			case <-w.Closed:
				return
			case <-cfg.Context.Done():
				w.Close()
				return
			}
		}
	}()

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, getHTML())
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
		Addr:    ":" + strconv.Itoa(cfg.Port),
		Handler: router,
	}

	eg := &errgroup.Group{}
	eg.Go(func() error {
		cfg.Log.Infof("listening on http://localhost:" + strconv.Itoa(cfg.Port))
		return server.ListenAndServe()
	})

	// start watching
	go w.Start(50 * time.Millisecond)

	<-cfg.Context.Done()
	server.Shutdown(context.Background())
	return nil
}

// html includes the client JavaScript
func getHTML() string {
	return `
<!doctype html>
<html>
<head>
</head>
<body>
  <script src="/bundle.js"></script>
  <script src="/livereload.js"></script>
</body>
</html>
`
}
