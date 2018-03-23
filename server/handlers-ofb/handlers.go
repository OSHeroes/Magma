package ofb

import (
	"net/http"

	"github.com/Synaxis/bfheroesMagma/tpl"
)

func (c *Controller) ofbProducts(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlProducts, nil)
}
