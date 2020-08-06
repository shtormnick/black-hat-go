package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com/robots.txt")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(body))
	resp.Body.Close()
}