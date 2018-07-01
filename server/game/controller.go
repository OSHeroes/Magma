package game

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/Synaxis/Magma/tpl"
)

type Controller struct {
	rdr *tpl.Renderer
}

func New(rdr *tpl.Renderer) *Controller {
	return &Controller{rdr}
}

func (c *Controller) Routing(r chi.Router) {
	r.Get("/store", c.store)
}


func (c *Controller) store(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlStore, nil)
}
