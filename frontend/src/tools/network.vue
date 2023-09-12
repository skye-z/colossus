<template>
    <div>
        <n-modal v-model:show="show" @update:show="updateShow">
            <div class="no-select">
                <n-button class="close" quaternary circle size="small" @click="close">
                    <template #icon>
                        <n-icon>
                            <CloseRound />
                        </n-icon>
                    </template>
                </n-button>
                <n-card style="width: 100vw;height: calc(100vh - 2px);" header-style="text-align: center" title="网络连接"
                    :bordered="false" size="small" role="dialog" aria-modal="true">
                    <n-data-table size="small" max-height="calc(100vh - 105px)" :loading="loading" virtual-scroll
                        :columns="columns" :data="result" :bordered="false" />
                </n-card>
            </div>
        </n-modal>
    </div>
</template>
  
<script>
import { CloseRound } from "@vicons/material";
import { host } from "../plugins/api";
import { NButton } from 'naive-ui'
import { h, getCurrentInstance } from "vue"

export default {
    name: "ToolNetword",
    components: { CloseRound },
    data: vm => ({
        id: 0,
        show: false,
        loading: false,
        command: "ss -np",
        columns: [{
            title: "类型",
            key: "netid",
            width: 80,
            sorter: 'default'
        }, {
            title: "状态",
            key: "state",
            width: 100,
            sorter: 'default'
        }, {
            title: "接收队列",
            key: "recv-q",
            width: 95,
            sorter: 'default'
        }, {
            title: "发送队列",
            key: "send-q",
            width: 95,
            sorter: 'default'
        }, {
            title: "本地地址",
            key: "local",
            sorter: 'default',
            ellipsis: {
                tooltip: true
            }
        }, {
            title: "本地端口",
            key: "localPort",
            width: 95,
            defaultSortOrder: 'ascend',
            sorter: 'default'
        }, {
            title: "远程地址",
            key: "peer",
            sorter: 'default',
            ellipsis: {
                tooltip: true
            }
        }, {
            title: "远程端口",
            key: "peerPort",
            width: 95,
            sorter: 'default'
        }, {
            title: "",
            key: "process",
            align: 'right',
            width: 70,
            render(row) {
                return h(
                    NButton,
                    {
                        size: 'small',
                        type: 'primary',
                        tertiary: true,
                        onClick: () => vm.showProcess(row)
                    },
                    { default: () => "进程" }
                )
            }
        }],
        result: []
    }),
    methods: {
        open(id) {
            this.id = id;
            this.show = true
            this.getData()
        },
        close() {
            this.show = false
        },
        updateShow(show) {
            if (!show) this.result = []
        },
        getData() {
            this.result = []
            this.loading = true
            host.runCMD(this.id, this.command).then(res => {
                this.loading = false
                if (res.state) {
                    let list = []
                    let cache = res.data.split('\n')
                    for (let i = 1; i < cache.length; i++) {
                        if (cache[i].length == 0) continue;
                        let item = cache[i].replaceAll('     ', ' ').replaceAll('   ', ' ').replaceAll('  ', ' ').replaceAll('  ', ' ')
                        item = item.replaceAll('* ', '*')
                        list.push(item.split(' '))
                    }

                    let dataList = []
                    for (let i in list) {
                        let num = list[i].length
                        let item = {
                            "netid": this.analysisNetid(list[i][0]),
                            "state": list[i][1],
                            "recv-q": parseInt(list[i][2]),
                            "send-q": parseInt(list[i][3]),
                            "process": list[i][num - 1]
                        }
                        if (num == 7) {
                            this.analysis(item, 'local', list[i][4])
                            this.analysis(item, 'peer', list[i][5])
                        } else if (num == 8) {
                            item.local = list[i][4]
                            item.localPort = list[i][5]
                            this.analysis(item, 'peer', list[i][6])
                        } else if (num == 9) {
                            item.local = list[i][4]
                            item.localPort = list[i][5]
                            item.peer = list[i][6]
                            item.peerPort = list[i][7]
                        }
                        dataList.push(item)
                    }

                    this.result = dataList
                } else {
                    window.$message.warning('执行失败, 系统暂不支持')
                }
            }).catch(err => {
                console.log(err)
                this.loading = false
                window.$message.warning('执行失败, 发生意料之外的错误')
            })
        },
        analysis(item, prefix, data) {
            if (data.indexOf('*') == 0) {
                item[prefix] = '*'
                item[prefix + 'Port'] = parseInt(data.substring(1))
            } else {
                item[prefix] = data.substring(0, data.lastIndexOf(":"))
                item[prefix + 'Port'] = parseInt(data.substring(data.lastIndexOf(":") + 1))
            }
        },
        analysisNetid(text) {
            if ('u_str' === text) return 'Socket'
            else return text.toUpperCase()
        },
        showProcess(row) {
            console.log(row)
        }
    }
};
</script>
<style scoped>
.close {
    position: fixed;
    z-index: 999;
    right: 10px;
    top: 10px;
}
</style>
  