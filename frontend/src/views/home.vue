<template>
    <div id="app-home">
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <host-search @btn-search="search" @btn-add="addHost" />
        <host-form ref="form" @update="refresh" />
        <div v-if="list.length == 0" class="null no-select flex align-center justify-center">
            <div>无可用主机</div>
        </div>
        <n-scrollbar v-else>
            <div id="host-list" class="pl-10 pt-10 no-select">
                <host-item v-for="item in list" :info="item" @edit-host="editHost" @remove-host="removeHost" />
            </div>
        </n-scrollbar>
    </div>
</template>
  
<script>
import HostSearch from '../components/hostSearch.vue'
import HostForm from '../components/hostForm.vue'
import HostItem from '../components/hostItem.vue'
import { host } from '../plugins/api'
import util from "../plugins/util"

export default {
    name: "Home",
    components: { HostSearch, HostForm, HostItem },
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
                                res.data[i].periodTxt = util.buildPeriod(now, res.data[i].period)
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
        removeHost(info){
            window.$dialog.warning({
                title: "删除主机",
                content: "你正在尝试删除‘" + info.name + "’主机, 此项操作不可逆, 确认要删除吗?",
                negativeText: '确认删除',
                positiveText: '取消',
                onNegativeClick: () => {
                    host.remove(info.id).then(res => {
                        if (res.state) {
                            window.$message.success("删除成功");
                            this.refresh()
                        } else {
                            window.$message.warning(
                                res.message ? res.message : "删除失败, 发生意料之外的错误"
                            );
                        }
                    }).catch(err => {
                        console.log(err)
                        window.$message.warning("主机删除失败, 发生意料之外的错误");
                    })
                }
            });
        }
    },
    mounted() {
        setTimeout(() => this.getHostList({}, 1, 30), 500)
    }
};
</script>
<style scoped>
#app-home{
    position: relative;
}

#app-home:deep(.n-scrollbar) {
    max-height: calc(100vh - 77px) !important;
    height: calc(100vh - 77px) !important;
    --wails-draggable: no-drag;
}

.null {
    color: rgba(255, 255, 255, .2);
    height: calc(100vh - 75px);
    position: absolute;
    font-weight: bold;
    font-size: 42px;
    cursor: default;
    width: 100%;
    z-index: 3;
    top: 75px;
    left: 0;
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