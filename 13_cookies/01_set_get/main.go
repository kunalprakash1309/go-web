package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func set(res http.ResponseWriter, req *http.Request){
	http.SetCookie(res, &http.Cookie{
		Name: "my-cookie",
		Value: "some value",
		Path: "/",
	})

	// For multiple cookies
	// http.SetCookie(res, &http.Cookie{
	// 	Name: "my-cookie",
	// 	Value: "some value",
	// })
}

func read(res http.ResponseWriter, req *http.Request){

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, "Your Cookie: ", c)
}