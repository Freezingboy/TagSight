<template>
  <div class="ai-assistant-page" :class="{ dark: isDarkTheme }">
    <!-- 对话区域 -->
    <div class="chat-container">
      <div class="messages-container" ref="messagesContainer">
        <div 
          v-for="message in messages" 
          :key="message.id" 
          class="message-wrapper"
          :class="{ 'user-message': message.type === 'user', 'ai-message': message.type === 'ai' }"
        >
          <div class="message">
            <div class="message-avatar">
              <div v-if="message.type === 'user'" class="user-avatar">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                  <circle cx="12" cy="7" r="4"/>
                </svg>
              </div>
              <div v-else class="ai-avatar">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/>
                </svg>
              </div>
            </div>
            <div class="message-content">
              <div class="message-header">
                <span class="sender-name">{{ message.type === 'user' ? '您' : 'TagSight AI' }}</span>
                <span class="message-time">{{ message.time }}</span>
              </div>
              <div class="message-text" v-html="message.content"></div>
              <div v-if="message.type === 'ai' && message.isTyping" class="typing-indicator">
                <div class="typing-dots">
                  <div class="dot"></div>
                  <div class="dot"></div>
                  <div class="dot"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 输入区域 -->
      <div class="input-container">
        <div class="input-wrapper">
          <div class="input-field">
            <textarea 
              v-model="inputMessage" 
              @keydown.enter.prevent="sendMessage"
              placeholder="输入您的问题..."
              class="message-input"
              rows="1"
              ref="messageInput"
            ></textarea>
            <div class="input-actions">
              <button 
                @click="sendMessage" 
                class="send-btn"
                :disabled="!inputMessage.trim() || isSending"
              >
                <span v-if="!isSending">发送</span>
                <div v-else class="sending-spinner"></div>
              </button>
            </div>
          </div>
          <div class="input-tips">
            <span>按 Enter 发送，Shift + Enter 换行</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick, watch } from 'vue'

interface Message {
  id: number
  type: 'user' | 'ai'
  content: string
  time: string
  isTyping?: boolean
}

interface Props {
  isDarkTheme?: boolean
}

interface Emits {
  (e: 'back'): void
}

const props = withDefaults(defineProps<Props>(), {
  isDarkTheme: false
})

const emit = defineEmits<Emits>()

const messages = ref<Message[]>([
  {
    id: 1,
    type: 'ai',
    content: '您好！我是 TagSight AI 助手，很高兴为您服务。我可以帮助您：<br>• 文件管理和分类<br>• 标签创建和管理<br>• 智能搜索和推荐<br>• 数据分析和建议<br><br>请告诉我您需要什么帮助？',
    time: '14:30'
  },
  {
    id: 2,
    type: 'user',
    content: '我想了解一下如何更好地管理我的文件标签',
    time: '14:31'
  },
  {
    id: 3,
    type: 'ai',
    content: '很好的问题！以下是一些文件标签管理的最佳实践：<br><br><strong>1. 创建层次化标签</strong><br>• 使用主标签（如"工作"、"个人"）<br>• 添加子标签（如"工作/项目A"、"个人/照片"）<br><br><strong>2. 使用描述性标签</strong><br>• 避免过于宽泛的标签<br>• 使用具体、有意义的名称<br><br><strong>3. 保持一致性</strong><br>• 建立标签命名规范<br>• 定期清理未使用的标签<br><br>您想了解哪个方面的具体操作？',
    time: '14:32'
  },
  {
    id: 4,
    type: 'user',
    content: '能帮我分析一下我最近上传的文件吗？',
    time: '14:33'
  },
  {
    id: 5,
    type: 'ai',
    content: '当然可以！根据您最近的文件活动，我为您分析如下：<br><br><strong>📊 文件统计</strong><br>• 总文件数：156 个<br>• 本月新增：23 个<br>• 主要类型：图片(45%)、文档(30%)、视频(15%)、其他(10%)<br><br><strong>🏷️ 标签使用情况</strong><br>• 最常用标签：工作(28%)、个人(25%)、项目(18%)<br>• 建议：考虑为"项目"标签添加更细分的子标签<br><br><strong>💡 优化建议</strong><br>• 为视频文件添加更多描述性标签<br>• 考虑创建"重要"标签来标记关键文件<br>• 定期归档旧文件以保持整洁<br><br>需要我帮您实施这些建议吗？',
    time: '14:34'
  }
])

const inputMessage = ref('')
const isSending = ref(false)
const messagesContainer = ref<HTMLElement>()
const messageInput = ref<HTMLTextAreaElement>()

// 自动调整输入框高度
const adjustTextareaHeight = () => {
  if (messageInput.value) {
    messageInput.value.style.height = 'auto'
    const lineHeight = 24 // 每行高度
    const maxLines = 3
    const maxHeight = lineHeight * maxLines
    const newHeight = Math.min(messageInput.value.scrollHeight, maxHeight)
    messageInput.value.style.height = Math.max(lineHeight, newHeight) + 'px'
  }
}

