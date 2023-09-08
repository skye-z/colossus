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
                <div v-for="item in list" class="code-item">
                    <div class="code-name line1">{{ item.name }}</div>
                    <div class="code-tags flex mt-10">
                        <n-tag size="small" :bordered="false" type="info" v-if="item.platform">{{ item.platform
                        }}</n-tag>
                        <n-tag size="small" :bordered="false" type="success" v-if="item.system">{{ item.system
                        }}</n-tag>
                    </div>
                    <div class="code-tool flex align-center justify-center">
                        <n-button class="mr-5" size="small" type="primary" @click="copyCode(item.content)">
                            <template #icon>
                                <n-icon size="20">
                                    <Copy16Regular />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="mr-5" size="small" type="warning" @click="editCode(item)">
                            <template #icon>
                                <n-icon size="20">
                                    <CalendarEdit16Regular />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button size="small" type="error" @click="removeCode(item)">
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
import CodeForm from '../components/codeForm.vue'
import { code } from "../plugins/api"
import { Base64 } from 'js-base64'

export default {
    name: "Code",
    components: { CodeForm, Copy16Regular, CalendarEdit16Regular, Delete16Regular },
    data: () => ({
        page: 1,
        number: 20,
        list: [],
        loading: true,
    }),
    methods: {
        getCodeList() {
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
        copyCode(content) {
            this.$goSetClipboard(Base64.decode(content))
            window.$message.success('命令已复制到剪切板')
        },
        removeCode(item) {
            window.$dialog.warning({
                title: "删除快捷命令",
                content: "你正在尝试删除‘" + item.name + "’快捷命令, 此项操作不可逆, 确认要删除吗?",
                negativeText: '确认删除',
                positiveText: '取消',
                onNegativeClick: () => {
                    code.remove(item.id).then(res => {
                        if (res.state) {
                            window.$message.success("删除成功");
                            this.getCodeList()
                        } else {
                            window.$message.warning(
                                res.message ? res.message : "删除失败, 发生意料之外的错误"
                            );
                        }
                    }).catch(err => {
                        console.log(err)
                        window.$message.warning("快捷命令删除失败, 发生意料之外的错误");
                    })
                }
            });
        },
        refresh() {
            this.getCodeList()
        }
    },
    mounted() {
        setTimeout(() => this.getCodeList(), 300)
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

.code-tool {
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

.code-tool>button {
    width: calc(33% - 7px);
    height: calc(100% - 10px);
}

.code-item:hover .code-tool {
    opacity: 1;
}
</style>