package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
)

func main() {

	baseurl := "http://app.myzaker.com/index.php?app_id="

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	queues, _ := queue.New(
		26, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "app.myzaker.com")
		r.Headers.Set("Referer", "http://www.myzaker.com")
		log.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach("#infinite_scroll a", func(_ int, el *colly.HTMLElement) {
			author := el.ChildText(".author span:first-child")
			//辅助函数(utils.go中)，解析汉字时间
			publishtime := el.ChildText(".date")
			log.Println(author)
			log.Println(publishtime)
			//数据入库
			// check = exist(db, url)
			// if _, ok := check.(bool); ok {
			// 	log.Println("此新闻已入库: ", url)
			// } else {
			// 	err := newsadd(db, field, title, author, publishtime, 0, 0, url, cover)
			// 	if err != nil {
			// 		log.Fatal("Line106-Error: ", err)
			// 	}
			// }
		})
	})
	pagelist := []int{0, 1}
	//pagelist := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 959, 981, 1014, 1039, 1067, 10376, 10386, 10530, 10802, 11195}
	for _, i := range pagelist {
		queues.AddURL(fmt.Sprintf("%s%d", baseurl, i))
	}

	queues.Run(c)
}
