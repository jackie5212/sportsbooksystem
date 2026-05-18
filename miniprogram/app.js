// app.js
App({
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 检查登录状态
    this.checkLoginStatus()
  },
  
  globalData: {
    userInfo: null,
    token: null,
    // baseURL: 'http://localhost:8080/api' // 开发环境本地测试
    baseURL: 'http://192.168.0.101:8080/api' // 局域网访问
  },
  
  // 检查登录状态
  checkLoginStatus() {
    const token = wx.getStorageSync('token')
    const userInfo = wx.getStorageSync('userInfo')
    
    if (token && userInfo) {
      this.globalData.token = token
      this.globalData.userInfo = userInfo
    }
  },
  
  // 设置Token
  setToken(token) {
    this.globalData.token = token
    wx.setStorageSync('token', token)
  },
  
  // 设置用户信息
  setUserInfo(userInfo) {
    this.globalData.userInfo = userInfo
    wx.setStorageSync('userInfo', userInfo)
  },
  
  // 清除登录信息
  clearLoginInfo() {
    this.globalData.token = null
    this.globalData.userInfo = null
    wx.removeStorageSync('token')
    wx.removeStorageSync('userInfo')
  }
})
