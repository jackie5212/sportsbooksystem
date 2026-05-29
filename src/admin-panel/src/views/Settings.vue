<template>
  <div class="settings-container">
    <el-card class="settings-card">
      <template #header>
        <div class="card-header">
          <span>系统设置</span>
        </div>
      </template>

      <el-tabs v-model="activeTab" type="border-card">
        <!-- 微信支付设置 -->
        <el-tab-pane label="微信支付配置" name="wechat">
          <el-form
            ref="wechatFormRef"
            :model="wechatForm"
            label-width="180px"
            class="settings-form"
          >
            <el-divider content-position="left">网页端微信支付（公众号/JSAPI）</el-divider>

            <el-form-item label="公众号 AppID">
              <el-input
                v-model="wechatForm.wechat_web_appid"
                placeholder="请输入公众号/服务号 AppID"
                clearable
              />
            </el-form-item>

            <el-form-item label="公众号 AppSecret">
              <el-input
                v-model="wechatForm.wechat_web_appsecret"
                placeholder="请输入公众号 AppSecret"
                clearable
                show-password
              />
            </el-form-item>

            <el-divider content-position="left">小程序微信支付</el-divider>

            <el-form-item label="小程序 AppID">
              <el-input
                v-model="wechatForm.wechat_mini_appid"
                placeholder="请输入小程序 AppID"
                clearable
              />
            </el-form-item>

            <el-form-item label="小程序 AppSecret">
              <el-input
                v-model="wechatForm.wechat_mini_appsecret"
                placeholder="请输入小程序 AppSecret"
                clearable
                show-password
              />
            </el-form-item>

            <el-divider content-position="left">通用商户配置</el-divider>

            <el-form-item label="微信支付商户号">
              <el-input
                v-model="wechatForm.wechat_mch_id"
                placeholder="请输入微信支付商户号（MchID）"
                clearable
              />
            </el-form-item>

            <el-form-item label="APIv3 密钥">
              <el-input
                v-model="wechatForm.wechat_api_v3_key"
                placeholder="请输入微信支付 APIv3 密钥"
                clearable
                show-password
              />
            </el-form-item>

            <el-form-item label="支付回调地址">
              <el-input
                v-model="wechatForm.wechat_notify_url"
                placeholder="https://your-domain.com/api/payments/notify"
                clearable
              />
              <div class="form-tip">
                回调地址需为外网可访问的 HTTPS 地址，用于接收微信支付结果通知
              </div>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" :loading="saving" @click="saveWechatConfig">
                保存配置
              </el-button>
              <el-button @click="resetWechatForm">重置</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <!-- 通用设置 -->
        <el-tab-pane label="通用设置" name="general">
          <el-form
            ref="generalFormRef"
            :model="generalForm"
            label-width="180px"
            class="settings-form"
          >
            <el-form-item label="订单超时时间(分钟)">
              <el-input-number
                v-model="generalForm.booking_timeout_minutes"
                :min="5"
                :max="120"
                controls-position="right"
              />
            </el-form-item>

            <el-form-item label="最多可预定天数">
              <el-input-number
                v-model="generalForm.max_booking_days"
                :min="1"
                :max="30"
                controls-position="right"
              />
            </el-form-item>

            <el-form-item label="营业开始时间">
              <el-time-picker
                v-model="generalForm.business_start_time"
                format="HH:mm"
                value-format="HH:mm"
                placeholder="选择开始时间"
              />
            </el-form-item>

            <el-form-item label="营业结束时间">
              <el-time-picker
                v-model="generalForm.business_end_time"
                format="HH:mm"
                value-format="HH:mm"
                placeholder="选择结束时间"
              />
            </el-form-item>

            <el-form-item label="时间段时长(分钟)">
              <el-input-number
                v-model="generalForm.slot_duration_minutes"
                :min="30"
                :max="180"
                :step="30"
                controls-position="right"
              />
            </el-form-item>

            <el-form-item>
              <el-button type="primary" :loading="saving" @click="saveGeneralConfig">
                保存配置
              </el-button>
              <el-button @click="resetGeneralForm">重置</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getSettings, updateSettings } from '@/api/admin'

const activeTab = ref('wechat')
const saving = ref(false)
const allSettings = ref({})

