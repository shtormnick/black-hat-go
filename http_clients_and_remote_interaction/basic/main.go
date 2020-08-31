package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	resp, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(body))
	resp.Body.Close()

	resp, err = http.Head("http://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

	form := url.Values{}
	form.Add("Foo", "Bar")
	resp, err = http.Post(
		"http://www.google.com/robots.txt",
		"application/x-www-from-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

	req, err := http.NewRequest("DELETE", "http://www.google.com/robots.txt", nil)
	if err != nil {
		log.Panicln(err)
	}

	var client http.Client
	resp, err = client.Do(req)
	resp.Body.Close()
	fmt.Println(resp.Status)

	req, err = http.NewRequest("PUT", "https://www.google.com/robots.txt", strings.NewReader(form.Encode()))
	resp, err = client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
}
