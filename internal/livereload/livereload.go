package livereload

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var _ http.Handler = (*Server)(nil)

// New livereload server
func New(port int) *Server {
	srv := &Server{}

	router := httprouter.New()
	srv.router = router
	router.HandlerFunc("GET", "/livereload.js", srv.serveJS)
	router.HandlerFunc("GET", "/livereload", srv.serveWS)

	srv.mu = &sync.RWMutex{}
	srv.conns = map[*conn]struct{}{}
	srv.js = fmt.Sprintf(js, port)

	return srv

}

// Server struct
type Server struct {
	router *httprouter.Router
	js     string

	// protects the connections
	mu    *sync.RWMutex
	conns map[*conn]struct{}
}

// Reload a file
func (s *Server) Reload(file string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for conn := range s.conns {
		conn.Reload(file)
	}

	return nil
}

// ServeHTTP implements the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// serveJS serves the client
func (s *Server) serveJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(s.js))
}

// serveWS serves the websocket
func (s *Server) serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// add the connection and defer it's removal
	s.mu.Lock()
	s.conns[conn] = struct{}{}
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		delete(s.conns, conn)
		s.mu.Unlock()
	}()

	// wait until the connection errors out
	if err := conn.Wait(); err != nil {
		return
	}
}
