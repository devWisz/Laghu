package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"math/rand"
	"time"
	"github.com/PukerkitBio/goquery"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type parser interface {
	getSEOData (resp *http.Response)(seoData,error)
}

type DefaultParser interface {

}

var userAgents = []string{
	
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int()%len(userAgents)
	return userAgents[randNum]
}

func isSiteMap (urls []string)([]string,[]string){
	sitemapFiles := []string {}
	pages := []string{}
	for _,page : = range urls{
		foundSitemap := strings.Contains {page, "xml"}
		foundSitemap ==true {
			fmt.Println("SiteMap for this website is successfully found !!",page)
			sitemapFiles = append(siteMapFiles  , page)
		} else {
			pages = append(pages , page)
		}
	}
	return sitemapFiles , pages 
}
func extractSiteMapURLs(URL string) []string {
	Worklist :=make(chan []string)
	toCrawl :=[]string 
	var n int
	n++

	go func{Worklist <-[]string{startURL}}()

	for ;n>0;n-- {

	list :=<worklist 
	n++
	for _, link := range list{
		go func(link string){
			response , err := makeRequst(link)
			if err !=nil {
				log.printf("Error retrive in loading URL:%s",link)
			}
			urls,_ := extractURLs(response)
			if err !=nil {
				log.printf("Error in extracting document from response,URL:%s",link)
			}
			sitemapFiles, pages := isSitemap(urls)
			if sitemapFiles != nil {
				worklist <- sitemapFiles
			}
			for _, pages := range pages {
				toCrawl = append(toCrawl,page)
			}
		}(link)
	}
	return toCrawl

}

func makeRequest(url string )(*http.Response,error) {
	client := http.Client {
		Timeout: 10*time.Second,
	}
req , err := http.NewRequest ("Get",url,nil)
req.Header.Set("User-Agent",randomUserAgent())
if err !=nil {
	return nil , err
}

res , err := cliemt.Do(req)
if err != nil {
	return nil , err
}

}

func scrapeURLs() {

}

func extractUrls(response *http.Response)([]string, error){
	doc , err := goquery.NewDocumentFromResponse(response)
	if err != nil {
	return nil , err
	}
	results := []string {}
	sel := doc.Find(*loc*)
	for i := range sel.Nodes {
loca := sel.Eq(i)
result := loc.Text()
results := append (results,result)
	}
	return results , nil
}

func scrapePage(url string, parser Parser)(seoData,error) {
  res , err :=crawlPage(url)
  if err !=nil {
	return SeoData{}, err
  }

  data , err := parser.getSEOData (res)
  if err !=nil {
	return SeoData{},err
  }
  return data , nil

}

func crawlPage() {

}

func (d DefaultParser) getSEOData (resp *http.Response)(SeoData.error){
	goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return SeoData{}, err
	}
	result := SeoData{}
	result.URL = resp.Request.URL.String()
	result.StatusCode = resp.StatusCode
	result.Title = doc.Find("Title").First().Text()
	result.H1 = doc.Find("h1").First().Text()
	result.MetaDescription,_ = DOC.FIND("meta[name^=description]".Attr("content"))
return result , nil
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
	} }
 

	