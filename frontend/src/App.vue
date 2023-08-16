<template>
  <n-config-provider :theme="darkTheme" :locale="i18n.main" :date-locale="i18n.date" :theme-overrides="theme">
    <n-dialog-provider>
      <n-message-provider>
        <global-api />
        <div class="flex">
          <left-bar />
          <div id="main">
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
    }
  }),
  methods: {
    init(){
      this.$router.push('/home')
    }
  },
  mounted(){
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