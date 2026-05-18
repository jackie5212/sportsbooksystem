// pages/profile/profile.js
const app = getApp()

Page({
  data: {
    userInfo: null
  },

  onShow() {
    this.setData({
      userInfo: app.globalData.userInfo
    })
  },

  goToBookings() {
    wx.switchTab({
      url: '/pages/my-bookings/my-bookings'
    })
  },

  bindPhone() {
    wx.showToast({
      title: '功能开发中',
      icon: 'none'
    })
  },

  contactService() {
    wx.showToast({
      title: '客服功能开发中',
      icon: 'none'
    })
  },

  logout() {
    wx.showModal({
      title: '提示',
      content: '确定要退出登录吗?',
      success: (res) => {
        if (res.confirm) {
          app.clearLoginInfo()
          wx.redirectTo({
            url: '/pages/login/login'
          })
        }
      }
    })
  }
})
