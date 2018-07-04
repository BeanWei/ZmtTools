package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
	"github.com/json-iterator/go"
)

// Result 自己需要的结构
type Result struct {
	Domain     string
	Title      string
	Type       string
	PublicTime string
	Read       int
	Comment    int
	Srcurl     string
}

// IDspider 为根据指定的用户ID进行爬虫的入口函数
func IDspider(reqinfo *IDpost) (results []Result) {

	platform := reqinfo.Platform
	authorid := reqinfo.AuthorID
	readlimit := reqinfo.Readlimit
	datefrom := fmt.Sprintf("%s 00:00:00", reqinfo.Datefrom)
	dateto := fmt.Sprintf("%s 23:59:59", reqinfo.Dateto)

	switch platform {
	case "大鱼号":
		results := Dayu(authorid, readlimit, datefrom, dateto)
		return results
	case "百家号":
		return nil
	case "企鹅号":
		return nil
	case "头条号":
		return nil
	case "十万爆文":
		results := Leguan(authorid, readlimit, datefrom, dateto)
		return results
	default:
		return nil
	}
}

/*================爬===虫===入===口==========================*/

/*根据前台传入的指定作者ID进行定向爬取文章*/

/*	*		*		*		*			大鱼号		*		 *		*/
// Resp 响应结构
type Resp struct {
	Data []Production `json:"data"`
}

// Production 作者创作列表json结构
type Production struct {
	ID          string `json:"content_id"`   //作品ID
	Type        int    `json:"format_type"`  //作品类型
	Category    string `json:"category"`     //作品分类
	Title       string `json:"title"`        //作品标题
	Publishtime string `json:"published_at"` //发表时间
}

// Dayu 大鱼号
func Dayu(AuthorID, Readlimit, Timefrom, Timeto string) (results []Result) {

	results = []Result{}

	c := colly.NewCollector(
		colly.Async(true),
	)
	//设置请求超时
	c.SetRequestTimeout(200 * time.Second)
	extensions.RandomUserAgent(c)
	extensions.Referrer(c)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		r.Request.Retry()
	})

	getHotCollector := c.Clone()
	getHotCollector.OnResponse(func(resp *colly.Response) {
		result := Result{}
		//TODO:分析最符合的阅读量
		// reg := regexp.MustCompile(`"click_total":(.*?),`)
		// result.Hot, _ = strconv.Atoi(reg.FindString(string(resp.Body)))
		reg1 := regexp.MustCompile(`"click1":(.*?)`)
		reg2 := regexp.MustCompile(`"click2":(.*?),`)
		reg3 := regexp.MustCompile(`"click3":(.*?),`)
		c1, err := strconv.Atoi(reg1.FindStringSubmatch(string(resp.Body))[1])
		c2, err := strconv.Atoi(reg2.FindStringSubmatch(string(resp.Body))[1])
		c3, err := strconv.Atoi(reg3.FindStringSubmatch(string(resp.Body))[1])
		if err != nil {
			log.Fatal(err)
		}
		result.Read = c1 + c2 + c3
		//TODO: 获取大鱼号图文的评论数
		result.Comment = 0
		id := resp.Request.Ctx.Get("id")
		result.Srcurl = fmt.Sprintf("http://a.mp.uc.cn/article.html?wm_cid=%s", id)
		result.Domain = resp.Request.Ctx.Get("category")
		result.Title = resp.Request.Ctx.Get("title")
		result.PublicTime = resp.Request.Ctx.Get("publishtime")

		theType := resp.Request.Ctx.Get("type")
		if theType == "1001" {
			result.Type = "图文"
		} else if theType == "1002" {
			result.Type = "视频"
		} else if theType == "1005" {
			result.Type = "图集"
		} else {
			result.Type = "未知类型" + theType
		}

		// 根据阅读量筛选
		hv, _ := strconv.Atoi(Readlimit)
		if hv <= result.Read {
			results = append(results, result)
		}

	})

	c.OnResponse(func(resp *colly.Response) {
		productionlist := &Resp{}
		if err := json.Unmarshal([]byte(string(resp.Body)), productionlist); err != nil {
			log.Fatal(err)
		}
		for _, oneproduction := range productionlist.Data {

			//根据发布的时间段筛选
			changetime := strings.Replace(strings.Split(oneproduction.Publishtime, ".")[0], "T", " ", -1)
			timeline, err := time.Parse("2006-01-02 15:04:05", changetime)

			tfrom, err := time.Parse("2006-01-02 15:04:05", Timefrom)
			tto, err := time.Parse("2006-01-02 15:04:05", Timeto)

			if err != nil {
				log.Fatal(err)
			}
			if tfrom.Before(timeline) && tto.After(timeline) {
				id := oneproduction.ID
				timel := timeline.Format("2006-01-02 15:04:05")

				targetURL := fmt.Sprintf("http://ff.dayu.com/contents/%s?biz_id=1002&_fetch_author=1&_fetch_incrs=1", id)
				ctx := colly.NewContext()
				ctx.Put("id", oneproduction.ID)
				ctx.Put("type", strconv.Itoa(oneproduction.Type))
				ctx.Put("category", oneproduction.Category)
				ctx.Put("title", oneproduction.Title)
				ctx.Put("publishtime", timel)
				getHotCollector.Request("GET", targetURL, nil, ctx, nil)
				log.Println("Visiting: ", targetURL)
			}
		}

	})

	visitUrl := fmt.Sprintf("http://ff.dayu.com/contents/author/%s?biz_id=1002&_size=1000000", AuthorID)
	c.Visit(visitUrl)

	getHotCollector.Wait()

	return

}

