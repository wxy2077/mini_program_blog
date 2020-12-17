/*
* @Time    : 2020-11-24 10:16
* @Author  : CoderCharm
* @File    : common.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// 请求文章分类
type ArticleCategory struct {
	CateId int `json:"cateId" form:"cateId"`
}

// 请求文章详情
type ArticleDetail struct {
	Href string `json:"href" form:"href"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string
}

type Empty struct{}
