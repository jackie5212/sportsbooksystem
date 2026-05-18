<template>
  <div class="page-container">
    <h2 class="page-title">场地管理</h2>
    
    <el-card>
      <div class="toolbar">
        <el-button type="primary" @click="handleAdd">
          <el-icon><Plus /></el-icon>
          新增场地
        </el-button>
        <el-input
          v-model="searchText"
          placeholder="搜索场地名称"
          style="width: 300px; margin-left: auto;"
          clearable
          @clear="loadCourts"
          @keyup.enter="loadCourts"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </div>
      
      <el-table :data="courts" stripe v-loading="loading">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="场地名称" width="150" />
        <el-table-column prop="location" label="位置" width="150" />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="price_per_hour" label="价格(元/小时)" width="130">
          <template #default="{ row }">
            ¥{{ row.price_per_hour }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="handleEdit(row)">
              编辑
            </el-button>
            <el-button 
              :type="row.status === 1 ? 'warning' : 'success'" 
              link 
              size="small"
              @click="handleToggleStatus(row)"
            >
              {{ row.status === 1 ? '禁用' : '启用' }}
            </el-button>
            <el-button type="danger" link size="small" @click="handleDelete(row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="pagination">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          :page-sizes="[10, 20, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="loadCourts"
          @current-change="loadCourts"
        />
      </div>
    </el-card>
    
    <!-- 编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑场地' : '新增场地'"
      width="600px"
    >
      <el-form :model="formData" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="场地名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入场地名称" />
        </el-form-item>
        
        <el-form-item label="位置" prop="location">
          <el-input v-model="formData.location" placeholder="请输入位置" />
        </el-form-item>
        
        <el-form-item label="描述" prop="description">
          <el-input 
            v-model="formData.description" 
            type="textarea"
            :rows="3"
            placeholder="请输入描述"
          />
        </el-form-item>
        
        <el-form-item label="价格" prop="price_per_hour">
          <el-input-number 
            v-model="formData.price_per_hour" 
            :min="0"
            :precision="2"
            style="width: 100%"
          />
        </el-form-item>
        
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitLoading">
          确定
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const searchText = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const formRef = ref(null)

// 模拟数据
const courts = ref([
  { id: 1, name: '标准网球场A', location: '东区1号场地', description: '标准硬地网球场,灯光照明良好', price_per_hour: 80, status: 1 },
  { id: 2, name: '标准网球场B', location: '东区2号场地', description: '标准硬地网球场,适合训练和比赛', price_per_hour: 80, status: 1 },
  { id: 3, name: 'VIP网球场', location: '西区VIP区', description: '高端VIP场地,配备休息区', price_per_hour: 150, status: 1 }
])

const formData = reactive({
  id: null,
  name: '',
  location: '',
  description: '',
  price_per_hour: 0,
  status: 1
})

const rules = {
  name: [{ required: true, message: '请输入场地名称', trigger: 'blur' }],
  location: [{ required: true, message: '请输入位置', trigger: 'blur' }],
  price_per_hour: [{ required: true, message: '请输入价格', trigger: 'blur' }]
}

const loadCourts = () => {
  loading.value = true
  // TODO: 调用API
  setTimeout(() => {
    loading.value = false
  }, 500)
}

const handleAdd = () => {
  isEdit.value = false
  Object.assign(formData, {
    id: null,
    name: '',
    location: '',
    description: '',
    price_per_hour: 0,
    status: 1
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  Object.assign(formData, row)
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      submitLoading.value = true
      try {
        // TODO: 调用API
        ElMessage.success(isEdit.value ? '修改成功' : '添加成功')
        dialogVisible.value = false
        loadCourts()
      } catch (error) {
        console.error('提交失败:', error)
      } finally {
        submitLoading.value = false
      }
    }
  })
}

const handleToggleStatus = (row) => {
  const action = row.status === 1 ? '禁用' : '启用'
  ElMessageBox.confirm(`确定要${action}该场地吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    // TODO: 调用API
    row.status = row.status === 1 ? 0 : 1
    ElMessage.success(`${action}成功`)
  })
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该场地吗?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(() => {
    // TODO: 调用API
    const index = courts.value.findIndex(item => item.id === row.id)
    if (index > -1) {
      courts.value.splice(index, 1)
      ElMessage.success('删除成功')
    }
  })
}

onMounted(() => {
  loadCourts()
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
