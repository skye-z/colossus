<template>
    <div class="host-item">
        <div class="host-name line1">{{ info.name }}</div>
        <div class="host-address line1">{{ info.address }}</div>
        <div class="host-tags flex">
            <n-tag size="small" :bordered="false" type="info" v-if="info.platform">{{ info.platform }}</n-tag>
            <n-tag size="small" :bordered="false" type="success" v-if="info.system">{{ info.system }}</n-tag>
            <n-tag size="small" :bordered="false" type="warning" v-if="info.region">{{ info.region }}</n-tag>
            <n-tag size="small" :bordered="false" type="error" v-if="info.periodTxt">{{ info.periodTxt }}</n-tag>
        </div>
        <div class="host-tool">
            <div class="pa-5">
                <n-button class="full-width" type="primary" @click="openHost()">
                    连接
                </n-button>
            </div>
            <div class="host-tool-sub flex align-center justify-center">
                <n-button class="mr-5" size="small" type="warning" @click="editHost">
                    <template #icon>
                        <n-icon>
                            <LinkEdit16Filled />
                        </n-icon>
                    </template>
                </n-button>
                <n-button size="small" type="error" @click="removeHost">
                    <template #icon>
                        <n-icon>
                            <Delete16Filled />
                        </n-icon>
                    </template>
                </n-button>
            </div>
        </div>
    </div>
</template>
<script>
import { PlugConnected20Filled, LinkEdit16Filled, Delete16Filled } from '@vicons/fluent'
export default {
    name: "HostItem",
    components: { PlugConnected20Filled, LinkEdit16Filled, Delete16Filled },
    props: {
        info: {
            type: Object,
            default: {}
        }
    },
    methods: {
        openHost() {
            window.dispatchEvent(new CustomEvent("cache:history", { detail: this.info }))
            this.$router.push('/host/' + this.info.id)
        },
        editHost() {
            this.$emit('editHost', this.info)
        },
        removeHost() {

        }
    }
};
</script>
<style scoped>
.host-item {
    background-color: rgba(255, 255, 255, .08);
    margin: 0 10px 10px 0;
    border-radius: 8px;
    position: relative;
    min-width: 180px;
}

.host-tool {
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

.host-tool-sub>button {
    width: calc(50% - 7px);
}

.host-item:hover .host-tool {
    opacity: 1;
}

.host-name {
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

.host-tags>div {
    justify-content: center;
    border-radius: 0;
    cursor: pointer;
    width: 100%;
}

.host-tags>div:first-child {
    border-radius: 0 0 0 8px;
}

.host-tags>div:last-child {
    border-radius: 0 0 8px 0;
}</style>