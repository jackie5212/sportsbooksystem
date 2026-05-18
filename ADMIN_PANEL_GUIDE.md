# 🎾 网球预定系统 - 完整项目说明

## 📦 项目组成

本项目包含三个部分:

1. **Go后端API服务** - 提供数据接口
2. **微信小程序** - 用户端界面
3. **Vue 3管理后台** - PC端管理系统

---

## 🚀 快速启动指南

### 1️⃣ 启动后端服务 (端口: 8080)

```bash
cd server
go run simple_server.go
```

访问测试: http://localhost:8080/health

### 2️⃣ 启动管理后台 (端口: 3000)

```bash
cd admin-panel
npm run dev
```

访问地址: http://localhost:3000
- 用户名: admin
- 密码: admin123

### 3️⃣ 打开微信小程序

使用微信开发者工具导入 `miniprogram` 目录

---

## 📊 系统架构

```
┌──────────────┐         ┌──────────────┐         ┌──────────────┐
│  微信小程序   │         │ Vue管理后台   │         │  Go后端服务   │
│  (用户端)     │ ←HTTP→ │  (PC端)       │ ←HTTP→ │  (API服务器)  │
│  Port: N/A   │  JSON   │  Port: 3000  │  JSON   │  Port: 8080  │
└──────────────┘         └──────────────┘         └──────────────┘
                                                         ↓
                                                  ┌──────────────┐
                                                  │   MySQL      │
                                                  │  (待配置)     │
                                                  └──────────────┘
```

---

## ✨ 功能清单

### 🔹 后端服务 (Go + Gin)

**已实现:**
- ✅ RESTful API架构
- ✅ 健康检查接口
- ✅ 场地列表API
- ✅ 场地详情API  
- ✅ 时间段查询API
- ✅ CORS跨域支持
- ✅ 统一JSON响应格式

**待完善:**
- ⏳ 数据库连接(MySQL)
- ⏳ JWT认证中间件
- ⏳ 微信登录集成
- ⏳ 订单管理API
- ⏳ 支付接口

### 🔹 微信小程序 (原生开发)

**已完成页面:**
- ✅ 登录页 - 微信授权登录
- ✅ 首页 - 场地列表(下拉刷新/上拉加载)
- ✅ 场地详情 - 图片轮播/日期选择/时间网格
- ✅ 确认预定 - 订单信息展示
- ✅ 我的预定 - 订单列表/状态筛选
- ✅ 个人中心 - 用户信息/功能菜单

**核心功能:**
- ✅ 网络请求封装
- ✅ Token自动管理
- ✅ 全局状态管理
- ✅ 错误处理
- ✅ 登录状态检查

### 🔹 管理后台 (Vue 3 + Element Plus)

**已完成页面:**
- ✅ 登录页 - 管理员认证
- ✅ 仪表盘 - 数据统计图表
  - 统计卡片(场地数/订单数/用户数/收入)
  - 订单趋势图(ECharts折线图)
  - 场地使用率(ECharts饼图)
  - 最近订单列表
- ✅ 场地管理 - CRUD操作
  - 新增/编辑/删除场地
  - 启用/禁用场地
  - 搜索和分页
- ✅ 订单管理
  - 订单列表展示
  - 按状态筛选
  - 查看订单详情
  - 取消订单
- ✅ 用户管理
  - 用户列表
  - 查看用户信息
  - 预定次数统计
- ✅ 时间段管理
  - 时间段列表
  - 按场地和日期筛选
  - 锁定/解锁时间段

**技术特性:**
- ✅ Vue 3 Composition API
- ✅ Vue Router路由守卫
- ✅ Axios请求拦截器
- ✅ Element Plus UI组件
- ✅ ECharts数据可视化
- ✅ 响应式布局

---

## 📁 项目结构

```
booksystemTongYi/
├── database/                    # 数据库
│   └── schema.sql              # 建表脚本
│
├── server/                     # Go后端
│   ├── cmd/
│   │   └── main.go            # 完整版入口(未运行)
│   ├── api/
│   │   └── routes.go          # 路由配置
│   ├── internal/
│   │   ├── config/            # 配置管理
│   │   ├── models/            # 数据模型
│   │   ├── services/          # 业务逻辑
│   │   ├── handlers/          # HTTP处理器
│   │   └── middleware/        # 中间件
│   ├── simple_server.go       # 演示版(运行中) ✅
│   ├── go.mod
│   └── config.yaml.example
│
├── miniprogram/                # 微信小程序
│   ├── app.js                 # 小程序入口
│   ├── app.json               # 配置
│   ├── pages/
│   │   ├── login/             # 登录页
│   │   ├── index/             # 首页
│   │   ├── court-detail/      # 场地详情
│   │   ├── booking-confirm/   # 确认预定
│   │   ├── my-bookings/       # 我的预定
│   │   ├── profile/           # 个人中心
│   │   ├── payment/           # 支付页
│   │   └── booking-detail/    # 订单详情
│   └── utils/
│       └── request.js         # 网络请求
│
├── admin-panel/                # Vue管理后台
│   ├── src/
│   │   ├── views/
│   │   │   ├── Login.vue      # 登录页
│   │   │   ├── Layout.vue     # 主布局
│   │   │   ├── Dashboard.vue  # 仪表盘
│   │   │   ├── Courts.vue     # 场地管理
│   │   │   ├── Bookings.vue   # 订单管理
│   │   │   ├── Users.vue      # 用户管理
│   │   │   └── TimeSlots.vue  # 时间段管理
│   │   ├── router/            # 路由
│   │   ├── api/               # API接口
│   │   ├── App.vue
│   │   └── main.js
│   ├── index.html
│   ├── vite.config.js
│   └── package.json
│
└── 文档/
    ├── README.md
    ├── QUICKSTART.md
    ├── MINIPROGRAM_GUIDE.md
    ├── HOW_TO_VIEW_MINIPROGRAM.md
    ├── CURRENT_STATUS.md
    └── ADMIN_PANEL_GUIDE.md   # 本文档
```

