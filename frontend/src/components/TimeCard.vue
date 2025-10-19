<template>
  <el-card class="time-card" :body-style="{ padding: '20px' }" shadow="hover">
    <template #header>
      <div class="card-header">
        <div class="title-section">
          <div 
            class="color-indicator" 
            :style="{ backgroundColor: record.color }"
          ></div>
          <h3 class="title">{{ record.title }}</h3>
        </div>
        <div class="actions">
          <el-button 
            :icon="Edit" 
            circle 
            size="small" 
            @click="$emit('edit', record)"
          />
          <el-button 
            :icon="Delete" 
            circle 
            size="small" 
            type="danger"
            @click="$emit('delete', record.id)"
          />
        </div>
      </div>
    </template>

    <div class="card-body">
      <div class="time-display" @click="cycleDisplayMode" :title="displayModeHint">
        <div class="time-content">
          <span class="time-number" :style="{ color: record.color }">
            {{ timeValue }}
          </span>
          <span class="time-unit">{{ timeUnit }}</span>
        </div>
        <div class="display-mode-indicator">
          <el-icon><Refresh /></el-icon>
        </div>
      </div>

      <div class="date-info">
        <el-icon><Calendar /></el-icon>
        <span>{{ formattedDate }}</span>
      </div>

      <div v-if="record.description" class="description">
        <el-icon><Document /></el-icon>
        <span>{{ record.description }}</span>
      </div>

      <div class="meta-info">
        <el-tag v-if="record.category" size="small" effect="plain">
          {{ record.category }}
        </el-tag>
        <el-tag 
          :type="record.record_type === 'countdown' ? 'warning' : 'success'" 
          size="small"
        >
          {{ record.record_type === 'countdown' ? '倒计时' : '累计天数' }}
        </el-tag>
      </div>
    </div>
  </el-card>
</template>

<script setup>
import { computed, ref, onMounted, onUnmounted } from 'vue'
import { Edit, Delete, Calendar, Document, Refresh } from '@element-plus/icons-vue'
import dayjs from 'dayjs'

const props = defineProps({
  record: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['edit', 'delete', 'updateDisplayMode'])

// 本地显示模式状态（用于点击切换，不保存到数据库）
const localDisplayMode = ref('yearMonthDay')

const now = ref(dayjs())
let timer = null

// 定时更新当前时间
onMounted(() => {
  timer = setInterval(() => {
    now.value = dayjs()
  }, 1000)
})

onUnmounted(() => {
  if (timer) {
    clearInterval(timer)
  }
})

// 计算时间差
const timeDiff = computed(() => {
  const target = dayjs(props.record.target_date)
  const current = now.value

  if (props.record.record_type === 'countdown') {
    // 倒计时：目标日期 - 当前时间
    return target.diff(current)
  } else {
    // 累计：当前时间 - 目标日期
    return current.diff(target)
  }
})

// 计算显示的时间值
const timeValue = computed(() => {
  const diff = Math.abs(timeDiff.value)
  const totalDays = Math.floor(diff / (1000 * 60 * 60 * 24))
  const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60))
  const seconds = Math.floor((diff % (1000 * 60)) / 1000)

  // 如果不到1天，显示时分秒（模式3）
  if (totalDays === 0) {
    if (hours > 0) {
      return `${hours}:${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`
    } else {
      return `${minutes}:${String(seconds).padStart(2, '0')}`
    }
  }

  const mode = localDisplayMode.value

  switch (mode) {
    case 'yearMonthDay':
      // 模式1: 年月天智能显示
      return formatYearMonthDay(totalDays, hours)
    
    case 'totalDays':
      // 模式2: 只显示总天数
      return totalDays
    
    default:
      // 默认使用模式1
      return formatYearMonthDay(totalDays, hours)
  }
})

