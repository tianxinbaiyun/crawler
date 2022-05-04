package main

import (
	"fmt"
	"github.com/tianxinbaiyun/crawler/config"
	"strings"
)

// @Title  请填写文件名称（需修改）
// @Description  请填写文件描述（需修改）
// @Author  clx  2022/5/4 9:36 上午
// @Update  clx  2022/5/4 9:36 上午

func GetRemoteURL() string {
	return config.C.Crawler.SearchUrl
}

func GetSearchUrl(keyword string, page int) (s string) {
	s = config.C.Crawler.SearchUrl
	if keyword != "" {
		s = strings.ReplaceAll(s, "{{keyword}}", keyword)

	} else {
		s = strings.ReplaceAll(s, "&keyword={{keyword}}", "")
	}

	if page > 0 {
		s = strings.ReplaceAll(s, "{{page}}", fmt.Sprintf("%d", page))
	} else {
		s = strings.ReplaceAll(s, "&page={{page}}", "")
	}

	return
}
