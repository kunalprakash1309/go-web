package main

import (
	"fmt"
	"net/http"
)

type hotdog string

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request){
	switch req.URL.Path {
	case "/dog":
		fmt.Fprintln(res, "doggy doggy doggy")
	
	case "/cat":
		fmt.Fprintln(res, "kitty kitty kitty")
	}
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080", d)
}