<template>
  <div class="tags-grid" :class="{ dark: isDarkTheme }">
    <div 
      v-for="tag in tags" 
      :key="tag.id" 
      class="tag-card"
      @click="handleTagClick(tag.id)"
    >
      <div class="tag-color" :style="{ backgroundColor: tag.color }"></div>
      <div class="tag-content">
        <h3>{{ tag.name }}</h3>
        <p class="tag-count">{{ tag.fileCount }}个文件</p>
      </div>
      <div class="tag-actions">
        <button class="action-btn view-btn" @click.stop="handleTagClick(tag.id)">查看</button>
        <button class="action-btn edit-btn">修改</button>
        <button class="action-btn delete-btn">删除</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// 定义标签接口
interface TagItem {
  id: string
  name: string
  color: string
  fileCount: number
}

// 定义组件的属性接口
interface Props {
  // 当前主题模式
  isDarkTheme?: boolean
}

// 定义组件的事件接口
interface Emits {
  (e: 'tag-click', tagId: string): void
}

// 定义组件的属性，设置默认值
const props = withDefaults(defineProps<Props>(), {
  isDarkTheme: false
})

const emit = defineEmits<Emits>()

// 标签数据
const tags = ref<TagItem[]>([
  { id: 'important', name: '重要项目', color: '#ef4444', fileCount: 15 },
  { id: 'work', name: '工作文档', color: '#3b82f6', fileCount: 42 },
  { id: 'personal', name: '个人资料', color: '#10b981', fileCount: 8 },
  { id: 'reference', name: '参考资料', color: '#8b5cf6', fileCount: 23 }
])

// 处理标签点击
const handleTagClick = (tagId: string) => {
  console.log('标签被点击:', tagId)
  emit('tag-click', tagId)
}
</script>

<style scoped>
.tags-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 1rem;
  padding: 1rem;
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
  background: linear-gradient(135deg, #1e293b 0%, #334155 100%) !important;
  border-color: #475569 !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
}

.tag-card:hover {
  background-color: #f8fafc;
  border-color: #cbd5e1;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

/* 暗色主题下的标签卡片悬停效果 */
.tags-grid.dark .tag-card:hover {
  background: linear-gradient(135deg, #334155 0%, #475569 100%) !important;
  border-color: #6366f1 !important;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.5) !important;
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
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

/* 暗色主题下的标签颜色 */
.tags-grid.dark .tag-color {
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
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
  color: #1f2937;
}

/* 暗色主题下的标签标题 */
.tags-grid.dark .tag-content h3 {
  color: #ffffff !important;
}

.tag-count {
  margin: 0;
  font-size: 0.75rem;
  color: #6b7280;
  transition: color 0.3s ease;
}

/* 暗色主题下的标签数量 */
.tags-grid.dark .tag-count {
  color: #ffffff !important;
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
  color: #ffffff !important;
}

.tags-grid.dark .action-btn:hover {
  background-color: rgba(79, 70, 229, 0.2);
}

.tags-grid.dark .action-btn.view-btn {
  color: #ffffff !important;
}

.tags-grid.dark .action-btn.view-btn:hover {
  background-color: rgba(99, 102, 241, 0.2);
}

.tags-grid.dark .action-btn.edit-btn {
  color: #ffffff !important;
}

.tags-grid.dark .action-btn.edit-btn:hover {
  background-color: rgba(245, 158, 11, 0.2);
}

.tags-grid.dark .action-btn.delete-btn {
  color: #ffffff !important;
}

.tags-grid.dark .action-btn.delete-btn:hover {
  background-color: rgba(239, 68, 68, 0.2);
}
</style> 