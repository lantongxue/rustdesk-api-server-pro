<template>
  <n-breadcrumb >
    <n-breadcrumb-item v-for="item in breadcrumbList" :key="item.name">
      <n-dropdown :options="item.children" v-if="item.children" @select="breadcrumbClick" key-field="name">
        <div class="flex-center">
          <component :is="item.icon" v-if="item.icon" class="mr-6px"/>
          {{ item.label }}
        </div>
      </n-dropdown>
      <div class="flex-center" v-else>
        <component :is="item.icon" v-if="item.icon" class="mr-6px"/>
        {{ item.label }}
      </div>
    </n-breadcrumb-item>
  </n-breadcrumb>
</template>
<script setup lang="ts">

// 生成面包屑数据
import {RouteLocationMatched, useRoute} from "vue-router";
import {System} from "@/typings/system";
import {renderIcon} from "@/utils";
import {computed} from "vue";
import {router} from "@/router";

const route = useRoute()
const generateBreadcrumbList = (matched: RouteLocationMatched[] | any) => {
  const breadcrumbList: System.GlobalBreadcrumb[] = [];
  if (matched && matched.length) {
    matched.forEach((item: RouteLocationMatched) => {
      const breadcrumb: System.GlobalBreadcrumb = {
        label: item?.meta?.title as any,
        name: item?.name as any,
        icon: item?.meta?.icon && renderIcon(item?.meta?.icon as any),
        disabled: item.name === route.name,
      };
      if (item?.children?.length && item?.children?.length > 1) {
        breadcrumb.children = generateBreadcrumbList(item.children);
      }
      breadcrumbList.push(breadcrumb);
    });
  }
  return breadcrumbList;
};
const breadcrumbList = computed(() => generateBreadcrumbList(route.matched))

const breadcrumbClick = (item:any) => {
  console.log(item)
  if (item !== route.name) {
    router.push({name: item})
  }
}
</script>

<style scoped>

</style>