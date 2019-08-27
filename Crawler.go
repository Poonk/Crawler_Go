package main

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/astaxie/beego/logs"
)

func download(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	res, err := client.Do(req)
	if err != nil {
		logs.Error(err)
	}
	defer res.Body.Close()

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

	// doc, err := goquery.NewDocument(url)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// logs.Debug(doc)

	// file, err := os.Create("./douban.txt")
	// if err != nil {
	// 	logs.Error(err)
	// }

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logs.Error(err)
	}
	logs.Debug(doc)

	doc.Find(".m_content .m_cont_3 .sub_cont_3 .company_details").Each(func(index int, content *goquery.Selection) {
		logs.Debug(content.Text())

		size := content.Find("dd").Length()
		logs.Debug(size)
		for i := 0; i < size; i++ {
			tm := content.Find("dt").Eq(i).Text()
			logs.Debug(tm)
			tmp := content.Find("dd").Eq(i).Text()
			logs.Debug(tmp)
		}
		// tmp := content.Find("dd").Text()
		// logs.Debug(tmp)
		// tmp = content.Find(".company_details title").Text()
		// logs.Debug(tmp)
		// logs.Debug("ttt")
		// title := content.Find(".jk").Text()
		// open := content.Find(".jk .topenprice").Text()
		// logs.Debug("Movie %d : %s\n", index, title)
		// err := ioutil.WriteFile("douban.txt", []byte(title), 0644)
		// if err != nil {
		// 	logs.Error(err)
		// }
		// file.Write([]byte(title + "\n"))
		// logs.Debug(title + "\t" + open)
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

	url := "http://stockpage.10jqka.com.cn/000001/#gegugp_zjjp"

	download(url)

}
