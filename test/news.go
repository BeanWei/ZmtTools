package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
	"github.com/json-iterator/go"
)

func main() {

	baseurl := "http://m.ynet.com/data/getlistinfo/"

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	queues, _ := queue.New(
		100, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "zixun.ynet.com")
		r.Headers.Set("Referer", "http://www.ynet.com")
		log.Println("Visiting: ", r.URL.String())
		r.Ctx.Put("url", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		jsonbyte := r.Body
		field := jsoniter.Get(jsonbyte, "channelName").ToString()
		if jsoniter.Get(jsonbyte, "articles", 0).ToString() != "" {
			channelId := jsoniter.Get(jsonbyte, "channelId").ToString()
			log.Println(field + " == " + channelId)
		}

		// url := jsoniter.Get(jsonbyte, "articles", index, "newsUrl").ToString()
		// title := jsoniter.Get(jsonbyte, "articles", index, "title").ToString()
		// author := jsoniter.Get(jsonbyte, "articles", index, "media").ToString()
		// publishtime := jsoniter.Get(jsonbyte, "articles", index, "updateTime").ToString()
		// cover := jsoniter.Get(jsonbyte, "articles", index, "images", 0, "url").ToString()

		// log.Println(field, title, author, publishtime, url, cover)

		// //数据入库
		// check = exist(db, url)
		// if _, ok := check.(bool); ok {
		// 	log.Println("此新闻已入库: ", url)
		// 	continue
		// } else {
		// 	err := newsadd(db, field, title, author, publishtime, 0, 0, url, cover)
		// 	if err != nil {
		// 		log.Fatal("Line54-Error: ", err)
		// 	}
		// }
	})

	//pagelist := []int{1, 2}
	//pagelist := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 959, 981, 1014, 1039, 1067, 10376, 10386, 10530, 10802, 11195}

	for id := 1; id < 101; id++ {
		queues.AddURL(fmt.Sprintf("%s?channelId=%d&page=1", baseurl, id))
	}

	queues.Run(c)
}
