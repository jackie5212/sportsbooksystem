package services

import (
	"encoding/json"
	"fmt"
	"tennis-booking-system/internal/config"
	"tennis-booking-system/internal/models"
)

// WechatConfig 微信支付配置结构
type WechatConfig struct {
	WebAppID     string `json:"wechat_web_appid"`
	WebAppSecret string `json:"wechat_web_appsecret"`
	MiniAppID    string `json:"wechat_mini_appid"`
	MiniAppSecret string `json:"wechat_mini_appsecret"`
	MchID        string `json:"wechat_mch_id"`
	APIV3Key     string `json:"wechat_api_v3_key"`
	NotifyURL    string `json:"wechat_notify_url"`
}

// GetAllSettings 获取所有系统配置（以 map 形式返回）
func GetAllSettings() (map[string]string, error) {
	var configs []models.SystemConfig
	if err := config.DB.Find(&configs).Error; err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, c := range configs {
		result[c.ConfigKey] = c.ConfigValue
	}
	return result, nil
}

// GetSetting 获取单个配置值
func GetSetting(key string) (string, error) {
	var cfg models.SystemConfig
	if err := config.DB.Where("config_key = ?", key).First(&cfg).Error; err != nil {
		return "", err
	}
	return cfg.ConfigValue, nil
}

// GetSettingWithDefault 获取单个配置值，不存在则返回默认值
func GetSettingWithDefault(key, defaultVal string) string {
	val, err := GetSetting(key)
	if err != nil {
		return defaultVal
	}
	return val
}

// UpdateSettings 批量更新系统配置
func UpdateSettings(settings map[string]string) error {
	for key, value := range settings {
		var cfg models.SystemConfig
		result := config.DB.Where("config_key = ?", key).First(&cfg)
		if result.Error != nil {
			// 不存在则创建
			cfg = models.SystemConfig{
				ConfigKey: key,
				ConfigValue: value,
			}
			if err := config.DB.Create(&cfg).Error; err != nil {
				return fmt.Errorf("create config %s failed: %v", key, err)
			}
		} else {
			// 存在则更新
			if err := config.DB.Model(&cfg).Update("config_value", value).Error; err != nil {
				return fmt.Errorf("update config %s failed: %v", key, err)
			}
		}
	}
	return nil
}

// GetWechatConfig 获取微信支付相关配置
func GetWechatConfig() (*WechatConfig, error) {
	settings, err := GetAllSettings()
	if err != nil {
		return nil, err
	}

	return &WechatConfig{
		WebAppID:      settings["wechat_web_appid"],
		WebAppSecret:  settings["wechat_web_appsecret"],
		MiniAppID:     settings["wechat_mini_appid"],
		MiniAppSecret: settings["wechat_mini_appsecret"],
		MchID:         settings["wechat_mch_id"],
		APIV3Key:      settings["wechat_api_v3_key"],
		NotifyURL:     settings["wechat_notify_url"],
	}, nil
}

// GetWechatConfigJSON 获取微信支付配置 JSON 字符串
func GetWechatConfigJSON() (string, error) {
	cfg, err := GetWechatConfig()
	if err != nil {
		return "", err
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
