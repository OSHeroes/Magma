package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	dc "gitlab.com/go-heroes/heroes-api/server/handlers-dc"
	game "gitlab.com/go-heroes/heroes-api/server/handlers-game"
	nucleus "gitlab.com/go-heroes/heroes-api/server/handlers-nucleus"
	ofb "gitlab.com/go-heroes/heroes-api/server/handlers-ofb"
	relationships "gitlab.com/go-heroes/heroes-api/server/handlers-relationships"
)

func (s *Server) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(s.logRequestMiddleware)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.Timeout(60 * time.Second))
	// r.Mount("/debug", middleware.Profiler())

	// TODO: user/updateUserProfile/%d

	// Nuclues (authentication and account data)
	r.Route("/relationships", relationships.New(s.rdr).Routing)

	// Client-relationship (i.e. friends, server bookmarks)
	r.Route("/nucleus", nucleus.New(s.rdr).Routing)

	// In-game overlay
	r.Route("/ofb", ofb.New(s.rdr).Routing)

	// Data collection
	r.Route("/dc", dc.New(s.rdr).Routing)

	// Game-client
	r.Route("/en/game", game.New(s.rdr).Routing)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		logrus.Warn("Not found URL: ", r.URL.String())
		logRequest(r)
		http.NotFound(w, r)
	})

	return r
}
