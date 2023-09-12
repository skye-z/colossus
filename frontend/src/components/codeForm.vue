<template>
  <n-drawer class="no-select no-drag" v-model:show="form.show" :mask-closable="false" :close-on-esc="false" :width="350"
    placement="right">
    <div class="form-head pa-10 flex justify-between border-bottom">
      <div>{{ form.edit ? "编辑命令" : "新建命令" }}</div>
      <n-button quaternary circle size="small" @click="close" :disabled="loading">
        <template #icon>
          <n-icon>
            <CloseRound />
          </n-icon>
        </template>
      </n-button>
    </div>
    <n-scrollbar style="max-height: calc(100vh - 105px)">
      <n-form ref="formRef" class="pa-20 pb-0" :model="form" :rules="rules" :disabled="loading" label-placement="left"
        label-width="auto" require-mark-placement="right-hanging">
        <n-form-item v-if="form.edit" label="编号" path="id">
          <n-input v-model:value="form.id" disabled style="width: 110px" />
        </n-form-item>
        <n-form-item label="命令名" path="name">
          <n-input v-model:value="form.name" placeholder="便于记忆与识别的命令名称" />
        </n-form-item>

        <n-form-item label="运行平台" path="platform">
          <n-select v-model:value="form.platform" placeholder="请选择适配平台" :options="options.platform"
            style="width: 110px" />
        </n-form-item>
        <n-form-item label="操作系统" path="system">
          <n-select v-model:value="form.system" placeholder="请选择适配系统" :options="options.system" style="width: 190px" />
        </n-form-item>
        <n-form-item label="命令内容" path="content" label-placement="top">
          <n-input v-model:value="form.content" :rows="13" type="textarea" placeholder="请输入命令内容" />
        </n-form-item>
      </n-form>
    </n-scrollbar>
    <div class="pa-10 flex justify-between border-top">
      <n-button @click="close" :disabled="loading">取消</n-button>
      <n-popconfirm @positive-click="submit">
        <template #trigger>
          <n-button type="primary" :loading="loading">保存</n-button>
        </template>
        确认要保存吗?
      </n-popconfirm>
    </div>
  </n-drawer>
</template>
<script>
import { CloseRound } from "@vicons/material";
import { code } from "../plugins/api";
import { Base64 } from 'js-base64';
import options from "../plugins/options"

export default {
  name: "CodeForm",
  components: { CloseRound },
  data: () => ({
    loading: false,
    form: {
      show: false,
      edit: false,

      name: null,
      platform: null,
      system: null,
      content: null,
    },
    options: {},
    rules: {
      name: {
        required: true,
        message: "请输入命令名",
        trigger: "blur",
      },
      platform: {
        required: true,
        message: "请选择命令运行平台",
        trigger: "blur",
      },
      content: {
        required: true,
        message: "请输入命令内容",
        trigger: "blur",
      },
    },
    certs: [],
  }),
  methods: {
    open(type, old) {
      this.form.show = true;
      this.form.edit = type == "edit";

      for (let key in old) {
        if (typeof old[key] == "string" && old[key] != "") this.form[key] = old[key];
        else if (typeof old[key] == "number") this.form[key] = old[key];
      }

      this.form.content = Base64.decode(this.form.content)
    },
    close() {
      this.form = {
        show: false,
        edit: false,

        name: null,
        platform: null,
        system: null,
        content: null
      };
    },
    submit() {
      this.$refs.formRef.validate((errors) => {
        if (errors) return false;
        this.loading = true;
        let form = JSON.parse(JSON.stringify(this.form));
        form.content = Base64.encode(form.content)
        if (this.form.edit) this.updateCode(form);
        else this.addCode(form);
      });
    },
    addCode(form) {
      code
        .add(form)
        .then((res) => {
          setTimeout(() => {
            this.loading = false;
            if (res.state) {
              window.$message.success("保存成功");
              this.$emit("update");
              this.close();
            } else {
              window.$message.warning(
                res.message ? res.message : "保存失败, 发生意料之外的错误"
              );
            }
          }, 500);
        })
        .catch((err) => {
          console.log(err);
          this.loading = false;
          window.$message.warning("保存失败, 发生意料之外的错误");
        });
    },
    updateCode(form) {
      code
        .edit(form.id, form)
        .then((res) => {
          setTimeout(() => {
            this.loading = false;
            if (res.state) {
              window.$message.success("保存成功");
              this.$emit("update");
              this.close();
            } else {
              window.$message.warning(
                res.message ? res.message : "保存失败, 发生意料之外的错误"
              );
            }
          }, 500);
        })
        .catch((err) => {
          console.log(err);
          this.loading = false;
          window.$message.warning("保存失败, 发生意料之外的错误");
        });
    },
  },
  mounted() {
    this.options = options(["platform", "system", "region"])
  }
};
</script>
<style scoped>
.form-head {
  font-size: 18px;
}
</style>
