package main

import "fmt"

type Page struct {
	number int
	pict   Picture
}

type Picture struct {
	orientation string
	content     string
}

func main() {

	p := Page{number: 7, pict: Picture{orientation: "album", content: "town"}}

	fmt.Println(p)

	var pict1 Picture
	pict1.orientation = "landscape"
	pict1.content = "nature"

	var page1 Page
	page1.number = 11
	page1.pict = pict1

	fmt.Println(page1, page1.pict.content)

}
