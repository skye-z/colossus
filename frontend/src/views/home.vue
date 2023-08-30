<template>
    <div id="app-home">
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <host-search @btn-search="search" @btn-add="addHost" />
        <host-form ref="form" @update="refresh" />
        <n-scrollbar>
            <div id="host-list" class="pl-10 pt-10 no-select">
                <host-item v-for="item in list" :info="item" @edit-host="editHost" />
            </div>
        </n-scrollbar>
    </div>
</template>
  
<script>
import HostSearch from '../components/hostSearch.vue'
import HostForm from '../components/hostForm.vue'
import HostItem from '../components/hostItem.vue'
import { host } from '../plugins/api'

export default {
    name: "Home",
    components: { HostSearch, HostItem },
    data: () => ({
        loading: true,
        now: null,
        list: []
    }),
    methods: {
        search(screen) {
            this.getHostList(screen, 1, 30)
        },
        refresh() {
            this.search({})
        },
        getHostList(screen, page, number) {
            this.loading = true
            let now = new Date().getTime()
            let form = {
                page, number, ...screen
            }
            host.getList(form).then(res => {
                if (res.data) {
                    setTimeout(() => {
                        for (let i in res.data) {
                            if(res.data[i].period)
                                res.data[i].periodTxt = this.buildPeriod(now, res.data[i].period)
                            else res.data[i].period = null
                        }
                        this.list = res.data
                        this.loading = false
                    }, 500)
                } else {
                    this.list = [];
                    this.loading = false
                }
            }).catch(err => {
                console.log(err)
                this.loading = false
                window.$message.warning('获取主机列表失败, 发生意料之外的错误')
            })
        },
        addHost() {
            this.$refs.form.open('add', undefined)
        },
        editHost(info) {
            this.$refs.form.open('edit', info)
        },
        buildPeriod(now, period) {
            let oneDay = 1000 * 60 * 60 * 24
            let offset = parseInt(period) - now
            if (offset <= 0) return '已过期'
            else {
                let txt = ''
                let num = 0
                if (offset >= (oneDay * 365)) {
                    txt += (offset / (oneDay * 365)).toFixed(0) + '年'
                    offset = offset % (oneDay * 365)
                    num++
                }
                if (offset >= (oneDay * 30)) {
                    txt += (offset / (oneDay * 30)).toFixed(0) + '月'
                    offset = offset % (oneDay * 30)
                    num++
                }
                if (num == 2) return txt
                if (offset >= (oneDay * 7)) {
                    txt += (offset / (oneDay * 7)).toFixed(0) + '周'
                    offset = offset % (oneDay * 7)
                    num++
                }
                if (num == 2) return txt
                if (offset >= oneDay) {
                    txt += (offset / oneDay).toFixed(0) + '天'
                    offset = offset % oneDay
                    num++
                }
                return txt
            }
        }
    },
    mounted() {
        setTimeout(() => this.getHostList({}, 1, 30), 500)
    }
};
</script>
<style scoped>
#app-home:deep(.n-scrollbar) {
    max-height: calc(100vh - 77px) !important;
    height: calc(100vh - 77px) !important;
    --wails-draggable: no-drag;
}

.loading {
    height: calc(100% - 80px);
    top: 78px;
}

#host-list {
    flex-wrap: wrap;
    display: flex;
}
</style>