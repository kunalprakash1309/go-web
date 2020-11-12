package main

import (
	//"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(res http.ResponseWriter, req *http.Request) {
	msg := "hello world"
	c := http.Cookie {
		Name: "flash",
		Value: msg,
	}

	http.SetCookie(res, &c)
}

func showMessage(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(res, "No message found")
		}
	} else {
		rc := http.Cookie {
			Name: "flash",
			MaxAge: -1,
			Expires : time.Unix(1, 0),
		}
		http.SetCookie(res, &rc)
		value := c.Value
		fmt.Fprintln(res, value)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set", setMessage)
	http.HandleFunc("/show", showMessage)
	server.ListenAndServe()
}