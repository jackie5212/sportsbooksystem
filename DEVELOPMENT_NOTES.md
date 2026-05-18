# 开发注意事项

## 环境准备

### 1. Go环境
- 确保Go版本 >= 1.21
- 配置GOPROXY加速依赖下载:
```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

### 2. MySQL配置
- 版本: MySQL 8.0+
- 字符集: utf8mb4
- 时区: 设置为Asia/Shanghai

```sql
-- 检查时区
SELECT @@time_zone;

-- 设置时区
SET GLOBAL time_zone = '+8:00';
```

### 3. 微信小程序
- 注册微信公众平台账号
- 获取AppID和AppSecret
- 配置服务器域名(生产环境)

## 开发规范

### Go代码规范

1. **命名规范**
   - 使用驼峰命名: `userName`, `getUserInfo`
   - 导出函数/变量首字母大写: `GetUser`, `Config`
   - 私有函数/变量首字母小写: `parseToken`, `db`

2. **错误处理**
```go
// 推荐方式
if err != nil {
    log.Printf("Error: %v", err)
    return err
}

// 不要忽略错误
result, err := doSomething()
if err != nil {
    // 处理错误
}
```

3. **注释规范**
```go
// GetUser 获取用户信息
// 参数: userID - 用户ID
// 返回: 用户对象和错误信息
func GetUser(userID uint64) (*User, error) {
    // ...
}
```

4. **项目结构**
```
internal/
├── config/      # 配置相关
├── middleware/  # 中间件
├── models/      # 数据模型
├── handlers/    # HTTP处理器(控制器)
├── services/    # 业务逻辑层
└── utils/       # 工具函数
```

### 小程序代码规范

1. **文件命名**
   - 页面文件放在pages目录下
   - 组件文件放在components目录下
   - 每个页面包含: .js, .json, .wxml, .wxss

2. **数据绑定**
```javascript
// 推荐: 使用setData更新数据
this.setData({
  userName: '张三',
  userList: [...this.data.userList, newItem]
})

// 避免: 直接修改data
this.data.userName = '张三' // 不推荐
```

3. **网络请求**
```javascript
// 使用封装的request方法
const { request } = require('../../utils/request')

async function loadData() {
  try {
    const data = await request({
      url: '/api/data',
      method: 'GET'
    })
    // 处理数据
  } catch (error) {
    console.error('请求失败:', error)
  }
}
```

## 常见问题与解决方案

### 1. 数据库连接失败

**症状**: 
```
Failed to connect database: dial tcp [::1]:3306: connectex: No connection could be made
```

**解决**:
- 检查MySQL服务是否启动
- 确认config.yaml中的host、port、user、password正确
- 测试命令行连接: `mysql -u root -p`

### 2. JWT Token验证失败

**症状**: 
```
401 Unauthorized: 无效的认证令牌
```

**解决**:
- 检查Token格式: `Bearer <token>`
- 确认JWT密钥(config.yaml中的jwt.secret)前后端一致
- 检查Token是否过期

### 3. 跨域问题

**症状**: 
```
Access to XMLHttpRequest has been blocked by CORS policy
```

**解决**:
- 确保后端CORS中间件已启用
- 检查请求头是否正确

### 4. 小程序真机调试失败

**症状**: 
```
request:fail url not in domain list
```

**解决**:
- 开发阶段: 勾选"不校验合法域名"
- 生产环境: 在微信公众平台配置服务器域名

### 5. 微信支付配置

**必要配置**:
1. 申请微信支付商户号
2. 配置API密钥(32位字符串)
3. 设置支付回调地址
4. 下载商户证书

**测试建议**:
- 先使用沙箱环境测试
- 小额测试(0.01元)
- 检查回调签名验证

## 性能优化建议

### 后端优化

1. **数据库优化**
```go
// 使用索引
db.Model(&User{}).Where("openid = ?", openID).First(&user)

// 批量插入
db.CreateInBatches(users, 100)

// 分页查询
db.Offset(offset).Limit(pageSize).Find(&results)
```

2. **连接池配置**
```go
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```

3. **避免N+1查询**
```go
// 不推荐: N+1查询
var bookings []Booking
db.Find(&bookings)
for i := range bookings {
    db.First(&bookings[i].User) // 每次循环都查询
}

