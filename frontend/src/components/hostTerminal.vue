<template>
    <div ref="xtermRef" class="xterm"></div>
</template>
  
<script>
import "xterm/css/xterm.css";
import { Terminal } from "xterm";
import { FitAddon } from 'xterm-addon-fit';
import { AttachAddon } from 'xterm-addon-attach';

export default {
    name: "HostTerminal",
    data: () => ({
        term: null,
        socket: null,
        socketURI: 'ws://127.0.0.1:18703/ws'
    }),
    methods: {
        init() {
            // 1.xterm终端初始化
            this.term = new Terminal()
            // 创建WebSocket连接
            this.socket = new WebSocket(this.socketURI)
            // 加载WebSocket插件
            this.term.loadAddon(new AttachAddon(this.socket))
            // 加载窗口自适应插件
            const fitAddon = new FitAddon()
            this.term.loadAddon(fitAddon)
            // 打开Dom元素
            this.term.open(this.$refs.xtermRef)
            // 自适应窗口大小
            fitAddon.fit()
            // 输入聚焦
            this.term.focus()

            this.term.onData(data => {
                console.log(data)
                this.term.write(data)
            })
        },
        close() {
            if (this.socket) this.socket.close()
            if (this.term) this.term.dispose()
        }
    }
};
</script>
<style scoped>
.xterm {
    height: calc(100vh - 50px);
}

.xterm:deep(.terminal) {
    height: calc(100vh - 50px);
}
</style>