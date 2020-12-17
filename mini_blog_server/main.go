/*
* @Time    : 2020-11-17 10:42
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package main

import (
	"mini_blog_server/core"
	"mini_blog_server/global"
	"mini_blog_server/initialize"
)

func main() {

	global.MiNI_VP = core.Viper()          // 初始化Viper
	global.MINI_LOG = core.Zap()           // 初始化zap日志库

	global.MINI_DB = initialize.Gorm()     // gorm连接数据库

	// 程序结束前关闭数据库链接
	db, _ := global.MINI_DB.DB()
	defer db.Close()

	core.RunWindowsServer("127.0.0.1:7000")

}

