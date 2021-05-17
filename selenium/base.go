package selenium

import (
	"net/url"
	"os"
)

// Job 任务结构体
type Job struct {
	URI      string            // 请求地址
	TryTimes int               // 重试次数
	File     *os.File          // 文件
	URL      *url.URL          // 请求地址解析数据
	Header   map[string]string // 请求头
}
