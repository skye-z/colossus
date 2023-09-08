<template>
    <div style="position: relative;">
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <div class="border-bottom flex align-center justify-between pa-10 no-select">
            <div class="code-title">快捷命令</div>
            <n-button-group>
                <n-button>导入</n-button>
                <n-button @click="addCode">创建</n-button>
            </n-button-group>
        </div>
        <code-form ref="form" @update="refresh" />

        <div v-if="list.length == 0" class="null no-select flex align-center justify-center">
            <div>无可用命令</div>
        </div>
        <n-scrollbar v-else>
            <div id="code-list" class="pl-10 pt-10 no-select">
                <div v-for="item in list" class="code-item" @click="editCode(item)">
                    <div class="code-name line1">{{ item.name }}</div>
                    <div class="code-tags flex mt-10">
                        <n-tag size="small" :bordered="false" type="info" v-if="item.platform">{{ item.platform
                        }}</n-tag>
                        <n-tag size="small" :bordered="false" type="success" v-if="item.system">{{ item.system
                        }}</n-tag>
                    </div>
                </div>
            </div>
        </n-scrollbar>
    </div>
</template>
  
<script>
import CodeForm from '../components/codeForm.vue'
import { code } from "../plugins/api"

export default {
    name: "Code",
    components: { CodeForm },
    data: () => ({
        page: 1,
        number: 20,
        list: [],
        loading: true,
    }),
    methods: {
        getHostList() {
            this.loading = true
            code.getList({ page: this.page, number: this.number }).then(res => {
                if (res.data) {
                    setTimeout(() => {
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
        addCode() {
            this.$refs.form.open('add', undefined)
        },
        editCode(info) {
            this.$refs.form.open('edit', info)
        },
        refresh() {
            this.getHostList()
        }
    },
    mounted() {
        setTimeout(() => this.getHostList(), 300)
    }
};
</script>
<style scoped>

.loading {
    height: calc(100% - 55px);
    top: 55px;
}

.code-title {
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

#code-list {
    flex-wrap: wrap;
    display: flex;
}

.code-item {
    background-color: rgba(255, 255, 255, .08);
    margin: 0 10px 10px 0;
    border-radius: 8px;
    position: relative;
    min-width: 180px;
    cursor: pointer;
}

.code-item:hover{
    background-color: rgba(255, 255, 255, .12);
}

.code-name {
    padding: 10px 10px 0 10px;
    max-width: 200px;
}

.code-tags>div {
    justify-content: center;
    border-radius: 0;
    cursor: pointer;
    width: 100%;
}

.code-tags>div:first-child {
    border-radius: 0 0 0 8px;
}

.code-tags>div:last-child {
    border-radius: 0 0 8px 0;
}
</style>