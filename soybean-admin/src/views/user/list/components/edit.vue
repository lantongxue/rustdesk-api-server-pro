<script setup lang="ts">
import { computed, reactive, ref, watch } from 'vue';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { addUser, editUser, getTOTP } from '@/service/api/user_management';
import { UserLoginVerifyOptions, UserStatusOptions, translateOptions } from '@/constants/business';

defineOptions({
  name: 'UserEdit'
});

interface Props {
  /** the type of operation */
  operateType: NaiveUI.TableOperateType;
  /** the edit row data */
  rowData?: Api.UserManagement.User | null;
}

const props = defineProps<Props>();

interface Emits {
  (e: 'submitted'): void;
}

const emit = defineEmits<Emits>();

const visible = defineModel<boolean>('visible', {
  default: false
});

const { formRef, validate, restoreValidation } = useNaiveForm();

const title = computed(() => {
  const titles: Record<NaiveUI.TableOperateType, string> = {
    add: $t('page.user.list.addUser'),
    edit: $t('page.user.list.editUser')
  };
  return titles[props.operateType];
});

type Model = Pick<
  Api.UserManagement.User,
  | 'id'
  | 'username'
  | 'password'
  | 'name'
  | 'email'
  | 'licensed_devices'
  | 'login_verify'
  | 'tfa_secret'
  | 'note'
  | 'is_admin'
  | 'admin_status'
  | 'status'
  | 'tfa_code'
>;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    id: 0,
    username: '',
    password: '',
    name: '',
    email: '',
    licensed_devices: 0,
    login_verify: 'access_token',
    tfa_secret: '',
    note: '',
    status: 1,
    is_admin: false,
    admin_status: 0,
    tfa_code: ''
  };
}

type RuleKey = Extract<
  keyof Model,
  'username' | 'password' | 'name' | 'email' | 'login_verify' | 'tfa_secret' | 'tfa_code' | 'status'
>;

const rules: Record<RuleKey, App.Global.FormRule> = {
  username: {
    required: true,
    message: $t('page.user.list.inputUsername')
  },
  password: {
    required: true,
    message: $t('page.user.list.inputPassword')
  },
  name: {
    required: true,
    message: $t('page.user.list.inputNickname')
  },
  email: {
    pattern: /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,
    message: $t('page.user.list.emailFormatError'),
    trigger: 'blur'
  },
  login_verify: {
    required: true,
    message: $t('page.user.list.selectUserStatus')
  },
  tfa_secret: {
    required: false,
    message: $t('page.user.list.require2FASecret')
  },
  tfa_code: {
    required: false,
    message: $t('page.user.list.require2FACode')
  },
  status: {
    required: true,
    message: $t('page.user.list.selectUserStatus')
  }
};

function handleInitModel() {
  Object.assign(model, createDefaultModel());

  if (props.operateType === 'edit' && props.rowData) {
    Object.assign(model, props.rowData);
    rules.password.required = false;
  }
}

function closeDrawer() {
  visible.value = false;
}

async function handleSubmit() {
  await validate();
  const method = props.operateType === 'add' ? addUser : editUser;
  const res = await method(model);

  if (res?.error === null) {
    const k = `api.${res.message}` as App.I18n.I18nKey;
    window.$message?.success($t(k));
    closeDrawer();
    emit('submitted');
  }
}

const tfaUrl = ref<string>();

async function handleLoginVerifyChanged(value: string) {
  if (value === 'tfa_check') {
    if (model.username === '') {
      window.$message?.error($t('page.user.list.inputUsername'));
      return;
    }
    const key = await getTOTP(model);
    tfaUrl.value = key.data.url;
    model.tfa_secret = key.data.key;
    rules.tfa_secret.required = true;
    rules.tfa_code.required = true;
  } else {
    rules.tfa_secret.required = false;
    rules.tfa_code.required = false;
  }
}

watch(visible, () => {
  if (visible.value) {
    handleInitModel();
    restoreValidation();
  }
});
</script>

<template>
  <NDrawer v-model:show="visible" display-directive="show" :width="360">
    <NDrawerContent :title="title" :native-scrollbar="false" closable>
      <NForm ref="formRef" :model="model" :rules="rules">
        <NFormItem :label="$t('dataMap.user.username')" path="username">
          <NInput v-model:value="model.username" :placeholder="$t('page.user.list.inputUsername')" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.user.password')" path="password">
          <NInput v-model:value="model.password" :placeholder="$t('page.user.list.inputPassword')" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.user.name')" path="name">
          <NInput v-model:value="model.name" :placeholder="$t('page.user.list.inputNickname')" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.user.email')" path="email">
          <NInput v-model:value="model.email" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.user.licensed_devices')" path="licensed_devices">
          <NInputNumber v-model:value="model.licensed_devices" />
        </NFormItem>
        <NFormItem
          v-if="!(operateType === 'edit' && model.is_admin)"
          :label="$t('dataMap.user.is_admin')"
          path="is_admin"
        >
          <NSwitch v-model:value="model.is_admin">
            <template #checked>{{ $t('common.yesOrNo.yes') }}</template>
            <template #unchecked>{{ $t('common.yesOrNo.no') }}</template>
          </NSwitch>
        </NFormItem>
        <NFormItem :label="$t('dataMap.user.login_verify')" path="login_verify">
          <NSelect
            v-model:value="model.login_verify"
            :options="translateOptions(UserLoginVerifyOptions)"
            @update-value="handleLoginVerifyChanged"
          />
        </NFormItem>
        <NFormItem
          v-if="model.login_verify === 'tfa_check'"
          :label="$t('page.user.list.tfa_secret_bind')"
          path="tfa_secret"
        >
          <NSpace vertical>
            <NButton v-if="model.tfa_secret !== ''" @click="() => handleLoginVerifyChanged('tfa_check')">
              重新绑定
            </NButton>
            <NQrCode :value="tfaUrl" :size="220" />
          </NSpace>
        </NFormItem>
        <NFormItem v-if="model.login_verify === 'tfa_check'" :label="$t('dataMap.user.tfa_secret')" path="tfa_secret">
          <NInput v-model:value="model.tfa_secret" type="textarea" readonly="true" />
        </NFormItem>
        <NFormItem v-if="model.login_verify === 'tfa_check'" :label="$t('dataMap.user.tfa_code')" path="tfa_code">
          <NInput v-model:value="model.tfa_code" />
        </NFormItem>
        <NFormItem v-if="!(operateType === 'edit' && model.is_admin)" :label="$t('dataMap.user.status')" path="status">
          <NSelect v-model:value="model.status" :options="translateOptions(UserStatusOptions)" />
        </NFormItem>
      </NForm>
      <template #footer>
        <NSpace :size="16">
          <NButton @click="closeDrawer">{{ $t('common.cancel') }}</NButton>
          <NButton type="primary" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
        </NSpace>
      </template>
    </NDrawerContent>
  </NDrawer>
</template>

<style scoped></style>
