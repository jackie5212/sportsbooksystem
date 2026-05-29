-- 多店铺管理系统数据库迁移脚本
-- 此脚本会在数据库初始化时自动执行
SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

USE tennis_booking;

-- 1. 创建店铺表
CREATE TABLE IF NOT EXISTS shops (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    shop_name VARCHAR(100) NOT NULL COMMENT '店铺名称',
    shop_code VARCHAR(50) UNIQUE NOT NULL COMMENT '店铺编码',
    description TEXT COMMENT '店铺描述',
    address VARCHAR(200) COMMENT '店铺地址',
    phone VARCHAR(20) COMMENT '联系电话',
    manager_name VARCHAR(50) COMMENT '负责人',
    manager_phone VARCHAR(20) COMMENT '负责人电话',
    status TINYINT DEFAULT 1 COMMENT '状态：1-营业中 2-停业 3-装修中',
    business_hours VARCHAR(100) COMMENT '营业时间',
    latitude DECIMAL(10, 7) COMMENT '纬度',
    longitude DECIMAL(10, 7) COMMENT '经度',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_status (status),
    INDEX idx_shop_code (shop_code)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='店铺表';

-- 2. 修改courts表，添加shop_id字段
ALTER TABLE courts 
ADD COLUMN IF NOT EXISTS shop_id BIGINT UNSIGNED COMMENT '所属店铺ID' AFTER id,
ADD INDEX IF NOT EXISTS idx_shop_id (shop_id);

-- 3. 修改bookings表，添加shop_id字段
ALTER TABLE bookings 
ADD COLUMN IF NOT EXISTS shop_id BIGINT UNSIGNED COMMENT '所属店铺ID' AFTER id,
ADD INDEX IF NOT EXISTS idx_shop_id (shop_id);

-- 4. 修改admins表，添加shop_id字段（支持店铺管理员）
ALTER TABLE admins 
ADD COLUMN IF NOT EXISTS shop_id BIGINT UNSIGNED COMMENT '所属店铺ID，NULL表示超级管理员' AFTER role,
ADD INDEX IF NOT EXISTS idx_shop_id (shop_id);

-- 5. 修改system_configs表，添加更多字段
ALTER TABLE system_configs 
ADD COLUMN IF NOT EXISTS config_type VARCHAR(20) DEFAULT 'string' COMMENT '配置类型：string/number/boolean/json' AFTER config_value,
ADD COLUMN IF NOT EXISTS is_encrypted TINYINT DEFAULT 0 COMMENT '是否加密：0-否 1-是' AFTER description,
ADD COLUMN IF NOT EXISTS category VARCHAR(50) COMMENT '配置分类：wechat/payment/system' AFTER is_encrypted,
ADD INDEX IF NOT EXISTS idx_category (category);

-- 6. 插入默认店铺数据
INSERT INTO shops (shop_name, shop_code, description, address, phone, manager_name, manager_phone, status, business_hours) VALUES
('总店', 'SHOP001', '网球预定系统总店', '北京市朝阳区网球中心', '010-12345678', '张经理', '13800138000', 1, '08:00-22:00'),
('东区分店', 'SHOP002', '东区分店', '北京市东城区体育馆', '010-87654321', '李经理', '13900139000', 1, '09:00-21:00')
ON DUPLICATE KEY UPDATE shop_name=shop_name;

-- 7. 更新现有场地，关联到默认店铺
UPDATE courts SET shop_id = 1 WHERE shop_id IS NULL;

-- 8. 更新现有订单，关联到对应店铺（根据场地的shop_id）
UPDATE bookings b
INNER JOIN courts c ON b.court_id = c.id
SET b.shop_id = c.shop_id
WHERE b.shop_id IS NULL;

-- 9. 插入微信支付配置示例
INSERT INTO system_configs (config_key, config_value, config_type, description, is_encrypted, category) VALUES
('wechat_appid', '', 'string', '微信小程序AppID', 0, 'wechat'),
('wechat_appsecret', '', 'string', '微信小程序AppSecret', 1, 'wechat'),
('wechat_mch_id', '', 'string', '微信支付商户号', 0, 'payment'),
('wechat_api_key', '', 'string', '微信支付API密钥', 1, 'payment'),
('wechat_cert_path', '', 'string', '微信支付证书路径', 0, 'payment'),
('wechat_key_path', '', 'string', '微信支付私钥路径', 0, 'payment'),
('payment_notify_url', 'https://yourdomain.com/api/payments/notify', 'string', '支付回调地址', 0, 'payment'),
('system_name', '网球预定系统', 'string', '系统名称', 0, 'system')
ON DUPLICATE KEY UPDATE config_value=config_value;

-- 10. 创建店铺统计表（可选，用于缓存统计数据）
CREATE TABLE IF NOT EXISTS shop_statistics (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    shop_id BIGINT UNSIGNED NOT NULL,
    stat_date DATE NOT NULL COMMENT '统计日期',
    total_orders INT DEFAULT 0 COMMENT '订单总数',
    paid_orders INT DEFAULT 0 COMMENT '已支付订单数',
    total_revenue DECIMAL(10, 2) DEFAULT 0.00 COMMENT '总收入',
    total_users INT DEFAULT 0 COMMENT '用户数',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_shop_date (shop_id, stat_date),
    INDEX idx_shop_id (shop_id),
    INDEX idx_stat_date (stat_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='店铺统计表';
