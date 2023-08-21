<template>
    <div>
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <div v-if="show">
            <div class="no-select">
                <div id="host-head" class="pa-10 flex align-center border-bottom">
                    <div class="host-name mr-10">测试服务器(广州)</div>
                    <n-tag size="small" :bordered="false" type="info">Linux</n-tag>
                    <n-tag size="small" :bordered="false" type="success">CentOS</n-tag>
                    <n-tag size="small" :bordered="false" type="warning">亚洲</n-tag>
                    <n-tag size="small" :bordered="false" type="error">3天后过期</n-tag>
                </div>
                <n-tabs id="host-tabs" type="line" @update:value="updateTab" animated>
                    <n-tab name="shell">终端</n-tab>
                    <n-tab name="file">文件</n-tab>
                    <n-tab name="tool">工具</n-tab>
                    <n-tab name="config">配置</n-tab>
                </n-tabs>
            </div>
            <div v-show="tab === 'shell'">
                <host-terminal ref="terminal" />
            </div>
            <div v-show="tab === 'file'">
                <div class="pa-10">
                    文件浏览、上传下载
                </div>
            </div>
            <div v-if="tab === 'tool'">
                <div class="pa-10">
                    配置检测、链路跟踪、性能测试、开放端口、路由表、笔记、流量监控、应用可视化安装器、Docker容器管理
                </div>
            </div>
            <div v-else-if="tab === 'config'">
                <div class="pa-10">
                    配置别名、终端颜色、压缩上传、压缩下载、探针、有效期限、认证信息
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import "xterm/css/xterm.css";
import HostTerminal from "../components/hostTerminal.vue";

export default {
    name: "Host",
    components: { HostTerminal },
    data: () => ({
        tab: 'shell',
        show: false,
        loading: true
    }),
    beforeDestroy() {
        if (this.$refs.terminal)
            this.$refs.terminal.close()
    },
    methods: {
        updateTab(e) {
            this.tab = e
        },
        initTerminal() {
            setTimeout(() => {
                this.show = true
            }, 100)
            setTimeout(() => {
                this.$refs.terminal.init()
                this.loading = false
            }, 300)
        }
    },
    mounted() {
        this.initTerminal()
    }
};
</script>
<style scoped>

#host-head {
    width: calc(100% - 201px);
    position: absolute;
}

.host-name {
    font-size: 18px;
}

#host-tabs:deep(.n-tabs-nav-scroll-content) {
    margin-left: calc(100% - 250px);
    --wails-draggable: no-drag;
    padding: 3.5px 0;
}
</style>