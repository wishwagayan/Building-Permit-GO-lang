package spider

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
var (
	// 	c.Visit("http://go-colly.org")
	// }
	count = 0
	url string = "https://permitsearch.jeffco.us/index.cfm?fuseaction=PropertySearchFormBldg"
)

func GetData(w http.ResponseWriter, r *http.Request) {

	//Ceating new Data Collector
	collector := colly.NewCollector()

	//slices to store the data
	var response []string

	collector.OnHTML("form[name]", func(e *colly.HTMLElement) {
		name := e.Attr("name")
		if strings.Contains(name, "CFForm_1") {
			err := collector.Post(url, map[string]string{"HouseNum": "admin", "password": "admin"})
			if err != nil {
				log.Fatal(err)
			}
		}
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
	// visit url
	collector.Visit(url)






}
