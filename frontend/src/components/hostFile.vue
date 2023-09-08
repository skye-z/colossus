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
                        <n-button size="small" tertiary @click="createDirectory">
                            <template #icon>
                                <n-icon>
                                    <CreateNewFolderOutlined />
                                </n-icon>
                            </template>
                        </n-button>
                    </template>
                    创建目录
                </n-tooltip>
                <n-dropdown trigger="click" :options="upOptions" @select="upload">
                    <n-button size="small" tertiary>
                        <template #icon>
                            <n-icon>
                                <CloudUpload />
                            </n-icon>
                        </template>
                    </n-button>
                </n-dropdown>
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
        <n-modal :mask-closable="false" :close-on-esc="false" v-model:show="rename.show">
            <n-card style="width: 300px" :bordered="false" size="small" role="dialog" aria-modal="true">
                <div class="text-center mb-10 text-gray">重命名</div>
                <div class="flex align-center mb-10">
                    <div class="file-label">原名称</div>
                    <n-input v-model:value="rename.file.name" disabled />
                </div>
                <div class="flex align-center mb-10">
                    <div class="file-label">新名称</div>
                    <n-input v-model:value="rename.name" placeholder="请输入新的名称" @keyup.enter="submitRename" />
                </div>
                <div>
                    <n-button type="primary" class="float-right" @click="submitRename">提交重命名</n-button>
                    <n-button @click="() => { rename.show = false }">取消</n-button>
                </div>
            </n-card>
        </n-modal>
        <n-modal :mask-closable="false" :close-on-esc="false" v-model:show="directory.show">
            <n-card style="width: 300px" :bordered="false" size="small" role="dialog" aria-modal="true">
                <div class="text-center mb-10 text-gray">创建目录</div>
                <div class="flex align-center mb-10">
                    <div class="file-label">目录名</div>
                    <n-input v-model:value="directory.name" placeholder="支持递归创建子文件夹" @keyup.enter="submitDirectory" />
                </div>
                <div>
                    <n-button type="primary" class="float-right" @click="submitDirectory">提交创建</n-button>
                    <n-button @click="() => { directory.show = false }">取消</n-button>
                </div>
            </n-card>
        </n-modal>
        <n-drawer class="no-select no-drag" v-model:show="fileInfo.show" :width="350" placement="right">
            <div class="tips" v-if="fileInfo.list.length == 0">文件详情</div>
            <!-- 文件名 索引节号 文件大小 I/O块大小 文件占用的块数 块大小 硬链接数量 文件类型 所有者id 所有者 
                用户组id 用户组 权限位 权限位和文件类型 主要设备类型 次要设备类型 最后访问时间 最后修改时间 最后更改时间 -->
            <div v-else class="pa-10">
                <div class="pa-10 line1">{{ fileInfo.list[0] }}</div>
                <div class="flex align-center pa-10">
                    <div class="info-label">索引</div>
                    <div>{{ fileInfo.list[1] }}</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">大小</div>
                    <div>{{ fileInfo.list[2] }}</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">存储块</div>
                    <div>{{ fileInfo.list[3] }} KB ({{ fileInfo.list[4] }} x {{ fileInfo.list[5] }})</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">挂载点</div>
                    <div>{{ fileInfo.list[14] }}</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">设备号</div>
                    <div>{{ fileInfo.list[15] }} / 0x{{ fileInfo.list[16] }}</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">硬链接</div>
                    <div>{{ fileInfo.list[6] }} 个</div>
                </div>
                <!-- <div class="flex align-center pa-10">
                    <div class="info-label">文件类型</div>
                    <div>{{ fileInfo.list[7] }}</div>
                </div> -->
                <div class="flex align-center pa-10">
                    <div class="info-label">所有者</div>
                    <div>{{ fileInfo.list[9] }} / {{ fileInfo.list[8] }}</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">用户组</div>
                    <div>{{ fileInfo.list[11] }} / {{ fileInfo.list[10] }}</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">权限</div>
                    <div>{{ fileInfo.list[13] }} / {{ fileInfo.list[12] }}</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">文件创建时间</div>
                    <div v-if="fileInfo.list[17]"><n-time :time="fileInfo.list[17]" /></div>
                    <div v-else>-</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">最后访问时间</div>
                    <div v-if="fileInfo.list[18]"><n-time :time="fileInfo.list[18]" /></div>
                    <div v-else>-</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">数据修改时间</div>
                    <div v-if="fileInfo.list[19]"><n-time :time="fileInfo.list[19]" /></div>
                    <div v-else>-</div>
                </div>
                <div class="flex align-center pa-10">
                    <div class="info-label">状态更改时间</div>
                    <div v-if="fileInfo.list[20]"><n-time :time="fileInfo.list[20]" /></div>
                    <div v-else>-</div>
                </div>
            </div>
        </n-drawer>
    </div>
