<template>
    <div class="host-item" @click="openHost()">
        <div class="host-name line1">{{ info.name }}</div>
        <div class="host-address line1">{{ info.address }}</div>
        <div class="host-tags flex">
            <n-tag size="small" :bordered="false" type="info" v-if="info.platform">{{ info.platform }}</n-tag>
            <n-tag size="small" :bordered="false" type="success" v-if="info.system">{{ info.system }}</n-tag>
            <n-tag size="small" :bordered="false" type="warning" v-if="info.region">{{ info.region }}</n-tag>
            <n-tag size="small" :bordered="false" type="error" v-if="info.period">{{ info.period }}</n-tag>
        </div>
    </div>
</template>
<script>
export default {
    name: "HostItem",
    props: {
        info: {
            type: Object,
            default: {}
        }
    },
    methods: {
        openHost() {
            this.$router.push('/host')
        }
    }
};
</script>
<style scoped>
.host-item {
    background-color: rgba(255, 255, 255, .08);
    transition: all ease-out 0.3s;
    margin: 0 10px 10px 0;
    border-radius: 8px;
    cursor: pointer;
}

.host-item:hover {
    background-color: rgba(255, 255, 255, .12);
}

.host-item:active {
    background-color: rgba(255, 255, 255, .08);
}

.host-name{
    padding: 10px 10px 0 10px;
    max-width: 200px;
}

.host-address {
    padding: 0 10px 10px 10px;
    line-height: 14px;
    max-width: 240px;
    font-size: 12px;
    color: #adadad;
}

.host-tags>div{
    justify-content: center;
    width: 100%;
}

.host-tags>div:first-child{
    border-radius: 0 0 0 8px;
}

.host-tags>div:last-child{
    border-radius: 0 0 8px 0;
}
</style>