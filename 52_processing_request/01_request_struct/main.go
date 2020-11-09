package main

import (
	"fmt"
	"net/http"
)

func headersStruct(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "--------start--------" )
	fmt.Fprintln(res, "Method -------", req.Method )
	fmt.Fprintln(res, "URL ----------", req.URL )
	fmt.Fprintln(res, "Header -------", req.Header )

	// fmt.Fprintln(res, "User Agent", req.Header["User-Agent"])

	fmt.Fprintln(res, "User-Agent -----", req.Header.Get("User-Agent"))
	fmt.Fprintln(res, "Body ---------", req.Body)
	fmt.Fprintln(res, "ContentLength ---", req.ContentLength )
	fmt.Fprintln(res, "Host", req.Host)
	fmt.Fprintln(res, "--------end-------" )
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", headersStruct)
	server.ListenAndServe()
}