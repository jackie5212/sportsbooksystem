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
        <el-table-column label="图片" width="120">
          <template #default="{ row }">
            <div v-if="row.images && row.images.length > 0" class="table-image">
              <el-image 
                :src="row.images[0]" 
                fit="cover"
                style="width: 80px; height: 60px; border-radius: 4px;"
                :preview-src-list="row.images"
                preview-teleported
              />
              <span v-if="row.images.length > 1" class="image-count">+{{ row.images.length - 1 }}</span>
            </div>
            <span v-else class="no-image">暂无图片</span>
          </template>
        </el-table-column>
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
        
        <el-form-item label="场地图片">
          <div class="image-upload-section">
            <el-upload
              action="#"
              list-type="picture-card"
              :file-list="imageList"
              :auto-upload="false"
              :on-change="handleImageChange"
              :on-remove="handleImageRemove"
              :limit="5"
              accept="image/*"
            >
              <el-icon><Plus /></el-icon>
            </el-upload>
            <div class="upload-tip">最多上传5张图片，支持jpg、png格式</div>
          </div>
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
import { getCourts, createCourt, updateCourt, deleteCourt } from '@/api/admin'

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const searchText = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const formRef = ref(null)

const courts = ref([])

const formData = reactive({
  id: null,
  name: '',
  location: '',
  description: '',
  price_per_hour: 0,
  status: 1,
  images: [] // 图片URL数组
})

const imageList = ref([]) // 上传组件显示的图片列表

const rules = {
  name: [{ required: true, message: '请输入场地名称', trigger: 'blur' }],
  location: [{ required: true, message: '请输入位置', trigger: 'blur' }],
  price_per_hour: [{ required: true, message: '请输入价格', trigger: 'blur' }]
}

// 处理图片URL，将相对路径转换为完整URL
const getImageUrl = (url) => {
  if (!url) return ''
  // 如果已经是完整URL（http/https开头），直接返回
  if (url.startsWith('http://') || url.startsWith('https://')) {
    return url
  }
  // 如果是相对路径（以/开头），添加后端地址
  if (url.startsWith('/')) {
    return `http://localhost:8080${url}`
  }
  // 其他情况也添加后端地址
  return `http://localhost:8080/${url}`
}

const loadCourts = async () => {
  loading.value = true
  try {
    const response = await getCourts({
      page: currentPage.value,
      page_size: pageSize.value,
      search: searchText.value
    })
    
    console.log('获取场地列表响应:', response)
    
    // request.js 的响应拦截器已经返回了 res.data
    // 所以 response 直接就是 { list, total, page, page_size }
    courts.value = (response.list || []).map(court => {
      // 如果 images 是字符串，解析为数组
      if (typeof court.images === 'string') {
        try {
          court.images = JSON.parse(court.images)
        } catch (e) {
          console.error('解析场地图片失败:', e)
          court.images = []
        }
      }
      // 确保 images 是数组
      if (!Array.isArray(court.images)) {
        court.images = []
      }
      // 处理图片URL，将相对路径转换为完整URL
      court.images = court.images.map(img => getImageUrl(img))
      return court
    })
    total.value = response.total || 0
    
    console.log('处理后的场地列表:', courts.value)
    console.log('总数:', total.value)
  } catch (error) {
    console.error('加载场地失败:', error)
    ElMessage.error(error.message || '网络错误，请检查后端服务')
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  isEdit.value = false
  Object.assign(formData, {
    id: null,
    name: '',
    location: '',
    description: '',
    price_per_hour: 0,
    status: 1,
    images: []
  })
  imageList.value = []
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  Object.assign(formData, row)
  // 将图片URL数组转换为上传组件需要的格式
  console.log('编辑场地 - row.images:', row.images)
  console.log('编辑场地 - images类型:', typeof row.images)
  
  let imagesArray = []
  // 处理images可能是JSON字符串的情况
  if (typeof row.images === 'string') {
    try {
      imagesArray = JSON.parse(row.images)
    } catch (e) {
      console.error('解析图片数据失败:', e)
      imagesArray = []
    }
  } else if (Array.isArray(row.images)) {
    imagesArray = row.images
  }
  
  console.log('编辑场地 - 解析后的图片数组:', imagesArray)
  
  if (imagesArray && imagesArray.length > 0) {
    imageList.value = imagesArray.map((url, index) => ({
      name: `image-${index}`,
      url: getImageUrl(url)  // 转换为完整URL
    }))
    console.log('编辑场地 - imageList:', imageList.value)
  } else {
    imageList.value = []
    console.log('编辑场地 - 没有图片数据')
  }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (valid) {
      submitLoading.value = true
      try {
        // 处理图片上传
        const uploadedUrls = []
        for (const file of imageList.value) {
          // 如果是新上传的文件（有raw属性且url是blob或不存在）
          if (file.raw && (!file.url || file.url.startsWith('blob:'))) {
            console.log('上传新图片:', file.name)
            const url = await uploadImage(file.raw)
            uploadedUrls.push(url)
          } else if (file.url && !file.url.startsWith('blob:')) {
            // 已存在的图片URL
            console.log('保留已有图片:', file.url)
            uploadedUrls.push(file.url)
          }
        }
        formData.images = uploadedUrls
        console.log('最终提交的图片数组:', uploadedUrls)
        
        // 调用API保存场地信息（request 拦截器已解包 data，成功即表示 code=200）
        if (isEdit.value) {
          await updateCourt(formData.id, formData)
        } else {
          await createCourt(formData)
        }

        ElMessage.success(isEdit.value ? '修改成功' : '添加成功')
        dialogVisible.value = false
        loadCourts()
      } catch (error) {
        console.error('提交失败:', error)
        ElMessage.error(error.message || '操作失败')
      } finally {
        submitLoading.value = false
      }
    }
  })
}

