<template>
  <n-space :vertical="true" :size="16">
    <data-card />
    <top-chart />
  </n-space>
</template>

<script lang="ts" setup>
import { DataCard, TopChart } from './components';
</script>

<style scoped></style>
