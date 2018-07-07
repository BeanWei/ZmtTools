package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ZEROKISEKI/go-astilectron-bootstrap"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
	// "github.com/jasonlvhit/gocron"
)

var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", true, "enables the debug mode")
	w       *astilectron.Window
)

//初始化数据库
var storage = &Storage{Filename: "./src.db"}

func main() {

	flag.Parse()
	astilog.FlagInit()

	// defer storage.Close()
	// err := storage.Init()
	// if err != nil {
	// 	astilog.Debugf("DB init Failed")
	// }

	if p, err := os.Executable(); err != nil {
		err = errors.Wrap(err, "os.Executable failed")
		return
	} else {
		p = filepath.Dir(p)
		log.Print(p)
	}

	astilog.Debugf("Running app built at %s", BuiltAt)

	if err := bootstrap.Run(bootstrap.Options{
		Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "ui/dist/icon.icns",
			AppIconDefaultPath: "ui/dist/icon.png",
		},
		Debug:         *debug,
		ResourcesPath: "ui/dist",
		Homepage:      "index.html",
		OnWait: func(_ *astilectron.Astilectron, iw *astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = iw
			err := w.OpenDevTools()
			if err != nil {
				astilog.Fatal(errors.Wrap(err, "OpenDevTools failed"))
			}

			//软件开启时开始新闻爬取

			//NewsSpider()
			// go func() {
			// 	tencent()
			// 	zaker()
			// 	jxnews()
			// 	ynet()
			// }()

			//将item传给前台vuex数据管理处
			go func() {
				time.Sleep(5 * time.Second)
				err := storage.Init()
				if err != nil {
					astilog.Debugf("DB init Failed")
				}
				//defer storage.Close()
				payload, err := storage.Domains()
				if err != nil {
					astilog.Error(errors.Wrap(err, "DB get the news's all domains failed"))
				} else {
					if err = bootstrap.SendMessage(w, "newsdomain", payload); err != nil {
						astilog.Error(errors.Wrap(err, "send the news's all domains failed"))
					}	
				}	
			}()

			//定时爬取新闻
			// gocron.NewScheduler().Every(20).Minutes().Do(NewsSpider)
			// <-gocron.Start()
			//go checkAPI()
			return nil
		},
		MessageHandler: handleMessages,
		RestoreAssets:  RestoreAssets,
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#fff"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(760),
			Width:           astilectron.PtrInt(1350),
			MinWidth:        astilectron.PtrInt(1050),
			AutoHideMenuBar:	 astilectron.PtrBool(true),
		},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}

}
