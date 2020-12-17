/*
* @Time    : 2020-11-17 11:13
* @Author  : CoderCharm
* @File    : article.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 路由分组
**/
package router


import (
	"github.com/gin-gonic/gin"
	"mini_blog_server/api/v1"
	"mini_blog_server/middleware"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	InitArticleGroup := Router.Group("article").Use(middleware.ZapLogger())
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		InitArticleGroup.GET("/get/recommend", v1.GetRecommendArticle)   // 获取所有文章
		InitArticleGroup.GET("/get/list", v1.GetArticleList)   // 获取所有文章
		InitArticleGroup.GET("/get/category", v1.GetCategoryList)   // 获取所有文章
		InitArticleGroup.GET("/get/detail", v1.GetArticleDetail)   // 获取文章详情
	}
}