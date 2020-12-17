# 微信小程序:个人博客接口

## 介绍
使用`Gin`框架构建个人博客小程序数据接口。

主要参考开源后台项目`gin-vue-admin`构建接口服务 https://github.com/flipped-aurora/gin-vue-admin 感谢大佬的项目模版。


## 特点

- `zap`   日志管理
- `viper` 读取配置文件
- `gorm`  数据库操作


## 目录说明

<details>
<summary>目录</summary>

```
.
|____api                接口目录
| |____v1                   v1
| | |____article.go             文章接口

|____config             配置文件
| |____config.go            汇聚配置文件结构体
| |____zap.go               定义读取配置文件yaml格式
| |____gorm.go              定义读取数据库配置文件yaml格式

|____core               核心的基础服务配置
| |____server.go            通过全局路由 启动http服务
| |____zap.go               日志文件配置
| |____viper.go             读取配置文件

|____global             定义全局变量
| |____global.go            

|____initialize         各种初始化文件
| |____router.go            汇聚各模块路由
| |____gorm.go              定义`gorm`数据库配置

|____middleware         定义中间价
| |____cors.go              跨域中间件
| |____zaplogger.go         记录日志中间价

|____model              定义各种model
| |____request              请求参数model
| | |____common.go              通用请求参数
| |____response             响应参数model
| | |____common.go              通用响应参数
| | |____response.go    
| |____article.go           文章表model

|____router             路由分组(配置中间价)
| |____article.go           文章model路由

|____service            表操作逻辑
| |____article.go           文章表的curd操作

|____utils              通用的工具方法
| |____rotatelogs_unix.go
| |____directory.go
| |____constant.go

|____config.yaml        配置文件
|____go.mod
|____go.sum
|____main.go            启动文件
|____README.md

```

</details>

## 如何运行

#### 调试

```shell
go run main.go
```



## 参考

- gorm文档 https://gorm.io/zh_CN/docs/index.html

