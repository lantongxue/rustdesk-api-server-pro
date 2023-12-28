<template>
  <BaseTable :columns="columns" :data="data"  :search-form-items="formFiledList" is-add-action title="基本列表" desc="基本列表秒杀"/>
</template>
<script setup lang="ts">

import {computed, reactive} from "vue";
import {ITableColumn} from "@/components/basic/table/index";
import BaseTable from "@/components/basic/table/index.vue";
import {BaseFormItemProps} from "@/components/basic/form/index";

const columns = reactive<ITableColumn[]>([
  {
    field: "name",
    title: "姓名",
    type: "string",
  },
  {
    field: "age",
    title: "年龄",
    type: "string",
  }
])
const formFiledList = reactive<BaseFormItemProps[]>([
  {
    label: '姓名',
    field: 'name',
    filedType: 'string',
    labelMessage: '这是一个提示',
    isSearch: true,
    filedOptions: {
      placeholder: '请输入姓名',
      required: true,
    },
  }, {
    label: '年龄',
    field: 'age',
    isSearch: true,
    filedType: 'string',
  }
])
const data = computed(()=>{
  return getTestTable()
})
function getTestTable(){
  // for 100次
  let data = []
  for (let i = 0; i < 100; i++) {
    data.push({
      name: '特产'+i,
      age: ""+i,
    })
  }
  return data
}
</script>

<style scoped>

</style>