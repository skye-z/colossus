<template>
    <div class="no-drag no-select">
        <div class="mask" v-if="menu.show" @click="closeMenu"></div>
        <div class="loading flex align-center justify-center" v-if="loading">
            <n-spin size="small" />
        </div>
        <div class="file-tool pa-5 border-bottom flex align-center justify-between">
            <n-input v-if="edit" class="mr-5" size="small" v-model:value="path" :autofocus="edit" type="text"
                placeholder="请输入目录路径" @keyup.enter="switchEdit(0)" />
            <n-breadcrumb class="full-width" v-else>
                <n-breadcrumb-item v-for="(item, index) in pathList" :clickable="item != '...'" @click="jump(index)">{{ item
                }}</n-breadcrumb-item>
            </n-breadcrumb>
            <n-tooltip v-if="edit" trigger="hover" :delay="1000" :show-arrow="false">
                <template #trigger>
                    <n-button size="small" class="mr-5" tertiary @click="switchEdit(0)">
                        <template #icon>
                            <n-icon>
                                <Check />
                            </n-icon>
                        </template>
                    </n-button>
                </template>
                跳转路径
            </n-tooltip>
            <n-tooltip v-else trigger="hover" :delay="1000" :show-arrow="false">
                <template #trigger>
                    <n-button size="small" class="mr-5" tertiary @click="switchEdit(1)">
                        <template #icon>
                            <n-icon>
                                <EditCircle />
                            </n-icon>
                        </template>
                    </n-button>
                </template>
                编辑路径
            </n-tooltip>
            <n-button-group class="mr-5">
                <n-tooltip trigger="hover" :delay="1000" :show-arrow="false">
                    <template #trigger>
                        <n-button size="small" tertiary>
                            <template #icon>
                                <n-icon>
                                    <CreateNewFolderOutlined />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    创建目录
                </n-tooltip>
                <n-tooltip trigger="hover" :delay="1000" :show-arrow="false">
                    <template #trigger>
                        <n-button size="small" tertiary>
                            <template #icon>
                                <n-icon>
                                    <CloudUpload />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    上传文件
                </n-tooltip>
                <n-tooltip trigger="hover" :delay="1000" :show-arrow="false">
                    <template #trigger>
                        <n-button size="small" tertiary>
                            <template #icon>
                                <n-icon>
                                    <Download />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    下载管理
                </n-tooltip>
            </n-button-group>
            <n-button-group>
                <n-tooltip trigger="hover" :delay="1000" :show-arrow="false" v-if="hide">
                    <template #trigger>
                        <n-button size="small" tertiary @click="switchHide">
                            <template #icon>
                                <n-icon>
                                    <Eye />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    显示不可见文件
                </n-tooltip>
                <n-tooltip trigger="hover" :delay="1000" :show-arrow="false" v-else>
                    <template #trigger>
                        <n-button size="small" tertiary @click="switchHide">
                            <template #icon>
                                <n-icon>
                                    <EyeOff />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    隐藏不可见文件
                </n-tooltip>
                <n-button size="small" tertiary @click="getFileList">
                    <template #icon>
                        <n-icon>
                            <Refresh />
                        </n-icon>
                    </template>
                </n-button>
            </n-button-group>
        </div>
        <div class="file-list">
            <n-data-table :columns="columns" :data="fileList" size="small" max-height="calc(100vh - 127px)"
                :bordered="false" :row-props="touchRow" />
        </div>
        <n-dropdown placement="bottom-start" trigger="manual" :x="menu.x" :y="menu.y" :options="menu.options"
            :show="menu.show" @select="selectMenu" />
    </div>
</template>
  
<script>
import { Check, EditCircle, CloudUpload, Download, Refresh, Eye, EyeOff } from '@vicons/tabler'
import { CreateNewFolderOutlined } from '@vicons/material'
import { file } from '../plugins/api';

