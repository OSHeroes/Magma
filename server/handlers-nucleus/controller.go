package nucleus

import (
	"github.com/go-chi/chi"

	"github.com/Synaxis/Magma/tpl"
)

type Controller struct {
	rdr *tpl.Renderer
}

func New(rdr *tpl.Renderer) *Controller {
	return &Controller{rdr}
}

func (s *Controller) Routing(r chi.Router) {
	// /authToken
	r.Get("/authToken", s.nucleusAuthToken)

	// /entitlements/%d
	// /entitlements/%d?entitlementTag=%s
	r.Get("/entitlements/{heroID}", s.nucleusEntitlements)
	// TODO: /entitlement/%s/useCount/%d
	// TODO: /entitlement/%s/status/%s

	// TODO: /wallets/%d/%s/%d/%s
	// TODO: /wallets/%d
	// GET /nucleus/wallets/5 HTTP/1.1
	// Host: 127.0.0.1
	// Cache-Control: no-cache
	// Connection: Keep-Alive
	// Cookie: magma=topsecret
	// Pragma: no-cache
	// User-Agent: BFHeroesINet
	r.Get("/wallets/{heroID}", s.walletsHandler)

	// TODO: /check/%s/%d
	r.Get("/check/user/{userID}", s.nucleusCheckUser)

	// TODO: /personas/%s
	// TODO: /refundAbilities/%d
	// TODO: /personas/%s
	// TODO: /name/%d
}
