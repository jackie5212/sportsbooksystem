<template>
  <div class="page-container">
    <h2 class="page-title">时间段管理</h2>
    
    <el-card>
      <div class="toolbar">
        <el-select v-model="selectedCourt" placeholder="选择场地" style="width: 200px">
          <el-option label="全部场地" :value="null" />
          <el-option label="标准网球场A" :value="1" />
          <el-option label="标准网球场B" :value="2" />
          <el-option label="VIP网球场" :value="3" />
        </el-select>
        <el-date-picker
          v-model="selectedDate"
          type="date"
          placeholder="选择日期"
          style="margin-left: 10px"
        />
      </div>
      
      <el-table :data="timeSlots" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="court_name" label="场地" width="150" />
        <el-table-column prop="date" label="日期" width="120" />
        <el-table-column prop="start_time" label="开始时间" width="120" />
        <el-table-column prop="end_time" label="结束时间" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : row.status === 2 ? 'warning' : 'info'">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button 
              v-if="row.status === 1"
              type="warning" 
              link 
              size="small"
              @click="handleBlock(row)"
            >
              锁定
            </el-button>
            <el-button 
              v-if="row.status === 3"
              type="success" 
              link 
              size="small"
              @click="handleUnlock(row)"
            >
              解锁
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const selectedCourt = ref(null)
const selectedDate = ref(new Date())

const timeSlots = ref([
  { id: 1, court_name: '标准网球场A', date: '2026-05-15', start_time: '08:00:00', end_time: '09:00:00', status: 1 },
  { id: 2, court_name: '标准网球场A', date: '2026-05-15', start_time: '09:00:00', end_time: '10:00:00', status: 2 },
  { id: 3, court_name: '标准网球场A', date: '2026-05-15', start_time: '10:00:00', end_time: '11:00:00', status: 1 },
  { id: 4, court_name: 'VIP网球场', date: '2026-05-15', start_time: '14:00:00', end_time: '15:00:00', status: 3 }
])

const getStatusText = (status) => {
  const texts = { 1: '可预定', 2: '已预定', 3: '已锁定' }
  return texts[status] || '未知'
}

const handleBlock = (row) => {
  row.status = 3
  ElMessage.success('时间段已锁定')
}

const handleUnlock = (row) => {
  row.status = 1
  ElMessage.success('时间段已解锁')
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
</style>
