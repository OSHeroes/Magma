package game

import (
	"net/http"

	"gitlab.com/go-heroes/heroes-api/tpl"
)

func (c *Controller) store(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlStore, nil)
}
