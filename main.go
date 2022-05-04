package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tianxinbaiyun/crawler/config"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func down(detailUrl string, group *sync.WaitGroup) {

	var c = colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", config.C.Header.Cookie)
		r.Headers.Set("user-agent", config.C.Header.UserAgent)
		r.Headers.Set("Host", config.C.Header.Host)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnHTML(config.C.Crawler.DetailImgSelector, func(e *colly.HTMLElement) {

		url := e.Attr(config.C.Crawler.DetailImgSelectorAttr)
		if strings.HasPrefix(url, `//`) {
			url = strings.ReplaceAll(url, `//`, "https://")
		}
		fmt.Println(url)
		var c = colly.NewCollector()
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("cookie", config.C.Header.Cookie)
			r.Headers.Set("user-agent", config.C.Header.UserAgent)
			r.Headers.Set("Host", config.C.Header.Host)

		})
		c.OnResponse(func(r *colly.Response) {
			reader := bytes.NewReader(r.Body)
			body, _ := ioutil.ReadAll(reader)
			//读取图片内容
			index := strings.LastIndex(url, "/")
			s := url[index+1:]
			err := ioutil.WriteFile(config.C.DownLoad.Path+s, body, 0755)
			fmt.Println(url, err)
		})
		c.Visit(url)

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnScraped(func(r *colly.Response) {
		group.Done()
	})
	c.Visit(detailUrl)

}

// GetPageList 列表页
func GetPageList(i int, group *sync.WaitGroup) {
	wg := sync.WaitGroup{}
	var c = colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", config.C.Header.Cookie)
		r.Headers.Set("user-agent", config.C.Header.UserAgent)
		r.Headers.Set("Host", config.C.Header.Host)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnHTML(config.C.Crawler.ListItemSelector, func(e *colly.HTMLElement) {

		url := e.Attr(config.C.Crawler.ListItemSelectorUrl)
		if strings.HasPrefix(url, `//`) {
			url = strings.ReplaceAll(url, `//`, "https://")
		}
		var c = colly.NewCollector()
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("cookie", config.C.Header.Cookie)
			r.Headers.Set("user-agent", config.C.Header.UserAgent)
			r.Headers.Set("Host", config.C.Header.Host)

		})
		c.OnResponse(func(r *colly.Response) {

			wg.Add(1)
			go down(url, &wg)

			//err := ioutil.WriteFile(config.C.DownLoad.Path+s, body, 0755)
			//fmt.Println(url, err)
		})
		c.Visit(url)

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnScraped(func(r *colly.Response) {
		wg.Wait()
		group.Done()
	})
	url := GetSearchUrl(config.C.Crawler.Keyword, i)
	c.Visit(url)

}

func main() {
	// 读取配置文件到struct,初始化变量
	config.InitConfig()

	group := sync.WaitGroup{}
	for i := 0; i <= config.C.Crawler.SearchMaxPage; i++ {
		group.Add(1)
		go GetPageList(i, &group)
	}
	group.Wait()
	fmt.Println("下载完整")
}
