Component({
  data: {
    iconfontBackgroundColor: "#f1f1f1"
  },
  properties: {
    width: Number,
    iconfontBackgroundColor: {
      type: String,
      value: "#f1f1f1"
    },
    tag: Object,
    isClick: {
      type: Boolean,
      value: false
    }
  },
  methods: {
    click() {
      // if (!this.data.isClick) {
      //   this.setData({
      //     isClick: true,
      //     iconfontBackgroundColor: "#00b38a"
      //   })
      // } else {
      //   this.setData({
      //     isClick: false,
      //     iconfontBackgroundColor: "#f1f1f1"
      //   })
      // }
      this.triggerEvent('clickOne', {
        tagId: this.data.tag.id,
      })
    }
  }
})