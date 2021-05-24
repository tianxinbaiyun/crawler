package chromedp

import (
	"github.com/chromedp/chromedp"
	"time"
)

// Login 登陆
func (j *Job) Login(url, nodeUsername, nodePassword, nodeSubmit string) (err error) {
	var body string
	// 访问登陆页面
	err = chromedp.Run(
		j.Ctx,
		chromedp.Navigate(url), // 跳转登录页面
		chromedp.WaitReady(`body`, chromedp.ByQuery), // 等待加载
		//chromedp.Click(nodeSwitch, chromedp.ByQuery, chromedp.NodeVisible), // 切换至账密登录
		//chromedp.Sleep(time.Second*5),
		chromedp.OuterHTML(`/html/body`, &body),
	)
	if err != nil {
		j.Error = err
		return
	}

	// 输入姓名密码
	err = chromedp.Run(
		j.Ctx,
		chromedp.SendKeys(
			nodeUsername,
			j.Account.UserName,
			chromedp.ByQuery,
		),
		chromedp.SendKeys(
			nodePassword,
			j.Account.Password,
			chromedp.ByQuery,
		),
	)
	if err != nil {
		j.Error = err
		return
	}

	// 点击登陆按钮
	err = chromedp.Run(
		j.Ctx,
		chromedp.Click(nodeSubmit),
		chromedp.WaitReady(`body`, chromedp.ByQuery),
		chromedp.Sleep(time.Hour),
	)
	if err != nil {
		j.Error = err
		return
	}
	return
}
