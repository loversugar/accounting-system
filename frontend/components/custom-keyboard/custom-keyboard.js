const systemInfo = wx.getSystemInfoSync();
const app = getApp();

Component({
  data: {
    keyWidth: Number,
    keyHeight: Number,
    money: "0.00",
    note: "",
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
        number: "今天"
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
        number: "+"
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
        number: "-"
      },
      {
        number: "."
      },
      {
        number: "0"
      },
      {
        number: "删除"
      },
      {
        number: "完成"
      }
    ]
  },
  lifetimes: {
    ready() {
      let screenWidth = systemInfo.screenWidth;
      let screenHeight = systemInfo.screenHeight;
      this.setData({
        keyWidth: screenWidth / 4,
        keyHeight: ((screenHeight * 0.4) - 35) / 4
      })
    },
  },
  observers: {
    'categoryId': function(categoryId) {
      this.init() 
    }
  },
  properties: {
    categoryId: Number
  },
  methods: {
    init() {
      this.setData({
        money: "0.00",
        note: "",
      })
    },
    click(e) {
      if (e.target.dataset.item.number === '完成') {
        
      }

      if (this.data.money === "0.00") {
        this.setData({
          money: e.target.dataset.item.number
        })
      } else {
        this.setData({
          money: this.data.money + e.target.dataset.item.number
        })
      }
    },

    finish() {

    }
  }
})