<template>
  <div class="home-container">
    <div class="header">
      <div class="header-content">
        <h1 class="title">
          <el-icon class="title-icon">
            <Timer />
          </el-icon>
          MyMem 时间记录器
        </h1>
        <el-button type="primary" size="large" :icon="Plus" @click="openCreateDialog">
          创建记录
        </el-button>
      </div>
    </div>

    <div class="main-content">
      <!-- 筛选栏 -->
      <div class="filter-bar">
        <el-radio-group v-model="filterType" @change="handleFilterChange">
          <el-radio-button label="all">全部</el-radio-button>
          <el-radio-button label="countdown">倒计时</el-radio-button>
          <el-radio-button label="elapsed">累计天数</el-radio-button>
        </el-radio-group>

        <el-input v-model="searchKeyword" placeholder="搜索标题或分类..." :prefix-icon="Search" clearable
          style="width: 300px" />
      </div>

      <!-- 加载状态 -->
      <div v-if="store.loading" class="loading-container">
        <el-icon class="is-loading" :size="50">
          <Loading />
        </el-icon>
      </div>

      <!-- 空状态 -->
      <el-empty v-else-if="filteredRecords.length === 0" description="还没有时间记录，点击上方按钮创建吧" :image-size="200" />

      <!-- 卡片网格 -->
      <div v-else class="card-grid">
        <TimeCard v-for="record in filteredRecords" :key="record.id" :record="record" @edit="openEditDialog"
          @delete="handleDelete" @updateDisplayMode="handleUpdateDisplayMode" />
      </div>
    </div>

    <!-- 表单对话框 -->
    <TimeForm v-model:visible="dialogVisible" :record="currentRecord" @submit="handleSubmit" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessageBox } from 'element-plus'
import { Plus, Timer, Search, Loading } from '@element-plus/icons-vue'
import TimeCard from '@/components/TimeCard.vue'
import TimeForm from '@/components/TimeForm.vue'
import { useTimeRecordStore } from '@/stores/timeRecord'

const store = useTimeRecordStore()

// 对话框状态
const dialogVisible = ref(false)
const currentRecord = ref(null)

// 筛选状态
const filterType = ref('all')
const searchKeyword = ref('')

// 过滤后的记录
const filteredRecords = computed(() => {
  let records = store.records

  // 按类型筛选
  if (filterType.value !== 'all') {
    records = records.filter(r => r.record_type === filterType.value)
  }

  // 按关键词搜索
  if (searchKeyword.value) {
    const keyword = searchKeyword.value.toLowerCase()
    records = records.filter(r =>
      r.title.toLowerCase().includes(keyword) ||
      (r.category && r.category.toLowerCase().includes(keyword))
    )
  }

  return records
})

// 打开创建对话框
const openCreateDialog = () => {
  currentRecord.value = null
  dialogVisible.value = true
}

// 打开编辑对话框
const openEditDialog = (record) => {
  currentRecord.value = record
  dialogVisible.value = true
}

// 处理表单提交
const handleSubmit = async (idOrData, data) => {
  try {
    if (data) {
      // 编辑模式
      await store.updateRecord(idOrData, data)
    } else {
      // 创建模式
      await store.createRecord(idOrData)
    }
  } catch (error) {
    console.error('提交失败:', error)
  }
}

// 删除记录
const handleDelete = async (id) => {
  try {
    await ElMessageBox.confirm(
      '确定要删除这条记录吗？',
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
    await store.deleteRecord(id)
  } catch (error) {
    // 用户取消删除
  }
}

// 更新显示模式（仅本地状态，不保存到数据库）
const handleUpdateDisplayMode = async (id, displayMode) => {
  // 点击切换是纯前端行为，不需要任何操作
  // 每个 TimeCard 组件自己管理本地状态
  console.log(`Card ${id} switched to ${displayMode} mode`)
}

// 筛选变化
const handleFilterChange = () => {
  // 可以添加额外的逻辑
}

// 初始化
onMounted(() => {
  store.fetchRecords()
})
</script>

<style scoped>
.home-container {
  max-width: 1400px;
  margin: 0 auto;
}

.header {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 24px 32px;
  margin-bottom: 24px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title {
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

.main-content {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 24px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  min-height: 600px;
}

.filter-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  flex-wrap: wrap;
  gap: 16px;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
  color: #409EFF;
}

.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 24px;
}

@media (max-width: 768px) {
  .header {
    padding: 16px;
  }

  .header-content {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }

  .title {
    font-size: 24px;
    justify-content: center;
  }

  .filter-bar {
    flex-direction: column;
    align-items: stretch;
  }

  .filter-bar .el-input {
    width: 100% !important;
  }

  .card-grid {
    grid-template-columns: 1fr;
  }
}
</style>
