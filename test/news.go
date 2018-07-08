package main

import (
	"encoding/json"
	//"io/ioutil"
	//"io/ioutil"
	"log"
	//"net/http"
	//"net/url"
	//"strings"
	"time"
	//"net/http"
	//"net/url"
	//"strings"
	"github.com/gocolly/colly"
	//"github.com/gocolly/colly/extensions"
	//"github.com/gocolly/colly/queue"
	//"github.com/json-iterator/go"
)

func main() {
	var title = "大一忙考研小学暑假仅半天 如此“抢跑”值不值？"

	//var content = "幼儿园越过小学知识量，小学越过初中词拼音等，上课就等于把之前学的知识再学一遍，很多学生就会出现了注意力不集中，读、写都不规范等情况。”“不管是哪个阶段的孩子，其理解能力、专注度都是有限的，如果超负荷安排学习计划，很容易产生逆反、排斥心理，不仅学不好，还会产生厌学情绪，结果会适得其反。"

	var content = `
	人民日报：美国贸易盲动症注定引火烧身
	人民网 2018-07-08 05:11:48
	“这一天或许会成为一个经济史上臭名昭著的日子。”7月6日,美国政府对华340亿美元商品加征关税的非理性举措立即遭到国际舆论抨击,就连美国众多有影响力的媒体也是一片责难之声。公道自在人心,白宫逞一时之“勇”发动贸易战,不仅冲击正常国际贸易、扰乱全球经济,还将严重损害美国自身利益。

	近段时间,美国将关税大棒奉若神明,予取予求、唯我独尊。让美国贸易盲动症日渐加重的原因很复杂,但有一件事是明摆着的,那就是某些人公然漠视“贸易战是把双刃剑”这样一个再浅显不过的道理,仿佛由美国次贷危机演化成的国际金融危机不曾有过,仿佛苦苦挣扎多年才走向复苏早已是几个世纪前的旧事,仿佛美国经济真的炼就了任凭国际大环境冲击的金刚不坏之身。不管是因为经济学常识缺失,还是揣着明白装糊涂,当今世界头号经济体如此行事,引发人们对全球经济前景担忧的同时,也把一个更深层面的问题提了出来:如果一家店铺里真的都是不堪一击的精美瓷器,鲁莽闯入的大象将创作出怎样一幅画面?好在地球村不是店铺,国际力量对比也没有失衡到让哪个国家可以像瓷器店里的大象那样无所顾忌、横冲直撞。

	即便在美国国内,对白宫贸易盲动症的质疑也是有力的,对引火烧身的担忧也是真切的。从哈雷—戴维森摩托车公司“逃离行动”,到通用汽车公司的收缩计划,白宫对外发动贸易战以重振国内制造业和保护国内就业的所谓初衷被打上一个个大大的问号。

	美国的企业急了——“这会给像我们公司一样的美国中小企业带来沉重打击”。贸易战不仅将阻隔供应链高效整合,还将推高产品价格、抑制企业投资,无异于在企业内部竖起一堵“墙”。美国信息技术创新基金会的研究报告显示,如果对从中国进口的信息和通讯技术产品加征25%关税,将导致美国经济未来10年损失约3320亿美元。

	美国的民众焦虑了——“说实话,我的工作是中国人给的”。美国全国对外贸易理事会统计,为反制美国关税举措,目前美国主要贸易伙伴已准备了约900亿美元关税清单,这将令数百万美国就业岗位受到威胁。

	美国的农民不干了——“我们需要继续同中国做生意,而不被这些关税措施所干扰”。贸易战开打后,每吨贵了约600元人民币的美国大豆失去价格优势,在中国市场难寻买主。美国大豆协会预计,贸易战将使今年美国大豆产值减少60亿美元。

	贸易战没有赢家。当下,华盛顿一些人甚至公然鼓噪“贸易战很好,而且很容易赢”。这既是对美国历史的无知,也是对民众利益的无视。单边关税策略在美国历史上从未成功过,反而造成就业损失等不利后果:2002年,美国政府对进口钢铁产品加征关税,导致美国净损失约20万个就业岗位；2009年,美国政府对从中国进口轮胎加征关税,导致美国净损失约2500个就业岗位、消费者增加11亿美元开支。更早之前的1930年,美国大幅提升进口商品关税,招致许多国家的报复性关税措施,美国进口额和出口额骤降50%以上,经济陷入长期萧条。

	国际贸易的重要性不言而喻,这份大家业需要各国一道用心呵护,需要在调整中不断做大做强,让每一方都有足够的获得感和舒适度。发言权确实也有美国一份,谁都无意于贬损美国世界第一大经济体的块头。但是,这绝不意味着大块头就拥有随心所欲、仗势欺人的权力。大家的事只能由大家商量着办。当家的不闹事。但愿美国能从这句充满东方智慧的大白话中有所醒悟,早日医治好自己有百害而无一利、到头来引火烧身的贸易盲动症。
	`
	result := originalCk(title, content)
	log.Println(result)
}
func originalCk(title, content string) string {

	var result string
	jsdata := map[string]string{
		"title":   title,
		"content": content}
	payload, err := json.Marshal(jsdata)
	if err != nil {
		return err.Error()
	}
	c := colly.NewCollector()
	c.SetRequestTimeout(3 * time.Minute)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		r.Headers.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		r.Headers.Set("Origin", "http://www.myleguan.com")
		r.Headers.Set("Referer", "http://www.myleguan.com/lgEditor/lgEditor.html")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

	})
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		result = string(r.Body)
	})
	err = c.PostRaw("http://120.35.10.209:10099/Keyword/checkoriginal", payload)
	if err != nil {
		log.Fatal(err)
	}
	return result
	// err := c.Post(
	// 	"http://120.35.10.209:10099/Keyword/checkoriginal",
	// 	map[string]string{"title": title, "content": content})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// c.OnResponse(func(r *colly.Response) {
	// 	log.Println("response received", r.StatusCode)
	// 	log.Println(string(r.Body))
	// })
	// c.Visit("http://120.35.10.209:10099/Keyword/checkoriginal")
}

// func originalCk(title, content string) string {
// 	client := &http.Client{
// 		Timeout: time.Minute,
// 	}
// 	var postData = url.Values{}
// 	postData.Add("title", title)
// 	postData.Add("content", content)
// 	data := postData.Encode()
// 	req, err := http.NewRequest(
// 		"POST",
// 		"http://120.35.10.209:10099/Keyword/checkoriginal",
// 		strings.NewReader(data))
// 	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("Origin", "http://www.myleguan.com")
// 	req.Header.Add("Referer", "http://www.myleguan.com/lgEditor/lgEditor.html")
// 	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err.Error()
// 	}

// 	//reObj := jsoniter.Get(body, "reObj").ToString()

// 	return string(body)

// }
