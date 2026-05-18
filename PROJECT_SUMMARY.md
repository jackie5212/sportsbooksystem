# 网球预定系统 - 项目总结

## 项目概述

这是一个完整的网球场地预定系统,包含微信小程序前端和Go语言后端服务。用户可以浏览场地、选择时间段进行预定、在线支付,管理员可以管理场地和订单。

## 已完成功能

### ✅ 后端 (Go + Gin)

#### 1. 基础架构
- [x] Go项目初始化
- [x] Gin Web框架搭建
- [x] GORM ORM集成
- [x] MySQL数据库连接池
- [x] JWT认证中间件
- [x] CORS跨域支持
- [x] 统一响应格式
- [x] 配置管理系统(Viper)
- [x] 日志系统(Zap)

#### 2. 数据模型
- [x] User (用户表)
- [x] Court (场地表)
- [x] TimeSlot (时间段表)
- [x] Booking (订单表)
- [x] Payment (支付表)
- [x] Admin (管理员表)
- [x] SystemConfig (系统配置表)

#### 3. API接口

**公开接口:**
- [x] POST /api/auth/wx-login - 微信登录
- [x] GET /api/courts - 获取场地列表(分页)
- [x] GET /api/courts/:id - 获取场地详情
- [x] GET /api/courts/:court_id/slots - 获取时间段

**用户接口(需认证):**
- [x] GET /api/user/profile - 获取用户信息
- [x] PUT /api/user/profile - 更新用户信息
- [ ] POST /api/bookings - 创建预定
- [ ] GET /api/bookings - 获取我的预定
- [ ] GET /api/bookings/:id - 获取订单详情
- [ ] PUT /api/bookings/:id/cancel - 取消订单
- [ ] POST /api/payments/create - 创建支付
- [ ] GET /api/payments/:booking_id - 查询支付状态

**管理员接口(需管理员权限):**
- [ ] POST /api/admin/login - 管理员登录
- [ ] POST /api/admin/courts - 创建场地
- [ ] PUT /api/admin/courts/:id - 更新场地
- [ ] DELETE /api/admin/courts/:id - 删除场地
- [ ] POST /api/admin/slots - 批量生成时间段
- [ ] PUT /api/admin/slots/:id - 更新时间段
- [ ] GET /api/admin/bookings - 获取所有订单
- [ ] GET /api/admin/dashboard - 数据统计

**回调接口:**
- [ ] POST /api/payments/notify - 微信支付回调

### ✅ 微信小程序

#### 1. 项目结构
- [x] app.json配置
- [x] app.js全局逻辑
- [x] app.wxss全局样式
- [x] 网络请求封装(utils/request.js)
- [x] TabBar配置(首页、我的预定、我的)

#### 2. 页面
- [x] 登录页 (pages/login/login)
  - 微信一键登录
  - 用户授权
  - Token管理
  
- [x] 首页 (pages/index/index)
  - 场地列表展示
  - 下拉刷新
  - 上拉加载更多
  - 搜索功能(UI已实现)
  - 跳转到详情页

- [ ] 场地详情页 (pages/court-detail/court-detail)
- [ ] 预定确认页 (pages/booking-confirm/booking-confirm)
- [ ] 支付页 (pages/payment/payment)
- [ ] 我的预定页 (pages/my-bookings/my-bookings)
- [ ] 订单详情页 (pages/booking-detail/booking-detail)
- [ ] 个人中心页 (pages/profile/profile)

### ✅ 数据库

- [x] 完整的SQL建表脚本
- [x] 索引优化
- [x] 外键约束
- [x] 初始数据(示例场地、管理员)
- [x] 系统配置数据

### ✅ 文档

- [x] README.md - 项目说明文档
- [x] QUICKSTART.md - 快速启动指南
- [x] .gitignore - Git忽略配置
- [x] start.sh - Linux/Mac启动脚本
- [x] start.bat - Windows启动脚本
- [x] config.yaml.example - 配置文件示例

## 技术栈

### 后端
- **语言**: Go 1.21+
- **Web框架**: Gin v1.12.0
- **ORM**: GORM v1.31.1
- **数据库**: MySQL 8.0+
- **认证**: JWT (golang-jwt/jwt/v5)
- **配置**: Viper v1.21.0
- **日志**: Zap v1.28.0
- **加密**: bcrypt (golang.org/x/crypto)

### 前端
- **框架**: 微信小程序原生
- **语言**: JavaScript (ES6+)
- **样式**: WXSS
- **网络**: wx.request封装

### 开发工具
- **版本控制**: Git
- **包管理**: Go Modules
- **IDE**: VS Code / GoLand / 微信开发者工具

## 项目结构

