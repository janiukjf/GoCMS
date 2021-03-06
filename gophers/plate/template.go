package plate

import (
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	Layout   string
	Template string
	Bag      map[string]interface{}
	Writer   http.ResponseWriter
	FuncMap  template.FuncMap
}

/* Templating |-- Using html/template library built into golang http://golang.org/pkg/html/template/ --|
   ------------------------------ */

func (this *Server) Template(w http.ResponseWriter) (templ Template, err error) {
	if w == nil {
		log.Printf("Template Error: %v", err.Error())
		return
	}
	templ.Writer = w
	templ.Bag = make(map[string]interface{})
	return
}

func (t Template) SinglePage(file_path string) (err error) {
	if t.Bag == nil {
		t.Bag = make(map[string]interface{})
	}
	if len(file_path) != 0 {
		t.Template = file_path
	}

	tmpl, err := template.New(t.Template).Funcs(t.FuncMap).ParseFiles(t.Template)
	err = tmpl.Execute(t.Writer, t.Bag)

	return
}

func (t Template) DisplayTemplate() (err error) {
	if t.Layout == "" {
		t.Layout = "layout.html"
	}
	if t.Bag == nil {
		t.Bag = make(map[string]interface{})
	}

	templ, err := template.New(t.Layout).Funcs(t.FuncMap).ParseFiles(t.Layout, t.Template)

	err = templ.Execute(t.Writer, t.Bag)

	return
}

func (t Template) DisplayMultiple(templates []string) (err error) {
	if t.Layout == "" {
		t.Layout = "layout.html"
	}
	if t.Bag == nil {
		t.Bag = make(map[string]interface{})
	}

	templ, err := template.New(t.Layout).Funcs(t.FuncMap).ParseFiles(t.Layout)
	for _, filename := range templates {
		templ.ParseFiles(filename)
	}
	err = templ.Execute(t.Writer, t.Bag)

	return
}
