<template>
  <div class="overflow-hidden">
    <n-card :title="$t('routes.user._value')" :bordered="false" class="h-full rounded-8px shadow-sm">
      <div class="flex-col h-full">
        <n-space class="pb-12px" justify="space-between">
          <n-space>
            <n-button type="primary" @click="handleAddTable">
              <icon-ic-round-plus class="mr-4px text-20px" />
              {{ $t('common.add') }}
            </n-button>
            <n-button type="error" @click="handleBatchDelete">
              <icon-ic-round-delete class="mr-4px text-20px" />
              {{ $t('common.delete') }}
            </n-button>
          </n-space>
          <n-space align="center" :size="18">
            <n-button size="small" type="primary" @click="getTableData">
              <icon-mdi-refresh class="mr-4px text-16px" :class="{ 'animate-spin': loading }" />
              {{ $t('common.refreshTable') }}
            </n-button>
          </n-space>
        </n-space>
        <n-data-table
          :columns="columns"
          :data="tableData"
          :remote="true"
          :loading="loading"
          :checked-row-keys="selectedRows"
          :pagination="pagination"
          :row-key="row => row.id"
          flex-height
          class="flex-1-hidden"
        />
        <user-edit-modal
          v-model:visible="visible"
          :type="modalType"
          :edit-data="editData"
          @update:refresh="getTableData()"
        />
      </div>
    </n-card>
  </div>
</template>
<script setup lang="tsx">
import { reactive, ref, onMounted } from 'vue';
import type { Ref } from 'vue';
import type { DataTableColumns, PaginationProps, DataTableRowKey } from 'naive-ui';
import { NButton, NPopconfirm, NSpace, NTag } from 'naive-ui';
import { PageSizes } from '@/constants';
import { useBoolean, useLoading } from '@/hooks';
import { fetchUserList, delUser } from '@/service/api/user';
import { $t } from '@/locales';
import { UserStatus } from '@/constants/business';
import UserEditModal from './components/edit.vue';
import type { ModalType } from './components/edit.vue';

const { loading, startLoading, endLoading } = useLoading(false);
const { bool: visible, setTrue: openModal } = useBoolean();

const columns: Ref<DataTableColumns<ApiUserManagement.User>> = ref([
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
      const tagTypes: any = {
        '0': '',
        '-1': 'error',
        '1': 'success'
      };

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
            {$t('common.yes')}
          </NTag>
        );
      }
      return <NTag bordered={false}>{$t('common.no')}</NTag>;
    }
  },
  {
    key: 'created_at',
    title: $t('dataMap.user.created_at'),
    align: 'center'
  },
  {
    key: 'actions',
    title: '操作',
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
]) as Ref<DataTableColumns<ApiUserManagement.User>>;
const selectedRows = ref<DataTableRowKey[]>([]);

const modalType = ref<ModalType>('add');

const editData = ref<ApiUserManagement.User | null>(null);

const tableData = ref<ApiUserManagement.User[]>([]);

const search = ref({
  page: 1,
  pageSize: 10
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

async function getTableData() {
  startLoading();
  const res = await fetchUserList(search.value);
  if (res.data?.total && res.data?.total > 0) {
    tableData.value = res.data?.list;
    pagination.itemCount = res.data?.total;
    endLoading();
  }
}

function handleAddTable() {
  modalType.value = 'add';
  openModal();
}

function handleEditTable(row: any) {
  editData.value = row;
  modalType.value = 'edit';
  openModal();
}

async function handleDeleteTable(row: any) {
  const res = await delUser({ ids: [row.id] });
  if (res.error === null) {
    const k = `backend.${res.message}` as I18nType.I18nKey;
    window.$message?.success($t(k));
    await getTableData();
  } else {
    window.$message?.error(res.error.msg);
  }
}

async function handleBatchDelete() {
  const res = await delUser({ ids: selectedRows.value });
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
