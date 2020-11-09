package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func jsonExample(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Kunal Prakash Lucky",
		Threads: []string{"first", "second", "third"},
	}

	json, _ := json.Marshal(post)
	res.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	

	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}