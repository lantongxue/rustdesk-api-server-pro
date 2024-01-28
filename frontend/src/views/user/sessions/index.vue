<template>
  <div class="overflow-hidden">
    <n-card :title="$t('routes.user.session')" :bordered="false" class="h-full rounded-8px shadow-sm">
      <div class="flex-col h-full">
        <n-space class="pb-12px" justify="space-between">
          <n-space>
            <n-button type="error" @click="handleBatchKill">
              <icon-ic-round-delete class="mr-4px text-20px" />
              {{ $t('page.session.kill') }}
            </n-button>
          </n-space>
          <n-space align="end" :size="18">
            <n-input v-model:value="search.kw" :placeholder="$t('dataMap.user.username')" size="small" />
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
import { NButton, NPopconfirm, NSpace } from 'naive-ui';
import { PageSizes } from '@/constants';
import { useLoading } from '@/hooks';
import { fetchSessionList, killSession } from '@/service/api/user';
import { $t } from '@/locales';

const { loading, startLoading, endLoading } = useLoading(false);

const columns: Ref<DataTableColumns<ApiUserManagement.Session>> = ref([
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
    title: $t('dataMap.user.username'),
    align: 'center'
  },
  {
    key: 'rustdesk_id',
    title: 'Rustdesk Id',
    align: 'center'
  },
  {
    key: 'expired',
    title: $t('dataMap.session.expired'),
    align: 'center'
  },
  {
    key: 'created_at',
    title: $t('dataMap.session.created_at'),
    align: 'center'
  },
  {
    key: 'actions',
    title: $t('common.action'),
    align: 'center',
    render: row => {
      return (
        <NSpace justify={'center'}>
          <NPopconfirm
            negativeText={$t('common.cancel')}
            positiveText={$t('common.confirm')}
            onPositiveClick={() => handleKill(row)}
          >
            {{
              default: () => $t('page.session.confirmKill'),
              trigger: () => (
                <NButton size={'small'} type={'error'}>
                  {$t('page.session.kill')}
                </NButton>
              )
            }}
          </NPopconfirm>
        </NSpace>
      );
    }
  }
]) as Ref<DataTableColumns<ApiUserManagement.Session>>;
const selectedRows = ref<DataTableRowKey[]>([]);

const tableData = ref<ApiUserManagement.Session[]>([]);

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
  const res = await fetchSessionList(search.value);
  if (res.error === null) {
    tableData.value = res.data?.list;
    pagination.itemCount = res.data?.total;
    endLoading();
  }
}

async function handleKill(row: any) {
  const res = await killSession({ ids: [row.id] });
  if (res.error === null) {
    const k = `backend.${res.message}` as I18nType.I18nKey;
    window.$message?.success($t(k));
    await getTableData();
  } else {
    window.$message?.error(res.error.msg);
  }
}

async function handleBatchKill() {
  const res = await killSession({ ids: selectedRows.value });
  if (res.error === null) {
    const k = `backend.${res.message}` as I18nType.I18nKey;
    window.$message?.success($t(k));
    await getTableData();
  } else {
    window.$message?.error(res.error.msg);
  }
}

onMounted(() => {
  getTableData();
});
</script>
