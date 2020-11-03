package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func dog(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "dog dog dog")
}

func me(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "hello %s", "kunal")
}

func defualt(res http.ResponseWriter, req *http.Request){


	err := tpl.ExecuteTemplate(res, "tpl.gohtml", "Kunal")
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	// http.HandleFunc("/", defualt)
	// http.HandleFunc("/dog/", dog)
	// http.HandleFunc("/me/", me)

	http.Handle("/", http.HandlerFunc(defualt))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}