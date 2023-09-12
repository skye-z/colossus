<template>
  <div class="no-select">
    <div class="border-bottom flex align-center justify-between pa-10">
      <div class="seting-title">设置</div>
      <div class="text-gray" v-if="version.stage">{{ stage[version.stage] }}</div>
    </div>
    <div class="setting-box flex justify-between pa-10">
      <div class="full-width mr-10">
        <div class="card pa-10 mb-10">
          <div class="flex align-center">
            <div class="setting-label">终端背景</div>
            <n-color-picker
              v-model:value="terminalBgColor"
              style="width: 200px"
              @complete="updateValue('terminal.background_color')"
            />
          </div>
          <div class="flex align-center mt-10">
            <div class="setting-label">文字颜色</div>
            <n-color-picker
              v-model:value="terminalTextColor"
              style="width: 200px"
              @complete="updateValue('terminal.text_color')"
            />
          </div>
          <div class="flex align-center mt-10">
            <div class="setting-label">文字大小</div>
            <n-input-number
              v-model:value="terminalTextSize"
              :min="12"
              :max="32"
              @blur="updateValue('terminal.text_size')"
              style="width: 80px"
            />
            <div class="text-small text-gray ml-10">px</div>
          </div>
          <div class="text-small text-gray mt-10">
            * 修改终端样式后已连接的终端需重新连接后生效
          </div>
        </div>
        <div class="card pa-10 mb-10">
          <div class="flex align-center">
            <div class="setting-label">下载目录</div>
            <n-input-group>
              <n-input
                v-model:value="downloadPath"
                :disabled="downloadSelect"
                type="text"
                placeholder="请选择下载目录"
                @blur="updateValue('download.directory')"
              />
              <n-button type="primary" :loading="downloadSelect" @click="selectDirectory"
                >选取目录</n-button
              >
            </n-input-group>
          </div>
          <div class="flex align-center mt-10">
            <div class="setting-label">自动解压</div>
            <n-switch
              v-model:value="autoUnzip"
              @update:value="updateValue('download.auto_unzip')"
            />
            <div class="text-small text-gray ml-10">打包下载完成后自动解压</div>
          </div>
          <div class="flex align-center mt-10">
            <div class="setting-label">显示隐藏</div>
            <n-switch
              v-model:value="showHide"
              @update:value="updateValue('file.show_hide')"
            />
            <div class="text-small text-gray ml-10">开启后默认显示隐藏文件</div>
          </div>
        </div>
      </div>
      <div>
        <div class="card setting-info">
          <div class="text-center">
            <img class="info-logo" src="../assets/logo.png" alt="logo" />
            <div class="info-name">Colossus</div>
            <div class="border-bottom pa-10 pt-0">
              Version {{ version.code }} {{ version.stage }} {{ version.number }}
            </div>
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
            <div class="text-center text-small text-gray mt-10">
              上为直接引用的主要组件
            </div>
          </div>
        </div>
        <n-button class="full-width mt-10" strong secondary>检查更新</n-button>
        <div class="text-small text-gray mt-5 text-center">
          Copyright &copy; 2023 Skye Zhang
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { LogoDevFilled, HomeRound } from "@vicons/material";
import { config } from "../plugins/api";

export default {
  name: "Setting",
  components: { LogoDevFilled, HomeRound },
  data: () => ({
    version: {
      code: "",
      stage: "",
      number: "",
    },
    stage: {
      alpha: "内部测试版",
      beta: "公开测试版",
      rc: "最终测试版",
    },
    terminalBgColor: "",
    terminalTextColor: "",
    terminalTextSize: 14,
    downloadPath: "",
    downloadSelect: false,
    autoUnzip: false,
    showHide: true,
  }),
  methods: {
    getConfig() {
      config
        .getAll()
        .then((res) => {
          if (res.state) {
            for (let key in res.data) {
              localStorage.setItem(key, res.data[key]);
            }
            this.version = {
              code: res.data["basic.version"],
              stage: res.data["basic.stage"],
              number: res.data["basic.number"],
            };
            this.terminalBgColor = res.data["terminal.background_color"];
            this.terminalTextColor = res.data["terminal.text_color"];
            this.terminalTextSize = parseInt(res.data["terminal.text_size"]);
            this.downloadPath = res.data["download.directory"];
            this.autoUnzip = res.data["download.auto_unzip"] === "true";
            this.showHide = res.data["file.show_hide"] === "true";
          } else window.$message.warning("加载应用设置失败");
        })
        .catch(() => {
          window.$message.warning("加载应用设置失败");
        });
    },
    updateConfig(key, value) {
      localStorage.setItem(key, value);
      window.dispatchEvent(
        new CustomEvent("update:config", {
          detail: {
            key,
            value,
          },
        })
      );
      config
        .update(key, value)
        .then((res) => {
          if (res.state) window.$message.success("设置已自动保存");
          else window.$message.warning("自动保存失败");
        })
        .catch(() => {
          window.$message.warning("自动保存失败");
        });
    },
    selectDirectory() {
      this.downloadSelect = true;
      this.$goSelectDirectory("选择下载目录").then((res) => {
        if (res) this.downloadPath = res + "/colossus";
        this.downloadSelect = false;
        this.updateConfig("download.directory", this.downloadPath);
      });
    },
    updateValue(key) {
      if (key === "terminal.background_color")
        this.updateConfig(key, this.terminalBgColor);
      else if (key === "terminal.text_color")
        this.updateConfig(key, this.terminalTextColor);
      else if (key === "terminal.text_size")
        this.updateConfig(key, "" + this.terminalTextSize);
      else if (key === "download.directory") this.updateConfig(key, this.downloadPath);
      else if (key === "download.auto_unzip") this.updateConfig(key, "" + this.autoUnzip);
      else if (key === "file.show_hide") this.updateConfig(key, "" + this.showHide);
    },
  },
  mounted() {
    this.getConfig();
  },
};
</script>
<style scoped>
.seting-title {
  font-size: 18px;
  padding: 2.5px 0;
}

.setting-box {
  margin: 0 auto;
  width: 823px;
}

.card {
  background-color: rgba(255, 255, 255, 0.08);
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
  margin-top: -8px;
  font-size: 16px;
}

.setting-label {
  margin-right: 10px;
  text-align: right;
  min-width: 60px;
  width: 60px;
}
</style>
