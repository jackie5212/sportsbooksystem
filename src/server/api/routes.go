package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"tennis-booking-system/internal/config"
	"tennis-booking-system/internal/handlers"
	"tennis-booking-system/internal/middleware"
	"tennis-booking-system/internal/models"
	"tennis-booking-system/internal/utils"
	"time"

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
		// 公共路由（无需认证）
		public := api.Group("/")
		{
			// 微信登录
			public.POST("/auth/wx-login", handlers.HandleWxLogin)
		
			// 管理员登录
			public.POST("/admin/login", handleAdminLogin)
		
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
			// 文件上传
			admin.POST("/upload", handleUploadImage)

			// 场地管理
			admin.GET("/courts", handleGetAdminCourts)
			admin.POST("/courts", handleCreateCourt)
			admin.PUT("/courts/:id", handleUpdateCourt)
			admin.DELETE("/courts/:id", handleDeleteCourt)

			// 时间段管理
			admin.POST("/slots", handleGenerateTimeSlots)
			admin.PUT("/slots/:id", handleUpdateTimeSlot)

			// 订单管理
			admin.GET("/bookings", handleGetAllBookings)
			admin.GET("/dashboard", handleGetDashboard)

			// 系统设置
			admin.GET("/settings", handlers.HandleGetSettings)
			admin.PUT("/settings", handlers.HandleUpdateSettings)
			admin.GET("/settings/wechat", handlers.HandleGetWechatConfig)
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
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 简单的账号密码验证（实际应该查数据库并加密验证）
	if loginData.Username == "admin" && loginData.Password == "admin123" {
		// 生成 JWT Token
		token, err := middleware.GenerateToken(1, "", true, "admin")
		if err != nil {
			utils.InternalError(c, "生成Token失败")
			return
		}

		utils.Success(c, gin.H{
			"token": token,
			"admin_info": gin.H{
				"id":       1,
				"username": "admin",
				"nickname": "系统管理员",
				"role":     "super_admin",
			},
		})
	} else {
		utils.Unauthorized(c, "用户名或密码错误")
	}
}

func handleCreateCourt(c *gin.Context) {
	// 使用 map 接收原始数据
	var rawData map[string]interface{}
	if err := c.ShouldBindJSON(&rawData); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 处理 images 字段
	var imagesJSON string
	if images, ok := rawData["images"]; ok {
		switch v := images.(type) {
		case []interface{}:
			// 将数组转换为JSON字符串
			jsonBytes, err := json.Marshal(v)
			if err == nil {
				imagesJSON = string(jsonBytes)
			} else {
				imagesJSON = "[]"
			}
		case string:
			imagesJSON = v
		default:
			imagesJSON = "[]"
		}
	} else {
		imagesJSON = "[]"
	}
	rawData["images"] = imagesJSON

	// 处理 facilities 字段（如果存在）
	if facilities, ok := rawData["facilities"]; ok {
		switch v := facilities.(type) {
		case []interface{}, map[string]interface{}:
			jsonBytes, err := json.Marshal(v)
			if err == nil {
				rawData["facilities"] = string(jsonBytes)
			}
		case string:
			// 已经是字符串
		default:
			rawData["facilities"] = "[]"
		}
	}

	// 将 map 转换为 Court 结构体
	var court models.Court
	court.Name = getStringValue(rawData, "name")
	court.Location = getStringValue(rawData, "location")
	court.Description = getStringValue(rawData, "description")
	court.PricePerHour = getFloatValue(rawData, "price_per_hour")
	court.Status = getIntValue(rawData, "status", 1)
	court.Images = imagesJSON
	court.Facilities = getStringValue(rawData, "facilities")
	court.SortOrder = getIntValue(rawData, "sort_order", 0)

	if err := config.DB.Create(&court).Error; err != nil {
		utils.InternalError(c, "创建场地失败: "+err.Error())
		return
	}

	utils.Success(c, court)
}

func handleUpdateCourt(c *gin.Context) {
	id := c.Param("id")
	var court models.Court

	// 查找场地
	if err := config.DB.First(&court, id).Error; err != nil {
		utils.NotFound(c, "场地不存在")
		return
	}

	// 更新字段
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 如果更新了 images，确保它是 JSON 字符串
	if images, ok := updateData["images"]; ok {
		switch v := images.(type) {
		case []interface{}:
			// 将数组转换为JSON字符串
			jsonBytes, err := json.Marshal(v)
			if err == nil {
				updateData["images"] = string(jsonBytes)
			} else {
				updateData["images"] = "[]"
			}
		case string:
			// 已经是字符串，检查是否是有效的JSON
			var js interface{}
			if err := json.Unmarshal([]byte(v), &js); err != nil {
				// 如果不是有效的JSON，设为空数组
				updateData["images"] = "[]"
			}
		default:
			updateData["images"] = "[]"
		}
	}

	// 如果更新了 facilities，确保它是 JSON 字符串
	if facilities, ok := updateData["facilities"]; ok {
		switch v := facilities.(type) {
		case []interface{}, map[string]interface{}:
			jsonBytes, err := json.Marshal(v)
			if err == nil {
				updateData["facilities"] = string(jsonBytes)
			}
		case string:
			// 已经是字符串
		default:
			updateData["facilities"] = "[]"
		}
	}

	if err := config.DB.Model(&court).Updates(updateData).Error; err != nil {
		utils.InternalError(c, "更新场地失败: "+err.Error())
		return
	}

	// 重新查询获取最新数据
	config.DB.First(&court, id)
	utils.Success(c, court)
}

