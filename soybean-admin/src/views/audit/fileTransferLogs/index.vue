<script setup lang="tsx">
import { NTag } from 'naive-ui';
import { fetchAuditFileTransferLogList } from '@/service/api/audit';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import { AuditFileTransferTypeOptions } from '@/constants/business';
import TableHeader from './components/table-header.vue';
import AuditFileTransferSearch from './components/search.vue';

const appStore = useAppStore();

const auditTypes: any = {
  '0': '',
  '1': 'success'
};

const {
  columns,
  columnChecks,
  data,
  getData,
  getDataByPage,
  loading,
  mobilePagination,
  searchParams,
  resetSearchParams
} = useTable({
  apiFn: fetchAuditFileTransferLogList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    type: null,
    rustdesk_id: null,
    peer_id: null,
    path: null,
    uuid: null,
    created_at: null
  },
  columns: () => [
    {
      key: 'id',
      title: 'ID',
      align: 'center'
    },
    {
      key: 'type',
      title: $t('dataMap.audit.type'),
      align: 'center',
      render: row => {
        let label = '';
        for (const option of AuditFileTransferTypeOptions) {
          if (option.value === row.type) {
            label = option.label;
          }
        }
        return (
          <NTag bordered={false} type={auditTypes[row.type]}>
            {$t(label as App.I18n.I18nKey)}
          </NTag>
        );
      }
    },
    {
      key: 'rustdesk_id',
      title: $t('dataMap.audit.rustdesk_id'),
      align: 'center'
    },
    {
      key: 'peer_id',
      title: $t('dataMap.audit.peer_id'),
      align: 'center'
    },
    {
      key: 'path',
      title: $t('dataMap.audit.path'),
      align: 'center'
    },
    {
      key: 'uuid',
      title: $t('dataMap.audit.uuid'),
      align: 'center'
    },
    {
      key: 'created_at',
      title: $t('dataMap.audit.created_at'),
      align: 'center'
    }
  ]
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <AuditFileTransferSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

    <NCard :title="$t('route.audit')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeader v-model:columns="columnChecks" :loading="loading" @refresh="getData" />
      </template>
      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="962"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
