<template>
  <el-dialog
    v-model="dialogVisible"
    :title="isEdit ? '编辑时间记录' : '创建时间记录'"
    width="600px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="rules"
      label-width="100px"
      label-position="left"
    >
      <el-form-item label="标题" prop="title">
        <el-input
          v-model="formData.title"
          placeholder="请输入标题"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="描述" prop="description">
        <el-input
          v-model="formData.description"
          type="textarea"
          :rows="3"
          placeholder="请输入描述（可选）"
          maxlength="500"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="目标日期" prop="target_date">
        <el-date-picker
          v-model="formData.target_date"
          type="datetime"
          placeholder="选择日期时间"
          format="YYYY-MM-DD HH:mm"
          value-format="YYYY-MM-DDTHH:mm:ss.000Z"
          style="width: 100%"
          @change="handleDateChange"
        />
      </el-form-item>

      <!-- <el-form-item label="记录类型">
        <el-tag :type="recordTypeTag.type" size="large">
          {{ recordTypeTag.text }}
        </el-tag>
        <span style="margin-left: 12px; color: #909399;">
          {{ recordTypeTag.hint }}
        </span>
      </el-form-item> -->

      <el-form-item label="分类" prop="category">
        <el-input
          v-model="formData.category"
          placeholder="请输入分类（可选）"
          maxlength="50"
        />
      </el-form-item>

      <el-form-item label="颜色" prop="color">
        <el-color-picker
          v-model="formData.color"
          show-alpha
          :predefine="predefineColors"
        />
        <span style="margin-left: 12px; color: #909399;">选择卡片主题颜色</span>
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="primary" @click="handleSubmit" :loading="loading">
        {{ isEdit ? '更新' : '创建' }}
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, watch, computed } from 'vue'
import dayjs from 'dayjs'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  record: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['update:visible', 'submit'])

const dialogVisible = ref(false)
const formRef = ref(null)
const loading = ref(false)

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
  '#f9ca24',
  '#6c5ce7',
  '#a29bfe',
  '#fd79a8',
]

// 表单数据
const formData = reactive({
  title: '',
  description: '',
  target_date: '',
  record_type: 'countdown',
  category: '',
  color: '#409EFF'
})

// 自动计算记录类型标签
// const recordTypeTag = computed(() => {
//   if (!formData.target_date) {
//     return {
//       type: 'info',
//       text: '请选择日期',
//       hint: '根据目标日期自动判断'
//     }
//   }
  
//   const targetDate = dayjs(formData.target_date)
//   const now = dayjs()
  
//   if (targetDate.isAfter(now)) {
//     return {
//       type: 'warning',
//       text: '⏰ 倒计时',
//       hint: '目标日期在未来'
//     }
//   } else {
//     return {
//       type: 'success',
//       text: '📅 累计天数',
//       hint: '目标日期已过去'
//     }
//   }
// })

// 处理日期变化 - 自动更新 record_type
const handleDateChange = (value) => {
  if (!value) {
    formData.record_type = 'countdown'
    return
  }
  
  const targetDate = dayjs(value)
  const now = dayjs()
  
  // 自动判断类型
  formData.record_type = targetDate.isAfter(now) ? 'countdown' : 'elapsed'
}

// 验证规则
const rules = {
  title: [
    { required: true, message: '请输入标题', trigger: 'blur' },
    { min: 1, max: 100, message: '长度在 1 到 100 个字符', trigger: 'blur' }
  ],
  target_date: [
    { required: true, message: '请选择目标日期', trigger: 'change' }
  ]
}

// 是否编辑模式
const isEdit = ref(false)

// 监听 visible 变化
watch(() => props.visible, (val) => {
  dialogVisible.value = val
  if (val) {
    if (props.record) {
      // 编辑模式
      isEdit.value = true
      Object.assign(formData, {
        title: props.record.title,
        description: props.record.description,
        target_date: props.record.target_date,
        record_type: props.record.record_type,
        category: props.record.category,
        color: props.record.color
      })
      // 编辑模式也重新计算类型（根据当前时间）
      handleDateChange(props.record.target_date)
    } else {
      // 创建模式
      isEdit.value = false
      resetForm()
    }
  }
})

// 重置表单
const resetForm = () => {
  Object.assign(formData, {
    title: '',
    description: '',
    target_date: '',
    record_type: 'countdown',
    category: '',
    color: '#409EFF'
  })
  formRef.value?.clearValidate()
}

// 关闭对话框
const handleClose = () => {
  emit('update:visible', false)
  resetForm()
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value.validate()
    loading.value = true
    
    const submitData = { ...formData }
    
    // 如果是编辑模式，传递 ID
    if (isEdit.value && props.record) {
      emit('submit', props.record.id, submitData)
    } else {
      emit('submit', submitData)
    }
    
    loading.value = false
    handleClose()
  } catch (error) {
    console.error('表单验证失败:', error)
    loading.value = false
  }
}
</script>

<style scoped>
:deep(.el-dialog) {
  border-radius: 12px;
}
</style>
