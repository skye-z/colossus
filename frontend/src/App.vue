<template>
  <n-config-provider :theme="darkTheme" :locale="i18n.main" :date-locale="i18n.date" :theme-overrides="theme">
    <n-dialog-provider>
      <n-message-provider>
        <global-api />
        <div class="flex">
          <left-bar :macOS="isMacOS" />
          <div id="main">
            <div class="loading flex align-center justify-center" v-if="loading">
              <n-spin size="small" />
            </div>
            <!-- <router-view /> -->
            <router-view v-slot="{ Component }">
              <keep-alive>
                <component :is="Component" />
              </keep-alive>
            </router-view>
          </div>
        </div>
      </n-message-provider>
    </n-dialog-provider>
  </n-config-provider>
</template>

<script>
import { darkTheme, zhCN, dateZhCN } from 'naive-ui'
import GlobalApi from './components/globalApi.vue'
import LeftBar from './components/leftBar.vue'
import { config } from './plugins/api'
import theme from './theme.json'

export default {
  name: "App",
  components: { GlobalApi, LeftBar },
  data: () => ({
    theme,
    darkTheme,
    i18n: {
      main: zhCN,
      date: dateZhCN
    },
    loading: true,
    isMacOS: false
  }),
  methods: {
    init() {
      this.initSystem();
      this.initConfig();
      setTimeout(() => {
        this.$router.push('/home')
      }, 100)
      setTimeout(() => {
        this.loading = false
      }, 300)
    },
    initSystem() {
      try {
        this.$goOSName().then(res => {
          this.isMacOS = res == 'darwin'
        });
      } catch (err) {
        this.isMacOS = true
        window.$message.warning('获取当前运行环境失败 ( 默认 darwin )')
      }
    },
    initConfig() {
      config.getAll().then(res => {
        if (res.state) {
          for (let key in res.data) {
            localStorage.setItem(key, res.data[key])
          }
        } else window.$message.warning('加载应用设置失败')
      }).catch(() => {
        window.$message.warning('加载应用设置失败')
      })
    }
  },
  mounted() {
    this.init()
  }
};
</script>
<style>
#main {
  width: calc(100vw - 181px);
  height: 100vh;
}
</style>