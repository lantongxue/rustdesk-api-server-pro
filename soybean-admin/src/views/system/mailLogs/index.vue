<script setup lang="tsx">
import { NButton, NSpace, NTag } from 'naive-ui';
import { fetchMailLogList } from '@/service/api/mail_logs';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { MailLogStatusOptions } from '@/constants/business';
import TableHeader from './components/table-header.vue';
import MailLogSearch from './components/search.vue';
import MailLogInfo from './components/info.vue';

const appStore = useAppStore();

const statusTypes: any = {
  1: 'success',
  2: 'error'
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
  apiFn: fetchMailLogList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    username: null,
    uuid: null,
    subject: null,
    to: null,
    from: null,
    status: null,
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
      title: $t('dataMap.mailLog.username'),
      align: 'center'
    },
    {
      key: 'uuid',
      title: $t('dataMap.mailLog.uuid'),
      align: 'center'
    },
    {
      key: 'from',
      title: $t('dataMap.mailLog.from'),
      align: 'center'
    },
    {
      key: 'to',
      title: $t('dataMap.mailLog.to'),
      align: 'center'
    },
    {
      key: 'subject',
      title: $t('dataMap.mailLog.subject'),
      align: 'center'
    },
    {
      key: 'status',
      title: $t('dataMap.mailLog.status'),
      align: 'center',
      render: row => {
        let label = '';
        for (const option of MailLogStatusOptions) {
          if (option.value === row.status) {
            label = option.label;
          }
        }
        return (
          <NTag bordered={false} type={statusTypes[row.status]}>
            {$t(label as App.I18n.I18nKey)}
          </NTag>
        );
      }
    },
    {
      key: 'created_at',
      title: $t('dataMap.mailLog.created_at'),
      align: 'center'
    },
    {
      key: 'operate',
      title: $t('common.action'),
      align: 'center',
      render: () => {
        return (
          <NSpace justify={'center'}>
            <NButton size={'small'} type={'info'} onClick={() => handleInfoTable()}>
              {$t('page.system.mailLog.info')}
            </NButton>
          </NSpace>
        );
      }
    }
  ]
});

const {
  drawerVisible,
  operateType,
  editingData,
  openDrawer
  // closeDrawer
} = useTableOperate(data, getData);

function handleInfoTable() {
  openDrawer();
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <MailLogSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

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
      <MailLogInfo
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getDataByPage"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
