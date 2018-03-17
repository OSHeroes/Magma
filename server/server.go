package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/go-heroes/heroes-api/tpl"
)

var (
	logger = logrus.WithField("pkg", "server")
)

type Server struct {
	rdr *tpl.Renderer
}

func New() (*Server, error) {
	rdr, err := tpl.NewRenderer()
	if err != nil {
		return nil, err
	}

	return &Server{rdr}, nil
}

func (s *Server) ListenAndServe(bind, bindSecure, certPath, pemPath string) {
	r := s.registerRoutes()

	// TODO: Pass context.Context to gracefully shutdown servers
	go s.ServeHTTP(bind, r)
	go s.ServeTLS(bindSecure, certPath, pemPath, r)

}

func (s *Server) ServeHTTP(bind string, r http.Handler) {
	srv := http.Server{
		Addr:         bind,
		Handler:      r,
		ErrorLog:     log.New(os.Stderr, "HTTP: ", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second, //todo check timeouts
		IdleTimeout:  60 * time.Second,
	}

	logrus.Info("Starting listening on HTTP")
	go logrus.Error(srv.ListenAndServe())
}

func (s *Server) ServeTLS(bind string, certPath, pemPath string, r http.Handler) {
	srv := &http.Server{
		Addr:         bind,
		Handler:      r,
		ErrorLog:     log.New(os.Stderr, "TLS: ", 0),
		ReadTimeout:  5 * time.Second, //todo check timeouts
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logrus.Info("Starting listening on HTTPS")
	logrus.Error(srv.ListenAndServeTLS(certPath, pemPath))
}
