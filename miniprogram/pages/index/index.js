// pages/index/index.js
const { request } = require('../../utils/request')
const app = getApp()

Page({
  data: {
    courts: [],
    searchKeyword: '',
    page: 1,
    pageSize: 10,
    loading: false,
    hasMore: true
  },

  onLoad() {
    this.loadCourts()
  },

  onShow() {
    // 不再强制要求登录，允许游客浏览场地
    // 只有在预定时才需要登录
  },

  // 加载场地列表
  async loadCourts(refresh = false) {
    if (this.data.loading) return
    
    this.setData({ loading: true })

    try {
      const page = refresh ? 1 : this.data.page
      const result = await request({
        url: '/courts',
        method: 'GET',
        data: {
          page: page,
          page_size: this.data.pageSize
        }
      })

      const courts = refresh ? result.list : [...this.data.courts, ...result.list]
      
      this.setData({
        courts,
        page: page + 1,
        hasMore: courts.length < result.total,
        loading: false
      })

    } catch (error) {
      console.error('加载场地失败:', error)
      this.setData({ loading: false })
    }
  },

  // 下拉刷新
  onPullDownRefresh() {
    this.loadCourts(true)
    setTimeout(() => {
      wx.stopPullDownRefresh()
    }, 1000)
  },

  // 加载更多
  loadMore() {
    if (this.data.hasMore && !this.data.loading) {
      this.loadCourts()
    }
  },

  // 搜索输入
  onSearchInput(e) {
    this.setData({
      searchKeyword: e.detail.value
    })
    // TODO: 实现搜索功能
  },

  // 跳转到详情页
  goToDetail(e) {
    const courtId = e.currentTarget.dataset.id
    wx.navigateTo({
      url: `/pages/court-detail/court-detail?id=${courtId}`
    })
  }
})
