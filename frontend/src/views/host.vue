<template>
    <div>
        <div id="host-head" class="pa-10 flex align-center border-bottom">
            <div class="host-name mr-10">测试服务器(广州)</div>
            <n-tag size="small" :bordered="false" type="info">Linux</n-tag>
            <n-tag size="small" :bordered="false" type="success">CentOS</n-tag>
            <n-tag size="small" :bordered="false" type="warning">亚洲</n-tag>
            <n-tag size="small" :bordered="false" type="error">3天后过期</n-tag>
        </div>
        <div>
            <n-tabs id="host-tabs" type="line" animated>
                <n-tab name="shell">终端</n-tab>
                <n-tab name="file">文件</n-tab>
                <n-tab name="tool">工具</n-tab>
                <n-tab name="setting">配置</n-tab>
            </n-tabs>
        </div>
    </div>
</template>
  
<script>
export default {
    name: "Host",
    data: () => ({

    }),
    methods: {

    }
};
</script>
<style scoped>
#host-head {
    width: calc(100% - 201px);
    position: absolute;
}

.host-name {
    font-size: 18px;
}

#host-tabs:deep(.n-tabs-nav-scroll-content){
    margin-left: calc(100% - 250px);
    padding: 3.5px 0;
}
</style>