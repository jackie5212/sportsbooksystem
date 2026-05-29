import request from './request'

// 管理员登录
export function adminLogin(data) {
  return request({
    url: '/admin/login',
    method: 'post',
    data
  })
}

// 获取统计数据
export function getStatistics() {
  return request({
    url: '/admin/statistics',
    method: 'get'
  })
}

// 获取场地列表
export function getCourts(params) {
  return request({
    url: '/admin/courts',
    method: 'get',
    params
  })
}

// 创建场地
export function createCourt(data) {
  return request({
    url: '/admin/courts',
    method: 'post',
    data
  })
}

// 更新场地
export function updateCourt(id, data) {
  return request({
    url: `/admin/courts/${id}`,
    method: 'put',
    data
  })
}

// 删除场地
export function deleteCourt(id) {
  return request({
    url: `/admin/courts/${id}`,
    method: 'delete'
  })
}

// 获取订单列表
export function getBookings(params) {
  return request({
    url: '/admin/bookings',
    method: 'get',
    params
  })
}

// 获取用户列表
export function getUsers(params) {
  return request({
    url: '/admin/users',
    method: 'get',
    params
  })
}

// 获取时间段列表
export function getTimeSlots(params) {
  return request({
    url: '/admin/timeslots',
    method: 'get',
    params
  })
}

// 更新时间段状态
export function updateTimeSlotStatus(id, status) {
  return request({
    url: `/admin/timeslots/${id}/status`,
    method: 'put',
    data: { status }
  })
}

// 获取系统配置
export function getSettings() {
  return request({
    url: '/admin/settings',
    method: 'get'
  })
}

// 更新系统配置
export function updateSettings(data) {
  return request({
    url: '/admin/settings',
    method: 'put',
    data
  })
}

// 获取微信支付配置（脱敏）
export function getWechatConfig() {
  return request({
    url: '/admin/settings/wechat',
    method: 'get'
  })
}
