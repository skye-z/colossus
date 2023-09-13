# 扩展工具开发指南

非常感谢您为 Colossus 开发扩展工具, 请 Fork 本项目, 然后准备好以下信息

```json
{
    "name": "唯一文件名",
    "title": "工具显示名",
    "profile": "工具简介",
    "author": "开发者",
    "version": "版本号"
}
```

## 创建工具文件

> 请注意, 无论你是多复杂的工具, 请保证有且仅有一个`.vue`文件

你需要在当前目录下创建工具的界面文件, 其命名应当遵循 `${name}.vue` 的规则

你可以在界面文件中使用`NaiveUI`的组件, 界面文件内容可参考下面的示例

```html
<template>
    <div>
        <n-card class="full-screen" title="示例工具" header-style="text-align: center" :bordered="false" size="small">
            这是一个全屏卡片, 带有居中的卡片标题
        </n-card>
    </div>
</template>
  
<script>
// 提供操作主机的相关接口
import { host } from "../plugins/api"

export default {
    name: "Example",
    data: () => ({
        id: 0
    }),
    methods: {
        // 必须提供此函数
        // 加载器会在动态挂载后调用此函数并传入主机编号
        init(id) {}
    }
};
</script>
<style scoped>

</style>
```

## 加入工具目录

当你创建工具文件后, 请将工具信息添加进目录下的`directory.js`中, 应用内的加载器会根据目录自动挂载

> 请勿修改他人的工具信息, 如遇到`name`冲突, 请尝试使用另外的`name`

## 开发与调试

当工具已加入目录, 你可以在应用中连接任意主机, 然后切换到`工具` Tab, 在这里会显示所有可用工具, 点击你的工具即可挂载启动

你可以在界面上右键打开控制台来进行调试

## 提交 PR

当你所有工作都完成后, 请向我们提交PR, 成功合并后你的扩展工具将会在下个版本中向用户推出