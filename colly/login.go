package colly

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"strings"
)

// Login 登录
func (j *Job) Login(uri string, body map[string]string) (res *colly.Response, err error) {
	// 收到响应后
	j.Collector.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Println("请求失败: ", uri)
			return
		}
		res = r
	})
	err = j.Collector.Post(uri, body)
	if err != nil {
		return
	}

	return
}

// CookieLogin cookie 登录
func (j *Job) CookieLogin(uri, cookie string) (err error) {

	// 在提出请求之前打印 "访问…"
	j.Collector.OnRequest(func(r *colly.Request) {
		err = j.Collector.SetCookies(uri, SplitCookieRaw(cookie))
		if err != nil {
			log.Println(err)
			return
		}
	})
	return
}

// SplitCookieRaw set cookies raw
func SplitCookieRaw(cookieRaw string) []*http.Cookie {
	// 可以添加多个cookie
	var cookies []*http.Cookie
	cookieList := strings.Split(cookieRaw, "; ")
	for _, item := range cookieList {
		keyValue := strings.Split(item, "=")

		name := keyValue[0]
		valueList := keyValue[1:]
		cookieItem := http.Cookie{
			Name:  name,
			Value: strings.Join(valueList, "="),
		}
		cookies = append(cookies, &cookieItem)
	}
	return cookies
}
