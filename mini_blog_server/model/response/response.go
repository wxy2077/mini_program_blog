/*
* @Time    : 2020-11-23 17:10
* @Author  : CoderCharm
* @File    : response.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


const (
	ERROR   = 400
	SUCCESS = 200
)

func Result(code int, msg string, data interface{},  c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS,"操作成功", map[string]interface{}{}, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS,  message, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, "操作成功", data, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, message, data, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "操作失败", map[string]interface{}{}, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, message, data, c)
}