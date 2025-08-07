<template>
  <!-- 文件管理器主组件 - 整个应用的核心容器 -->
  <div class="file-manager" :class="{ dark: isDarkTheme }">
    <!-- 顶部导航栏 -->
    <header class="header">
      <!-- 左侧区域 - 标题和图标 -->
      <div class="header-left">
        <!-- 根据当前页面显示不同的图标 -->
        <svg v-if="selectedSidebarItem === 'home'" class="folder-icon" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
          <polyline points="9,22 9,12 15,12 15,22"/>
        </svg>
        <svg v-else class="folder-icon" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.89l-.812-1.22A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/>
        </svg>
        <!-- 动态标题 - 根据当前页面显示不同标题 -->
        <h1>{{ selectedSidebarItem === 'home' ? '首页' : '文件管理系统' }}</h1>
      </div>
      
      <!-- 右侧区域 - 搜索栏和操作按钮 -->
      <div class="header-right">
        <!-- 搜索栏 - 只在非首页显示 -->
        <div class="search-bar" v-if="selectedSidebarItem !== 'home'">
          <svg class="search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="11" cy="11" r="8"/>
            <path d="m21 21-4.35-4.35"/>
          </svg>
          <!-- 动态搜索占位符 - 根据当前页面显示不同提示 -->
          <input type="text" :placeholder="selectedSidebarItem === 'tags' ? '搜索标签...' : '搜索文件...'" />
        </div>
        
        <!-- 文件分类下拉框 - 只在文件管理页面显示 -->
        <div v-if="selectedSidebarItem === 'files'" class="category-select" @click="toggleCategoryDropdown">
          <div class="select-display">
            <span>{{ selectedCategoryLabel }}</span>
            <svg class="chevron-down" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="6,9 12,15 18,9"/>
            </svg>
          </div>
          <!-- 分类选项下拉菜单 -->
          <div class="dropdown-menu" v-show="categoryDropdownOpen">
            <div 
              v-for="category in categoryOptions" 
              :key="category.value"
              class="dropdown-item"
              :class="{ active: selectedCategory === category.value }"
              @click="selectCategory(category.value)"
            >
              <!-- 不同分类的图标 -->
              <svg class="category-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect v-if="category.value === 'all'" x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                <rect v-if="category.value === 'images'" x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                <circle v-if="category.value === 'images'" cx="8.5" cy="8.5" r="1.5"/>
                <polyline v-if="category.value === 'images'" points="21,15 16,10 5,21"/>
                <path v-if="category.value === 'audio'" d="M9 18V5l12-2v13"/>
                <circle v-if="category.value === 'audio'" cx="6" cy="18" r="3"/>
                <circle v-if="category.value === 'audio'" cx="18" cy="16" r="3"/>
                <polygon v-if="category.value === 'video'" points="23,7 16,12 23,17 23,7"/>
                <rect v-if="category.value === 'video'" x="1" y="5" width="15" height="14" rx="2" ry="2"/>
                <path v-if="category.value === 'documents'" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline v-if="category.value === 'documents'" points="14,2 14,8 20,8"/>
                <line v-if="category.value === 'documents'" x1="16" y1="13" x2="8" y2="13"/>
                <line v-if="category.value === 'documents'" x1="16" y1="17" x2="8" y2="17"/>
                <polyline v-if="category.value === 'documents'" points="10,9 9,9 8,9"/>
                <rect v-if="category.value === 'archives'" x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                <path v-if="category.value === 'archives'" d="M16 2v4"/>
                <path v-if="category.value === 'archives'" d="M8 2v4"/>
                <path v-if="category.value === 'archives'" d="M3 10h18"/>
              </svg>
              {{ category.label }}
            </div>
          </div>
        </div>
        
        <!-- 文件大小筛选下拉框 - 只在文件管理页面显示 -->
        <div v-if="selectedSidebarItem === 'files'" class="filter-select" @click="toggleSizeDropdown">
          <div class="select-display">
            <span>{{ selectedSize }}</span>
            <svg class="chevron-down" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="6,9 12,15 18,9"/>
            </svg>
          </div>
          <!-- 文件大小筛选选项 -->
          <div class="dropdown-menu" v-show="sizeDropdownOpen">
            <div 
              v-for="option in sizeOptions" 
              :key="option.value"
              class="dropdown-item"
              :class="{ active: selectedSize === option.label }"
              @click="selectSize(option.label)"
            >
              {{ option.label }}
            </div>
          </div>
        </div>
        
        <!-- 修改日期筛选下拉框 - 只在文件管理页面显示 -->
        <div v-if="selectedSidebarItem === 'files'" class="filter-select" @click="toggleDateDropdown">
          <div class="select-display">
            <span>{{ selectedDate }}</span>
            <svg class="chevron-down" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="6,9 12,15 18,9"/>
            </svg>
          </div>
          <!-- 修改日期筛选选项 -->
          <div class="dropdown-menu" v-show="dateDropdownOpen">
            <div 
              v-for="option in dateOptions" 
              :key="option.value"
              class="dropdown-item"
              :class="{ active: selectedDate === option.label }"
              @click="selectDate(option.label)"
            >
              {{ option.label }}
            </div>
          </div>
        </div>
        
        <!-- 新建标签按钮 - 只在标签管理页面显示 -->
        <button v-if="selectedSidebarItem === 'tags'" class="new-tag-btn">
          <svg class="plus-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="12" y1="5" x2="12" y2="19"/>
            <line x1="5" y1="12" x2="19" y2="12"/>
          </svg>
          新建标签
        </button>
        
        <!-- 上传按钮 - 只在文件管理页面显示 -->
        <button v-if="selectedSidebarItem === 'files'" class="upload-btn">
          <svg class="upload-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
            <polyline points="7,10 12,15 17,10"/>
            <line x1="12" y1="15" x2="12" y2="3"/>
          </svg>
          上传
        </button>
        
        <!-- 主题切换按钮 - 全局显示 -->
        <button class="theme-toggle-btn" @click="toggleTheme">
          <!-- 亮色主题图标 -->
          <svg v-if="!isDarkTheme" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="5"/>
            <path d="M12 1v2"/>
            <path d="M12 21v2"/>
            <path d="M4.22 4.22l1.42 1.42"/>
            <path d="M18.36 18.36l1.42 1.42"/>
            <path d="M1 12h2"/>
            <path d="M21 12h2"/>
            <path d="M4.22 19.78l1.42-1.42"/>
            <path d="M18.36 5.64l1.42-1.42"/>
          </svg>
          <!-- 暗色主题图标 -->
          <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"/>
          </svg>
          {{ isDarkTheme ? '开灯' : '关灯' }}
        </button>
      </div>
    </header>

    <!-- 主要内容区域 -->
    <div class="main-content">
      <!-- 侧边栏组件 -->
      <Sidebar 
        :selectedSidebarItem="selectedSidebarItem"
        :isDarkTheme="isDarkTheme"
        @update:selectedSidebarItem="selectedSidebarItem = $event"
      />

      <!-- 主内容区域 -->
      <main class="main-area">
        <!-- 首页内容 -->
        <div v-if="selectedSidebarItem === 'home'">
          <Homepage :isDarkTheme="isDarkTheme" />
        </div>

        <!-- 文件管理页面 -->
        <div v-if="selectedSidebarItem === 'files'">
          <!-- 文件管理页面头部 -->
          <div class="main-header">
            <div class="main-title">
              <h2>全部文件</h2>
              <span class="file-count">共128个文件</span>
            </div>
            <!-- 视图切换选项 -->
            <div class="view-options">
              <!-- 网格视图按钮 -->
              <button class="view-btn active">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <rect x="3" y="3" width="7" height="7"/>
                  <rect x="14" y="3" width="7" height="7"/>
                  <rect x="14" y="14" width="7" height="7"/>
                  <rect x="3" y="14" width="7" height="7"/>
                </svg>
              </button>
              <!-- 列表视图按钮 -->
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

          <!-- 上传进度组件 -->
          <UploadProgress :isDarkTheme="isDarkTheme" />

          <!-- 文件网格组件 -->
          <FileGrid :isDarkTheme="isDarkTheme" @file-click="handleFileCardClick" />

          <!-- 分页组件 -->
          <Pagination :isDarkTheme="isDarkTheme" />
        </div>

        <!-- 标签管理页面 -->
        <div v-if="selectedSidebarItem === 'tags'" class="tags-page">
          <div class="tags-header">
            <h2>标签管理</h2>
          </div>
          
          <!-- 标签网格组件 -->
          <TagGrid :isDarkTheme="isDarkTheme" />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
