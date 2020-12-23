/*
* @Time    : 2020-11-17 11:35
* @Author  : CoderCharm
* @File    : router.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 初始化路由
**/
package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "mini_blog_server/docs"
	"mini_blog_server/middleware"
	"mini_blog_server/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 跨域
	Router.Use(middleware.CORSMiddleware())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("/mini/api")

	router.InitArticleRouter(ApiGroup) // 注册文章路由

	return Router
}
