<template>
    <div id="app-home">
        <host-search @btn-add="addHost" />
        <host-form ref="form" />
        <n-scrollbar>
            <div id="host-list" class="pl-10 pt-10 no-select">
                <host-item v-for="item in list" :info="item" />
            </div>
        </n-scrollbar>
    </div>
</template>
  
<script>
import HostSearch from '../components/hostSearch.vue'
import HostForm from '../components/hostForm.vue'
import HostItem from '../components/hostItem.vue'

export default {
    name: "Home",
    components: { HostSearch, HostItem },
    data: () => ({
        list: [
            {
                id: 1,
                name: '测试服务器(广州)',
                address: '192.168.123.223',
                platform: 'Linux',
                system: 'CentOS',
                region: '亚洲',
                usage: '企业测试',
                period: '3天后过期'
            },
            {
                id: 2,
                name: '个人学习服务器',
                address: 'ces.hangzhou.e3.aliyun.com',
                platform: 'Linux',
                system: 'CentOS',
                usage: '个人学习'
            }
        ]
    }),
    methods: {
        addHost(){
            this.$refs.form.open('add',undefined)
        }
    }
};
</script>
<style scoped>
#app-home:deep(.n-scrollbar) {
    max-height: calc(100vh - 77px) !important;
    height: calc(100vh - 77px) !important;
    --wails-draggable: no-drag;
}

#host-list {
    flex-wrap: wrap;
    display: flex;
}
</style>