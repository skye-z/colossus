<template>
    <div>
        <n-input-group>
            <n-input id="search" v-model:value="screen.keyword" type="text" placeholder="搜索主机名称或地址...">
                <template #prefix>
                    <n-icon>
                        <Search />
                    </n-icon>
                </template>
            </n-input>
            <n-button strong secondary>搜索</n-button>
            <n-button id="btn-add-host" type="primary" tertiary>
                <n-icon size="20">
                    <Add12Regular />
                </n-icon>
                <div class="ml-5">新增主机</div>
            </n-button>
        </n-input-group>
        <div class="pa-10 border-bottom">
            <n-popselect v-model:value="screen.platform" :options="options.platform" trigger="click">
                <n-button size="tiny" round>平台: {{ screen.platform || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.system" :options="options.system" trigger="click">
                <n-button class="ml-10" size="tiny" round>系统: {{ screen.system || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.region" :options="options.region" trigger="click">
                <n-button class="ml-10" size="tiny" round>区域: {{ screen.region || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.usage" :options="options.usage" trigger="click">
                <n-button class="ml-10" size="tiny" round>用途: {{ screen.usage || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.period" :options="options.period" trigger="click">
                <n-button class="ml-10" size="tiny" round>有效期: {{ screen.period || 'All' }}</n-button>
            </n-popselect>
            <n-button class="float-right" size="tiny" type="error" strong secondary round
                @click="cleanScreen">清除所有条件</n-button>
        </div>
        <n-scrollbar id="host-scrollbar">
            <div id="host-list" class="pl-10 pt-10">
                <div class="host-item" @click="openHost()">
                    <div class="host-name line1">测试服务器(广州)</div>
                    <div class="host-address line1">192.168.123.223</div>
                    <div class="mt-10">
                        <n-tag size="small" :bordered="false" type="info">Linux</n-tag>
                        <n-tag size="small" :bordered="false" type="success">CentOS</n-tag>
                        <n-tag size="small" :bordered="false" type="warning">亚洲</n-tag>
                        <n-tag size="small" :bordered="false" type="error">3天后过期</n-tag>
                    </div>
                </div>
                <div class="host-item">
                    <div class="host-name line1">个人测试服务器</div>
                    <div class="host-address line1">ces.hangzhou.e3.aliyun.com</div>
                    <div class="mt-10">
                        <n-tag size="small" :bordered="false" type="info">Linux</n-tag>
                        <n-tag size="small" :bordered="false" type="success">CentOS</n-tag>
                        <n-tag size="small" :bordered="false" type="warning">亚洲</n-tag>
                    </div>
                </div>
            </div>
        </n-scrollbar>
    </div>
</template>
  
<script>
import { Add12Regular } from '@vicons/fluent'
import { Search } from '@vicons/tabler'

export default {
    name: "Home",
    components: { Search, Add12Regular },
    data: () => ({
        screen: {
            keyword: '',
            platform: '',
            system: '',
            region: '',
            usage: '',
            period: ''
        },
        options: {
            platform: [
                {
                    label: "All",
                    value: ""
                },
                {
                    label: "Linux",
                    value: "Linux"
                },
                {
                    label: "MacOS",
                    value: "MacOS"
                },
                {
                    label: "Windows",
                    value: "Windows"
                },
            ],
            system: [
                {
                    label: "All",
                    value: ""
                },
                {
                    label: "RedHat",
                    value: "RedHat"
                },
                {
                    label: "Ubuntu",
                    value: "Ubuntu"
                },
                {
                    label: "CentOS",
                    value: "CentOS"
                },
                {
                    label: "Debian",
                    value: "Debian"
                },
                {
                    label: "Oracle",
                    value: "Oracle"
                },
                {
                    label: "SUSE",
                    value: "SUSE"
                },
                {
                    label: "Fedora",
                    value: "Fedora"
                },
                {
                    label: "Arch",
                    value: "Arch"
                }, ,
                {
                    label: "FreeBSD",
                    value: "FreeBSD"
                }
            ],
            region: [
                {
                    label: "All",
                    value: ""
                },
                {
                    label: "亚洲",
                    value: "亚洲"
                },
                {
                    label: "欧洲",
                    value: "欧洲"
                },
                {
                    label: "北美洲",
                    value: "北美洲"
                },
                {
                    label: "南美洲",
                    value: "南美洲"
                },
                {
                    label: "大洋洲",
                    value: "大洋洲"
                },
                {
                    label: "非洲",
                    value: "非洲"
                },
            ],
            usage: [
                {
                    label: "All",
                    value: ""
                },
                {
                    label: "个人学习",
                    value: "个人学习"
                },
                {
                    label: "个人项目",
                    value: "个人项目"
                },
                {
                    label: "个人测试",
                    value: "个人测试"
                },
                {
                    label: "家庭服务",
                    value: "家庭服务"
                },
                {
                    label: "企业项目",
                    value: "企业项目"
                },
                {
                    label: "企业测试",
                    value: "企业测试"
                },
            ],
            period: [
                {
                    label: "All",
                    value: ""
                },
                {
                    label: "长期有效",
                    value: "长期有效"
                },
                {
                    label: "近7天过期",
                    value: "近7天过期"
                },
                {
                    label: "近15天过期",
                    value: "近15天过期"
                },
                {
                    label: "近30天过期",
                    value: "近30天过期"
                },
            ]
        }
    }),
    methods: {
        cleanScreen() {
            this.screen = {
                keyword: '',
                platform: '',
                system: '',
                region: '',
                usage: '',
                period: ''
            }
        },
        openHost(){
            this.$router.push('/host')
        }
    }
};
</script>
<style scoped>
#search,
#btn-add-host {
    border-radius: 0;
}

#host-scrollbar {
    height: calc(100vh - 77px);
}

#host-list {
    flex-wrap: wrap;
    display: flex;
}

.host-item {
    background-color: rgba(255, 255, 255, .08);
    transition: all ease-out 0.3s;
    margin: 0 10px 10px 0;
    border-radius: 8px;
    cursor: pointer;
    padding: 10px;
}

.host-item:hover {
    background-color: rgba(255, 255, 255, .12);
}

.host-item:active {
    background-color: rgba(255, 255, 255, .08);
}

.host-address {
    line-height: 14px;
    font-size: 12px;
    color: #adadad;
}
</style>