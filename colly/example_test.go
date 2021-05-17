package colly

import (
	"errors"
	"github.com/gocolly/colly"
	"log"
	"testing"
)

// Example 获取百度热搜版
func Example() {
	var err error
	uri := "http://www.baidu.com"
	job := NewJob(uri, nil, nil)
	job.Collector.OnHTML(".s-hotsearch-content .hotsearch-item .title-content", func(element *colly.HTMLElement) {
		text := element.Text
		log.Println(text)
	})

	// 收到响应后
	job.Collector.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			err = errors.New("访问失败: " + job.URI)
			return
		}
	})

	err = job.Collector.Visit(job.URI)
	if err != nil {
		return
	}

	// 采集等待结束
	job.Collector.Wait()
	return
}

func TestGetBaiduHotSearch(t *testing.T) {
	Example()
	//=== RUN   TestGetBaiduHotSearch
	//2021/05/14 19:29:03 1国家卫健委派出专家组前往安徽热
	//2021/05/14 19:29:03 4恒河出现大量浮尸 印媒给出原因
	//2021/05/14 19:29:03 2以军方致电加沙居民称导弹将炸你家
	//2021/05/14 19:29:03 5国内灭活疫苗对多数变异株有效新
	//2021/05/14 19:29:03 3安徽确诊病例曾2次停留北京
	//2021/05/14 19:29:03 6加沙遭以军空袭前后卫星地图对比
	//--- PASS: TestGetBaiduHotSearch (0.32s)
	//PASS
}

func BenchmarkGetBaiduHotSearch(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Example()
		}
	})
	//BenchmarkGetBaiduHotSearch-20    	     146	   9522083 ns/op
}
