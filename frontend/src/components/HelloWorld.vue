<script setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <main>
    <div class="result">{{ data.resultText }}</div>
    <div class="input-box">
      <n-input v-model:value="data.name" type="text" placeholder="请输入名称" />
      <n-button type="primary" @click="greet">Greet</n-button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box {
  margin: 0 auto;
  display: flex;
  width: 300px;
}
</style>
