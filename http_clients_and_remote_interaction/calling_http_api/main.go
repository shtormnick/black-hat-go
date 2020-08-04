package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main() {

	rl, _ := http.Get("http://www.goolge.com/robots.txt")
	defer rl.Body.Close()

	r2, _ := http.Head("http://www.goolge.com/robots.txt")
	defer r2.Body.Close()

	form := url.Values{}
	form.Add("foo", "bar")

	r3, _ := http.Post(
		"http://www.goolge.com/robots.txt", 
		"aplication/x-www-form-urlencoded", 
		strings.NewReader(form.Encode()),
	)
	defer r3.Body.Close()
}