<template>
  <n-config-provider
    :theme="theme.naiveTheme"
    :theme-overrides="theme.naiveThemeOverrides"
    :locale="zhCN"
    :date-locale="dateZhCN"
    class="h-full"
  >
    <naive-provider>
      <router-view v-if="isRouterAlive" />
    </naive-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import { dateZhCN, zhCN } from 'naive-ui';
import { subscribeStore, useThemeStore } from '@/store';
import { useGlobalEvents } from '@/composables';
import { nextTick, ref } from 'vue';
import { provide } from 'vue';

const theme = useThemeStore();

const isRouterAlive = ref(true)

subscribeStore();
useGlobalEvents();

const reload = () => {
  isRouterAlive.value = false;
  nextTick(() => {
    isRouterAlive.value = true;
  })
}

provide('reload', reload);
</script>

<style scoped></style>