// 导入Vue 3的组合式API
import { ref, onMounted, onUnmounted, computed } from 'vue'
// 导入子组件
import Sidebar from './Sidebar.vue'
import FileGrid from './FileGrid.vue'
import TagGrid from './TagGrid.vue'
import Pagination from './Pagination.vue'
import UploadProgress from './UploadProgress.vue'
import Homepage from './Homepage.vue'

// ==================== 响应式数据定义 ====================

// 当前选中的侧边栏项目 - 控制页面切换
const selectedSidebarItem = ref('home')

// 当前选中的文件分类
const selectedCategory = ref('all')

// 主题模式状态 - 控制明暗主题切换
const isDarkTheme = ref(false)

// ==================== 配置数据 ====================

// 侧边栏选项配置
const sidebarOptions = [
  { value: 'files', label: '文件管理', icon: 'file' },
  { value: 'tags', label: '标签管理', icon: 'tag' }
]

// 文件分类选项配置
const categoryOptions = [
  { value: 'all', label: '全部文件', icon: 'folder' },
  { value: 'images', label: '图片', icon: 'image' },
  { value: 'audio', label: '音频', icon: 'music' },
  { value: 'video', label: '视频', icon: 'video' },
  { value: 'documents', label: '文档', icon: 'file-text' },
  { value: 'archives', label: '压缩包', icon: 'archive' }
]