/*	*		*		*		*		乐观号平台		*		 *		*/
// 乐观号平台 json 数据结构的构造
type Obj struct {
	Content []LeguanObj `json:"content"`
}
type LeguanObj struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Platform    string `json:"platform"`
	Account     string `json:"account"`
	Domain      string `json:"domain"`
	Status      int    `json:"status"`
	Read        int    `json:"read"`
	Like        int    `json:"like"`
	Comment     int    `json:"comment"`
	Favor       int    `json:"favor"`
	Origin      int    `json:"origin"`
	MediaType   int    `json:"mediaType"`
	Coverpic    string `json:"coverPic"`
	Sn          string `json:"sn"`
	Publictime  string `json:"publicTime"`
	Createtime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	Lptime      int64  `json:"lptime"`
	IsCrawl     int    `json:"isCrawl"`
	Crawl       int    `json:"crawl"`
	Area        string `json:"area"`
	Province    string `json:"province"`
	ChildDomain string `json:"childDomain"`
	UserID      string `json:"userId"`
	App         int    `json:"app"`
	Apponline   int    `json:"appOnline"`
	Priority    int    `json:"priority"`
	Social      int    `json:"social"`
	Havehtml    int    `json:"haveHtml"`
	ES          int    `json:"es"`
	Logo        string `json:"logo"`
	Nickname    string `json:"nickName"`
}

// Leguan 乐观号平台的爬取
func Leguan(AuthorID, Readlimit, Timefrom, Timeto string) (results []Result) {
	results = []Result{}
	relimit, err := strconv.Atoi(Readlimit)
	tfrom, err := time.Parse("2006-01-02 15:04:05", Timefrom)
	tto, err := time.Parse("2006-01-02 15:04:05", Timeto)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.PostForm(
		"http://www.myleguan.com/lg_res/focus/flr/gaibui",
		url.Values{
			"ids":          {AuthorID},
			"lgCustomerId": {"1000183092"},
			"lptimeOrder":  {"DESC"},
			"size":         {"10000"}})

	if err != nil {
		log.Println(err.Error())
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}

	reObj := jsoniter.Get(body, "reObj").ToString()

	content := &Obj{}
	err = jsoniter.UnmarshalFromString(reObj, content)
	if err != nil {
		log.Fatal(err)
	}
	result := Result{}
	for _, item := range content.Content {
		result.Domain = item.Domain
		result.Title = item.Title
		result.PublicTime = item.Publictime
		result.Read = item.Read
		result.Comment = item.Comment
		result.Srcurl = item.URL

		if item.MediaType == 1 {
			result.Type = "图文"
		} else if item.MediaType == 2 {
			result.Type = "视频"
		} else {
			result.Type = "未知"
		}

		timeline, err := time.Parse("2006-01-02 15:04:05", result.PublicTime)
		if err != nil {
			log.Fatal(err)
		}

		if relimit <= result.Read && tfrom.Before(timeline) && tto.After(timeline) {
			results = append(results, result)
		}
	}
	return results
}



