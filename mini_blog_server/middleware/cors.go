/*
* @Time    : 2020-11-17 11:16
* @Author  : CoderCharm
* @File    : cors.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 处理跨域
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理跨域请求,支持options访问
// 参考 https://stackoverflow.com/questions/29418478/go-gin-framework-cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}