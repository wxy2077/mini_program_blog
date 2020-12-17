// miniprogram/pages/cate/cate.js

import {getArticleList} from '../../api/article'
import {getArticleCategory} from '../../api/category.js'


Page({

  /**
   * 页面的初始数据
   */
  data: {
    selectedTab: 0,   // 当前选中的分类下角标
    tabs: [],         // 分类列表
    page: 1,
    pageSize: 10,
    articleList: []   // 文章列表
  },
  // 切换分类
  switchTab: function (event) {
    if (event.target.dataset.tabindex !== this.data.selectedTab) {
        this.setData({
            selectedTab: event.target.dataset.tabindex,
            page: 1,
            pageSize: 10
        });
        this.fetchCateArticle(this.data.tabs[this.data.selectedTab].id)
    }
  },
  // 滑动
  swiperChange: function(event){
    this.setData({
      selectedTab: event.detail.current
    });
  },

  // 获取分类
  fetchCategory(){
    getArticleCategory().then(res => {
      this.setData({tabs: res.data})
      // 同步获取分类
      this.fetchCateArticle(this.data.tabs[this.data.selectedTab].id)
    })
  },
  // 获取分类文章
  fetchCateArticle(cateId){
    getArticleList({cateId: cateId, page: this.data.page, pageSize: this.data.pageSize}).then(res=>{
      this.setData({
        articleList: res.data.data
      });
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
  // 加载更多分类文章
  loadMore(event){
      // console.log(event.target.dataset.cateId)
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.fetchCategory()
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})