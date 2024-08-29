<script setup lang="tsx">
import { fetchAuditLogList } from '@/service/api/audit';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import TableHeader from './components/table-header.vue';
import AuditSearch from './components/search.vue';

const appStore = useAppStore();

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
    username: null,
    action: null,
    conn_id: null,
    rustdesk_id: null,
    ip: null,
    session_id: null,
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
      key: 'username',
      title: $t('dataMap.audit.username'),
      align: 'center'
    },
    {
      key: 'action',
      title: $t('dataMap.audit.action'),
      align: 'center'
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
    }
  ]
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <AuditSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

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
