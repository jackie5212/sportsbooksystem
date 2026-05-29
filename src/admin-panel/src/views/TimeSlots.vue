<template>
  <div class="page-container">
    <h2 class="page-title">时间段管理</h2>
    
    <!-- 营业时间工具栏 -->
    <el-card class="business-hours-bar" shadow="never">
      <div class="business-hours-inner">
        <div class="bh-title">
          <el-icon><Clock /></el-icon>
          <span>营业时间</span>
        </div>
        <el-form :inline="true" :model="globalBusinessHours" class="business-hours-form">
          <el-form-item label="开始">
            <el-time-picker
              v-model="globalBusinessHours.start_time"
              format="HH:mm"
              value-format="HH:mm:ss"
              placeholder="开始"
              style="width: 110px"
              size="small"
            />
          </el-form-item>
          <el-form-item label="结束">
            <el-time-picker
              v-model="globalBusinessHours.end_time"
              format="HH:mm"
              value-format="HH:mm:ss"
              placeholder="结束"
              style="width: 110px"
              size="small"
            />
          </el-form-item>
          <el-form-item label="间隔">
            <el-select v-model="globalBusinessHours.interval" style="width: 90px" size="small">
              <el-option label="30分" :value="30" />
              <el-option label="1小时" :value="60" />
              <el-option label="2小时" :value="120" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" size="small" @click="applyGlobalBusinessHours">
              <el-icon><Check /></el-icon>
              应用到所有场地
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-card>

    <!-- 预订日历 -->
    <el-row class="calendar-row">
      <el-col :span="24">
        <el-card class="calendar-card" shadow="never">
          <template #header>
            <span class="calendar-header"><el-icon><Calendar /></el-icon> 预订日历</span>
          </template>
          <el-calendar v-model="calendarDate">
            <template #date-cell="{ data }">
              <div 
                class="calendar-day"
                :class="{ 'has-booking': isDateHasBooking(data.day) }"
                @click="selectCalendarDate(data.day)"
              >
                <div class="day-number">{{ data.day.split('-').slice(2).join('-') }}</div>
                <div v-if="isDateHasBooking(data.day)" class="booking-indicator">
                  <el-icon><User /></el-icon>
                  <span>{{ getBookingCount(data.day) }}</span>
                </div>
              </div>
            </template>
          </el-calendar>
          <div class="calendar-legend">
            <div class="legend-item">
              <div class="legend-color normal"></div>
              <span>无预订</span>
            </div>
            <div class="legend-item">
              <div class="legend-color booked"></div>
              <span>有预订</span>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    
    <!-- 按场地分块显示 -->
    <el-row :gutter="20">
      <el-col 
        v-for="court in courts" 
        :key="court.id" 
        :span="24"
        class="court-block"
      >
        <el-card>
          <template #header>
            <div class="court-header">
              <span class="court-name">{{ court.name }}</span>
              <div class="court-actions">
                <el-button type="primary" size="small" @click="showUnavailableDialog(court.id)">
                  <el-icon><Clock /></el-icon>
                  设置不营业时间
                </el-button>
              </div>
            </div>
          </template>
          
          <div v-loading="loadingStates[court.id]">
            <div class="time-slots-grid">
              <div 
                v-for="slot in getCourtSlots(court.id)" 
                :key="slot.id"
                class="time-slot-card"
                :class="{ 'available': slot.status === 1, 'unavailable': slot.status !== 1 }"
                @click="toggleSlotStatus(slot)"
              >
                <div class="slot-time">
                  {{ formatTime(slot.start_time) }} - {{ formatTime(slot.end_time) }}
                </div>
                <div class="slot-status">
                  <el-tag :type="getStatusType(slot.status)" size="small">
                    {{ getStatusText(slot.status) }}
                  </el-tag>
                </div>
              </div>
            </div>
          </div>
          
          <el-divider />
          
          <div class="batch-operations">
            <el-button size="small" type="success" @click="batchSetCourtStatus(court.id, 1)">
              全部可预订
            </el-button>
            <el-button size="small" type="danger" @click="batchSetCourtStatus(court.id, 3)">
              全部不可预订
            </el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 设置不营业时间对话框 -->
    <el-dialog 
      v-model="unavailableDialogVisible" 
      title="设置不营业时间" 
      width="600px"
    >
      <el-form :model="unavailableForm" label-width="120px">
        <el-form-item label="选择日期类型">
          <el-radio-group v-model="unavailableForm.dateType">
            <el-radio label="specific">指定日期</el-radio>
            <el-radio label="weekly">每周重复</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item v-if="unavailableForm.dateType === 'specific'" label="选择日期">
          <el-date-picker
            v-model="unavailableForm.specificDate"
            type="date"
            placeholder="选择日期"
            style="width: 100%"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        
        <el-form-item v-if="unavailableForm.dateType === 'weekly'" label="选择星期">
          <el-checkbox-group v-model="unavailableForm.weekDays">
            <el-checkbox :label="1">周一</el-checkbox>
            <el-checkbox :label="2">周二</el-checkbox>
            <el-checkbox :label="3">周三</el-checkbox>
            <el-checkbox :label="4">周四</el-checkbox>
            <el-checkbox :label="5">周五</el-checkbox>
            <el-checkbox :label="6">周六</el-checkbox>
            <el-checkbox :label="7">周日</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        
        <el-form-item label="时间范围">
          <el-time-picker
            v-model="unavailableForm.startTime"
            format="HH:mm"
            value-format="HH:mm:ss"
            placeholder="开始时间"
            style="width: 45%; margin-right: 10%"
          />
          <span style="margin: 0 10px">至</span>
          <el-time-picker
            v-model="unavailableForm.endTime"
            format="HH:mm"
            value-format="HH:mm:ss"
            placeholder="结束时间"
            style="width: 45%"
          />
        </el-form-item>
        
        <el-form-item label="说明">
          <el-input 
            v-model="unavailableForm.reason" 
            type="textarea" 
            :rows="2"
            placeholder="例如：维护、休息等（可选）"
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <el-button @click="unavailableDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveUnavailableHours">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getCourts } from '@/api/admin'

