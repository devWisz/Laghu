Go SEO Sitemap Crawler

A lightweight, concurrent SEO crawler written in Go that starts from a sitemap URL, discovers nested sitemaps, extracts page URLs, and collects essential SEO metadata from each page.



Features

* Recursive sitemap discovery from a starting sitemap URL
* Concurrent crawling with configurable concurrency limits
* Pluggable SEO parser via a clean interface
* Extraction of core SEO metadata (title, H1, meta description, status code)
* Rotating User-Agent headers
* HTTP request timeouts for stability

## Installation

### Prerequisites

* Go **1.20+** (recommended)
* Internet connection

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-seo-sitemap-crawler.git
cd go-seo-sitemap-crawler
```

### Install Dependencies

```bash
go mod init go-seo-sitemap-crawler
go get github.com/PuerkitoBio/goquery
```

## How to Run

1. Update the sitemap URL in `main.go`

```go
results := ScrapeSitemap("https://example.com/sitemap.xml", p, 10)
```

2. Run the program

```bash
go run main.go
```

The crawler will fetch all sitemap URLs, scrape pages concurrently, and print SEO data to standard output.


## Benefits

* Quickly audit SEO metadata across all sitemap pages
* Efficient crawling using controlled concurrency
* Easy to extend with custom parsers or exporters
* Suitable for learning and experimenting with Go concurrency patterns

---

## Limitations

* No robots.txt handling
* No URL deduplication
* No retry or backoff strategy
* Results are printed to stdout only

---

## Disclaimer

Use responsibly. Crawling aggressive sites without permission can get your IP blocked.

---
## License
Can be used for any purposes. Open source to use and modify.

## Developed and Concept By
devWisZ aka Sarjak Khanal 