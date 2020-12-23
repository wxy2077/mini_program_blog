/*
* @Time    : 2020-12-18 15:58
* @Author  : CoderCharm
* @File    : redis.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
