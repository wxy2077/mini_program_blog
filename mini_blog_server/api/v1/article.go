/*
* @Time    : 2020-11-17 11:26
* @Author  : CoderCharm
* @File    : article.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

**/

package v1

import (
	"github.com/gin-gonic/gin"
	"mini_blog_server/model/request"
	"mini_blog_server/model/response"
	"mini_blog_server/service"
	"mini_blog_server/utils"
)

// @Tags ArticleAPI
// @Summary 获取推荐文章
// @accept application/json
// @Produce application/json
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/recommend [get]
func GetRecommendArticle(c *gin.Context) {
	RecommendArticleList := service.FetchRecommendArticleList()

	//global.MINI_REDIS.Set()

	response.OkWithData(RecommendArticleList, c)
}

// @Tags ArticleAPI
// @Summary 获取文章列表
// @accept application/json
// @Produce application/json
// @Param cateInfo body request.ArticleCategory true "文章分类"
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/list [get]
func GetArticleList(c *gin.Context) {
	pageInfo := request.PageInfo{Page: 1, PageSize: 10}
	cateInfo := request.ArticleCategory{}

	_ = c.ShouldBindQuery(&pageInfo)
	_ = c.ShouldBindQuery(&cateInfo)

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	articleList, total := service.FetchArticleIndexList(pageInfo, cateInfo)

	response.OkWithDetailed(response.PageResult{
		Data:     articleList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "success", c)

}

// @Tags ArticleAPI
// @Summary 获取文章分类
// @accept application/json
// @Produce application/json
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/category [get]
func GetCategoryList(c *gin.Context) {
	CategoryList := service.FetchCategoryList()
	response.OkWithData(CategoryList, c)
}

// @Tags ArticleAPI
// @Summary 获取文章详情
// @accept application/json
// @Produce application/json
// @Param data body request.ArticleDetail true "文章href链接"
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/detail [get]
func GetArticleDetail(c *gin.Context) {
	href := request.ArticleDetail{}
	_ = c.ShouldBindQuery(&href)
	if err := utils.Verify(href, utils.ArticleDetailVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	DetailArticle := service.FetchArticleDetail(href)
	response.OkWithData(DetailArticle, c)
}
