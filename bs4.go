package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://pypi.org/search/?q=telethon")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	for _, link := range doc.FindAll("span", "class", "package-snippet__name") {
		fmt.Println(link.Text())
	}
}
