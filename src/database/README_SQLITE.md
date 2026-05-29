# SQLite数据库使用说明

## ✅ 数据库已创建成功!

### 📊 数据库信息

- **文件位置**: `D:\github\booksystemTongYi\tennis_booking.db`
- **数据库类型**: SQLite 3
- **无需安装**: 不需要MySQL,开箱即用
- **文件大小**: 约100KB

### 📋 已创建的表

| 表名 | 说明 | 记录数 |
|------|------|--------|
| users | 用户表 | 0 |
| courts | 场地表 | 3 |
| time_slots | 时间段表 | 210 |
| bookings | 订单表 | 0 |
| payments | 支付表 | 0 |
| admins | 管理员表 | 1 |
| system_configs | 系统配置表 | 5 |

**总计**: 8个表, 219条初始数据

### 🎯 初始数据

#### 场地数据 (3个)
1. 标准网球场A - ¥80/小时
2. 标准网球场B - ¥80/小时
3. VIP网球场 - ¥150/小时

#### 时间段数据 (210个)
- 未来7天
- 每天10个时段 (08:00-21:00)
- 3个场地 × 7天 × 10时段 = 210个

#### 管理员账号
- 用户名: **admin**
- 密码: **admin123**

#### 系统配置
- 网站名称: 网球预定系统
- 客服电话: 400-123-4567
- 营业时间: 08:00-22:00
- 可提前预定: 7天
- 取消提前时间: 2小时

---

## 🔧 如何查看数据库

### 方法1: DB Browser for SQLite (推荐)

1. **下载安装**: https://sqlitebrowser.org/dl/
2. **打开软件**: 选择 "Open Database"
3. **选择文件**: `tennis_booking.db`
4. **浏览数据**: 点击 "Browse Data" 标签页

### 方法2: VS Code插件

1. **安装插件**: "SQLite Viewer" 或 "SQLite"
2. **打开文件**: 右键点击 `tennis_booking.db`
3. **查看数据**: 选择 "Open Database"

### 方法3: Python脚本查询

```python
import sqlite3

conn = sqlite3.connect('tennis_booking.db')
cursor = conn.cursor()

# 查询所有场地
cursor.execute("SELECT * FROM courts")
courts = cursor.fetchall()
for court in courts:
    print(court)

conn.close()
```

### 方法4: 命令行工具

```bash
# Windows PowerShell
sqlite3 tennis_booking.db ".tables"
sqlite3 tennis_booking.db "SELECT * FROM courts;"
```

---

## 🚀 重新初始化数据库

如果需要重置数据库:

```bash
cd d:\github\booksystemTongYi
python database\init_sqlite.py
```

脚本会询问是否删除旧数据库,输入 `y` 确认即可。

---

## 📝 数据库结构

### users (用户表)
```sql
id, openid, session_key, nickname, avatar, phone, 
gender, country, province, city, status, 
created_at, updated_at
```

### courts (场地表)
```sql
id, name, location, description, price_per_hour, 
images, status, sort_order, created_at, updated_at
```

### time_slots (时间段表)
```sql
id, court_id, date, start_time, end_time, status, 
created_at, updated_at
```

### bookings (订单表)
```sql
id, order_no, user_id, court_id, time_slot_id, 
booking_date, start_time, end_time, duration_hours, 
total_amount, status, remark, created_at, updated_at
```

### payments (支付表)
```sql
id, booking_id, transaction_id, amount, payment_method, 
status, paid_at, created_at, updated_at
```

### admins (管理员表)
```sql
id, username, password, nickname, role, status, 
last_login_at, created_at, updated_at
```

### system_configs (系统配置表)
```sql
id, config_key, config_value, description, 
created_at, updated_at
```

---

## 💡 SQLite优势

### 优点
✅ **零配置** - 无需安装服务器  
✅ **单文件** - 整个数据库就是一个文件  
✅ **跨平台** - Windows/Mac/Linux通用  
✅ **轻量级** - 占用资源极少  
✅ **开发友好** - 适合开发和测试  

### 适用场景
- ✅ 开发环境
- ✅ 小型项目
- ✅ 原型演示
- ✅ 单元测试
- ⚠️ 不适合高并发生产环境

---

## 🔄 迁移到MySQL

当需要迁移到MySQL时:

1. **导出SQLite数据**:
   ```bash
   sqlite3 tennis_booking.db .dump > export.sql
   ```

2. **修改SQL语法** (SQLite和MySQL有些差异)

3. **导入MySQL**:
   ```bash
   mysql -u root -p tennis_booking < export.sql
   ```

或者使用专门的迁移工具。

---

## ⚙️ 后端配置

如果使用Go后端连接SQLite,需要:

1. **安装SQLite驱动**:
   ```bash
   go get github.com/mattn/go-sqlite3
   ```

2. **修改配置** (`config.yaml`):
   ```yaml
   database:
     driver: sqlite3
     dsn: tennis_booking.db
   ```

3. **修改GORM配置**:
   ```go
   import _ "github.com/mattn/go-sqlite3"
   
   db, err := gorm.Open(sqlite.Open("tennis_booking.db"), &gorm.Config{})
   ```

---

## 📞 常见问题

### Q: SQLite和MySQL有什么区别?
A: SQLite是文件型数据库,适合开发;MySQL是服务器型,适合生产。

### Q: 可以同时多人访问吗?
A: 可以,但并发性能不如MySQL。建议开发用SQLite,生产用MySQL。

### Q: 数据库文件在哪里?
A: `D:\github\booksystemTongYi\tennis_booking.db`

### Q: 如何备份数据库?
A: 直接复制 `tennis_booking.db` 文件即可。

### Q: 忘记密码怎么办?
A: 重新运行 `python database\init_sqlite.py` 重置数据库。

---

**数据库状态**: ✅ 已就绪  
**最后更新**: 2026-05-15
