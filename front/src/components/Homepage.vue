<template>
  <!-- 首页组件 - 展示项目介绍和功能概览 -->
  <div class="homepage" :class="{ dark: isDarkTheme }">
    <!-- 英雄区域 - 项目标题和简介 -->
    <div class="hero-section">
      <div class="hero-content">
        <!-- 项目图标 -->
        <div class="hero-icon">
          <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.89l-.812-1.22A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/>
          </svg>
        </div>
        <!-- 项目标题 -->
        <h1>TagSight 前端项目</h1>
        <!-- 项目副标题 -->
        <p class="hero-subtitle">基于 Vue.js 的现代化文件管理系统</p>
      </div>
    </div>

    <!-- 内容网格 - 展示项目各个方面的信息卡片 -->
    <div class="content-grid">
      <!-- 项目概述卡片 -->
      <div class="content-card">
        <div class="card-header">
          <!-- 信息图标 -->
          <div class="card-icon">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="10"/>
              <path d="M12 16v-4"/>
              <path d="M12 8h.01"/>
            </svg>
          </div>
          <h2>项目概述</h2>
        </div>
        <div class="card-content">
          <p>Vue.js 文件管理系统，支持文件管理、标签管理、主题切换等功能。</p>
        </div>
      </div>

      <!-- 技术栈卡片 -->
      <div class="content-card">
        <div class="card-header">
          <!-- 技术栈图标 -->
          <div class="card-icon">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 2L2 7l10 5 10-5-10-5z"/>
              <path d="M2 17l10 5 10-5"/>
              <path d="M2 12l10 5 10-5"/>
            </svg>
          </div>
          <h2>技术栈</h2>
        </div>
        <div class="card-content">
          <!-- 技术标签列表 -->
          <div class="tech-list">
            <span class="tech-tag vue">Vue.js 3</span>
            <span class="tech-tag vite">Vite</span>
            <span class="tech-tag css">CSS3</span>
          </div>
        </div>
      </div>

      <!-- 主要功能卡片 -->
      <div class="content-card">
        <div class="card-header">
          <!-- 功能图标 -->
          <div class="card-icon">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
            </svg>
          </div>
          <h2>主要功能</h2>
        </div>
        <div class="card-content">
          <!-- 功能列表 -->
          <div class="features-list">
            <div class="feature-item">📁 文件管理 - 网格布局、搜索筛选</div>
            <div class="feature-item">🏷️ 标签管理 - 卡片布局、颜色标识</div>
            <div class="feature-item">🎨 主题切换 - 明暗双主题</div>
          </div>
        </div>
      </div>

      <!-- 组件结构卡片 -->
      <div class="content-card">
        <div class="card-header">
          <!-- 组件图标 -->
          <div class="card-icon">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
              <line x1="9" y1="9" x2="15" y2="9"/>
              <line x1="9" y1="12" x2="15" y2="12"/>
              <line x1="9" y1="15" x2="15" y2="15"/>
            </svg>
          </div>
          <h2>组件结构</h2>
        </div>
        <div class="card-content">
          <!-- 组件列表 -->
          <div class="components-list">
            <code>App.vue</code>
            <code>FileManager.vue</code>
            <code>Header.vue</code>
            <code>Sidebar.vue</code>
          </div>
        </div>
      </div>

      <!-- 开发说明卡片 -->
      <div class="content-card">
        <div class="card-header">
          <!-- 文档图标 -->
          <div class="card-icon">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
              <polyline points="14,2 14,8 20,8"/>
              <line x1="16" y1="13" x2="8" y2="13"/>
              <line x1="16" y1="17" x2="8" y2="17"/>
              <polyline points="10,9 9,9 8,9"/>
            </svg>
          </div>
          <h2>开发说明</h2>
        </div>
        <div class="card-content">
          <!-- 开发命令 -->
          <div class="dev-commands">
            <div class="command">
              <span class="prompt">$</span>
              <code>npm run dev</code>
            </div>
          </div>
        </div>
      </div>

      <!-- 更新日志卡片 -->
      <div class="content-card">
        <div class="card-header">
          <!-- 编辑图标 -->
          <div class="card-icon">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 20h9"/>
              <path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z"/>
            </svg>
          </div>
          <h2>更新日志</h2>
        </div>
        <div class="card-content">
          <!-- 更新记录列表 -->
          <div class="changelog-list">
            <div class="changelog-item">✅ 文件管理和标签管理功能</div>
            <div class="changelog-item">✅ 主题切换功能</div>
            <div class="changelog-item">✅ 组件化重构完成</div>
          </div>
        </div>
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

// Homepage组件 - 展示项目介绍和功能概览
// 这是一个静态展示页面，用于介绍项目的功能和技术栈
</script>

<style scoped>
/* 首页容器样式 */
.homepage {
  padding: 0.375rem;
  max-width: 1200px;
  margin: 0 auto;
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

/* 英雄区域样式 */
.hero-section {
  text-align: center;
  margin-bottom: 0.375rem;
  padding: 0.75rem;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-light) 100%);
  border-radius: 10px;
  color: white;
  box-shadow: 0 3px 12px rgba(79, 70, 229, 0.25);
}

