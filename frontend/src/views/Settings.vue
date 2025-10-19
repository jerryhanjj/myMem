<template>
  <div class="settings-container">
    <div class="page-header">
      <h1 class="page-title">
        <el-icon class="title-icon"><Setting /></el-icon>
        系统设置
      </h1>
    </div>

    <div class="settings-content">
      <!-- 外观设置 -->
      <el-card class="setting-card" shadow="hover">
        <template #header>
          <div class="card-header">
            <el-icon><Brush /></el-icon>
            <span>外观设置</span>
          </div>
        </template>

        <el-form label-width="120px" label-position="left">
          <el-form-item label="主题模式">
            <el-radio-group v-model="settings.theme">
              <el-radio value="light">浅色</el-radio>
              <el-radio value="dark">深色</el-radio>
              <el-radio value="auto">跟随系统</el-radio>
            </el-radio-group>
          </el-form-item>

          <el-form-item label="主题颜色">
            <el-color-picker 
              v-model="settings.primaryColor" 
              show-alpha 
              :predefine="predefineColors"
            />
            <span style="margin-left: 12px; color: #909399;">
              当前: {{ settings.primaryColor }}
            </span>
          </el-form-item>
        </el-form>
      </el-card>

      <!-- 通知设置 -->
      <el-card class="setting-card" shadow="hover">
        <template #header>
          <div class="card-header">
            <el-icon><Bell /></el-icon>
            <span>通知设置</span>
          </div>
        </template>

        <el-form label-width="120px" label-position="left">
          <el-form-item label="桌面通知">
            <el-switch v-model="settings.enableNotifications" />
            <div class="form-tip">倒计时到期时发送桌面通知</div>
          </el-form-item>

          <el-form-item label="提醒音效">
            <el-switch v-model="settings.enableSound" />
            <div class="form-tip">通知时播放提示音</div>
          </el-form-item>

          <el-form-item label="提前提醒">
            <el-input-number 
              v-model="settings.reminderDays" 
              :min="0" 
              :max="30"
              controls-position="right"
            />
            <span style="margin-left: 8px; color: #909399;">天前提醒</span>
          </el-form-item>
        </el-form>
      </el-card>

      <!-- 数据管理 -->
      <el-card class="setting-card" shadow="hover">
        <template #header>
          <div class="card-header">
            <el-icon><FolderOpened /></el-icon>
            <span>数据管理</span>
          </div>
        </template>

        <el-form label-width="120px" label-position="left">
          <el-form-item label="数据导出">
            <el-button type="primary" :icon="Download" @click="handleExportData">
              导出为 JSON
            </el-button>
            <div class="form-tip">导出所有时间记录数据</div>
          </el-form-item>

          <el-form-item label="清空数据">
            <el-popconfirm
              title="确定要清空所有数据吗？此操作不可恢复！"
              confirm-button-text="确定"
              cancel-button-text="取消"
              @confirm="handleClearData"
            >
              <template #reference>
                <el-button type="danger" :icon="Delete">
                  清空所有记录
                </el-button>
              </template>
            </el-popconfirm>
            <div class="form-tip warning">危险操作：将删除所有时间记录</div>
          </el-form-item>
        </el-form>
      </el-card>

      <!-- 关于信息 -->
      <el-card class="setting-card" shadow="hover">
        <template #header>
          <div class="card-header">
            <el-icon><InfoFilled /></el-icon>
            <span>关于</span>
          </div>
        </template>

        <div class="about-content">
          <div class="about-item">
            <span class="about-label">应用名称</span>
            <span class="about-value">MyMem 时间记录器</span>
          </div>
          <div class="about-item">
            <span class="about-label">版本号</span>
            <span class="about-value">v1.0.0</span>
          </div>
          <div class="about-item">
            <span class="about-label">技术栈</span>
            <span class="about-value">Vue 3 + Element Plus + Go + SQLite</span>
          </div>
          <div class="about-item">
            <span class="about-label">开发日期</span>
            <span class="about-value">2025年10月18日</span>
          </div>
        </div>
      </el-card>

      <!-- 保存按钮 -->
      <div class="save-button-container">
        <el-button type="primary" size="large" @click="handleSaveSettings" :loading="saving">
          <el-icon><Select /></el-icon>
          保存设置
        </el-button>
        <el-button size="large" @click="handleResetSettings">
          <el-icon><RefreshLeft /></el-icon>
          恢复默认
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { 
  Setting, Brush, Bell, FolderOpened, InfoFilled,
  Download, Delete, Select, RefreshLeft
} from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useTimeRecordStore } from '@/stores/timeRecord'
import { timeRecordApi } from '@/api'

