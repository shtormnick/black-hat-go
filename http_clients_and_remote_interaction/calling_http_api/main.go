package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	rl, err := http.Get("https://www.goolge.com/robots.txt")
	defer rl.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	r2, err := http.Head("https://www.goolge.com/robots.txt")
	defer r2.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	form := url.Values{}
	form.Add("foo", "bar")

	r3, err := http.NewRequest(
		"PUT",
		"https://www.goolge.com/robots.txt", 
		strings.NewReader(form.Encode()),
	)

	if err != nil {
		log.Fatalln(err)
	}

	defer r3.Body.Close()

	req, err := http.NewRequest("DELETE", "http://www.goolge.com/robots.txt", nil)

	var client http.Client

	resp, err := client.Do(req)

	fmt.Println(resp)
}