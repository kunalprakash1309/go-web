package main

import (
	"fmt"
	"net/http"
)

func writeExample(res http.ResponseWriter, req *http.Request){
	res.Header().Set("my name", "lucky")
	res.Header().Set("Content-type", "text/html; charset=utf-8")
	
	req.Header.Set("my name", "Kunal")
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`

	res.Write([]byte(str))
	fmt.Fprintln(res, str)
	
}

func headerExample(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Location", "https://www.google.com")
	res.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/redirect", headerExample)

	server.ListenAndServe()
}