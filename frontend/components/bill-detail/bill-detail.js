const util = require("../../utils/util")

Component({
  data: {
      dateName: String, 
      countNumber: Number,
  },
  lifetimes: {
    ready:function() {
      console.log(this.data.item.length)
      var countNumber = 0;
      for (var i=0; i<this.data.item.length; i++) {
        countNumber += this.data.item[i].consumption
      }
      this.setData({
        countNumber: countNumber,
        dateName: util.getDateName(new Date(Date.parse(this.data.time)))
      })
    }
  },
  properties: {
    time: String,
    item: Object
  },
  methods: {}
})