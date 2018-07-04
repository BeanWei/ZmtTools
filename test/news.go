package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var AuthorID = "5a8a9f033817831a28afd45b"
var Readlimit = "100000"
var Timefrom = "2018-06-01 00:00:00"
var Timeto = "2018-07-02 00:00:00"

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

func main() {
	results := []Result{}
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
	log.Println(results)

}
