<template>
  <hover-container
    class="w-40px h-full"
    :tooltip-content="$t('layout.settingDrawer.title')"
    :inverted="theme.header.inverted"
    @click="app.toggleSettingDrawerVisible"
  >
    <icon-ant-design-setting-outlined class="text-20px" />
  </hover-container>
</template>

<script setup lang="ts">
import { useAppStore, useThemeStore } from '@/store';
import { $t } from '@/locales';

defineOptions({ name: 'SettingButton' });

const app = useAppStore();
const theme = useThemeStore();
</script>

<style scoped></style>
