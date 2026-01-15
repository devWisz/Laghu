package main

import (
	"fmt"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type parser interface {
}

func extractSiteMapURLs(URL string) []string {
	worklist :=make(chan []string)
	toCrawl :=[]string 

	go func{worklist <-[]string{startURL}}()

	

}

func makeRequest() {

}

func scrapeURLs() {

}

func scrapePage() {

}

func crawlPage() {

}

func scrapeSiteMap(url string)[]SeoData {
	results := extractSiteMapURLs (url)
	res :=scrapeURLs(results)

}

func main() {
	fmt.Println("Welcome to the site !!")
	p := DefaultParser{}
	results := scrapeSiteMap("")
	for _,res := range results {
		fmt.Println(res)
	}


}


