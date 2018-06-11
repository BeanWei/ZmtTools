package main

import (
	"log"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	w, err := window.New(
		sciter.SW_TITLEBAR|
			sciter.SW_RESIZEABLE|
			sciter.SW_CONTROLS|
			sciter.SW_MAIN|
			sciter.SW_ENABLE_DEBUG,
		nil)
	if err != nil {
		log.Fatal("Line_20-Error:", err)
	}

	w.LoadFile("./static/html/index.html")
	//w.LoadFile("./static/html/upload.html")
	w.SetTitle("登录")
	w.Show()
	w.Run()
}
