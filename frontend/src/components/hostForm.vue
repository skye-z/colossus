<template>
    <n-drawer class="no-select no-drag" v-model:show="form.show" :mask-closable="false" :close-on-esc="false" :width="350"
        placement="right">
        <div class="form-head pa-10 flex justify-between border-bottom">
            <div>{{ form.edit ? '编辑主机' : '新建主机' }}</div>
            <n-button quaternary circle size="small" @click="close">
                <template #icon>
                    <n-icon><CloseRound /></n-icon>
                </template>
            </n-button>
        </div>
            <n-scrollbar style="max-height: calc(100vh - 105px)">
                <n-form ref="formRef" class="pa-20 pb-0" :model="form" :rules="rules" label-placement="left" label-width="auto"
                    require-mark-placement="right-hanging">
                    <n-form-item v-if="form.edit" label="编号" path="id">
                        <n-input v-model:value="form.id" disabled />
                    </n-form-item>
                    <n-form-item label="主机名" path="name">
                        <n-input v-model:value="form.name" placeholder="便于记忆与识别的主机名称" />
                    </n-form-item>
                    <n-form-item label="访问地址" path="address">
                        <n-input v-model:value="form.address" placeholder="主机网络地址(支持IP或域名)" />
                    </n-form-item>
                    <n-form-item label="连接端口" path="port">
                        <n-input v-model:value="form.port" placeholder="SSH连接端口(默认22)" style="width: 170px;" />
                    </n-form-item>
                    <n-form-item label="登录用户" path="user">
                        <n-input v-model:value="form.user" placeholder="登录主机的用户名称" style="width: 170px;" />
                    </n-form-item>
                    <n-form-item label="授权方式" path="type">
                        <n-radio-group v-model:value="form.type" name="type">
                            <n-space>
                                <n-radio value="cert">
                                    数字证书
                                </n-radio>
                                <n-radio value="pwd">
                                    账户密码
                                </n-radio>
                            </n-space>
                        </n-radio-group>
                    </n-form-item>
                    <n-form-item label="数字证书" path="cert" v-if="form.type === 'cert'">
                        <div class="full-width">
                            <n-select class="mb-10 full-width" v-model:value="form.certs" placeholder="选取已有证书"
                                :options="options.certs" />
                            <n-button class="mr-10">本地上传</n-button>
                            <n-button>生成新证书</n-button>
                        </div>
                    </n-form-item>
                    <n-form-item label="登录密码" path="secret" v-if="form.type === 'pwd'">
                        <n-input v-model:value="form.secret" type="password" show-password-on="mousedown"
                            placeholder="登录主机的用户密码" />
                    </n-form-item>

                    <n-form-item label="运行平台" path="platform">
                        <n-select v-model:value="form.platform" placeholder="连接成功后可自动获取" :options="options.platform"
                            style="width: 190px;" />
                    </n-form-item>
                    <n-form-item label="操作系统" path="system">
                        <n-select v-model:value="form.system" placeholder="连接成功后可自动获取" :options="options.system"
                            style="width: 190px;" />
                    </n-form-item>
                    <n-form-item label="所在地区" path="region">
                        <n-select v-model:value="form.region" placeholder="主机的网络地区" :options="options.region"
                            style="width: 150px;" />
                    </n-form-item>
                    <n-form-item label="主要用途" path="usage">
                        <n-select v-model:value="form.usage" placeholder="主机的使用场景" :options="options.usage"
                            style="width: 150px;" />
                    </n-form-item>
                    <n-form-item label="有效期限" path="period">
                        <n-date-picker v-model:value="form.period" type="date" placeholder="可使用的最后期限"
                            style="width: 170px;" />
                    </n-form-item>
                </n-form>
            </n-scrollbar>
        <div class="pa-10 flex justify-between border-top">
            <n-button @click="close">取消</n-button>
            <n-button type="primary">保存</n-button>
        </div>
    </n-drawer>
</template>
<script>
import { CloseRound } from '@vicons/material'

export default {
    name: "HostForm",
    components:{CloseRound},
    data: () => ({
        form: {
            show: false,
            edit: false,

            platform: null,
            system: null,
            region: null,
            usage: null,
            period: null,
            type: 'cert'
        },
        options: {
            platform: [
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
        },
        rules: {
            name: {
                required: true,
                message: "请输入主机名",
                trigger: "blur"
            },
            address: {
                required: true,
                message: "请输入主机访问地址",
                trigger: "blur"
            },
            port: {
                required: true,
                message: "请输入主机SSH连接端口",
                trigger: "blur"
            },
            user: {
                required: true,
                message: "请输入主机登录用户名",
                trigger: "blur"
            },
            cert: {
                required: true,
                message: "请选择数字证书",
                trigger: "blur"
            },
            secret: {
                required: true,
                message: "请输入主机登录用户密码",
                trigger: "blur"
            }
        },
        certs: []
    }),
    methods: {
        open(type, old) {
            this.form.show = true
            this.form.edit = type == 'edit'
        },
        close() {
            this.form.show = false
        }
    }
};
</script>
<style scoped>
.form-head {
    font-size: 18px;
}
</style>