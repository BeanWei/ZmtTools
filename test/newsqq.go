package main

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/json-iterator/go"
)

func main() {

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "ent.163.com")
		r.Headers.Set("Referer", "http://ent.163.com/")
		log.Println("Visiting: ", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		jsonbyte := []byte(r.Body[15:(len(r.Body) - 1)])
		total := jsoniter.Get(jsonbyte, "datanum").ToInt()
		for index := 0; index < total; index++ {
			title := jsoniter.Get(jsonbyte, "data", index, "title").ToString()
			url := jsoniter.Get(jsonbyte, "data", index, "surl").ToString()
			log.Println(title)
			log.Println(url)
		}
	})

	c.Visit("http://ent.163.com/special/000380VU/newsdata_index.js?callback=data_callback")
}
