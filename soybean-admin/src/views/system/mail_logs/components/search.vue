<script setup lang="ts">
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { MailLogStatusOptions, translateOptions } from '@/constants/business';

defineOptions({
  name: 'MailLogSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

const emit = defineEmits<Emits>();

const { formRef } = useNaiveForm();

const model = defineModel<Api.System.MailLogSearchParams>('model', { required: true });

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
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.mailLog.username')" path="username">
              <NInput v-model:value="model.username" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.mailLog.uuid')" path="uuid">
              <NInput v-model:value="model.uuid" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.mailLog.to')" path="to">
              <NInput v-model:value="model.to" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.mailLog.status')" path="status">
              <NSelect v-model:value="model.status" :options="translateOptions(MailLogStatusOptions)" clearable />
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
