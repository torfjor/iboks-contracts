package contracts

import (
	"net/http"
	"time"
)

type Server struct {
	S    *http.Server
	path string
}

func NewServer(port, path string) *Server {
	m := http.NewServeMux()
	m.Handle("/contracts/",
		http.StripPrefix("/contracts/", http.FileServer(http.Dir(path))))

	s := &http.Server{
		Addr:         "0.0.0.0:" + port,
		Handler:      m,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return &Server{
		S:    s,
		path: path,
	}
}
