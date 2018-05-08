package game

import (
	"net/http"

	"github.com/Synaxis/Magma/tpl"
)

func (c *Controller) store(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlStore, nil)
}
