

import requests from "../utils/request"

// 获取文章分类
export function getArticleCategory() {
  return requests(
    '/article/get/category',
    'GET',
  )
} 
