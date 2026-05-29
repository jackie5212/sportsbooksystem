# 网球预定系统 - PC管理后台

## 📋 功能特性

- ✅ 管理员登录认证
- ✅ 数据仪表盘(统计图表)
- ✅ 场地管理(CRUD)
- ✅ 订单管理(查看/取消)
- ✅ 用户管理
- ✅ 时间段管理(锁定/解锁)

## 🚀 快速启动

### 1. 安装依赖

```bash
cd admin-panel
npm install
```

### 2. 启动开发服务器

```bash
npm run dev
```

服务将运行在: http://localhost:3000

### 3. 访问管理后台

浏览器打开: http://localhost:3000

**默认账号:**
- 用户名: admin
- 密码: admin123

## 📁 项目结构

```
admin-panel/
├── src/
│   ├── views/           # 页面组件
│   │   ├── Login.vue    # 登录页
│   │   ├── Layout.vue   # 主布局
│   │   ├── Dashboard.vue # 仪表盘
│   │   ├── Courts.vue   # 场地管理
│   │   ├── Bookings.vue # 订单管理
│   │   ├── Users.vue    # 用户管理
│   │   └── TimeSlots.vue # 时间段管理
│   ├── components/      # 公共组件
│   ├── router/          # 路由配置
│   ├── api/             # API接口
│   ├── store/           # 状态管理
│   ├── utils/           # 工具函数
│   ├── App.vue          # 根组件
│   └── main.js          # 入口文件
├── index.html           # HTML模板
├── vite.config.js       # Vite配置
└── package.json         # 依赖配置
```

## 🛠️ 技术栈

- **框架**: Vue 3 (Composition API)
- **构建工具**: Vite 5
- **UI组件库**: Element Plus
- **路由**: Vue Router 4
- **状态管理**: Pinia
- **HTTP客户端**: Axios
- **图表**: ECharts 5

## 📊 页面说明

### 1. 登录页 (/login)
- 管理员身份验证
- JWT Token存储

### 2. 仪表盘 (/dashboard)
- 数据统计卡片(场地数/订单数/用户数/收入)
- 订单趋势图(折线图)
- 场地使用率(饼图)
- 最近订单列表

### 3. 场地管理 (/courts)
- 场地列表展示
- 新增/编辑/删除场地
- 启用/禁用场地
- 搜索和分页

### 4. 订单管理 (/bookings)
- 订单列表
- 按状态筛选
- 查看订单详情
- 取消订单

### 5. 用户管理 (/users)
- 用户列表
- 查看用户信息
- 预定次数统计

### 6. 时间段管理 (/timeslots)
- 时间段列表
- 按场地和日期筛选
- 锁定/解锁时间段

## 🔌 API集成

当前使用模拟数据,需要与后端API对接:

1. 修改 `src/api/admin.js` 中的接口地址
2. 确保后端服务运行在 http://localhost:8080
3. Vite代理已配置,自动转发 `/api` 请求到后端

## 🎨 UI设计

- **主题色**: #409EFF (Element Plus默认蓝)
- **侧边栏**: 深色背景 (#304156)
- **布局**: 响应式设计
- **图标**: Element Plus Icons

## 📝 开发注意事项

1. **路由守卫**: 未登录自动跳转到登录页
2. **Token管理**: 存储在localStorage
3. **错误处理**: 统一的Axios拦截器
4. **权限控制**: 基于Token验证

## 🔧 后续优化

- [ ] 接入真实后端API
- [ ] 添加权限管理(角色/权限)
- [ ] 完善表单验证
- [ ] 添加数据导出功能
- [ ] 实现消息通知
- [ ] 添加操作日志
- [ ] 优化移动端适配

---

**开发完成时间**: 2026-05-15
