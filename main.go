package main

import (
	"log"
	"net/http"
	"text/template"
)

type myHandler int

var tpl *template.Template

var ansFunc = template.FuncMap{
	"a1": answer1,
}

func init() {
	tpl = template.Must(template.New("").Funcs(ansFunc).ParseFiles("index.gohtml"))
}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.PostForm)
}

func main() {
	var myVar myHandler
	http.ListenAndServe(":8080", myVar)
}

func answer1(s string) string {
	var reply = ""
	if s == "root-node" {
		reply = `Correct`
	} else {
		reply = `Incorrect`
	}
	return reply
}
