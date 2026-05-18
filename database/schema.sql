-- 网球预定系统数据库脚本
-- 创建数据库
CREATE DATABASE IF NOT EXISTS tennis_booking DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE tennis_booking;

-- 1. 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    openid VARCHAR(64) NOT NULL UNIQUE COMMENT '微信openid',
    session_key VARCHAR(128) DEFAULT NULL COMMENT '会话密钥',
    nickname VARCHAR(100) DEFAULT '' COMMENT '昵称',
    avatar VARCHAR(500) DEFAULT '' COMMENT '头像URL',
    phone VARCHAR(20) DEFAULT '' COMMENT '手机号',
    gender TINYINT DEFAULT 0 COMMENT '性别 0-未知 1-男 2-女',
    country VARCHAR(50) DEFAULT '' COMMENT '国家',
    province VARCHAR(50) DEFAULT '' COMMENT '省份',
    city VARCHAR(50) DEFAULT '' COMMENT '城市',
    status TINYINT DEFAULT 1 COMMENT '状态 1-正常 0-禁用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_openid (openid),
    INDEX idx_phone (phone)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 2. 场地表
CREATE TABLE IF NOT EXISTS courts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '场地名称',
    location VARCHAR(200) DEFAULT '' COMMENT '位置描述',
    description TEXT COMMENT '场地描述',
    price_per_hour DECIMAL(10,2) NOT NULL DEFAULT 0.00 COMMENT '每小时价格',
    status TINYINT DEFAULT 1 COMMENT '状态 1-启用 0-停用',
    images TEXT COMMENT '图片URL列表(JSON数组)',
    facilities TEXT COMMENT '设施信息(JSON)',
    sort_order INT DEFAULT 0 COMMENT '排序权重',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='场地表';

-- 3. 时间段表
CREATE TABLE IF NOT EXISTS time_slots (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    court_id BIGINT UNSIGNED NOT NULL COMMENT '场地ID',
    date DATE NOT NULL COMMENT '日期',
    start_time TIME NOT NULL COMMENT '开始时间',
    end_time TIME NOT NULL COMMENT '结束时间',
    status TINYINT DEFAULT 1 COMMENT '状态 1-可用 2-已预定 3-维护中',
    booking_id BIGINT UNSIGNED DEFAULT NULL COMMENT '预定的订单ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (court_id) REFERENCES courts(id) ON DELETE CASCADE,
    UNIQUE KEY uk_court_date_start (court_id, date, start_time),
    INDEX idx_date (date),
    INDEX idx_status (status),
    INDEX idx_court_date (court_id, date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='时间段表';

-- 4. 预定订单表
CREATE TABLE IF NOT EXISTS bookings (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_no VARCHAR(32) NOT NULL UNIQUE COMMENT '订单号',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    court_id BIGINT UNSIGNED NOT NULL COMMENT '场地ID',
    time_slot_id BIGINT UNSIGNED NOT NULL COMMENT '时间段ID',
    booking_date DATE NOT NULL COMMENT '预定日期',
    start_time TIME NOT NULL COMMENT '开始时间',
    end_time TIME NOT NULL COMMENT '结束时间',
    duration_hours DECIMAL(5,2) NOT NULL DEFAULT 1.00 COMMENT '时长(小时)',
    total_amount DECIMAL(10,2) NOT NULL DEFAULT 0.00 COMMENT '总金额',
    status TINYINT DEFAULT 1 COMMENT '状态 1-待支付 2-已支付 3-已取消 4-已完成 5-已退款',
    payment_id BIGINT UNSIGNED DEFAULT NULL COMMENT '支付记录ID',
    cancel_reason VARCHAR(200) DEFAULT '' COMMENT '取消原因',
    remark VARCHAR(500) DEFAULT '' COMMENT '备注',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (court_id) REFERENCES courts(id) ON DELETE CASCADE,
    FOREIGN KEY (time_slot_id) REFERENCES time_slots(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_order_no (order_no),
    INDEX idx_status (status),
    INDEX idx_booking_date (booking_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='预定订单表';

-- 5. 支付记录表
CREATE TABLE IF NOT EXISTS payments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    booking_id BIGINT UNSIGNED NOT NULL UNIQUE COMMENT '订单ID',
    transaction_id VARCHAR(64) DEFAULT '' COMMENT '微信支付交易号',
    prepay_id VARCHAR(128) DEFAULT '' COMMENT '预支付交易会话标识',
    amount DECIMAL(10,2) NOT NULL DEFAULT 0.00 COMMENT '支付金额',
    status TINYINT DEFAULT 1 COMMENT '状态 1-待支付 2-已支付 3-支付失败 4-已退款',
    pay_type TINYINT DEFAULT 1 COMMENT '支付方式 1-微信支付',
    paid_at TIMESTAMP NULL DEFAULT NULL COMMENT '支付时间',
    notify_data TEXT COMMENT '微信支付回调数据(JSON)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE,
    INDEX idx_transaction_id (transaction_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付记录表';

-- 6. 管理员表
CREATE TABLE IF NOT EXISTS admins (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名',
    password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
    real_name VARCHAR(50) DEFAULT '' COMMENT '真实姓名',
    phone VARCHAR(20) DEFAULT '' COMMENT '手机号',
    role TINYINT DEFAULT 1 COMMENT '角色 1-超级管理员 2-普通管理员',
    status TINYINT DEFAULT 1 COMMENT '状态 1-启用 0-禁用',
    last_login_at TIMESTAMP NULL DEFAULT NULL COMMENT '最后登录时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- 7. 系统配置表
CREATE TABLE IF NOT EXISTS system_configs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    config_key VARCHAR(100) NOT NULL UNIQUE COMMENT '配置键',
    config_value TEXT COMMENT '配置值',
    description VARCHAR(200) DEFAULT '' COMMENT '配置描述',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_config_key (config_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统配置表';

-- 插入初始数据

-- 示例场地数据
INSERT INTO courts (name, location, description, price_per_hour, status, images, facilities, sort_order) VALUES
('标准网球场A', '东区1号场地', '标准硬地网球场,灯光照明良好', 80.00, 1, '["https://example.com/court1.jpg"]', '{"surface": "硬地", "lighting": true}', 1),
('标准网球场B', '东区2号场地', '标准硬地网球场,适合训练和比赛', 80.00, 1, '["https://example.com/court2.jpg"]', '{"surface": "硬地", "lighting": true}', 2),
('VIP网球场', '西区VIP区', '高端VIP场地,配备休息区', 150.00, 1, '["https://example.com/court3.jpg"]', '{"surface": "红土", "lighting": true, "rest_area": true}', 3);

-- 默认管理员账号 (密码: admin123, 需要在应用中加密)
INSERT INTO admins (username, password_hash, real_name, role, status) VALUES
('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '系统管理员', 1, 1);

-- 系统配置
INSERT INTO system_configs (config_key, config_value, description) VALUES
('booking_timeout_minutes', '30', '订单超时取消时间(分钟)'),
('max_booking_days', '7', '最多可预定天数'),
('business_start_time', '08:00', '营业开始时间'),
('business_end_time', '22:00', '营业结束时间'),
('slot_duration_minutes', '60', '每个时间段时长(分钟)');