// 辅助函数
func getStringValue(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getFloatValue(data map[string]interface{}, key string) float64 {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case float64:
			return v
		case int:
			return float64(v)
		}
	}
	return 0
}

func getIntValue(data map[string]interface{}, key string, defaultVal int) int {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case float64:
			return int(v)
		case int:
			return v
		}
	}
	return defaultVal
}

func handleDeleteCourt(c *gin.Context) {
	id := c.Param("id")
	var court models.Court

	// 查找场地
	if err := config.DB.First(&court, id).Error; err != nil {
		utils.NotFound(c, "场地不存在")
		return
	}

	// 删除场地
	if err := config.DB.Delete(&court).Error; err != nil {
		utils.InternalError(c, "删除场地失败: "+err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "删除成功"})
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

// handleUploadImage 处理图片上传
func handleUploadImage(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("image")
	if err != nil {
		utils.BadRequest(c, "获取文件失败")
		return
	}

	// 验证文件类型
	if !isImageFile(file) {
		utils.BadRequest(c, "只支持图片文件(jpg, jpeg, png, gif, webp)")
		return
	}

	// 验证文件大小（最大5MB）
	if file.Size > 5*1024*1024 {
		utils.BadRequest(c, "文件大小不能超过5MB")
		return
	}

	// 创建上传目录
	uploadDir := "./uploads/courts"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		utils.InternalError(c, "创建目录失败")
		return
	}

	// 生成文件名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), generateRandomString(8), ext)
	filepath := fmt.Sprintf("%s/%s", uploadDir, filename)

	// 保存文件
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		utils.InternalError(c, "保存文件失败")
		return
	}

	// 生成访问URL（使用相对路径，让浏览器自动拼接当前域名和协议）
	url := fmt.Sprintf("/uploads/courts/%s", filename)

	utils.Success(c, gin.H{
		"url":      url,
		"filename": filename,
		"size":     file.Size,
	})
}

// isImageFile 验证是否为图片文件
func isImageFile(file *multipart.FileHeader) bool {
	allowedTypes := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	for _, t := range allowedTypes {
		if ext == t {
			return true
		}
	}
	return false
}

// generateRandomString 生成随机字符串
func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}

// handleGetAdminCourts 管理员获取场地列表（支持分页和搜索）
func handleGetAdminCourts(c *gin.Context) {
	// 获取分页参数
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	search := c.Query("search")

	// 转换为整数
	pageInt := 1
	pageSizeInt := 10
	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageInt = p
	}
	if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
		pageSizeInt = ps
	}

	// 构建查询
	query := config.DB.Model(&models.Court{})

	// 如果有搜索关键词，添加搜索条件
	if search != "" {
		query = query.Where("name LIKE ? OR location LIKE ? OR description LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// 获取总数
	var total int64
	query.Count(&total)

	// 分页查询
	var courts []models.Court
	offset := (pageInt - 1) * pageSizeInt
	if err := query.Order("id ASC").Offset(offset).Limit(pageSizeInt).Find(&courts).Error; err != nil {
		utils.InternalError(c, "获取场地列表失败: "+err.Error())
		return
	}

	// 转换数据格式，解析 images JSON 字符串
	result := make([]gin.H, len(courts))
	for i, court := range courts {
		var images []string
		if court.Images != "" {
			if err := json.Unmarshal([]byte(court.Images), &images); err != nil {
				images = []string{}
			}
		}

		var facilities interface{}
		if court.Facilities != "" {
			json.Unmarshal([]byte(court.Facilities), &facilities)
		}

		result[i] = gin.H{
			"id":             court.ID,
			"name":           court.Name,
			"location":       court.Location,
			"description":    court.Description,
			"price_per_hour": court.PricePerHour,
			"status":         court.Status,
			"images":         images,
			"facilities":     facilities,
			"sort_order":     court.SortOrder,
			"created_at":     court.CreatedAt,
			"updated_at":     court.UpdatedAt,
		}
	}

	utils.Success(c, gin.H{
		"list":      result,
		"total":     total,
		"page":      pageInt,
		"page_size": pageSizeInt,
	})
}
