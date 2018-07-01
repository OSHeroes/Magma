package nucleus

import (
	"net/http"

	"github.com/Synaxis/Magma/server/auth"
	"github.com/Synaxis/Magma/server/servertest"
	"github.com/Synaxis/Magma/tpl"
	"github.com/go-chi/chi"
)

type dtSession struct {
	Token string
}

// nucleusAuthToken authorizes the client by assigning a `magma` cookie.
func (s *Controller) nucleusAuthToken(w http.ResponseWriter, r *http.Request) {
	if serverKey := auth.GetServerHeader(r); serverKey != "" {
		s.rdr.RenderXML(w, r, tpl.XmlSession, nil)
		return
	}

	userKey, err := r.Cookie("magma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	s.rdr.RenderXML(w, r, tpl.XmlSessionNew, dtSession{userKey.Value})
}

// nucleusCheckUser is requested by both: game-server and game-client.
//
// See also TestServer_nucleusCheckUser as an example request
// made by game-client.
// func (s *Controller) nucleusCheckUser(w http.ResponseWriter, r *http.Request) {

// 	req, _ := http.NewRequest(
// 		http.MethodGet,
// 		ts.URL+`/nucleus/check/user/`+nucleusID,
// 		nil,
// 	)
// }

type dtHero struct {
	HeroID string
}

func (s *Controller) nucleusEntitlements(w http.ResponseWriter, r *http.Request) {
	s.rdr.RenderXML(w, r, tpl.XmlEntitlements, &dtHero{chi.URLParam(r, "heroID")})
}

func (s *Controller) walletsHandler(w http.ResponseWriter, r *http.Request) {
	s.rdr.RenderXML(w, r, tpl.XmlWallets, nil)
}

func (s *Controller) nucleusCheckUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam("userID")

	ts := servertest.StartTestServer()
	defer ts.Close()

	req, _ := http.NewRequest(
		http.MethodGet,
		ts.URL+`/nucleus/check/user/`+userID,
		nil,
	)
	servertest.SetCommonTestHeaders(req)
	servertest.AddTestGameClientHeaders(req)

	res, _ := ts.Client().Do(req)

	t.Log(res)

	req, _ := http.NewRequest(
		http.MethodGet,
		ts.URL+`/nucleus/check/user/`+heroID,
		nil,
	)
}

func (s *Controller) nucleusCheckUser(w http.ResponseWriter, r *http.Request) {
	// userID := chi.URLParam("userID")
}
