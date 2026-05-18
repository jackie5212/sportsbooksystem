// pages/court-detail/court-detail.js
const { request } = require('../../utils/request')

Page({
  data: {
    courtId: null,
    court: {},
    selectedDate: '',
    dateList: [],
    slots: [],
    selectedSlot: null
  },

  onLoad(options) {
    const courtId = options.id
    this.setData({ courtId })
    this.loadCourtDetail(courtId)
    this.generateDateList()
  },

  // 加载场地详情
  async loadCourtDetail(courtId) {
    try {
      const court = await request({
        url: `/courts/${courtId}`,
        method: 'GET'
      })
      this.setData({ court })
    } catch (error) {
      console.error('加载场地详情失败:', error)
    }
  },

  // 生成日期列表
  generateDateList() {
    const dates = []
    const today = new Date()
    const weeks = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']

    for (let i = 0; i < 7; i++) {
      const date = new Date(today)
      date.setDate(today.getDate() + i)
      
      const year = date.getFullYear()
      const month = date.getMonth() + 1
      const day = date.getDate()
      const week = weeks[date.getDay()]

      dates.push({
        date: `${year}-${month.toString().padStart(2, '0')}-${day.toString().padStart(2, '0')}`,
        day: i === 0 ? '今天' : `${month}/${day}`,
        week: week
      })
    }

    this.setData({
      dateList: dates,
      selectedDate: dates[0].date
    })

    this.loadTimeSlots(dates[0].date)
  },

  // 选择日期
  selectDate(e) {
    const date = e.currentTarget.dataset.date
    this.setData({ selectedDate: date })
    this.loadTimeSlots(date)
  },

  // 加载时间段
  async loadTimeSlots(date) {
    try {
      const slots = await request({
        url: `/courts/${this.data.courtId}/slots`,
        method: 'GET',
        data: { date }
      })
      this.setData({ slots })
    } catch (error) {
      console.error('加载时间段失败:', error)
    }
  },

  // 选择时间段
  selectSlot(e) {
    const slot = e.currentTarget.dataset.slot
    if (slot.status === 1) {
      this.setData({ selectedSlot: slot })
    }
  },

  // 确认预定
  confirmBooking() {
    if (!this.data.selectedSlot) {
      wx.showToast({
        title: '请选择时间段',
        icon: 'none'
      })
      return
    }

    wx.navigateTo({
      url: `/pages/booking-confirm/booking-confirm?courtId=${this.data.courtId}&slotId=${this.data.selectedSlot.id}`
    })
  }
})
