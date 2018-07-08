package main

import (
	"encoding/json"

	"github.com/ZEROKISEKI/go-astilectron-bootstrap"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilog"
)

// UploadInfo 视频上传信息接口
type UploadInfo []struct {
	Filename	string	`json:"filename"`
	Title    string `json:"title"`
	Describe string `json:"describe"`
	Tags     string `json:"tags"`
	Field    string `json:"field"`
}

// IDpost为IDspider 获取前台提交的数据的结构
type IDpost struct {
	Platform  string `json:"platform"`
	AuthorID  string `json:"authorid"`
	Readlimit string `json:"readlimit"`
	Datefrom  string `json:"datefrom"`
	Dateto    string `json:"dateto"`
}

// CkReq 原创检测的post数据结构
type CkReq struct {
	Title		string	`json:"title"`
	Content	string	`json:"content"`
}

// handleMessages handles messages
func handleMessages(a *astilectron.Astilectron, _ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "Upload":
		{
			uploadinfo := &UploadInfo{}
			err := json.Unmarshal(m.Payload, uploadinfo)
			if err != nil {
				return err.Error(), nil
			}
			//TODO: 添加视频上传模块
			astilog.Debug(uploadinfo)
			return "I HAVE RECIVEID THIS MESSAGE", nil
		}
	case "IDspider":
		{
			idpost := &IDpost{}
			err := json.Unmarshal(m.Payload, idpost)
			if err != nil {
				return err.Error(), nil
			}
			payload = IDspider(idpost)
			return payload, nil
		}
	case "Urlwindow":
		{
			var url string
			err := json.Unmarshal(m.Payload, &url)
			if err != nil {
				return err.Error(), nil
			}

			var w, _ = a.NewWindow(url, &astilectron.WindowOptions{
				Center: astilectron.PtrBool(true),
				Height: astilectron.PtrInt(600),
				Width:  astilectron.PtrInt(600),
			})
			w.Create()
		}
	case "News":
		{
			var field string
			err := json.Unmarshal(m.Payload, &field)
			if err != nil {
				return err.Error(), nil
			}
			err = storage.Init()
			if err != nil {
				astilog.Debugf("DB init Failed")
			}
			//defer storage.Close()
			payload, err = storage.Newsquery(field)
			if err != nil {
				return err.Error(), nil
			}
			return payload, nil
		}
	case "Domains":
		{
			err := storage.Init()
			if err != nil {
				astilog.Debugf("DB init Failed")
			}
			//defer storage.Close()
			payload, err = storage.Domains()
			if err != nil {
				return err.Error(), nil
			}
			return payload, nil
		}
	case "ClearOldNews":
		{
			// err := storage.Init()
			// if err != nil {
			// 	astilog.Debugf("DB init Failed")
			// }
			// //defer storage.Close()
			// err = storage.ClearOldNews()
			// if err != nil {
			// 	astilog.Error(err)
			// 	return map[string]string{"infotype": "error", "alertInfo": "清理历史新闻失败"}, nil
			// }
			// return map[string]string{"infotype": "success", "alertInfo": "已为您成功清理历史新闻"}, nil
			return map[string]string{"infotype": "info", "alertInfo": "抱歉！暂不支持清理"}, nil
		}
	case "NewsSpider":
		{
			//NewsSpider()
			go func() {
				tencent()
				zaker()
				jxnews()
				ynet()
			}()
			return map[string]string{"infotype": "info", "alertInfo": "正在为您爬取更多最新新闻"}, nil	
		}
	case "OriginalCk": 
		{
			CkData := &CkReq{}
			err := json.Unmarshal(m.Payload, CkData)
			if err != nil {
				return err.Error(), nil
			}
			title := CkData.Title
			content := CkData.Content
			resultStr := originalCk(title, content)
			return resultStr, nil
		}
	}	
	return
}
