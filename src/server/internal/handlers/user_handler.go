package handlers

import (
	"tennis-booking-system/internal/config"
	"tennis-booking-system/internal/middleware"
	"tennis-booking-system/internal/services"
	"tennis-booking-system/internal/utils"

	"github.com/gin-gonic/gin"
)

var userService *services.UserService

// InitHandlers 初始化所有handler
func InitHandlers() {
	db := config.GetDB()
	userService = services.NewUserService(db)
}

// HandleWxLogin 微信登录
func HandleWxLogin(c *gin.Context) {
	var req services.WXLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误: "+err.Error())
		return
	}

	resp, err := userService.WXLogin(&req)
	if err != nil {
		utils.InternalError(c, "登录失败: "+err.Error())
		return
	}

	// 生成JWT Token
	token, err := middleware.GenerateToken(resp.UserInfo.ID, resp.UserInfo.OpenID, false, "")
	if err != nil {
		utils.InternalError(c, "生成Token失败")
		return
	}

	resp.Token = token
	utils.Success(c, resp)
}

// HandleGetUserProfile 获取用户信息
func HandleGetUserProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)

	user, err := userService.GetProfile(userID)
	if err != nil {
		utils.NotFound(c, "用户不存在")
		return
	}

	utils.Success(c, user)
}

// HandleUpdateUserProfile 更新用户信息
func HandleUpdateUserProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		Phone    string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	if err := userService.UpdateProfile(userID, req.Nickname, req.Avatar, req.Phone); err != nil {
		utils.InternalError(c, "更新失败: "+err.Error())
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}
