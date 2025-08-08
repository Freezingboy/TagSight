<template>
  <div class="tag-detail-page" :class="{ dark: isDarkTheme }">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-left">
        <button class="back-btn" @click="goBack">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="15,18 9,12 15,6"/>
          </svg>
          返回标签管理
        </button>
        <div class="tag-info">
          <div class="tag-color" :style="{ backgroundColor: tagInfo.color }"></div>
          <div class="tag-details">
            <h2>{{ tagInfo.name }}</h2>
            <p class="file-count">{{ tagInfo.fileCount }}个文件</p>
          </div>
        </div>
      </div>
      
      <div class="header-actions">
        <button class="action-btn edit-btn">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
          </svg>
          编辑标签
        </button>
        <button class="action-btn delete-btn">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="3,6 5,6 21,6"/>
            <path d="M19,6v14a2,2,0,0,1-2,2H7a2,2,0,0,1-2-2V6m3,0V4a2,2,0,0,1,2-2h4a2,2,0,0,1,2,2V6"/>
          </svg>
          删除标签
        </button>
      </div>
    </div>

    <!-- 文件筛选和排序 -->
    <div class="file-controls">
      <div class="filter-section">
        <div class="filter-group">
          <label>文件类型:</label>
          <select v-model="selectedFileType" class="filter-select">
            <option value="all">全部类型</option>
            <option value="images">图片</option>
            <option value="documents">文档</option>
            <option value="audio">音频</option>
            <option value="video">视频</option>
            <option value="archives">压缩包</option>
          </select>
        </div>
        
        <div class="filter-group">
          <label>排序方式:</label>
          <select v-model="sortBy" class="filter-select">
            <option value="name">按名称</option>
            <option value="size">按大小</option>
            <option value="date">按日期</option>
          </select>
        </div>
      </div>
      
      <div class="view-options">
        <button class="view-btn active">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="3" width="7" height="7"/>
            <rect x="14" y="3" width="7" height="7"/>
            <rect x="14" y="14" width="7" height="7"/>
            <rect x="3" y="14" width="7" height="7"/>
          </svg>
        </button>
        <button class="view-btn">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="8" y1="6" x2="21" y2="6"/>
            <line x1="8" y1="12" x2="21" y2="12"/>
            <line x1="8" y1="18" x2="21" y2="18"/>
            <line x1="3" y1="6" x2="3.01" y2="6"/>
            <line x1="3" y1="12" x2="3.01" y2="12"/>
            <line x1="3" y1="18" x2="3.01" y2="18"/>
          </svg>
        </button>
      </div>
    </div>

    <!-- 文件网格 -->
    <div class="files-section">
      <FileGrid 
        :isDarkTheme="isDarkTheme" 
        :files="filteredFiles"
        @file-click="handleFileClick"
      />
    </div>

    <!-- 分页 -->
    <Pagination :isDarkTheme="isDarkTheme" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import FileGrid from './FileGrid.vue'
import Pagination from './Pagination.vue'

// 定义组件的属性接口
interface Props {
  isDarkTheme?: boolean
  tagId?: string
}

// 定义组件的事件接口
interface Emits {
  (e: 'back'): void
  (e: 'file-click', fileName: string): void
}

const props = withDefaults(defineProps<Props>(), {
  isDarkTheme: false,
  tagId: ''
})

const emit = defineEmits<Emits>()

// 标签信息
const tagInfo = ref({
  name: '重要项目',
  color: '#ef4444',
  fileCount: 15
})

// 筛选和排序状态
const selectedFileType = ref('all')
const sortBy = ref('name')

// 模拟文件数据
const files = ref([
  {
    id: 1,
    name: 'project_report.pdf',
    type: 'documents',
    size: '5.7 MB',
    date: '2024-03-10',
    thumbnail: 'pdf'
  },
  {
    id: 2,
    name: 'presentation.mp4',
    type: 'video',
    size: '156 MB',
    date: '2024-02-28',
    thumbnail: 'video'
  },
  {
    id: 3,
    name: 'financial_data.xlsx',
    type: 'documents',
    size: '1.8 MB',
    date: '2024-02-15',
    thumbnail: 'excel'
  },
  {
    id: 4,
    name: 'meeting_notes.docx',
    type: 'documents',
    size: '2.1 MB',
    date: '2024-03-12',
    thumbnail: 'document'
  },
  {
    id: 5,
    name: 'background_music.mp3',
    type: 'audio',
    size: '8.2 MB',
    date: '2024-03-05',
    thumbnail: 'audio'
  },
  {
    id: 6,
    name: 'screenshot.png',
    type: 'images',
    size: '1.2 MB',
    date: '2024-03-14',
    thumbnail: 'image'
  }
])

