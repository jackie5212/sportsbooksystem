# 快速启动指南

## 前置要求

1. **Go环境**: Go 1.21+
2. **MySQL**: MySQL 8.0+
3. **微信开发者工具**: 用于运行小程序
4. **Git**: 版本控制

## 第一步: 数据库设置

### 1. 创建数据库

```bash
# 登录MySQL
mysql -u root -p

# 执行建表脚本
source database/schema.sql
```

或者直接使用命令行:
```bash
mysql -u root -p < database/schema.sql
```

### 2. 验证数据库

```sql
USE tennis_booking;
SHOW TABLES;
```

应该看到以下表:
- users
- courts
- time_slots
- bookings
- payments
- admins
- system_configs

## 第二步: 后端配置

### 1. 进入server目录

```bash
cd server
```

### 2. 复制配置文件

```bash
cp config.yaml.example config.yaml
```

### 3. 编辑config.yaml

```yaml
server:
  port: "8080"
  mode: "debug"

database:
  host: "localhost"     # 修改为你的MySQL地址
  port: "3306"
  user: "root"          # 修改为你的MySQL用户名
  password: "root"      # 修改为你的MySQL密码
  dbname: "tennis_booking"
  charset: "utf8mb4"

jwt:
  secret: "tennis-booking-secret-key-2026-change-in-production"
  expire_hour: 72

wechat:
  app_id: ""            # 填入你的微信小程序AppID
  app_secret: ""        # 填入你的微信小程序AppSecret
  mch_id: ""            # 微信支付商户号(可选)
  api_key: ""           # 微信支付API密钥(可选)
```

**注意**: 
- 如果只是测试,可以暂时不填写微信相关配置
- JWT密钥在生产环境务必改为强随机字符串

### 4. 安装依赖

```bash
go mod download
```

### 5. 运行后端服务

```bash
go run cmd/main.go
```

如果看到以下输出,说明启动成功:
```
Server starting on :8080
Database connected successfully
```

### 6. 测试API

打开浏览器或使用curl测试:

```bash
# 健康检查
curl http://localhost:8080/health

# 获取场地列表
curl http://localhost:8080/api/courts
```

应该返回JSON数据。

## 第三步: 小程序配置

### 1. 打开微信开发者工具

- 下载并安装[微信开发者工具](https://developers.weixin.qq.com/miniprogram/dev/devtools/download.html)
- 使用微信扫码登录

### 2. 导入项目

- 点击"+"创建新项目
- 选择目录: `miniprogram`
- AppID: 使用测试号或填入你的正式AppID
- 项目名称: 网球预定

### 3. 配置服务器域名

在微信开发者工具中:
- 点击右上角"详情"
- 选择"本地设置"
- 勾选"不校验合法域名、web-view(业务域名)、TLS版本以及HTTPS证书"(仅开发时)

**生产环境需要**:
- 在微信公众平台配置服务器域名白名单
- 必须使用HTTPS

### 4. 修改API地址

编辑 `miniprogram/app.js`:

```javascript
globalData: {
  userInfo: null,
  token: null,
  baseURL: 'http://localhost:8080/api' // 开发环境
  // 生产环境改为: 'https://your-domain.com/api'
}
```

### 5. 编译运行

- 点击"编译"按钮
- 如果一切正常,应该能看到登录页面

## 第四步: 测试完整流程

### 1. 用户登录

- 点击"微信一键登录"按钮
- 授权获取用户信息
- 登录成功后跳转到首页

### 2. 浏览场地

- 首页显示场地列表
- 可以下拉刷新和上拉加载更多
- 点击场地卡片进入详情页

### 3. 查看场地详情

- 显示场地详细信息
- 可以选择日期查看可用时间段
- 点击"立即预定"按钮

## 常见问题

### Q1: 后端启动失败,提示数据库连接错误

**解决方案**:
1. 确认MySQL服务已启动
2. 检查config.yaml中的数据库配置是否正确
3. 确认数据库tennis_booking已创建

```bash
# 检查MySQL是否运行
mysqladmin -u root -p status

# 手动创建数据库
mysql -u root -p -e "CREATE DATABASE tennis_booking DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

### Q2: 小程序登录失败

**解决方案**:
1. 确认后端服务正在运行
2. 检查app.js中的baseURL是否正确
3. 查看控制台错误信息
4. 如果是跨域问题,确保后端CORS中间件已启用

### Q3: 微信登录提示AppID无效

**解决方案**:
1. 在config.yaml中配置正确的AppID和AppSecret
2. 或在微信公众平台申请测试号
3. 开发阶段可以使用测试号

### Q4: 如何生成测试数据?

执行以下SQL插入测试数据:

```sql
USE tennis_booking;

-- 插入测试场地
INSERT INTO courts (name, location, description, price_per_hour, status, images) VALUES
('测试场地1', '北京市朝阳区', '标准硬地网球场', 80.00, 1, '["https://via.placeholder.com/400x300"]'),
('测试场地2', '北京市海淀区', '红土网球场', 100.00, 1, '["https://via.placeholder.com/400x300"]');
```

### Q5: 端口被占用

**解决方案**:
修改config.yaml中的端口:
```yaml
server:
  port: "8081"  # 改为其他可用端口
```

## 开发建议

### 后端开发

```bash
# 热重载开发(需要安装air)
go install github.com/air-verse/air@latest
air

# 构建可执行文件
go build -o tennis-booking cmd/main.go
./tennis-booking
```

### 小程序调试

- 使用微信开发者工具的"调试器"查看Console和Network
- 使用"模拟器"预览效果
- 使用真机预览检查兼容性

## 下一步

系统基础框架已完成,接下来可以:

1. **完善预定功能**: 实现创建订单、支付等核心业务
2. **管理员后台**: 开发场地管理、订单管理等功能
3. **优化UI/UX**: 美化界面,提升用户体验
4. **添加更多功能**: 如评价系统、会员系统等
5. **部署上线**: 配置服务器、域名、SSL证书等

## 技术支持

遇到问题可以:
1. 查看README.md文档
2. 检查控制台错误日志
3. 提交Issue到GitHub仓库

祝开发顺利! 🎾
