package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Submissions url.Values
		Header http.Header
		Host string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		// before calling r.Form we need to call r.ParseForm()
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}