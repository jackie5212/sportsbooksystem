// pages/booking-confirm/booking-confirm.js
const { request } = require('../../utils/request')
const app = getApp()

Page({
  data: {
    courtId: null,
    slotId: null,
    courtName: '',
    bookingDate: '',
    timeSlot: '',
    price: 0
  },

  onLoad(options) {
    this.setData({
      courtId: options.courtId,
      slotId: options.slotId
    })
    // TODO: 加载详细信息
    this.setData({
      courtName: '标准网球场A',
      bookingDate: '2026-05-15',
      timeSlot: '08:00-09:00',
      price: 80
    })
  },

  confirmOrder() {
    wx.showToast({
      title: '预定成功',
      icon: 'success'
    })
    
    setTimeout(() => {
      wx.switchTab({
        url: '/pages/my-bookings/my-bookings'
      })
    }, 1500)
  }
})
