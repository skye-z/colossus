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
                <component :is="component" ref="dynamic"></component>
            </div>
        </n-modal>
    </div>
</template>
  
<script>
import { markRaw } from 'vue';
import { CloseRound } from "@vicons/material";

export default {
    name: "DynamicTool",
    components: { CloseRound },
    data: () => ({
        name: '',
        show: false,
        component: null
    }),
    props: {
        id: {
            type: Number,
            default: 0
        }
    },
    methods: {
        loadComponent: (name) => import(`../tools/${name}.vue`),
        open(name) {
            this.name = name
            this.component =
                this.loadComponent(this.name).then(res => {
                    this.show = true
                    this.component = markRaw(res.default)
                    setTimeout(()=>{
                        this.$refs.dynamic.init(this.id)
                    }, 200)
                }).catch(() => {
                    this.show = false
                    window.$message.warning('工具加载失败');
                })
        },
        close() {
            this.show = false
        },
        updateShow(show) {
            if (!show) this.result = []
        },
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
  