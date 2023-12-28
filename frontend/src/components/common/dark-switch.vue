<template>
  <base-hover-tip text="主题切换" >
    <div class="flex text-24px cursor-pointer">
      <icon-line-md:sunny-filled-loop class="text-#FFCC00" v-if="theme.mode=='light'" @click="theme.setThemeMode('dark')"/>
      <icon-line-md:moon-filled-loop v-else @click="theme.setThemeMode('light')"/>
    </div>
  </base-hover-tip>
</template>
<script setup lang="ts">
import {useThemeStore} from "@/store";

const  theme = useThemeStore()
</script>

<style scoped>

</style>