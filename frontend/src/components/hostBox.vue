<template>
    <div>
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <div v-if="show">
            <div class="no-select" style="overflow: hidden;">
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
                    <div class="host-address text-small text-gray">{{ info.user }}@{{ info.address }}:{{ info.port }}</div>
                </div>
                <n-tabs v-model:value="tab" id="host-tabs" type="line" @update:value="updateTab" animated>
                    <n-tab name="shell">终端</n-tab>
                    <n-tab name="file">文件</n-tab>
                    <n-tab name="tool">工具</n-tab>
                </n-tabs>
            </div>
            <div v-show="tab === 'shell'">
                <host-terminal :ref="'terminal' + hostId" />
            </div>
            <div v-show="tab === 'file'">
                <host-file :ref="'file' + hostId" @send="sendTerminal" />
            </div>
            <div v-if="tab === 'tool'" id="tool-list" class="pa-10">
                <div class="tool-item pa-5 no-drag" @click="openTool('toolNetwork', info.id)">
                    <div>网络连接扫描</div>
                    <div class="text-small text-gray">检测主机开放端口与建立的连接</div>
                    <tool-network ref="toolNetwork" />
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import HostTerminal from "../components/hostTerminal.vue";
import HostFile from "../components/hostFile.vue";
import ToolNetwork from "../tools/network.vue";
import { host } from "../plugins/api";

export default {
    name: "HostBox",
    components: { HostTerminal, HostFile, ToolNetwork },
    data: () => ({
        info: {},
        tab: 'shell',
        show: false,
        loading: true,
        lastUpdate: 0
    }),
    props: {
        hostId: {
            type: String,
            default: ""
        }
    },
    beforeDestroy() {
        if (this.$refs["terminal" + this.hostId])
            this.$refs["terminal" + this.hostId].close()
    },
    methods: {
        init() {
            this.show = true
            this.getHostInfo()
        },
        getHostInfo() {
            host.getInfo(this.hostId).then(res => {
                this.loading = false
                if (res.state) {
                    this.$refs["terminal" + this.hostId].init(this.hostId, res.data.platform, res.data.system)
                    this.lastUpdate = new Date().getTime()
                    this.info = res.data
                } else {
                    this.$refs["terminal" + this.hostId].init(this.hostId, null, null)
                    this.initError()
                }
            }).catch(err => {
                console.log(err)
                this.initError()
            })
        },
        initError() {
            window.$dialog.error({
                closable: false,
                closeOnEsc: false,
                maskClosable: false,
                title: "无效主机",
                content: "主机参数配置错误或主机不存在",
                positiveText: "返回",
                onPositiveClick: () => {
                    this.$router.back()
                }
            });
        },
        sendTerminal(e) {
            this.$refs["terminal" + this.hostId].send('cd ' + e)
            this.tab = 'shell'
        },
        updateTab(e) {
            if (e == 'file') this.$refs["file" + this.hostId].init(this.hostId)
        },
        openTool(tag, param) {
            this.$refs[tag].open(param)
        }
    },
    mounted() {
        setTimeout(() => {
            this.init()
        }, 300)
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
    margin-left: calc(100% - 180px);
    --wails-draggable: no-drag;
    padding: 3.5px 0;
}

.tips {
    padding-top: 40vh;
}

#tool-list {
    flex-wrap: wrap;
    display: flex;
}

.tool-item {
    background-color: rgba(255, 255, 255, .08);
    margin: 0 10px 10px 0;
    border-radius: 8px;
    position: relative;
    min-width: 180px;
    cursor: pointer;
}


.tool-item:hover {
    background-color: rgba(255, 255, 255, .12);
}
</style>