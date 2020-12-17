/*
* @Time    : 2020-11-19 10:43
* @Author  : CoderCharm
* @File    : directory.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package utils

import (
	"go.uber.org/zap"
	"mini_blog_server/global"
	"os"
)

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @title    createDir
// @description   批量创建文件夹
// @auth                     （2020/04/05  20:22）
// @param     dirs            string
// @return    err             error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.MINI_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.MINI_LOG.Error("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
