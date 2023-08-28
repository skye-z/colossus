<template>
    <div>
        <div ref="xtermRef" class="xterm border-bottom"></div>
        <div class="pa-10 flex align-center">
            <n-input v-model:value="fastCode" type="text" placeholder="查找快捷命令..." clearable>
                <template #prefix>
                    <n-icon>
                        <Code16Filled />
                    </n-icon>
                </template>
            </n-input>
            <n-button strong secondary v-if="connect" type="error" class="ml-10" @click="close">断开连接</n-button>
            <n-button strong secondary v-else type="primary" class="ml-10" @click="init()">重新连接</n-button>
        </div>
    </div>
</template>
  
<script>
import "xterm/css/xterm.css";
import { Code16Filled } from '@vicons/fluent'
import { Terminal } from "xterm";
import { FitAddon } from 'xterm-addon-fit';
import { AttachAddon } from 'xterm-addon-attach';
import { CanvasAddon } from 'xterm-addon-canvas';
import { WebglAddon } from 'xterm-addon-webgl';

export default {
    name: "HostTerminal",
    components: { Code16Filled },
    data: () => ({
        id: 0,
        term: null,
        socket: null,
        socketURI: 'ws://127.0.0.1:18703/ws',
        fastCode: '',
        connect: true
    }),
    methods: {
        init(hostId) {
            this.connect = true
            if(hostId) this.id = hostId
            this.term = new Terminal()
            // 加载Canvas渲染
            this.term.loadAddon(new CanvasAddon())
            // 加载WebGL渲染
            this.term.loadAddon(new WebglAddon())
            // 加载窗口自适应插件
            const fitAddon = new FitAddon()
            this.term.loadAddon(fitAddon)
            // 打开Dom元素
            this.term.open(this.$refs.xtermRef)
            // 自适应窗口大小
            fitAddon.fit()
            // 创建WebSocket连接
            this.socket = new WebSocket(this.socketURI + '?cols=' + this.term.cols + '&rows' + this.term.rows + '&id=' + this.id)
            this.socket.onopen = () => {
                window.dispatchEvent(new CustomEvent("cache:connect", { detail: { id: this.id, connect: true } }))
            };
            this.socket.onclose = () => {
                window.dispatchEvent(new CustomEvent("cache:connect", { detail: { id: this.id, connect: false } }))
            };
            this.socket.onmessage = e => {
                console.log(e)
            };

            // 加载WebSocket插件
            this.term.loadAddon(new AttachAddon(this.socket))
            // 输入聚焦
            this.term.focus()

            let timeout = 0
            window.addEventListener('resize', () => {
                fitAddon.fit();
                clearTimeout(timeout)
                timeout = setTimeout(() => {
                    this.socket.send("!~" + this.term.cols + ":" + this.term.rows)
                }, 500)
            });
        },
        close() {
            this.connect = false
            // 关闭连接
            if (this.socket) this.socket.close()
            // 销毁终端
            if (this.term) this.term.dispose()
            window.dispatchEvent(new CustomEvent("cache:connect", { detail: { id: hostId, connect: false } }))
            console.log('Terminal Close')
        },
    }
};
</script>
<style scoped>
.xterm {
    height: calc(100vh - 107px);
}

.xterm:deep(.terminal) {
    height: calc(100vh - 110px);
    padding: 5px;
}
</style>