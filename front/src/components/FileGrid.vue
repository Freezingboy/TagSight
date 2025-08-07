<template>
  <div class="file-grid" :class="{ dark: isDarkTheme }">
    <div class="file-card" @click="handleFileCardClick('mountain_view.jpg')">
      <div class="file-thumbnail">
        <img src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjQiIGhlaWdodD0iNjQiIHZpZXdCb3g9IjAgMCA2NCA2NCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHJlY3Qgd2lkdGg9IjY0IiBoZWlnaHQ9IjY0IiBmaWxsPSIjRjNGNEY2Ii8+CjxwYXRoIGQ9Ik0xNiAxNkg0OFY0OEgxNlYxNloiIGZpbGw9IiM5Q0EzQUYiLz4KPHBhdGggZD0iTTI0IDI0SDQwVjMySDI0VjI0WiIgZmlsbD0iI0U1RTdFQiIvPgo8L3N2Zz4K" alt="mountain_view.jpg" />
      </div>
      <div class="file-info">
        <h4>mountain_view.jpg</h4>
        <p>2.4 MB</p>
        <p>2024-03-15</p>
      </div>
    </div>

    <div class="file-card" @click="handleFileCardClick('project_report.pdf')">
      <div class="file-thumbnail pdf">
        <span>PDF</span>
      </div>
      <div class="file-info">
        <h4>project_report.pdf</h4>
        <p>5.7 MB</p>
        <p>2024-03-10</p>
      </div>
    </div>

    <div class="file-card" @click="handleFileCardClick('summer_vibes.mp3')">
      <div class="file-thumbnail audio">
        <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M9 18V5l12-2v13"/>
          <circle cx="6" cy="18" r="3"/>
          <circle cx="18" cy="16" r="3"/>
        </svg>
      </div>
      <div class="file-info">
        <h4>summer_vibes.mp3</h4>
        <p>8.2 MB</p>
        <p>2024-03-05</p>
      </div>
    </div>

    <div class="file-card" @click="handleFileCardClick('presentation.mp4')">
      <div class="file-thumbnail video">
        <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <polygon points="23,7 16,12 23,17 23,7"/>
          <rect x="1" y="5" width="15" height="14" rx="2" ry="2"/>
        </svg>
      </div>
      <div class="file-info">
        <h4>presentation.mp4</h4>
        <p>156 MB</p>
        <p>2024-02-28</p>
      </div>
    </div>

    <div class="file-card">
      <div class="file-thumbnail">
        <img src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjQiIGhlaWdodD0iNjQiIHZpZXdCb3g9IjAgMCA2NCA2NCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHJlY3Qgd2lkdGg9IjY0IiBoZWlnaHQ9IjY0IiBmaWxsPSIjRjNGNEY2Ii8+CjxwYXRoIGQ9Ik0xNiAxNkg0OFY0OEgxNlYxNloiIGZpbGw9IiM5Q0EzQUYiLz4KPHBhdGggZD0iTTI0IDI0SDQwVjMySDI0VjI0WiIgZmlsbD0iI0U1RTdFQiIvPgo8L3N2Zz4K" alt="city_night.jpg" />
      </div>
      <div class="file-info">
        <h4>city_night.jpg</h4>
        <p>3.1 MB</p>
        <p>2024-02-20</p>
      </div>
    </div>

    <div class="file-card">
      <div class="file-thumbnail excel">
        <span>X</span>
      </div>
      <div class="file-info">
        <h4>financial_data.xlsx</h4>
        <p>1.8 MB</p>
        <p>2024-02-15</p>
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

// 定义组件的事件接口
interface Emits {
  (e: 'file-click', fileName: string): void
}

// 定义组件的属性，设置默认值
const props = withDefaults(defineProps<Props>(), {
  isDarkTheme: false
})

const emit = defineEmits<Emits>()

const handleFileCardClick = (fileName: string) => {
  console.log('文件被点击:', fileName)
  emit('file-click', fileName)
  
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
</script>

<style scoped>
.file-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.file-card {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 1rem;
  cursor: pointer;
  transition: all 0.2s ease;
  position: relative;
  transform: translateY(0);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

/* 暗色主题下的文件卡片 */
.file-grid.dark .file-card {
  background-color: #2d2d2d;
  border-color: #404040;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
}

.file-card:hover {
  background-color: #f8fafc;
  border-color: #cbd5e1;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

/* 暗色主题下的文件卡片悬停效果 */
.file-grid.dark .file-card:hover {
  background-color: #404040;
  border-color: #525252;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4);
}

.file-card:active {
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
  transition: all 0.1s ease;
}

/* 暗色主题下的文件卡片激活效果 */
.file-grid.dark .file-card:active {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.5);
}

.file-thumbnail {
  width: 100%;
  height: 120px;
  background-color: #f1f5f9;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 0.75rem;
  overflow: hidden;
  transition: background-color 0.3s ease;
}

/* 暗色主题下的文件缩略图 */
.file-grid.dark .file-thumbnail {
  background-color: #404040;
}

.file-thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.file-thumbnail.pdf {
  background-color: #ef4444;
  color: white;
  font-weight: 600;
}

.file-thumbnail.audio {
  background-color: #8b5cf6;
  color: white;
}

.file-thumbnail.video {
  background-color: #3b82f6;
  color: white;
}

.file-thumbnail.excel {
  background-color: #10b981;
  color: white;
  font-weight: 600;
}

.file-info h4 {
  margin: 0 0 0.25rem 0;
  font-size: 0.875rem;
  font-weight: 500;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.file-info p {
  margin: 0;
  font-size: 0.75rem;
  color: #6b7280;
  transition: color 0.3s ease;
}

/* 暗色主题下的文件信息文字 */
.file-grid.dark .file-info p {
  color: #9ca3af;
}
</style> 