// 计算属性：获取当前选中分类的显示标签
const selectedCategoryLabel = computed(() => {
  const category = categoryOptions.find(cat => cat.value === selectedCategory.value)
  return category ? category.label : '全部文件'
})

// 分类下拉框显示状态
const categoryDropdownOpen = ref(false)

// 文件大小筛选选项配置
const sizeOptions = [
  { label: '全部', value: 'all' },
  { label: '小于1MB', value: 'lt1mb' },
  { label: '1MB-5MB', value: '1mb5mb' },
  { label: '5MB-10MB', value: '5mb10mb' },
  { label: '大于10MB', value: 'gt10mb' },
]
const selectedSize = ref('全部')
const sizeDropdownOpen = ref(false)

// 修改日期筛选选项配置
const dateOptions = [
  { label: '全部', value: 'all' },
  { label: '今天', value: 'today' },
  { label: '本周', value: 'week' },
  { label: '本月', value: 'month' },
  { label: '今年', value: 'year' },
]
const selectedDate = ref('全部')
const dateDropdownOpen = ref(false)

// ==================== 方法定义 ====================

/**
 * 切换主题模式
 * 在明暗主题之间切换
 */
const toggleTheme = () => {
  isDarkTheme.value = !isDarkTheme.value
}

/**
 * 处理用户信息点击事件
 * 预留的用户菜单功能
 */
const handleUserClick = () => {
  console.log('用户信息被点击')
  // 这里可以添加用户菜单或其他功能
}

/**
 * 切换文件分类下拉框显示状态
 * 同时关闭其他下拉框以保持界面整洁
 */
const toggleCategoryDropdown = () => {
  categoryDropdownOpen.value = !categoryDropdownOpen.value
  if (categoryDropdownOpen.value) {
    sizeDropdownOpen.value = false
    dateDropdownOpen.value = false
  }
}

/**
 * 选择文件分类
 * @param value 分类值
 */
const selectCategory = (value: string) => {
  selectedCategory.value = value
  categoryDropdownOpen.value = false
}

/**
 * 切换文件大小筛选下拉框显示状态
 * 同时关闭其他下拉框
 */
const toggleSizeDropdown = () => {
  sizeDropdownOpen.value = !sizeDropdownOpen.value
  if (sizeDropdownOpen.value) {
    categoryDropdownOpen.value = false
    dateDropdownOpen.value = false
  }
}

/**
 * 选择文件大小筛选条件
 * @param label 大小标签
 */
const selectSize = (label: string) => {
  selectedSize.value = label
  sizeDropdownOpen.value = false
}

/**
 * 切换修改日期筛选下拉框显示状态
 * 同时关闭其他下拉框
 */
const toggleDateDropdown = () => {
  dateDropdownOpen.value = !dateDropdownOpen.value
  if (dateDropdownOpen.value) {
    categoryDropdownOpen.value = false
    sizeDropdownOpen.value = false
  }
}

/**
 * 选择修改日期筛选条件
 * @param label 日期标签
 */
const selectDate = (label: string) => {
  selectedDate.value = label
  dateDropdownOpen.value = false
}

/**
 * 处理文件卡片点击事件
 * 提供视觉反馈和后续操作
 * @param fileName 文件名
 */
const handleFileCardClick = (fileName: string) => {
  console.log('文件被点击:', fileName)
  // 这里可以添加文件预览、下载或其他操作
  // 可以添加一个短暂的视觉反馈
  const card = event?.target?.closest('.file-card') as HTMLElement
  if (card) {
    card.style.transform = 'translateY(-3px) scale(1.02)'
    card.style.boxShadow = '0 6px 20px rgba(0, 0, 0, 0.25)'
    setTimeout(() => {
      card.style.transform = ''
      card.style.boxShadow = ''
    }, 150)
  }
}

