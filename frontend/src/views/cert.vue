<template>
    <div>
        <div class="cert-head border-bottom flex align-center justify-between pa-10 no-select">
            <div class="cert-title">凭证管理</div>
            <n-button-group>
                <n-button>导入</n-button>
                <n-button>创建</n-button>
            </n-button-group>
        </div>
        <div class="tips text-center">
            <div>等待开发</div>
            <div class="text-small text-gray">创建证书、管理证书、管理密码</div>
        </div>
    </div>
</template>
  
<script>
export default {
    name: "Cert",
    data: () => ({
    }),
    methods: {

    },
    mounted() {

    }
};
</script>
<style scoped>
.cert-title{
    font-size: 18px;
}


.tips {
    padding-top: 40vh;
}
</style>