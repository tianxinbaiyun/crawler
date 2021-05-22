package selenium

import (
	"sourcegraph.com/sourcegraph/go-selenium"
)

// 常量定义
const (
	// WebDriverURL 为docker容器启动的地址
	// docker run -d -p 4444:4444 -v /dev/shm:/dev/shm selenium/standalone-chrome:4.0.0-beta-4-prerelease-20210517
	WebDriverURL     = "http://localhost:4444/wd/hub"
	ChromeDriverPath = "chromedriver"
)

// Job 任务结构体
type Job struct {
	URI      string                // 请求地址
	TryTimes int                   // 重试次数
	Header   map[string]string     // 请求头
	Driver   selenium.WebDriver    //
	Caps     selenium.Capabilities //
}

// NewJob 任务初始化
func NewJob(uri string) *Job {
	var err error

	// job 初始化
	job := &Job{
		URI:      uri,
		TryTimes: 0,
		Header:   nil,
		Driver:   nil,
	}

	// 设置浏览器
	job.SetChromeCaps()

	// 请求远程WebDriver
	driver, err := selenium.NewRemote(job.Caps, WebDriverURL)
	if err != nil {
		panic(err)
	}
	job.Driver = driver

	err = job.Driver.Get(uri)
	if err != nil {
		panic(err)
	}

	return job
}

// Close 关闭浏览器
func (j *Job) Close() (err error) {
	err = j.Driver.Quit()
	return
}

// SetCaps 设置浏览器，此处选择chrome
func (j *Job) SetChromeCaps() {
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	j.Caps = caps
	return
}