// 推荐: 预加载
db.Preload("User").Find(&bookings)
```

### 小程序优化

1. **图片优化**
   - 使用webp格式
   - 压缩图片大小
   - 使用CDN加速

2. **列表优化**
   - 使用虚拟列表
   - 分页加载
   - 图片懒加载

3. **减少setData调用**
```javascript
// 不推荐: 频繁setData
for (let i = 0; i < 100; i++) {
  this.setData({ count: i })
}

// 推荐: 合并setData
this.setData({ count: 99 })
```

## 安全注意事项

### 1. 敏感信息保护

**不要硬编码**:
```go
// 错误示例
secret := "my-secret-key"

// 正确示例
secret := config.AppConfig.JWT.Secret
```

**.gitignore配置**:
```
config.yaml
.env
*.pem
*.key
```

### 2. SQL注入防护

GORM已自动处理,但仍需注意:
```go
// 安全: 使用参数化查询
db.Where("name = ?", userInput).Find(&users)

// 危险: 字符串拼接(不要用!)
db.Where("name = '" + userInput + "'").Find(&users)
```

### 3. XSS防护

小程序已自动转义,显示HTML时使用:
```html
<!-- 安全 -->
<text>{{content}}</text>

<!-- 需要显示富文本时 -->
<rich-text nodes="{{content}}"></rich-text>
```

### 4. 权限控制

```go
// 管理员接口必须使用AdminMiddleware
admin := api.Group("/admin")
admin.Use(middleware.AuthMiddleware())
admin.Use(middleware.AdminMiddleware())
{
    // 管理员路由
}
```

## 调试技巧

### Go调试

1. **日志输出**
```go
import "log"

log.Printf("Debug: user=%v, action=%s", user, action)
```

2. **使用Delve调试器**
```bash
# 安装
go install github.com/go-delve/delve/cmd/dlv@latest

# 启动调试
dlv debug cmd/main.go
```

3. **查看SQL语句**
```go
// 在config中启用日志
Logger: logger.New(..., logger.Config{
    LogLevel: logger.Info, // 显示所有SQL
})
```

### 小程序调试

1. **Console调试**
```javascript
console.log('数据:', data)
console.error('错误:', error)
```

2. **Network面板**
- 查看请求详情
- 检查请求头和响应
- 分析性能

3. **Wxml面板**
- 查看DOM结构
- 实时编辑样式
- 检查数据绑定

## 部署前检查清单

### 后端
- [ ] 修改JWT密钥为强随机字符串
- [ ] 配置正确的数据库连接(生产环境)
- [ ] 设置server.mode为"release"
- [ ] 配置微信AppID和AppSecret
- [ ] 配置HTTPS证书
- [ ] 设置防火墙规则
- [ ] 配置日志轮转
- [ ] 备份数据库脚本

### 小程序
- [ ] 修改baseURL为生产域名
- [ ] 配置服务器域名白名单
- [ ] 移除console.log(生产环境)
- [ ] 测试真机兼容性
- [ ] 提交代码审核
- [ ] 准备隐私协议
- [ ] 配置客服消息

### 数据库
- [ ] 修改默认管理员密码
- [ ] 创建数据库备份
- [ ] 设置定时备份任务
- [ ] 监控数据库性能

## 版本管理

### Git提交规范

```bash
# 功能开发
git commit -m "feat: 添加用户登录功能"

# Bug修复
git commit -m "fix: 修复Token验证失败问题"

# 文档更新
git commit -m "docs: 更新API文档"

# 重构
git commit -m "refactor: 重构用户服务层"

# 性能优化
git commit -m "perf: 优化数据库查询性能"
```

### 分支策略

```
main          # 主分支(生产环境)
├── develop   # 开发分支
│   ├── feature/login      # 功能分支
│   ├── feature/booking    # 功能分支
│   └── bugfix/token-fix   # Bug修复分支
```

## 资源链接

- [Go官方文档](https://go.dev/doc/)
- [Gin框架文档](https://gin-gonic.com/zh-cn/docs/)
- [GORM文档](https://gorm.io/zh_CN/docs/)
- [微信小程序文档](https://developers.weixin.qq.com/miniprogram/dev/framework/)
- [MySQL文档](https://dev.mysql.com/doc/)

---

**持续更新中...**
