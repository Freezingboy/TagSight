<template>
  <!-- 头部组件 - 提供搜索、筛选和操作功能 -->
  <header class="header" :class="{ dark: isDarkTheme }">
    <!-- 搜索和筛选区域 -->
    <div class="search-container">
      <!-- 搜索栏 -->
      <div class="search-bar" :class="{ dark: isDarkTheme }">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="11" cy="11" r="8"/>
          <path d="m21 21-4.35-4.35"/>
        </svg>
        <input 
          type="text" 
          :placeholder="searchPlaceholder" 
          class="search-input"
          :class="{ dark: isDarkTheme }"
        />
      </div>
      
      <!-- 筛选选项区域 -->
      <div class="header-filters">
        <!-- 分类选择下拉框 -->
        <div class="category-select">
          <button class="dropdown-btn" :class="{ dark: isDarkTheme }" @click="toggleCategoryDropdown">
            {{ selectedCategory }}
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="6,9 12,15 18,9"/>
            </svg>
          </button>
          <!-- 分类下拉菜单 -->
          <div class="dropdown-menu" :class="{ dark: isDarkTheme }" v-if="showCategoryDropdown">
            <div 
              v-for="category in categories" 
              :key="category.name"
              class="dropdown-item"
              :class="{ dark: isDarkTheme }"
              @click="selectCategory(category.name)"
            >
              <!-- 不同分类的图标 -->
              <svg v-if="category.icon === 'image'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect width="18" height="18" x="3" y="3" rx="2" ry="2"/>
                <circle cx="9" cy="9" r="2"/>
                <path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/>
              </svg>
              <svg v-else-if="category.icon === 'audio'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M9 18V5l12-2v13"/>
                <circle cx="6" cy="18" r="3"/>
                <circle cx="18" cy="16" r="3"/>
              </svg>
              <svg v-else-if="category.icon === 'video'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="m22 8-6 4 6 4V8Z"/>
                <rect width="14" height="12" x="2" y="6" rx="2" ry="2"/>
              </svg>
              <svg v-else-if="category.icon === 'document'" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/>
                <polyline points="14,2 14,8 20,8"/>
              </svg>
              {{ category.name }}
            </div>
          </div>
        </div>
        
        <!-- 筛选条件下拉框 -->
        <div class="filter-select">
          <button class="dropdown-btn" :class="{ dark: isDarkTheme }" @click="toggleFilterDropdown">
            {{ selectedFilter }}
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="6,9 12,15 18,9"/>
            </svg>
          </button>
          <!-- 筛选条件下拉菜单 -->
          <div class="dropdown-menu" :class="{ dark: isDarkTheme }" v-if="showFilterDropdown">
            <div 
              v-for="filter in filters" 
              :key="filter.name"
              class="dropdown-item"
              :class="{ dark: isDarkTheme }"
              @click="selectFilter(filter.name)"
            >
              {{ filter.name }}
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 操作按钮区域 -->
    <div class="header-actions">
      <!-- 上传文件按钮 - 根据showUploadButton属性显示 -->
      <button class="upload-btn" :class="{ dark: isDarkTheme }" v-if="showUploadButton">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
          <polyline points="7,10 12,15 17,10"/>
          <line x1="12" y1="15" x2="12" y2="3"/>
        </svg>
        上传文件
      </button>
      
      <!-- 新建标签按钮 - 根据showNewTagButton属性显示 -->
      <button class="new-tag-btn" :class="{ dark: isDarkTheme }" v-if="showNewTagButton">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <line x1="12" y1="5" x2="12" y2="19"/>
          <line x1="5" y1="12" x2="19" y2="12"/>
        </svg>
        新建标签
      </button>
      
      <!-- 主题切换按钮 -->
      <button class="theme-toggle-btn" :class="{ dark: isDarkTheme }" @click="toggleTheme">
        <!-- 暗色主题图标 -->
        <svg v-if="isDarkTheme" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="4"/>
          <path d="M12 2v2"/>
          <path d="M12 20v2"/>
          <path d="m4.93 4.93 1.41 1.41"/>
          <path d="m17.66 17.66 1.41 1.41"/>
          <path d="M2 12h2"/>
          <path d="M20 12h2"/>
          <path d="m6.34 17.66-1.41 1.41"/>
          <path d="m19.07 4.93-1.41 1.41"/>
        </svg>
        <!-- 亮色主题图标 -->
        <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"/>
        </svg>
        {{ isDarkTheme ? '开灯' : '关灯' }}
      </button>
    </div>
  </header>
</template>

<script setup lang="ts">
// 导入Vue 3的组合式API
import { ref, onMounted, onUnmounted } from 'vue'

// ==================== 组件属性定义 ====================

// 定义组件的属性接口
interface Props {
  // 是否显示上传按钮
  showUploadButton?: boolean
  // 是否显示新建标签按钮
  showNewTagButton?: boolean
  // 搜索占位符文本
  searchPlaceholder?: string
  // 当前主题模式
  isDarkTheme?: boolean
}

// 定义组件的属性，设置默认值
const props = withDefaults(defineProps<Props>(), {
  showUploadButton: false,
  showNewTagButton: false,
  searchPlaceholder: '搜索...',
  isDarkTheme: false
})

// ==================== 响应式数据 ====================

// 当前选中的分类
const selectedCategory = ref('全部文件')

// 当前选中的筛选条件
const selectedFilter = ref('文件大小')

// 分类下拉框显示状态
const showCategoryDropdown = ref(false)

// 筛选下拉框显示状态
const showFilterDropdown = ref(false)

// ==================== 配置数据 ====================

// 分类选项配置
const categories = [
  { name: '全部文件', icon: 'folder' },
  { name: '图片', icon: 'image' },
  { name: '音频', icon: 'audio' },
  { name: '视频', icon: 'video' },
  { name: '文档', icon: 'document' }
]

