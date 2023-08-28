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
            <div class="connect-item no-drag" v-for="item in connect" @click="openHost(item)">
                <div class="line1 connect-name">{{ item.name }}</div>
                <div class="line1 connect-address">{{ item.address }}</div>
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
        connect: []
    }),
    methods: {
        init() {
            let cache = localStorage.getItem('cache:history')
            if (cache == undefined || cache == '') localStorage.setItem('cache:history', "[]")
            else this.connect = JSON.parse(cache)

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
                this.connect = cache;
                localStorage.setItem('cache:history', JSON.stringify(cache))
            })
        },
        openHost(item) {
            this.$router.push({
                path: '/host/' + item.id,
                query: { info: JSON.stringify(item) }
            })
        },
        jump(path) {
            this.$router.push('/' + path)
        }
    },
    mounted() {
        this.init()
    }
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
    transition: all ease-out 0.3s;
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

.connect-address {
    line-height: 14px;
    font-size: 12px;
    color: #adadad;
}
</style>