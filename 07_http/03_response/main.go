package main

import (
	"fmt"
	"net/http"
)

type hotdog int 

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Kunal-Key", "this is from Kunal")
	w.Header().Set("Conten-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any Code goes here</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}