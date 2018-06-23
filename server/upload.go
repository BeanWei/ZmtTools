package main

import (
	"fmt"
	"encoding/json"

	"time"
	"bytes"
	"log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
		
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var uploadUrl = ""
var filename = ""

func main() {

	a := colly.NewCollector(
		colly.Async(true),
	)
	a.SetRequestTimeout(200 * time.Second)
	extensions.RandomUserAgent(a)
	extensions.Referrer(a)

	b := a.Clone()
	c := b.Clone()
	d := c.Clone()
	e := d.Clone()
	f := e.Clone()
	g := f.Clone()
	h := g.Clone()
	i := g.Clone()
	j := g.Clone()
	k := g.Clone()
	m := g.Clone()

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile("userfile", filename)
	if err != nil {
			fmt.Println("error writing to buffer")
			return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
			fmt.Println("error opening file")
			return nil, err
	}
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
			log.Fatal("Line-56_Error:", err)
	}

	if err := a.Post(uploadUrl, map[string]byte{"data": request_reader}); err != nil {
		log.Fatal("Line-61_Error:", err)
	}
	a.OnResponse(func(r *colly.Response) {
		var oneresp struct {
			Code	int	`json:"code"`
			Message	string	`json:"message"`
			Data struct {
				ID string `json:"id"`
			}	`json:"data"`
		}
		ainfo := &oneresp
		if err := json.Unmarshal([]byte(string(r.Body)), ainfo); err != nil {
			log.Fatal("Line-73_Error:", err)
		}
		bpostinfo := map[string]string{
			"videoId": ainfo.Data.ID
			"videoName": filename
			"videoDesc": ""
			"appid": "wemedia"
			"uid": "941878fe5eb1434a8f2afe13a3874c51 "       //UID和utoken在user信息里
			"utoken": "1d24fa6122532c469e18cf838ef0838b"
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", ainfo.Data.ID)
		b.Request("POST", "https://mp.dayu.com/confirm_info", bpostinfo, ctx, nil)
	})

	b.OnResponse(func(r *colly.Response) {
		var tworesp struct {
			Succeed int	`json:"succeed"`
		}
		binfo := &tworesp
		if err := json.Unmarshal([]byte(string(r.Body)), binfo); err != nil {
			log.Fatal("Line-101_Error:", err)
		}
		if bresult := tworesp.Succeed; bresult != 1 {
			log.Fatal("第二步上传失败:", bresult)
		}
		vid :=  r.Request.Ctx.Get("videoID")
		cpostinfo := map[string]string{
			"type": "video"
			"title": filename
			"video[video_id]": vid
			"utoken": "1d24fa6122532c469e18cf838ef0838b"
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		c.Request("POST", "https://mp.dayu.com/dashboard/addMaterial", cpostinfo, ctx, nil)
	})

	c.OnResponse(func(r *colly.Response) {
		var threeresp struct {
			Data struct {
				Createdat	string	`json:"_created_at"`
				MaterialId	string 	`json:"_id"`
				Updatedat	string	`json:"_updated_at"`
			}	`json:"data"`
		}
		cinfo := &threeresp
		if err := json.Unmarshal([]byte(string(r.Body)), cinfo); err != nil {
			log.Fatal("Line-128_Error:", err)
		}
		materialId := cinfo.Data.MaterialId
		dpostinfo := map[string]string {
			"id": r.Request.Ctx.Get("videoID")
			"utoken": "1d24fa6122532c469e18cf838ef0838b"
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", materialId)
		d.Request("POST", "https://mp.dayu.com/dashboard/getVideoPlayUrl", dpostinfo, ctx, nil)
	})

	d.OnResponse(func(r *colly.Response) {
		var fourresp struct {
			PlayUrl string `json:"playUrl"`
		}
		dinfo := &fourresp
		if err := json.Unmarshal([]byte(string(r.Body)), dinfo); err != nil {
			log.Fatal("Line-128_Error:", err)
		}
		playurl := dinfo.PlayUrl
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", playurl)
		vid := r.Request.Ctx.Get("videoID")
		t := time.Now().UnixNano()
		erequrl := fmt.Sprintf("https://mp.dayu.com/media_keyframe?id=%s&_=%s", vid, t)
		epostinfo := map[string]string {
			"id": vid
			"_": t
		}
		e.Request("POST", erequrl, epostinfo, ctx, nil)
	})

	e.OnResponse(func(r *colly.Response) {
		var fiveresp struct {
			Succeed	int	`json:"succeed"`
			Api		string	`json:"api"`
			Data	struct {
				ID	string	`json:"id"`
				Reqcount	int	`json:"req_count"`
				Respcount	int	`json:"resp_count"`
				Totalcount	int	`json:"total_count"`
				Videokeyframes struct {
					Frames	[]struct {
						Imgurl	string	`json:"img_url"`
						Width	int	`json:"width"`
						Height	int	`josn:"height"`
						BlackBorderPercent	int `json:"blackBorderPercent"`
						BlackBorderLevel	string 	`josn:"blackBorderLevel"`
						BlurMark		bool	`josn:"blurMark"`
						BlackBorderMark	bool	`json:"blackBorderMark"`
						Keyframeindex	int		`json:"keyframe_index"`
					}	`json:"frames"`
				}	`json:"video_keyframes"`
			}	`json:"data"`
		}
		einfo := &fiveresp
		if err := json.Unmarshal([]byte(string(r.Body)), einfo); err != nil {
			log.Fatal("Line-185_Error:", err)
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		//默认选择第五张图为视频封面
		keyframeindex := einfo.Data.Videokeyframes.Frames[4].Keyframeindex
		imgSrc := einfo.Data.Videokeyframes.Frames[4].Imgurl
		ctx.Put("keyframeindex", keyframeindex)
		fpostinfo := map[string]string {
			"imgSrc": imgSrc
			"utoken": "1d24fa6122532c469e18cf838ef0838b"
		}
		//TODO: 构造封面选择的地址
		imgpostlink := "https://ns.dayu.com/article/outsiteImgResave?
		appid=website&wmid=941878fe5eb1434a8f2afe13a3874c51&
		sign=edc37ce0f1938f9b91aff565f7e1c51b"
		f.Request("POST", imgpostlink, fpostinfo, ctx, nil)
	})

	f.OnResponse(func(r *colly.Response) {
		var sixresp struct {
			Code 	int	`json:"code"`
			Message 	string `json:"message"`
			Data struct {
				Status	int	`json:"status"`
				Url		string	`json:"url"`
			}	`json:"data"`
		}
		finfo := &sixresp
		if err := json.Unmarshal([]byte(string(r.Body)), finfo); err != nil {
			log.Fatal("Line-223_Error:", err)
		}
		stcImg := finfo.Data.Url
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		ctx.Put("keyframeindex", r.Request.Ctx.Get("keyframeindex"))
		gpostinfo := map[string]string {
			"imgSrc": stcImg
			"cutX":"0"
			"cutY":"0"
			"oriWidth":"1770"
			"oriHeight":"1000"
			"saveWidth":"1024"
			"saveHeight":"578"
			"utoken":"1d24fa6122532c469e18cf838ef0838b"
		}
		//构造保存封面的请求地址
		savecoverlink := "https://ns.dayu.com/article/imagecut?
		appid=website&wmid=941878fe5eb1434a8f2afe13a3874c51&
		sign=9338f1f93886fb374ec03654d51581f6"
		g.Request("POST", savecoverlink, gpostinfo, ctx, nil)
	})

	h.OnResponse(func(r *colly.Response) {
		var sevenresp struct {
			code	int	`json:"code"`
			Message	int	`json:"message"`
			Data	struct {
				Status	int	`json:"status"`
				Url		string	`json:"url"`
			}	`json:"data"`
		}
		hinfo := &sevenresp
		if err := json.Unmarshal([]byte(string(r.Body)), hinfo); err != nil {
			log.Fatal("Line-261_Error:", err)
		}
		imgSrc := hinfo.Data.Url
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		ctx.Put("keyframeindex", r.Request.Ctx.Get("keyframeindex"))
		ctx.Put("imgSrc", imgSrc)
		ipostinfo := map[string]string {
			"author":	""
      "title": filename
      "content": ""
      "utoken":"1d24fa6122532c469e18cf838ef0838b"
		}
		i.Request("POST", "https://mp.dayu.com/filterContent", ipostinfo, ctx, nil)
	})
	
	i.OnResponse(func(r *colly.Response) {
		var eightresp struct {
			Result	int	`json:"result"`
		}
		iinfo := &eightresp
		if err := json.Unmarshal([]byte(string(r.Body)), iinfo); err != nil {
			log.Fatal("Line-285_Error:", err)
		}
		if r := iinfo.Result; r != 0 {
			log.Fatal("9.信息提交失败：", r)
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		ctx.Put("keyframeindex", r.Request.Ctx.Get("keyframeindex"))
		ctx.Put("imgSrc", r.Request.Ctx.Get("imgSrc"))
		jpostinfo := map[string]string {
			"keyframe_index": r.Request.Ctx.Get("keyframeindex")    
			"cover_from": 3
			"isCloseAdManual": false
			"article_category": 1
			"draft_id": ""
			"article_id": ""
			"materialId": r.Request.Ctx.Get("materialId")
			"mid": ""
			"videoId": r.Request.Ctx.Get("videoID")
			"videoFileName": filename
			"title": filename
			"sub_title": ""
			"coverImg": r.Request.Ctx.Get("imgSrc")
			"article_type": 1
			"category": "娱乐"
			"aoyun": false
			"weixin_promote": false
			"curDaySubmit": false
			"customize_tags": ["水下生活"]
			"is_show_ad": true
			"is_video_ad_share": false
			"is_original": 0
			"article_activity_title": ""
			"article_activity_id": ""
			"is_timed_release": false
			"time_for_release": 1528650469483
			"vertical_cover_url": ""
			"keyword": ""
			"is_exclusive": false
			"use_vertical_cover": false
			"source": 1
			"originCoverImg": ""
			"utoken":"1d24fa6122532c469e18cf838ef0838b"
		}
		j.Request("POST", "https://mp.dayu.com/dashboard/save-draft", jpostinfo, ctx, nil)
	})

	k.OnResponse(func(r *colly.Response) {
		var nineresp struct {
			Data  struct {
				Createdat string `json:"_created_at"`
				DataListID string `json:"_id"`
				Updatedat	string	`json:"_updated_at"`
				Wmupdateat 	string	`json:"wm_update_at"`
				PreviewUrl	string	`json:"previewUrl"`
				Wmaid	string	`json:"wm_aid"`
				Acl	string	`json:"acl"`
			}	`json:"data"`
		}
		kinfo := &nineresp
		if err := json.Unmarshal([]byte(string(r.Body)), kinfo); err != nil {
			log.Fatal("Line-348_Error:", err)
		}
		dataListID := kinfo.Data.DataListID
		mpostinfo := map[string]string {
			"dataList[0][_id]": dataListID
			"dataList[0][isDraft]":1
			"curDaySubmit":false
			"utoken":"1d24fa6122532c469e18cf838ef0838b"
		}
		err := m.Post("https://mp.dayu.com/dashboard/submit-article", mpostinfo)
		if err != nil {
			log.Fatal("Line-360_Error:", err)
		}
	})
}