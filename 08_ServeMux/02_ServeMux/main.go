package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (d hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	m := http.NewServeMux()

	m.Handle("/dog/", d)
	m.Handle("/cat", c)

	http.ListenAndServe(":8080", m)
}