const loading = ref(false)
const courts = ref([])
const selectedDate = ref(new Date())
const allTimeSlots = ref([])
const loadingStates = ref({})

// 统一营业时间设置
const globalBusinessHours = ref({
  start_time: '08:00:00',
  end_time: '22:00:00',
  interval: 60
})

// 日历相关
const calendarDate = ref(new Date())
const bookingData = ref({})

// 不营业时间对话框
const unavailableDialogVisible = ref(false)
const currentCourtId = ref(null)
const unavailableForm = ref({
  dateType: 'specific', // specific: 指定日期, weekly: 每周重复
  specificDate: '',
  weekDays: [],
  startTime: '12:00:00',
  endTime: '14:00:00',
  reason: ''
})

// 应用全局营业时间到所有场地
const applyGlobalBusinessHours = () => {
  if (!globalBusinessHours.value.start_time || !globalBusinessHours.value.end_time) {
    ElMessage.warning('请设置完整的营业时间')
    return
  }
  
  const [startHour] = globalBusinessHours.value.start_time.split(':').map(Number)
  const [endHour] = globalBusinessHours.value.end_time.split(':').map(Number)
  
  if (startHour >= endHour) {
    ElMessage.warning('结束时间必须大于开始时间')
    return
  }
  
  // 重新生成所有场地的时间段
  regenerateAllSlots()
  
  ElMessage.success(`已将所有场地营业时间设置为 ${formatTime(globalBusinessHours.value.start_time)}-${formatTime(globalBusinessHours.value.end_time)}`)
}

// 重新生成所有时间段
const regenerateAllSlots = () => {
  generateDefaultSlots()
}

// 检查日期是否有预订
const isDateHasBooking = (dateStr) => {
  // 模拟数据：某些日期有预订
  const datesWithBookings = ['2026-05-28', '2026-05-29', '2026-05-30', '2026-06-01', '2026-06-05']
  return datesWithBookings.includes(dateStr)
}

// 获取日期的预订数量
const getBookingCount = (dateStr) => {
  // 模拟数据
  const counts = {
    '2026-05-28': 3,
    '2026-05-29': 5,
    '2026-05-30': 2,
    '2026-06-01': 8,
    '2026-06-05': 4
  }
  return counts[dateStr] || 0
}

// 选择日历日期
const selectCalendarDate = (dateStr) => {
  selectedDate.value = new Date(dateStr)
  regenerateAllSlots()
  ElMessage.info(`已切换到 ${dateStr}`)
}

// 加载场地列表
const loadCourts = async () => {
  try {
    const response = await getCourts({ page: 1, page_size: 100 })
    courts.value = response.list || []
    
    // 初始化加载状态
    courts.value.forEach(court => {
      loadingStates.value[court.id] = false
    })
    
    // 为每个场地生成默认时间段
    generateDefaultSlots()
  } catch (error) {
    console.error('加载场地失败:', error)
  }
}

