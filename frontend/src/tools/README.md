# Extensions Tool Development Guide

Thank you very much for developing extensions for Colossus, please Fork this project, and have the following information ready

```json
{
    "name": "Unique Filename",
    "title": "Tool Display Name",
    "profile": "Tool Profile",
    "author": "Developer",
    "version": "Version",
}
```

## Creating tool files

> Please note that no matter how complex your tool is, make sure there is only one `.vue` file.

You need to create the tool's interface file in the current directory, and name it according to the `${name}.vue` rule.

You can use `NaiveUI` components in the interface file, the content of the interface file can be seen in the following example

```html
<template>
    <div>
        <n-card class="full-screen" title="Example" header-style="text-align: center" :bordered="false" size="small">
            This is a full-screen card with a centered card title
        </n-card>
    </div>
</template>
  
<script>
// Provide interfaces to operate the host
import { host } from "../plugins/api"

export default {
    name: "Example",
    data: () => ({ }),
    methods: {
        // This function must be provided
        // The loader will call this function after a dynamic mount and pass in the host number
        init(id) { }
    }
};
</script>
<style scoped>

</style>
```

## Add tools to the directory

After you create a tool file, please add the tool information into the `directory.js` in the directory, the application loader will automatically mount it according to the directory.

> Do not modify other people's tool information, if you encounter `name` conflict, please try to use another `name`.

## Development and Debugging

Once the tool is added to the directory, you can connect to any host in the app, then switch to the `Tools` Tab, where all available tools will be displayed, click on your tool to mount and start it.

You can right-click on the interface to open the console for debugging.

## Submit PR

When you are done, please submit a PR to us, and your extension will be available to users in the next version after successful merge.