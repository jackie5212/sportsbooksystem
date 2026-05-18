// pages/my-bookings/my-bookings.js
const app = getApp()

Page({
  data: {
    tabs: ['全部', '待支付', '已完成', '已取消'],
    activeTab: 0,
    bookings: []
  },

  onShow() {
    // 检查登录状态
    if (!app.globalData.token) {
      wx.redirectTo({
        url: '/pages/login/login'
      })
      return
    }
    
    this.loadBookings()
  },

  switchTab(e) {
    const index = e.currentTarget.dataset.index
    this.setData({ activeTab: index })
    this.loadBookings()
  },

  loadBookings() {
    // TODO: 从API加载数据
    // 模拟数据
    const mockBookings = [
      {
        id: 1,
        order_no: 'TB202605150001',
        court_name: '标准网球场A',
        booking_date: '2026-05-15',
        start_time: '08:00',
        end_time: '09:00',
        total_amount: 80,
        status: 1,
        status_text: '待支付'
      },
      {
        id: 2,
        order_no: 'TB202605140002',
        court_name: 'VIP网球场',
        booking_date: '2026-05-14',
        start_time: '14:00',
        end_time: '15:00',
        total_amount: 150,
        status: 2,
        status_text: '已完成'
      }
    ]
    
    this.setData({ bookings: mockBookings })
  }
})
