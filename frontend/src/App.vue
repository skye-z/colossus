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
            <router-view />
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
      this.$goOSName().then(res => {
        this.isMacOS = res == 'darwin'
      });
      setTimeout(() => {
        this.$router.push('/home')
      }, 100)
      setTimeout(() => {
        this.loading = false
      }, 300)
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