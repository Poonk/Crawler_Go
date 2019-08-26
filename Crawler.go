package main

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
)

func download(url string) {
	// client := &http.Client{}
	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	// res, err := client.Do(req)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	logs.Debug(err)
	// }
	// logs.Debug(string(body))

	//link
	// links := collectlinks.All(res.Body)
	// for _, link := range links {
	// 	fmt.Println("parse url", link)
	// }

	doc, err := goquery.NewDocument(url)
	if err != nil {
		logs.Error(err)
	}
	logs.Debug(doc)
	doc.Find(".article .grid_view .item .info").Each(func(index int, content *goquery.Selection) {
		title := content.Find(".hd a span").Text()
		logs.Debug("Movie %d : %s\n", index, title)
		time.Sleep(time.Second * 2)
	})
	// doc.Find(".sidebar-reviews article .content-block").Each(func(index int, content *goquery.Selection) {
	// 	band := content.Find("a").Text()
	// 	tilte := content.Find("i").Text()
	// 	logs.Debug("Band %d: %s - %s\n", index, band, tilte)
	// })

	// time.Sleep(time.Second * 5)

}

func main() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	url := "https://movie.douban.com/top250"

	download(url)

}
