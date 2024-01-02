<template>
  <n-drawer :show="app.settingDrawerVisible" display-directive="show" :width="330" @mask-click="app.closeSettingDrawer">
    <n-drawer-content :title="$t('layout.settingDrawer.title')" :native-scrollbar="false">
      <dark-mode />
      <layout-mode />
      <theme-color-select />
      <page-func />
      <page-view />
      <theme-config />
    </n-drawer-content>
  </n-drawer>
</template>

<script setup lang="ts">
import { useAppStore } from '@/store';
import { $t } from '@/locales';
import { DarkMode, LayoutMode, PageFunc, PageView, ThemeColorSelect, ThemeConfig } from './components';

defineOptions({ name: 'SettingDrawer' });

const app = useAppStore();

</script>

<style scoped></style>
