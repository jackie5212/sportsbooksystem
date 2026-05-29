package handlers

import (
	"strconv"
	"tennis-booking-system/internal/config"
	"tennis-booking-system/internal/models"
	"tennis-booking-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// HandleGetCourts 获取场地列表
func HandleGetCourts(c *gin.Context) {
	db := config.GetDB()

	var courts []models.Court
	query := db.Where("status = ?", 1).Order("sort_order ASC, id ASC")

	// 分页
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var total int64
	db.Model(&models.Court{}).Where("status = ?", 1).Count(&total)

	query.Offset(offset).Limit(pageSize).Find(&courts)

	utils.Success(c, gin.H{
		"list":      courts,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// HandleGetCourtDetail 获取场地详情
func HandleGetCourtDetail(c *gin.Context) {
	db := config.GetDB()

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequest(c, "无效的场地ID")
		return
	}

	var court models.Court
	if err := db.First(&court, id).Error; err != nil {
		utils.NotFound(c, "场地不存在")
		return
	}

	utils.Success(c, court)
}

// HandleGetTimeSlots 获取时间段列表
func HandleGetTimeSlots(c *gin.Context) {
	db := config.GetDB()

	courtIDStr := c.Param("court_id")
	courtID, err := strconv.ParseUint(courtIDStr, 10, 64)
	if err != nil {
		utils.BadRequest(c, "无效的场地ID")
		return
	}

	date := c.Query("date")
	if date == "" {
		utils.BadRequest(c, "请提供日期参数")
		return
	}

	var slots []models.TimeSlot
	db.Where("court_id = ? AND date = ?", courtID, date).
		Order("start_time ASC").
		Find(&slots)

	utils.Success(c, slots)
}
