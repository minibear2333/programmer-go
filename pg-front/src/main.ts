import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { registerApp } from './global'

const app = createApp(App)
// 注册ant-design
registerApp(app)

app.use(store)
app.use(router)

app.mount('#app')
