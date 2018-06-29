package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)


// Result 自己需要的结构
type Result struct {
	Domain	string	
	Title		string
	Type		string
	PublicTime	string
	Read		int
	Comment		int
	Srcurl		string
}

// IDspider 为根据指定的用户ID进行爬虫的入口函数
func IDspider( reqinfo *IDpost) (results []Result) {

	platform := reqinfo.Platform
	authorid := reqinfo.AuthorID
	readlimit := reqinfo.Readlimit
	datefrom := reqinfo.Datefrom
	dateto := reqinfo.Dateto

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
	default:
		return nil 
	}
}




/*================爬===虫===入===口==========================*/

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
	Coverlink   string `json:"cover_url"`    //封面链接
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
		result.Srcurl = fmt.Sprintf("html: http://a.mp.uc.cn/article.html?uc_param_str=frdnsnpfvecpntnwprdssskt&from=media#!wm_cid=%s", id)
		result.Type = resp.Request.Ctx.Get("type")
		result.Domain = resp.Request.Ctx.Get("category")
		result.Title = resp.Request.Ctx.Get("title")
		result.PublicTime = resp.Request.Ctx.Get("publishtime")

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
			if tfrom.Before(timeline) && tto.Before(timeline) {
				id := oneproduction.ID
				timel := timeline.Format("2006-01-02 15:04:05")

				targetURL := fmt.Sprintf("http://ff.dayu.com/contents/%s?biz_id=1002&_fetch_author=1&_fetch_incrs=1", id)
				ctx := colly.NewContext()
				ctx.Put("id", oneproduction.ID)
				ctx.Put("type", strconv.Itoa(oneproduction.Type))
				ctx.Put("category", oneproduction.Category)
				ctx.Put("title", oneproduction.Title)
				ctx.Put("coverlink", oneproduction.Coverlink)
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