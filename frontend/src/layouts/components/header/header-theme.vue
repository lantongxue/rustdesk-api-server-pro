<template>
  <base-hover-tip text="主题设置" >
    <div class="flex text-24px cursor-pointer" @click="showSetting=!showSetting">
      <icon-line-md:sunny-filled-loop class="text-#FFCC00" v-if="theme.mode=='light'" />
      <icon-line-md:moon-filled-loop v-else/>
    </div>
  </base-hover-tip>
  <HeaderThemeSetting v-model:show="showSetting"/>
</template>
<script setup lang="ts">

import {useThemeStore} from "@/store";
import {ref} from "vue";
import HeaderThemeSetting from "@/layouts/components/header/header-theme-setting.vue";

const theme = useThemeStore()
const showSetting = ref(false)
</script>

<style scoped>

</style>