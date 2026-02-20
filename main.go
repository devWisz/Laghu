package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SeoData struct {
	URL             string `json:"url"`
	Title           string `json:"title"`
	H1              string `json:"h1"`
	MetaDescription string `json:"meta_description"`
	StatusCode      int    `json:"status_code"`
}

type SeoReport struct {
	ScrapedAt time.Time `json:"scraped_at"`
	TotalURLs int       `json:"total_urls"`
	Results   []SeoData `json:"results"`
}

type Parser interface {
	GetSeoData(resp *http.Response) (SeoData, error)
}

type DefaultParser struct{}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13)",
	"Mozilla/5.0 (X11; Linux x86_64)",
}

func randomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return userAgents[rand.Intn(len(userAgents))]
}

func makeRequest(url string) (*http.Response, error) {
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", randomUserAgent())
	return client.Do(req)
}

func extractUrls(response *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	var urls []string
	doc.Find("loc").Each(func(_ int, s *goquery.Selection) {
		urls = append(urls, strings.TrimSpace(s.Text()))
	})
	return urls, nil
}

func isSitemap(urls []string) ([]string, []string) {
	var sitemaps, pages []string
	for _, u := range urls {
		if strings.Contains(u, ".xml") {
			sitemaps = append(sitemaps, u)
		} else {
			pages = append(pages, u)
		}
	}
	return sitemaps, pages
}

func extractSitemapURLs(startURL string) []string {
	worklist := make(chan []string)
	var result []string
	var n int

	n++
	go func() { worklist <- []string{startURL} }()

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			n++
			go func(link string) {
				defer func() { worklist <- nil }()
				resp, err := makeRequest(link)
				if err != nil {
					return
				}
				urls, _ := extractUrls(resp)
				sitemaps, pages := isSitemap(urls)
				if len(sitemaps) > 0 {
					worklist <- sitemaps
				}
				result = append(result, pages...)
			}(link)
		}
	}
	return result
}

func scrapePage(url string, parser Parser) (SeoData, error) {
	resp, err := makeRequest(url)
	if err != nil {
		return SeoData{}, err
	}
	return parser.GetSeoData(resp)
}

func scrapeUrls(urls []string, parser Parser, concurrency int) []SeoData {
	tokens := make(chan struct{}, concurrency)
	results := []SeoData{}

	for _, url := range urls {
		tokens <- struct{}{}
		go func(u string) {
			defer func() { <-tokens }()
			data, err := scrapePage(u, parser)
			if err == nil {
				results = append(results, data)
			}
		}(url)
	}

	for i := 0; i < cap(tokens); i++ {
		tokens <- struct{}{}
	}

	return results
}

func (d DefaultParser) GetSeoData(resp *http.Response) (SeoData, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return SeoData{}, err
	}

	data := SeoData{
		URL:        resp.Request.URL.String(),
		StatusCode: resp.StatusCode,
		Title:      doc.Find("title").First().Text(),
		H1:         doc.Find("h1").First().Text(),
	}
	data.MetaDescription, _ = doc.Find("meta[name=description]").Attr("content")
	return data, nil
}

func SaveToJSON(data []SeoData, filename string) error {
	report := SeoReport{
		ScrapedAt: time.Now(),
		TotalURLs: len(data),
		Results:   data,
	}
	bytes, _ := json.MarshalIndent(report, "", "  ")
	return os.WriteFile(filename, bytes, 0644)
}

func SaveURLsOnlyToJSON(data []SeoData, filename string) error {
	var urls []string
	for _, d := range data {
		urls = append(urls, d.URL)
	}
	report := struct {
		ScrapedAt time.Time `json:"scraped_at"`
		TotalURLs int       `json:"total_urls"`
		URLs      []string  `json:"urls"`
	}{time.Now(), len(urls), urls}

	bytes, _ := json.MarshalIndent(report, "", "  ")
	return os.WriteFile(filename, bytes, 0644)
}

func askYesNo(question string) bool {
	var input string
	fmt.Printf("%s (y/n): ", question)
	fmt.Scanln(&input)
	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "yes"
}

func printSummary(results []SeoData, start time.Time) {
	fmt.Println("\nScraping Summary")
	fmt.Println("-----------------------------")
	fmt.Printf("Total scraped URLs : %d\n", len(results))
	fmt.Printf("Time taken         : %v\n", time.Since(start))
	fmt.Println("-----------------------------")
}

func ScrapeSitemap(url string, parser Parser, concurrency int) []SeoData {
	urls := extractSitemapURLs(url)
	return scrapeUrls(urls, parser, concurrency)
}

func main() {
	start := time.Now()
	parser := DefaultParser{}

	fmt.Println("Starting sitemap scraping...\n")

	results := ScrapeSitemap("https://openai.com/sitemap.xml", parser, 10)

	printSummary(results, start)

	if len(results) == 0 {
		fmt.Println("No URLs scraped. Exiting.")
		return
	}

	if askYesNo("Do you want to download the scraped reports") {
		SaveToJSON(results, "seo_report.json")
		SaveURLsOnlyToJSON(results, "urls_only.json")
		fmt.Println("Files saved successfully.")
	} else {
		fmt.Println("Thanks for using the Laghu.")
	}
}