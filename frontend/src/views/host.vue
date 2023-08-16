<template>
    <div>
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
            <div id="xterm" class="xterm"></div>
        </div>
        <div v-show="tab === 'file'">
            文件
        </div>
        <div v-if="tab === 'tool'">
            工具
        </div>
        <div v-else-if="tab === 'config'">
            配置
        </div>
    </div>
</template>
  
<script>
import "xterm/css/xterm.css";
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";

export default {
    name: "Host",
    data: () => ({
        tab: 'shell',
        term: null,
        prefix: "[root@colossus ~]# "
    }),
    methods: {
        updateTab(e){
            this.tab = e
        },
        initTerm() {
            let fitAddon = new FitAddon();
            this.term = new Terminal({
                fontSize: 14,
                cursorBlink: true,
                allowProposedApi: true,
                disableStdin: false,
                LogLevel: "debug"
            });
            this.term.loadAddon(fitAddon);

            this.term.open(document.getElementById("xterm"));
            this.term.writeln("\x1b[1;1;32mwellcom to colossus!\x1b[0m");
            this.term.write(this.prefix);

            setTimeout(() => {
                fitAddon.fit();
            }, 5)
        },
    },
    mounted() {
        setTimeout(() => {
            this.initTerm()
        }, 500);
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

#xterm {
    height: calc(100vh - 50px);
}

#xterm:deep(.terminal) {
    height: calc(100vh - 50px);
}
</style>