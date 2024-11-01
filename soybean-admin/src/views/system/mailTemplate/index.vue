<script setup lang="tsx">
import { NButton, NSpace, NTag } from 'naive-ui';
import { fetchMailTemplateList } from '@/service/api/system';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { MailTemplateOptions } from '@/constants/business';
import MailTemplateEdit from './components/edit.vue';
import MailTemplateSearch from './components/search.vue';
import TableHeader from './components/table-header.vue';
const appStore = useAppStore();

const tagTypes: any = {
  1: 'success',
  2: 'warning',
  3: 'info'
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
  apiFn: fetchMailTemplateList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    name: null,
    type: null,
    subject: null,
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
      key: 'name',
      title: $t('dataMap.mailTemplate.name'),
      align: 'center'
    },
    {
      key: 'type',
      title: $t('dataMap.mailTemplate.type'),
      align: 'center',
      render: row => {
        let label = '';
        for (const option of MailTemplateOptions) {
          if (option.value === row.type) {
            label = option.label;
          }
        }
        return (
          <NTag bordered={false} type={tagTypes[row.type]}>
            {$t(label as App.I18n.I18nKey)}
          </NTag>
        );
      }
    },
    {
      key: 'subject',
      title: $t('dataMap.mailTemplate.subject'),
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
            <NButton size={'small'} type={'success'} onClick={() => handleEditTable(row)}>
              {$t('common.edit')}
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
  handleAdd,
  handleEdit,
  editingData,
  checkedRowKeys
  // closeDrawer
} = useTableOperate(data, getData);

function handleEditTable(row: Api.System.MailTemplate) {
  handleEdit(row.id);
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <MailTemplateSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />
    <NCard
      :title="$t('route.system_mailtemplate')"
      :bordered="false"
      size="small"
      class="sm:flex-1-hidden card-wrapper"
    >
      <template #header-extra>
        <TableHeader
          v-model:columns="columnChecks"
          :disabled-kill="checkedRowKeys.length === 0"
          :loading="loading"
          @add="handleAdd"
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
      <MailTemplateEdit
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getDataByPage"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
