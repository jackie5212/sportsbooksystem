package models

import (
	"time"
)

// User 用户模型
type User struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	OpenID     string    `gorm:"column:openid;uniqueIndex;size:64;not null" json:"openid"`
	SessionKey string    `gorm:"column:session_key;size:128" json:"-"`
	Nickname   string    `gorm:"size:100;default:''" json:"nickname"`
	Avatar     string    `gorm:"size:500;default:''" json:"avatar"`
	Phone      string    `gorm:"size:20;default:'';index" json:"phone"`
	Gender     int       `gorm:"type:tinyint;default:0" json:"gender"`
	Country    string    `gorm:"size:50;default:''" json:"country"`
	Province   string    `gorm:"size:50;default:''" json:"province"`
	City       string    `gorm:"size:50;default:''" json:"city"`
	Status     int       `gorm:"type:tinyint;default:1" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

// Court 场地模型
type Court struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Location     string    `gorm:"size:200;default:''" json:"location"`
	Description  string    `gorm:"type:text" json:"description"`
	PricePerHour float64   `gorm:"type:decimal(10,2);not null;default:0.00" json:"price_per_hour"`
	Status       int       `gorm:"type:tinyint;default:1;index" json:"status"`
	Images       string    `gorm:"type:text" json:"images"`     // JSON数组字符串
	Facilities   string    `gorm:"type:text" json:"facilities"` // JSON对象字符串
	SortOrder    int       `gorm:"default:0" json:"sort_order"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Court) TableName() string {
	return "courts"
}

// TimeSlot 时间段模型
type TimeSlot struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CourtID   uint64    `gorm:"column:court_id;not null;index" json:"court_id"`
	Date      string    `gorm:"column:date;type:date;not null;index" json:"date"`
	StartTime string    `gorm:"column:start_time;type:time;not null" json:"start_time"`
	EndTime   string    `gorm:"column:end_time;type:time;not null" json:"end_time"`
	Status    int       `gorm:"type:tinyint;default:1;index" json:"status"` // 1-可用 2-已预定 3-维护中
	BookingID *uint64   `gorm:"column:booking_id" json:"booking_id,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// 关联
	Court Court `gorm:"foreignKey:CourtID" json:"court,omitempty"`
}

func (TimeSlot) TableName() string {
	return "time_slots"
}

// Booking 预定订单模型
type Booking struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderNo       string    `gorm:"column:order_no;uniqueIndex;size:32;not null" json:"order_no"`
	UserID        uint64    `gorm:"column:user_id;not null;index" json:"user_id"`
	CourtID       uint64    `gorm:"column:court_id;not null;index" json:"court_id"`
	TimeSlotID    uint64    `gorm:"column:time_slot_id;not null" json:"time_slot_id"`
	BookingDate   string    `gorm:"column:booking_date;type:date;not null;index" json:"booking_date"`
	StartTime     string    `gorm:"column:start_time;type:time;not null" json:"start_time"`
	EndTime       string    `gorm:"column:end_time;type:time;not null" json:"end_time"`
	DurationHours float64   `gorm:"column:duration_hours;type:decimal(5,2);not null;default:1.00" json:"duration_hours"`
	TotalAmount   float64   `gorm:"column:total_amount;type:decimal(10,2);not null;default:0.00" json:"total_amount"`
	Status        int       `gorm:"type:tinyint;default:1;index" json:"status"` // 1-待支付 2-已支付 3-已取消 4-已完成 5-已退款
	PaymentID     *uint64   `gorm:"column:payment_id" json:"payment_id,omitempty"`
	CancelReason  string    `gorm:"column:cancel_reason;size:200;default:''" json:"cancel_reason"`
	Remark        string    `gorm:"size:500;default:''" json:"remark"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// 关联
	User     *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Court    *Court    `gorm:"foreignKey:CourtID" json:"court,omitempty"`
	TimeSlot *TimeSlot `gorm:"foreignKey:TimeSlotID" json:"time_slot,omitempty"`
	Payment  *Payment  `gorm:"foreignKey:PaymentID" json:"payment,omitempty"`
}

func (Booking) TableName() string {
	return "bookings"
}

// Payment 支付记录模型
type Payment struct {
	ID            uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	BookingID     uint64     `gorm:"column:booking_id;uniqueIndex;not null" json:"booking_id"`
	TransactionID string     `gorm:"column:transaction_id;size:64;default:'';index" json:"transaction_id"`
	PrepayID      string     `gorm:"column:prepay_id;size:128;default:''" json:"prepay_id"`
	Amount        float64    `gorm:"type:decimal(10,2);not null;default:0.00" json:"amount"`
	Status        int        `gorm:"type:tinyint;default:1;index" json:"status"`             // 1-待支付 2-已支付 3-支付失败 4-已退款
	PayType       int        `gorm:"column:pay_type;type:tinyint;default:1" json:"pay_type"` // 1-微信支付
	PaidAt        *time.Time `gorm:"column:paid_at" json:"paid_at,omitempty"`
	NotifyData    string     `gorm:"column:notify_data;type:text" json:"-"` // JSON字符串,不返回给前端
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// 关联
	Booking *Booking `gorm:"foreignKey:BookingID" json:"booking,omitempty"`
}

func (Payment) TableName() string {
	return "payments"
}

// Admin 管理员模型
type Admin struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string     `gorm:"size:50;uniqueIndex;not null" json:"username"`
	PasswordHash string     `gorm:"column:password_hash;size:255;not null" json:"-"`
	RealName     string     `gorm:"column:real_name;size:50;default:''" json:"real_name"`
	Phone        string     `gorm:"size:20;default:''" json:"phone"`
	Role         int        `gorm:"type:tinyint;default:1" json:"role"` // 1-超级管理员 2-普通管理员
	Status       int        `gorm:"type:tinyint;default:1" json:"status"`
	LastLoginAt  *time.Time `gorm:"column:last_login_at" json:"last_login_at,omitempty"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Admin) TableName() string {
	return "admins"
}

// SystemConfig 系统配置模型
type SystemConfig struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigKey   string    `gorm:"column:config_key;uniqueIndex;size:100;not null" json:"config_key"`
	ConfigValue string    `gorm:"column:config_value;type:text" json:"config_value"`
	Description string    `gorm:"size:200;default:''" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (SystemConfig) TableName() string {
	return "system_configs"
}
