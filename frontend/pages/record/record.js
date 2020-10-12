// pages/record/record.js
const app =  getApp();
const systemInfo = wx.getSystemInfoSync();

var currentSelectedKeyId = 0

Page({
  /**
   * 页面的初始数据
   */
  data: {
    navTitle: "发布",
    prePageUrl: String,
    tagWidth: Number,
    isShowKeyboard: false,
    categoryId: Number,
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
    this.getCategories()
    this.setData({
      tagWidth: (systemInfo.screenWidth - 100) / 4
    })
  },
  click(e) {
    let selectedItemId = e.target.dataset.item.id
    if (currentSelectedKeyId === selectedItemId) {
      return
    }
    currentSelectedKeyId = selectedItemId;
    var dataArray = [];
    for (let i=0; i<this.data.tags.length; i++) {
      let currentTag = this.data.tags[i]
      if (selectedItemId == currentTag.id) {
        currentTag.iconfontBackgroundColor = "#00b38a"
      } else {
        currentTag.iconfontBackgroundColor = "#f1f1f1"
      }
      dataArray = dataArray.concat(currentTag)
    }

    this.onClickOne(e)

    this.setData({
      categoryId: selectedItemId,
      isShowKeyboard: true,
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
  },

  sendBill() {
    wx.showLoading({
      title: "发送中...",
      mask: true,
    });
    wx.request({
      url: app.globalData.remoteAddress + '/bill/addBill',
      data: {
        userId: app.globalData.userId,
        consumption: 9,
        categoryId: 1, 
      },
      header: {'content-type':'application/json'},
      method: 'post',
      dataType: 'json',
      responseType: 'text',
      success: (result)=>{
        wx.hideLoading();    
      },
      fail: ()=>{},
      complete: ()=>{}
    });
  },

  getCategories() {
    wx.request({
      url: app.globalData.remoteAddress + "/category/getCategories",
      data: {},
      header: {'content-type':'application/json'},
      method: 'GET',
      dataType: 'json',
      responseType: 'text',
      success: (result)=>{
        this.setData({
          tags: result.data
        })
        console.log(result.data)
      },
      fail: ()=>{},
      complete: ()=>{}
    });
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