// 导入Vite的配置定义函数
import { defineConfig } from 'vite'
// 导入Vue插件，用于处理.vue单文件组件
import vue from '@vitejs/plugin-vue'

// Vite配置文件
// 参考文档: https://vite.dev/config/
export default defineConfig({
  // 插件配置
  // vue() - 启用Vue单文件组件支持
  plugins: [vue()],
})
