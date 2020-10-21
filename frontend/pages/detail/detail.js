// pages/detail/detail.js
const util = require('../../utils/util.js')
			 
const app  =  getApp();
			 
Page({

  /**
   * 页面的初始数据
   */
  data: {
    year: util.getYear(new Date()),
    month: util.getMonth(new Date()),
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
     
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  selectDate: function() {
    console.log("selectDate")
  },

  getDateTime: function(e) {
    if (e.detail.value) {
      console.log(e)
      var dataArray = e.detail.value.split('-')
      this.setData({
        year: dataArray[0],
        month: dataArray[1]
      })
    }
  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
    app.globalData.preTabUrl = "/pages/detail/detail"
    wx.showTabBar({
      animation: true,
    });
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