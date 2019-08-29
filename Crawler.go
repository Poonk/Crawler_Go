package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/astaxie/beego/logs"
)

func download(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Xsrftoken", "b6d695bbdcd111e8b681002324e63af81")
	req.Header.Add("Cookie", "sessionid=5edb1f18c5a0cb334b42b2383c899e01")
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

	// name := url[32:37]
	// logs.Debug(name)
	// f, err := os.Create(name + ".csv")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// defer f.Close()

	// f.WriteString("\xEF\xBB\xBF")
	// w := csv.NewWriter(f)

	// data := &protocol.Basicdata{}

	doc, err := goquery.NewDocumentFromResponse(res)
	// doc, err := goquery.NewDocumentFromReader(strings.NewReader(url))
	if err != nil {
		logs.Error(err)
	}
	logs.Debug(doc.Text())

	// doc.Find("body div>.in_squote").Each(func(index int, content *goquery.Selection) {
	// 	logs.Debug(content.Text())
	// })

	// doc.Find(".m_content .m_cont_3 .sub_cont_3 .company_details").Each(func(index int, content *goquery.Selection) {
	// 	// logs.Debug(content.Text())

	// 	size := content.Find("dd").Length()
	// 	logs.Debug(size)
	// 	for i := 0; i < size; i++ {
	// 		tm := content.Find("dt").Eq(i).Text()
	// 		logs.Debug(tm)
	// 		logs.Debug(utf8.RuneCountInString(tm))
	// 		tmp := content.Find("dd").Eq(i).Text()
	// 		logs.Debug(tmp)

	// 		len := utf8.RuneCountInString(tmp)
	// 		if tm == "总股本：" {
	// 			// logs.Debug(tmp[:len-1])
	// 			num, _ := strconv.ParseFloat(tmp[:len-1], 64)
	// 			// logs.Debug(num)
	// 			data.CapitalAmount = num
	// 		} else if tm == "流通股：" {
	// 			num, _ := strconv.ParseFloat(tmp[:len-1], 64)
	// 			data.FloatingStocks = num
	// 		} else if tm == "每股收益：" {
	// 			num, _ := strconv.ParseFloat(tmp[:len-1], 64)
	// 			data.EPS = num
	// 		}
	// 	}
	// 	logs.Debug(data)

	// 	// tmp := content.Find("dd").Text()
	// 	// logs.Debug(tmp)
	// 	// tmp = content.Find(".company_details title").Text()
	// 	// logs.Debug(tmp)
	// 	// logs.Debug("ttt")
	// 	// title := content.Find(".jk").Text()
	// 	// open := content.Find(".jk .topenprice").Text()
	// 	// logs.Debug("Movie %d : %s\n", index, title)
	// 	// err := ioutil.WriteFile("douban.txt", []byte(title), 0644)
	// 	// if err != nil {
	// 	// 	logs.Error(err)
	// 	// }
	// 	// file.Write([]byte(title + "\n"))
	// 	// logs.Debug(title + "\t" + open)
	// 	time.Sleep(time.Second * 2)
	// })

	// doc.Find(".sidebar-reviews article .content-block").Each(func(index int, content *goquery.Selection) {
	// 	band := content.Find("a").Text()
	// 	tilte := content.Find("i").Text()
	// 	logs.Debug("Band %d: %s - %s\n", index, band, tilte)
	// })

	// time.Sleep(time.Second * 5)

}

// func storeToCsv(filename string, posts map[int]*Post) {
// 	// 创建文件
// 	csvFile, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer csvFile.Close()

// 	// 获取csv的Writer
// 	writer := csv.NewWriter(csvFile)

// 	// 将map中的Post转换成slice，因为csv的Write需要slice参数
// 	// 并写入csv文件
// 	for _, post := range posts {
// 		record := []string{strconv.Itoa(post.Id), post.Content, post.Author}
// 		err1 := writer.Write(record)
// 		if err1 != nil {
// 			panic(err1)
// 		}
// 	}

// 	// 确保所有内存数据刷到csv文件
// 	writer.Flush()
// }

func main() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	// url := "http://stockpage.10jqka.com.cn/000001/#gegugp_zjjp"
	// url := "http://stockpage.10jqka.com.cn/realHead_v2.html#hs_000001"
	// url := "http://stockpage.10jqka.com.cn/realHead_v2.html"
	url := "http://d.10jqka.com.cn/v2/realhead/hs_000001/"

	download(url)

}
