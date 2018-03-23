package relationships

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Synaxis/bfheroesMagma/tpl"
)

type Controller struct {
	rdr *tpl.Renderer
}

func New(rdr *tpl.Renderer) *Controller {
	return &Controller{rdr}
}

func (s *Controller) Routing(r chi.Router) {
	// /roster/nucleus:%d
	r.Get("/roster/nucleus:{userID}", s.relationshipRosterNucleus)
	// /roster/server:%s
	r.Get("/roster/server:{serverName}", s.relationshipRosterServer)
	// /roster/server:%s/bvip/1,3
	r.Get("/roster/server:{serverName}/bvip/1,3", s.relationshipRosterServer)

	// TODO: /acknowledge/nucleus:%d/%d
	// TODO: /acknowledge/server:%s/%d
	// TODO: /increase/nucleus:%d/nucleus:%d/%s
	// TODO: /increase/server:%s/nucleus:%d/%s
	// TODO: /decrease/nucleus:%d/nucleus:%d/%s
	// TODO: /decrease/nucleus:%d/server:%s/%s
	// TODO: /decrease/server:%s/nucleus:%d/%s
	// TODO: /status/nucleus:%d
	// TODO: /status/server:%s

	// Example request payload (client, after login)
	// <update>
	// 	<status>{"O":1,"N":"heroName","T":0,"K":0,"L":0,"S":""}</status>
	// </update>
	// Probably stats like kills, deaths, suicides, team...
	r.Post("/status/nucleus:{userId}", http.NotFound)

	r.Post("/increase/nucleus:{playerID}/server:{serverName}/bvip", s.playerBookmarkServer)
}
