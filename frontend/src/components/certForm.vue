<template>
  <n-drawer class="no-select no-drag" v-model:show="form.show" :mask-closable="false" :close-on-esc="false" :width="350"
    placement="right">
    <div class="form-head pa-10 flex justify-between border-bottom">
      <div>{{ form.edit ? "编辑凭证" : "新建凭证" }}</div>
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
        <n-form-item label="凭证名" path="name">
          <n-input v-model:value="form.name" placeholder="便于记忆与识别的凭证名称" />
        </n-form-item>

        <n-form-item label="算法" path="type">
          <n-select v-model:value="form.type" placeholder="请选择" :options="options.type" style="width: 80px" />
          <n-select v-model:value="form.level" placeholder="请选择" :options="options.level" class="ml-10"
            style="width: 80px" />
          <n-button class="ml-10" :disabled="form.type != 'RSA'" @click="getKey" v-if="!form.edit">生成</n-button>
        </n-form-item>
        <n-form-item label="公钥" path="publicKey">
          <n-input v-model:value="form.publicKey" :rows="4" type="textarea" placeholder="请输入公钥" />
        </n-form-item>
        <n-form-item label="私钥" path="privateKey">
          <n-input v-model:value="form.privateKey" :rows="8" type="textarea" placeholder="请输入私钥" />
        </n-form-item>
        <n-form-item label="密码" path="passphrase">
          <n-input v-model:value="form.passphrase" placeholder="请输入PEM密码(如有)" />
        </n-form-item>
        <n-form-item label="有效期限" path="period">
          <n-date-picker v-model:value="form.period" type="date" placeholder="留空即长期有效" clearable style="width: 170px" />
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
import { cert } from "../plugins/api";
import { Base64 } from 'js-base64';

export default {
  name: "CertForm",
  components: { CloseRound },
  data: () => ({
    loading: false,
    form: {
      show: false,
      edit: false,

      name: null,
      type: "RSA",
      level: 2048,
      publicKey: null,
      privateKey: null,
      passphrase: null,
      period: null
    },
    options: {
      type: [
        {
          label: "RSA",
          value: "Linux",
        },
        {
          label: "EC",
          value: "MacOS",
        },
        {
          label: "DSA",
          value: "Windows",
        },
      ],
      level: [
        {
          label: "256",
          value: 256,
        },
        {
          label: "512",
          value: 512,
        },
        {
          label: "1024",
          value: 1024,
        },
        {
          label: "2048",
          value: 2048,
        },
        {
          label: "4096",
          value: 4096,
        },
        {
          label: "8192",
          value: 8192,
        }
      ],
    },
    rules: {
      name: {
        required: true,
        message: "凭证名称不能为空",
        trigger: "blur",
      },
      type: {
        required: true,
        message: "算法类型不能为空",
        trigger: "blur",
      },
      level: {
        required: true,
        message: "算法复杂度不能为空",
        trigger: "blur",
      },
      privateKey: {
        required: true,
        message: "私钥不能为空",
        trigger: "blur",
      },
    }
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
        type: "RSA",
        level: 2048,
        publicKey: null,
        privateKey: null,
        passphrase: null,
        period: null
      };
    },
    getKey() {
      this.loading = true;
      cert.initAdd(this.form.level).then((res) => {
        setTimeout(() => {
          this.loading = false;
          if (res.state) {
            window.$message.success("生成成功");
            this.form.type = res.data.type
            this.form.level = res.data.level
            this.form.publicKey = res.data.publicKey.trim()
            this.form.privateKey = res.data.privateKey.trim()
          } else {
            window.$message.warning(
              res.message ? res.message : "生成失败, 发生意料之外的错误"
            );
          }
        }, 500);
      })
        .catch((err) => {
          console.log(err);
          this.loading = false;
          window.$message.warning("生成失败, 发生意料之外的错误");
        });
    },
    submit() {
      this.$refs.formRef.validate((errors) => {
        if (errors) return false;
        this.loading = true;
        if (this.form.edit) this.updateCert();
        else this.addCert();
      });
    },
    addCert() {
      cert
        .add(this.form)
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
    updateCert() {
      cert
        .edit(this.form.id, this.form)
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
};
</script>
<style scoped>
.form-head {
  font-size: 18px;
}
</style>
