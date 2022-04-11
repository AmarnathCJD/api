package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func PyPi(query string) []Package {
	baseUrl := "https://pypi.org/search/?q=" + url.QueryEs cape(query)
	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Print(err)
	}
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	var R []Package
	doc.Find(".package-snippet").Each(func(i int, s *goquery.Selection) {
		R = append(R, Package{s.Find(".package-snippet__name").Text(), s.Find(".package-snippet__version").Text(), s.Find(".package-snippet__released").Text(), s.Find(".package-snippet__description").Text()})
	})
	return R

}

func main() {
	Q := PyPi("requests")
	for _, v := range Q {
		fmt.Printf("%s", v)
	}
}

type Package struct {
	Name        string
	Version     string
	Released    string
	Description string
}
