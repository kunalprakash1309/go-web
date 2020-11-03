package main

import (
	"io"
	"fmt"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "dog dog dog")
}

func me(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "hello %s", "kunal")
}

func defualt(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "home page")
}

func main(){
	http.HandleFunc("/", defualt)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
