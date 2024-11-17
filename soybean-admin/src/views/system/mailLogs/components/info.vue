<script setup lang="ts">
import { ref, watch } from 'vue';
import { MailLogStatusOptions, translateOptions } from '@/constants/business';
import { $t } from '@/locales';
import { fetchMailLogInfo } from '@/service/api/mail_logs';

defineOptions({
  name: 'MailLogInfo'
});

interface Props {
  /** the edit row data */
  rowData: Api.System.MailLog;
}

const props = defineProps<Props>();

const visible = defineModel<boolean>('visible', {
  default: false
});

function closeDrawer() {
  visible.value = false;
}

async function init() {
  const a = await fetchMailLogInfo({ uuid: props.rowData.uuid });
  console.log(a);
}

watch(visible, () => {
  if (visible.value) {
    init();
  }
});
</script>

<template>
  <NModal v-model:show="visible" display-directive="show">
    <NCard :title="$t('page.system.mailLog.info')" class="n-card">
      <NForm>
        <NFormItem :label="$t('dataMap.mailLog.username')" path="username">
          <NInput v-model:value="rowData.username!" :disabled="true" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailLog.uuid')" path="uuid">
          <NInput v-model:value="rowData.uuid!" :disabled="true" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailLog.from')" path="from">
          <NInput v-model:value="rowData.from!" :disabled="true" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailLog.to')" path="to">
          <NInput v-model:value="rowData.to!" :disabled="true" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailLog.status')" path="status">
          <NSelect v-model:value="rowData.status!" :options="translateOptions(MailLogStatusOptions)" :disabled="true" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailLog.subject')" path="subject">
          <NInput v-model:value="rowData.subject!" :disabled="true" />
        </NFormItem>
      </NForm>
      <template #action>
        <NSpace :size="16" justify="end">
          <NButton @click="closeDrawer">{{ $t('common.cancel') }}</NButton>
        </NSpace>
      </template>
    </NCard>
  </NModal>
</template>

<style scoped>
.n-card {
  width: 800px;
}
</style>
