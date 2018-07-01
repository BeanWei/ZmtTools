package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
	"github.com/json-iterator/go"
)

//初始化数据库
var db = dbinit()

// 检查准备入库的新闻是否已存在
var check interface{}

// tencent 腾讯新闻 TODO: 暂时只分析了娱乐版块
func tencent() {

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
		field := "娱乐"
		for index := 0; index < total; index++ {
			url := jsoniter.Get(jsonbyte, "data", index, "surl").ToString()
			title := jsoniter.Get(jsonbyte, "data", index, "title").ToString()
			author := jsoniter.Get(jsonbyte, "data", index, "source").ToString()
			publishtime := jsoniter.Get(jsonbyte, "data", index, "publish_time").ToString()
			views := jsoniter.Get(jsonbyte, "data", index, "view_count").ToInt()
			comments := jsoniter.Get(jsonbyte, "data", index, "comment_num").ToInt()
			cover := jsoniter.Get(jsonbyte, "data", index, "bimg").ToString()
			//数据入库
			check = exist(db, url)
			if _, ok := check.(bool); ok {
				log.Println("此新闻已入库: ", url)
				continue
			} else {
				err := newsadd(db, field, title, author, publishtime, views, comments, url, cover)
				if err != nil {
					log.Fatal("Line51-Error: ", err)
				}
			}
		}
	})

	c.Visit("https://pacaio.match.qq.com/irs/rcd?cid=56&ext=ent&token=c786875b8e04da17b24ea5e332745e0f&num=1000&page=0&expIds=&callback=__jp1")
}

// zaker zaker新闻所有版块第一页36个数据
func zaker() {

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
		field := e.ChildText(".zk-wap-logotile")
		e.ForEach("#infinite_scroll a", func(_ int, el *colly.HTMLElement) {
			url := el.Attr("href")
			title := el.ChildText(".title")
			author := strings.Split(el.ChildText(".author"), " ")[0]
			//辅助函数(utils.go中)，解析汉字时间
			publishtime, err := dateparse(el.ChildText(".date"))
			if err != nil {
				log.Fatal(err)
			}
			bg := regexp.MustCompile(`url\(\/\/(.*?)\);`).FindStringSubmatch(el.ChildAttr(".pic-cover", "style"))
			var cover string
			if len(bg) == 2 {
				cover = bg[1]
			} else {
				cover = "http://zkres.myzaker.com/static/zaker_web2/img/logo.png?v=20170726"
			}
			//数据入库
			check = exist(db, url)
			if _, ok := check.(bool); ok {
				log.Println("此新闻已入库: ", url)
			} else {
				err := newsadd(db, field, title, author, publishtime, 0, 0, url, cover)
				if err != nil {
					log.Fatal("Line105-Error: ", err)
				}
			}
		})
	})

	pagelist := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 959, 981, 1014, 1039, 1067, 10376, 10386, 10530, 10802, 11195}
	for _, i := range pagelist {
		queues.AddURL(fmt.Sprintf("%s%d", baseurl, i))
	}

	queues.Run(c)
}
