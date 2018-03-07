package dc

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/go-heroes/heroes-api/tpl"
)

type Controller struct {
	rdr *tpl.Renderer
}

func New(rdr *tpl.Renderer) *Controller {
	return &Controller{rdr}
}

func (s *Controller) Routing(r chi.Router) {
	// Example request payload (from server):
	// <dataset>
	// 	<insert table="error">
	// 		<code>2</code>
	// 		<context>backend</context>
	// 		<session>0</session>
	// 		<stack>AutoLogin Timeout on login attempt #1</stack>
	// 		<ts>2018-02-03 00:00:45</ts>
	// 	</insert>
	// 	<insert table="error">
	// 		<code>2</code>
	// 		<context>backend</context>
	// 		<session>0</session>
	// 		<stack>AutoLogin Timeout on login attempt #2</stack>
	// 		<ts>2018-02-03 00:00:45</ts>
	// 	</insert>
	// 	<insert table="error">
	// 		<code>2</code>
	// 		<context>backend</context>
	// 		<session>0</session>
	// 		<stack>AutoLogin Timeout on login attempt #3</stack>
	// 		<ts>2018-02-03 00:00:45</ts>
	// 	</insert>
	// </dataset>
	r.Post("/submit", http.NotFound)
}
