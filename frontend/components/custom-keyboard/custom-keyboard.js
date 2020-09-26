const systemInfo = wx.getSystemInfoSync();

Component({
  data: {
    keyWidth: Number,
    boardKeys: [
      {
        number: "7"
      },
      {
        number: "8"
      },
      {
        number: "9"
      },
      {
        number: "8"
      },
      {
        number: "4"
      },
      {
        number: "5"
      },
      {
        number: "6"
      },
      {
        number: "8"
      },
      {
        number: "1"
      },
      {
        number: "2"
      },
      {
        number: "3"
      },
      {
        number: "8"
      },
      {
        number: "8"
      },
      {
        number: "0"
      },
      {
        number: "8"
      },
      {
        number: "8"
      }
    ]
  },
  lifetimes: {
    ready() {
      let screenWidth = systemInfo.screenWidth;
      this.setData({
        keyWidth: screenWidth / 4
      })
    }
  },
  properties: {},
  methods: {}
})