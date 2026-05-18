# API测试示例

本文档提供常用API的测试示例,可以使用curl、Postman或Apifox等工具进行测试。

## 基础信息

- **Base URL**: `http://localhost:8080/api`
- **Content-Type**: `application/json`

## 1. 健康检查

```bash
curl http://localhost:8080/health
```

**响应:**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "status": "ok"
  }
}
```

## 2. 获取场地列表

```bash
curl http://localhost:8080/api/courts?page=1&page_size=10
```

**响应:**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "name": "标准网球场A",
        "location": "东区1号场地",
        "description": "标准硬地网球场,灯光照明良好",
        "price_per_hour": 80.00,
        "status": 1,
        "images": "[\"https://example.com/court1.jpg\"]",
        "sort_order": 1
      }
    ],
    "total": 3,
    "page": 1,
    "page_size": 10
  }
}
```

## 3. 获取场地详情

```bash
curl http://localhost:8080/api/courts/1
```

## 4. 获取时间段列表

```bash
curl "http://localhost:8080/api/courts/1/slots?date=2026-05-15"
```

**响应:**
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "court_id": 1,
      "date": "2026-05-15",
      "start_time": "08:00:00",
      "end_time": "09:00:00",
      "status": 1
    }
  ]
}
```

## 5. 微信登录

**注意**: 需要真实的微信code,开发时可以使用测试工具生成

```bash
curl -X POST http://localhost:8080/api/auth/wx-login \
  -H "Content-Type: application/json" \
  -d '{
    "code": "test_code_here",
    "nickname": "测试用户",
    "avatar": "https://example.com/avatar.jpg",
    "gender": 1,
    "country": "中国",
    "province": "北京",
    "city": "北京"
  }'
```

**响应:**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user_info": {
      "id": 1,
      "openid": "xxx",
      "nickname": "测试用户",
      "avatar": "https://example.com/avatar.jpg"
    }
  }
}
```

## 6. 获取用户信息 (需要认证)

```bash
curl http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## 7. 更新用户信息 (需要认证)

```bash
curl -X PUT http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "新昵称",
    "phone": "13800138000"
  }'
```

## Postman集合示例

可以导入以下JSON到Postman:

```json
{
  "info": {
    "name": "网球预定系统API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "健康检查",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "http://localhost:8080/health",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["health"]
        }
      }
    },
    {
      "name": "获取场地列表",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "http://localhost:8080/api/courts?page=1&page_size=10",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "courts"],
          "query": [
            {"key": "page", "value": "1"},
            {"key": "page_size", "value": "10"}
          ]
        }
      }
    }
  ]
}
```

## 错误响应格式

所有错误响应遵循统一格式:

```json
{
  "code": 400,
  "message": "错误描述信息"
}
```

常见错误码:
- `200`: 成功
- `400`: 请求参数错误
- `401`: 未授权(Token无效或过期)
- `403`: 禁止访问(权限不足)
- `404`: 资源不存在
- `500`: 服务器内部错误

## 测试建议

1. **先测试公开接口**: 健康检查、场地列表
2. **再测试认证接口**: 登录后获取Token
3. **使用环境变量**: 在Postman中设置base_url和token变量
4. **自动化测试**: 可以编写脚本批量测试

## 性能测试

使用Apache Bench进行简单压力测试:

```bash
# 100个请求,10个并发
ab -n 100 -c 10 http://localhost:8080/api/courts

# 查看QPS和响应时间
```

使用wrk进行更专业的测试:

```bash
wrk -t12 -c400 -d30s http://localhost:8080/api/courts
```

---

**提示**: 实际使用时,请将`YOUR_TOKEN_HERE`替换为真实登录获得的Token。
