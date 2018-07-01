package tpl

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const (
	XmlEntitlements = "entitlements.xml"
	XmlProducts     = "products.xml"
	XmlRelationship = "relationship.xml"
	XmlSession      = "session.xml"
	XmlSessionNew   = "session_new.xml"
	XmlStore        = "store.xml"
	XmlWallets      = "wallets.xml"
)

type Renderer struct {
	tpl *template.Template
}

func NewRenderer() (*Renderer, error) {
	tmpl, err := template.New("noop").Parse(``)
	if err != nil {
		return nil, err
	}

	err = parseTemplates(
		tmpl,
		bindataTemplates(),
		XmlEntitlements,
		XmlProducts,
		XmlRelationship,
		XmlSessionNew,
		XmlSession,
		XmlWallets,
		XmlStore,
	)
	if err != nil {
		return nil, err
	}

	return &Renderer{tmpl}, nil
}

func parseTemplates(tmpl *template.Template, templates map[string]string, filenames ...string) error {
	for _, filename := range filenames {
		encoded, ok := templates[filename]
		if !ok {
			return fmt.Errorf("renderer: Cannot find template '%s'", filename)
		}

		bys, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return fmt.Errorf("renderer: Cannot decode template '%s'", filename)
		}

		_, err = tmpl.New(filename).Parse(string(bys))
		if err != nil {
			return fmt.Errorf("renderer: Cannot parse template '%s'", filename)
		}
	}

	return nil
}

func (rdr *Renderer) RenderXML(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	buf := new(bytes.Buffer)
	err := rdr.tpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		logrus.WithError(err).Error("Failed render XML", name)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// logrus.Info("Response", buf.String())
	respondXML(w, r, buf.Bytes())
}

func respondXML(w http.ResponseWriter, r *http.Request, v []byte) {
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	if status, ok := r.Context().Value(render.StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	w.Write(v)
}
