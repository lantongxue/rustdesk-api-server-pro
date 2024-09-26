<script setup lang="ts">
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';

defineOptions({
  name: 'AuditSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

const emit = defineEmits<Emits>();

const { formRef } = useNaiveForm();

const model = defineModel<Api.Audit.AuditLogSearchParams>('model', { required: true });

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
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.audit.action')" path="action">
              <NInput v-model:value="model.action" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.audit.conn_id')" path="conn_id">
              <NInput v-model:value="model.conn_id" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.audit.rustdesk_id')" path="rustdesk_id">
              <NInput v-model:value="model.rustdesk_id" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.audit.ip')" path="ip">
              <NInput v-model:value="model.ip" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.audit.session_id')" path="session_id">
              <NInput v-model:value="model.session_id" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.audit.uuid')" path="uuid">
              <NInput v-model:value="model.uuid" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.user.created_at')" path="created_at">
              <NDatePicker
                v-model:formatted-value="model.created_at"
                type="daterange"
                value-format="yyyy-MM-dd HH:mm:ss"
                clearable
              />
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
