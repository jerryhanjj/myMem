import { defineStore } from 'pinia'
import { ref } from 'vue'
import { timeRecordApi } from '@/api'
import { ElMessage } from 'element-plus'

export const useTimeRecordStore = defineStore('timeRecord', () => {
  const records = ref([])
  const loading = ref(false)

  // 获取所有记录
  const fetchRecords = async () => {
    loading.value = true
    try {
      const response = await timeRecordApi.getAll()
      records.value = response.data || []
    } catch (error) {
      ElMessage.error('获取记录失败')
      console.error(error)
    } finally {
      loading.value = false
    }
  }

  // 创建记录
  const createRecord = async (data) => {
    try {
      const response = await timeRecordApi.create(data)
      records.value.unshift(response.data)
      ElMessage.success('创建成功')
      return response.data
    } catch (error) {
      ElMessage.error('创建失败')
      throw error
    }
  }

  // 更新记录
  const updateRecord = async (id, data) => {
    try {
      const response = await timeRecordApi.update(id, data)
      const index = records.value.findIndex(r => r.id === id)
      if (index !== -1) {
        // 使用 Object.assign 确保响应式更新
        Object.assign(records.value[index], response.data)
      }
      ElMessage.success('更新成功')
      return response.data
    } catch (error) {
      ElMessage.error('更新失败')
      console.error('Update error:', error)
      throw error
    }
  }

  // 删除记录
  const deleteRecord = async (id) => {
    try {
      await timeRecordApi.delete(id)
      records.value = records.value.filter(r => r.id !== id)
      ElMessage.success('删除成功')
    } catch (error) {
      ElMessage.error('删除失败')
      throw error
    }
  }

  return {
    records,
    loading,
    fetchRecords,
    createRecord,
    updateRecord,
    deleteRecord
  }
})
