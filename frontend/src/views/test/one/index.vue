<template>
  <n-space vertical>
    <n-card :bordered="false" title="测试1" class=" w-full ">
    </n-card>
    <n-card :bordered="false" title="" class=" w-full ">
      <BaseForm :items="formFiledList" :grid-props="{cols:2}" inline v-model:data="data"></BaseForm>

    </n-card>
    <n-card :bordered="false" title="" class=" w-full ">
      <BaseForm :items="formFiledList" :grid-props="{cols:1}" :data="data"></BaseForm>
    </n-card>
    <n-card>
    </n-card>
  </n-space>
</template>
<script setup lang="ts">
import BaseForm from "@/components/basic/form/index.vue"
import {reactive} from "vue";
import {BaseFormItemProps} from "@/components/basic/form/index";

const formFiledList = reactive<BaseFormItemProps[]>([
  {
    label: '姓名',
    field: 'name',
    filedType: 'string',
    labelMessage: '这是一个提示',
    filedOptions: {
      placeholder: '请输入姓名',
      required: true,
    },
  }, {
    label: '年龄',
    field: 'age',
    filedType: 'string',
  }
])
const data = reactive({
  name: '123',
  age: '123',
})
</script>

<style scoped>

</style>