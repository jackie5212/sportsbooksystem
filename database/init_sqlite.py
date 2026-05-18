#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
网球预定系统 - SQLite数据库初始化脚本
无需安装MySQL,直接运行即可创建数据库
"""

import sqlite3
import os
from datetime import datetime, timedelta

DB_FILE = 'tennis_booking.db'

def init_database():
    """初始化数据库"""
    print("=" * 60)
    print("网球预定系统 - SQLite数据库初始化")
    print("=" * 60)
    
    # 删除旧数据库文件(如果存在)
    if os.path.exists(DB_FILE):
        print(f"\n发现现有数据库文件: {DB_FILE}")
        choice = input("是否删除并重新创建? (y/n): ")
        if choice.lower() == 'y':
            os.remove(DB_FILE)
            print("已删除旧数据库")
        else:
            print("取消操作")
            return
    
    # 连接数据库(自动创建)
    conn = sqlite3.connect(DB_FILE)
    cursor = conn.cursor()
    
    print(f"\n正在创建数据库: {DB_FILE}...")
    
    # 读取SQL文件
    sql_file = 'database/schema_sqlite.sql'
    if not os.path.exists(sql_file):
        print(f"错误: SQL文件不存在: {sql_file}")
        return
    
    with open(sql_file, 'r', encoding='utf-8') as f:
        sql_script = f.read()
    
    # 执行SQL脚本
    try:
        cursor.executescript(sql_script)
        conn.commit()
        print("✓ 数据库表结构创建成功")
    except Exception as e:
        print(f"✗ 创建表结构失败: {e}")
        conn.rollback()
        return
    
    # 生成时间段数据
    print("\n正在生成时间段数据...")
    generate_time_slots(cursor, conn)
    
    # 验证数据
    print("\n验证数据...")
    verify_data(cursor)
    
    # 关闭连接
    conn.close()
    
    print("\n" + "=" * 60)
    print("✓ 数据库初始化完成!")
    print("=" * 60)
    print(f"\n数据库文件: {os.path.abspath(DB_FILE)}")
    print("\n可以使用以下工具查看数据库:")
    print("  - DB Browser for SQLite (推荐)")
    print("  - SQLite Studio")
    print("  - VS Code SQLite插件")
    print("\n默认管理员账号:")
    print("  用户名: admin")
    print("  密码: admin123")
    print("=" * 60)

def generate_time_slots(cursor, conn):
    """生成未来7天的时间段数据"""
    courts = [(1,), (2,), (3,)]  # 3个场地ID
    
    # 时间段定义 (每小时一个时段)
    time_ranges = [
        ('08:00:00', '09:00:00'),
        ('09:00:00', '10:00:00'),
        ('10:00:00', '11:00:00'),
        ('11:00:00', '12:00:00'),
        ('14:00:00', '15:00:00'),
        ('15:00:00', '16:00:00'),
        ('16:00:00', '17:00:00'),
        ('17:00:00', '18:00:00'),
        ('19:00:00', '20:00:00'),
        ('20:00:00', '21:00:00'),
    ]
    
    today = datetime.now().date()
    count = 0
    
    for day in range(7):  # 未来7天
        date = today + timedelta(days=day)
        date_str = date.strftime('%Y-%m-%d')
        
        for court in courts:
            court_id = court[0]
            
            for start_time, end_time in time_ranges:
                try:
                    cursor.execute(
                        "INSERT OR IGNORE INTO time_slots (court_id, date, start_time, end_time, status) VALUES (?, ?, ?, ?, 1)",
                        (court_id, date_str, start_time, end_time)
                    )
                    count += 1
                except Exception as e:
                    print(f"插入时间段失败: {e}")
    
    conn.commit()
    print(f"✓ 已生成 {count} 个时间段")

def verify_data(cursor):
    """验证数据是否正确插入"""
    # 检查表数量
    cursor.execute("SELECT name FROM sqlite_master WHERE type='table'")
    tables = cursor.fetchall()
    print(f"✓ 创建了 {len(tables)} 个表")
    
    # 检查各表数据量
    table_checks = [
        ('users', '用户'),
        ('courts', '场地'),
        ('time_slots', '时间段'),
        ('bookings', '订单'),
        ('payments', '支付'),
        ('admins', '管理员'),
        ('system_configs', '系统配置'),
    ]
    
    for table, name in table_checks:
        cursor.execute(f"SELECT COUNT(*) FROM {table}")
        count = cursor.fetchone()[0]
        print(f"  - {name} ({table}): {count} 条记录")

if __name__ == '__main__':
    init_database()
