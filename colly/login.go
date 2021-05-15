package colly

import (
	"fmt"
	"github.com/gocolly/colly"
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
func (j *Job) CookieLogin(uri, cookie string) {

	// 在提出请求之前打印 "访问…"
	j.Collector.OnRequest(func(r *colly.Request) {
		cookie := "session=eyJjc3JmX3Rva2VuIjoiR3doU0FqTG5RV3pkTkp2RE1vS1lCZmdPeFJ0VFZGaVVFQ2xYYmtadW1yYVBzeXFwY2VISSIsInVzZXJuYW1lIjoicXFxIn0.EmHo_w.pMG9x7Vdwd2INAw1O25NLw6saRk"
		err := j.Collector.SetCookies(uri, setCookieRaw(cookie))
		if err != nil {
			fmt.Println(err)
		}
	})
}

// setCookieRaw set cookies raw
func setCookieRaw(cookieRaw string) []*http.Cookie {
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
