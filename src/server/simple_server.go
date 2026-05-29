package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置Gin为发布模式
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    gin.H{"status": "ok"},
		})
	})

	// 模拟场地数据 - 丰富的场地列表
	courts := []gin.H{
		{
			"id":             1,
			"name":           "标准硬地场A",
			"location":       "东区1号",
			"description":    "国际标准硬地场地,灯光照明优良,适合各类比赛和训练。配备专业防滑地面,提供舒适的运动体验。",
			"price_per_hour": 80.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=1"},
			"facilities":     []string{"灯光照明", "休息区", "淋浴间"},
			"rating":         4.8,
		},
		{
			"id":             2,
			"name":           "标准硬地场B",
			"location":       "东区2号",
			"description":    "专业级硬地网球场,地面平整度高,球速适中。适合初学者和进阶球员日常练习使用。",
			"price_per_hour": 80.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=2"},
			"facilities":     []string{"灯光照明", "饮水处"},
			"rating":         4.6,
		},
		{
			"id":             3,
			"name":           "红土场C",
			"location":       "西区1号",
			"description":    "法式红土场地,球速较慢,回合多。适合喜欢拉锯战的球员,体验法网同款场地感受。",
			"price_per_hour": 120.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=3"},
			"facilities":     []string{"红土地面", "遮阳棚", "休息椅"},
			"rating":         4.9,
		},
		{
			"id":             4,
			"name":           "草地场D",
			"location":       "西区2号",
			"description":    "天然草皮场地,球速快,弹跳低。温网同款场地类型,适合进攻型打法球员。",
			"price_per_hour": 150.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=4"},
			"facilities":     []string{"天然草皮", "专业维护", "VIP通道"},
			"rating":         5.0,
		},
		{
			"id":             5,
			"name":           "室内恒温场E",
			"location":       "室内馆1号",
			"description":    "全天候室内场地,恒温空调系统,不受天气影响。全年无休,随时可约。",
			"price_per_hour": 100.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=5"},
			"facilities":     []string{"空调系统", "室内照明", "更衣室", "储物柜"},
			"rating":         4.7,
		},
		{
			"id":             6,
			"name":           "VIP豪华场F",
			"location":       "贵宾区",
			"description":    "顶级VIP场地,独立空间,私密性强。配备专属休息区、淋浴间和私人教练服务。",
			"price_per_hour": 200.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=6"},
			"facilities":     []string{"独立空间", "私人教练", "淋浴间", "茶水区", "按摩椅"},
			"rating":         5.0,
		},
		{
			"id":             7,
			"name":           "青少年训练场G",
			"location":       "训练区1号",
			"description":    "专为青少年设计的训练场地,尺寸略小,网高可调。配备专业教练指导,适合儿童和青少年学习网球。",
			"price_per_hour": 60.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=7"},
			"facilities":     []string{"可调网高", "教学设备", "安全防护"},
			"rating":         4.5,
		},
		{
			"id":             8,
			"name":           "夜光场H",
			"location":       "夜景區",
			"description":    "特色夜光场地,夜间照明效果极佳。荧光标记线,营造独特的夜间运动氛围。",
			"price_per_hour": 90.00,
			"status":         1,
			"images":         []string{"https://picsum.photos/400/300?random=8"},
			"facilities":     []string{"LED照明", "荧光标线", "音响系统"},
			"rating":         4.8,
		},
	}

	// 获取场地列表
	r.GET("/api/courts", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("page_size", "10")

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"list":      courts,
				"total":     len(courts),
				"page":      page,
				"page_size": pageSize,
			},
		})
	})

	// 获取场地详情
	r.GET("/api/courts/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, court := range courts {
			if fmt.Sprintf("%.0f", court["id"]) == id {
				c.JSON(200, gin.H{
					"code":    200,
					"message": "success",
					"data":    court,
				})
				return
			}
		}

		c.JSON(404, gin.H{
			"code":    404,
			"message": "场地不存在",
		})
	})

	// 获取时间段
	r.GET("/api/courts/:id/slots", func(c *gin.Context) {
		date := c.DefaultQuery("date", "2026-05-15")
		courtId := c.Param("id")

		// 根据不同场地生成不同的时间段数据
		slots := []gin.H{}

		// 基础时间段 (8:00 - 21:00)
		timeRanges := []struct {
			start string
			end   string
		}{
			{"08:00:00", "09:00:00"},
			{"09:00:00", "10:00:00"},
			{"10:00:00", "11:00:00"},
			{"11:00:00", "12:00:00"},
			{"14:00:00", "15:00:00"},
			{"15:00:00", "16:00:00"},
			{"16:00:00", "17:00:00"},
			{"17:00:00", "18:00:00"},
			{"19:00:00", "20:00:00"},
			{"20:00:00", "21:00:00"},
		}

		// 根据场地ID生成不同的可用状态
		for i, tr := range timeRanges {
			status := 1 // 默认可用

			// 模拟部分时段已被预定
			if courtId == "1" && (i == 2 || i == 5) {
				status = 2 // 已预定
			} else if courtId == "3" && (i == 1 || i == 4 || i == 7) {
				status = 2
			} else if courtId == "6" && i < 3 {
				status = 2 // VIP场地上午较忙
			}

			slots = append(slots, gin.H{
				"id":         i + 1,
				"court_id":   courtId,
				"date":       date,
				"start_time": tr.start,
				"end_time":   tr.end,
				"status":     status,
			})
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    slots,
		})
	})

	// ========== 管理后台API ==========

	// 管理员登录
	r.POST("/api/admin/login", func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "请求参数错误",
			})
			return
		}

		// 简单的账号密码验证(实际应该查数据库并加密验证)
		if loginData.Username == "admin" && loginData.Password == "admin123" {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "登录成功",
				"data": gin.H{
					"token": "admin_token_" + fmt.Sprintf("%d", time.Now().Unix()),
					"admin_info": gin.H{
						"id":       1,
						"username": "admin",
						"nickname": "系统管理员",
						"role":     "super_admin",
					},
				},
			})
		} else {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "用户名或密码错误",
			})
		}
	})

	// 获取统计数据
	r.GET("/api/admin/statistics", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"totalCourts":   len(courts),
				"todayBookings": 12,
				"totalUsers":    156,
				"todayRevenue":  960,
				"weekBookings":  85,
				"monthBookings": 342,
			},
		})
	})

	// 获取订单列表
	r.GET("/api/admin/bookings", func(c *gin.Context) {
		// 模拟订单数据
		bookings := []gin.H{
			{
				"id":           1,
				"order_no":     "TB202605150001",
				"user_name":    "郑钦文",
				"court_name":   "标准硬地场A",
				"booking_date": "2026-05-15",
				"time_range":   "08:00-09:00",
				"total_amount": 80,
				"status":       1,
				"created_at":   "2026-05-15 07:30:00",
			},
			{
				"id":           2,
				"order_no":     "TB202605150002",
				"user_name":    "张德培",
				"court_name":   "VIP豪华场F",
				"booking_date": "2026-05-15",
				"time_range":   "14:00-15:00",
				"total_amount": 200,
				"status":       2,
				"created_at":   "2026-05-15 08:15:00",
			},
			{
				"id":           3,
				"order_no":     "TB202605140003",
				"user_name":    "桑普拉斯",
				"court_name":   "红土场C",
				"booking_date": "2026-05-14",
				"time_range":   "10:00-11:00",
				"total_amount": 120,
				"status":       3,
				"created_at":   "2026-05-14 09:20:00",
			},
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"list":  bookings,
				"total": len(bookings),
			},
		})
	})

	// 获取用户列表
	r.GET("/api/admin/users", func(c *gin.Context) {
		// 模拟用户数据
		users := []gin.H{
			{
				"id":            1,
				"nickname":      "郑钦文",
				"phone":         "138****1234",
				"openid":        "oXXXX1",
				"booking_count": 5,
				"created_at":    "2026-05-01 10:00:00",
			},
			{
				"id":            2,
				"nickname":      "张德培",
				"phone":         "139****5678",
				"openid":        "oXXXX2",
				"booking_count": 3,
				"created_at":    "2026-05-05 14:30:00",
			},
			{
				"id":            3,
				"nickname":      "桑普拉斯",
				"phone":         "137****9012",
				"openid":        "oXXXX3",
				"booking_count": 8,
				"created_at":    "2026-05-10 09:15:00",
			},
			{
				"id":            4,
				"nickname":      "阿加西",
				"phone":         "136****3456",
				"openid":        "oXXXX4",
				"booking_count": 12,
				"created_at":    "2026-05-12 16:45:00",
			},
			{
				"id":            5,
				"nickname":      "纳达尔",
				"phone":         "135****7890",
				"openid":        "oXXXX5",
				"booking_count": 15,
				"created_at":    "2026-05-13 11:20:00",
			},
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"list":  users,
				"total": len(users),
			},
		})
	})

	// 获取时间段列表
	r.GET("/api/admin/timeslots", func(c *gin.Context) {
		courtId := c.DefaultQuery("court_id", "1")
		date := c.DefaultQuery("date", "2026-05-15")

		// 生成时间段数据
		slots := []gin.H{}
		timeRanges := []struct {
			start string
			end   string
		}{
			{"08:00:00", "09:00:00"},
			{"09:00:00", "10:00:00"},
			{"10:00:00", "11:00:00"},
			{"11:00:00", "12:00:00"},
			{"14:00:00", "15:00:00"},
			{"15:00:00", "16:00:00"},
			{"16:00:00", "17:00:00"},
			{"17:00:00", "18:00:00"},
			{"19:00:00", "20:00:00"},
			{"20:00:00", "21:00:00"},
		}

		for i, tr := range timeRanges {
			status := 1
			if i%3 == 0 {
				status = 2 // 模拟部分已预定
			}

			slots = append(slots, gin.H{
				"id":         i + 1,
				"court_id":   courtId,
				"date":       date,
				"start_time": tr.start,
				"end_time":   tr.end,
				"status":     status,
			})
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"list":  slots,
				"total": len(slots),
			},
		})
	})

	fmt.Println("========================================")
	fmt.Println("  网球预定系统 - 演示版本")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("服务已启动!")
	fmt.Println("访问地址: http://localhost:8888")
	fmt.Println()
	fmt.Println("测试接口:")
	fmt.Println("  - http://localhost:8888/health")
	fmt.Println("  - http://localhost:8888/api/courts")
	fmt.Println("  - http://localhost:8888/api/courts/1")
	fmt.Println("  - http://localhost:8888/api/courts/1/slots?date=2026-05-15")
	fmt.Println()
	fmt.Println("按 Ctrl+C 停止服务")
	fmt.Println("========================================")

	log.Fatal(r.Run(":8888"))
}
