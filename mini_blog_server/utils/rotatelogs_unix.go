/*
* @Time    : 2020-11-19 10:45
* @Author  : CoderCharm
* @File    : rotatelogs_unix.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package utils


import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"mini_blog_server/global"
	"os"
	"path"
	"time"
)

// GetWriteSyncer zap logger中加入file-rotatelogs
func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.MINI_CONFIG.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.MINI_CONFIG.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.MINI_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}