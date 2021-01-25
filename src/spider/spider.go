package spider

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

// func Phase() {

// 	c := colly.NewCollector()

// 	// Find and visit all links
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr("href"))
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	c.Visit("http://go-colly.org")
// }

func GetData(w http.ResponseWriter, r *http.Request) {
	//verify the param "URL" Exists
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("missing URL Argument")
		return
	}
	log.Println("visiting", URL)

	//Ceating new Data Collector
	collector := colly.NewCollector()

	//slices to store the data
	var response []string

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			response = append(response, link)
			log.Println("adding..."+link)
		}
		// visit url
		collector.Visit("http://go-colly.org")


	})


	//handeling errors
	collector.OnError(func(response *colly.Response, err error) {
		log.Println("error is %v",err)
	})

	// parse our response slice into json Format
	b, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to serilize response", err)
		return
	}
	// add some header and write the body for our endpoint
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
	collector.Visit(URL)





}
