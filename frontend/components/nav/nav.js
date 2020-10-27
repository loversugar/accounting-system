const systemInfo = wx.getSystemInfoSync();
Component({
  data: {
    statusHeight: Number,
    navHeight: 44
  },
  properties: {
    navTitle: {
      type: String,
      value: "导航"
    },
    prePageUrl: String  
  },

  lifetimes: {
    attached: function() {
      this.setData({
        statusHeight: systemInfo.statusBarHeight 
      })
    }
  },
  methods: {
      back() {
        wx.reLaunch({
          url: this.data.prePageUrl
        });
      } 
  }
})