// 生成默认时间段（根据全局营业时间）
const generateDefaultSlots = () => {
  const slots = []
  const baseDate = new Date(selectedDate.value)
  const dateStr = `${baseDate.getFullYear()}-${String(baseDate.getMonth() + 1).padStart(2, '0')}-${String(baseDate.getDate()).padStart(2, '0')}`
  
  const [startHour, startMin] = globalBusinessHours.value.start_time.split(':').map(Number)
  const [endHour, endMin] = globalBusinessHours.value.end_time.split(':').map(Number)
  const intervalMinutes = globalBusinessHours.value.interval
  
  courts.value.forEach(court => {
    let currentMinutes = startHour * 60 + startMin
    const endTotalMinutes = endHour * 60 + endMin
    
    while (currentMinutes + intervalMinutes <= endTotalMinutes) {
      const slotStartHour = Math.floor(currentMinutes / 60)
      const slotStartMin = currentMinutes % 60
      const slotEndMinutes = currentMinutes + intervalMinutes
      const slotEndHour = Math.floor(slotEndMinutes / 60)
      const slotEndMin = slotEndMinutes % 60
      
      slots.push({
        id: `${court.id}_${currentMinutes}`,
        court_id: court.id,
        date: dateStr,
        start_time: `${String(slotStartHour).padStart(2, '0')}:${String(slotStartMin).padStart(2, '0')}:00`,
        end_time: `${String(slotEndHour).padStart(2, '0')}:${String(slotEndMin).padStart(2, '0')}:00`,
        status: 1 // 默认可预订
      })
      
      currentMinutes += intervalMinutes
    }
  })
  
  allTimeSlots.value = slots
}

// 获取指定场地的时间段
const getCourtSlots = (courtId) => {
  return allTimeSlots.value.filter(slot => slot.court_id === courtId)
}

// 格式化时间显示
const formatTime = (time) => {
  if (!time) return ''
  return time.substring(0, 5) // HH:mm
}

// 获取状态类型
const getStatusType = (status) => {
  const types = { 1: 'success', 2: 'warning', 3: 'info' }
  return types[status] || ''
}

// 获取状态文本
const getStatusText = (status) => {
  const texts = { 1: '可预订', 2: '已预订', 3: '不可预订' }
  return texts[status] || '未知'
}

// 切换时间段状态
const toggleSlotStatus = (slot) => {
  if (slot.status === 2) {
    ElMessage.warning('该时间段已被预订，无法修改')
    return
  }
  
  slot.status = slot.status === 1 ? 3 : 1
  ElMessage.success(`已${slot.status === 1 ? '设为可预订' : '设为不可预订'}`)
  
  // TODO: 调用API保存
  // updateTimeSlotStatus(slot.id, slot.status)
}

// 显示不营业时间对话框
const showUnavailableDialog = (courtId) => {
  currentCourtId.value = courtId
  unavailableForm.value = {
    dateType: 'specific',
    specificDate: '',
    weekDays: [],
    startTime: '12:00:00',
    endTime: '14:00:00',
    reason: ''
  }
  unavailableDialogVisible.value = true
}

// 保存不营业时间
const saveUnavailableHours = () => {
  if (!currentCourtId.value) {
    ElMessage.warning('请选择场地')
    return
  }
  
  if (unavailableForm.value.dateType === 'specific' && !unavailableForm.value.specificDate) {
    ElMessage.warning('请选择日期')
    return
  }
  
  if (unavailableForm.value.dateType === 'weekly' && unavailableForm.value.weekDays.length === 0) {
    ElMessage.warning('请选择星期')
    return
  }
  
  // 解析时间
  const [startHour] = unavailableForm.value.startTime.split(':').map(Number)
  const [endHour] = unavailableForm.value.endTime.split(':').map(Number)
  
  if (startHour >= endHour) {
    ElMessage.warning('结束时间必须大于开始时间')
    return
  }
  
  // 根据设置将对应时间段标记为不可预订
  let affectedCount = 0
  
  if (unavailableForm.value.dateType === 'specific') {
    // 指定日期
    const targetDate = unavailableForm.value.specificDate
    
    allTimeSlots.value.forEach(slot => {
      if (slot.court_id === currentCourtId.value && slot.date === targetDate) {
        const slotHour = parseInt(slot.start_time.split(':')[0])
        if (slotHour >= startHour && slotHour < endHour && slot.status !== 2) {
          slot.status = 3
          affectedCount++
        }
      }
    })
    
    ElMessage.success(`已将 ${affectedCount} 个时间段设为不可预订`)
  } else {
    // 每周重复
    const targetWeekDays = unavailableForm.value.weekDays
    
    // 这里简化处理，实际应该根据日期计算星期几
    // 暂时将所有符合时间范围的都设为不可预订
    allTimeSlots.value.forEach(slot => {
      if (slot.court_id === currentCourtId.value) {
        const slotHour = parseInt(slot.start_time.split(':')[0])
        if (slotHour >= startHour && slotHour < endHour && slot.status !== 2) {
          slot.status = 3
          affectedCount++
        }
      }
    })
    
    const weekDayText = targetWeekDays.map(d => `周${['日','一','二','三','四','五','六'][d]}`).join('、')
    ElMessage.success(`已将每周${weekDayText}的 ${affectedCount} 个时间段设为不可预订`)
  }
  
  unavailableDialogVisible.value = false
}

