package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var uploadUrl = ""
var filename = ""
var now = time.Now().UnixNano() / 1e6

// UploadVideo 视频上传 TODO: 目前还未完善

func UploadVideo() {

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
	i := h.Clone()
	j := i.Clone()

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile("userfile", filename)
	if err != nil {
		log.Println("error writing to buffer")
		return
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		log.Println("error opening file")
		return
	}
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)

	a.Request("POST", uploadUrl, request_reader, nil, nil)

	a.OnResponse(func(r *colly.Response) {
		var oneresp struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				ID string `json:"id"`
			} `json:"data"`
		}
		ainfo := &oneresp
		if err := json.Unmarshal([]byte(string(r.Body)), ainfo); err != nil {
			log.Fatal("Line-73_Error:", err)
			return
		}
		bpostinfo := map[string]string{
			"videoId":   ainfo.Data.ID,
			"videoName": filename,
			"videoDesc": "",
			"appid":     "wemedia",
			"uid":       "941878fe5eb1434a8f2afe13a3874c51 ", //UID和utoken在user信息里
			"utoken":    "1d24fa6122532c469e18cf838ef0838b",
		}
		jsonbpostinfo, err := json.Marshal(bpostinfo)
		if err != nil {
			log.Fatal("Line-101_Error:", err)
			return
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", ainfo.Data.ID)
		b.Request("POST", "https://mp.dayu.com/confirm_info", bytes.NewBuffer(jsonbpostinfo), ctx, nil)
	})

	b.OnResponse(func(r *colly.Response) {
		var tworesp struct {
			Succeed int `json:"succeed"`
		}
		binfo := &tworesp
		if err := json.Unmarshal([]byte(string(r.Body)), binfo); err != nil {
			log.Fatal("Line-101_Error:", err)
			return
		}
		if bresult := tworesp.Succeed; bresult != 1 {
			log.Fatal("第二步上传失败:", bresult)
			return
		}
		vid := r.Request.Ctx.Get("videoID")
		cpostinfo := map[string]string{
			"type":            "video",
			"title":           filename,
			"video[video_id]": vid,
			"utoken":          "1d24fa6122532c469e18cf838ef0838b",
		}
		jsoncpostinfo, err := json.Marshal(cpostinfo)
		if err != nil {
			log.Fatal("Line-131_Error:", err)
			return
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		c.Request("POST", "https://mp.dayu.com/dashboard/addMaterial", bytes.NewBuffer(jsoncpostinfo), ctx, nil)
	})

	c.OnResponse(func(r *colly.Response) {
		var threeresp struct {
			Data struct {
				Createdat  string `json:"_created_at"`
				MaterialId string `json:"_id"`
				Updatedat  string `json:"_updated_at"`
			} `json:"data"`
		}
		cinfo := &threeresp
		if err := json.Unmarshal([]byte(string(r.Body)), cinfo); err != nil {
			log.Fatal("Line-128_Error:", err)
			return
		}
		materialId := cinfo.Data.MaterialId
		dpostinfo := map[string]string{
			"id":     r.Request.Ctx.Get("videoID"),
			"utoken": "1d24fa6122532c469e18cf838ef0838b",
		}
		jsondpostinfo, err := json.Marshal(dpostinfo)
		if err != nil {
			log.Fatal("Line-159_Error:", err)
			return
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", materialId)
		d.Request("POST", "https://mp.dayu.com/dashboard/getVideoPlayUrl", bytes.NewBuffer(jsondpostinfo), ctx, nil)
	})

	d.OnResponse(func(r *colly.Response) {
		var fourresp struct {
			PlayUrl string `json:"playUrl"`
		}
		dinfo := &fourresp
		if err := json.Unmarshal([]byte(string(r.Body)), dinfo); err != nil {
			log.Fatal("Line-128_Error:", err)
			return
		}
		playurl := dinfo.PlayUrl
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", playurl)
		vid := r.Request.Ctx.Get("videoID")
		erequrl := fmt.Sprintf("https://mp.dayu.com/media_keyframe?id=%s&_=%s", vid, strconv.FormatInt(now, 10))
		epostinfo := map[string]interface{}{
			"id": vid,
			"_":  now,
		}
		jsonepostinfo, err := json.Marshal(epostinfo)
		if err != nil {
			log.Fatal("Line-159_Error:", err)
			return
		}
		e.Request("POST", erequrl, bytes.NewBuffer(jsonepostinfo), ctx, nil)
	})

	e.OnResponse(func(r *colly.Response) {
		var fiveresp struct {
			Succeed int    `json:"succeed"`
			Api     string `json:"api"`
			Data    struct {
				ID             string `json:"id"`
				Reqcount       int    `json:"req_count"`
				Respcount      int    `json:"resp_count"`
				Totalcount     int    `json:"total_count"`
				Videokeyframes struct {
					Frames []struct {
						Imgurl             string `json:"img_url"`
						Width              int    `json:"width"`
						Height             int    `josn:"height"`
						BlackBorderPercent int    `json:"blackBorderPercent"`
						BlackBorderLevel   string `josn:"blackBorderLevel"`
						BlurMark           bool   `josn:"blurMark"`
						BlackBorderMark    bool   `json:"blackBorderMark"`
						Keyframeindex      int    `json:"keyframe_index"`
					} `json:"frames"`
				} `json:"video_keyframes"`
			} `json:"data"`
		}
		einfo := &fiveresp
		if err := json.Unmarshal([]byte(string(r.Body)), einfo); err != nil {
			log.Fatal("Line-185_Error:", err)
			return
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		//默认选择第五张图为视频封面
		keyframeindex := einfo.Data.Videokeyframes.Frames[4].Keyframeindex
		imgSrc := einfo.Data.Videokeyframes.Frames[4].Imgurl
		ctx.Put("keyframeindex", keyframeindex)
		fpostinfo := map[string]string{
			"imgSrc": imgSrc,
			"utoken": "1d24fa6122532c469e18cf838ef0838b",
		}
		jsonfpostinfo, err := json.Marshal(fpostinfo)
		if err != nil {
			log.Fatal("Line-238_Error:", err)
			return
		}
		//TODO: 构造封面选择的地址
		imgpostlink := "https://ns.dayu.com/article/outsiteImgResave?appid=website&wmid=xxxxx&sign=xxxxx"
		f.Request("POST", imgpostlink, bytes.NewBuffer(jsonfpostinfo), ctx, nil)
	})

	f.OnResponse(func(r *colly.Response) {
		var sixresp struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Status int    `json:"status"`
				Url    string `json:"url"`
			} `json:"data"`
		}
		finfo := &sixresp
		if err := json.Unmarshal([]byte(string(r.Body)), finfo); err != nil {
			log.Fatal("Line-223_Error:", err)
			return
		}
		stcImg := finfo.Data.Url
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		ctx.Put("keyframeindex", r.Request.Ctx.Get("keyframeindex"))
		gpostinfo := map[string]interface{}{
			"imgSrc":     stcImg,
			"cutX":       0,
			"cutY":       0,
			"oriWidth":   1770,
			"oriHeight":  1000,
			"saveWidth":  1024,
			"saveHeight": 578,
			"utoken":     "1d24fa6122532c469e18cf838ef0838b",
		}
		jsongpostinfo, err := json.Marshal(gpostinfo)
		if err != nil {
			log.Fatal("Line-238_Error:", err)
			return
		}
		//构造保存封面的请求地址
		savecoverlink := "https://ns.dayu.com/article/imagecut?appid=website&wmid=xxxxx&sign=xxxxx"
		g.Request("POST", savecoverlink, bytes.NewBuffer(jsongpostinfo), ctx, nil)
	})

	g.OnResponse(func(r *colly.Response) {
		var sevenresp struct {
			code    int `json:"code"`
			Message int `json:"message"`
			Data    struct {
				Status int    `json:"status"`
				Url    string `json:"url"`
			} `json:"data"`
		}
		hinfo := &sevenresp
		if err := json.Unmarshal([]byte(string(r.Body)), hinfo); err != nil {
			log.Fatal("Line-261_Error:", err)
			return
		}
		imgSrc := hinfo.Data.Url
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		ctx.Put("keyframeindex", r.Request.Ctx.Get("keyframeindex"))
		ctx.Put("imgSrc", imgSrc)
		hpostinfo := map[string]string{
			"author":  "",
			"title":   filename,
			"content": "",
			"utoken":  "1d24fa6122532c469e18cf838ef0838b",
		}
		jsonhpostinfo, err := json.Marshal(hpostinfo)
		if err != nil {
			log.Fatal("Line-315_Error:", err)
			return
		}
		h.Request("POST", "https://mp.dayu.com/filterContent", bytes.NewBuffer(jsonhpostinfo), ctx, nil)
	})

	h.OnResponse(func(r *colly.Response) {
		var eightresp struct {
			Result int `json:"result"`
		}
		iinfo := &eightresp
		if err := json.Unmarshal([]byte(string(r.Body)), iinfo); err != nil {
			log.Fatal("Line-285_Error:", err)
			return
		}
		if r := iinfo.Result; r != 0 {
			log.Fatal("9.信息提交失败：", r)
			return
		}
		ctx := colly.NewContext()
		ctx.Put("videoID", r.Request.Ctx.Get("videoID"))
		ctx.Put("materialId", r.Request.Ctx.Get("materialId"))
		ctx.Put("playUrl", r.Request.Ctx.Get("playUrl"))
		ctx.Put("keyframeindex", r.Request.Ctx.Get("keyframeindex"))
		ctx.Put("imgSrc", r.Request.Ctx.Get("imgSrc"))
		ipostinfo := map[string]interface{}{
			"keyframe_index":         r.Request.Ctx.Get("keyframeindex"),
			"cover_from":             3,
			"isCloseAdManual":        false,
			"article_category":       1,
			"draft_id":               "",
			"article_id":             "",
			"materialId":             r.Request.Ctx.Get("materialId"),
			"mid":                    "",
			"videoId":                r.Request.Ctx.Get("videoID"),
			"videoFileName":          filename,
			"title":                  filename,
			"sub_title":              "",
			"coverImg":               r.Request.Ctx.Get("imgSrc"),
			"article_type":           1,
			"category":               "娱乐",
			"aoyun":                  false,
			"weixin_promote":         false,
			"curDaySubmit":           false,
			"customize_tags":         `["水下生活"]`,
			"is_show_ad":             true,
			"is_video_ad_share":      false,
			"is_original":            0,
			"article_activity_title": "",
			"article_activity_id":    "",
			"is_timed_release":       false,
			"time_for_release":       1528650469483,
			"vertical_cover_url":     "",
			"keyword":                "",
			"is_exclusive":           false,
			"use_vertical_cover":     false,
			"source":                 1,
			"originCoverImg":         "",
			"utoken":                 "1d24fa6122532c469e18cf838ef0838b",
		}
		jsonipostinfo, err := json.Marshal(ipostinfo)
		if err != nil {
			log.Fatal("Line-315_Error:", err)
			return
		}
		i.Request("POST", "https://mp.dayu.com/dashboard/save-draft", bytes.NewBuffer(jsonipostinfo), ctx, nil)
	})

	i.OnResponse(func(r *colly.Response) {
		var nineresp struct {
			Data struct {
				Createdat  string `json:"_created_at"`
				DataListID string `json:"_id"`
				Updatedat  string `json:"_updated_at"`
				Wmupdateat string `json:"wm_update_at"`
				PreviewUrl string `json:"previewUrl"`
				Wmaid      string `json:"wm_aid"`
				Acl        string `json:"acl"`
			} `json:"data"`
		}
		kinfo := &nineresp
		if err := json.Unmarshal([]byte(string(r.Body)), kinfo); err != nil {
			log.Fatal("Line-395_Error:", err)
			return
		}
		dataListID := kinfo.Data.DataListID
		jpostinfo := map[string]interface{}{
			"dataList[0][_id]":     dataListID,
			"dataList[0][isDraft]": 1,
			"curDaySubmit":         false,
			"utoken":               "1d24fa6122532c469e18cf838ef0838b",
		}
		jsonjpostinfo, err := json.Marshal(jpostinfo)
		if err != nil {
			log.Fatal("Line-407_Error:", err)
			return
		}
		j.Request("POST", "https://mp.dayu.com/dashboard/submit-article", bytes.NewBuffer(jsonjpostinfo), nil, nil)
	})

}
