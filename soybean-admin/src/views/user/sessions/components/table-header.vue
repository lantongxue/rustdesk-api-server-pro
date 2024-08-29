<script setup lang="ts">
import { $t } from '@/locales';

defineOptions({
  name: 'TableHeader'
});

interface Props {
  itemAlign?: NaiveUI.Align;
  disabledKill?: boolean;
  loading?: boolean;
}

defineProps<Props>();

interface Emits {
  (e: 'kill'): void;
  (e: 'refresh'): void;
}

const emit = defineEmits<Emits>();

const columns = defineModel<NaiveUI.TableColumnCheck[]>('columns', {
  default: () => []
});

function batchKill() {
  emit('kill');
}

function refresh() {
  emit('refresh');
}
</script>

<template>
  <NSpace :align="itemAlign" wrap justify="end" class="lt-sm:w-200px">
    <slot name="prefix"></slot>
    <slot name="default">
      <NPopconfirm @positive-click="batchKill">
        <template #trigger>
          <NButton size="small" ghost type="error" :disabled="disabledKill">
            <template #icon>
              <icon-ic-round-delete class="text-icon" />
            </template>
            {{ $t('page.user.sessions.kill') }}
          </NButton>
        </template>
        {{ $t('common.confirmDelete') }}
      </NPopconfirm>
    </slot>
    <NButton size="small" @click="refresh">
      <template #icon>
        <icon-mdi-refresh class="text-icon" :class="{ 'animate-spin': loading }" />
      </template>
      {{ $t('common.refresh') }}
    </NButton>
    <TableColumnSetting v-model:columns="columns" />
    <slot name="suffix"></slot>
  </NSpace>
</template>

<style scoped></style>
