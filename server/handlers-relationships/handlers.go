package relationships

import (
	"net/http"

	"github.com/Synaxis/bfheroesMagma/tpl"

	"github.com/go-chi/chi"
)

type dtRelationship struct {
	ServerID string
}

// relationshipHandler is used by game-server
//
// FIXME: This request is retried 5 times in 10sec interval.
//
// See also TestServer_relationshipRosterServer as an example request.
func (s *Controller) relationshipRosterServer(w http.ResponseWriter, r *http.Request) {
	s.rdr.RenderXML(w, r, tpl.XmlRelationship, &dtRelationship{
		ServerID: chi.URLParam(r, "serverName"),
	})
}

// relationshipRosterNucleus is used by game-client.
//
// FIXME: This request is retried 5 times in 10sec interval.
//
// See also TestServer_relationshipRosterNucleus as an example request.
func (s *Controller) relationshipRosterNucleus(w http.ResponseWriter, r *http.Request) {
	s.rdr.RenderXML(w, r, tpl.XmlRelationship, &dtRelationship{
		ServerID: chi.URLParam(r, "userID"),
	})
}

func (s *Controller) playerBookmarkServer(w http.ResponseWriter, r *http.Request) {
	// playerID := chi.URLParam(r, "playerID")
	// serverName := chi.URLParam(r, "serverName")

	// POST /relationships/increase/nucleus:2/server:Server/bvip HTTP/1.1
	// Host: opnhr.ddns.net
	// Cache-Control: no-cache
	// Connection: Keep-Alive
	// Content-Length: 0
	// Cookie: magma=topsecret
	// Pragma: no-cache
	// User-Agent: BFHeroesINet

	w.WriteHeader(http.StatusOK)
}
