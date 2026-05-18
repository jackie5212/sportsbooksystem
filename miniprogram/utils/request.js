// utils/request.js
const app = getApp()

/**
 * 封装网络请求
 */
function request(options) {
  return new Promise((resolve, reject) => {
    const {
      url,
      method = 'GET',
      data = {},
      header = {}
    } = options

    // 添加基础URL
    const fullUrl = url.startsWith('http') ? url : `${app.globalData.baseURL}${url}`

    // 添加认证头
    if (app.globalData.token) {
      header['Authorization'] = `Bearer ${app.globalData.token}`
    }

    wx.request({
      url: fullUrl,
      method,
      data,
      header: {
        'Content-Type': 'application/json',
        ...header
      },
      success(res) {
        if (res.statusCode === 200) {
          if (res.data.code === 200) {
            resolve(res.data.data)
          } else {
            wx.showToast({
              title: res.data.message || '请求失败',
              icon: 'none'
            })
            reject(new Error(res.data.message))
          }
        } else if (res.statusCode === 401) {
          // Token过期,跳转登录
          wx.showToast({
            title: '登录已过期',
            icon: 'none'
          })
          app.clearLoginInfo()
          wx.redirectTo({
            url: '/pages/login/login'
          })
          reject(new Error('未授权'))
        } else {
          wx.showToast({
            title: '网络错误',
            icon: 'none'
          })
          reject(new Error('网络错误'))
        }
      },
      fail(err) {
        wx.showToast({
          title: '网络连接失败',
          icon: 'none'
        })
        reject(err)
      }
    })
  })
}

module.exports = {
  request
}
