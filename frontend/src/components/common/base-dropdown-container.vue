<template>
  <n-dropdown v-if="options.length" :options="props.options" @select="handleDropdown" :key-field="props.keyField" @clickoutside="handleClickOutSide" :label-field="props.labelField" :size="size">
    <slot></slot>
  </n-dropdown>
  <slot v-else></slot>
</template>
<script setup lang="ts">
import {PropType} from "vue";
import {DropdownOption} from "naive-ui";
import {System} from "@/typings/system";

const props = defineProps({
  options: {
    type: Array as PropType<System.GlobalDropdown[]>,
    default: () => []
  },
  size: {
    type: String as PropType<'small' | 'medium' | 'large' | 'huge'>,
    default: 'medium'
  },
  keyField: {
    type: String,
    default: 'key'
  },
  labelField: {
    type: String,
    default: 'label'
  },
})
const emits = defineEmits(['select', 'clickOutside'])
const handleDropdown = (key: string | number, option: DropdownOption) => {
  emits('select', key, option)
}
const handleClickOutSide = (e: MouseEvent) => {
  emits('clickOutside',e)
}
</script>

<style scoped>

</style>