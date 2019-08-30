package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

type movie struct {
	Directors []string `json:"directors"`
	Rate      string   `json:"rate"`
	Rover_x   int      `json:"cover_x"`
	Star      string   `json:"star"`
	Title     string   `json:"title"`
	Url       string   `json:"url"`
	Casts     []string `json:"casts"`
	Cover     string   `json:"cover"`
	Id        string   `json:"id"`
	Cover_y   int      `json:"cover_y"`
}

type data struct {
	Data []movie `json:"data"`
}

var (
	file *xlsx.File
	//爬到的影片总数
	number   int
	filePath string = "C:\\Users\\Administrator\\Desktop\\豆瓣.xlsx"

	//影视评分
	Score = flag.Float64("score", 0, "score")
	//分页
	Page = flag.Int("page", 0, "page")
	//影视类型
	Type = flag.String("type", "电影", "type")
)

func getAndParseHtml(movieUrl string) error {

	doc, err := goquery.NewDocument(movieUrl)
	if err != nil {
		return err
	}

	sheet := file.Sheet["Sheet1"]
	row := sheet.AddRow()

	year := row.AddCell()
	doc.Find(".year").Each(func(i int, s *goquery.Selection) {
		year.Value = s.Text()
	})

	name := row.AddCell()
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		v, ok := s.Attr("property")
		if ok {
			if strings.Contains(v, "v:itemreviewed") {
				name.Value = s.Text()
				number++
				glog.Infoln(number, name.Value)
			}
		}
	})

	//顺序不能乱
	director := row.AddCell()
	screenwriter := row.AddCell()
	star := row.AddCell()
	movietype := row.AddCell()
	country := row.AddCell()
	language := row.AddCell()
	releasedate := row.AddCell()
	alias := row.AddCell()
	summary := row.AddCell()

	info := doc.Find("#info")
	for _, v := range strings.Split(info.Text(), "\n") {
		v = strings.TrimSpace(v)
		if strings.Contains(v, "导演:") {
			director.Value = v
		}
		if strings.Contains(v, "编剧:") {
			screenwriter.Value = v
		}
		if strings.Contains(v, "主演:") {
			star.Value = v
		}
		if strings.Contains(v, "类型:") {
			movietype.Value = v
		}
		if strings.Contains(v, "制片国家/地区:") {
			country.Value = v
		}
		if strings.Contains(v, "语言:") {
			language.Value = v
		}
		if strings.Contains(v, "上映日期:") {
			releasedate.Value = v
		}
		if strings.Contains(v, "又名:") {
			alias.Value = v
		}
	}

	//剧情简介
	doc.Find("#link-report").Find("span").Each(func(i int, s *goquery.Selection) {
		_, ok := s.Attr("property")
		if ok {
			summary.Value = strings.Replace(s.Text(), " ", "", -1)
		}
	})

	return nil
}

func spider(pageUrl string) (end bool, err error) {

	cli := &http.Client{}
	req, _ := http.NewRequest("GET", pageUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0")
	glog.Infoln(pageUrl)

	res, e := cli.Do(req)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()

	if e != nil {
		return false, e
	}

	if res.StatusCode == 403 {
		return false, fmt.Errorf("403 Fobidden")
	}

	var d data
	content, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(content, &d)

	//每发一个请求，暂停一段时间
	sleep()

	for _, mv := range d.Data {
		movieUrl := mv.Url
		e := getAndParseHtml(movieUrl)
		if e != nil {
			return false, e
		}

		//每发一个请求，暂停一段时间
		sleep()
	}

	//总数小于20，表示该评分的影片已经爬取完毕
	if len(d.Data) < 20 {
		return true, nil
	}

	return false, nil
}

func parse() {

	for score := *Score; score <= 10; {

		for page := *Page; page < 9980; {

			pageUrl := fmt.Sprintf("https://movie.douban.com/j/new_search_subjects?sort=R&range=%.1f,%.1f&tags=%s&start=%v", score, score, *Type, page)

			end, e := spider(pageUrl)

			//每解析完一页保存一次,每页20个影片
			err := file.Save(filePath)
			if err != nil {
				glog.Errorln(e.Error())
				return
			}

			if e != nil {
				glog.Errorln(e.Error())
				//出错的话通常是被封了，封禁时间大约为3个小时
				time.Sleep(time.Second * 3600)
				continue
			}
			if end {
				break
			}

			page += 20
		}
		*Page = 0
		score += 0.1
	}
}

func main() {

	var e error
	file, e = xlsx.OpenFile(filePath)
	if e != nil {
		file = xlsx.NewFile()
		_, e = file.AddSheet("Sheet1")
		if e != nil {
			glog.Errorln(e.Error())
		}
	}

	flag.Parse()
	parse()

	err := file.Save(filePath)
	if err != nil {
		glog.Errorln(err.Error())
	}

	glog.Flush()
	time.Sleep(time.Second * 5)
}

func sleep() {
	pause := rand.Intn(8) + 3
	time.Sleep(time.Second * time.Duration(pause))

	//五十分之一的几率暂停一段时间
	if rand.Intn(50) == 10 {
		err := file.Save(filePath)
		if err != nil {
			glog.Errorln(err.Error())
		}
		time.Sleep(time.Second * time.Duration(200+pause*5))
	}
}
