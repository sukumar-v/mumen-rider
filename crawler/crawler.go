package crawler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocolly/colly/v2"
)

// Ping to test server status
func Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Pong")
	w.Write([]byte("Pong"))
}

// GetUrls to get all urls in a page
func GetUrls(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		log.Println("Missing 'url' argument")
		return
	}
	log.Println("Visiting", url)

	c := colly.NewCollector()

	var urls []string

	c.OnHTML("a[href]", func (e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			urls = append(urls, link)
		}
	})

	c.Visit(url)

	b, err := json.Marshal(urls)
	if err != nil {
		log.Println("Failed to serialize response: ", err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(b)
}
