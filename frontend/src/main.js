import {createApp} from 'vue'
import router from './plugins/router'
// 非等宽字体
import 'vfonts/FiraSans.css'
import App from './App.vue'

createApp(App).use(router).mount('#app')
