<template>
  <n-space vertical>
    <n-form :rules="rules" ref="formRef" :model="formValue">
      <n-form-item-row label="手机号码" path="phone">
        <n-input placeholder="请输入手机号码" v-model:value="formValue.phone"/>
      </n-form-item-row>
      <n-form-item-row label="密码" path="password">
        <n-input placeholder="请输入密码" type="password" v-model:value="formValue.password" show-password-on="click"/>
      </n-form-item-row>
    </n-form>
    <n-button type="primary" block @click="handleSubmit" :disabled="loading" :loading="loading">登录</n-button>
  </n-space>
</template>
<script setup lang="ts">

import {reactive, ref} from "vue";
import {useAuthStore} from "@/store";
import {useRouter} from "vue-router";

const ua = useAuthStore()
const router = useRouter()
const rules = reactive({
  phone: {required: true, pattern: /^1[3456789]\d{9}$/, message: '请输入正确的手机号码', trigger: ['input', 'blur']},
  password: {required: true, message: '请输入密码', trigger: ['input', 'blur']}
})
const formRef = ref(null)
const loading = ref(false)
const formValue = reactive({
  phone: "18145000000",
  password: "123123"
})

const handleSubmit = async () => {
  loading.value = true
  let validate = false
  try {
    await formRef.value?.validate((errors) => {
      validate = !errors
    })
  } catch (error) {
    error.forEach((item) => {
      // @ts-ignore
      window.$message?.error(item[0].message)
    })
    loading.value = false
  }

  if (validate) {
    await ua.login()
    // @ts-ignore
    window.$message?.success('登录成功')
    // todo 实际登录
    setTimeout(() => {
      loading.value = false
      goHome()
    }, 1000)
  }
}

function goHome() {
  const {query} = router.currentRoute.value
  if (query?.redirect) {
    router.push({path: query.redirect as string})
  } else {
    router.push({name: 'home_index'})
  }
}

</script>

<style scoped>

</style>