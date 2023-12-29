<template>
  <n-form ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
    <n-form-item path="userName">
      <n-input
        v-model:value="model.userName"
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
    <n-form-item path="imgCode">
      <n-input v-model:value="model.imgCode" :clearable="true" :placeholder="$t('page.login.common.codePlaceholder')" />
      <div class="pl-8px">
        <image-verify v-model:code="imgCode" />
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
import { reactive, ref } from 'vue';
import type { FormInst, FormRules } from 'naive-ui';
import { useAuthStore } from '@/store';
import { formRules } from '@/utils';
import { fetchCaptcha } from '@/service/api/auth';

const auth = useAuthStore();
const { login } = useAuthStore();

const formRef = ref<HTMLElement & FormInst>();

const model = reactive({
  userName: '',
  password: '',
  imgCode: ''
});

const imgCode = ref('');

const captcha = fetchCaptcha();
console.log(captcha);

const rules: FormRules = {
  password: formRules.pwd
};

async function handleSubmit() {
  await formRef.value?.validate();

  const { userName, password } = model;

  login(userName, password);
}
</script>

<style scoped></style>
