package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// dateparse 时间解析
func dateparse(datestr string) (string, error) {

	// var eg = []string{
	// 	"1分钟前",
	// 	"54分钟前",
	// 	"1小时前",
	// 	"13小时前",
	// 	"昨天",
	// 	"前天",
	// 	"06-25",
	// 	"12-20"}

	// func main() {
	// 	for _, i := range eg {
	// 		t, err := dateparse(i)
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		log.Println(t)
	// 	}
	// }

	if strings.Contains(datestr, "分钟前") {
		min, err := strconv.Atoi(strings.Split(datestr, "分钟前")[0])
		if err != nil {
			return datestr, err
		}

		return time.Now().Add(-time.Minute * time.Duration(min)).Format("2006-01-02 15:04:05"), nil

	} else if strings.Contains(datestr, "小时前") {
		hour, err := strconv.Atoi(strings.Split(datestr, "小时前")[0])
		if err != nil {
			return datestr, err
		}

		return time.Now().Add(-time.Hour * time.Duration(hour)).Format("2006-01-02 15:04:05"), nil

	} else if datestr == "昨天" {
		return time.Now().Add(-time.Hour * 24).Format("2006-01-02 15:04:05"), nil

	} else if datestr == "前天" {
		return time.Now().Add(-time.Hour * 24 * 2).Format("2006-01-02 15:04:05"), nil

	} else if strings.Contains(datestr, "-") && len(datestr) == 5 {
		return fmt.Sprintf("%d-%s 00:00:00", time.Now().Year(), datestr), nil

	} else {
		//panic(fmt.Sprintf("不支持此时间格式的解析=> %s", datestr))
		return datestr, nil
	}
}
