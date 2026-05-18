// pages/login/login.js
const app = getApp()
const { request } = require('../../utils/request')

Page({
  data: {
    loading: false
  },

  onLoad() {
    // 检查是否已登录
    if (app.globalData.token) {
      wx.switchTab({
        url: '/pages/index/index'
      })
    }
  },

  // 获取用户信息并登录
  async onGetUserInfo(e) {
    if (e.detail.userInfo) {
      this.setData({ loading: true })
      
      try {
        // ========== 模拟登录开始 ==========
        // 开发环境使用模拟数据,避免依赖后端服务
        console.log('使用模拟登录模式')
        
        // 模拟延迟
        await new Promise(resolve => setTimeout(resolve, 800))
        
        // 模拟用户数据
        const mockUserInfo = {
          id: 1,
          openid: 'mock_openid_' + Date.now(),
          nickname: e.detail.userInfo.nickName || '测试用户',
          avatar: e.detail.userInfo.avatarUrl || '',
          phone: '138****8888',
          gender: e.detail.userInfo.gender || 0,
          country: e.detail.userInfo.country || '中国',
          province: e.detail.userInfo.province || '',
          city: e.detail.userInfo.city || '',
          created_at: new Date().toISOString()
        }
        
        // 模拟Token
        const mockToken = 'mock_token_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9)
        
        // 保存登录信息
        app.setToken(mockToken)
        app.setUserInfo(mockUserInfo)
        
        console.log('模拟登录成功:', mockUserInfo)
        // ========== 模拟登录结束 ==========
        
        /* ========== 真实登录代码(已注释) ==========
        // 获取微信登录code
        const loginRes = await this.getWxLoginCode()
        
        if (!loginRes.code) {
          throw new Error('获取登录code失败')
        }

        // 调用后端登录接口
        const userInfo = e.detail.userInfo
        const result = await request({
          url: '/auth/wx-login',
          method: 'POST',
          data: {
            code: loginRes.code,
            nickname: userInfo.nickName,
            avatar: userInfo.avatarUrl,
            gender: userInfo.gender,
            country: userInfo.country,
            province: userInfo.province,
            city: userInfo.city
          }
        })

        // 保存登录信息
        app.setToken(result.token)
        app.setUserInfo(result.user_info)
        ========== 真实登录代码结束 ========== */

        wx.showToast({
          title: '登录成功',
          icon: 'success'
        })

        // 跳转到首页
        setTimeout(() => {
          wx.switchTab({
            url: '/pages/index/index'
          })
        }, 1500)

      } catch (error) {
        console.error('登录失败:', error)
        wx.showToast({
          title: '登录失败,请重试',
          icon: 'none'
        })
      } finally {
        this.setData({ loading: false })
      }
    } else {
      wx.showToast({
        title: '需要授权才能登录',
        icon: 'none'
      })
    }
  },

  // 获取微信登录code
  getWxLoginCode() {
    return new Promise((resolve, reject) => {
      wx.login({
        success: resolve,
        fail: reject
      })
    })
  }
})
