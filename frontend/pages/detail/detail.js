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
    items: [{
      id: 1
    },{
      id: 2
    }]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    console.log(this.data.year)
     this.getBills(
       this.getStartDate(this.data.year, this.data.month), 
       this.getEndDate(this.data.year, this.data.month)
       )
  },

  getBills: function(startDate, endDate) {
    console.log(startDate)
    wx.showLoading({
      title: "加载数据中...",
      mask: true,
    });
    wx.request({
      url: app.globalData.remoteAddress + "/bill/getBillByDate?userId=1&startDate=" + startDate
       + "&endDate=" + endDate,
      data: {},
      header: {'content-type':'application/json'},
      method: 'GET',
      dataType: 'json',
      responseType: '',
      success: (result)=>{
        for (var attrName in result.data.data) {
          console.log(attrName)
        }
        this.setData({
          items: result.data.data
        })
        console.log(result.data.data)
      },
      fail: ()=>{
      },
      complete: ()=>{}
    });
    wx.hideLoading();
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
      var dataArray = e.detail.value.split('-')

      var year = dataArray[0]
      var month = dataArray[1]

      this.setData({
        year: year,
        month: month
      })

      this.getBills(
        this.getStartDate(year, month), 
        this.getEndDate(year, month)
        )
    }
  },

  getStartDate: function(year, month) {
      var suffix = year + "-" + month + "-"
      return suffix + "01"
  },

  getEndDate: function(year, month) {
      var suffix = year + "-" + month + "-"

      var currentMonthDay = new Date(parseInt(year), parseInt(month), 0).getDate()

      return suffix + currentMonthDay

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