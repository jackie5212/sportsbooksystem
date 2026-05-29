package handlers

import (
	"tennis-booking-system/internal/services"
	"tennis-booking-system/internal/utils"

	"github.com/gin-gonic/gin"
)

// HandleGetSettings 获取所有系统配置
func HandleGetSettings(c *gin.Context) {
	settings, err := services.GetAllSettings()
	if err != nil {
		utils.InternalError(c, "获取系统配置失败")
		return
	}
	utils.Success(c, settings)
}

// HandleUpdateSettings 批量更新系统配置
func HandleUpdateSettings(c *gin.Context) {
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	if err := services.UpdateSettings(req); err != nil {
		utils.InternalError(c, "更新系统配置失败: "+err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "更新成功"})
}

// HandleGetWechatConfig 获取微信支付配置（仅返回配置项，不包含敏感密钥）
func HandleGetWechatConfig(c *gin.Context) {
	settings, err := services.GetAllSettings()
	if err != nil {
		utils.InternalError(c, "获取配置失败")
		return
	}

	// 只返回配置项，密钥字段做脱敏处理
	data := gin.H{
		"wechat_web_appid":       settings["wechat_web_appid"],
		"wechat_web_appsecret":   maskSecret(settings["wechat_web_appsecret"]),
		"wechat_mini_appid":      settings["wechat_mini_appid"],
		"wechat_mini_appsecret":  maskSecret(settings["wechat_mini_appsecret"]),
		"wechat_mch_id":          settings["wechat_mch_id"],
		"wechat_api_v3_key":      maskSecret(settings["wechat_api_v3_key"]),
		"wechat_notify_url":      settings["wechat_notify_url"],
	}

	utils.Success(c, data)
}

// maskSecret 对密钥进行脱敏处理
func maskSecret(s string) string {
	if len(s) <= 8 {
		if s == "" {
			return ""
		}
		return "****"
	}
	return s[:4] + "****" + s[len(s)-4:]
}
