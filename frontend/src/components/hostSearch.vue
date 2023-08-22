<template>
    <div class="no-select">
        <n-input-group class="no-drag">
            <n-input id="search" v-model:value="screen.keyword" type="text" placeholder="搜索主机名称或地址...">
                <template #prefix>
                    <n-icon>
                        <Search />
                    </n-icon>
                </template>
            </n-input>
            <n-button strong secondary>搜索</n-button>
            <n-button id="btn-add-host" type="primary" tertiary @click="addHost">
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
    </div>
</template>
<script>
import { Add12Regular } from '@vicons/fluent'
import { Search } from '@vicons/tabler'

export default {
    name: "HostSearch",
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
        addHost(){
	      this.$emit('btnAdd')
        }
    }
};
</script>
<style scoped>
#search,
#btn-add-host {
    border-radius: 0;
}
</style>