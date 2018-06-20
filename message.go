package main

import (
	"encoding/json"
	"log"

	"github.com/ZEROKISEKI/go-astilectron-bootstrap"
	"github.com/asticode/go-astilectron"
)

// handleMessages handles messages
func handleMessages(a *astilectron.Astilectron, _ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "Upload":
		{
			// UploadInfo 视频上传信息接口
			var UploadInfo struct {
				Title    string `json:"title"`
				Describe string `json:"describe"`
				Tags     string `json:"tags"`
				Field    string `json:"field"`
			}

			err := json.Unmarshal(m.Payload, &UploadInfo)
			if err != nil {
				return err.Error(), nil
			}
			log.Println(UploadInfo)
		}
	}
	return
}
