// Using cookies, track how many times a user has been to your website domain.

package main

import (
	"io"
	"strconv"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path: "/",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	// if err != nil {
	// log.Fatalln(err)
	// }
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)




	// this also runs 
	// var cookie *http.Cookie

	// http.SetCookie(res, &http.Cookie{
	// 	Name: "my-cookie",
	// 	Value: "0",
	// 	Path: "/",
	// })

	// cookie, _ = req.Cookie("my-cookie")

	// fmt.Fprintln(res, "------------", cookie)
}