/**
 * 点击外部关闭下拉框
 * 提升用户体验，避免下拉框一直显示
 * @param event 点击事件
 */
const handleClickOutside = (event: Event) => {
  const target = event.target as Element
  if (!target.closest('.custom-select') && !target.closest('.category-select') && !target.closest('.filter-select')) {
    sizeDropdownOpen.value = false
    dateDropdownOpen.value = false
    categoryDropdownOpen.value = false
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
/* ==================== 主容器样式 ==================== */

/* 文件管理器主容器 */
.file-manager {
  width: 100vw;  /* 占满视口宽度 */
  height: 100vh; /* 占满视口高度 */
  min-height: 100vh;
  min-width: 100vw;
  background-color: #ffffff; /* 亮色主题背景 */
  color: #1f2937;           /* 亮色主题文字颜色 */
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  transition: all 0.3s ease; /* 平滑过渡动画 */
}

/* 暗色主题样式覆盖 */
.file-manager.dark {
  background-color: #1a1a1a; /* 暗色主题背景 */
  color: #ffffff;            /* 暗色主题文字颜色 */
}

/* ==================== 顶部导航栏样式 ==================== */

.header {
  display: flex;
  justify-content: space-between; /* 左右两端对齐 */
  align-items: center;
  padding: 1rem 2rem;
  background-color: #f8fafc; /* 导航栏背景色 */
  border-bottom: 1px solid #e2e8f0; /* 底部边框 */
  transition: all 0.3s ease;
}

/* 暗色主题导航栏样式 */
.file-manager.dark .header {
  background-color: #2d2d2d;
  border-bottom: 1px solid #404040;
}

/* 导航栏左侧区域 */
.header-left {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

/* 文件夹图标样式 */
.folder-icon {
  color: #4f46e5; /* 主题色 */
}

/* 导航栏标题 */
.header-left h1 {
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0;
  color: #1f2937;
}

/* 暗色主题标题颜色 */
.file-manager.dark .header-left h1 {
  color: #ffffff;
}

/* 导航栏右侧区域 */
.header-right {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* ==================== 搜索栏样式 ==================== */

.search-bar {
  display: flex;
  align-items: center;
  background-color: #f1f5f9;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 0.5rem 0.75rem;
  gap: 0.5rem;
  transition: all 0.3s ease;
}

/* 搜索栏悬停效果 */
.search-bar:hover {
  border-color: #cbd5e1;
  background-color: #ffffff;
}

/* 搜索图标 */
.search-icon {
  color: #64748b;
  flex-shrink: 0;
}

/* 搜索输入框 */
.search-bar input {
  border: none;
  background: none;
  outline: none;
  font-size: 0.875rem;
  color: #1f2937;
  width: 200px;
}

/* 搜索输入框占位符 */
.search-bar input::placeholder {
  color: #9ca3af;
}

/* ==================== 下拉框样式 ==================== */

/* 分类选择下拉框 */
.category-select {
  position: relative;
  cursor: pointer;
}

/* 筛选选择下拉框 */
.filter-select {
  position: relative;
  cursor: pointer;
}

/* 下拉框显示区域 */
.select-display {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background-color: #f1f5f9;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  color: #1f2937;
  transition: all 0.3s ease;
  min-width: 120px;
}

/* 暗色主题下的下拉框显示区域 */
.file-manager.dark .select-display {
  background-color: #1f2937;
  border-color: #374151;
  color: #e5e7eb;
}

/* 下拉框悬停效果 */
.select-display:hover {
  border-color: #cbd5e1;
  background-color: #ffffff;
}

/* 暗色主题下的下拉框悬停效果 */
.file-manager.dark .select-display:hover {
  border-color: #4f46e5;
  background-color: #2d2d2d;
}

/* 下拉箭头图标 */
.chevron-down {
  color: #64748b;
  transition: transform 0.3s ease;
}

/* 下拉菜单 */
.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  margin-top: 0.25rem;
  max-height: 200px;
  overflow-y: auto;
}

/* 暗色主题下的下拉菜单 */
.file-manager.dark .dropdown-menu {
  background-color: #1f2937;
  border-color: #374151;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4);
}

/* 下拉菜单项 */
.dropdown-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  color: #1f2937;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

/* 暗色主题下的下拉菜单项 */
.file-manager.dark .dropdown-item {
  color: #e5e7eb;
}

/* 下拉菜单项悬停效果 */
.dropdown-item:hover {
  background-color: #f1f5f9;
}

/* 暗色主题下的下拉菜单项悬停效果 */
.file-manager.dark .dropdown-item:hover {
  background-color: #374151;
}

/* 激活状态的下拉菜单项 */
.dropdown-item.active {
  background-color: #eef2ff;
  color: #4f46e5;
}

/* 暗色主题下的激活状态下拉菜单项 */
.file-manager.dark .dropdown-item.active {
  background-color: #4f46e5;
  color: #ffffff;
}

/* 分类图标 */
.category-icon {
  color: #64748b;
  flex-shrink: 0;
}

/* 暗色主题下的分类图标 */
.file-manager.dark .category-icon {
  color: #9ca3af;
}

/* ==================== 按钮样式 ==================== */

/* 新建标签按钮 */
.new-tag-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background-color: #4f46e5;
  color: white;
  border: none;
  border-radius: 6px;
  padding: 0.5rem 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

/* 新建标签按钮悬停效果 */
.new-tag-btn:hover {
  background-color: #4338ca;
  transform: translateY(-1px);
}

/* 上传按钮 */
.upload-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background-color: #10b981;
  color: white;
  border: none;
  border-radius: 6px;
  padding: 0.5rem 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

/* 上传按钮悬停效果 */
.upload-btn:hover {
  background-color: #059669;
  transform: translateY(-1px);
}

/* 主题切换按钮 */
.theme-toggle-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background-color: #f1f5f9;
  color: #1f2937;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

/* 主题切换按钮悬停效果 */
.theme-toggle-btn:hover {
  background-color: #e2e8f0;
  border-color: #cbd5e1;
}

/* ==================== 主要内容区域样式 ==================== */

.main-content {
  display: flex;
  height: calc(100vh - 80px); /* 减去导航栏高度 */
}

/* 主内容区域 */
.main-area {
  flex: 1;
  padding: 1.5rem;
  overflow-y: auto;
  background-color: #ffffff;
}

/* 暗色主题主内容区域 */
.file-manager.dark .main-area {
  background-color: #1a1a1a;
}

/* ==================== 文件管理页面样式 ==================== */

/* 文件管理页面头部 */
.main-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

/* 主标题区域 */
.main-title {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* 主标题 */
.main-title h2 {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
  color: #1f2937;
}

/* 文件数量显示 */
.file-count {
  font-size: 0.875rem;
  color: #6b7280;
  background-color: #f3f4f6;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
}

/* 视图切换选项 */
.view-options {
  display: flex;
  gap: 0.5rem;
}

/* 视图切换按钮 */
.view-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  background-color: #f1f5f9;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #64748b;
}

/* 视图切换按钮悬停效果 */
.view-btn:hover {
  background-color: #e2e8f0;
  border-color: #cbd5e1;
}

/* 激活状态的视图切换按钮 */
.view-btn.active {
  background-color: #4f46e5;
  border-color: #4f46e5;
  color: white;
}

/* ==================== 标签管理页面样式 ==================== */

/* 标签页面容器 */
.tags-page {
  height: 100%;
}

/* 标签页面头部 */
.tags-header {
  margin-bottom: 1.5rem;
}

/* 标签页面标题 */
.tags-header h2 {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
  color: #1f2937;
}

/* ==================== 响应式设计 ==================== */

/* 平板设备适配 */
@media (max-width: 768px) {
  .header {
    padding: 0.75rem 1rem;
  }
  
  .header-right {
    gap: 0.5rem;
  }
  
  .search-bar input {
    width: 150px;
  }
  
  .main-area {
    padding: 1rem;
  }
  
  .main-header {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }
}

/* 手机设备适配 */
@media (max-width: 480px) {
  .header-left h1 {
    font-size: 1rem;
  }
  
  .search-bar {
    display: none;
  }
  
  .category-select,
  .filter-select {
    display: none;
  }
  
  .main-area {
    padding: 0.75rem;
  }
}

/* ==================== 暗色主题适配 ==================== */

/* 暗色主题下的文字颜色 */
.file-manager.dark .main-title h2,
.file-manager.dark .tags-header h2 {
  color: #ffffff;
}

/* 暗色主题下的文件数量 */
.file-manager.dark .file-count {
  color: #9ca3af;
  background-color: #374151;
}

/* 暗色主题下的视图切换按钮 */
.file-manager.dark .view-btn {
  background-color: #374151;
  border-color: #4b5563;
  color: #9ca3af;
}

.file-manager.dark .view-btn:hover {
  background-color: #4b5563;
  border-color: #6b7280;
}

.file-manager.dark .view-btn.active {
  background-color: #6366f1;
  border-color: #6366f1;
  color: white;
}
</style> 