// 根据标签ID加载数据
const loadTagData = (tagId: string) => {
  // 这里可以根据tagId从API加载数据
  // 现在使用模拟数据
  const tagData = {
    'important': { name: '重要项目', color: '#ef4444', fileCount: 15 },
    'work': { name: '工作文档', color: '#3b82f6', fileCount: 42 },
    'personal': { name: '个人资料', color: '#10b981', fileCount: 8 },
    'reference': { name: '参考资料', color: '#8b5cf6', fileCount: 23 }
  }
  
  if (tagData[tagId as keyof typeof tagData]) {
    tagInfo.value = tagData[tagId as keyof typeof tagData]
  }
}

// 筛选后的文件
const filteredFiles = computed(() => {
  let result = [...files.value]
  
  // 按文件类型筛选
  if (selectedFileType.value !== 'all') {
    result = result.filter(file => file.type === selectedFileType.value)
  }
  
  // 排序
  result.sort((a, b) => {
    switch (sortBy.value) {
      case 'name':
        return a.name.localeCompare(b.name)
      case 'size':
        return parseFloat(a.size) - parseFloat(b.size)
      case 'date':
        return new Date(b.date).getTime() - new Date(a.date).getTime()
      default:
        return 0
    }
  })
  
  return result
})

// 返回标签管理页面
const goBack = () => {
  emit('back')
}

// 处理文件点击
const handleFileClick = (fileName: string) => {
  emit('file-click', fileName)
}

// 组件挂载时加载数据
onMounted(() => {
  if (props.tagId) {
    loadTagData(props.tagId)
  }
})
</script>

<style scoped>
.tag-detail-page {
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* 页面头部 */
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
  border-bottom: 1px solid #e2e8f0;
  background: white;
}

.page-header.dark {
  background: #1e293b;
  border-color: #475569;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: #f1f5f9;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  color: #64748b;
  cursor: pointer;
  transition: all 0.2s ease;
}

.back-btn:hover {
  background: #e2e8f0;
  color: #475569;
}

.tag-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.tag-color {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  flex-shrink: 0;
}

.tag-details h2 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
  color: #1f2937;
}

.tag-details.dark h2 {
  color: #e5e7eb;
}

.file-count {
  margin: 0.25rem 0 0 0;
  font-size: 0.875rem;
  color: #6b7280;
}

.header-actions {
  display: flex;
  gap: 0.5rem;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  background: white;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.875rem;
}

.edit-btn {
  color: #3b82f6;
  border-color: #3b82f6;
}

.edit-btn:hover {
  background: #3b82f6;
  color: white;
}

.delete-btn {
  color: #ef4444;
  border-color: #ef4444;
}

.delete-btn:hover {
  background: #ef4444;
  color: white;
}

/* 文件控制区域 */
.file-controls {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  background: #f8fafc;
  border-bottom: 1px solid #e2e8f0;
}

.file-controls.dark {
  background: #334155;
  border-color: #475569;
}

.filter-section {
  display: flex;
  gap: 1.5rem;
}

.filter-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.filter-group label {
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  white-space: nowrap;
}

.filter-group.dark label {
  color: #e5e7eb;
}

.filter-select {
  padding: 0.375rem 0.75rem;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  background: white;
  font-size: 0.875rem;
  color: #374151;
}

.filter-select.dark {
  background: #1e293b;
  border-color: #475569;
  color: #e5e7eb;
}

.view-options {
  display: flex;
  gap: 0.25rem;
}

.view-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  background: #f1f5f9;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #64748b;
}

.view-btn:hover {
  background: #e2e8f0;
  border-color: #cbd5e1;
}

.view-btn.active {
  background: #3b82f6;
  border-color: #3b82f6;
  color: white;
}

/* 文件区域 */
.files-section {
  flex: 1;
  padding: 1.5rem;
  overflow-y: auto;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }
  
  .header-left {
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
  }
  
  .file-controls {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }
  
  .filter-section {
    flex-direction: column;
    gap: 1rem;
  }
}
</style>
