<script setup lang="tsx">
import { NButton, NPopconfirm, NSpace } from 'naive-ui';
import { fetchSessionList, killSession } from '@/service/api/user_management';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable, useTableOperate } from '@/hooks/common/table';
import SessionSearch from './components/search.vue';
import TableHeader from './components/table-header.vue';
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
  apiFn: fetchSessionList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    username: null,
    created_at: null
  },
  columns: () => [
    {
      type: 'selection',
      align: 'center',
      disabled: row => {
        return row.id === 1;
      }
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
      key: 'operate',
      title: $t('common.action'),
      align: 'center',
      render: row => {
        return (
          <NSpace justify={'center'}>
            <NPopconfirm
              negativeText={$t('common.cancel')}
              positiveText={$t('common.confirm')}
              onPositiveClick={() => handleKillTable(row)}
            >
              {{
                default: () => $t('page.user.sessions.confirmKill'),
                trigger: () => (
                  <NButton size={'small'} type={'error'}>
                    {$t('page.user.sessions.kill')}
                  </NButton>
                )
              }}
            </NPopconfirm>
          </NSpace>
        );
      }
    }
  ]
});

const { checkedRowKeys } = useTableOperate(data, getData);

async function handleBatchKill() {
  const res = await killSession({ ids: checkedRowKeys.value });
  if (res.error === null) {
    getDataByPage();
  }
}

async function handleKillTable(row: Api.UserManagement.Session) {
  const res = await killSession({ ids: [row.id] });
  if (res.error === null) {
    getDataByPage();
  }
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <SessionSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

    <NCard :title="$t('route.user_sessions')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeader
          v-model:columns="columnChecks"
          :disabled-kill="checkedRowKeys.length === 0"
          :loading="loading"
          @kill="handleBatchKill"
          @refresh="getData"
        />
      </template>
      <NDataTable
        v-model:checked-row-keys="checkedRowKeys"
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