export default {
    name: "HostFile",
    components: { Check, EditCircle, CloudUpload, Download, Refresh, CreateNewFolderOutlined, Eye, EyeOff },
    data: () => ({
        id: 0,
        path: '~',
        hide: true,
        edit: false,
        pathList: [],
        fileList: [],
        fileInfo: {},
        loading: true,
        columns: [
            {
                title: '名称',
                key: 'name',
                ellipsis: {
                    tooltip: true
                }
            },
            {
                title: '类型',
                key: 'typeName',
                width: '50px',
                align: 'center'
            },
            {
                title: "大小",
                key: "sizeText",
                width: '120px'
            },
            {
                title: "权限",
                key: "power",
                width: '90px'
            },
            {
                title: "链接数",
                key: "link",
                width: '70px',
                align: 'center'
            },
            {
                title: "所有者",
                key: "owner",
                width: '120px'
            }
        ],
        menu: {
            x: 0,
            y: 0,
            show: false,
            file: {},
            options: []
        }
    }),
    methods: {
        // 初始化文件管理器
        init(hostId) {
            if (this.id == 0) {
                this.id = hostId
                this.buildPathList()
                this.getFileList()
            }
        },
        // 构建路径列表
        buildPathList() {
            let list = this.path.split('/')
            if (list.length > 6) {
                this.pathList = list.slice(list.length - 4)
                this.pathList.unshift("...")
                this.pathList.unshift(list[0])
            } else this.pathList = list
        },
        // 面包屑跳转
        jump(index) {
            let path = ''
            let list = this.path.split('/')
            if (list.length > 6) {
                for (let i = 0; i <= list.length - 5; i++) {
                    path += '/' + list[i]
                }
            }
            for (let i = (list.length > 6 ? 2 : 0); i <= index; i++) {
                path += '/' + this.pathList[i]
            }
            if (path.length > 1) {
                let check = path.substring(0, 2);
                if (check == '/~' || check == '//') path = path.substring(1)
            }

            this.path = path
            this.buildPathList()
            this.getFileList()
        },
        // 获取文件列表
        getFileList() {
            this.loading = true
            file.getList(this.id, this.path, this.hide).then(res => {
                if (res.state) {
                    setTimeout(() => {
                        let list = []
                        for (let i in res.data) {
                            this.analysis(list, res.data[i])
                        }
                        this.fileList = list
                        this.loading = false
                    })
                } else {
                    window.$message.warning(res.message)
                    this.loading = false
                }
            }).catch(err => {
                console.log(err)
                this.loading = false
                window.$message.warning('获取文件列表失败, 发生意料之外的错误')
            })
        },
        // 获取文件详情
        getFileInfo(path) {
            file.getInfo(this.id, path).then(res => {
                if (res.state) {
                    console.log(res)
                } else {
                    window.$message.warning(res.message)
                }
            }).catch(err => {
                console.log(err)
                window.$message.warning('获取文件详情失败, 发生意料之外的错误')
            })
        },
        // 解析文件数据
        analysis(list, data) {
            if (data.name == '.' || data.name == '..') return false;
            if (data.type == 1) data.typeName = '文件'
            else if (data.type == 2) data.typeName = '目录'
            else if (data.type == 3) data.typeName = '程序'
            else if (data.type == 4) data.typeName = '链接'

            if (data.name.indexOf('->') != -1) {
                data.type = 4
                data.typeName = '链接'
                data.path = data.name.substring(data.name.indexOf('->') + 2, data.name.length)
                data.name = data.name.substring(0, data.name.indexOf('->'))
            }

            data.owner = data.user + ' / ' + data.group
            if (data.type == 2 || data.type == 4) data.sizeText = '-'
            else data.sizeText = this.bytes2Size(data.size)
            list.push(data)
        },
        // 字节转文件大小
        bytes2Size(number) {
            const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
            if (number === 0) return '0 Byte';
            const i = parseInt(Math.floor(Math.log(number) / Math.log(1024)));
            return Math.round(number / Math.pow(1024, i), 2) + ' ' + sizes[i];
        },
        // 切换地址编辑模式
        switchEdit(state) {
            this.edit = state == 1
            if (state == 0) {
                this.buildPathList()
                this.getFileList()
            }
        },
        // 切换不可见文件的显示状态
        switchHide() {
            this.hide = !this.hide;
            this.getFileList();
        },
        // 点击右键菜单
        selectMenu(key) {
            console.log(key)
            this.menu.show = false
            if (key === 'directDownload') {
                let item = this.menu.file
                let directory = localStorage.getItem('download.directory')
                file.download(this.id, directory, this.path, item.name).then(res => {
                    console.log(res)
                }).catch(() => {
                    window.$message.warning('文件下载失败, 发生意料之外的错误')
                })
            } else if (key === 'copyPath') {
                let path = ''
                if (this.menu.file.type == 4) path = this.menu.file.path
                else path = this.menu.file.name
                this.$goSetClipboard(this.path + '/' + path)
                window.$message.success('路径已复制到剪切板')
            } else if (key === 'sendTerminal') {
                let path = ''
                if (this.menu.file.type == 2) path = this.menu.file.name
                else if (this.menu.file.type == 4) path = this.menu.file.path
                this.$emit('send', this.path + '/' + path)
            }
        },
        closeMenu() {
            this.menu.show = false
        },
        // 点击项目
        touchRow(row) {
            return {
                onClick: () => {
                    let path = this.path
                    if (path.substring(0, path.length - 1) != '/') path += '/';
                    if (row.type == 2 || row.type == 4) {
                        // 打开目录
                        path = path + (row.type == 2 ? row.name : row.path)
                        if (path.length > 1) {
                            let check = path.substring(0, 2);
                            if (check == '/~' || check == '//') path = path.substring(1)
                        }
                        this.path = path
                        this.buildPathList()
                        this.getFileList()
                    } else this.getFileInfo(path + row.name)
                },
                onContextmenu: e => {
                    e.preventDefault();
                    let options = [
                        {
                            title: "重命名",
                            key: "rename"
                        },
                        {
                            title: "删除",
                            key: "remove"
                        },
                        {
                            type: 'divider',
                            key: 'd1'
                        },
                        {
                            label: '工具',
                            key: 'tools',
                            children: [
                                {
                                    label: '暂无可用工具',
                                    key: 'none',
                                    disabled: true
                                },
                            ]
                        },
                        {
                            title: "复制路径",
                            key: "copyPath"
                        }
                    ]
                    if (row.type == 2 || row.type == 4) {
                        options.unshift({
                            title: "打包下载",
                            key: "zipDownload"
                        })
                        options.push({
                            title: "终端中查看",
                            key: "sendTerminal"
                        })
                    } else {
                        options.unshift({
                            title: "下载",
                            key: "download",
                            children: [
                                {
                                    title: "直接下载",
                                    key: "directDownload"
                                },
                                {
                                    title: "打包下载",
                                    key: "zipDownload"
                                },
                                {
                                    title: "分块下载",
                                    key: "blockDownload"
                                },
                            ]
                        })
                    }

                    this.menu.options = options
                    this.menu.x = e.clientX
                    this.menu.y = e.clientY
                    this.menu.show = true
                    this.menu.file = row
                }
            };
        }
    }
};
</script>
<style scoped>
.loading {
    height: calc(100% - 90px);
    top: 88px;
}

.mask {
    position: fixed;
    height: 100vh;
    width: 100vw;
    z-index: 99;
    left: 0;
    top: 0;
}

.file-tool:deep(.n-breadcrumb-item__separator) {
    margin: 0;
}

.file-tool:deep(.n-breadcrumb-item__link) {
    -webkit-box-orient: vertical;
    text-overflow: ellipsis;
    overflow: hidden;
    max-width: 200px;
}

.file-list:deep(.n-data-table-tr) {
    cursor: pointer;
}
</style>