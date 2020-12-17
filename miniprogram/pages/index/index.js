//index.js
const app = getApp()

import {getRecommendArticle, getArticleList} from '../../api/article'
// import { debounce } from '../../utils/common'

Page({
    data: {
        recommendList: [],
        articleList: [],
        page: 1,
        pageSize: 10,
        overLine: false    
    },
    onLoad: function (options) {
        // 获取文章列表
        this.fetchArticlList()
    },
    /**
    * 生命周期函数--监听页面初次渲染完成
    */
    onReady: function () {
         // 获取推荐文章  
        getRecommendArticle().then(res=>{
            this.setData({recommendList: res.data})
        })
    },

    /**
    * 页面上拉触底事件的处理函数
    */
    onReachBottom: function () {
        this.fetchArticlList()
    },
    fetchArticlList: function(){
        getArticleList({page: this.data.page, pageSize: this.data.pageSize}).then(res => {
            if(res.data.data){
                this.setData({
                    articleList: this.data.articleList.concat(res.data.data),
                    page: res.data.page + 1
                })
            }else{
                this.setData({
                    overLine: true
                })
            }
        })
    },
    
    onGoToDetail: function(event){
        // var $this = this;
        const href = event.currentTarget.dataset.href

        wx.navigateTo({
          url: '/pages/detail/detail',
          success: function(res) {
            // 通过eventChannel向被打开页面传送数据
            res.eventChannel.emit('acceptDataFromOpenerPage', { href: href })
          }
        })
    },
      
})
