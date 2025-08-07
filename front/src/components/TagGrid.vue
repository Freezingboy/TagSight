<template>
  <div class="tags-grid" :class="{ dark: isDarkTheme }">
    <div class="tag-card">
      <div class="tag-color" style="background-color: #ef4444;"></div>
      <div class="tag-content">
        <h3>重要项目</h3>
        <p class="tag-count">15个文件</p>
      </div>
      <div class="tag-actions">
        <button class="action-btn view-btn">查看</button>
        <button class="action-btn edit-btn">修改</button>
        <button class="action-btn delete-btn">删除</button>
      </div>
    </div>
    
    <div class="tag-card">
      <div class="tag-color" style="background-color: #3b82f6;"></div>
      <div class="tag-content">
        <h3>工作文档</h3>
        <p class="tag-count">42个文件</p>
      </div>
      <div class="tag-actions">
        <button class="action-btn view-btn">查看</button>
        <button class="action-btn edit-btn">修改</button>
        <button class="action-btn delete-btn">删除</button>
      </div>
    </div>
    
    <div class="tag-card">
      <div class="tag-color" style="background-color: #10b981;"></div>
      <div class="tag-content">
        <h3>个人资料</h3>
        <p class="tag-count">8个文件</p>
      </div>
      <div class="tag-actions">
        <button class="action-btn view-btn">查看</button>
        <button class="action-btn edit-btn">修改</button>
        <button class="action-btn delete-btn">删除</button>
      </div>
    </div>
    
    <div class="tag-card">
      <div class="tag-color" style="background-color: #8b5cf6;"></div>
      <div class="tag-content">
        <h3>参考资料</h3>
        <p class="tag-count">23个文件</p>
      </div>
      <div class="tag-actions">
        <button class="action-btn view-btn">查看</button>
        <button class="action-btn edit-btn">修改</button>
        <button class="action-btn delete-btn">删除</button>
      </div>
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
.tags-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 1rem;
}

.tag-card {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 1rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  cursor: pointer;
  transition: all 0.2s ease;
  position: relative;
  overflow: visible;
  min-width: 0;
}

/* 暗色主题下的标签卡片 */
.tags-grid.dark .tag-card {
  background-color: #2d2d2d;
  border-color: #404040;
}

.tag-card:hover {
  background-color: #f8fafc;
  border-color: #cbd5e1;
}

/* 暗色主题下的标签卡片悬停效果 */
.tags-grid.dark .tag-card:hover {
  background-color: #404040;
  border-color: #525252;
}

.tag-card:hover .tag-actions {
  opacity: 1;
  transform: translateX(0);
}

.tag-color {
  width: 24px;
  height: 24px;
  border-radius: 4px;
  flex-shrink: 0;
}

.tag-content {
  flex: 1;
  min-width: 0;
  overflow: hidden;
}

.tag-content h3 {
  margin: 0 0 0.25rem 0;
  font-size: 0.875rem;
  font-weight: 500;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.tag-count {
  margin: 0;
  font-size: 0.75rem;
  color: #6b7280;
  transition: color 0.3s ease;
}

/* 暗色主题下的标签数量 */
.tags-grid.dark .tag-count {
  color: #9ca3af;
}

.tag-actions {
  display: flex;
  gap: 0.25rem;
  opacity: 0;
  transform: translateX(10px);
  transition: all 0.2s ease;
  flex-shrink: 0;
  min-width: fit-content;
  position: relative;
  z-index: 2;
}

.action-btn {
  padding: 0.25rem 0.5rem;
  border: none;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
  background-color: transparent;
  color: inherit;
  position: relative;
  z-index: 1;
  font-size: 0.75rem;
  font-weight: 500;
  min-width: 40px;
}

.action-btn:hover {
  transform: scale(1.05);
}

.action-btn.view-btn {
  color: #3b82f6;
}

.action-btn.view-btn:hover {
  background-color: rgba(59, 130, 246, 0.1);
}

.action-btn.edit-btn {
  color: #f59e0b;
}

.action-btn.edit-btn:hover {
  background-color: rgba(245, 158, 11, 0.1);
}

.action-btn.delete-btn {
  color: #ef4444;
}

.action-btn.delete-btn:hover {
  background-color: rgba(239, 68, 68, 0.1);
}

/* 暗色主题下的操作按钮 */
.tags-grid.dark .action-btn {
  color: #9ca3af;
}

.tags-grid.dark .action-btn:hover {
  background-color: rgba(59, 130, 246, 0.2);
}

.tags-grid.dark .action-btn.view-btn {
  color: #60a5fa;
}

.tags-grid.dark .action-btn.view-btn:hover {
  background-color: rgba(96, 165, 250, 0.2);
}

.tags-grid.dark .action-btn.edit-btn {
  color: #fbbf24;
}

.tags-grid.dark .action-btn.edit-btn:hover {
  background-color: rgba(251, 191, 36, 0.2);
}

.tags-grid.dark .action-btn.delete-btn {
  color: #f87171;
}

.tags-grid.dark .action-btn.delete-btn:hover {
  background-color: rgba(248, 113, 113, 0.2);
}
</style> 