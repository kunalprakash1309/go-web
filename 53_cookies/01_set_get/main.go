package main

import (
	"fmt"
	"net/http"
)

func setCookie(res http.ResponseWriter, req *http.Request){
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go web ",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Manning Publication",
		HttpOnly: true,
	}

	http.SetCookie(res, &c1)
	http.SetCookie(res, &c2)
}

func getCookie(res http.ResponseWriter, req *http.Request){
	// one of the method
	// h := req.Header["Cookie"]
	// j := req.Header.Get("Cookie")

	c1, err := req.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(res, "cannot get cookie", err)
	}

	cs := req.Cookies()
	fmt.Fprintln(res, c1)
	fmt.Fprintln(res, cs)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/get", getCookie)
	server.ListenAndServe()
}