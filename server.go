package contracts

import (
	"net/http"
	"strings"
	"time"
)

// Server represents the HTTP file server
type Server struct {
	S    *http.Server
	path string
}

// NewServer returns a configured server
func NewServer(port, path string) *Server {
	m := http.NewServeMux()
	m.Handle("/contracts/",
		http.StripPrefix("/contracts", noDirectoryListing(http.FileServer(http.Dir(path)))))

	s := &http.Server{
		Addr:         "127.0.0.1:" + port,
		Handler:      m,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return &Server{
		S:    s,
		path: path,
	}
}

func noDirectoryListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
