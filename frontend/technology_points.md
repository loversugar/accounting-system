### 配置
>自定义顶部title  
>分为全局自定义，和局部自定义。分别修改的是app.json中的window标签模块和局部模块的.json文件  
>都是在当中添加***"navigationStyle": "custom"***,具体配置请看[官网](https://developers.weixin.qq.com/miniprogram/dev/reference/configuration/page.html)
>顶部整个导航栏=状态栏+标题栏，状态栏高度获取[官网](https://developers.weixin.qq.com/miniprogram/dev/api/base/system/system-info/wx.getSystemInfo.html). 
>Android顶部整个导航栏=32px + 8px * 2 = 48px
>IOS顶部整个导航栏=32px + 6px * 2 = 44px
***
>getCurrentPages()获取当前页面属性
>switchTab跳转到tabBar页面，关闭其他***非tabBar***页面 reLanuch关闭所有页面，打开应用中某个页面能触发onLoad,但是前者不能
***
>scroll-view组件不能在该组件上定义class
***
>微信小程序组件之间的通信，分为三种   
> 1. 数据绑定，用于父组件向子组件传递数据  
> 2. 事件(triggerEvent),用于子组件向父组件传递数据
> 3. 以上两种不满足，父组件可以通过this.selectComponent方法获取子组件实例
***
> 在wxml中为方法传入参数，需要用data-xxx={{item}}的形式来获取
***
> div中让内容文字居中，text-align:center 水平居中，padding: 50px 20px垂直居中(高度不固定的情况下)
### 小程序问题
>自定义顶部导航栏，与原生顶部导航栏字体样式在微信开发者工具中，显示一致，但是在真机环境下不一致。
> 如font-weight


### 自己傻逼时刻
> 自定义组件中，data和properties中数据都可以通过this.data来获取