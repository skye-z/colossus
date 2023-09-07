<template>
    <div style="position: relative;">
        <div :ref="'xterm' + id" :id="'xterm' + id" class="xterm border-bottom"></div>
        <div v-if="!connect" class="notconnect no-select flex align-center justify-center">
            <div>未连接</div>
        </div>
        <div class="no-select pa-10 flex align-center">
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
// import { WebglAddon } from 'xterm-addon-webgl';

export default {
    name: "HostTerminal",
    components: { Code16Filled },
    data: () => ({
        id: 0,
        term: null,
        socket: null,
        fitAddon: null,
        socketURI: 'ws://127.0.0.1:18703/ws',
        fastCode: '',
        connect: true
    }),
    methods: {
        init(hostId) {
            this.connect = true
            if (hostId) this.id = hostId
            try {
                let background = localStorage.getItem('terminal.background_color')
                let foreground = localStorage.getItem('terminal.text_color')
                let fontSize = parseInt(localStorage.getItem('terminal.text_size'))
                this.term = new Terminal({
                    theme: {
                        background,
                        foreground
                    },
                    fontSize
                })
                // 加载插件
                this.addPlugins();
                // 打开Dom元素
                setTimeout(()=>{
                    this.term.open(this.$refs['xterm' + hostId])
                },100)
                // 自适应窗口大小
                this.fitAddon.fit()
                // 创建连接
                this.addSocket()
                // 输入聚焦
                this.term.focus()
                // 加载大小变动事件
                this.addResizeEvent();
            } catch (err) {
                console.log(err)
            }
        },
        addPlugins() {
            // 加载Canvas渲染
            this.term.loadAddon(new CanvasAddon())
            // 加载WebGL渲染 (卡顿暂时屏蔽)
            // this.term.loadAddon(new WebglAddon())
            // 加载窗口自适应插件
            this.fitAddon = new FitAddon()
            this.term.loadAddon(this.fitAddon)
        },
        addSocket() {
            // 创建WebSocket连接
            this.socket = new WebSocket(this.socketURI + '?cols=' + this.term.cols + '&rows' + this.term.rows + '&id=' + this.id)
            // 连接开启事件
            this.socket.onopen = () => {
                window.dispatchEvent(new CustomEvent("cache:connect", { detail: { id: this.id, connect: true } }))
            };
            // 连接关闭事件
            this.socket.onclose = () => {
                window.dispatchEvent(new CustomEvent("cache:connect", { detail: { id: this.id, connect: false } }))
            };
            this.socket.onerror = () => this.close()
            // 加载WebSocket插件
            this.term.loadAddon(new AttachAddon(this.socket))
        },
        addResizeEvent() {
            let timeout = 0
            window.addEventListener('resize', () => {
                this.fitAddon.fit();
                clearTimeout(timeout)
                timeout = setTimeout(() => {
                    this.socket.send("!~" + this.term.cols + ":" + this.term.rows)
                }, 500)
            });
        },
        send(cmd) {
            console.log(cmd)
            this.socket.send(cmd + '\n')
        },
        close() {
            this.connect = false
            // 关闭连接
            try {
                if (this.socket) this.socket.close()
            } catch (err) {
                console.log(err)
            }
            // 销毁终端
            // try {
            //     if (this.term) this.term.dispose()
            // } catch (err) {
            //     console.log(err)
            // }

            document.getElementById('xterm' + this.id).innerHTML = "";
            // document.getElementsByClassName("xterm")[0].innerHTML = "";
            window.dispatchEvent(new CustomEvent("cache:connect", { detail: { id: this.id, connect: false } }))
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
    height: calc(100vh - 117px);
    padding: 5px;
}

.notconnect {
    color: rgba(255, 255, 255, .2);
    height: calc(100vh - 107px);
    position: absolute;
    font-weight: bold;
    font-size: 42px;
    cursor: default;
    width: 100%;
    z-index: 3;
    left: 0;
    top: 0;
}
</style>