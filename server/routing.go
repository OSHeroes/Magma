package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	dc "github.com/Synaxis/Magma/server/handlers-dc"
	game "github.com/Synaxis/Magma/server/handlers-game"
	nucleus "github.com/Synaxis/Magma/server/handlers-nucleus"
	ofb "github.com/Synaxis/Magma/server/handlers-ofb"
	relationships "github.com/Synaxis/Magma/server/handlers-relationships"
)

func (s *Server) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(s.logRequestMiddleware)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
//  'nucleus/wallets/%I64d
//  'ofb/purchase/%I64d/%s
//  'nucleus/wallets/%I64d/%s/%d/%s
//  'nucleus/refundAbilities/%I64d
//  'relationships/acknowledge/nucleus:%I64d/%I64d
//  'relationships/acknowledge/server:%s/%I64
//  'relationships/increase/nucleus:%I64d/nucleus:%I64d/%s
//  'relationships/increase/nucleus:%I64d/server:%s/%s
//  'relationships/increase/server:%s/nucleus:%I64d/%s
//  'relationships/decrease/nucleus:%I64d/nucleus:%I64d/%s
//  'relationships/decrease/nucleus:%I64d/server:%s/%s
//  'relationships/decrease/server:%s/nucleus:%I64d/%s
//  'relationships/status/nucleus:%I64d'
//  'relationships/status/server:%s
//  'nucleus/personas/%s
//  'relationships/roster/nucleus:%I64d
//  'relationships/roster/server:%s/bvip/1,3
//  'relationships/roster/nucleus:%I64d
//  'relationships/roster/server:%s
//  'nucleus/entitlement/%s/useCount/%d
//  'nucleus/entitlement/%s/status/%s'
//  'persona
//  'nucleus/check/%s/%I64d
//  'nucleus/authToken
//  'nucleus/name/%I64d'
//  'nucleus/entitlements/%I64d'
//  'nucleus/entitlements/%I64d?entitlementTag=%s'

	// TODO: user/updateUserProfile/%d
	// r.Use(middleware.Timeout(60 * time.Second))
	// r.Mount("/debug", middleware.Profiler())


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
		logrus.Warn("New Request , check  Dump: ", r.URL.String())
		logRequest(r)
		http.NotFound(w, r)
	})

	return r
}
