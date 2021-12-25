package view

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var (
	tpIndex = template.New("")
)

func init() {
	tpIndex.Funcs(template.FuncMap{})
	_, err := tpIndex.ParseFiles("templates/root.tmpl", "templates/index.tmpl")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tpIndex = tpIndex.Lookup("root")
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

//index renders index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}
