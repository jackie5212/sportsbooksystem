-- ============================================
-- 网球预定系统 - SQLite数据库脚本
-- 适用于开发测试,无需安装MySQL
-- ============================================

-- 启用外键支持
PRAGMA foreign_keys = ON;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    openid TEXT NOT NULL UNIQUE,
    session_key TEXT,
    nickname TEXT DEFAULT '',
    avatar TEXT DEFAULT '',
    phone TEXT DEFAULT '',
    gender INTEGER DEFAULT 0,
    country TEXT DEFAULT '',
    province TEXT DEFAULT '',
    city TEXT DEFAULT '',
    status INTEGER DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_openid ON users(openid);
CREATE INDEX IF NOT EXISTS idx_users_phone ON users(phone);

-- 场地表
CREATE TABLE IF NOT EXISTS courts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    location TEXT NOT NULL,
    description TEXT,
    price_per_hour REAL NOT NULL DEFAULT 0,
    images TEXT,
    status INTEGER DEFAULT 1,
    sort_order INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_courts_status ON courts(status);

-- 时间段表
CREATE TABLE IF NOT EXISTS time_slots (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    court_id INTEGER NOT NULL,
    date TEXT NOT NULL,
    start_time TEXT NOT NULL,
    end_time TEXT NOT NULL,
    status INTEGER DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (court_id) REFERENCES courts(id) ON DELETE CASCADE,
    UNIQUE(court_id, date, start_time)
);

CREATE INDEX IF NOT EXISTS idx_time_slots_court_date ON time_slots(court_id, date);
CREATE INDEX IF NOT EXISTS idx_time_slots_status ON time_slots(status);

-- 订单表
CREATE TABLE IF NOT EXISTS bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    order_no TEXT NOT NULL UNIQUE,
    user_id INTEGER NOT NULL,
    court_id INTEGER NOT NULL,
    time_slot_id INTEGER NOT NULL,
    booking_date TEXT NOT NULL,
    start_time TEXT NOT NULL,
    end_time TEXT NOT NULL,
    duration_hours REAL NOT NULL DEFAULT 1,
    total_amount REAL NOT NULL DEFAULT 0,
    status INTEGER DEFAULT 1,
    remark TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (court_id) REFERENCES courts(id) ON DELETE CASCADE,
    FOREIGN KEY (time_slot_id) REFERENCES time_slots(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_bookings_user ON bookings(user_id);
CREATE INDEX IF NOT EXISTS idx_bookings_court ON bookings(court_id);
CREATE INDEX IF NOT EXISTS idx_bookings_status ON bookings(status);
CREATE INDEX IF NOT EXISTS idx_bookings_date ON bookings(booking_date);

-- 支付表
CREATE TABLE IF NOT EXISTS payments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    booking_id INTEGER NOT NULL UNIQUE,
    transaction_id TEXT,
    amount REAL NOT NULL,
    payment_method TEXT DEFAULT 'wechat',
    status INTEGER DEFAULT 1,
    paid_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_payments_booking ON payments(booking_id);
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(status);

-- 管理员表
CREATE TABLE IF NOT EXISTS admins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    nickname TEXT DEFAULT '',
    role TEXT DEFAULT 'admin',
    status INTEGER DEFAULT 1,
    last_login_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 系统配置表
CREATE TABLE IF NOT EXISTS system_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    config_key TEXT NOT NULL UNIQUE,
    config_value TEXT,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- 插入初始数据
-- ============================================

-- 插入默认管理员 (密码: admin123, 实际应使用加密后的密码)
INSERT OR IGNORE INTO admins (username, password, nickname, role) 
VALUES ('admin', 'admin123', '系统管理员', 'super_admin');

-- 插入示例场地
INSERT OR IGNORE INTO courts (name, location, description, price_per_hour, status, sort_order) VALUES
('标准网球场A', '东区1号场地', '标准硬地网球场,灯光照明良好', 80.00, 1, 1),
('标准网球场B', '东区2号场地', '标准硬地网球场,适合训练和比赛', 80.00, 1, 2),
('VIP网球场', '西区VIP区', '高端VIP场地,配备休息区', 150.00, 1, 3);

-- 插入系统配置
INSERT OR IGNORE INTO system_configs (config_key, config_value, description) VALUES
('site_name', '网球预定系统', '网站名称'),
('contact_phone', '400-123-4567', '客服电话'),
('business_hours', '08:00-22:00', '营业时间'),
('booking_advance_days', '7', '可提前预定天数'),
('cancel_before_hours', '2', '取消订单提前小时数');

-- 生成未来7天的时间段数据 (为每个场地每天生成时间段)
-- 这里只生成示例数据,实际应该通过程序生成
