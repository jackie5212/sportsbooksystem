package api

import (
	"fmt"
	"net/http"
	"tennis-booking-system/internal/handlers"
	"tennis-booking-system/internal/middleware"
	"tennis-booking-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine) {
	// 初始化handlers
	handlers.InitHandlers()

	// 全局中间件
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		utils.Success(c, gin.H{"status": "ok"})
	})

	// API路由组
	api := r.Group("/api")
	{
		// 公开路由(无需认证)
		public := api.Group("/")
		{
			// 微信登录
			public.POST("/auth/wx-login", handlers.HandleWxLogin)

			// 获取场地列表（使用模拟数据）
			public.GET("/courts", handleGetCourts)
			public.GET("/courts/:id", handleGetCourtDetail)
		}

		// 时间段路由（独立路径，避免冲突）
		slots := api.Group("/slots")
		{
			slots.GET("/court/:court_id", handlers.HandleGetTimeSlots)
		}

		// 需要认证的路由
		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware())
		{
			// 用户相关
			auth.GET("/user/profile", handlers.HandleGetUserProfile)
			auth.PUT("/user/profile", handlers.HandleUpdateUserProfile)

			// 预定相关
			auth.POST("/bookings", handleCreateBooking)
			auth.GET("/bookings", handleGetMyBookings)
			auth.GET("/bookings/:id", handleGetBookingDetail)
			auth.PUT("/bookings/:id/cancel", handleCancelBooking)

			// 支付相关
			auth.POST("/payments/create", handleCreatePayment)
			auth.GET("/payments/:booking_id", handleGetPaymentStatus)
		}

		// 管理员路由
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		admin.Use(middleware.AdminMiddleware())
		{
			// 管理员登录
			admin.POST("/login", handleAdminLogin)

			// 场地管理
			admin.POST("/courts", handleCreateCourt)
			admin.PUT("/courts/:id", handleUpdateCourt)
			admin.DELETE("/courts/:id", handleDeleteCourt)

			// 时间段管理
			admin.POST("/slots", handleGenerateTimeSlots)
			admin.PUT("/slots/:id", handleUpdateTimeSlot)

			// 订单管理
			admin.GET("/bookings", handleGetAllBookings)
			admin.GET("/dashboard", handleGetDashboard)
		}

		// 微信支付回调(无需认证)
		public.POST("/payments/notify", handlePaymentNotify)
	}
}

// 占位handler函数 - 后续实现
func handleWxLogin(c *gin.Context) {
	utils.Success(c, gin.H{"message": "wx login endpoint"})
}

func handleGetCourts(c *gin.Context) {
	// 模拟场地数据
	courts := []gin.H{
		{
			"id":             1,
			"name":           "标准硬地场A",
			"location":       "东区1号",
			"description":    "国际标准硬地场地,灯光照明优良,适合各类比赛和训练。配备专业防滑地面,提供舒适的运动体验。",
			"price_per_hour": 80.00,
			"status":         1,
			"images":         []string{"https://images.unsplash.com/photo-1622163642998-1ea36b1dde3b?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1554068865-24cecd4e34b8?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1595435934249-5df7ed86e1c0?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1530915516966-a30d20d9f5db?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1599474924187-334a4ae5bd3c?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1551698618-1dfe5d97d256?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1626245358043-4f9103f9e8b1?w=800&q=80"},
			"facilities":     []string{"可调网高", "教练指导", "安全防护"},
			"rating":         4.5,
		},
	}

	utils.Success(c, gin.H{
		"list":      courts,
		"total":     len(courts),
		"page":      1,
		"page_size": len(courts),
	})
}

func handleGetCourtDetail(c *gin.Context) {
	idStr := c.Param("id")
	
	// 模拟场地数据（与 handleGetCourts 中的数据保持一致）
	courts := []gin.H{
		{
			"id":             1,
			"name":           "标准硬地场A",
			"location":       "东区1号",
			"description":    "国际标准硬地场地,灯光照明优良,适合各类比赛和训练。配备专业防滑地面,提供舒适的运动体验。",
			"price_per_hour": 80.00,
			"status":         1,
			"images":         []string{"https://images.unsplash.com/photo-1622163642998-1ea36b1dde3b?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1554068865-24cecd4e34b8?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1595435934249-5df7ed86e1c0?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1530915516966-a30d20d9f5db?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1599474924187-334a4ae5bd3c?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1551698618-1dfe5d97d256?w=800&q=80"},
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
			"images":         []string{"https://images.unsplash.com/photo-1626245358043-4f9103f9e8b1?w=800&q=80"},
			"facilities":     []string{"可调网高", "教练指导", "安全防护"},
			"rating":         4.5,
		},
	}
	
	// 查找对应的场地
	for _, court := range courts {
		if fmt.Sprintf("%.0f", court["id"]) == idStr {
			utils.Success(c, court)
			return
		}
	}
	
	utils.NotFound(c, "场地不存在")
}

func handleGetTimeSlots(c *gin.Context) {
	utils.Success(c, gin.H{"message": "get time slots endpoint"})
}

func handleGetUserProfile(c *gin.Context) {
	utils.Success(c, gin.H{"message": "get user profile endpoint"})
}

func handleUpdateUserProfile(c *gin.Context) {
	utils.Success(c, gin.H{"message": "update user profile endpoint"})
}

func handleCreateBooking(c *gin.Context) {
	utils.Success(c, gin.H{"message": "create booking endpoint"})
}

func handleGetMyBookings(c *gin.Context) {
	utils.Success(c, gin.H{"message": "get my bookings endpoint"})
}

func handleGetBookingDetail(c *gin.Context) {
	utils.Success(c, gin.H{"message": "get booking detail endpoint"})
}

func handleCancelBooking(c *gin.Context) {
	utils.Success(c, gin.H{"message": "cancel booking endpoint"})
}

func handleCreatePayment(c *gin.Context) {
	utils.Success(c, gin.H{"message": "create payment endpoint"})
}

func handleGetPaymentStatus(c *gin.Context) {
	utils.Success(c, gin.H{"message": "get payment status endpoint"})
}

func handleAdminLogin(c *gin.Context) {
	utils.Success(c, gin.H{"message": "admin login endpoint"})
}

func handleCreateCourt(c *gin.Context) {
	utils.Success(c, gin.H{"message": "create court endpoint"})
}

func handleUpdateCourt(c *gin.Context) {
	utils.Success(c, gin.H{"message": "update court endpoint"})
}

func handleDeleteCourt(c *gin.Context) {
	utils.Success(c, gin.H{"message": "delete court endpoint"})
}

func handleGenerateTimeSlots(c *gin.Context) {
	utils.Success(c, gin.H{"message": "generate time slots endpoint"})
}

func handleUpdateTimeSlot(c *gin.Context) {
	utils.Success(c, gin.H{"message": "update time slot endpoint"})
}

func handleGetAllBookings(c *gin.Context) {
	utils.Success(c, gin.H{"message": "get all bookings endpoint"})
}

func handleGetDashboard(c *gin.Context) {
	utils.Success(c, gin.H{"message": "dashboard endpoint"})
}

func handlePaymentNotify(c *gin.Context) {
	c.String(http.StatusOK, "success")
}
