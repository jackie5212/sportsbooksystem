<template>
  <div class="dashboard">
    <h2 class="page-title">仪表盘</h2>
    
    <!-- 统计卡片 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #409EFF;">
              <el-icon size="30"><Grid /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.totalCourts || 0 }}</div>
              <div class="stat-label">场地总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #67C23A;">
              <el-icon size="30"><Document /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.todayBookings || 0 }}</div>
              <div class="stat-label">今日订单</div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #E6A23C;">
              <el-icon size="30"><User /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stats.totalUsers || 0 }}</div>
              <div class="stat-label">用户总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-icon" style="background-color: #F56C6C;">
              <el-icon size="30"><Money /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">¥{{ stats.todayRevenue || 0 }}</div>
              <div class="stat-label">今日收入</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 图表区域 -->
    <el-row :gutter="20" class="chart-row">
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>近7天订单趋势</span>
          </template>
          <div ref="orderChartRef" style="height: 300px;"></div>
        </el-card>
      </el-col>
      
      <el-col :span="12">
        <el-card>
          <template #header>
            <span>场地使用率</span>
          </template>
          <div ref="courtChartRef" style="height: 300px;"></div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 最近订单 -->
    <el-card class="recent-bookings">
      <template #header>
        <div class="card-header">
          <span>最近订单</span>
          <el-button type="primary" link @click="$router.push('/bookings')">查看全部</el-button>
        </div>
      </template>
      
      <el-table :data="recentBookings" stripe>
        <el-table-column prop="order_no" label="订单号" width="180" />
        <el-table-column prop="user_name" label="用户" width="120" />
        <el-table-column prop="court_name" label="场地" width="150" />
        <el-table-column prop="booking_date" label="日期" width="120" />
        <el-table-column prop="time_range" label="时间" width="150" />
        <el-table-column prop="total_amount" label="金额" width="100">
          <template #default="{ row }">
            ¥{{ row.total_amount }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as echarts from 'echarts'

const stats = ref({
  totalCourts: 3,
  todayBookings: 12,
  totalUsers: 5,
  todayRevenue: 960
})

const recentBookings = ref([
  {
    order_no: 'TB202605150001',
    user_name: '郑钦文',
    court_name: '标准网球场A',
    booking_date: '2026-05-15',
    time_range: '08:00-09:00',
    total_amount: 80,
    status: 1
  },
  {
    order_no: 'TB202605150002',
    user_name: '张德培',
    court_name: 'VIP网球场',
    booking_date: '2026-05-15',
    time_range: '14:00-15:00',
    total_amount: 150,
    status: 2
  },
  {
    order_no: 'TB202605140003',
    user_name: '桑普拉斯',
    court_name: '标准网球场B',
    booking_date: '2026-05-14',
    time_range: '10:00-11:00',
    total_amount: 80,
    status: 3
  }
])

const orderChartRef = ref(null)
const courtChartRef = ref(null)

// 初始化订单趋势图
const initOrderChart = () => {
  const chart = echarts.init(orderChartRef.value)
  const option = {
    tooltip: { trigger: 'axis' },
    xAxis: {
      type: 'category',
      data: ['5/9', '5/10', '5/11', '5/12', '5/13', '5/14', '5/15']
    },
    yAxis: { type: 'value' },
    series: [{
      data: [8, 12, 15, 10, 14, 18, 12],
      type: 'line',
      smooth: true,
      itemStyle: { color: '#409EFF' },
      areaStyle: {
        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          { offset: 0, color: 'rgba(64, 158, 255, 0.3)' },
          { offset: 1, color: 'rgba(64, 158, 255, 0.1)' }
        ])
      }
    }]
  }
  chart.setOption(option)
}

// 初始化场地使用率图
const initCourtChart = () => {
  const chart = echarts.init(courtChartRef.value)
  const option = {
    tooltip: { trigger: 'item' },
    legend: { bottom: '5%', left: 'center' },
    series: [{
      type: 'pie',
      radius: ['40%', '70%'],
      data: [
        { value: 65, name: '已使用', itemStyle: { color: '#67C23A' } },
        { value: 35, name: '空闲', itemStyle: { color: '#E6E6E6' } }
      ]
    }]
  }
  chart.setOption(option)
}

const getStatusType = (status) => {
  const types = { 1: 'warning', 2: 'success', 3: 'info' }
  return types[status] || ''
}

const getStatusText = (status) => {
  const texts = { 1: '待支付', 2: '已完成', 3: '已取消' }
  return texts[status] || '未知'
}

onMounted(() => {
  initOrderChart()
  initCourtChart()
})
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.page-title {
  margin-bottom: 20px;
  color: #303133;
}

.stats-row {
  margin-bottom: 20px;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 15px;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 5px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.chart-row {
  margin-bottom: 20px;
}

.recent-bookings {
  margin-top: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
