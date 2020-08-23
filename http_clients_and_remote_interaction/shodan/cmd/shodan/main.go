package main

import (
	"fmt"
	"log"
	"os"
	
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan <searchterm>")
	}

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf(
		"Querry Credits: %d\nScan Credits: %d\n\n",
		info.QuerryCredits,
		info.ScanCredits,
	)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}