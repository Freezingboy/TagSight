<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import LoginPage from './components/LoginPage.vue'
import FileManager from './components/FileManager.vue'
import LoadingScreen from './components/LoadingScreen.vue'

// 登录状态管理
const isLoggedIn = ref(false)
const isLoading = ref(false)

// 检查登录状态
const checkLoginStatus = () => {
  const loginStatus = localStorage.getItem('isLoggedIn')
  if (loginStatus === 'true') {
    isLoggedIn.value = true
  }
}

// 处理登录成功
const handleLoginSuccess = () => {
  isLoading.value = true
  
  // 2秒后完成加载
  setTimeout(() => {
    isLoading.value = false
    isLoggedIn.value = true
  }, 2000)
}

// 处理登出
const handleLogout = () => {
  isLoggedIn.value = false
}

// 监听存储变化
const handleStorageChange = () => {
  checkLoginStatus()
}

// 监听登录成功事件
const handleLoginSuccessEvent = () => {
  handleLoginSuccess()
}

// 监听登出事件
const handleLogoutEvent = () => {
  handleLogout()
}

onMounted(() => {
  checkLoginStatus()
  window.addEventListener('storage', handleStorageChange)
  window.addEventListener('loginSuccess', handleLoginSuccessEvent)
  window.addEventListener('logout', handleLogoutEvent)
})

onUnmounted(() => {
  window.removeEventListener('storage', handleStorageChange)
  window.removeEventListener('loginSuccess', handleLoginSuccessEvent)
  window.removeEventListener('logout', handleLogoutEvent)
})
</script>

<template>
  <!-- 加载屏幕 -->
  <LoadingScreen v-if="isLoading" />
  
  <!-- 根据登录状态显示不同组件 -->
  <LoginPage v-else-if="!isLoggedIn" />
  <FileManager v-else @logout="handleLogout" />
</template>

<style>
/* 全局样式设置 */
/* 确保应用占满整个视口 */
html, body, #app {
  width: 100vw;  /* 视口宽度 */
  height: 100vh; /* 视口高度 */
  margin: 0;     /* 移除默认边距 */
  padding: 0;    /* 移除默认内边距 */
  background: white; /* 白色背景 */
}

/* 全局盒模型设置 */
* {
  box-sizing: border-box; /* 使用border-box盒模型，包含padding和border */
}
</style>
