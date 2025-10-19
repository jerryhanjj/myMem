import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Statistics from '@/views/Statistics.vue'
import Settings from '@/views/Settings.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: '首页',
      icon: 'HomeFilled'
    }
  },
  {
    path: '/statistics',
    name: 'Statistics',
    component: Statistics,
    meta: {
      title: '统计',
      icon: 'DataAnalysis'
    }
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings,
    meta: {
      title: '设置',
      icon: 'Setting'
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 设置页面标题
router.beforeEach((to, from, next) => {
  document.title = `${to.meta.title} - MyMem 时间记录器` || 'MyMem 时间记录器'
  next()
})

export default router
