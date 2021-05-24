package chromedp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"log"
	"net/http"
	"time"
)

// 常量定义
const (
	UserName  = "admin"
	Password  = "admin"
	RemoteURL = "http://localhost:9222"
)

// Job 任务结构体
type Job struct {
	Ctx     context.Context
	Account *Account // 账号
	Steps   []*Step  // 步骤
	Error   error    // 错误
	Cancel  context.CancelFunc
	Debug   bool
}

// Account 账号
type Account struct {
	ID       string
	UserName string // 用户名
	Password string // 密码
	Nickname string // 显示昵称
}

// Step 步骤
type Step struct {
	URL      string        // 地址
	Action   string        // 动作
	Selector string        // 选择器
	Value    string        // 内容
	WantResp bool          // 是否需要结果
	Sleep    time.Duration // 等待时间
	Cron     string        // 定时操作
}

// NewJob 实例化任务
func NewJob(debug bool) *Job {
	ctx, cancel := chromedp.NewContext(context.Background())

	job := &Job{
		Ctx: ctx,
		Account: &Account{
			UserName: UserName,
			Password: Password,
		},
		Steps:  make([]*Step, 0),
		Error:  nil,
		Cancel: cancel,
		Debug:  debug,
	}
	if job.Debug {
		job.GetLocalContext()
	} else {
		job.GetProxyContext()
	}
	return job
}

// SetSteps 设置步骤
func (j *Job) SetSteps(steps []*Step) {
	if len(steps) <= 0 {
		return
	}
	for _, step := range steps {
		j.Steps = append(j.Steps, step)
	}
	return
}

// Close 关闭
func (j *Job) Close() {
	defer j.Cancel()
	v := chromedp.FromContext(j.Ctx)
	if v.Target != nil {
		target := fmt.Sprintf("%s/close/%s", RemoteURL, v.Target.TargetID)
		_, err := http.Get(target)
		if err != nil {
			return
		}
	}
	return
}

// GetLocalContext 用于本地测试的 context
func (j *Job) GetLocalContext() {
	log.Println("use local context.")
	useragent := `Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 Edg/87.0.664.66`

	// 设置应用配置
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", !j.Debug), // debug使用
		chromedp.UserAgent(useragent),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// 设置上下文
	ctx, cancel := chromedp.NewContext(c)
	j.Ctx = ctx
	j.Cancel = cancel
	err := chromedp.Run(
		j.Ctx,
		emulation.SetDeviceMetricsOverride(1920, 1024, 1.0, false),
	)
	if err != nil {
		panic(err)
	}

	return
}

// BrowserInfo 代理远程信息
// 指向一个可用的 browser
type BrowserInfo struct {
	Browser              string
	ProtocolVersion      string `json:"Protocol-Version"`
	UserAgent            string `json:"User-Agent"`
	V8Version            string `json:"V8-Version"`
	WebKitVersion        string `json:"WebKit-Version"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// GetProxyContext 使用代理的 context
func (j *Job) GetProxyContext() {
	log.Println("use proxy context.")
	var err error
	res := BrowserInfo{}
	target := fmt.Sprintf("%s/version", RemoteURL)

	resp, err := http.Get(target)
	if err != nil {
		panic(err)
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		panic(err)
	}
	ctx, _ := chromedp.NewRemoteAllocator(context.Background(), res.WebSocketDebuggerURL)
	ctx, cancel := chromedp.NewContext(ctx)
	j.Ctx = ctx
	j.Cancel = cancel

	err = chromedp.Run(
		j.Ctx,
		emulation.SetDeviceMetricsOverride(1920, 1024, 1.0, false),
	)
	if err != nil {
		panic(err)
	}

	return
}
