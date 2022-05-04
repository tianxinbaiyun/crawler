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

var (
	Cookie    = "__yjs_duid=1_d683a02f304244394b8b744db2beda391651629234114; yjs_js_security_passport=2a88ad0a5a361deee7a00574d0f9df2b24c86509_1651629236_js; zkhanecookieclassrecord=%2C54%2C; Hm_lvt_c59f2e992a863c2744e1ba985abaea6c=1651629237; Hm_lpvt_c59f2e992a863c2744e1ba985abaea6c=1651629632"
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"
)

func down(i int, group *sync.WaitGroup) {
	//func down(i int) {
	var c = colly.NewCollector()
	//c.OnRequest(func(r *colly.Request) {
	//	//r.Headers.Set("cookie", config.C.Header.Cookie)
	//	//r.Headers.Set("user-agent", UserAgent)
	//
	//})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Println("返回不是200")
			return
		}
	})
	c.OnHTML(config.C.Crawler.Selector, func(e *colly.HTMLElement) {

		//url := "https://pic.netbian.com" + e.Attr("src")
		url := e.Attr("src")
		fmt.Println(e.Attr("src"))
		var c = colly.NewCollector()
		//c.OnRequest(func(r *colly.Request) {
		//	r.Headers.Set("cookie", Cookie)
		//	r.Headers.Set("user-agent", UserAgent)
		//
		//})
		c.OnResponse(func(r *colly.Response) {
			reader := bytes.NewReader(r.Body)
			body, _ := ioutil.ReadAll(reader)
			//读取图片内容
			index := strings.LastIndex(url, "/")
			s := url[index+1:]
			err := ioutil.WriteFile("./img/"+s, body, 0755)
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
	url := GetSearchUrl(config.C.Crawler.Keyword, i)
	c.Visit(url)

}

func main() {
	// 读取配置文件到struct,初始化变量
	config.InitConfig()

	group := sync.WaitGroup{}
	for i := 1; i < 100; i++ {
		group.Add(1)
		go down(i, &group)
	}
	group.Wait()
	fmt.Println("下载完整")
}
