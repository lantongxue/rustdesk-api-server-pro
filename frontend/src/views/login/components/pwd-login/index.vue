<template>
  <n-form ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
    <n-form-item path="username">
      <n-input
        v-model:value="model.username"
        :clearable="true"
        :placeholder="$t('page.login.common.userNamePlaceholder')"
      />
    </n-form-item>
    <n-form-item path="password">
      <n-input
        v-model:value="model.password"
        type="password"
        :clearable="true"
        show-password-on="click"
        :placeholder="$t('page.login.common.passwordPlaceholder')"
      />
    </n-form-item>
    <n-form-item path="code">
      <n-input v-model:value="model.code" :clearable="true" :placeholder="$t('page.login.common.codePlaceholder')" />
      <div class="pl-8px">
        <img width="152" height="40" class="cursor-pointer" :src="captcha?.data?.img" @click="handleCaptcha" />
      </div>
    </n-form-item>
    <n-space :vertical="true" :size="24">
      <n-button
        type="primary"
        size="large"
        :block="true"
        :round="true"
        :loading="auth.loginLoading"
        @click="handleSubmit"
      >
        {{ $t('page.login.common.confirm') }}
      </n-button>
    </n-space>
  </n-form>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue';
import type { FormInst, FormRules } from 'naive-ui';
import { useAuthStore } from '@/store';
import { createRequiredFormRule } from '@/utils';
import { fetchCaptcha } from '@/service/api/auth';

const auth = useAuthStore();
const { login } = useAuthStore();
const captcha = ref<Service.RequestResult<ApiAuth.Captcha>>();

const formRef = ref<HTMLElement & FormInst>();

const model = reactive({
  username: '',
  password: '',
  code: ''
});

async function handleCaptcha() {
  captcha.value = await fetchCaptcha();
}

onMounted(() => {
  handleCaptcha();
});

const rules: FormRules = {
  username: createRequiredFormRule('用户名不能为空'),
  password: createRequiredFormRule('密码不能为空'),
  code: createRequiredFormRule('验证码不能为空'),
  id: createRequiredFormRule('')
};

async function handleSubmit() {
  await formRef.value?.validate();
  await login({
    ...model,
    captchaId: captcha.value?.data?.id
  });
}
</script>

<style scoped></style>
