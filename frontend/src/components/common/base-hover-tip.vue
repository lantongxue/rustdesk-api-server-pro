<template>
  <n-tooltip placement="bottom" trigger="hover">
    <template #trigger>
      <slot></slot>
    </template>
    <span>{{props.text}}</span>
  </n-tooltip>
</template>
<script setup lang="ts">
const props = defineProps({
  text: {
    type: String,
    default: ''
  },
})
</script>

<style scoped>

</style>