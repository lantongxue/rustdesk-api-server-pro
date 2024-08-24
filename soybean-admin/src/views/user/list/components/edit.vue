<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { addUser, editUser } from '@/service/api/user_management';

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

const userStatusOptions = [
  {
    value: -1,
    label: $t('dataMap.user.statusLabel.unverified')
  },
  {
    value: 0,
    label: $t('dataMap.user.statusLabel.disabled')
  },
  {
    value: 1,
    label: $t('dataMap.user.statusLabel.normal')
  }
];

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
  'username' | 'password' | 'name' | 'email' | 'licensed_devices' | 'note' | 'is_admin' | 'status'
>;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    username: '',
    password: '',
    name: '',
    email: '',
    licensed_devices: 0,
    note: '',
    status: 1,
    is_admin: false
  };
}

type RuleKey = Extract<keyof Model, 'username' | 'password' | 'name' | 'email' | 'status'>;

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
  status: {
    required: true,
    message: $t('page.user.list.selectUserStatus')
  }
};

function handleInitModel() {
  Object.assign(model, createDefaultModel());

  if (props.operateType === 'edit' && props.rowData) {
    Object.assign(model, props.rowData);
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
        <NFormItem v-if="!(operateType === 'edit' && model.is_admin)" :label="$t('dataMap.user.status')" path="status">
          <NSelect v-model:value="model.status" :options="userStatusOptions" />
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
