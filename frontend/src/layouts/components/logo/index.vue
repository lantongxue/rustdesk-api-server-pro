<template>
  <div class="flex-center nowrap-hidden h-56px ">
    <eg-icon local-icon="logo" v-if="props.collapsed" class="text-32px " :class="{ 'mr-2':!props.collapsed}"/>
    <n-image v-else :src="logo" :preview-disabled="true" width="200"/>
  </div>
</template>
<script setup lang="ts">
import logo from '@/assets/images/rustdesk-logo.png'

const props = defineProps({
  // 侧边栏是否折叠
  collapsed: {
    type: Boolean,
  },
});

</script>

<style scoped lang="scss">
</style>