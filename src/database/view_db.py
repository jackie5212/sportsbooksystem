#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
快速查看SQLite数据库内容
"""

import sqlite3

def show_database_info():
    conn = sqlite3.connect('tennis_booking.db')
    cursor = conn.cursor()
    
    print("=" * 60)
    print("网球预定系统 - 数据库内容概览")
    print("=" * 60)
    
    # 1. 查看所有表
    print("\n📋 数据库表:")
    cursor.execute("SELECT name FROM sqlite_master WHERE type='table' ORDER BY name")
    tables = cursor.fetchall()
    for table in tables:
        print(f"  - {table[0]}")
    
    # 2. 查看场地信息
    print("\n🏟️ 场地列表:")
    cursor.execute("SELECT id, name, location, price_per_hour, status FROM courts")
    courts = cursor.fetchall()
    for court in courts:
        status = "启用" if court[4] == 1 else "禁用"
        print(f"  [{court[0]}] {court[1]} - {court[2]} - ¥{court[3]}/小时 - {status}")
    
    # 3. 查看管理员
    print("\n👤 管理员账号:")
    cursor.execute("SELECT username, nickname, role FROM admins")
    admins = cursor.fetchall()
    for admin in admins:
        print(f"  用户名: {admin[0]}, 昵称: {admin[1]}, 角色: {admin[2]}")
    
    # 4. 查看时间段统计
    print("\n⏰ 时间段统计:")
    cursor.execute("""
        SELECT date, COUNT(*) as count 
        FROM time_slots 
        GROUP BY date 
        ORDER BY date
    """)
    slots = cursor.fetchall()
    for slot in slots:
        print(f"  {slot[0]}: {slot[1]}个时段")
    
    # 5. 查看系统配置
    print("\n⚙️ 系统配置:")
    cursor.execute("SELECT config_key, config_value FROM system_configs")
    configs = cursor.fetchall()
    for config in configs:
        print(f"  {config[0]}: {config[1]}")
    
    # 6. 数据统计
    print("\n📊 数据统计:")
    cursor.execute("SELECT COUNT(*) FROM users")
    print(f"  用户总数: {cursor.fetchone()[0]}")
    
    cursor.execute("SELECT COUNT(*) FROM bookings")
    print(f"  订单总数: {cursor.fetchone()[0]}")
    
    cursor.execute("SELECT COUNT(*) FROM time_slots WHERE status = 1")
    print(f"  可用时段: {cursor.fetchone()[0]}")
    
    conn.close()
    
    print("\n" + "=" * 60)

if __name__ == '__main__':
    show_database_info()
