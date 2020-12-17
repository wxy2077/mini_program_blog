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


/*
获取推荐文章
*/
func GetRecommendArticle(c *gin.Context){
	RecommendArticleList := service.FetchRecommendArticleList()
	response.OkWithData(RecommendArticleList, c)
}

/*
获取文章列表
*/
func GetArticleList(c *gin.Context){
	pageInfo := request.PageInfo{Page:1, PageSize:10}
	cateInfo := request.ArticleCategory{}

	_ = c.ShouldBindQuery(&pageInfo)
	_ = c.ShouldBindQuery(&cateInfo)

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	articleList, total:= service.FetchArticleIndexList(pageInfo, cateInfo)

	response.OkWithDetailed(response.PageResult{
		Data: articleList,
		Total: total,
		Page: pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "success", c)

}

/*
获取所有分类
*/
func GetCategoryList(c *gin.Context){
	CategoryList := service.FetchCategoryList()
	response.OkWithData(CategoryList,c)
}

/*
获取文章详情
*/
func GetArticleDetail(c *gin.Context){
	href := request.ArticleDetail{}
	_ = c.ShouldBindQuery(&href)
	if err := utils.Verify(href, utils.ArticleDetailVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	DetailArticle := service.FetchArticleDetail(href)
	response.OkWithData(DetailArticle, c)
}