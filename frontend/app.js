//app.js
App({
  onLaunch: function () {
    this.globalData.preTabUrl = '/pages/detail/detail'
    // 展示本地存储能力

    var openid = wx.getStorageSync('openid')
    if (!openid || openid === nil || openid === "") {
      // 登录
      wx.login({
        success: res => {
          this.globalData.currentCode = res.code
          // 发送 res.code 到后台换取 openId, sessionKey, unionId
          // wx.request({
          //   url: 'http://localhost:8888/accounting-system/user/login',
          //   data: {
          //     code: res.code
          //   }
          // })
        }
      })
    } 
  },
  globalData: {
    userInfo: null,
    preTabUrl: null,
    openid: null
  }
})