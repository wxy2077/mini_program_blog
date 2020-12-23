/*
* @Time    : 2020-11-19 10:31
* @Author  : CoderCharm
* @File    : global.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 全局变量
**/
package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mini_blog_server/config"
)

var (
	MINI_DB     *gorm.DB
	MiNI_VP     *viper.Viper
	MINI_CONFIG config.Server
	MINI_LOG    *zap.Logger
	MINI_REDIS  *redis.Client
)
