<template>
    <div>
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <div v-if="show">
            <div class="no-select">
                <div id="host-head" class="border-bottom">
                    <div class="flex align-center">
                        <div class="host-name mr-10">{{ info.name }}</div>
                        <n-tag size="small" class="mr-5" :bordered="false" type="info" v-if="info.platform">{{ info.platform
                        }}</n-tag>
                        <n-tag size="small" class="mr-5" :bordered="false" type="success" v-if="info.system">{{ info.system
                        }}</n-tag>
                        <n-tag size="small" class="mr-5" :bordered="false" type="warning" v-if="info.region">{{ info.region
                        }}</n-tag>
                        <n-tag size="small" :bordered="false" type="error" v-if="info.periodTxt">{{ info.periodTxt
                        }}</n-tag>
                    </div>
                    <div class="host-address text-small text-gray">{{ info.address }}:{{ info.port }}</div>
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
                <host-file ref="file" />
            </div>
            <div v-if="tab === 'tool'">
                <div class="tips text-center">
                    <div>等待开发</div>
                    <div class="text-small text-gray">配置检测、链路跟踪、性能测试、开放端口、路由表、笔记、流量监控、应用可视化安装器、Docker容器管理</div>
                </div>
            </div>
            <div v-else-if="tab === 'config'">
                <div class="tips text-center">
                    <div>等待开发</div>
                    <div class="text-small text-gray">配置别名、终端颜色、压缩上传、压缩下载、探针、有效期限、认证信息</div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import HostTerminal from "../components/hostTerminal.vue";
import HostFile from "../components/hostFile.vue";

export default {
    name: "Host",
    components: { HostTerminal, HostFile },
    data: () => ({
        info: {},
        tab: 'shell',
        show: false,
        loading: true
    }),
    beforeDestroy() {
        if (this.$refs.terminal)
            this.$refs.terminal.close()
    },
    methods: {
        init() {
            this.saveConnectLog()
            setTimeout(() => {
                this.show = true
            }, 100)
            setTimeout(() => {
                this.$refs.terminal.init(this.info.id)
                this.loading = false
            }, 300)
        },
        saveConnectLog() {

        },
        updateTab(e) {
            this.tab = e
            if (e == 'file') this.$refs.file.init(this.info.id)
        }
    },
    mounted() {
        if (this.$route.query.info) {
            this.info = JSON.parse(this.$route.query.info)
            this.init()
        } else {
            window.$dialog.error({
                closable: false,
                closeOnEsc: false,
                maskClosable: false,
                title: "无效主机",
                content: "入口错误, 未传入有效的主机数据,",
                positiveText: "返回",
                onPositiveClick: () => {
                    this.$router.back()
                }
            });
        }
    }
};
</script>
<style scoped>
#host-head {
    width: calc(100% - 201px);
    position: absolute;
    padding: 6px 10px;
}

.host-name {
    line-height: 16px;
    font-size: 18px;
}

.host-address {
    line-height: 14px;
    font-size: 12px;
}

#host-tabs:deep(.n-tabs-nav-scroll-content) {
    margin-left: calc(100% - 250px);
    --wails-draggable: no-drag;
    padding: 3.5px 0;
}

.tips {
    padding-top: 40vh;
}
</style>