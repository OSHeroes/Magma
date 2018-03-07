package servertest

import (
	"net/http"
	"net/http/httptest"
)

var (
	// It is ID of the player (Nucleus ID).
	nucleusID = "2"

	// serverName is the name of the server.
	serverName = "Test-Server"
)

func StartTestServer(h http.Handler) *httptest.Server {
	// s, _ := server.New()
	// h := s.registerRoutes()
	ts := httptest.NewServer(h)
	return ts
}

func SetCommonTestHeaders(req *http.Request) {
	req.Header.Set("Host", "example.com")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("User-Agent", "BFHeroesINet")
}

func AddTestGameServerHeaders(req *http.Request) {
	// serverKey is value of the `+key` argument,
	// which is provided when starting the game-server.
	serverKey := "some-server-identifier"

	req.Header.Set("X-Server-Key", serverKey)
}

func AddTestGameClientHeaders(req *http.Request) {
	// magma is a cookie assgined in GET /nucleus/authToken
	magma := "some-user-cookie"

	req.Header.Set("Cookie", `magma=`+magma)
}
