<template>
  <div class="overflow-hidden">
    <n-card :title="$t('routes.audit.logs')" :bordered="false" class="h-full rounded-8px shadow-sm">
      <div class="flex-col h-full">
        <n-space class="pb-12px" justify="end">
          <n-space>
            <n-input
              v-model:value="search.kw"
              :placeholder="$t('page.audit.logsSearchPlaceholder')"
              size="small"
              style="width: 300px"
            />
            <n-button size="small" type="primary" @click="getTableData()">
              <icon-mdi-search class="mr-4px text-16px" :class="{ 'animate-spin': loading }" />
              {{ $t('common.search') }}
            </n-button>
            <n-button size="small" type="warning" @click="getTableData(true)">
              <icon-mdi-refresh class="mr-4px text-16px" :class="{ 'animate-spin': loading }" />
              {{ $t('common.refreshTable') }}
            </n-button>
          </n-space>
        </n-space>
        <n-data-table
          v-model:checked-row-keys="selectedRows"
          :columns="columns"
          :data="tableData"
          :remote="true"
          :loading="loading"
          :pagination="pagination"
          :row-key="row => row.id"
          flex-height
          class="flex-1-hidden"
        />
      </div>
    </n-card>
  </div>
</template>
<script setup lang="tsx">
import { reactive, ref, onMounted } from 'vue';
import type { Ref } from 'vue';
import type { DataTableColumns, PaginationProps, DataTableRowKey } from 'naive-ui';
import { NButton, NSpace } from 'naive-ui';
import { PageSizes } from '@/constants';
import { useLoading } from '@/hooks';
import { fetchLogs } from '@/service/api/audit';
import { $t } from '@/locales';

const { loading, startLoading, endLoading } = useLoading(false);

const columns: Ref<DataTableColumns<ApiAudit.Audit>> = ref([
  {
    type: 'selection',
    align: 'center'
  },
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
]) as Ref<DataTableColumns<ApiAudit.Audit>>;
const selectedRows = ref<DataTableRowKey[]>([]);

const tableData = ref<ApiAudit.Audit[]>([]);

const search = ref({
  page: 1,
  pageSize: 10,
  kw: ''
});

const pagination: PaginationProps = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: PageSizes,
  onChange: (page: number) => {
    pagination.page = page;
    search.value.page = page;
    getTableData();
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.pageSize = pageSize;
    pagination.page = 1;
    search.value.page = pagination.page;
    search.value.pageSize = pageSize;
    getTableData();
  }
});

async function getTableData(reset: boolean = false) {
  if (reset) {
    search.value.kw = '';
  }
  startLoading();
  const res = await fetchLogs(search.value);
  if (res.error === null) {
    tableData.value = res.data?.list;
    pagination.itemCount = res.data?.total;
    endLoading();
  }
}

onMounted(() => {
  getTableData();
});
</script>
