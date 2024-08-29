<script setup lang="ts">
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { UserIsAdminOptions, UserStatusOptions } from '@/constants/business';

defineOptions({
  name: 'UserSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

const emit = defineEmits<Emits>();

const { formRef } = useNaiveForm();

const model = defineModel<Api.UserManagement.UserSearchParams>('model', { required: true });

async function reset() {
  emit('reset');
}

async function search() {
  emit('search');
}
</script>

<template>
  <NCard :bordered="false" size="small" class="card-wrapper">
    <NCollapse>
      <NCollapseItem :title="$t('common.search')" name="user-search">
        <NForm ref="formRef" :model="model" label-placement="left" :label-width="80">
          <NGrid responsive="screen" item-responsive>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.user.username')" path="username">
              <NInput v-model:value="model.username" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.user.name')" path="name">
              <NInput v-model:value="model.name" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.user.email')" path="email">
              <NInput v-model:value="model.email" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.user.created_at')" path="created_at">
              <NDatePicker
                v-model:formatted-value="model.created_at"
                type="daterange"
                value-format="yyyy-MM-dd HH:mm:ss"
                clearable
              />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.user.is_admin')" path="admin_status">
              <NSelect v-model:value="model.admin_status" :options="UserIsAdminOptions" clearable />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.user.status')" path="status">
              <NSelect v-model:value="model.status" :options="UserStatusOptions" clearable />
            </NFormItemGi>
            <NFormItemGi span="24 m:12">
              <NSpace class="w-full" justify="end">
                <NButton @click="reset">
                  <template #icon>
                    <icon-ic-round-refresh class="text-icon" />
                  </template>
                  {{ $t('common.reset') }}
                </NButton>
                <NButton type="primary" ghost @click="search">
                  <template #icon>
                    <icon-ic-round-search class="text-icon" />
                  </template>
                  {{ $t('common.search') }}
                </NButton>
              </NSpace>
            </NFormItemGi>
          </NGrid>
        </NForm>
      </NCollapseItem>
    </NCollapse>
  </NCard>
</template>

<style scoped></style>
