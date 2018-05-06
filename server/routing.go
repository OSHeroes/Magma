package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	dc "github.com/Synaxis/bfheroesMagma/server/handlers-dc"
	game "github.com/Synaxis/bfheroesMagma/server/handlers-game"
	nucleus "github.com/Synaxis/bfheroesMagma/server/handlers-nucleus"
	ofb "github.com/Synaxis/bfheroesMagma/server/handlers-ofb"
	relationships "github.com/Synaxis/bfheroesMagma/server/handlers-relationships"
)

func (s *Server) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(s.logRequestMiddleware)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
//  'nucleus/entitlements/%I64d',0
//  'ofb/products',0
//  'nucleus/wallets/%I64d',0
//  'ofb/purchase/%I64d/%s',0
//  'nucleus/wallets/%I64d/%s/%d/%s',0
//  'nucleus/refundAbilities/%I64d',0
//  'relationships/acknowledge/nucleus:%I64d/%I64d',0
//  'relationships/acknowledge/server:%s/%I64d',0
//  'relationships/increase/nucleus:%I64d/nucleus:%I64d/%s',0
//  'relationships/increase/nucleus:%I64d/server:%s/%s',0
//  'relationships/increase/server:%s/nucleus:%I64d/%s',0
//  'relationships/decrease/nucleus:%I64d/nucleus:%I64d/%s',0
//  'relationships/decrease/nucleus:%I64d/server:%s/%s',0
//  'relationships/decrease/server:%s/nucleus:%I64d/%s',0
//  'relationships/status/nucleus:%I64d',0
//  '<update>',0Ah,0
//  '</status>',0Ah,0
//  09h,'<status>',0
//  '</update>',0
//  'relationships/status/server:%s',0
//  '<update>',0Ah,0
//  '</status>',0Ah,0
//  09h,'<status>',0
//  '</name>',0Ah,0
//  09h,'<name>',0
//  '</update>',0
//  'nucleus/personas/%s',0
//  'relationships/roster/nucleus:%I64d',0
//  'relationships/roster/server:%s/bvip/1,3',0
//  'relationships/roster/nucleus:%I64d',0
//  'relationships/roster/server:%s',0
// 'user/updateUserProfile/%I64d',0
//  'nucleus/entitlement/%s/useCount/%d',0
// : 'nucleus/entitlement/%s/status/%s',0
// 'user',0
//  'persona',0
//  'nucleus/check/%s/%I64d',0
//  'nucleus/authToken',0
//  'nucleus/name/%I64d',0
//  'dc/submit',0
//  'nucleus/entitlements/%I64d',0
//  'nucleus/entitlements/%I64d?entitlementTag=%s',0

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
		logrus.Warn("Not found URL: ", r.URL.String())
		logRequest(r)
		http.NotFound(w, r)
	})

	return r
}
