package colly

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/url"
	"os"
	"time"
)

/*
请求执行之前调用
	- OnRequest
响应返回之后调用
	- OnResponse
监听执行 selector
	- OnHTML
监听执行 selector
	- OnXML
错误回调
	- OnError
完成抓取后执行，完成所有工作后执行
	- OnScraped
取消监听，参数为 selector 字符串
	- OnHTMLDetach
取消监听，参数为 selector 字符串
	- OnXMLDetach
*/

// Job 任务结构体
type Job struct {
	URI       string            // 请求地址
	TryTimes  int               // 重试次数
	Collector *colly.Collector  // 收集器
	File      *os.File          // 文件
	URL       *url.URL          // 请求地址解析数据
	Header    map[string]string // 请求头
}

// JobInit 任务初始化
func NewJob(uri string, header map[string]string, limit *colly.LimitRule) *Job {
	// 获取域名
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	// job 初始化
	job := &Job{
		URI:       uri,
		TryTimes:  0,
		Collector: colly.NewCollector(),
		File:      nil,
		URL:       u,
		Header:    header,
	}

	// 仅访问域
	job.Collector.AllowedDomains = []string{u.Host}

	// 允许重复访问
	job.Collector.AllowURLRevisit = true

	// 表示抓取时异步的
	// c.Collector.Async = true

	// 模拟浏览器
	if v, ok := header["user-agent"]; ok {
		job.Collector.UserAgent = v
	} else {
		job.Collector.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"
	}

	// 设置限制规则
	err = job.SetLimitRule(limit)
	if err != nil {
		panic(err)
	}
	// 随机UserAgent
	extensions.RandomUserAgent(job.Collector)

	return job
}

// SetLimitRule 设置限制规则
func (j *Job) SetLimitRule(limit *colly.LimitRule) (err error) {
	// 限制采集规则
	/*
		在Colly里面非常方便控制并发度，只抓取符合某个(些)规则的URLS
		colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 5}，表示限制只抓取域名是douban(域名后缀和二级域名不限制)的地址，当然还支持正则匹配某些符合的 URLS

		Limit方法中也限制了并发是5。为什么要控制并发度呢？因为抓取的瓶颈往往来自对方网站的抓取频率的限制，如果在一段时间内达到某个抓取频率很容易被封，所以我们要控制抓取的频率。
		另外为了不给对方网站带来额外的压力和资源消耗，也应该控制你的抓取机制。
	*/
	if limit != nil {
		err = j.Collector.Limit(limit)
		if err != nil {
			return
		}
	} else {
		err = j.Collector.Limit(&colly.LimitRule{
			// Filter domains affected by this rule
			// 筛选受此规则影响的域
			DomainGlob: j.URL.Host + "/*",
			// Set a delay between requests to these domains
			// 设置对这些域的请求之间的延迟
			Delay: 1 * time.Second,
			// Add an additional random delay
			// 添加额外的随机延迟
			RandomDelay: 1 * time.Second,
			// 设置并发
			Parallelism: 5,
		})
		if err != nil {
			return
		}
	}
	return
}

// SetHeaders 设置 headers
func (j *Job) SetHeaders(header map[string]string) {
	if _, ok := header["host"]; !ok {
		header["host"] = j.URL.Host
	}
	if _, ok := header["accept"]; !ok {
		header["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	}
	if _, ok := header["user-agent"]; !ok {
		header["accept"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"
	}

	// 在提出请求之前打印 "访问…"
	j.Collector.OnRequest(func(r *colly.Request) {
		for key, value := range header {
			r.Headers.Add(key, value)
		}
	})
	return
}