// 滚动到底部
const scrollToBottom = () => {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
    }
  })
}

// 发送消息
const sendMessage = async () => {
  if (!inputMessage.value.trim() || isSending.value) return

  const userMessage: Message = {
    id: Date.now(),
    type: 'user',
    content: inputMessage.value.trim(),
    time: new Date().toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
  }

  messages.value.push(userMessage)
  const userInput = inputMessage.value.trim()
  inputMessage.value = ''
  isSending.value = true

  // 模拟AI回复
  setTimeout(() => {
    const aiMessage: Message = {
      id: Date.now() + 1,
      type: 'ai',
      content: '',
      time: new Date().toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' }),
      isTyping: true
    }
    messages.value.push(aiMessage)
    scrollToBottom()

    // 模拟打字效果
    const responses = [
      `我理解您的问题。关于"${userInput}"，我可以为您提供以下建议：<br><br>• 首先，建议您检查相关的设置选项<br>• 其次，可以尝试重新整理文件结构<br>• 最后，如果问题持续存在，可以联系技术支持<br><br>您觉得这些建议如何？`,
      `很好的问题！基于您提到的"${userInput}"，我建议：<br><br><strong>1. 系统优化</strong><br>定期清理缓存和临时文件<br><br><strong>2. 功能使用</strong><br>充分利用TagSight的智能分类功能<br><br><strong>3. 最佳实践</strong><br>建立统一的文件命名规范<br><br>需要我详细解释任何一点吗？`,
      `关于"${userInput}"，我为您整理了以下解决方案：<br><br>📋 <strong>步骤指南</strong><br>1. 打开相关设置面板<br>2. 检查配置选项<br>3. 应用更改并测试<br><br>🎯 <strong>预期结果</strong><br>操作完成后，您应该能看到明显的改善<br><br>如果还有疑问，请随时告诉我！`
    ]

    const randomResponse = responses[Math.floor(Math.random() * responses.length)]
    
    setTimeout(() => {
      aiMessage.content = randomResponse
      aiMessage.isTyping = false
      isSending.value = false
      scrollToBottom()
    }, 1500)
  }, 1000)
}

// 监听输入变化
watch(inputMessage, adjustTextareaHeight)

onMounted(() => {
  scrollToBottom()
})
</script>

<style scoped>
.ai-assistant-page {
  height: 100%;
  display: flex;
  flex-direction: column;
  background: #ffffff;
  transition: all 0.3s ease;
}

.ai-assistant-page.dark {
  background: #0f172a;
}



/* 对话区域 */
.chat-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  padding: 0 1.5rem;
  position: relative;
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 1rem 0;
  scroll-behavior: smooth;
  padding-bottom: 140px;
}

.messages-container::-webkit-scrollbar {
  width: 6px;
}

.messages-container::-webkit-scrollbar-track {
  background: transparent;
}

.messages-container::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.1);
  border-radius: 3px;
}

.ai-assistant-page.dark .messages-container::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.1);
}

.message-wrapper {
  margin-bottom: 1.25rem;
}

.message {
  display: flex;
  gap: 0.75rem;
  max-width: 60%;
}

.user-message .message {
  margin-left: auto;
  flex-direction: row-reverse;
}

/* 头像设计 */
.message-avatar {
  flex-shrink: 0;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  position: relative;
  overflow: hidden;
  border: 3px solid rgba(255, 255, 255, 0.3);
  transition: all 0.3s ease;
}

.user-avatar {
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  box-shadow: 0 6px 20px rgba(99, 102, 241, 0.4);
  border-color: rgba(99, 102, 241, 0.3);
}

.ai-avatar {
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%);
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.4);
  border-color: rgba(59, 130, 246, 0.3);
}

.message-avatar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.2) 50%, transparent 70%);
  animation: shimmer 4s infinite;
}

.message-avatar::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 70%;
  height: 70%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.15) 0%, transparent 70%);
  transform: translate(-50%, -50%);
  border-radius: 50%;
}

