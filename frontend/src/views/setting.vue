<template>
    <div class="no-select">
        <div class="border-bottom flex align-center justify-between pa-10">
            <div class="seting-title">设置</div>
            <div class="text-gray">内部预览版</div>
        </div>
        <div class="setting-box flex justify-between pa-10">
            <div class="full-width mr-10">
                <div class="card pa-10 mb-10">
                    <div>终端背景</div>
                    <div>终端字号</div>
                    <div>终端字色</div>
                </div>
                <div class="card pa-10 mb-10">
                    <div>下载目录</div>
                </div>
                <div class="card pa-10 mb-10">
                    <div>加密策略</div>
                </div>
                <div class="card pa-10">
                    <div>缓存清理</div>
                </div>
            </div>
            <div>
                <div class="card setting-info">
                    <div class="text-center">
                        <img class="info-logo" src="../assets/logo.png" />
                        <div class="info-name">Colossus</div>
                        <div class="border-bottom pa-10 pt-0">Version 0.0.1 Alpha-3</div>
                    </div>
                    <div class="border-bottom pa-10">
                        <div class="flex align-center mb-5">
                            <n-icon size="20" class="mr-5">
                                <LogoDevFilled />
                            </n-icon>
                            <div class="mr-5">Skye</div>
                            <div class="text-small text-gray">github.com/skye-z</div>
                        </div>
                        <div class="flex align-center">
                            <n-icon size="20" class="mr-5">
                                <HomeRound />
                            </n-icon>
                            <div>colossus.betax.dev</div>
                        </div>
                    </div>
                    <div class="pa-10 text-small">
                        <div class="info-name mb-5">致谢 & 开源组件</div>
                        <div class="flex justify-between">
                            <div>golang/go</div>
                            <div class="text-gray">BSD 3 Clause</div>
                        </div>
                        <div class="flex justify-between">
                            <div>wailsapp/wails</div>
                            <div class="text-gray">MIT</div>
                        </div>
                        <div class="flex justify-between">
                            <div>gin-gonic/gin</div>
                            <div class="text-gray">MIT</div>
                        </div>
                        <div class="flex justify-between">
                            <div>gorilla/websocket</div>
                            <div class="text-gray">BSD 3 Clause</div>
                        </div>
                        <div class="flex justify-between">
                            <div>go/pkg/sftp</div>
                            <div class="text-gray">BSD 2 Clause</div>
                        </div>
                        <div class="flex justify-between">
                            <div>go/x/crypto</div>
                            <div class="text-gray">BSD 3 Clause</div>
                        </div>
                        <div class="flex justify-between">
                            <div>modernc.org/sqlite</div>
                            <div class="text-gray">BSD 3 Clause</div>
                        </div>
                        <div class="flex justify-between">
                            <div>xorm.io/xorm</div>
                            <div class="text-gray">BSD 3 Clause</div>
                        </div>
                        <div class="flex justify-between">
                            <div>vuejs/core</div>
                            <div class="text-gray">MIT</div>
                        </div>
                        <div class="flex justify-between">
                            <div>xtermjs</div>
                            <div class="text-gray">MIT</div>
                        </div>
                        <div class="flex justify-between">
                            <div>tusen-ai/naive-ui</div>
                            <div class="text-gray">MIT</div>
                        </div>
                        <div class="flex justify-between">
                            <div>axios</div>
                            <div class="text-gray">MIT</div>
                        </div>
                        <div class="flex justify-between">
                            <div>vuejs/pinia</div>
                            <div class="text-gray">MIT</div>
                        </div>
                        <div class="text-center text-small text-gray mt-10">上为直接引用的主要组件</div>
                    </div>
                </div>
                <n-button class="full-width mt-10" strong secondary>检查更新</n-button>
                <div class="text-small text-gray mt-5 text-center">Copyright &copy; 2023 Skye Zhang</div>
            </div>
        </div>
    </div>
</template>
  
<script>
import { LogoDevFilled, HomeRound } from '@vicons/material'
export default {
    name: "Setting",
    components: { LogoDevFilled, HomeRound },
    data: () => ({
    }),
    methods: {

    },
    mounted() {

    }
};
</script>
<style scoped>
.seting-title {
    font-size: 18px;
}

.setting-box {
    margin: 0 auto;
    width: 823px;
}

.card {
    background-color: rgba(255, 255, 255, .08);
    border-radius: 8px;
}

.setting-info {
    width: 200px;
}

.info-logo {
    margin-top: 10px;
    width: 100px;
}

.info-name {
    margin-top: -5px;
    font-size: 16px;
}
</style>