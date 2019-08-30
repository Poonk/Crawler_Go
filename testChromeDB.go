package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	// var buf []byte

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		//访问打开必应页面
		chromedp.Navigate(`http://stockpage.10jqka.com.cn/000001/#gegugp_zjjp`),
		// 等待右下角图标加载完成
		// chromedp.WaitVisible(`#sh_cp_in`),
		//搜索框内输入zhangguanzhang
		// chromedp.SendKeys(`#sb_form_q`, `zhangguanzhang`, chromedp.ByID),
		// 点击搜索图标
		// chromedp.Click(`#sb_form_go`, chromedp.NodeVisible),
		// 获取第一个搜索结构的超链接
		chromedp.Text(`#in_squote>.new_trading fl>li`, &example),
		// chromedp.CaptureScreenshot(&buf),
		// chromedp.OuterHTML("#in_squote",example,chromedp.ByID)
	)
	if err != nil {
		log.Fatal(err)
	}

	// if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
	// 	log.Fatal(err)
	// }
	log.Printf("example: %s", example)

}
