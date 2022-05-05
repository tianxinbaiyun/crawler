package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config 配置信息 yaml 结构体
type Config struct {
	Header   Header   `yaml:"header"`
	Crawler  Crawler  `yaml:"crawler"`
	DownLoad DownLoad `yaml:"download"`
}

// Header 请求头
type Header struct {
	Origin    string `yaml:"origin"`
	Referer   string `yaml:"referer"`
	UserAgent string `yaml:"user_agent"`
	Cookie    string `yaml:"cookie"`
	Host      string `yaml:Host`
}

// Crawler 抓包配置
type Crawler struct {
	SearchURL           string `yaml:"search_url"`
	SearchMaxPage       int    `yaml:"search_max_page"`
	SearchKeyword       string `yaml:"search_keyword"`
	ListItemSelector    string `yaml:"list_item_selector"`
	ListItemSelectorURL string `yaml:"list_item_selector_url"`
	ImgCommentMaxPage   int    `yaml:"img_comment_max_page"`
	ImgCommentPageSize  int    `yaml:"img_comment_page_size"`
}

// DownLoad 下载配置结构体
type DownLoad struct {
	Path             string `path`
	ImagesName       string `images_name`
	CommentExcelName string `comment_excel_name`
}

// C 全局配置信息
var C = Config{}

// InitConfig 初始化配置
func InitConfig() {
	fileName := "./config.yaml"
	//目录不存在，从指定的目录找
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fileName = "../config.yaml"
	}
	ret, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(ret, &C)
	if err != nil {
		panic(err)
	}
	fmt.Println(C)

	// 如果路径不存在，创建

}
