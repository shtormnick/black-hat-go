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
	r1, err := http.Get("http://www.google.com/robots.txt")
	defer r1.Body.Close()

	if err != nil {
		log.Panicln(err)
	}

	r2, err := http.Head("http://www.google.com/robots.txt")
	defer r2.Body.Close()

	if err != nil {
		log.Panicln(err)
	}

	form := url.Values{}
	form.Add("Foo", "Bar")
	req, err := http.NewRequest(
		"PUT",
		"https://www.google.com/robots.txt",
		strings.NewReader(form.Encode()),
	)

	fmt.Println(req)

	if err != nil {
		log.Panicln(err)
	}

	req, err = http.NewRequest("DELETE", "https://www.google.com/robots.txt",
		nil)

	resp, err := http.Get("https://www.google.com/robots.txt")
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

}
