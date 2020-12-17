/*
* @Time    : 2020-11-24 10:30
* @Author  : CoderCharm
* @File    : article.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 对应表数据的 结构体
**/
package model

import "time"

// 推荐文章 轮播
type RecommendArticle struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	SeoMetaKey  string `json:"seo_meta_key"`
	Cover  string `json:"cover"`
	ArticleUrl  string `json:"article_url"`
}

/*
 列表页 文章查询
 */
type ArticleIndex struct {
	Title string `json:"title"`
	Cover  string `json:"cover"`
	ArticleUrl  string `json:"article_url"`
	ClickCount  int64 `json:"click_count"`
	AddTime  time.Time `json:"add_time"`
}

/*
文章分类
*/
type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

/*
文章详情
*/
type ArticleDetail struct {
	Title string `json:"title"`
	Cover  string `json:"cover"`
	Content  string `json:"content"`
	ArticleUrl  string `json:"article_url"`
	ClickCount  int64 `json:"click_count"`
	AddTime  time.Time `json:"add_time"`
}