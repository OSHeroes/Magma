package tpl

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const (
	Entitlements = "entitlements.xml"
	Products     = "products.xml"
	Relationship = "relationship.xml"
	Session      = "session.xml"
	SessionNew   = "session_new.xml"
	Store        = "store.xml"
	Wallets      = "wallets.xml"
	check        = "check.xml"
)

type Renderer struct {
	tpl *template.Template
}

func newRenderer() Renderer {
	tpl := template.Must(
		template.ParseFiles(
			addPathPrefix(
				Entitlements,
				Products,
				Relationship,
				SessionNew,
				Session,
				Wallets,
				Store,
			)...,
		),
	)
	return Renderer{tpl}
}

func addPathPrefix(path ...string) []string {
	prefixed := []string{}
	for _, p := range path {
		prefixed = append(prefixed, fmt.Sprintf("server/tpl/%s", p))
	}
	return prefixed
}

func (rdr *Renderer) RenderXML(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	buf := new(bytes.Buffer)
	err := rdr.tpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		logrus.WithError(err).Error("Failed render XML", name)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.Info("Response\n", buf.String())
	respondXML(w, r, buf.Bytes())
}

func respondXML(w http.ResponseWriter, r *http.Request, v []byte) {
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	if status, ok := r.Context().Value(render.StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	logrus.Println("->Answer : ")
	w.Write(v)
}
