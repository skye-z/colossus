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
            // 监听输入事件
            this.inputEvent()
        },
        close() {
            if (this.socket) this.socket.close()
            if (this.term) this.term.dispose()
        },
        inputEvent() {
            const TERMINAL_INPUT_KEY = {
                BACK: 8, // 删除键
                TAB: 9, // TAB键
                ENTER: 13, // 回车键
                UP: 38, // 方向上键
                DOWN: 40, // 方向下键
                LEFT: 37, // 方向左键
                RIGHT: 39, // 方向右键
            };
            let prefix = 0;
            let inputText = "";
            let currentIndex = 0;
            let inputTextList = [];
            this.term.onKey((e) => {
                const { key, domEvent } = e;
                const { keyCode, altKey, altGraphKey, ctrlKey, metaKey } = domEvent;
                // 计算前缀长度
                if (inputText == "" && prefix == 0) prefix = this.term._core.buffer.x;

                const printAble = !(altKey || altGraphKey || ctrlKey || metaKey); // 禁止相关按键
                const totalOffsetLength = inputText.length + prefix; // 总偏移量
                const currentOffsetLength = this.term._core.buffer.x; // 当前x偏移量

                switch (keyCode) {
                    // 删除键
                    case TERMINAL_INPUT_KEY.BACK: {
                        if (currentOffsetLength > prefix) {
                            const cursorOffSetLength = this.getCursorOffsetLength(totalOffsetLength - currentOffsetLength, "\x1b[D"); // 保留原来光标位置

                            this.term._core.buffer.x = currentOffsetLength - 1;
                            this.term.write("\x1b[?K" + inputText.slice(currentOffsetLength - prefix));
                            this.term.write(cursorOffSetLength);
                            inputText = `${inputText.slice(0, currentOffsetLength - prefix - 1)}${inputText.slice(
                                currentOffsetLength - prefix
                            )}`;
                        }
                        break;
                    }
                    // TAB键
                    case TERMINAL_INPUT_KEY.TAB:{
                        // TODO
                    }
                    // 回车键
                    case TERMINAL_INPUT_KEY.ENTER: {
                        this.term.write("\r\n");
                        console.log("inputText", inputText);

                        if (!inputText.trim()) {
                            // this.term.prompt();
                            return;
                        }

                        if (inputTextList.indexOf(inputText) === -1) {
                            inputTextList.push(inputText);
                            currentIndex = inputTextList.length;
                        }

                        // this.term.prompt();
                        inputText = "";
                        break;
                    }
                    // 方向上键
                    case TERMINAL_INPUT_KEY.UP: {
                        if (!inputTextList[currentIndex - 1]) break;

                        const offsetLength = this.getCursorOffsetLength(inputText.length, "\x1b[D");

                        inputText = inputTextList[currentIndex - 1];
                        this.term.write(offsetLength + "\x1b[?K");
                        this.term.write(inputTextList[currentIndex - 1]);
                        this.term._core.buffer.x = totalOffsetLength;
                        currentIndex--;

                        break;
                    }
                    // 方向下键
                    case TERMINAL_INPUT_KEY.DOWN: {
                        if (!inputTextList[currentIndex + 1]) break;

                        const offsetLength = this.getCursorOffsetLength(inputText.length, "\x1b[D");

                        inputText = inputTextList[currentIndex + 1];
                        this.term.write(offsetLength + "\x1b[?K");
                        this.term.write(inputTextList[currentIndex + 1]);
                        this.term._core.buffer.x = totalOffsetLength;
                        currentIndex++;

                        break;
                    }
                    // 方向左键
                    case TERMINAL_INPUT_KEY.LEFT: {
                        if (currentOffsetLength > prefix) {
                            this.term.write(key); // '\x1b[D'
                        }
                        break;
                    }
                    // 方向右键
                    case TERMINAL_INPUT_KEY.RIGHT:{
                        if (currentOffsetLength < totalOffsetLength) {
                            this.term.write(key); // '\x1b[C'
                        }
                        break;
                    }
                    default: {
                        if (!printAble) break;
                        if (totalOffsetLength >= this.term.cols) break;
                        if (currentOffsetLength >= totalOffsetLength) {
                            this.term.write(key);
                            inputText += key;
                            break;
                        }
                        let cursorOffSetLength = this.getCursorOffsetLength(totalOffsetLength - currentOffsetLength, "\x1b[D");
                        this.term.write("\x1b[?K" + `${key}${inputText.slice(currentOffsetLength - prefix)}`);
                        this.term.write(cursorOffSetLength);
                        inputText = inputText.slice(0, currentOffsetLength) + key + inputText.slice(totalOffsetLength - currentOffsetLength);
                        break;
                    }
                }
            });
        },
        getCursorOffsetLength(offsetLength, subString) {
            let cursorOffsetLength = "";
            for (let offset = 0; offset < offsetLength; offset++) {
                cursorOffsetLength += subString;
            }
            return cursorOffsetLength;
        },
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