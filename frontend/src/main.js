import {createApp} from 'vue'
import router from './plugins/router'
// 非等宽字体
import 'vfonts/FiraSans.css'
import App from './App.vue'
import { GetOSName } from '../wailsjs/go/main/App'

const app = createApp(App)

app.config.globalProperties.$goOSName = GetOSName

app.use(router)
app.mount('#app')
