// chaining is useful if we need same work to be done with multiple function
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "Hello!")
}


// it also a middleware
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request){

		// it prints the name of function which is passed (in this case h(main.hello))
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		fmt.Println("Kunal:- ", h)

		// dont know yet, but use this
		h(res, req)
	}
}


func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}