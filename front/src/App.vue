<script setup lang="ts">
// 导入主要的文件管理组件
// FileManager是整个应用的核心组件，包含了所有功能模块
import FileManager from './components/FileManager.vue'
</script>

<template>
  <!-- 根组件模板 -->
  <!-- 整个应用只包含一个FileManager组件，它是所有功能的入口点 -->
  <FileManager />
</template>

<style>
/* 全局样式设置 */
/* 确保应用占满整个视口 */
html, body, #app {
  width: 100vw;  /* 视口宽度 */
  height: 100vh; /* 视口高度 */
  margin: 0;     /* 移除默认边距 */
  padding: 0;    /* 移除默认内边距 */
  background: #1a1a1a; /* 深色背景 */
}

/* 全局盒模型设置 */
* {
  box-sizing: border-box; /* 使用border-box盒模型，包含padding和border */
}
</style>
