package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", foo)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(res http.ResponseWriter, req *http.Request){
	fname := req.FormValue("f")
	lname := req.FormValue("l")
	fmt.Fprintln(res, "hello: "  +fname +lname)
}