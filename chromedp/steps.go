package chromedp

import (
	"context"
	"github.com/chromedp/chromedp"
)

func (s *Step) Do(ctx context.Context) (resp string, err error) {
	switch s.Action {
	case "Click":
		err = s.Click(ctx)
	case "Navigate":
		resp, err = s.Navigate(ctx)
	case "Reload":
		err = s.Reload(ctx)
	case "SendKeys":
		err = s.SendKeys(ctx)
	case "Title":
		resp, err = s.Title(ctx)

	}
	return
}

// Click 点击步骤
func (s *Step) Click(ctx context.Context) (err error) {

	err = chromedp.Run(
		ctx,
		chromedp.Click(s.Selector),
		chromedp.WaitReady(`body`, chromedp.ByQuery), // 等待加载
		chromedp.Sleep(s.Sleep),
	)
	if err != nil {
		return
	}
	return
}

// Navigate 访问操作
func (s *Step) Navigate(ctx context.Context) (body string, err error) {
	err = chromedp.Run(
		ctx,
		chromedp.Navigate(s.URL), // 跳转登录页面
		chromedp.WaitReady(`body`, chromedp.ByQuery), // 等待加载
		chromedp.Sleep(s.Sleep),
		chromedp.OuterHTML(`/html/body`, &body),
	)
	return
}

// Reload 刷新
func (s *Step) Reload(ctx context.Context) (err error) {
	err = chromedp.Run(
		ctx,
		chromedp.Reload(),
		chromedp.Sleep(s.Sleep),
	)
	return
}

// SendKeys 输入信息
func (s *Step) SendKeys(ctx context.Context) (err error) {
	err = chromedp.Run(
		ctx,
		chromedp.SendKeys(
			s.Selector,
			s.Value,
			chromedp.ByQuery,
		),
		chromedp.Sleep(s.Sleep),
	)
	return
}

// TextContent 获取内容
func (s *Step) TextContent(ctx context.Context) (resp string, err error) {
	err = chromedp.Run(
		ctx,
		chromedp.TextContent(s.Selector, &resp),
		chromedp.Sleep(s.Sleep),
	)
	return
}

// Title 获取title
func (s *Step) Title(ctx context.Context) (title string, err error) {
	err = chromedp.Run(
		ctx,
		chromedp.Navigate(s.URL), // 跳转登录页面
		chromedp.WaitReady(`body`, chromedp.ByQuery), // 等待加载
		chromedp.Title(&title),
		chromedp.Sleep(s.Sleep),
	)
	return
}
