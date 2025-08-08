<template>
  <div class="login-container">
    <!-- 左侧品牌区域 -->
    <div class="brand-section">
      <div class="brand-content">
        <!-- Logo -->
        <div class="logo">
          <svg class="logo-icon" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"/>
            <line x1="7" y1="7" x2="7.01" y2="7"/>
          </svg>
          <span class="logo-text">TagSight</span>
        </div>
        
        <!-- 欢迎信息 -->
        <div class="welcome-content">
          <h1 class="welcome-title">欢迎回来</h1>
          <p class="welcome-subtitle">登录您的账户, 开始探索数据洞察</p>
        </div>
        
        <!-- 功能图标 -->
        <div class="feature-icons">
          <div class="icon-circle">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="22,12 18,12 15,21 9,3 6,12 2,12"/>
            </svg>
          </div>
          <div class="icon-circle">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
              <circle cx="8.5" cy="8.5" r="1.5"/>
              <polyline points="21,15 16,10 5,21"/>
            </svg>
          </div>
          <div class="icon-circle">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M18 10h-1.26A8 8 0 1 0 9 20h9a5 5 0 0 0 0-10z"/>
            </svg>
          </div>
        </div>
        
        <!-- 版权信息 -->
        <div class="copyright">
          © 2024 TagSight Analytics. All rights reserved.
        </div>
      </div>
    </div>
    
    <!-- 右侧登录表单区域 -->
    <div class="login-section">
      <div class="login-content">
        <h2 class="login-title">登录您的账户</h2>
        <p class="login-subtitle">请输入您的凭据以继续</p>
        
        <form @submit.prevent="handleLogin" class="login-form">
          <!-- 用户名输入 -->
          <div class="form-group">
            <label for="username" class="form-label">用户名</label>
            <input
              id="username"
              v-model="formData.username"
              type="text"
              class="form-input"
              placeholder="请输入您的用户名"
              required
            />
          </div>
          
          <!-- 密码输入 -->
          <div class="form-group">
            <label for="password" class="form-label">密码</label>
            <input
              id="password"
              v-model="formData.password"
              type="password"
              class="form-input"
              placeholder="请输入您的密码"
              required
            />
          </div>
          
          <!-- 记住我和忘记密码 -->
          <div class="form-options">
            <label class="checkbox-label">
              <input type="checkbox" v-model="formData.rememberMe" />
              <span class="checkmark"></span>
              记住我
            </label>
            <a href="#" class="forgot-password">忘记密码?</a>
          </div>
          
          <!-- 登录按钮 -->
          <button type="submit" class="login-button" :disabled="isLoading">
            {{ isLoading ? '登录中...' : '登录' }}
          </button>
          
          <!-- 错误信息 -->
          <div v-if="errorMessage" class="error-message">
            {{ errorMessage }}
          </div>
        </form>
        
        <!-- 分割线 -->
        <div class="divider">
          <span>或使用以下方式登录</span>
        </div>
        
        <!-- 社交登录 -->
        <div class="social-login">
          <button class="social-button">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
              <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
              <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
              <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
              <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
            </svg>
          </button>
          <button class="social-button">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </button>
          <button class="social-button">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
              <path d="M3 3h18v18H3V3zm16 16V5H5v14h14z"/>
              <path d="M7 7h4v4H7V7zm0 6h4v4H7v-4zm6-6h4v4h-4V7zm0 6h4v4h-4v-4z"/>
            </svg>
          </button>
        </div>
        
        <!-- 注册链接 -->
        <div class="register-link">
          <span>还没有账户?</span>
          <a href="#" class="register-text">立即注册</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

// 表单数据
const formData = reactive({
  username: '',
  password: '',
  rememberMe: false
})

// 状态管理
const isLoading = ref(false)
const errorMessage = ref('')

