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
        number: "7",
      },
      {
        id: 2,
        display: "8",
        number: "8"
      },
      {
        id: 3,
        display: "9",
        number: "9"
      },
      {
        id: 4,
        display: util.formatTime(new Date()).split(" ")[0],
        name: "calendar"
      },
      {
        id: 5,
        display: "4",
        number: "4",
      },
      {
        id: 6,
        display: "5",
        number: "5"
      },
      {
        id: 7,
        display: "6",
        number: "6"
      },
      {
        id: 8,
        display: "+"
      },
      {
        id: 9,
        display: "1",
        number: "1"
      },
      {
        id: 10,
        display: "2",
        number: "2"
      },
      {
        id: 11,
        display: "3",
        number: "3"
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
        display: "0",
        number: "0"
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
        keyHeight: ((screenHeight * 0.4) - 35) / 4,
        selectedTime: new Date(),
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
        isCalendarShow: false,
        selectedTime: new Date(),
        // 'boardKeys[3].display': util.formatTime(new Date()).split(" ")[0]
      })
    },
    click(e) {
      var dataMoney = this.data.money
      var inputDisplayName = e.target.dataset.item.display;
      if (dataMoney.includes("+") && dataMoney.split("+").length > 1) {
        var dataMoneySplit = dataMoney.split("+")
        if ((dataMoneySplit[1] != "" && inputDisplayName === "+") 
              || (dataMoneySplit[1] != "" && inputDisplayName === '-') 
              || (dataMoneySplit[1] != "" && inputDisplayName === '=')) {
          this.setData({
            money: parseFloat(dataMoneySplit[0]) + parseFloat(dataMoneySplit[1]) + ""
          })
        }
      }

      if (dataMoney.includes("-") && dataMoney.split("-").length > 1) {
        var dataMoneySplit = dataMoney.split("-")
        if ((dataMoneySplit[1] != "" && inputDisplayName === "+") 
              || (dataMoneySplit[1] != "" && inputDisplayName === '-') 
              || (dataMoneySplit[1] != "" && inputDisplayName === '=')) {
          this.setData({
            money: parseFloat(dataMoneySplit[0]) - parseFloat(dataMoneySplit[1]) + ""
          })
        }
      }

      if (e.target.dataset.item.name === 'calendar') {
        this.setData({
          isCalendarShow: true,
        })
        return
      }

      if(e.target.dataset.item.display === "删除") {
        this.setData({
          money: this.data.money.slice(0, this.data.money.length - 1).length > 0? 
            this.data.money.slice(0, this.data.money.length - 1) : "0.00"
        })
        return
      }
      if (e.target.dataset.item.display === '完成') {
        this.finish()
        return
      }

      if (this.data.money === "0.00") {
        this.setData({
          money: e.target.dataset.item.display
        })
      } else {
        if (e.target.dataset.item.display != '=') {
          this.setData({
            money: this.data.money + e.target.dataset.item.display
          })
        }
      }

      dataMoney = this.data.money
      var tempStr = 'boardKeys[15].display'
      if (dataMoney.includes("+")) {
        var dataMoneySplit = dataMoney.split("+")
        if (dataMoneySplit[1] != 0) {
          this.setData({
            [tempStr]: "="
          })
        } else {
          this.setData({
            [tempStr]: "完成"
          })
        }
      }

      if (dataMoney.includes("-")) {
       var dataMoneySplit = dataMoney.split("-")
        if (dataMoneySplit[1] != 0) {
          this.setData({
            [tempStr]: "="
          })
        } else {
          this.setData({
            [tempStr]: "完成"
          })
        }
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

    inputEdit(e) {
      let _this = this;
      let dataset = e.currentTarget.dataset
      let value = e.detail.value
      let name = dataset.name
      _this.data[name] = value
      _this.setData({
        name: _this.data[name]
      })
    },

    finish() {
      var dataMoney = this.data.money
      var consumption = "";
      if (dataMoney.includes("+")) {
        consumption = dataMoney.split("+")[0]
      } else if (dataMoney.includes("-")) {
        consumption = dataMoney.split("-")[0]
      } else {
        consumption = dataMoney
      }

      this.triggerEvent("sendBill", {
        categoryId: this.data.categoryId,
        selectedTime: util.getYear(this.data.selectedTime) 
          + "-" + util.getMonth(this.data.selectedTime) 
          + "-" + util.getDay(this.data.selectedTime),
        consumption: parseFloat(consumption),
        note: this.data.note,
        userId: app.globalData.userId
      })
    }
  }
})