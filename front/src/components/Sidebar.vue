<template>
  <!-- 侧边栏组件 - 提供导航菜单和用户信息 -->
  <aside class="sidebar" :class="{ dark: isDarkTheme }">
    <!-- 导航菜单区域 -->
    <div class="sidebar-section">
      <h3>导航菜单</h3>
      <ul>
        <!-- 首页菜单项 -->
        <li 
          :class="{ active: selectedSidebarItem === 'home' }"
          @click="$emit('update:selectedSidebarItem', 'home')"
        >
          <svg class="sidebar-icon" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
            <polyline points="9,22 9,12 15,12 15,22"/>
          </svg>
          <span class="sidebar-text">首页</span>
        </li>
        
        <!-- 文件管理菜单项 -->
        <li 
          :class="{ active: selectedSidebarItem === 'files' }"
          @click="$emit('update:selectedSidebarItem', 'files')"
        >
          <svg class="sidebar-icon" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
            <polyline points="14,2 14,8 20,8"/>
            <line x1="16" y1="13" x2="8" y2="13"/>
            <line x1="16" y1="17" x2="8" y2="17"/>
            <polyline points="10,9 9,9 8,9"/>
          </svg>
          <span class="sidebar-text">文件管理</span>
        </li>
        
        <!-- 标签管理菜单项 -->
        <li 
          :class="{ active: selectedSidebarItem === 'tags' }"
          @click="$emit('update:selectedSidebarItem', 'tags')"
        >
          <svg class="sidebar-icon" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"/>
            <line x1="7" y1="7" x2="7.01" y2="7"/>
          </svg>
          <span class="sidebar-text">标签管理</span>
        </li>
      </ul>
    </div>
    
    <!-- 用户信息区域 -->
    <div class="user-info" @click="handleUserClick">
      <!-- 用户头像 -->
      <div class="user-avatar">
        <svg class="user-icon" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
          <circle cx="12" cy="7" r="4"/>
        </svg>
      </div>
      <!-- 用户详细信息 -->
      <div class="user-details">
        <span class="user-name">管理员</span>
        <span class="user-role">系统管理员</span>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
// ==================== 组件属性定义 ====================

// 定义组件的属性接口
interface Props {
  // 当前选中的侧边栏项目
  selectedSidebarItem: string
  // 当前主题模式
  isDarkTheme?: boolean
}

// 定义组件的事件接口
interface Emits {
  // 更新选中的侧边栏项目事件
  (e: 'update:selectedSidebarItem', value: string): void
}

// 声明组件属性和事件
const props = withDefaults(defineProps<Props>(), {
  isDarkTheme: false
})
defineEmits<Emits>()

// ==================== 方法定义 ====================

/**
 * 处理用户信息点击事件
 * 预留的用户菜单功能
 */
const handleUserClick = () => {
  console.log('用户信息被点击')
  // 这里可以添加用户菜单或其他功能
}
</script>

<style scoped>
/* ==================== 侧边栏主容器样式 ==================== */

/* 侧边栏主容器 */
.sidebar {
  width: 280px; /* 固定宽度 */
  background-color: #f8fafc; /* 亮色主题背景 */
  border-right: 1px solid #e2e8f0; /* 右侧边框 */
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  justify-content: space-between; /* 上下分布 */
  transition: all 0.3s ease; /* 平滑过渡动画 */
  position: relative;
}

/* 暗色主题侧边栏样式 */
.sidebar.dark {
  background-color: #1f2937;
  border-right: 1px solid #374151;
}

/* ==================== 导航菜单区域样式 ==================== */

/* 侧边栏区域 */
.sidebar-section h3 {
  margin: 0 0 1rem 0;
  font-size: 0.875rem;
  font-weight: 600;
  color: #6b7280; /* 次要文字颜色 */
  text-transform: uppercase; /* 大写字母 */
  letter-spacing: 0.05em; /* 字母间距 */
}

/* 暗色主题下的标题颜色 */
.sidebar.dark .sidebar-section h3 {
  color: #d1d5db;
}

/* 导航菜单列表 */
.sidebar-section ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

/* 导航菜单项 */
.sidebar-section li {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  margin-bottom: 0.25rem;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #374151; /* 主要文字颜色 */
}

/* 暗色主题下的导航菜单项 */
.sidebar.dark .sidebar-section li {
  color: #e5e7eb;
}