</template>
  
<script>
import { Check, EditCircle, CloudUpload, Download, Refresh, Eye, EyeOff } from '@vicons/tabler'
import { CreateNewFolderOutlined } from '@vicons/material'
import { file } from '../plugins/api';
import util from '../plugins/util';

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
        fileInfo: {
            show: false,
            list: []
        },
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
        },
        rename: {
            show: false,
            file: {},
            name: ''
        },
        directory: {
            show: false,
            name: "",
        },
        upOptions: [
            { title: "上传文件", key: "uploadFile" },
            { title: "上传文件夹", key: "uploadDirectory" }
        ]
    }),
    methods: {
        // 初始化文件管理器
        init(hostId) {
            if (this.id == 0) {
                this.id = parseInt(hostId)
                this.getHomePath()
            }
        },
        // 获取用户目录
        getHomePath() {
            this.loading = true
            file.getHome(this.id).then(res => {
                if (res.state) {
                    this.path = res.message
                } else {
                    window.$message.warning(res.message)
                    this.loading = false
                }
                this.buildPathList()
                this.getFileList()
            }).catch(err => {
                console.log(err)
                this.loading = false
                window.$message.warning('获取用户目录失败, 发生意料之外的错误')
            })
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
            this.path = util.buildFilePath(this.pathList, this.path.split('/'), index)
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
                            let data = util.analysisFile(res.data[i])
                            if (data != null) list.push(data)
                        }
                        this.fileList = list
                        this.loading = false
                    })
                } else {
                    window.$message.warning(res.message)
                    this.loading = false
                }
            }).catch(err => util.funcError(err, this.loading, '获取文件列表'))
        },
        // 获取文件详情
        getFileInfo(path) {
            this.fileInfo = {
                show: true,
                list: []
            }
            file.getInfo(this.id, path).then(res => {
                if (res.state) {
                    if (res.data[0]) res.data[0] = res.data[0].substring(res.data[0].lastIndexOf('/')+1)
                    if (res.data[2]) res.data[2] = util.bytes2Size(parseInt(res.data[2]))
                    if (res.data[17]) res.data[17] = parseInt(res.data[17]+'000')
                    if (res.data[18]) res.data[18] = parseInt(res.data[18]+'000')
                    if (res.data[19]) res.data[19] = parseInt(res.data[19]+'000')
                    if (res.data[20]) res.data[20] = parseInt(res.data[20]+'000')
                    this.fileInfo.list = res.data
                } else {
                    window.$message.warning(res.message)
                }
            }).catch(err => util.funcErrorNoLoading(err, '获取文件详情'))
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
            this.menu.show = false
            if (key === 'directDownload' || key === 'zipDownload') {
                let item = this.menu.file
                let directory = localStorage.getItem('download.directory')
                let tips = window.$message.loading("正在下载中...", { duration: (1000 * 60 * 30) });
                file.download(this.id, key.replace('Download', ''), directory, this.path, item.name).then(res => {
                    if (res.state) {
                        tips.type = 'success'
                        tips.content = '下载成功'
                    } else {
                        tips.type = 'warning'
                        tips.content = res.message ? res.message : "下载失败, 发生意料之外的错误"
                    }
                    setTimeout(() => {
                        tips.destroy()
                    }, 2000)
                }).catch(err => util.funcErrorDIYTips(err, tips, '文件下载'))
            } else if (key === 'rename') {
                this.rename = {
                    show: true,
                    file: this.menu.file,
                    name: ''
                }
            } else if (key === 'remove') {
                window.$dialog.warning({
                    title: "危险操作警告",
                    content: "你正在尝试删除‘" + this.menu.file.name + "’" + (this.menu.file.type == 2 || this.menu.file.type == 4 ? ', 删除目录将会一并删除目录下的所有文件' : '') + ", 此项操作不可逆, 确认要删除吗?",
                    negativeText: '确认删除',
                    positiveText: '取消',
                    onNegativeClick: () => {
                        file.remove(this.id, this.path, this.menu.file.name).then(res => {
                            if (res.state) {
                                window.$message.success("删除成功");
                                this.getFileList()
                            } else {
                                window.$message.warning(
                                    res.message ? res.message : "删除失败, 发生意料之外的错误"
                                );
                            }
                        }).catch(err => util.funcErrorNoLoading(err, '文件删除'))
                    }
                });
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
        // 关闭右键菜单难
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
                                    key: "blockDownload",
                                    disabled: true
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
        },
        // 提交重命名
        submitRename() {
            let item = this.rename.file
            let newName = this.rename.name
            if (item.name === newName) {
                window.$message.warning('新旧名称不能相同')
                return false
            }
            this.rename.show = false
            file.move(this.id, this.path + "/" + item.name, this.path + "/" + newName).then(res => {
                if (res.state) {
                    window.$message.success("重命名成功");
                    this.getFileList()
                } else
                    window.$message.warning(res.message ? res.message : "重命名失败, 发生意料之外的错误")
            }).catch(err => util.funcErrorNoLoading(err, '重命名'))
        },
        // 上传准备
        upload(key) {
            let tips = window.$message.loading("等待选取...", { duration: (1000 * 60 * 30) });
            setTimeout(() => {
                if (key === 'uploadFile') {
                    this.$goSelectFile("选择待上传的文件").then((res) => {
                        if (res) {
                            tips.content = "正在上传中..."
                            this.submitUpload(tips, "file", res)
                        } else tips.destroy()
                        console.log(res)
                    });
                } else {
                    this.$goSelectDirectory("选择待上传的文件目录").then((res) => {
                        if (res) {
                            tips.content = "正在上传中..."
                            this.submitUpload(tips, "directory", res)
                        } else tips.destroy()
                        console.log(res)
                    });
                }
            }, 300)
        },
        // 提交上传
        submitUpload(tips, type, path) {
            let name = path.substring(path.lastIndexOf('/') + 1)
            path = path.substring(0, path.lastIndexOf('/'))
            file.upload(this.id, type, path, this.path, name).then(res => {
                if (res.state) {
                    tips.type = 'success'
                    tips.content = '上传成功'
                    this.getFileList()
                } else {
                    tips.type = 'warning'
                    tips.content = res.message ? res.message : "上传失败, 发生意料之外的错误"
                }
                setTimeout(() => {
                    tips.destroy()
                }, 2000)
            }).catch(err => util.funcErrorDIYTips(err, tips, '文件上传'))
        },
        // 显示创建目录窗口
        createDirectory() {
            this.directory = {
                show: true,
                name: "",
            }
        },
        // 提交创建目录请求
        submitDirectory() {
            if (this.directory.name.indexOf('/') == 0) {
                window.$message.warning('不允许使用斜杠开头')
                return false
            }
            this.directory.show = false
            file.createDirectory(this.id, this.path + "/" + this.directory.name).then(res => {
                if (res.state) {
                    window.$message.success("目录创建成功");
                    this.getFileList()
                } else
                    window.$message.warning(res.message ? res.message : "目录创建失败, 发生意料之外的错误")
            }).catch(err => util.funcErrorNoLoading(err, '目录创建'))
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

.file-label {
    margin-right: 10px;
    text-align: right;
    min-width: 45px;
    width: 45px;
}

.tips {
    color: rgba(255, 255, 255, .2);
    justify-content: center;
    align-items: center;
    font-size: 24px;
    display: flex;
    height: 100vh;
    width: 100%;
}

.info-label {
    width: 100px;
}
</style>