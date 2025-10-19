<template>
  <div class="layout-container">
    <!-- 侧边导航栏 -->
    <el-aside :width="isCollapse ? '64px' : '200px'" class="sidebar">
      <div class="logo-container">
        <el-icon v-if="!isCollapse" class="logo-icon"><Timer /></el-icon>
        <h2 v-if="!isCollapse" class="logo-text">MyMem</h2>
      </div>

      <el-menu
        :default-active="activeRoute"
        class="sidebar-menu"
        :collapse="isCollapse"
        :collapse-transition="false"
        router
      >
        <el-menu-item index="/">
          <el-icon><HomeFilled /></el-icon>
          <template #title>首页</template>
        </el-menu-item>

        <el-menu-item index="/statistics">
          <el-icon><DataAnalysis /></el-icon>
          <template #title>统计</template>
        </el-menu-item>

        <el-menu-item index="/settings">
          <el-icon><Setting /></el-icon>
          <template #title>设置</template>
        </el-menu-item>
      </el-menu>

      <!-- 折叠按钮 -->
      <div class="collapse-button" @click="toggleCollapse">
        <el-icon>
          <component :is="isCollapse ? 'Expand' : 'Fold'" />
        </el-icon>
      </div>
    </el-aside>

    <!-- 主内容区域 -->
    <el-main class="main-content">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </el-main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { 
  Timer, HomeFilled, DataAnalysis, Setting, Expand, Fold 
} from '@element-plus/icons-vue'

const route = useRoute()
const isCollapse = ref(false)

// 当前激活的路由
const activeRoute = computed(() => route.path)

// 切换侧边栏折叠状态
const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}
</script>

<style scoped>
.layout-container {
  display: flex;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.sidebar {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  transition: width 0.3s ease;
  display: flex;
  flex-direction: column;
  position: relative;
}

.logo-container {
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  padding: 0 20px;
  border-bottom: 1px solid #ebeef5;
}

.logo-icon {
  font-size: 32px;
  color: #409EFF;
}

.logo-text {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin: 0;
}

.sidebar-menu {
  flex: 1;
  border-right: none;
  background: transparent;
}

.sidebar-menu :deep(.el-menu-item) {
  height: 56px;
  line-height: 56px;
  margin: 4px 8px;
  border-radius: 8px;
}

.sidebar-menu :deep(.el-menu-item.is-active) {
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.sidebar-menu :deep(.el-menu-item.is-active .el-icon) {
  color: white;
}

.collapse-button {
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-top: 1px solid #ebeef5;
  color: #909399;
  transition: all 0.3s;
}

.collapse-button:hover {
  background: #f5f7fa;
  color: #409EFF;
}

.main-content {
  flex: 1;
  padding: 0;
  overflow-y: auto;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@media (max-width: 768px) {
  .sidebar {
    position: fixed;
    left: 0;
    top: 0;
    bottom: 0;
    z-index: 1000;
  }

  .main-content {
    margin-left: 64px;
  }
}
</style>
