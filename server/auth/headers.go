package auth

import (
	"net/http"
)

const (
	xServerKey = "X-SERVER-KEY"
)

func GetServerHeader(r *http.Request) string {
	return r.Header.Get(xServerKey)
}
