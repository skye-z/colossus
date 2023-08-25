<template>
    <div ref="xtermRef" class="xterm"></div>
</template>
  
<script>
import "xterm/css/xterm.css";
import { Terminal } from "xterm";
import { FitAddon } from 'xterm-addon-fit';
import { AttachAddon } from 'xterm-addon-attach';
import { CanvasAddon } from 'xterm-addon-canvas';
import { WebglAddon } from 'xterm-addon-webgl';

export default {
    name: "HostTerminal",
    data: () => ({
        term: null,
        socket: null,
        socketURI: 'ws://127.0.0.1:18703/ws'
    }),
    methods: {
        init(hostId) {
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
            this.socket = new WebSocket(this.socketURI + '?cols=' + this.term.cols + '&rows' + this.term.rows + '&id=' + hostId)
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
            // 关闭连接
            if (this.socket) this.socket.close()
            // 销毁终端
            if (this.term) this.term.dispose()
            console.log('Terminal Close')
        },
    }
};
</script>
<style scoped>
.xterm {
    height: calc(100vh - 50px);
}

.xterm:deep(.terminal) {
    height: calc(100vh - 60px);
    padding: 5px;
}
</style>