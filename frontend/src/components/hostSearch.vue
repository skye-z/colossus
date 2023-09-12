<template>
    <div class="no-select">
        <n-input-group class="no-drag">
            <n-input id="search" v-model:value="screen.keyword" type="text" placeholder="搜索主机名称或地址..." clearable>
                <template #prefix>
                    <n-icon>
                        <Search />
                    </n-icon>
                </template>
            </n-input>
            <n-button strong secondary @click="submit">搜索</n-button>
            <n-button id="btn-add-host" type="primary" tertiary @click="addHost">
                <n-icon size="20">
                    <Add12Regular />
                </n-icon>
                <div class="ml-5">新增主机</div>
            </n-button>
        </n-input-group>
        <div class="pa-10 border-bottom">
            <n-popselect v-model:value="screen.platform" :options="options.platform" trigger="click" @update:value="submit">
                <n-button size="tiny" round>平台: {{ screen.platform || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.system" :options="options.system" virtual-scroll trigger="click" @update:value="submit">
                <n-button class="ml-10" size="tiny" round>系统: {{ screen.system || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.region" :options="options.region" trigger="click" @update:value="submit">
                <n-button class="ml-10" size="tiny" round>区域: {{ screen.region || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.usage" :options="options.usage" trigger="click" @update:value="submit">
                <n-button class="ml-10" size="tiny" round>用途: {{ screen.usage || 'All' }}</n-button>
            </n-popselect>
            <n-popselect v-model:value="screen.period" :options="options.period" trigger="click" @update:value="submit">
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
import options from "../plugins/options"

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
        options: {}
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
            this.submit()
        },
        submit() {
            this.$emit('btnSearch', this.screen)
        },
        getScreen(){
            return screen
        },
        addHost() {
            this.$emit('btnAdd')
        }
    },
    mounted(){
        this.options = options(["platform","system","region","usage","period"],true)
    }
};
</script>
<style scoped>
#search,
#btn-add-host {
    border-radius: 0;
}
</style>