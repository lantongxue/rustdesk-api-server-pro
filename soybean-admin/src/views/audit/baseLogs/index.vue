<script setup lang="tsx">
import { NTag } from 'naive-ui';
import { fetchAuditLogList } from '@/service/api/audit';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import { AuditTypeOptions } from '@/constants/business';
import TableHeader from './components/table-header.vue';
import AuditBaseLogsSearch from './components/search.vue';

const appStore = useAppStore();

const auditTypes: any = {
  '0': '',
  '1': 'error',
  '2': 'success'
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
  apiFn: fetchAuditLogList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    conn_id: null,
    type: null,
    rustdesk_id: null,
    ip: null,
    session_id: null,
    uuid: null,
    created_at: null,
    closed_at: null
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
        for (const option of AuditTypeOptions) {
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
      key: 'conn_id',
      title: $t('dataMap.audit.conn_id'),
      align: 'center'
    },
    {
      key: 'rustdesk_id',
      title: $t('dataMap.audit.rustdesk_id'),
      align: 'center'
    },
    {
      key: 'ip',
      title: $t('dataMap.audit.ip'),
      align: 'center'
    },
    {
      key: 'session_id',
      title: $t('dataMap.audit.session_id'),
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
    },
    {
      key: 'closed_at',
      title: $t('dataMap.audit.closed_at'),
      align: 'center'
    }
  ]
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <AuditBaseLogsSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

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
