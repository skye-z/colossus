import { createApp } from 'vue'
import router from './plugins/router'
// 非等宽字体
import 'vfonts/FiraSans.css'
import App from './App.vue'
import { GetOSName, SelectDirectory, SelectFile } from '../wailsjs/go/main/App'
import { ClipboardSetText } from '../wailsjs/runtime/runtime'

const app = createApp(App)

app.config.globalProperties.$goSetClipboard = ClipboardSetText
app.config.globalProperties.$goOSName = GetOSName
app.config.globalProperties.$goSelectDirectory = SelectDirectory
app.config.globalProperties.$goSelectFile = SelectFile

app.use(router)
app.mount('#app')
