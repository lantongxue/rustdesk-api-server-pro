<template>
  <dark-mode-container class="flex-center h-full" :inverted="theme.footer.inverted">
    <span>
      Copyright ©2024
      <a href="https://github.com/lantongxue/rustdesk-api-server-pro" target="_blank">GitHub@lantongxue</a>
    </span>
  </dark-mode-container>
</template>

<script setup lang="ts">
import { useThemeStore } from '@/store';

defineOptions({ name: 'GlobalFooter' });

const theme = useThemeStore();
</script>

<style scoped></style>
