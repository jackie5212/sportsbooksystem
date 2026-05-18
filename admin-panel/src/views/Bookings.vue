<template>
  <div class="page-container">
    <h2 class="page-title">订单管理</h2>
    
    <el-card>
      <div class="toolbar">
        <el-select v-model="statusFilter" placeholder="订单状态" style="width: 150px" clearable>
          <el-option label="全部" :value="null" />
          <el-option label="待支付" :value="1" />
          <el-option label="已完成" :value="2" />
          <el-option label="已取消" :value="3" />
        </el-select>
        <el-input
          v-model="searchText"
          placeholder="搜索订单号或用户名"
          style="width: 300px; margin-left: auto;"
          clearable
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </div>
      
      <el-table :data="bookings" stripe v-loading="loading">
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
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="handleView(row)">
              查看详情
            </el-button>
            <el-button 
              v-if="row.status === 1"
              type="danger" 
              link 
              size="small"
              @click="handleCancel(row)"
            >
              取消订单
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="pagination">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          layout="total, sizes, prev, pager, next"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const statusFilter = ref(null)
const searchText = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(3)

const bookings = ref([
  {
    order_no: 'TB202605150001',
    user_name: '郑钦文',
    court_name: '标准网球场A',
    booking_date: '2026-05-15',
    time_range: '08:00-09:00',
    total_amount: 80,
    status: 1,
    created_at: '2026-05-15 07:30:00'
  },
  {
    order_no: 'TB202605150002',
    user_name: '张德培',
    court_name: 'VIP网球场',
    booking_date: '2026-05-15',
    time_range: '14:00-15:00',
    total_amount: 150,
    status: 2,
    created_at: '2026-05-15 08:15:00'
  },
  {
    order_no: 'TB202605140003',
    user_name: '桑普拉斯',
    court_name: '标准网球场B',
    booking_date: '2026-05-14',
    time_range: '10:00-11:00',
    total_amount: 80,
    status: 3,
    created_at: '2026-05-14 09:20:00'
  }
])

const getStatusType = (status) => {
  const types = { 1: 'warning', 2: 'success', 3: 'info' }
  return types[status] || ''
}

const getStatusText = (status) => {
  const texts = { 1: '待支付', 2: '已完成', 3: '已取消' }
  return texts[status] || '未知'
}

const handleView = (row) => {
  ElMessage.info(`查看订单: ${row.order_no}`)
}

const handleCancel = (row) => {
  ElMessageBox.confirm('确定要取消该订单吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    row.status = 3
    ElMessage.success('订单已取消')
  })
}

onMounted(() => {
  // TODO: 加载数据
})
</script>

<style scoped>
.page-container {
  padding: 20px;
}

.page-title {
  margin-bottom: 20px;
  color: #303133;
}

.toolbar {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
