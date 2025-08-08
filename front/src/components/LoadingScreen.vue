<template>
  <div class="loading-screen">
    <!-- 背景渐变 -->
    <div class="loading-background"></div>
    
    <!-- 主要内容 -->
    <div class="loading-content">
      <!-- Logo动画 -->
      <div class="logo-container">
        <div class="logo-spinner">
          <svg class="logo-icon" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"/>
            <line x1="7" y1="7" x2="7.01" y2="7"/>
          </svg>
        </div>
        <div class="logo-text">TagSight</div>
      </div>
      
      <!-- 加载文字 -->
      <div class="loading-text">
        <h2>正在加载您的数据...</h2>
        <p>请稍候，我们正在为您准备最佳体验</p>
      </div>
      
      <!-- 进度条容器 -->
      <div class="progress-container">
        <div class="progress-bar">
          <div class="progress-fill" :style="{ width: progress + '%' }"></div>
        </div>
        <div class="progress-text">{{ Math.round(progress) }}%</div>
      </div>
      
      <!-- 功能图标动画 -->
      <div class="feature-animation">
        <div class="feature-item" v-for="(feature, index) in features" :key="index" :style="{ animationDelay: index * 0.2 + 's' }">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline v-if="feature.type === 'chart'" points="22,12 18,12 15,21 9,3 6,12 2,12"/>
            <rect v-if="feature.type === 'image'" x="3" y="3" width="18" height="18" rx="2" ry="2"/>
            <circle v-if="feature.type === 'image'" cx="8.5" cy="8.5" r="1.5"/>
            <polyline v-if="feature.type === 'image'" points="21,15 16,10 5,21"/>
            <path v-if="feature.type === 'cloud'" d="M18 10h-1.26A8 8 0 1 0 9 20h9a5 5 0 0 0 0-10z"/>
          </svg>
          <span>{{ feature.text }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// 进度状态
const progress = ref(0)

// 功能列表
const features = [
  { type: 'chart', text: '数据分析' },
  { type: 'image', text: '文件管理' },
  { type: 'cloud', text: '云端同步' }
]

// 进度动画
const startProgress = () => {
  const duration = 2000 // 2秒
  const interval = 50 // 每50ms更新一次
  const steps = duration / interval
  const increment = 100 / steps
  
  const timer = setInterval(() => {
    progress.value += increment
    if (progress.value >= 100) {
      progress.value = 100
      clearInterval(timer)
    }
  }, interval)
}

// 组件挂载时开始进度动画
onMounted(() => {
  startProgress()
})
</script>

<style scoped>
.loading-screen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

.loading-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  opacity: 0.95;
}

.loading-content {
  position: relative;
  z-index: 1;
  text-align: center;
  color: white;
  max-width: 500px;
  padding: 2rem;
}

/* Logo动画 */
.logo-container {
  margin-bottom: 2rem;
}

.logo-spinner {
  display: inline-block;
  margin-bottom: 1rem;
  animation: spin 2s linear infinite;
}

.logo-icon {
  color: white;
}

.logo-text {
  font-size: 2rem;
  font-weight: 700;
  letter-spacing: 1px;
}

/* 旋转动画 */
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

/* 加载文字 */
.loading-text {
  margin-bottom: 2rem;
}

.loading-text h2 {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  opacity: 0.9;
}

.loading-text p {
  font-size: 1rem;
  opacity: 0.7;
  margin: 0;
}

/* 进度条 */
.progress-container {
  margin-bottom: 2rem;
}

.progress-bar {
  width: 100%;
  height: 8px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 4px;
  overflow: hidden;
  margin-bottom: 0.5rem;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #ffffff 0%, #f0f0f0 100%);
  border-radius: 4px;
  transition: width 0.1s ease;
  position: relative;
}

.progress-fill::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent 0%, rgba(255, 255, 255, 0.3) 50%, transparent 100%);
  animation: shimmer 1.5s infinite;
}

@keyframes shimmer {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

.progress-text {
  font-size: 0.9rem;
  font-weight: 500;
  opacity: 0.8;
}

/* 功能动画 */
.feature-animation {
  display: flex;
  justify-content: center;
  gap: 2rem;
  flex-wrap: wrap;
}

.feature-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  opacity: 0;
  animation: fadeInUp 0.6s ease forwards;
}

.feature-item svg {
  color: white;
  opacity: 0.8;
}

.feature-item span {
  font-size: 0.8rem;
  opacity: 0.7;
  white-space: nowrap;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .loading-content {
    padding: 1rem;
  }
  
  .logo-text {
    font-size: 1.5rem;
  }
  
  .loading-text h2 {
    font-size: 1.25rem;
  }
  
  .feature-animation {
    gap: 1rem;
  }
  
  .feature-item span {
    font-size: 0.7rem;
  }
}
</style>
