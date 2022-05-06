package main

// @Title  请填写文件名称（需修改）
// @Description  请填写文件描述（需修改）
// @Author  clx  2022/5/4 9:36 上午
// @Update  clx  2022/5/4 9:36 上午

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tianxinbaiyun/crawler/config"
	"github.com/tianxinbaiyun/crawler/excel"
	"github.com/tianxinbaiyun/crawler/tool"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// GetDetail 获取详情
func GetDetail(detail *tool.ProductDetail, group *sync.WaitGroup) (exportList []*excel.ExportData) {
	defer func() {
		group.Done()
	}()
	exportList = make([]*excel.ExportData, 0)
	pageSie := config.C.Crawler.ImgCommentPageSize
	maxPage := config.C.Crawler.SearchMaxPage

	// 请求评论接口
	rsp, err := tool.GetProductPageImageCommentList(detail.ProductID, 1, pageSie)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 保存图片
	imgPath := tool.GetImagePath(config.C.DownLoad.Path, config.C.DownLoad.ImagesName) + "/"
	if rsp.ImgComments.ImgCommentCount <= 0 || len(rsp.ImgComments.ImgList) <= 0 {
		return
	}
	list := SaveImgList(detail, rsp.ImgComments.ImgList, imgPath)
	exportList = append(exportList, list...)

	// 更多页处理
	if rsp.ImgComments.ImgCommentCount > pageSie {
		for i := 2; i <= maxPage && i < rsp.ImgComments.ImgCommentCount/pageSie+1; i++ {
			rsp, err := tool.GetProductPageImageCommentList(detail.ProductID, i, pageSie)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(rsp.ImgComments.ImgList) <= 0 {
				return
			}

			list := SaveImgList(detail, rsp.ImgComments.ImgList, imgPath)
			exportList = append(exportList, list...)
		}
	}
	return
	//fmt.Println(rsp)
}

// GetPageList 列表页
func GetPageList(i int, group *sync.WaitGroup) {

	exportList := make([]*excel.ExportData, 0)
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
		title := tool.TrimTitle(e.Text)

		// 通过地址获取产品id
		detailURL := e.Attr(config.C.Crawler.ListItemSelectorURL)
		if strings.HasPrefix(detailURL, `//`) {
			detailURL = strings.ReplaceAll(detailURL, `//`, "https://")
		}
		urlParams, _ := url.Parse(detailURL)
		productID := strings.TrimRight(strings.TrimLeft(urlParams.Path, "/"), ".html")

		// 请求详情
		wg.Add(1)
		go func() {
			list := GetDetail(&tool.ProductDetail{
				ProductID: productID,
				Title:     title,
			}, &wg)
			exportList = append(exportList, list...)
		}()

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// 加载完处理
	c.OnScraped(func(r *colly.Response) {
		wg.Wait()

		// excel数据导出
		f := excelize.NewFile()
		err := excel.Save2Excel(exportList, f, "Sheet1")
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = f.SaveAs(fmt.Sprintf("%s/jd_search_%s_%d.xlsx", config.C.DownLoad.Path, config.C.Crawler.SearchKeyword, i))

		group.Done()
	})
	url := tool.GetSearchURL(config.C.Crawler.SearchURL, config.C.Crawler.SearchKeyword, i)
	c.Visit(url)

}

// SaveImgList 保存图片列表
func SaveImgList(detail *tool.ProductDetail, imgList []*tool.ImgListItem, imgPath string) (list []*excel.ExportData) {
	list = make([]*excel.ExportData, 0)
	for _, item := range imgList {
		// 图片地址处理
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

		// 请求图片地址，下载
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		urlArr := strings.Split(item.ImageURL, `/`)
		fileName := urlArr[len(urlArr)-1]
		newImgPath := imgPath + fileName
		err = ioutil.WriteFile(newImgPath, body, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 导出数据
		item := &excel.ExportData{
			ProductID:        detail.ProductID,
			ProductName:      detail.Title,
			SearchKeyword:    config.C.Crawler.SearchKeyword,
			CommentTime:      item.CommentVo.CreationTime,
			CommentKeyword:   config.C.Crawler.ImgCommentKeyword,
			CommentText:      item.CommentVo.Content,
			IsHasKeyword:     false,
			CommentImg:       newImgPath,
			CommentRemoteImg: item.ImageURL,
		}
		if strings.Contains(item.CommentText, item.CommentKeyword) {
			item.IsHasKeyword = true
		}
		list = append(list, item)

	}
	return
}

func main() {
	// 读取配置文件到struct,初始化变量
	config.InitConfig()

	group := sync.WaitGroup{}
	for i := 1; i <= config.C.Crawler.SearchMaxPage; i++ {
		group.Add(1)
		go GetPageList(i, &group)
	}
	group.Wait()
	fmt.Println("下载完整")
}
