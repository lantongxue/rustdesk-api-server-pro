<script setup lang="ts">
import { computed, reactive, ref, watch } from 'vue';
import type { FormInst, FormRules } from 'naive-ui';
import type { SelectMixedOption } from 'naive-ui/es/select/src/interface';
import { createRequiredFormRule } from '@/utils';
import { $t } from '@/locales';
import { UserStatus } from '@/constants/business';
import { addUser, editUser } from '@/service/api/user';

export interface Props {
  /** 弹窗可见性 */
  visible: boolean;
  type?: 'add' | 'edit';
  /** 编辑的表格行数据 */
  editData?: ApiUserManagement.User | null;
}

export type ModalType = NonNullable<Props['type']>;

defineOptions({ name: 'UserEditModal' });

const props = withDefaults(defineProps<Props>(), {
  type: 'add',
  editData: null
});

interface Emits {
  (e: 'update:visible', visible: boolean): void;
  (e: 'update:refresh'): void;
}

const emit = defineEmits<Emits>();

const modalVisible = computed({
  get() {
    return props.visible;
  },
  set(visible) {
    emit('update:visible', visible);
  }
});
const closeModal = () => {
  modalVisible.value = false;
  handleUpdateFormModel(createDefaultFormModel());
};

const title = computed(() => {
  const titles: Record<ModalType, string> = {
    add: $t('page.users.addUser'),
    edit: $t('page.users.editUser')
  };
  return titles[props.type];
});

const userStatusOptions: SelectMixedOption[] = [];
UserStatus.forEach((value, key) => {
  userStatusOptions.push({
    label: value,
    value: key
  });
});

const formRef = ref<HTMLElement & FormInst>();

const formModel = reactive<ApiUserManagement.User>(createDefaultFormModel());

const REGEXP_EMAIL = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;

const rules: FormRules = {
  name: createRequiredFormRule($t('page.users.inputNickname')),
  email: [{ pattern: REGEXP_EMAIL, message: $t('page.users.emailFormatError'), trigger: 'blur' }],
  status: createRequiredFormRule($t('page.users.selectUserStatus'))
};

function updateRules() {
  if (props.type === 'add') {
    rules.username = createRequiredFormRule($t('page.users.inputUsername'));
    rules.password = createRequiredFormRule($t('page.users.inputPassword'));
  } else {
    delete rules.username;
    delete rules.password;
  }
}

function createDefaultFormModel(): ApiUserManagement.User {
  return {
    id: 0,
    username: '',
    password: '',
    name: '',
    email: '',
    licensed_devices: 0,
    status: 1,
    note: '',
    is_admin: false,
    created_at: ''
  };
}

function handleUpdateFormModel(model: ApiUserManagement.User) {
  Object.assign(formModel, model);
}

function handleUpdateFormModelByModalType() {
  const handlers: Record<ModalType, () => void> = {
    add: () => {
      const defaultFormModel = createDefaultFormModel();
      handleUpdateFormModel(defaultFormModel);
    },
    edit: () => {
      if (props.editData) {
        handleUpdateFormModel(props.editData);
      }
    }
  };

  handlers[props.type]();
}

async function handleSubmit() {
  await formRef.value?.validate();
  const requestMethod = props.type === 'add' ? addUser : editUser;
  const res = await requestMethod(formModel);
  if (res.error === null) {
    const k = `backend.${res.message}` as I18nType.I18nKey;
    window.$message?.success($t(k));
  } else {
    window.$message?.error(res.error.msg);
  }
  emit('update:refresh');
  closeModal();
}

watch(
  () => props.visible,
  newValue => {
    if (newValue) {
      updateRules();
      handleUpdateFormModelByModalType();
    }
  }
);
</script>

<template>
  <NModal v-model:show="modalVisible" preset="card" :title="title" class="w-700px">
    <NForm ref="formRef" label-placement="left" :label-width="80" :model="formModel" :rules="rules">
      <NGrid :cols="24" :x-gap="18">
        <NFormItemGridItem v-if="type == 'add'" :span="12" :label="$t('dataMap.user.username')" path="username">
          <NInput v-model:value="formModel.username" />
        </NFormItemGridItem>
        <NFormItemGridItem :span="12" :label="$t('dataMap.user.password')" path="password">
          <NInput v-model:value="formModel.password" type="password" clearable />
        </NFormItemGridItem>
        <NFormItemGridItem :span="12" :label="$t('dataMap.user.name')" path="name">
          <NInput v-model:value="formModel.name" clearable />
        </NFormItemGridItem>
        <NFormItemGridItem :span="12" :label="$t('dataMap.user.email')" path="email">
          <NInput v-model:value="formModel.email" />
        </NFormItemGridItem>
        <NFormItemGridItem :span="12" :label="$t('dataMap.user.licensed_devices')" path="licensed_devices">
          <NInputNumber v-model:value="formModel.licensed_devices" />
        </NFormItemGridItem>
        <NFormItemGridItem :span="12" :label="$t('dataMap.user.status')" path="status">
          <NSelect v-model:value="formModel.status" :options="userStatusOptions" />
        </NFormItemGridItem>
        <NFormItemGridItem :span="12" :label="$t('dataMap.user.is_admin')" path="is_admin">
          <NSwitch v-model:value="formModel.is_admin">
            <template #checked>{{ $t('common.yes') }}</template>
            <template #unchecked>{{ $t('common.no') }}</template>
          </NSwitch>
        </NFormItemGridItem>
      </NGrid>
      <NSpace class="w-full pt-16px" :size="24" justify="end">
        <NButton class="w-72px" @click="closeModal">{{ $t('common.close') }}</NButton>
        <NButton class="w-72px" type="primary" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </NForm>
  </NModal>
</template>

<style scoped></style>
