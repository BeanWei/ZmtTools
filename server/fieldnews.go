package main

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/json-iterator/go"
)

func qqnews() {

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "pacaio.match.qq.com")
		r.Headers.Set("Referer", "https://xw.qq.com")
		log.Println("Visiting: ", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		jsonbyte := []byte(r.Body[6:(len(r.Body) - 1)])
		total := jsoniter.Get(jsonbyte, "datanum").ToInt()
		for index := 0; index < total; index++ {
			url := jsoniter.Get(jsonbyte, "data", index, "surl").ToString()
			title := jsoniter.Get(jsonbyte, "data", index, "title").ToString()
			author := jsoniter.Get(jsonbyte, "data", index, "source").ToString()
			authorID := jsoniter.Get(jsonbyte, "data", index, "source_id").ToString()
			publishtime := jsoniter.Get(jsonbyte, "data", index, "publish_time").ToString()
			views := jsoniter.Get(jsonbyte, "data", index, "view_count").ToInt()
			comments := jsoniter.Get(jsonbyte, "data", index, "comment_num").ToInt()
			log.Println(url, title, author, authorID, publishtime, views, comments)
			//TODO: 数据入库
		}
	})

	c.Visit("https://pacaio.match.qq.com/irs/rcd?cid=56&ext=ent&token=c786875b8e04da17b24ea5e332745e0f&num=1000&page=0&expIds=&callback=__jp1")
}
