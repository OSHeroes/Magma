package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	var (
		tplDirectory = flag.String("dir", "./tpl/xmls/", "Path to the server/tpl/ directory")
		outFilepath  = flag.String("out", "./tpl/templates.go", "Path where server should be saved")
	)
	flag.Parse()

	rf := rendererFile{
		Contents: map[string]string{},
	}
	rf.readContents(*tplDirectory)

	render(&rf, *outFilepath)
}

type rendererFile struct {
	Contents map[string]string
}

func (rf *rendererFile) readContents(dirpath string) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			rf.encode(dirpath, f.Name())
		}
	}
}

func (rf *rendererFile) encode(dirpath, filepath string) {
	f, err := os.Open(dirpath + filepath)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	if err != nil {
		log.Fatal(err, filepath)
	}
	defer f.Close()

	rf.Contents[filepath] = base64.StdEncoding.EncodeToString(buf.Bytes())
}

func render(rf *rendererFile, filepath string) {
	t := template.Must(template.New("out").Parse(`package tpl

func bindataTemplates() map[string]string {
	return map[string]string{
		// Base64-encoded templates
{{ range $name, $contents := .Contents }}
		"{{ $name }}": "{{ $contents }}",{{ end }}
	}
}
`))
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(f, rf); err != nil {
		log.Fatal(err)
	}
}
