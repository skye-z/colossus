<template>
    <div id="left-bar" class="no-select">
        <div id="logo" :class="{ macos: macOS }">Colossus_</div>
        <n-button id="btn-add" class="no-drag" strong secondary type="primary" @click="jump('home')">
            <n-icon size="28">
                <InsertLinkOutlined />
            </n-icon>
            <div>创建连接</div>
        </n-button>
        <div id="connect-list" :class="{ macos: macOS }">
            <div class="connect-item no-drag" v-for="item in connect" @click="openHost(item)"
                :class="{ 'select': select == item.id }">
                <div class="line1 connect-name">{{ item.name }}</div>
                <div class="line1 connect-address flex align-center">
                    <div :class="'dot mr-5 ' + (item.connect ? 'dot-connect' : '')"></div>
                    <div>{{ item.address }}</div>
                </div>
            </div>
        </div>
        <div class="flex align-center justify-between pl-10 no-drag">
            <n-button-group>
                <n-button round @click="jump('cert')">
                    <template #icon>
                        <n-icon size="22">
                            <Key16Regular />
                        </n-icon>
                    </template>
                </n-button>
                <n-button round @click="jump('code')">
                    <template #icon>
                        <n-icon size="22">
                            <Code16Filled />
                        </n-icon>
                    </template>
                </n-button>
                <n-button round @click="jump('setting')">
                    <template #icon>
                        <n-icon size="22">
                            <Settings16Regular />
                        </n-icon>
                    </template>
                </n-button>
            </n-button-group>
        </div>
    </div>
</template>
<script>
import { Key16Regular, Code16Filled, Settings16Regular } from '@vicons/fluent'
import { InsertLinkOutlined } from '@vicons/material'

export default {
    name: "LeftBar",
    components: { InsertLinkOutlined, Key16Regular, Code16Filled, Settings16Regular },
    props: {
        macOS: {
            type: Boolean,
            default: false
        }
    },
    data: () => ({
        connect: [],
        select: null
    }),
    methods: {
        init() {
            this.initHistory()
            this.initHistoryEvent();
            this.initConnectEvent();
        },
        initHistory(){
            let cache = localStorage.getItem('cache:history')
            if (cache == undefined || cache == '') localStorage.setItem('cache:history', "[]")
            else {
                let list = JSON.parse(cache)
                for (let i in list) list[i].connect = false
                this.connect = list
            }
        },
        initHistoryEvent(){
            window.addEventListener("cache:history", ({ detail }) => {
                let cache = localStorage.getItem('cache:history')
                if (cache == undefined || cache == '') cache = []
                else cache = JSON.parse(cache)

                let exist = false
                for (let i in cache) {
                    if (cache[i].id == detail.id) {
                        exist = true;
                        break;
                    }
                }
                if (!exist) cache.unshift(detail)
                let connect = []
                for (let x in cache) {
                    for (let y in this.connect) {
                        if (cache[x].id == this.connect[y].id) cache[x].connect = this.connect[y].connect
                    }
                    connect.push(cache[x])
                }
                this.connect = connect;
                localStorage.setItem('cache:history', JSON.stringify(cache))
            })
        },
        initConnectEvent(){
            window.addEventListener("cache:connect", ({ detail }) => {
                for (let i in this.connect) {
                    if (this.connect[i].id == detail.id) {
                        this.connect[i].connect = detail.connect
                        break;
                    }
                }
            })
        },
        openHost(item) {
            this.$router.push('/host/' + item.id)
        },
        jump(path) {
            this.$router.push('/' + path)
        },
        illuminate(id) {
            this.select = id
        }
    },
    mounted() {
        this.init()
    },
    watch: {
        $route: {
            handler(to) {
                if (to.params.id) this.illuminate(to.params.id)
                else this.illuminate(null)
            },
            deep: true,
        },
    },
};
</script>
<style scoped>
#left-bar {
    border-right: rgba(255, 255, 255, 0.12) 1px solid;
    height: 100vh;
    width: 180px;
}

#logo {
    text-align: center;
    font-weight: bold;
    font-size: 28px;
    cursor: default;
    padding: 5px 0;
}

#logo.macos {
    padding: 30px 0 5px 0;
}

#btn-add {
    width: calc(100% - 20px);
    margin: 0 0 10px 10px;
}

#connect-list {
    height: calc(100vh - 155px);
    padding: 0 10px 10px 10px;
}

#connect-list.macos {
    height: calc(100vh - 180px);
}

.connect-item {
    background-color: rgba(255, 255, 255, .08);
    transition: background-color ease-out 0.3s;
    padding: 6px 5px 8px 5px;
    margin-bottom: 10px;
    border-radius: 8px;
    cursor: pointer;
}

.connect-item:hover {
    background-color: rgba(255, 255, 255, .12);
}

.connect-item:active {
    background-color: rgba(255, 255, 255, .08);
}

.connect-item.select {
    background-color: rgba(99, 226, 183, .12);
    border: 1px solid rgba(99, 226, 183, .3);
    padding: 5px 4px 7px 4px;
}

.connect-address {
    line-height: 14px;
    font-size: 12px;
    color: #adadad;
}

.dot {
    background-color: #dc143c;
    border-radius: 8px;
    height: 8px;
    width: 8px;
}

.dot-connect {
    background-color: #14dc28;
}
</style>