// 登录处理
const handleLogin = async () => {
  isLoading.value = true
  errorMessage.value = ''
  
  try {
    // 模拟登录验证
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // 检查用户名和密码
    if (formData.username === 'admin' && formData.password === '123456') {
      // 登录成功，存储登录状态
      localStorage.setItem('isLoggedIn', 'true')
      localStorage.setItem('username', formData.username)
      
      // 触发登录成功事件
      window.dispatchEvent(new CustomEvent('loginSuccess'))
    } else {
      errorMessage.value = '用户名或密码错误，请重试'
    }
  } catch (error) {
    errorMessage.value = '登录失败，请稍后重试'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  width: 100vw;
  height: 100vh;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  overflow: hidden;
}

/* 左侧品牌区域 */
.brand-section {
  flex: 1;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  position: relative;
  overflow: hidden;
}

.brand-content {
  text-align: center;
  max-width: 400px;
  padding: 2rem;
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 3rem;
}

.logo-icon {
  margin-right: 0.5rem;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 700;
}

.welcome-content {
  margin-bottom: 3rem;
}

.welcome-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
  line-height: 1.2;
}

.welcome-subtitle {
  font-size: 1.1rem;
  opacity: 0.9;
  line-height: 1.6;
}

.feature-icons {
  display: flex;
  justify-content: center;
  gap: 1.5rem;
  margin-bottom: 3rem;
}

.icon-circle {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(10px);
}

.copyright {
  position: absolute;
  bottom: 2rem;
  left: 50%;
  transform: translateX(-50%);
  font-size: 0.9rem;
  opacity: 0.8;
}

/* 右侧登录区域 */
.login-section {
  flex: 1;
  background: white;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
}

.login-content {
  max-width: 400px;
  width: 100%;
}

.login-title {
  font-size: 2rem;
  font-weight: 700;
  color: #1a1a1a;
  margin-bottom: 0.5rem;
  text-align: center;
}

.login-subtitle {
  color: #666;
  text-align: center;
  margin-bottom: 2rem;
}

.login-form {
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #333;
}

.form-input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 2px solid #e1e5e9;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-input:focus {
  outline: none;
  border-color: #667eea;
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.checkbox-label {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-size: 0.9rem;
  color: #666;
}

.checkbox-label input[type="checkbox"] {
  margin-right: 0.5rem;
}

.forgot-password {
  color: #667eea;
  text-decoration: none;
  font-size: 0.9rem;
}

.forgot-password:hover {
  text-decoration: underline;
}

.login-button {
  width: 100%;
  padding: 0.75rem;
  background: #667eea;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.login-button:hover:not(:disabled) {
  background: #5a6fd8;
}

.login-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.error-message {
  margin-top: 1rem;
  padding: 0.75rem;
  background: #fee;
  color: #c53030;
  border-radius: 6px;
  font-size: 0.9rem;
  text-align: center;
}

.divider {
  text-align: center;
  margin: 2rem 0;
  position: relative;
}

.divider::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: 1px;
  background: #e1e5e9;
}

.divider span {
  background: white;
  padding: 0 1rem;
  color: #666;
  font-size: 0.9rem;
}

.social-login {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 2rem;
}

.social-button {
  width: 50px;
  height: 50px;
  border: 2px solid #e1e5e9;
  border-radius: 50%;
  background: white;
  color: #666;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.social-button:hover {
  border-color: #667eea;
  color: #667eea;
}

.register-link {
  text-align: center;
  font-size: 0.9rem;
  color: #666;
}

.register-text {
  color: #667eea;
  text-decoration: none;
  margin-left: 0.5rem;
}

.register-text:hover {
  text-decoration: underline;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-container {
    flex-direction: column;
  }
  
  .brand-section {
    min-height: 40vh;
  }
  
  .login-section {
    min-height: 60vh;
  }
  
  .welcome-title {
    font-size: 2rem;
  }
  
  .login-title {
    font-size: 1.75rem;
  }
}
</style>