// 批量设置场地状态
const batchSetCourtStatus = (courtId, status) => {
  const text = status === 1 ? '可预订' : '不可预订'
  
  allTimeSlots.value.forEach(slot => {
    if (slot.court_id === courtId && slot.status !== 2) {
      slot.status = status
    }
  })
  
  ElMessage.success(`已将该场地所有时间段设为${text}`)
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

.business-hours-bar {
  margin-bottom: 16px;
}

.business-hours-bar :deep(.el-card__body) {
  padding: 12px 16px;
}

.business-hours-inner {
  display: flex;
  align-items: center;
  gap: 16px;
  flex-wrap: wrap;
}

.bh-title {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  font-weight: 600;
  color: #303133;
  white-space: nowrap;
}

.business-hours-form {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 4px;
}

.business-hours-form :deep(.el-form-item) {
  margin-bottom: 0;
  margin-right: 8px;
}

.business-hours-form :deep(.el-form-item__label) {
  font-size: 12px;
  padding-right: 4px;
}

.calendar-row {
  margin-bottom: 16px;
}

.calendar-card {
  flex: 1;
}

.calendar-card :deep(.el-card__header) {
  padding: 10px 16px;
  border-bottom: 1px solid #ebeef5;
}

.calendar-header {
  font-size: 14px;
  font-weight: 600;
  color: #303133;
}

.calendar-card :deep(.el-calendar) {
  height: auto;
}

.calendar-card :deep(.el-calendar__header) {
  padding: 8px 12px;
}

.calendar-card :deep(.el-calendar__title) {
  font-size: 14px;
}

.calendar-card :deep(.el-calendar__body) {
  padding: 8px 12px;
}

.calendar-card :deep(.el-calendar-table thead th) {
  padding: 4px 0;
  font-size: 12px;
}

.calendar-card :deep(.el-calendar-table .el-calendar-day) {
  height: auto;
  padding: 2px;
}

.calendar-day {
  height: 36px;
  padding: 3px;
  cursor: pointer;
  transition: all 0.3s;
  border-radius: 4px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
}

.calendar-day:hover {
  background-color: #f5f7fa;
}

.calendar-day.has-booking {
  background-color: #fef0f0;
  border: 1px solid #f56c6c;
}

.calendar-day.has-booking:hover {
  background-color: #fde2e2;
}

.day-number {
  font-size: 12px;
  color: #606266;
  line-height: 1.2;
}

.booking-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 2px;
  font-size: 10px;
  color: #f56c6c;
  font-weight: bold;
  line-height: 1.2;
}

.calendar-legend {
  margin-top: 8px;
  padding-top: 8px;
  border-top: 1px solid #ebeef5;
  display: flex;
  gap: 16px;
  justify-content: center;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #606266;
}

.legend-color {
  width: 12px;
  height: 12px;
  border-radius: 2px;
}

.legend-color.normal {
  background-color: #fff;
  border: 1px solid #dcdfe6;
}

.legend-color.booked {
  background-color: #fef0f0;
  border: 1px solid #f56c6c;
}

.court-block {
  margin-bottom: 20px;
}

.court-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.court-name {
  font-size: 18px;
  font-weight: bold;
  color: #303133;
}

.time-slots-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 12px;
  max-height: 400px;
  overflow-y: auto;
}

.time-slot-card {
  padding: 15px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  border: 2px solid transparent;
  text-align: center;
}

.time-slot-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.time-slot-card.available {
  background-color: #f0f9ff;
  border-color: #67c23a;
}

.time-slot-card.unavailable {
  background-color: #fef0f0;
  border-color: #f56c6c;
}

.slot-time {
  font-size: 16px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 8px;
}

.slot-status {
  margin-top: 8px;
}

.batch-operations {
  display: flex;
  gap: 10px;
  justify-content: center;
}
</style>
