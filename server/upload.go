package main

import (
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

	err := a.Post(uploadUrl, map[string]byte{"data": request_reader})
	if err != nil {
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
		b.Request("POST", "https://mp.dayu.com/confirm_info", bpostinfo, CTX, nil)
	})

	b.OnResponse(func(r *colly.Response) {
		var tworesp struct {
			
		}
	})


}