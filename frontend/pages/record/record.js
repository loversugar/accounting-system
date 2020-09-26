// pages/record/record.js
const app =  getApp();
const systemInfo = wx.getSystemInfoSync();

Page({
  /**
   * 页面的初始数据
   */
  data: {
    navTitle: "发布",
    prePageUrl: String,
    tagWidth: Number,
    tags: [
      {
        id: 1,
        name: "蔬菜",
        iconfontName: "icon-chongwu"
      },
      {
        id: 2,
        name: "水果",
        iconfontName: "icon-shucai"
      },
      {
        id: 3,
        name :"交通",
        iconfontName: "icon-jiaotong"
      },
      {
        id: 4,
        name : "宠物",
        iconfontName: "icon-shuiguo"
      }
    ]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.setData({
      tagWidth: (systemInfo.screenWidth - 100) / 4
    })
  },
  click(e) {
    var dataArray = [];
    let selectedItemId = e.target.dataset.item.id
    for (let i=0; i<this.data.tags.length; i++) {
      let currentTag = this.data.tags[i]
      if (selectedItemId == currentTag.id) {
        currentTag.iconfontBackgroundColor = "#00b38a"
      } else {
        currentTag.iconfontBackgroundColor = "#f1f1f1"
      }
      dataArray = dataArray.concat(currentTag)
    }
    console.log(dataArray)
    this.setData({
      tags: dataArray
    })
  },

  onClickOne(e) {
    for (var i=0; i<this.data.tags.length; i++) {
      var tag = this.data.tags[i];
      if(tag.id === e.detail.tagId) {
        this.data.tags[i].isClick = true
      } else {
        this.data.tags[i].isClick = false
      }
    }
    console.log(this.data.tags)
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
    this.setData({
      prePageUrl: app.globalData.preTabUrl
    })
    wx.hideTabBar({
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