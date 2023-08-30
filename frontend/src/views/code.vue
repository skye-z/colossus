<template>
    <div>
        <div class="border-bottom flex align-center justify-between pa-10 no-select">
            <div class="code-title">快捷命令</div>
            <n-button-group>
                <n-button>导入</n-button>
                <n-button>创建</n-button>
            </n-button-group>
        </div>
        <div class="tips text-center">
            <div>等待开发</div>
            <div class="text-small text-gray">命令管理、命令导入、历史记录</div>
        </div>
    </div>
</template>
  
<script>
export default {
    name: "Code",
    data: () => ({
    }),
    methods: {

    },
    mounted() {

    }
};
</script>
<style scoped>
.code-title{
    font-size: 18px;
}

.tips {
    padding-top: 40vh;
}
</style>