.hero-content {
  max-width: 400px;
  margin: 0 auto;
}

.hero-icon {
  margin-bottom: 0.375rem;
  opacity: 0.9;
}

.hero-content h1 {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 0.125rem;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.hero-subtitle {
  font-size: 0.8rem;
  opacity: 0.9;
  margin: 0;
}

/* 内容网格布局 */
.content-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 0.375rem;
  flex: 1;
}

/* 内容卡片样式 */
.content-card {
  background: var(--bg-color);
  border-radius: 8px;
  padding: 0.25rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--border-color);
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
}

.content-card:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.12);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 0.375rem;
  margin-bottom: 0.125rem;
  justify-content: center;
  min-height: 1.5rem;
  padding: 0.25rem 0;
}

.card-icon {
  width: 20px;
  height: 20px;
  background: var(--primary-light);
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--primary-color);
  box-shadow: 0 1px 3px rgba(79, 70, 229, 0.15);
}

.card-header h2 {
  font-size: 1rem;
  font-weight: 600;
  margin: 0;
  color: #000000;
}

.card-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.content-card p {
  font-size: 0.8rem;
  color: var(--text-color);
  margin: 0;
  line-height: 1.3;
  font-weight: 500;
}

/* 技术标签列表样式 */
.tech-list {
  display: flex;
  flex-wrap: wrap;
  gap: 0.25rem;
  justify-content: center;
}

.tech-tag {
  padding: 0.125rem 0.375rem;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 600;
  color: white;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.tech-tag.vue {
  background: linear-gradient(135deg, #42b883, #3aa876);
}

.tech-tag.vite {
  background: linear-gradient(135deg, #646cff, #535bf2);
}

.tech-tag.css {
  background: linear-gradient(135deg, #1572b6, #0d5aa7);
}

/* 功能列表样式 */
.features-list {
  display: flex;
  flex-direction: column;
  gap: 0.125rem;
  width: 100%;
}

.feature-item {
  font-size: 0.8rem;
  color: var(--text-color);
  line-height: 1.2;
  font-weight: 500;
  padding: 0.125rem 0;
}

/* 组件列表样式 */
.components-list {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 0.125rem;
  width: 100%;
}

.components-list code {
  padding: 0.125rem 0.25rem;
  background: var(--code-bg);
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--primary-color);
  border: 1px solid var(--border-light);
  text-align: center;
  transition: all 0.2s ease;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.components-list code:hover {
  background: var(--primary-light);
  border-color: var(--primary-color);
  transform: translateY(-1px);
}

/* 开发命令样式 */
.dev-commands {
  display: flex;
  flex-direction: column;
  gap: 0.125rem;
  width: 100%;
}

.command {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.375rem;
  padding: 0.125rem 0.375rem;
  background: var(--code-bg);
  border-radius: 6px;
  font-family: 'Courier New', monospace;
  border: 1px solid var(--border-light);
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.prompt {
  color: var(--primary-color);
  font-weight: 700;
  font-size: 0.8rem;
}

.command code {
  background: none;
  padding: 0;
  color: var(--text-color);
  font-size: 0.8rem;
  font-weight: 600;
}

/* 更新日志样式 */
.changelog-list {
  display: flex;
  flex-direction: column;
  gap: 0.125rem;
  width: 100%;
}

.changelog-item {
  font-size: 0.8rem;
  color: var(--text-color);
  line-height: 1.2;
  font-weight: 500;
  padding: 0.125rem 0;
}

/* 亮色主题变量定义 */
.homepage {
  --bg-color: #ffffff;
  --text-color: #1a202c;
  --title-color: #000000;
  --text-secondary: #4a5568;
  --primary-color: #4f46e5;
  --primary-light: #eef2ff;
  --card-bg: #f8fafc;
  --border-color: #e2e8f0;
  --border-light: #f1f5f9;
  --code-bg: #f1f5f9;
  --code-color: #1e293b;
}

/* 暗色主题变量定义 */
.homepage.dark {
  --bg-color: #111827;
  --text-color: #e5e7eb;
  --title-color: #f9fafb;
  --text-secondary: #d1d5db;
  --primary-color: #4f46e5;
  --primary-light: #312e81;
  --card-bg: #1f2937;
  --border-color: #374151;
  --border-light: #374151;
  --code-bg: #2d2d2d;
  --code-color: #e5e7eb;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .homepage {
    padding: 0.5rem;
  }
  
  .hero-content h1 {
    font-size: 1.5rem;
  }
  
  .content-grid {
    grid-template-columns: 1fr;
  }
  
  .components-list {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .tech-list {
    justify-content: center;
  }
}

/* 暗色主题标题颜色覆盖 */
.homepage.dark .card-header h2 {
  color: #e5e7eb;
}
</style> 