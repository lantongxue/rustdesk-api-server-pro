<template>
  <div :style="style" class="w-full flex-col " :class="{'bg-#fff dark:bg-dark':!props.transparent}">
    <slot></slot>
  </div>
</template>
<script setup lang="ts">
import {computed} from "vue";

/**
 * @description 通用容器组件
 *
 */
const props = defineProps({
  /** 是否透明 */
  transparent: {
    type: Boolean,
    default: false
  },
  /** 是否自动高度 true 自动计算 否100% */
  height: {
    type: String,
    default: 'calc(100vh - 108px)'
  },
})
const style = computed(() => {
  return {
    height: props.height
  }
})

</script>

<style scoped>

</style>