/*==========================================================================================*/

/*利用目前的站点采集所有领域的文章并存入数据库*/

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
					log.Fatal("Line53-Error: ", err)
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
			author := el.ChildText(".author span:first-child")
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
					log.Fatal("Line106-Error: ", err)
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

// jxnews 头条新闻，来自江西的一个新闻聚合站点，爬取每个领域的前一百条新闻
func jxnews() {

	baseurl := "http://toutiao.jxnews.com.cn/m/channelnewslist.php?pagesize=100&cate="

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	queues, _ := queue.New(
		32, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "toutiao.jxnews.com.cn")
		r.Headers.Set("Referer", "http://toutiao.jxnews.com.cn")
		log.Println("Visiting: ", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		jsonbyte := r.Body
		for index := 0; index < 100; index++ {
			field := jsoniter.Get(jsonbyte, index, "classname").ToString()
			url := fmt.Sprintf("http://toutiao.jxnews.com.cn%s", jsoniter.Get(jsonbyte, index, "titleurl").ToString())
			title := jsoniter.Get(jsonbyte, index, "title").ToString()
			author := jsoniter.Get(jsonbyte, index, "befrom").ToString()
			publishtime, err := dateparse(jsoniter.Get(jsonbyte, index, "newstime").ToString())
			if err != nil {
				log.Fatal(err)
			}
			cover := jsoniter.Get(jsonbyte, index, "titlepic").ToString()
			//数据入库
			check = exist(db, url)
			if _, ok := check.(bool); ok {
				log.Println("此新闻已入库: ", url)
				continue
			} else {
				err := newsadd(db, field, title, author, publishtime, 0, 0, url, cover)
				if err != nil {
					log.Fatal("Line159-Error: ", err)
				}
			}
		}
	})

	pagelist := []int{47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 71, 75, 76, 77, 78, 79, 80, 81, 83}
	for _, i := range pagelist {
		queues.AddURL(fmt.Sprintf("%s%d", baseurl, i))
	}

	queues.Run(c)
}

// ynet 北青网，爬取各个领域前1000条新闻
func ynet() {

	baseurl := "http://m.ynet.com/data/getlistinfo/"

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	queues, _ := queue.New(
		520, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "zixun.ynet.com")
		r.Headers.Set("Referer", "http://www.ynet.com")
		log.Println("Visiting: ", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		jsonbyte := r.Body
		field := jsoniter.Get(jsonbyte, "channelName").ToString()
		for index := 0; index < 50; index++ {
			url := jsoniter.Get(jsonbyte, "articles", index, "newsUrl").ToString()
			title := jsoniter.Get(jsonbyte, "articles", index, "title").ToString()
			author := jsoniter.Get(jsonbyte, "articles", index, "media").ToString()
			publishtime := jsoniter.Get(jsonbyte, "articles", index, "updateTime").ToString()
			cover := jsoniter.Get(jsonbyte, "articles", index, "images", 0, "url").ToString()

			//数据入库
			check = exist(db, url)
			if _, ok := check.(bool); ok {
				log.Println("此新闻已入库: ", url)
				continue
			} else {
				err := newsadd(db, field, title, author, publishtime, 0, 0, url, cover)
				if err != nil {
					log.Fatal("Line210-Error: ", err)
				}
			}
		}
	})

	idlist := []int{30, 4, 11, 37, 19, 1, 24, 12, 2, 5, 3, 23, 13, 10, 8, 18, 26, 16, 25, 15, 27, 36, 28, 29, 6, 20}
	for _, id := range idlist {
		for page := 1; page < 21; page++ {
			queues.AddURL(fmt.Sprintf("%s?channelId=%d&page=%d", baseurl, id, page))
		}
	}

	queues.Run(c)
}
