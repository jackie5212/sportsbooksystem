package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tennis-booking-system/internal/config"
	"tennis-booking-system/internal/models"

	"gorm.io/gorm"
)

// WXLoginRequest 微信登录请求
type WXLoginRequest struct {
	Code     string `json:"code" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
}

// WXLoginResponse 微信登录响应
type WXLoginResponse struct {
	Token    string       `json:"token"`
	UserInfo *models.User `json:"user_info"`
}

// WeChatSessionResponse 微信session响应
type WeChatSessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// UserService 用户服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// WXLogin 微信小程序登录
func (s *UserService) WXLogin(req *WXLoginRequest) (*WXLoginResponse, error) {
	// 调用微信API获取openid和session_key
	sessionResp, err := s.getWeChatSession(req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to get wechat session: %v", err)
	}

	if sessionResp.ErrCode != 0 {
		return nil, fmt.Errorf("wechat api error: %s", sessionResp.ErrMsg)
	}

	// 查找或创建用户
	var user models.User
	err = s.db.Where("openid = ?", sessionResp.OpenID).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		// 创建新用户
		user = models.User{
			OpenID:     sessionResp.OpenID,
			SessionKey: sessionResp.SessionKey,
			Nickname:   req.Nickname,
			Avatar:     req.Avatar,
			Gender:     req.Gender,
			Country:    req.Country,
			Province:   req.Province,
			City:       req.City,
			Status:     1,
		}

		if err := s.db.Create(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to create user: %v", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to query user: %v", err)
	} else {
		// 更新用户信息
		user.SessionKey = sessionResp.SessionKey
		if req.Nickname != "" {
			user.Nickname = req.Nickname
		}
		if req.Avatar != "" {
			user.Avatar = req.Avatar
		}
		s.db.Save(&user)
	}

	// 生成JWT Token
	token, err := s.generateUserToken(user.ID, user.OpenID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &WXLoginResponse{
		Token:    token,
		UserInfo: &user,
	}, nil
}

// GetProfile 获取用户信息
func (s *UserService) GetProfile(userID uint64) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// 隐藏敏感信息
	user.SessionKey = ""
	return &user, nil
}

// UpdateProfile 更新用户信息
func (s *UserService) UpdateProfile(userID uint64, nickname, avatar, phone string) error {
	updates := make(map[string]interface{})
	if nickname != "" {
		updates["nickname"] = nickname
	}
	if avatar != "" {
		updates["avatar"] = avatar
	}
	if phone != "" {
		updates["phone"] = phone
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	if err := s.db.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update profile: %v", err)
	}

	return nil
}

// getWeChatSession 调用微信API获取session
func (s *UserService) getWeChatSession(code string) (*WeChatSessionResponse, error) {
	// 优先从数据库 system_configs 读取小程序配置
	appID := GetSettingWithDefault("wechat_mini_appid", config.AppConfig.WeChat.AppID)
	appSecret := GetSettingWithDefault("wechat_mini_appsecret", config.AppConfig.WeChat.AppSecret)

	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appID, appSecret, code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sessionResp WeChatSessionResponse
	if err := json.Unmarshal(body, &sessionResp); err != nil {
		return nil, err
	}

	return &sessionResp, nil
}

// generateUserToken 生成用户Token
func (s *UserService) generateUserToken(userID uint64, openID string) (string, error) {
	// 这里需要导入middleware包,暂时返回空
	// 实际应该在handler层调用middleware.GenerateToken
	return "", nil
}
