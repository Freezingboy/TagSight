<template>
  <div class="upload-progress" :class="{ dark: isDarkTheme }">
    <div class="progress-info">
      <span>上传中: project_document.pdf</span>
      <span class="progress-percentage">45%</span>
    </div>
    <div class="progress-bar" :class="{ dark: isDarkTheme }">
      <div class="progress-fill" style="width: 45%"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 定义组件的属性接口
interface Props {
  // 当前主题模式
  isDarkTheme?: boolean
}

// 定义组件的属性，设置默认值
const props = withDefaults(defineProps<Props>(), {
  isDarkTheme: false
})
</script>

<style scoped>
.upload-progress {
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1.5rem;
  transition: all 0.3s ease;
}

/* 暗色主题下的上传进度条 */
.upload-progress.dark {
  background: linear-gradient(135deg, #1e293b 0%, #334155 100%) !important;
  border-color: #475569 !important;
}

.progress-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
  color: #1f2937;
}

/* 暗色主题下的进度信息文字 */
.upload-progress.dark .progress-info {
  color: #ffffff !important;
}

.progress-percentage {
  color: #3b82f6;
  font-weight: 600;
}

/* 暗色主题下的进度百分比文字 */
.upload-progress.dark .progress-percentage {
  color: #ffffff !important;
}

.progress-bar {
  height: 4px;
  background-color: #e2e8f0;
  border-radius: 2px;
  overflow: hidden;
  transition: background-color 0.3s ease;
}

/* 暗色主题下的进度条背景 */
.progress-bar.dark {
  background: linear-gradient(90deg, #475569 0%, #64748b 100%) !important;
}

.progress-fill {
  height: 100%;
  background-color: #3b82f6;
  transition: width 0.3s ease;
}
</style> 