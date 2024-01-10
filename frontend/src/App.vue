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
import { useAppVueContext } from '@/context';

const theme = useThemeStore();

const { useProvide } = useAppVueContext();
const { isRouterAlive } = useProvide();

subscribeStore();
useGlobalEvents();



</script>

<style scoped></style>