const store = useTimeRecordStore()
const saving = ref(false)

// 预定义颜色
const predefineColors = [
  '#409EFF',
  '#67C23A',
  '#E6A23C',
  '#F56C6C',
  '#909399',
  '#ff6b6b',
  '#4ecdc4',
  '#45b7d1',
]

// 设置数据
const settings = reactive({
  theme: 'light',
  primaryColor: '#409EFF',
  enableNotifications: false,
  enableSound: true,
  reminderDays: 1
})

// 加载设置
const loadSettings = () => {
  const savedSettings = localStorage.getItem('myMemSettings')
  if (savedSettings) {
    Object.assign(settings, JSON.parse(savedSettings))
  }
}

// 保存设置
const handleSaveSettings = () => {
  saving.value = true
  
  setTimeout(() => {
    localStorage.setItem('myMemSettings', JSON.stringify(settings))
    saving.value = false
    ElMessage.success('设置已保存')
  }, 500)
}

// 恢复默认设置
const handleResetSettings = () => {
  ElMessageBox.confirm(
    '确定要恢复默认设置吗？',
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    settings.theme = 'light'
    settings.primaryColor = '#409EFF'
    settings.enableNotifications = false
    settings.enableSound = true
    settings.reminderDays = 1
    
    localStorage.removeItem('myMemSettings')
    ElMessage.success('已恢复默认设置')
  }).catch(() => {})
}

// 导出数据
const handleExportData = () => {
  const data = JSON.stringify(store.records, null, 2)
  const blob = new Blob([data], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `mymem-export-${new Date().getTime()}.json`
  a.click()
  URL.revokeObjectURL(url)
  ElMessage.success('数据导出成功')
}

// 清空数据
const handleClearData = async () => {
  try {
    // 调用后端批量删除API
    await timeRecordApi.clearAll()
    // 清空前端store
    store.records = []
    ElMessage.success('数据已清空')
  } catch (error) {
    console.error('清空失败:', error)
    ElMessage.error('清空失败')
  }
}

// 初始化
onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
.settings-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

.page-header {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 24px 32px;
  margin-bottom: 24px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.page-title {
  margin: 0;
  font-size: 32px;
  font-weight: bold;
  color: #303133;
  display: flex;
  align-items: center;
  gap: 12px;
}

.title-icon {
  font-size: 36px;
  color: #409EFF;
}

.settings-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.setting-card {
  border-radius: 12px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
}

.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

.form-tip.warning {
  color: #F56C6C;
}

.about-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.about-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 0;
  border-bottom: 1px solid #EBEEF5;
}

.about-item:last-child {
  border-bottom: none;
}

.about-label {
  color: #909399;
  font-size: 14px;
}

.about-value {
  color: #303133;
  font-size: 14px;
  font-weight: 500;
}

.save-button-container {
  display: flex;
  justify-content: center;
  gap: 16px;
  padding: 24px 0;
}

@media (max-width: 768px) {
  .settings-container {
    padding: 12px;
  }

  .page-header {
    padding: 16px;
  }

  .page-title {
    font-size: 24px;
  }
}
</style>
