# golang爬虫框架总结


## colly

[colly](https://github.com/gocolly/colly)
是快如闪电而优雅的爬虫框架，提供简洁的API能够帮助你构建爬虫应用。使用Colly，你可以轻松地从网站中提取结构化数据，这些数据可用于广泛的应用程序，如数据挖掘，数据处理或归档。

colly 的特性：

- 简单的API
- 快速（单核上> 1k请求/秒）
- 控制请求延迟和每个域名的最大并发数
- 自动cookie和session处理
- 同步/异步/并行抓取
- 高速缓存
- 对非unicode响应自动编码
- Robots.txt支持
- 分布式抓取
- 支持通过环境变量配置
- 随意扩展



## goquery

[goquery](https://github.com/PuerkitoBio/goquery)
带来了类似于Go语言的jQuery的语法和一组特性。它是基于Go的net/html包和CSS选择器库cascadia。因为net/html解析器返回的是节点，而不是功能齐全的DOM树，所以jQuery的有状态操作函数(如height()、css()、detach())被省略了。

goquery 的特性

- 与jQuery类似，语法简单
- 爬虫效率高
- goquery只支持utf-8编码

## selenium

[go-selenium](https://sourcegraph.com/sourcegraph/go-selenium)
是Go的Selenium WebDriver客户端。Selenium可以让浏览器自动执行各种Web应用。
它目前主要用于Web端的自动化测试，但它并不仅仅局限于此。
它还可以用于自动化管理基于Web的各种无聊费时的任务。

selenium 特性：

- 官网文档详情，可供参考
- 多语言
- 多平台
- 多浏览器
- 简单、灵活
- 因为是模拟操作WebDriver，它的性能一般
- golang版本与docker容器兼容不好，如FindElements，FindElement就不能正常获取


## chromedp









