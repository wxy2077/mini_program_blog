// miniprogram/pages/welcome/welcome.js

import {getAllArticle} from '../../api/article'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    avatar: '/images/avatar.jpeg',
    welcome_text: '欢迎来到王小右小程序博客',
    blog_theme: 'Just For Fun'
  },
  inblog: function(){
    wx.switchTab({
      url: "/pages/index/index"
    })
    // wx.redirectTo({
    //   url: '/pages/index/index',
    // })
    // setTimeout(function(){
    //   console.log("延时进入博客")
    //   wx.switchTab({url: "/pages/index/index"})
    // }, 500)
  },
  
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    // setTimeout(function(){
    //   // console.log("1.5秒进入博客")
    //   wx.switchTab({url: "/pages/index/index"})
    // }, 1500)
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