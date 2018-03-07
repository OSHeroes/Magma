package ofb

import (
	"github.com/go-chi/chi"

	"gitlab.com/go-heroes/heroes-api/tpl"
)

type Controller struct {
	rdr *tpl.Renderer
}

func New(rdr *tpl.Renderer) *Controller {
	return &Controller{rdr}
}

func (c *Controller) Routing(r chi.Router) {
	r.Get("/products", c.ofbProducts)
	// TODO: purchase/%d/%s
}