---

## 🛠️ 技术栈

### 后端
- **语言**: Go 1.21.13
- **Web框架**: Gin v1.9.1
- **ORM**: GORM v1.25.0
- **配置**: Viper v1.16.0
- **日志**: Zap v1.24.0
- **认证**: JWT (golang-jwt/jwt/v5)

### 小程序
- **框架**: 微信小程序原生
- **语言**: JavaScript (ES6+)
- **样式**: WXSS
- **模板**: WXML

### 管理后台
- **框架**: Vue 3.4.0
- **构建工具**: Vite 5.0
- **UI库**: Element Plus 2.5.0
- **路由**: Vue Router 4.2.0
- **状态管理**: Pinia 2.1.0
- **HTTP**: Axios 1.6.0
- **图表**: ECharts 5.4.0

---

## 📝 当前状态

### ✅ 已完成
- [x] 数据库设计(7个表)
- [x] 后端API框架搭建
- [x] 后端演示服务(运行中)
- [x] 小程序UI开发(8个页面)
- [x] 小程序核心功能
- [x] 管理后台开发(7个页面)
- [x] 管理后台数据可视化
- [x] 完整文档编写

### ⏳ 待完善
- [ ] 后端连接MySQL数据库
- [ ] 实现完整的JWT认证
- [ ] 微信登录集成
- [ ] 订单创建API
- [ ] 微信支付集成
- [ ] 管理员登录API
- [ ] 前后端API对接
- [ ] 真机测试
- [ ] 性能优化
- [ ] 部署上线

---

## 🎯 使用场景

### 用户端(小程序)
1. 浏览网球场列表
2. 查看场地详情和图片
3. 选择日期和时间段
4. 提交预定订单
5. 查看我的预定记录
6. 个人中心管理

### 管理端(PC后台)
1. 查看系统数据统计
2. 管理场地信息(增删改查)
3. 查看所有订单
4. 管理用户信息
5. 控制时间段可用性
6. 监控收入和订单趋势

---

## 🔧 配置说明

### 后端配置
编辑 `server/config.yaml`:
```yaml
server:
  port: 8080
database:
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: tennis_booking
jwt:
  secret: your_jwt_secret
wechat:
  appid: your_appid
  secret: your_secret
```

### 管理后台配置
Vite代理已配置,自动转发 `/api` 到 `http://localhost:8080`

### 小程序配置
编辑 `miniprogram/app.js`:
```javascript
globalData: {
  baseURL: 'http://localhost:8080/api' // 开发环境
}
```

---

## 📈 下一步计划

### 短期(1-2周)
1. 安装MySQL并导入数据库
2. 完善后端API(订单/支付/认证)
3. 前后端API对接测试
4. 修复已知bug

### 中期(1个月)
1. 实现微信支付
2. 添加消息通知
3. 完善管理员权限
4. 添加评价系统

### 长期(3个月)
1. 性能优化和缓存
2. 数据统计分析
3. 优惠券系统
4. 会员体系
5. 多场馆支持

---

## 💡 常见问题

### Q1: 如何查看小程序界面?
A: 使用微信开发者工具导入 `miniprogram` 目录,选择测试号即可。

### Q2: 管理后台无法登录?
A: 当前使用模拟数据,直接输入 admin/admin123 即可登录。后续需接入真实API。

### Q3: 后端为什么用simple_server.go?
A: 由于Go模块系统问题,完整版暂时无法运行。simple_server.go是绕过问题的演示版本。

### Q4: 如何部署到生产环境?
A: 
- 后端: 编译成二进制文件,配置systemd服务
- 小程序: 提交审核发布
- 管理后台: npm run build生成静态文件,使用Nginx部署

---

## 📞 联系方式

如有问题,请查看相关文档或继续开发完善。

**项目完成度**: 约70%
**可用状态**: ✅ 演示版本可运行
**生产就绪**: ❌ 需要完善API和数据库

---

**最后更新**: 2026-05-15
**开发者**: AI Assistant