// 上传图片到服务器
const uploadImage = (file) => {
  return new Promise((resolve, reject) => {
    const formData = new FormData()
    formData.append('image', file)
    
    // 获取token
    const token = localStorage.getItem('admin_token')
    
    fetch('/api/admin/upload', {
      method: 'POST',
      body: formData,
      headers: {
        'Authorization': `Bearer ${token}`
      }
    })
    .then(res => res.json())
    .then(data => {
      if (data.code === 200) {
        resolve(data.data.url)
      } else {
        reject(new Error(data.message || '上传失败'))
      }
    })
    .catch(err => {
      console.error('上传错误:', err)
      reject(new Error('网络错误，请检查后端服务'))
    })
  })
}

// 图片选择变化
const handleImageChange = (file, fileList) => {
  // 为新选择的文件生成本地预览URL
  if (file.raw && !file.url) {
    file.url = URL.createObjectURL(file.raw)
  }
  imageList.value = fileList
}

// 图片删除
const handleImageRemove = (file, fileList) => {
  imageList.value = fileList
}

const handleToggleStatus = async (row) => {
  const action = row.status === 1 ? '禁用' : '启用'
  ElMessageBox.confirm(`确定要${action}该场地吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      const newStatus = row.status === 1 ? 0 : 1
      await updateCourt(row.id, { status: newStatus })
      row.status = newStatus
      ElMessage.success(`${action}成功`)
      loadCourts()
    } catch (error) {
      console.error('更新状态失败:', error)
      ElMessage.error(error.message || '操作失败')
    }
  })
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该场地吗?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'error'
  }).then(async () => {
    try {
      await deleteCourt(row.id)
      ElMessage.success('删除成功')
      loadCourts()
    } catch (error) {
      console.error('删除失败:', error)
      ElMessage.error(error.message || '删除失败')
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

.image-upload-section {
  width: 100%;
}

.upload-tip {
  margin-top: 10px;
  font-size: 12px;
  color: #999;
}

:deep(.el-upload--picture-card) {
  width: 100px;
  height: 100px;
}

:deep(.el-upload-list--picture-card .el-upload-list__item) {
  width: 100px;
  height: 100px;
}

.table-image {
  position: relative;
  display: inline-block;
}

.image-count {
  position: absolute;
  bottom: 2px;
  right: 2px;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  font-size: 10px;
  padding: 2px 4px;
  border-radius: 2px;
}

.no-image {
  color: #999;
  font-size: 12px;
}
</style>
