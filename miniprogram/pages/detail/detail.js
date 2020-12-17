// miniprogram/pages/detail/detail.js
import {getArticleDetail} from '../../api/article'
Page({

  /**
   * 页面的初始数据
   */
  data: {
      article: {}
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var _this = this
    const eventChannel = this.getOpenerEventChannel()

    eventChannel.on('acceptDataFromOpenerPage', function(data) {
      getArticleDetail(data).then(res=>{
        _this.setData({article: res.data})
      })
    })
  },
  previewImage(event){
      var src = event.currentTarget.dataset.src;//获取data-src;
      // console.log(src)
      wx.previewImage({
        current: src, // 当前显示图片的http链接
        urls: [src]
      })
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