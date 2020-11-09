package main

import (
	"fmt"
	"net/http"
)

func process(res http.ResponseWriter, req *http.Request){
	req.ParseForm()
	fmt.Fprintln(res, req.PostForm)

	fmt.Println("------------")

	fmt.Fprintln(res, "key :- ", req.PostForm.Get("hello"))
	fmt.Fprintln(res, "key :- ", req.PostForm["hello"])
	fmt.Fprintln(res, "key :- ", req.FormValue("hello"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}