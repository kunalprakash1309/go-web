package main

import (
	"io"
	"log"
	"net/http"
	"html/template"
)

func foo(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request){
	var tpl *template.Template

	tpl = template.Must(tpl.ParseFiles("dog.gohtml"))

	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func image(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "dog.jpg")
}

func main(){

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", image)
	log.Fatalln(http.ListenAndServe(":8080", nil))
	
}
