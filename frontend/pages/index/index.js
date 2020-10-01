const app = getApp();

Page({
  data: {
      //判断小程序的API，回调，参数，组件等是否在当前版本可用。
      canIUse: wx.canIUse('button.open-type.getUserInfo'),
      isHide: false
  },

  onLoad: function() {
      var that = this;
      // 查看是否授权
      wx.getSetting({
          success: function(res) {
              if (res.authSetting['scope.userInfo']) {
                  var p1 = new Promise((resolve, reject) => 
                    wx.getUserInfo({
                      success: function(res) {
                        console.log(res)        
                        app.globalData.userInfo = res.userInfo
                        
                        resolve(res)
                      },
                      fail: function(err) {
                          reject(err)
                      }
                    })
                  );
                  // 发送 res.code 到后台换取 openId, sessionKey, unionId
                  var p2 = new Promise((resolve, reject) => {
                    wx.request({
                        url: 'http://localhost:8888/accounting-system/user/login',
                        method: 'GET',
                        data: {
                          code: app.globalData.currentCode
                        },
                        success: (res) => {
                            app.globalData.openid = res.data.openid
                            resolve(res)
                        },
                        fail:(err) => {
                            reject(err)
                        }
                      })
                  })

                  Promise.all([p1, p2]).then(values => {
                    console.log(app.globalData)
                    wx.switchTab({
                        url: '/pages/detail/detail'
                    });
                  })
                  
                  
              } else {
                  // 用户没有授权
                  that.setData({
                      isHide: true
                  });
              }
          }
      });
  },

  bindGetUserInfo: function(e) {
      if (e.detail.userInfo) {
          //用户按了允许授权按钮
          // 获取到用户的信息了，打印到控制台上看下
          console.log("用户的信息如下：");
          console.log(e.detail.userInfo);
          wx.request({
            url: 'http://localhost:8888/accounting-system/user/login',
            method: 'GET',
            data: {
              code: app.globalData.currentCode
            },
            success: (res) => {
                console.log(res)
                app.globalData.openid = res.openid
            },
            fail:(err) => {
            }
          })
          wx.switchTab({
            url: '/pages/detail/detail'
          });
      } else {
          //用户按了拒绝按钮
          wx.showModal({
              title: '警告',
              content: '您点击了拒绝授权，将无法进入小程序，请授权之后再进入!!!',
              showCancel: false,
              confirmText: '返回授权',
              success: function(res) {
                  // 用户没有授权成功，不需要改变 isHide 的值
                  if (res.confirm) {
                      console.log('用户点击了“返回授权”');
                  }
              }
          });
      }
  }
})
