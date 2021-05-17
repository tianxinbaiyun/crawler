package colly

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// 登陆例子使用go-admin框架
func TestLogin(t *testing.T) {

	uri := "http://localhost:9033/admin/signin"
	job := NewJob(uri, nil, nil)

	res, err := job.Login(uri, map[string]string{"username": "admin", "password": "admin"})
	assert.NoError(t, err)
	cookie := res.Headers.Get("Set-Cookie")
	assert.NotEmpty(t, cookie)
}

func TestCookieLogin(t *testing.T) {
	// 初始化任务
	loginURL := "http://localhost:9033/admin/signin"
	job := NewJob(loginURL, nil, nil)

	// 登陆
	res, err := job.Login(loginURL, map[string]string{"username": "admin", "password": "admin"})
	assert.NoError(t, err)
	cookie := res.Headers.Get("Set-Cookie")
	assert.NotEmpty(t, cookie)

	// 设置cookie
	visitURL := "http://localhost:9033/admin/info/manager"
	err = job.CookieLogin(visitURL, cookie)
	assert.NoError(t, err)

	// 获取用户在线信息
	job.Collector.OnHTML(".user-panel .info", func(element *colly.HTMLElement) {
		text := element.Text
		text = strings.ReplaceAll(text, "\n", "")
		text = strings.ReplaceAll(text, " ", "")
		assert.Equal(t, "admin在线", text)
	})

	// 收到响应后
	job.Collector.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			err = errors.New("访问失败: " + job.URI)
			return
		}
	})

	// 访问地址
	err = job.Collector.Visit(visitURL)
	assert.NoError(t, err)
}
