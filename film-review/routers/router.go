package routers

import (
	"github.com/astaxie/beego"
	"github.com/tianxinbaiyun/crawler/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/crawler", &controllers.CrawlMovieController{}, "*:CrawlMovie")
}
