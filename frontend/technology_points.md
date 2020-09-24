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
>scroll-view组件不能在该组件上定义class
***
### 小程序问题
>自定义顶部导航栏，与原生顶部导航栏字体样式在微信开发者工具中，显示一致，但是在真机环境下不一致。
> 如font-weight