package main

import (
	
	"encoding/json"

	"github.com/asticode/go-astilog"
	"github.com/ZEROKISEKI/go-astilectron-bootstrap"
	"github.com/asticode/go-astilectron"
)

// UploadInfo 视频上传信息接口   
type UploadInfo []struct {
	Title    string `json:"title"`
	Describe string `json:"describe"`
	Tags     string `json:"tags"`
	Field    string `json:"field"`
}

// IDpost为IDspider 获取前台提交的数据的结构
type IDpost struct {
	Platform	string `json:"platform"`
	AuthorID	string	`json:"authorid"`
	Readlimit string	`json:"readlimit"`
	Datefrom	string	`json:"datefrom"`
	Dateto		string	`json:"dateto"`
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
	case "Baowen":
		{
			var url string
			err := json.Unmarshal(m.Payload, &url)
			if err != nil {
				return err.Error(), nil
			}

			_, _ = a.NewWindow(url, &astilectron.WindowOptions{
				Center: astilectron.PtrBool(true),
				Height: astilectron.PtrInt(600),
        Width:  astilectron.PtrInt(600),
			})
		}
	}
	return
}