```
booksystemTongYi/
├── database/
│   └── schema.sql              # 数据库脚本
├── server/                     # Go后端
│   ├── cmd/
│   │   └── main.go            # 入口文件
│   ├── api/
│   │   └── routes.go          # 路由配置
│   ├── internal/
│   │   ├── config/            # 配置
│   │   │   ├── config.go
│   │   │   └── database.go
│   │   ├── middleware/        # 中间件
│   │   │   ├── auth.go
│   │   │   └── cors.go
│   │   ├── models/            # 数据模型
│   │   │   └── models.go
│   │   ├── handlers/          # HTTP处理器
│   │   │   ├── user_handler.go
│   │   │   └── court_handler.go
│   │   ├── services/          # 业务逻辑
│   │   │   └── user_service.go
│   │   └── utils/             # 工具函数
│   │       └── response.go
│   ├── config.yaml.example    # 配置示例
│   └── go.mod                 # Go模块
├── miniprogram/               # 微信小程序
│   ├── pages/
│   │   ├── login/            # 登录页
│   │   └── index/            # 首页
│   ├── utils/
│   │   └── request.js        # 网络请求
│   ├── app.js                # 应用入口
│   ├── app.json              # 应用配置
│   ├── app.wxss              # 全局样式
│   ├── project.config.json   # 项目配置
│   └── sitemap.json          # 站点地图
├── README.md                  # 项目说明
├── QUICKSTART.md             # 快速启动
├── .gitignore                # Git忽略
├── start.sh                  # Linux启动脚本
└── start.bat                 # Windows启动脚本
```

## 核心功能流程

### 1. 用户登录流程
```
小程序端                          后端                         微信服务器
  |                                |                              |
  |-- wx.login() --> code ---------|                              |
  |                                |-- code + AppSecret --------->|
  |                                |<-- openid + session_key -----|
  |                                |                              |
  |-- 用户信息 + code ------------->|                              |
  |                                |-- 查找/创建用户              |
  |                                |-- 生成JWT Token              |
  |<-- Token + 用户信息 -----------|                              |
  |                                |                              |
  |-- 保存Token到本地存储          |                              |
```

### 2. 场地预定流程(待实现)
```
用户选择场地 -> 选择日期和时间段 -> 确认订单 -> 调用支付API -> 
微信支付 -> 支付回调 -> 更新订单状态 -> 通知用户
```

## 数据库设计亮点

1. **防超卖设计**: time_slots表使用唯一索引(court_id, date, start_time)防止重复预定
2. **订单追踪**: bookings表记录完整订单信息,支持状态流转
3. **支付记录**: payments表独立存储,便于对账和问题排查
4. **软删除**: 使用status字段而非物理删除,保留历史数据
5. **时间戳**: 所有表都有created_at和updated_at自动维护

## 安全特性

1. **JWT认证**: 无状态认证,支持分布式部署
2. **密码加密**: 管理员密码使用bcrypt加密
3. **CORS控制**: 跨域请求严格控制
4. **参数验证**: 使用Gin的binding进行参数校验
5. **SQL防注入**: GORM使用参数化查询

## 性能优化

1. **数据库索引**: 关键字段建立索引加速查询
2. **连接池**: MySQL连接池复用连接
3. **分页查询**: 列表接口支持分页,避免大数据量
4. **懒加载**: GORM关联查询按需加载

## 待完成功能

### 高优先级
1. **预定订单API**: 创建订单、查询订单、取消订单
2. **支付集成**: 微信支付统一下单、回调处理
3. **场地详情页**: 日历组件、时间段选择
4. **预定流程**: 完整的下单和支付流程
5. **我的预定**: 订单列表和详情展示

### 中优先级
6. **管理员后台**: 场地管理、订单管理
7. **定时任务**: 订单超时自动取消
8. **消息通知**: 模板消息推送
9. **数据统计**: 订单统计、收入统计

### 低优先级
10. **评价系统**: 用户对场地评价
11. **会员系统**: 积分、优惠券
12. **缓存优化**: Redis缓存热点数据
13. **搜索引擎**: Elasticsearch场地搜索

## 部署建议

### 开发环境
- 本地运行MySQL
- 后端localhost:8080
- 小程序使用测试号

### 测试环境
- 云服务器安装MySQL
- 域名解析 + Nginx反向代理
- HTTPS证书(Let's Encrypt)
- 小程序配置正式AppID

### 生产环境
- 数据库主从复制
- 后端多实例负载均衡
- Redis缓存层
- CDN加速静态资源
- 监控告警(Prometheus + Grafana)
- 日志收集(ELK)

## 扩展方向

1. **多运动类型**: 支持羽毛球、乒乓球等其他场地
2. **社交功能**: 约球、组队、排行榜
3. **智能推荐**: 基于用户偏好推荐场地
4. **数据分析**: 用户行为分析、业务报表
5. **物联网集成**: 智能门禁、灯光控制
6. **多端支持**: H5、Android、iOS App

## 学习价值

本项目适合学习:
- Go Web开发最佳实践
- 微信小程序开发
- RESTful API设计
- JWT认证机制
- 微信支付集成
- 数据库设计与优化
- 前后端分离架构

## 总结

本项目已完成基础框架搭建,包括:
- ✅ 完整的Go后端架构
- ✅ 数据库设计与初始化
- ✅ 用户认证系统
- ✅ 场地查询功能
- ✅ 小程序登录和首页

后续可以继续完善预定、支付等核心业务功能,即可成为一个可用的网球预定系统。代码结构清晰,易于扩展和维护。

---

**开发时间**: 2026年5月  
**版本**: v0.1.0 (基础版)  
**许可证**: MIT
