<template>
  <div class="page-container">
    <h2 class="page-title">用户管理</h2>
    
    <el-card>
      <el-table :data="users" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="nickname" label="昵称" width="150" />
        <el-table-column prop="phone" label="手机号" width="150" />
        <el-table-column prop="openid" label="OpenID" min-width="200" show-overflow-tooltip />
        <el-table-column prop="booking_count" label="预定次数" width="100" />
        <el-table-column prop="created_at" label="注册时间" width="180" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="handleView(row)">
              查看详情
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
import { ElMessage } from 'element-plus'

const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(5)

const users = ref([
  { id: 1, nickname: '郑钦文', phone: '138****1234', openid: 'oXXXX1', booking_count: 5, created_at: '2026-05-01 10:00:00' },
  { id: 2, nickname: '张德培', phone: '139****5678', openid: 'oXXXX2', booking_count: 3, created_at: '2026-05-05 14:30:00' },
  { id: 3, nickname: '桑普拉斯', phone: '137****9012', openid: 'oXXXX3', booking_count: 8, created_at: '2026-05-10 09:15:00' },
  { id: 4, nickname: '阿加西', phone: '136****3456', openid: 'oXXXX4', booking_count: 12, created_at: '2026-05-12 16:45:00' },
  { id: 5, nickname: '纳达尔', phone: '135****7890', openid: 'oXXXX5', booking_count: 15, created_at: '2026-05-13 11:20:00' }
])

const handleView = (row) => {
  ElMessage.info(`查看用户: ${row.nickname}`)
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

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
