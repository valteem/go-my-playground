package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {

	urlRaw := "http://example.com/stores/main/items/thing?color=black&size=medium,large"
	urlParsed, err := url.Parse(urlRaw)
	if err != nil {
		log.Fatal("error parsing URL")
		return
	}
	fmt.Println(urlParsed.Host, urlParsed.Path, urlParsed.RawQuery)
	qryParam, err := url.ParseQuery(urlParsed.RawQuery)
	if err != nil {
		log.Fatal("error parsing query params")
		return
	}
	fmt.Println(qryParam)

	} 