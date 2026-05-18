# 网球预定系统

一个完整的网球场地预定系统,包含微信小程序端和Go后端服务。

## 技术栈

- **前端**: 微信小程序
- **后端**: Go + Gin框架 + GORM ORM
- **数据库**: MySQL 8.0+
- **认证**: JWT Token
- **支付**: 微信支付API v3

## 项目结构

```
booksystemTongYi/
├── database/              # 数据库脚本
│   └── schema.sql        # 数据库建表脚本
├── server/               # Go后端服务
│   ├── cmd/             # 应用入口
│   │   └── main.go      # 主程序
│   ├── api/             # API路由
│   │   └── routes.go    # 路由配置
│   ├── internal/        # 内部包
│   │   ├── config/      # 配置管理
│   │   ├── middleware/  # 中间件
│   │   ├── models/      # 数据模型
│   │   ├── handlers/    # HTTP处理器
│   │   ├── services/    # 业务逻辑
│   │   └── utils/       # 工具函数
│   ├── config.yaml.example  # 配置文件示例
│   └── go.mod           # Go模块文件
└── miniprogram/         # 微信小程序(待开发)
```

## 快速开始

### 1. 数据库设置

```bash
# 导入数据库脚本
mysql -u root -p < database/schema.sql
```

### 2. 后端配置

```bash
# 进入server目录
cd server

# 复制配置文件
cp config.yaml.example config.yaml

# 编辑config.yaml,修改数据库连接和微信配置
```

### 3. 安装依赖

```bash
cd server
go mod download
```

### 4. 运行后端服务

```bash
cd server
go run cmd/main.go
```

服务将在 `http://localhost:8080` 启动

### 5. 测试API

```bash
# 健康检查
curl http://localhost:8080/health

# 获取场地列表
curl http://localhost:8080/api/courts
```

## API接口文档

### 公开接口(无需认证)

- `POST /api/auth/wx-login` - 微信小程序登录
- `GET /api/courts` - 获取场地列表
- `GET /api/courts/:id` - 获取场地详情
- `GET /api/courts/:court_id/slots?date=2026-05-15` - 获取时间段
- `POST /api/payments/notify` - 微信支付回调

### 需要认证的接口

在请求头中添加: `Authorization: Bearer <token>`

#### 用户相关
- `GET /api/user/profile` - 获取用户信息
- `PUT /api/user/profile` - 更新用户信息

#### 预定相关
- `POST /api/bookings` - 创建预定
- `GET /api/bookings` - 获取我的预定列表
- `GET /api/bookings/:id` - 获取订单详情
- `PUT /api/bookings/:id/cancel` - 取消订单

#### 支付相关
- `POST /api/payments/create` - 创建支付订单
- `GET /api/payments/:booking_id` - 查询支付状态

### 管理员接口

需要管理员权限

- `POST /api/admin/login` - 管理员登录
- `POST /api/admin/courts` - 创建场地
- `PUT /api/admin/courts/:id` - 更新场地
- `DELETE /api/admin/courts/:id` - 删除场地
- `POST /api/admin/slots` - 批量生成时间段
- `PUT /api/admin/slots/:id` - 更新时间段状态
- `GET /api/admin/bookings` - 获取所有订单
- `GET /api/admin/dashboard` - 数据统计

## 配置说明

### config.yaml

```yaml
server:
  port: "8080"          # 服务端口
  mode: "debug"         # 运行模式: debug/release

database:
  host: "localhost"     # 数据库地址
  port: "3306"          # 数据库端口
  user: "root"          # 数据库用户名
  password: "root"      # 数据库密码
  dbname: "tennis_booking"  # 数据库名称
  charset: "utf8mb4"    # 字符集

jwt:
  secret: "your-secret-key"  # JWT密钥
  expire_hour: 72            # Token过期时间(小时)

wechat:
  app_id: "your_appid"       # 微信小程序AppID
  app_secret: "your_secret"  # 微信小程序AppSecret
  mch_id: "your_mch_id"      # 商户号
  api_key: "your_api_key"    # API密钥
```

## 数据库表说明

1. **users** - 用户表
2. **courts** - 场地表
3. **time_slots** - 时间段表
4. **bookings** - 预定订单表
5. **payments** - 支付记录表
6. **admins** - 管理员表
7. **system_configs** - 系统配置表

## 开发进度

- [x] 数据库设计与初始化
- [x] Go后端基础架构
- [x] 配置管理系统
- [x] JWT认证中间件
- [x] 用户模块API
- [x] 场地管理API(基础)
- [ ] 预定订单API
- [ ] 支付模块API
- [ ] 微信小程序开发
- [ ] 定时任务
- [ ] 消息通知
- [ ] 测试与部署

## 注意事项

1. **生产环境配置**: 
   - 修改JWT密钥为强随机字符串
   - 使用环境变量管理敏感配置
   - 启用HTTPS

2. **微信支付配置**:
   - 需要在微信公众平台申请小程序
   - 配置支付商户号
   - 设置支付回调地址

3. **安全性**:
   - 定期更新依赖包
   - 使用参数化查询防止SQL注入
   - 实施速率限制防止暴力攻击

## 常见问题

### Q: 如何重置管理员密码?
A: 直接修改数据库中admins表的password_hash字段,使用bcrypt加密新密码。

### Q: 微信登录失败?
A: 检查config.yaml中的app_id和app_secret是否正确配置。

### Q: 数据库连接失败?
A: 确认MySQL服务已启动,检查config.yaml中的数据库配置。

## 许可证

禁止商用

## 联系方式
如有问题,请提交Issue或联系开发者。
<img width="1396" height="948" alt="ScreenShot_2026-05-15_171837_753" src="https://github.com/user-attachments/assets/09813448-115c-44cd-9ada-4fa9da541714" />
<img width="1255" height="683" alt="ScreenShot_2026-05-15_171523_361" src="https://github.com/user-attachments/assets/5ffe3da3-928d-4b41-9097-b9d94a5a780a" />
<img width="1254" height="812" alt="ScreenShot_2026-05-15_171513_643" src="https://github.com/user-attachments/assets/2559300b-03f5-49b6-9f2e-312213130b81" />
<img width="891" height="790" alt="ScreenShot_2026-05-15_171435_919" src="https://github.com/user-attachments/assets/04010f6a-bce0-40bb-9952-633cc3267ab3" />
<img width="1253" height="666" alt="ScreenShot_2026-05-15_171402_305" src="https://github.com/user-attachments/assets/b70bf15e-4395-4073-8dbf-41586179b1ec" />



<img width="1258" height="834" alt="ScreenShot_2026-05-15_171347_523" src="https://github.com/user-attachments/assets/ebd2863f-4dee-425b-9b39-b30d4ae5ccd9" />