// 格式化年月天（智能显示）
const formatYearMonthDay = (totalDays, hours) => {
  const years = Math.floor(totalDays / 365)
  const remainingDays1 = totalDays % 365
  const months = Math.floor(remainingDays1 / 30)
  const days = remainingDays1 % 30

  if (years > 0) {
    // 大于1年：显示年月天
    if (months > 0 && days > 0) {
      return `${years}年 ${months}月 ${days}天`
    } else if (months > 0) {
      return `${years}年 ${months}月`
    } else if (days > 0) {
      return `${years}年 ${days}天`
    } else {
      return `${years}年`
    }
  } else if (months > 0) {
    // 不够1年但大于1月：显示月天
    if (days > 0) {
      return `${months}月 ${days}天`
    } else {
      return `${months}月`
    }
  } else {
    // 不够1月：只显示天数（不带单位，统一用 timeUnit 显示）
    return totalDays
  }
}

// 时间单位
const timeUnit = computed(() => {
  const diff = Math.abs(timeDiff.value)
  const totalDays = Math.floor(diff / (1000 * 60 * 60 * 24))
  
  // 不到1天，不显示单位
  if (totalDays === 0) {
    return ''
  }
  
  const mode = localDisplayMode.value
  
  // 模式1：年月天
  if (mode === 'yearMonthDay') {
    const years = Math.floor(totalDays / 365)
    const remainingDays = totalDays % 365
    const months = Math.floor(remainingDays / 30)
    
    // 只有不到1个月时才显示"天"单位
    if (years === 0 && months === 0) {
      return '天'
    }
    return ''
  }
  
  // 模式2：总天数，始终显示"天"
  if (mode === 'totalDays') {
    return '天'
  }
  
  return ''
})

// 显示模式提示
const displayModeHint = computed(() => {
  const mode = localDisplayMode.value
  switch (mode) {
    case 'yearMonthDay':
      return '点击切换为总天数模式'
    case 'totalDays':
      return '点击切换为年月天模式'
    default:
      return '点击切换显示模式'
  }
})

// 点击切换显示模式
const cycleDisplayMode = () => {
  // 切换模式
  if (localDisplayMode.value === 'yearMonthDay') {
    localDisplayMode.value = 'totalDays'
  } else {
    localDisplayMode.value = 'yearMonthDay'
  }
  
  // 触发更新事件（可选：保存到后端）
  emit('updateDisplayMode', props.record.id, localDisplayMode.value)
}

// 格式化目标日期
const formattedDate = computed(() => {
  return dayjs(props.record.target_date).format('YYYY年MM月DD日 HH:mm')
})
</script>

<style scoped>
.time-card {
  width: 100%;
  transition: all 0.3s ease;
  border-radius: 12px;
  overflow: hidden;
}

.time-card:hover {
  transform: translateY(-5px);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title-section {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.color-indicator {
  width: 4px;
  height: 32px;
  border-radius: 2px;
}

.title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
}

.actions {
  display: flex;
  gap: 8px;
}

.card-body {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.time-display {
  position: relative;
  text-align: center;
  padding: 24px 0;
  background: linear-gradient(135deg, #f5f7fa 0%, #ffffff 100%);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  user-select: none;
}

.time-display:hover {
  background: linear-gradient(135deg, #e8f0fe 0%, #f5f7fa 100%);
  transform: scale(1.02);
}

.time-display:active {
  transform: scale(0.98);
}

.display-mode-indicator {
  position: absolute;
  top: 8px;
  right: 8px;
  color: #909399;
  font-size: 16px;
  opacity: 0.5;
  transition: all 0.3s ease;
}

.time-display:hover .display-mode-indicator {
  opacity: 1;
  color: #409EFF;
}

.time-content {
  display: flex;
  align-items: baseline;
  justify-content: center;
  gap: 4px;
}

.time-number {
  font-size: 48px;
  font-weight: bold;
  line-height: 1;
}

.time-unit {
  font-size: 20px;
  color: #909399;
  font-weight: 500;
}

.date-info,
.description {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #606266;
  font-size: 14px;
}

.description {
  color: #909399;
  line-height: 1.6;
}

.meta-info {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  padding-top: 8px;
  border-top: 1px solid #ebeef5;
}
</style>
