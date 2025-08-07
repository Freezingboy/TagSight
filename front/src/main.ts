// 导入Vue 3的createApp函数 - 用于创建Vue应用实例
import { createApp } from 'vue'
// 导入全局样式文件
import './style.css'
// 导入根组件App.vue
import App from './App.vue'

// 创建Vue应用实例并挂载到DOM
// createApp(App) - 创建以App组件为根组件的Vue应用
// .mount('#app') - 将应用挂载到id为'app'的DOM元素上
createApp(App).mount('#app')