/* 导航菜单项悬停效果 */
.sidebar-section li:hover {
  background-color: #e5e7eb;
  transform: translateX(4px); /* 轻微右移效果 */
}

/* 暗色主题下的导航菜单项悬停效果 */
.sidebar.dark .sidebar-section li:hover {
  background-color: #374151;
}

/* 激活状态的导航菜单项 */
.sidebar-section li.active {
  background-color: #4f46e5; /* 主题色背景 */
  color: white;
  box-shadow: 0 2px 4px rgba(79, 70, 229, 0.2);
}

/* 激活状态悬停效果 */
.sidebar-section li.active:hover {
  background-color: #4338ca;
  transform: translateX(4px);
}

/* 侧边栏图标 */
.sidebar-icon {
  color: #6b7280; /* 默认图标颜色 */
  transition: color 0.3s ease;
  flex-shrink: 0; /* 防止图标被压缩 */
}

/* 暗色主题下的侧边栏图标 */
.sidebar.dark .sidebar-icon {
  color: #9ca3af;
}

/* 悬停时的图标颜色 */
.sidebar-section li:hover .sidebar-icon {
  color: #4f46e5;
}

/* 暗色主题下的悬停图标颜色 */
.sidebar.dark .sidebar-section li:hover .sidebar-icon {
  color: #6366f1;
}

/* 激活状态的图标颜色 */
.sidebar-section li.active .sidebar-icon {
  color: white;
}

/* 侧边栏文字 */
.sidebar-text {
  font-size: 0.875rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

/* ==================== 用户信息区域样式 ==================== */

/* 用户信息容器 */
.user-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem;
  background-color: #ffffff; /* 用户信息背景 */
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid #e5e7eb;
  margin-top: auto; /* 推到底部 */
}

/* 用户信息悬停效果 */
.user-info:hover {
  background-color: #f9fafb;
  border-color: #d1d5db;
  transform: translateY(-1px);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

/* 暗色主题下的用户信息 */
.sidebar.dark .user-info {
  background-color: #2d2d2d;
  border-color: #374151;
}

.sidebar.dark .user-info:hover {
  background-color: #374151;
  border-color: #4f46e5;
}

/* 用户头像 */
.user-avatar {
  width: 2.5rem;
  height: 2.5rem;
  background-color: #e5e7eb;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

/* 暗色主题下的用户头像 */
.sidebar.dark .user-avatar {
  background-color: #4b5563;
}

/* 用户图标 */
.user-icon {
  color: #6b7280;
}

/* 暗色主题下的用户图标 */
.sidebar.dark .user-icon {
  color: #9ca3af;
}

/* 用户详细信息 */
.user-details {
  display: flex;
  flex-direction: column;
  gap: 0.125rem;
  flex: 1;
}

/* 用户姓名 */
.user-name {
  font-size: 0.875rem;
  font-weight: 600;
  color: #1f2937;
  line-height: 1.2;
}

/* 暗色主题下的用户姓名 */
.sidebar.dark .user-name {
  color: #e5e7eb;
}

/* 用户角色 */
.user-role {
  font-size: 0.75rem;
  color: #6b7280;
  line-height: 1.2;
}

/* 暗色主题下的用户角色 */
.sidebar.dark .user-role {
  color: #d1d5db;
}

/* ==================== 响应式设计 ==================== */

/* 平板设备适配 */
@media (max-width: 768px) {
  .sidebar {
    width: 240px;
    padding: 1rem;
  }
  
  .sidebar-text {
    font-size: 0.8rem;
  }
  
  .user-info {
    padding: 0.75rem;
  }
  
  .user-name {
    font-size: 0.8rem;
  }
  
  .user-role {
    font-size: 0.7rem;
  }
}

/* 手机设备适配 */
@media (max-width: 480px) {
  .sidebar {
    width: 200px;
    padding: 0.75rem;
  }
  
  .sidebar-section li {
    padding: 0.5rem 0.75rem;
    gap: 0.5rem;
  }
  
  .sidebar-icon {
    width: 16px;
    height: 16px;
  }
  
  .sidebar-text {
    font-size: 0.75rem;
  }
  
  .user-avatar {
    width: 2rem;
    height: 2rem;
  }
  
  .user-icon {
    width: 16px;
    height: 16px;
  }
}
</style> 