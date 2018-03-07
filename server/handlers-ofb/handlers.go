package ofb

import (
	"net/http"

	"gitlab.com/go-heroes/heroes-api/tpl"
)

func (c *Controller) ofbProducts(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlProducts, nil)
}