// 筛选条件配置
const filters = [
  { name: '文件大小' },
  { name: '修改日期' },
  { name: '创建日期' },
  { name: '文件类型' }
]

// ==================== 方法定义 ====================

/**
 * 切换主题模式
 * 触发父组件的主题切换事件
 */
const emit = defineEmits<{
  'toggle-theme': []
  'select-category': [category: string]
  'select-filter': [filter: string]
}>()

const toggleTheme = () => {
  emit('toggle-theme')
}

/**
 * 切换分类下拉框显示状态
 * 同时关闭筛选下拉框以保持界面整洁
 */
const toggleCategoryDropdown = () => {
  showCategoryDropdown.value = !showCategoryDropdown.value
  if (showCategoryDropdown.value) {
    showFilterDropdown.value = false
  }
}

/**
 * 选择分类
 * @param categoryName 分类名称
 */
const selectCategory = (categoryName: string) => {
  selectedCategory.value = categoryName
  showCategoryDropdown.value = false
  emit('select-category', categoryName)
}

/**
 * 切换筛选下拉框显示状态
 * 同时关闭分类下拉框以保持界面整洁
 */
const toggleFilterDropdown = () => {
  showFilterDropdown.value = !showFilterDropdown.value
  if (showFilterDropdown.value) {
    showCategoryDropdown.value = false
  }
}

/**
 * 选择筛选条件
 * @param filterName 筛选条件名称
 */
const selectFilter = (filterName: string) => {
  selectedFilter.value = filterName
  showFilterDropdown.value = false
  emit('select-filter', filterName)
}

/**
 * 点击外部关闭下拉框
 * 提升用户体验，避免下拉框一直显示
 */
const handleClickOutside = (event: Event) => {
  const target = event.target as HTMLElement
  if (!target.closest('.category-select') && !target.closest('.filter-select')) {
    showCategoryDropdown.value = false
    showFilterDropdown.value = false
  }
}

// ==================== 生命周期钩子 ====================

/**
 * 组件挂载时添加全局点击事件监听
 * 用于关闭下拉框
 */
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

/**
 * 组件卸载时移除全局点击事件监听
 * 防止内存泄漏
 */
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: #ffffff;
  border-bottom: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

/* 暗色主题下的头部样式 */
.header.dark {
  background-color: #1f2937;
  border-bottom-color: #374151;
}

.search-container {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 1;
  max-width: 600px;
}

.search-bar {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background-color: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 0.5rem 0.75rem;
  flex: 1;
  transition: all 0.2s ease;
}

/* 暗色主题下的搜索栏 */
.search-bar.dark {
  background-color: #374151;
  border-color: #525252;
}

.search-bar:focus-within {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

/* 暗色主题下的搜索栏焦点状态 */
.search-bar.dark:focus-within {
  border-color: #60a5fa;
  box-shadow: 0 0 0 3px rgba(96, 165, 250, 0.1);
}

.search-input {
  border: none;
  outline: none;
  background: transparent;
  font-size: 0.875rem;
  color: #1f2937;
  width: 100%;
}

/* 暗色主题下的搜索输入框 */
.search-input.dark {
  color: #ffffff;
}

.search-input::placeholder {
  color: #9ca3af;
}

/* 暗色主题下的搜索输入框占位符 */
.search-input.dark::placeholder {
  color: #6b7280;
}

.header-filters {
  display: flex;
  gap: 0.5rem;
}

.category-select,
.filter-select {
  position: relative;
}

.dropdown-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  font-size: 0.875rem;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s ease;
  white-space: nowrap;
}

/* 暗色主题下的下拉按钮 */
.dropdown-btn.dark {
  background-color: #374151;
  border-color: #525252;
  color: #ffffff;
}

.dropdown-btn:hover {
  border-color: #3b82f6;
}

/* 暗色主题下的下拉按钮悬停效果 */
.dropdown-btn.dark:hover {
  border-color: #60a5fa;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  z-index: 10;
  margin-top: 0.25rem;
  min-width: 120px;
}

/* 暗色主题下的下拉菜单 */
.dropdown-menu.dark {
  background-color: #2d2d2d;
  border-color: #525252;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  color: #374151;
  cursor: pointer;
  transition: background-color 0.2s ease;
  border-bottom: 1px solid #e2e8f0;
}

/* 暗色主题下的下拉菜单项 */
.dropdown-item.dark {
  color: #ffffff;
  border-bottom-color: #525252;
}

.dropdown-item:last-child {
  border-bottom: none;
}

.dropdown-item:hover {
  background-color: #f1f5f9;
}

/* 暗色主题下的下拉菜单项悬停效果 */
.dropdown-item.dark:hover {
  background-color: #404040;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.upload-btn,
.new-tag-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.upload-btn:hover,
.new-tag-btn:hover {
  background-color: #2563eb;
  transform: translateY(-1px);
}

/* 暗色主题下的上传和新建标签按钮 */
.upload-btn.dark,
.new-tag-btn.dark {
  background-color: #6366f1;
}

.upload-btn.dark:hover,
.new-tag-btn.dark:hover {
  background-color: #4f46e5;
}

.theme-toggle-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  background-color: transparent;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  color: #374151;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.875rem;
}

/* 暗色主题下的主题切换按钮 */
.theme-toggle-btn.dark {
  border-color: #525252;
  color: #ffffff;
}

.theme-toggle-btn:hover {
  background-color: #f1f5f9;
  border-color: #3b82f6;
}

/* 暗色主题下的主题切换按钮悬停效果 */
.theme-toggle-btn.dark:hover {
  background-color: #404040;
  border-color: #60a5fa;
}
</style> 