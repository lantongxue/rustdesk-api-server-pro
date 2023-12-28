<template>
  <div class="h-full flex-center w-30px cursor-pointer">
    <icon-line-md-menu-unfold-right class="text-20px" v-if="props.collapsed" @click="emits('click',false)"/>
    <icon-line-md-menu-unfold-left class="text-20px" v-if="!props.collapsed" @click="emits('click',true)"/>
  </div>
</template>
<script setup lang="ts">

const props = defineProps({
  collapsed: Boolean
});
const emits = defineEmits(["click"])
</script>

<style scoped>

</style>