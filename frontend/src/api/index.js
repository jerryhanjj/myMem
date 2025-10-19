import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 响应拦截器
api.interceptors.response.use(
  response => response.data,
  error => {
    console.error('API Error:', error)
    return Promise.reject(error)
  }
)

// 时间记录 API
export const timeRecordApi = {
  // 获取所有记录
  getAll() {
    return api.get('/records')
  },
  
  // 获取单个记录
  getById(id) {
    return api.get(`/records/${id}`)
  },
  
  // 创建记录
  create(data) {
    return api.post('/records', data)
  },
  
  // 更新记录
  update(id, data) {
    return api.put(`/records/${id}`, data)
  },
  
  // 删除记录
  delete(id) {
    return api.delete(`/records/${id}`)
  },

  // 批量删除记录
  batchDelete(ids) {
    return api.post('/records/batch-delete', ids)
  },

  // 清空所有记录
  clearAll() {
    return api.delete('/records/clear')
  }
}

// 统计 API
export const statisticsApi = {
  // 获取统计数据
  getStatistics() {
    return api.get('/statistics')
  }
}

export default api
