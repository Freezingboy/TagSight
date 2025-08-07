<template>
  <div class="pagination" :class="{ dark: isDarkTheme }">
    <div class="pagination-info">
      显示1-6条, 共128条
    </div>
    <div class="pagination-controls">
      <button class="page-btn">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <polyline points="15,18 9,12 15,6"/>
        </svg>
      </button>
      <button class="page-btn active">1</button>
      <button class="page-btn">2</button>
      <button class="page-btn">3</button>
      <span class="page-dots">...</span>
      <button class="page-btn">21</button>
      <button class="page-btn">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <polyline points="9,18 15,12 9,6"/>
        </svg>
      </button>
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
.pagination {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 1px solid #e2e8f0;
  transition: border-color 0.3s ease;
}

/* 暗色主题下的分页组件 */
.pagination.dark {
  border-top: 1px solid #475569;
}

.pagination-info {
  color: #6b7280;
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

/* 暗色主题下的分页信息 */
.pagination.dark .pagination-info {
  color: #ffffff !important;
}

.pagination-controls {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

.page-btn {
  background-color: #f1f5f9;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  padding: 0.5rem 0.75rem;
  cursor: pointer;
  color: #6b7280;
  transition: all 0.2s ease;
  min-width: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 暗色主题下的分页按钮 */
.pagination.dark .page-btn {
  background: linear-gradient(135deg, #1e293b 0%, #334155 100%) !important;
  border-color: #475569 !important;
  color: #ffffff !important;
}

.page-btn.active {
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  color: white;
  border-color: #6366f1;
}

.page-btn:hover {
  background-color: #e2e8f0;
  border-color: #cbd5e1;
}

/* 暗色主题下的分页按钮悬停效果 */
.pagination.dark .page-btn:hover {
  background: linear-gradient(135deg, #334155 0%, #475569 100%) !important;
  border-color: #6366f1 !important;
}

.page-dots {
  color: #6b7280;
  padding: 0 0.25rem;
  transition: color 0.3s ease;
}

/* 暗色主题下的分页省略号 */
.pagination.dark .page-dots {
  color: #ffffff !important;
}
</style> 