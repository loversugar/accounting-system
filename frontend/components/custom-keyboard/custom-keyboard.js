var util = require('../../utils/util.js')

const systemInfo = wx.getSystemInfoSync();
const app = getApp();

Component({
  data: {
    keyWidth: Number,
    keyHeight: Number,
    money: "0.00",
    note: "",
    innerCategoryId: Number,
    selectedTime: {
      type: Date,
      value: null
    },
    isCalendarShow: false,
    boardKeys: [
      {
        id: 1,
        display: "7",
      },
      {
        id: 2,
        display: "8"
      },
      {
        id: 3,
        display: "9"
      },
      {
        id: 4,
        display: util.formatTime(new Date()).split(" ")[0],
        name: "calendar"
      },
      {
        id: 5,
        display: "4"
      },
      {
        id: 6,
        display: "5"
      },
      {
        id: 7,
        display: "6"
      },
      {
        id: 8,
        display: "+"
      },
      {
        id: 9,
        display: "1"
      },
      {
        id: 10,
        display: "2"
      },
      {
        id: 11,
        display: "3"
      },
      {
        id: 12,
        display: "-"
      },
      {
        id: 13,
        display: "."
      },
      {
        id: 14,
        display: "0"
      },
      {
        id: 15,
        display: "删除"
      },
      {
        id: 16,
        display: "完成"
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

      this.setData({
        innerCategoryId: this.data.categoryId
      })
    },
  },
  observers: {
    'categoryId': function(categoryId) {
      this.init(categoryId) 
    }
  },
  properties: {
    categoryId: Number
  },
  methods: {
    init(categoryId) {
      this.setData({
        innerCategoryId: categoryId,
        money: "0.00",
        note: "",
        isCalendarShow: false
      })
    },
    click(e) {
      if (e.target.dataset.item.name === 'calendar') {
        this.setData({
          isCalendarShow: true,
        })
        return
      }
      if (e.target.dataset.item.display === '完成') {
        console.log(this.data.selectedTime)
        // this.finish()
        return
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

    dayClick(e) {
      var tempStr = 'boardKeys[3].display'
      var selectedTime = e.detail.year+'/'+e.detail.month+'/'+e.detail.day
      this.setData({
        selectedTime: new Date(selectedTime),
        isCalendarShow: false,
        [tempStr]: util.formatTime(new Date(selectedTime)).split(" ")[0]
      })
    },

    finish() {
      console.log(this.data.selectedTime)
      // this.triggerEvent("doSendBill", {
      //   categoryId: this.data.categoryId,
      //   selectedTime: "",
      //   consumption: this.data.consumption,
      //   note: this.data.note
      // })
    }
  }
})