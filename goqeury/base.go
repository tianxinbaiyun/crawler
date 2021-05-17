package goqeury

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// Job 任务结构体
type Job struct {
	URL      string // 请求地址
	TryTimes int    // 重试次数
}

// JobInit 任务初始化
func NewJob(url string) *Job {
	return &Job{
		URL:      url,
		TryTimes: 0,
	}
}

// FetchData 抓取数据
func (job *Job) FetchData() (doc *goquery.Document, err error) {
	res, err := http.Get(job.URL)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return
	}
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	return
}
