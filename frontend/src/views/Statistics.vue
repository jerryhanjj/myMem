<template>
  <div class="statistics-container">
    <div class="page-header">
      <h1 class="page-title">
        <el-icon class="title-icon"><DataAnalysis /></el-icon>
        数据统计
      </h1>
    </div>

    <div class="stats-grid">
      <!-- 总览卡片 -->
      <el-card class="stat-card" shadow="hover">
        <div class="stat-content">
          <div class="stat-icon" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%)">
            <el-icon :size="32"><Collection /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-label">总记录数</div>
            <div class="stat-value">{{ totalRecords }}</div>
          </div>
        </div>
      </el-card>

      <!-- 倒计时卡片 -->
      <el-card class="stat-card" shadow="hover">
        <div class="stat-content">
          <div class="stat-icon" style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%)">
            <el-icon :size="32"><Timer /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-label">倒计时记录</div>
            <div class="stat-value">{{ countdownRecords }}</div>
          </div>
        </div>
      </el-card>

      <!-- 累计天数卡片 -->
      <el-card class="stat-card" shadow="hover">
        <div class="stat-content">
          <div class="stat-icon" style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)">
            <el-icon :size="32"><Calendar /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-label">累计天数记录</div>
            <div class="stat-value">{{ elapsedRecords }}</div>
          </div>
        </div>
      </el-card>

      <!-- 分类数量卡片 -->
      <el-card class="stat-card" shadow="hover">
        <div class="stat-content">
          <div class="stat-icon" style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%)">
            <el-icon :size="32"><Folder /></el-icon>
          </div>
          <div class="stat-info">
            <div class="stat-label">分类数量</div>
            <div class="stat-value">{{ categoryCount }}</div>
          </div>
        </div>
      </el-card>
    </div>

    <!-- 分类统计表格 -->
    <el-card class="chart-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <span class="card-title">分类统计</span>
        </div>
      </template>
      <el-table :data="categoryStats" style="width: 100%">
        <el-table-column prop="category" label="分类" width="200">
          <template #default="{ row }">
            <el-tag v-if="row.category" type="info">{{ row.category }}</el-tag>
            <span v-else style="color: #909399;">未分类</span>
          </template>
        </el-table-column>
        <el-table-column prop="count" label="记录数量" width="150">
          <template #default="{ row }">
            <el-text type="primary" size="large">{{ row.count }}</el-text>
          </template>
        </el-table-column>
        <el-table-column label="占比">
          <template #default="{ row }">
            <div class="progress-bar">
              <el-progress 
                :percentage="row.percentage" 
                :color="row.color"
                :show-text="true"
              />
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 最近创建 -->
    <el-card class="chart-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <span class="card-title">最近创建的记录</span>
        </div>
      </template>
      <el-timeline>
        <el-timeline-item
          v-for="record in recentRecords"
          :key="record.id"
          :timestamp="formatDate(record.created_at)"
          placement="top"
          :color="record.color"
        >
          <el-card>
            <h4>{{ record.title }}</h4>
            <p v-if="record.description" style="margin: 8px 0; color: #606266;">
              {{ record.description }}
            </p>
            <div style="display: flex; gap: 8px; margin-top: 8px;">
              <el-tag :type="record.record_type === 'countdown' ? 'warning' : 'success'" size="small">
                {{ record.record_type === 'countdown' ? '倒计时' : '累计天数' }}
              </el-tag>
              <el-tag v-if="record.category" size="small" effect="plain">
                {{ record.category }}
              </el-tag>
            </div>
          </el-card>
        </el-timeline-item>
      </el-timeline>
      <el-empty v-if="recentRecords.length === 0" description="暂无记录" />
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { DataAnalysis, Collection, Timer, Calendar, Folder } from '@element-plus/icons-vue'
import { useTimeRecordStore } from '@/stores/timeRecord'
import { statisticsApi } from '@/api'
import dayjs from 'dayjs'

const store = useTimeRecordStore()
const useBackendStats = ref(false) // 是否使用后端统计（可切换）
const backendStats = ref(null)

// 总记录数
const totalRecords = computed(() => {
  if (useBackendStats.value && backendStats.value) {
    return backendStats.value.total_records
  }
  return store.records.length
})

// 倒计时记录数
const countdownRecords = computed(() => {
  if (useBackendStats.value && backendStats.value) {
    return backendStats.value.countdown_records
  }
  return store.records.filter(r => r.record_type === 'countdown').length
})

// 累计天数记录数
const elapsedRecords = computed(() => {
  if (useBackendStats.value && backendStats.value) {
    return backendStats.value.elapsed_records
  }
  return store.records.filter(r => r.record_type === 'elapsed').length
})

// 分类数量
const categoryCount = computed(() => {
  const categories = new Set(
    store.records
      .filter(r => r.category)
      .map(r => r.category)
  )
  return categories.size
})

// 分类统计
const categoryStats = computed(() => {
  const stats = {}
  
  store.records.forEach(record => {
    const category = record.category || '未分类'
    if (!stats[category]) {
      stats[category] = {
        category: record.category,
        count: 0,
        color: record.color || '#409EFF'
      }
    }
    stats[category].count++
  })

  const total = store.records.length || 1
  return Object.values(stats).map(stat => ({
    ...stat,
    percentage: Math.round((stat.count / total) * 100)
  })).sort((a, b) => b.count - a.count)
})

// 最近创建的记录（前5条）
const recentRecords = computed(() => {
  return [...store.records]
    .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
    .slice(0, 5)
})

// 格式化日期
const formatDate = (date) => {
  return dayjs(date).format('YYYY-MM-DD HH:mm:ss')
}

// 初始化
onMounted(() => {
  if (store.records.length === 0) {
    store.fetchRecords()
  }
})
</script>

<style scoped>
.statistics-container {
  max-width: 1400px;
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

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 24px;
}

.stat-card {
  border-radius: 12px;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-5px);
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 20px;
}

.stat-icon {
  width: 64px;
  height: 64px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  flex-shrink: 0;
}

.stat-info {
  flex: 1;
}

.stat-label {
  font-size: 14px;
  color: #909399;
  margin-bottom: 8px;
}

.stat-value {
  font-size: 32px;
  font-weight: bold;
  color: #303133;
}

.chart-card {
  border-radius: 12px;
  margin-bottom: 24px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-title {
  font-size: 18px;
  font-weight: 600;
  color: #303133;
}

.progress-bar {
  padding-right: 20px;
}

:deep(.el-timeline-item__timestamp) {
  color: #909399;
  font-size: 13px;
}

@media (max-width: 768px) {
  .statistics-container {
    padding: 12px;
  }

  .page-header {
    padding: 16px;
  }

  .page-title {
    font-size: 24px;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }
}
</style>