// 微信支付表单
const wechatForm = ref({
  wechat_web_appid: '',
  wechat_web_appsecret: '',
  wechat_mini_appid: '',
  wechat_mini_appsecret: '',
  wechat_mch_id: '',
  wechat_api_v3_key: '',
  wechat_notify_url: ''
})

// 通用设置表单
const generalForm = ref({
  booking_timeout_minutes: 30,
  max_booking_days: 7,
  business_start_time: '08:00',
  business_end_time: '22:00',
  slot_duration_minutes: 60
})

// 加载配置
const loadSettings = async () => {
  try {
    const data = await getSettings()
    allSettings.value = data || {}

    // 填充微信支付表单
    wechatForm.value.wechat_web_appid = data.wechat_web_appid || ''
    wechatForm.value.wechat_web_appsecret = data.wechat_web_appsecret || ''
    wechatForm.value.wechat_mini_appid = data.wechat_mini_appid || ''
    wechatForm.value.wechat_mini_appsecret = data.wechat_mini_appsecret || ''
    wechatForm.value.wechat_mch_id = data.wechat_mch_id || ''
    wechatForm.value.wechat_api_v3_key = data.wechat_api_v3_key || ''
    wechatForm.value.wechat_notify_url = data.wechat_notify_url || ''

    // 填充通用设置表单
    generalForm.value.booking_timeout_minutes = parseInt(data.booking_timeout_minutes) || 30
    generalForm.value.max_booking_days = parseInt(data.max_booking_days) || 7
    generalForm.value.business_start_time = data.business_start_time || '08:00'
    generalForm.value.business_end_time = data.business_end_time || '22:00'
    generalForm.value.slot_duration_minutes = parseInt(data.slot_duration_minutes) || 60
  } catch (error) {
    ElMessage.error('加载配置失败')
  }
}

// 保存微信支付配置
const saveWechatConfig = async () => {
  saving.value = true
  try {
    const payload = { ...wechatForm.value }
    await updateSettings(payload)
    ElMessage.success('微信支付配置保存成功')
    loadSettings()
  } catch (error) {
    ElMessage.error('保存失败')
  } finally {
    saving.value = false
  }
}

// 保存通用配置
const saveGeneralConfig = async () => {
  saving.value = true
  try {
    const payload = {
      booking_timeout_minutes: String(generalForm.value.booking_timeout_minutes),
      max_booking_days: String(generalForm.value.max_booking_days),
      business_start_time: generalForm.value.business_start_time,
      business_end_time: generalForm.value.business_end_time,
      slot_duration_minutes: String(generalForm.value.slot_duration_minutes)
    }
    await updateSettings(payload)
    ElMessage.success('通用配置保存成功')
    loadSettings()
  } catch (error) {
    ElMessage.error('保存失败')
  } finally {
    saving.value = false
  }
}

// 重置表单
const resetWechatForm = () => {
  wechatForm.value = {
    wechat_web_appid: allSettings.value.wechat_web_appid || '',
    wechat_web_appsecret: allSettings.value.wechat_web_appsecret || '',
    wechat_mini_appid: allSettings.value.wechat_mini_appid || '',
    wechat_mini_appsecret: allSettings.value.wechat_mini_appsecret || '',
    wechat_mch_id: allSettings.value.wechat_mch_id || '',
    wechat_api_v3_key: allSettings.value.wechat_api_v3_key || '',
    wechat_notify_url: allSettings.value.wechat_notify_url || ''
  }
}

const resetGeneralForm = () => {
  generalForm.value = {
    booking_timeout_minutes: parseInt(allSettings.value.booking_timeout_minutes) || 30,
    max_booking_days: parseInt(allSettings.value.max_booking_days) || 7,
    business_start_time: allSettings.value.business_start_time || '08:00',
    business_end_time: allSettings.value.business_end_time || '22:00',
    slot_duration_minutes: parseInt(allSettings.value.slot_duration_minutes) || 60
  }
}

onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
.settings-container {
  padding: 20px;
}

.settings-card {
  max-width: 900px;
}

.card-header {
  font-size: 18px;
  font-weight: bold;
}

.settings-form {
  padding: 20px 10px;
}

.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 5px;
  line-height: 1.4;
}

:deep(.el-divider__text) {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
}
</style>
