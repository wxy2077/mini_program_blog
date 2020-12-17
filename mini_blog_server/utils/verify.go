/*
* @Time    : 2020-11-25 21:34
* @Author  : CoderCharm
* @File    : verify.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 参数验证
**/
package utils

var (
	PageInfoVerify = Rules{"Page": {Ge("1")}, "PageSize": {Le("50")}}
	ArticleDetailVerify = Rules{"Href": {NotEmpty()}}
)
