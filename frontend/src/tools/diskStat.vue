<template>
    <div>
        <n-card class="full-screen" title="磁盘统计" header-style="text-align: center" :bordered="false" size="small">
            <div class="loading flex align-center justify-center" v-if="loading">
                <n-spin size="small" />
            </div>
            <n-scrollbar style="max-height: calc(100vh - 65px)">
                <div id="group-list">
                    <div v-for="(sub, key) in group" class="group-item">
                        <div class="pa-10 float-right text-gray" v-if="tag[key]">{{ tag[key] }}</div>
                        <div class="pa-10">{{ key }}</div>
                        <div v-for="item in sub" class="border-top">
                            <div class="pa-5">{{ item.position }}</div>
                            <n-progress type="line" :percentage="item.proportion" color="#e88080" rail-color="#63e2b7"
                                :height="1" :border-radius="0" :show-indicator="false" status="success" />
                            <div class="flex align-center justify-between text-small text-gray pa-5">
                                <div>{{ item.used }}</div>
                                <div>{{ item.avail }}</div>
                            </div>
                        </div>
                    </div>
                </div>
            </n-scrollbar>
        </n-card>
    </div>
</template>
  
<script>
import { host } from "../plugins/api"

export default {
    name: "DiskStat",
    data: () => ({
        id: 0,
        loading: true,
        command: "df -h",
        group: {},
        tag: {
            "devtmpfs": "驱动文件",
            "tmpfs": "内存文件",
            "overlay": "联合文件",
        }
    }),
    methods: {
        init(id) {
            this.id = id
            this.getData()
        },
        getData() {
            host.runCMD(this.id, this.command).then(res => {
                this.loading = false
                if (res.state) {
                    let list = []
                    let cache = res.data.split('\n')
                    for (let i = 1; i < cache.length; i++) {
                        if (cache[i].length == 0) continue;
                        let item = cache[i].replaceAll('     ', ' ').replaceAll('   ', ' ').replaceAll('  ', ' ').replaceAll('  ', ' ')
                        list.push(item.split(' '))
                    }
                    let group = {}
                    for (let i in list) {
                        let item = list[i]
                        let sub = group[item[0]]
                        if (sub == undefined) sub = []
                        sub.push({
                            "size": item[1] + 'B',
                            "used": item[2] + 'B',
                            "avail": item[3] + 'B',
                            "proportion": parseInt(item[4].substring(0, item[4].length - 1)),
                            "position": item[5]
                        })
                        group[item[0]] = sub
                    }
                    this.group = group
                } else {
                    window.$message.warning('执行失败, 系统暂不支持')
                }
            }).catch(err => {
                console.log(err)
                this.loading = false
                window.$message.warning('执行失败, 发生意料之外的错误')
            })
        }
    }
};
</script>
<style scoped>
#group-list {
    flex-wrap: wrap;
    display: flex;
}

.group-item {
    background-color: rgba(255, 255, 255, .08);
    margin: 0 10px 10px 0;
    border-radius: 8px;
    position: relative;
    min-width: 188px;
}
</style>