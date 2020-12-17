// 文章操作api


import requests from "../utils/request"

// 获取轮播图推荐文章
export function getRecommendArticle() {
  return requests(
    '/article/get/recommend',
    'GET',
  )
} 

// 获取文章列表
export function getArticleList(data) {
  return requests(
    '/article/get/list',
    'GET',
    data
  )
} 

// 获取文章详情
export function getArticleDetail(data) {
  return requests(
    '/article/get/detail',
    'GET',
    data
  )
} 