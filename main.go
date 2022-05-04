package main

// @Title  请填写文件名称（需修改）
// @Description  请填写文件描述（需修改）
// @Author  clx  2022/5/4 9:36 上午
// @Update  clx  2022/5/4 9:36 上午

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tianxinbaiyun/crawler/config"
	"github.com/tianxinbaiyun/crawler/tool"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func GetDetail(detailUrl string, group *sync.WaitGroup) {
	defer func() {
		group.Done()
	}()
	pageSie := config.C.Crawler.ImgCommentPageSize
	maxPage := config.C.Crawler.SearchMaxPage
	urlParams, _ := url.Parse(detailUrl)
	productID := strings.TrimRight(strings.TrimLeft(urlParams.Path, "/"), ".html")

	//fmt.Println(productID)
	rsp, err := tool.GetProductPageImageCommentList(productID, 1, pageSie)
	if err != nil {
		fmt.Println(err)
		return
	}

	if rsp.ImgComments.ImgCommentCount <= 0 || len(rsp.ImgComments.ImgList) <= 0 {
		return
	}
	SaveImgList(rsp.ImgComments.ImgList)

	// 更多页处理
	if rsp.ImgComments.ImgCommentCount > pageSie {
		for i := 2; i <= maxPage && i < rsp.ImgComments.ImgCommentCount/pageSie+1; i++ {
			rsp, err := tool.GetProductPageImageCommentList(productID, i, pageSie)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(rsp.ImgComments.ImgList) <= 0 {
				return
			}
			SaveImgList(rsp.ImgComments.ImgList)
		}
	}
	return
	//fmt.Println(rsp)
}

func SaveImgList(list []*tool.ImgListItem) {
	for _, item := range list {
		if strings.HasPrefix(item.ImageURL, `//`) {
			item.ImageURL = strings.ReplaceAll(item.ImageURL, `//`, "https://")
		}
		fmt.Println(item.ImageURL)
		resp, err := http.Get(item.ImageURL)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
			break
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		urlArr := strings.Split(item.ImageURL, `/`)
		fileName := urlArr[len(urlArr)-1]
		err = ioutil.WriteFile(config.C.DownLoad.Path+fileName, body, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
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
			go GetDetail(url, &wg)

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
	url := tool.GetSearchUrl(config.C.Crawler.Keyword, i)
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
