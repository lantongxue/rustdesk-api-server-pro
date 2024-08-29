<script setup lang="tsx">
import { NButton, NPopconfirm, NSpace, NTag } from 'naive-ui';
import { delUser, fetchUserList } from '@/service/api/user_management';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { UserStatus } from '@/constants/business';
import UserEdit from './components/edit.vue';
import UserSearch from './components/search.vue';
const appStore = useAppStore();

const tagTypes: any = {
  '0': '',
  '-1': 'error',
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
  apiFn: fetchUserList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    username: null,
    name: null,
    email: null,
    created_at: null,
    admin_status: null,
    status: null
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
      key: 'name',
      title: $t('dataMap.user.name'),
      align: 'center'
    },
    {
      key: 'email',
      title: $t('dataMap.user.email'),
      align: 'center'
    },
    {
      key: 'licensed_devices',
      title: $t('dataMap.user.licensed_devices'),
      align: 'center'
    },
    {
      key: 'status',
      title: $t('dataMap.user.status'),
      align: 'center',
      render: row => {
        const label = UserStatus.get(row.status);
        return (
          <NTag bordered={false} type={tagTypes[row.status.toString()]}>
            {label}
          </NTag>
        );
      }
    },
    {
      key: 'is_admin',
      title: $t('dataMap.user.is_admin'),
      align: 'center',
      render: row => {
        if (row.is_admin) {
          return (
            <NTag bordered={false} type={'success'}>
              {$t('common.yesOrNo.yes')}
            </NTag>
          );
        }
        return <NTag bordered={false}>{$t('common.yesOrNo.no')}</NTag>;
      }
    },
    {
      key: 'created_at',
      title: $t('dataMap.user.created_at'),
      align: 'center'
    },
    {
      key: 'operate',
      title: $t('common.action'),
      align: 'center',
      render: row => {
        return (
          <NSpace justify={'center'}>
            <NButton size={'small'} onClick={() => handleEditTable(row)}>
              {$t('common.edit')}
            </NButton>
            {row.id === 1 ? null : (
              <NPopconfirm
                negativeText={$t('common.cancel')}
                positiveText={$t('common.confirm')}
                onPositiveClick={() => handleDeleteTable(row)}
              >
                {{
                  default: () => $t('common.confirmDelete'),
                  trigger: () => (
                    <NButton size={'small'} type={'error'}>
                      {$t('common.delete')}
                    </NButton>
                  )
                }}
              </NPopconfirm>
            )}
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
  handleAdd,
  handleEdit,
  checkedRowKeys,
  onBatchDeleted,
  onDeleted
  // closeDrawer
} = useTableOperate(data, getData);

async function handleBatchDelete() {
  const res = await delUser({ ids: checkedRowKeys.value });
  if (res.error === null) {
    onBatchDeleted();
  }
}

async function handleDeleteTable(row: Api.UserManagement.User) {
  const res = await delUser({ ids: [row.id] });
  if (res.error === null) {
    onDeleted();
  }
}

function handleEditTable(row: Api.UserManagement.User) {
  handleEdit(row.id);
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <UserSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />
    <NCard :title="$t('route.user_list')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :disabled-delete="checkedRowKeys.length === 0"
          :loading="loading"
          @add="handleAdd"
          @delete="handleBatchDelete"
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
      <UserEdit
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getDataByPage"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
