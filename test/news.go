package main

import (
	//"io/ioutil"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	//"net/http"
	//"net/url"
	//"strings"
)

var title = "大一忙考研小学暑假仅半天 如此“抢跑”值不值？"
var content = "幼儿园越过小学知识量，小学越过初中词拼音等，上课就等于把之前学的知识再学一遍，很多学生就会出现了注意力不集中，读、写都不规范等情况。”“不管是哪个阶段的孩子，其理解能力、专注度都是有限的，如果超负荷安排学习计划，很容易产生逆反、排斥心理，不仅学不好，还会产生厌学情绪，结果会适得其反。"

// func main() {
// 	postData := fmt.Sprintf(`{"title": %s, "content": %s}`, title, content)
// 	//postData := `{"content": "幼儿园越过小学知识量，小学越过初中词拼音等，上课就等于把之前学的知识再学一遍，很多学生就会出现了注意力不集中，读、写都不规范等情况。”“不管是哪个阶段的孩子，其理解能力、专注度都是有限的，如果超负荷安排学习计划，很容易产生逆反、排斥心理，不仅学不好，还会产生厌学情绪，结果会适得其反。", "title": "大一忙考研小学暑假仅半天 如此“抢跑”值不值？ "}`
// 	payload := []byte(postData)
// 	c := colly.NewCollector()
// 	c.SetRequestTimeout(3 * time.Minute)
// 	c.OnRequest(func(r *colly.Request) {
// 		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
// 		r.Headers.Set("Accept", "application/json, text/javascript, */*; q=0.01")
// 		r.Headers.Set("Origin", "http://www.myleguan.com")
// 		r.Headers.Set("Referer", "http://www.myleguan.com/lgEditor/lgEditor.html")
// 		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

// 	})
// 	c.OnResponse(func(r *colly.Response) {
// 		log.Println("response received", r.StatusCode)
// 		log.Println(string(r.Body))
// 	})
// 	err := c.PostRaw("http://120.35.10.209:10099/Keyword/checkoriginal", payload)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// err := c.Post(
// 	// 	"http://120.35.10.209:10099/Keyword/checkoriginal",
// 	// 	map[string]string{"title": title, "content": content})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// c.OnResponse(func(r *colly.Response) {
// 	// 	log.Println("response received", r.StatusCode)
// 	// 	log.Println(string(r.Body))
// 	// })
// 	// c.Visit("http://120.35.10.209:10099/Keyword/checkoriginal")
// }

func main() {
	client := &http.Client{}
	var postData = url.Values{}
	postData.Add("title", title)
	postData.Add("content", content)
	data := postData.Encode()
	log.Println(data)
	req, err := http.NewRequest(
		"POST",
		"http://120.35.10.209:10099/Keyword/checkoriginal",
		strings.NewReader(data))
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "http://www.myleguan.com")
	req.Header.Add("Referer", "http://www.myleguan.com/lgEditor/lgEditor.html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	//reObj := jsoniter.Get(body, "reObj").ToString()

	log.Println(string(body))

}