.message-avatar:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.message-content {
  flex: 1;
  min-width: 0;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.sender-name {
  font-weight: 700;
  font-size: 0.9rem;
  color: #374151;
  letter-spacing: 0.025em;
}

.ai-assistant-page.dark .sender-name {
  color: #f1f5f9;
}

.ai-message .sender-name {
  color: #3b82f6;
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.ai-assistant-page.dark .ai-message .sender-name {
  color: #60a5fa;
  background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.message-time {
  font-size: 0.75rem;
  color: #6b7280;
  font-weight: 500;
}

.ai-assistant-page.dark .message-time {
  color: #9ca3af;
}

/* 消息气泡 */
.message-text {
  background: white;
  padding: 1rem 1.25rem;
  border-radius: 18px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
  line-height: 1.6;
  color: #374151;
  border: 1px solid rgba(0, 0, 0, 0.06);
  position: relative;
  transition: all 0.3s ease;
}

.ai-assistant-page.dark .message-text {
  background: #1e293b;
  color: #f1f5f9;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.25);
  border-color: rgba(255, 255, 255, 0.12);
}

.user-message .message-text {
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  color: white;
  border: none;
  box-shadow: 0 6px 20px rgba(99, 102, 241, 0.35);
}

.ai-message .message-text {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border: 1px solid rgba(59, 130, 246, 0.1);
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.08);
}

.ai-assistant-page.dark .ai-message .message-text {
  background: linear-gradient(135deg, #1e293b 0%, #334155 100%);
  border-color: rgba(59, 130, 246, 0.2);
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.15);
}

.message-text:hover {
  transform: translateY(-1px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.12);
}

.ai-assistant-page.dark .message-text:hover {
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
}

.user-message .message-text:hover {
  box-shadow: 0 8px 25px rgba(99, 102, 241, 0.4);
}

.ai-message .message-text:hover {
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.12);
}

.ai-assistant-page.dark .ai-message .message-text:hover {
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.2);
}

/* 打字指示器 */
.typing-indicator {
  margin-top: 0.75rem;
}

.typing-dots {
  display: flex;
  gap: 6px;
  padding: 1rem 1.25rem;
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.1) 0%, rgba(29, 78, 216, 0.1) 100%);
  border-radius: 16px;
  width: fit-content;
  border: 1px solid rgba(59, 130, 246, 0.15);
}

.ai-assistant-page.dark .typing-dots {
  background: linear-gradient(135deg, rgba(59, 130, 246, 0.15) 0%, rgba(29, 78, 216, 0.15) 100%);
  border-color: rgba(59, 130, 246, 0.25);
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%);
  animation: typing 1.4s infinite ease-in-out;
  box-shadow: 0 2px 4px rgba(59, 130, 246, 0.3);
}

.dot:nth-child(1) { animation-delay: -0.32s; }
.dot:nth-child(2) { animation-delay: -0.16s; }

@keyframes typing {
  0%, 80%, 100% { transform: scale(0.8); opacity: 0.5; }
  40% { transform: scale(1); opacity: 1; }
}

/* 输入区域 */
.input-container {
  position: absolute;
  bottom: 0;
  left: 1.5rem;
  right: 1.5rem;
  background: #ffffff;
  border-top: 1px solid #e2e8f0;
  padding: 1.5rem;
  z-index: 10;
}

.ai-assistant-page.dark .input-container {
  background: #1e293b;
  border-top-color: #334155;
}

.input-wrapper {
  max-width: 800px;
  margin: 0 auto;
}

.input-field {
  display: flex;
  gap: 0.75rem;
  align-items: flex-end;
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  padding: 1rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.ai-assistant-page.dark .input-field {
  background: #1e293b;
  border-color: #475569;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.input-field:focus-within {
  border-color: #6366f1;
  box-shadow: 0 4px 20px rgba(99, 102, 241, 0.12);
  transform: translateY(-1px);
}

.ai-assistant-page.dark .input-field:focus-within {
  border-color: #8b5cf6;
  box-shadow: 0 4px 20px rgba(139, 92, 246, 0.12);
}

.message-input {
  flex: 1;
  border: none;
  outline: none;
  resize: none;
  font-family: inherit;
  font-size: 0.875rem;
  line-height: 1.5;
  color: #374151;
  background: transparent;
  max-height: 120px;
  padding: 0;
}

.ai-assistant-page.dark .message-input {
  color: #f1f5f9;
}

.message-input::placeholder {
  color: #9ca3af;
}

.ai-assistant-page.dark .message-input::placeholder {
  color: #64748b;
}

.input-actions {
  display: flex;
  align-items: center;
}

.send-btn {
  width: 80px;
  height: 42px;
  border: none;
  border-radius: 14px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  flex-shrink: 0;
  position: relative;
  overflow: hidden;
  font-size: 0.875rem;
  font-weight: 500;
}

.send-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.send-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(99, 102, 241, 0.4);
}

.send-btn:hover:not(:disabled)::before {
  left: 100%;
}

.send-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

.sending-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid transparent;
  border-top: 2px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.input-tips {
  text-align: center;
  font-size: 0.75rem;
  color: #6b7280;
  margin-top: 0.75rem;
}

.ai-assistant-page.dark .input-tips {
  color: #9ca3af;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .chat-container {
    padding: 0 1rem;
  }
  
  .message {
    max-width: 75%;
  }
  
  .input-container {
    left: 1rem;
    right: 1rem;
    padding: 1rem;
  }
  
  .input-field {
    padding: 0.75rem;
  }
  
  .send-btn {
    width: 70px;
    height: 38px;
  }
}
</style>
