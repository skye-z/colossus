<template>
    <div style="position: relative;">
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <div class="cert-head border-bottom flex align-center justify-between pa-10 no-select">
            <div class="cert-title">凭证管理</div>
            <n-button-group>
                <n-button @click="getCertList">刷新</n-button>
                <n-button @click="addCert">创建</n-button>
            </n-button-group>
        </div>
        <cert-form ref="form" @update="refresh" />
        <div v-if="list.length == 0" class="null no-select flex align-center justify-center">
            <div>无可用凭证</div>
        </div>
        <n-scrollbar v-else>
            <div id="cert-list" class="pl-10 pt-10 no-select">
                <div v-for="item in list" class="cert-item">
                    <div class="cert-name line1">{{ item.name }}</div>
                    <div class="cert-tags flex mt-10">
                        <n-tag size="small" :bordered="false" type="info" v-if="item.type">{{ item.type }} {{ item.level
                        }}</n-tag>
                        <n-tag size="small" :bordered="false" type="error" v-if="item.periodTxt">{{ item.periodTxt
                        }}</n-tag>
                    </div>
                    <div class="cert-tool flex align-center justify-center">
                        <n-button class="mr-5" size="small" type="primary" @click="copyCert(item)">
                            <template #icon>
                                <n-icon size="20">
                                    <Copy16Regular />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="mr-5" size="small" type="warning" @click="editCert(item)">
                            <template #icon>
                                <n-icon size="20">
                                    <CalendarEdit16Regular />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button size="small" type="error" @click="removeCert(item)">
                            <template #icon>
                                <n-icon size="20">
                                    <Delete16Regular />
                                </n-icon>
                            </template>
                        </n-button>
                    </div>
                </div>
            </div>
        </n-scrollbar>
    </div>
</template>
  
<script>
import { Copy16Regular, CalendarEdit16Regular, Delete16Regular } from '@vicons/fluent'
import CertForm from '../components/certForm.vue'
import { cert } from "../plugins/api";

export default {
    name: "Cert",
    components: { CertForm, Copy16Regular, CalendarEdit16Regular, Delete16Regular },
    data: () => ({
        page: 1,
        number: 20,
        list: [],
        loading: true,
    }),
    methods: {
        getCertList() {
            this.loading = true
            let now = new Date().getTime()
            cert.getList({ page: this.page, number: this.number }).then(res => {
                if (res.data) {
                    setTimeout(() => {
                        for (let i in res.data) {
                            if (res.data[i].period)
                                res.data[i].periodTxt = this.buildPeriod(now, res.data[i].period)
                            else {
                                res.data[i].periodTxt = "长期"
                                res.data[i].period = null
                            }
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
                window.$message.warning('获取命令列表失败, 发生意料之外的错误')
            })
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
        },
        addCert() {
            this.$refs.form.open('add', undefined)
        },
        editCert(info) {
            this.$refs.form.open('edit', info)
        },
        removeCert(item) {
            window.$dialog.warning({
                title: "删除凭证",
                content: "你正在尝试删除‘" + item.name + "’凭证, 此项操作不可逆, 确认要删除吗?",
                negativeText: '确认删除',
                positiveText: '取消',
                onNegativeClick: () => {
                    cert.remove(item.id).then(res => {
                        if (res.state) {
                            window.$message.success("删除成功");
                            this.getCertList()
                        } else {
                            window.$message.warning(
                                res.message ? res.message : "删除失败, 发生意料之外的错误"
                            );
                        }
                    }).catch(err => {
                        console.log(err)
                        window.$message.warning("凭证删除失败, 发生意料之外的错误");
                    })
                }
            });
        },
        copyCert(item) {
            let content = item.privateKey
            if (item.publicKey) content += item.publicKey
            this.$goSetClipboard(content)
            window.$message.success('凭证已复制到剪切板')
        },
        refresh() {
            this.getCertList()
        }
    },
    mounted() {
        setTimeout(() => this.getCertList(), 300)
    }
};
</script>
<style scoped>
.loading {
    height: calc(100% - 55px);
    top: 55px;
}

.cert-title {
    font-size: 18px;
}

.null {
    color: rgba(255, 255, 255, .2);
    height: calc(100vh - 55px);
    position: absolute;
    font-weight: bold;
    font-size: 42px;
    cursor: default;
    width: 100%;
    z-index: 3;
    top: 55px;
    left: 0;
}

#cert-list {
    flex-wrap: wrap;
    display: flex;
}

.cert-item {
    background-color: rgba(255, 255, 255, .08);
    margin: 0 10px 10px 0;
    border-radius: 8px;
    position: relative;
    min-width: 180px;
    cursor: pointer;
}

.cert-name {
    padding: 10px 10px 0 10px;
    max-width: 200px;
}

.cert-tags>div {
    justify-content: center;
    border-radius: 0;
    cursor: pointer;
    width: 100%;
}

.cert-tags>div:first-child {
    border-radius: 0 0 0 8px;
}

.cert-tags>div:last-child {
    border-radius: 0 0 8px 0;
}

.cert-tool {
    background-color: rgba(255, 255, 255, .08);
    transition: all ease-out 0.3s;
    border-radius: 8px;
    position: absolute;
    height: 100%;
    width: 100%;
    opacity: 0;
    left: 0;
    top: 0;
}

.cert-tool>button {
    width: calc(33% - 7px);
    height: calc(100% - 10px);
}

.cert-item:hover .cert-tool {
    opacity: 1;
}
</style>